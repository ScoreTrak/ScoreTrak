package runner

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type Runner struct {
	db           *gorm.DB
	q            queue.WorkerQueue
	r            *util.Store
	dsync        time.Duration
	staticConfig config.StaticConfig
}

// refreshDsync retrieves current timestamp from the database, and ensures that it is within 2 seconds of range when compared to the host's clock.
func (d Runner) refreshDsync() error {
	var tm time.Time
	// query current timestamp
	res, err := d.db.Raw("SELECT current_timestamp;").Rows()
	if err != nil {
		panic(err)
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {
			log.Printf("unable to close rows: %v", err)
		}
	}(res)
	if res.Err() != nil {
		panic(res.Err())
	}
	for res.Next() {
		_ = res.Scan(&tm)
	}
	err = util.DatabaseOutOfSync(tm)
	if err != nil {
		return err
	}
	return nil
}

func NewRunner(db *gorm.DB, q queue.WorkerQueue, r *util.Store, staticConfig config.StaticConfig) *Runner {
	return &Runner{
		db: db, q: q, r: r, staticConfig: staticConfig,
	}
}


func (d *Runner) run(configLoop *time.Ticker, scoringLoop *time.Timer, cnf *config.DynamicConfig, lastRound *round.Round) (err error) {
	for {
	runnerSelect:
		select {
		case <-configLoop.C:
			// When config timer kicks in, we re-pull the new config.
			newConfig, err := d.r.Config.Get(context.Background())
			if err != nil {
				return err
			}
			// We then update contents of cnf, with contents of new config
			*cnf = *newConfig
			// Update last known round
			lastRound, err = d.r.Round.GetLastRound(context.Background())
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			err = d.refreshDsync()
			if err != nil {
				return err
			}
			// restart scoring loop based on lastRound
			scoringLoop.Stop()
			if lastRound == nil {
				// If no round exists, retry
				scoringLoop = time.NewTimer((time.Duration(d.staticConfig.DynamicConfigPullSeconds) * time.Second) / 2)
			} else {
				scoringLoop = time.NewTimer(d.durationUntilNextRound(lastRound, cnf.RoundDuration))
			}
			if err != nil {
				return err
			}
		case <-scoringLoop.C:
			// When scoring timer kicks in, stop the timer (so we can later assign a new timer)
			scoringLoop.Stop()
			// re-pull config
			dcc, err := d.r.Config.Get(context.Background())
			if err != nil {
				return err
			}
			*cnf = *dcc
			var rnd *round.Round
			// If we are allowed to score, aka competition is enabled
			if *cnf.Enabled {
				// Update last known round
				lastRound, err = d.r.Round.GetLastRound(context.Background())
				switch {
				case err != nil:
					if errors.Is(err, gorm.ErrRecordNotFound) {
						// If there are no known rounds, start with round 1
						rnd = &round.Round{ID: 1}
					} else {
						return err
					}
				case lastRound.Finish != nil:
					// if last known round is already finished (aka, currently no, new rounds are elapsing), start the new round with ID of old round incremented
					rnd = &round.Round{ID: lastRound.ID + 1}
				case time.Now().After(lastRound.Start.Add(time.Duration(cnf.RoundDuration) * time.Second).Add(config.MinRoundDuration)):
					// If last round did not finish, and some long time has passed since the start (Likely because master of the previous round has failed to score),
					// then we create a new transaction, that
					err = d.db.Transaction(func(tx *gorm.DB) error {
						delayedLastRound := &round.Round{}
						// gets the last known round
						err := tx.Last(delayedLastRound).Error
						if err != nil {
							return err
						}
						// compares it to the lastRound, and ensures that Finish is still nil (aka, at the time of the transaction, the round is still not complete),
						if lastRound.ID == delayedLastRound.ID && delayedLastRound.Finish == nil {
							now := time.Now()
							// finalizes the round that failed to score
							err = tx.Model(delayedLastRound).Updates(round.Round{Finish: &now, Note: "This round did not score!", Err: "this round is now skipped because it took to long to score. This could be due to too low round duration, dead master, or a time desync between master instances."}).Error
							if err != nil {
								return err
							}
						}
						return nil
					})
					if err != nil {
						return err
					}
					// and attempts to start a new round
					rnd = &round.Round{ID: lastRound.ID + 1}
				default:
					// else, if there is an elapsing round, but that round is not due to be completed yet, we just wait
					scoringLoop = time.NewTimer(config.MinRoundDuration)
					break runnerSelect
				}
				// After everything is figured out, we are reading to move to the next round
				// Create context that will attempt to finish the round a little than Start time + RoundDuration to account for networking delay.
				ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(cnf.RoundDuration)*time.Second*9/10)
				// Attempt to Store a new round
				err = d.r.Round.Store(ctx, rnd)
				if err != nil {
					// If attempt to store round failed
					var serr *pgconn.PgError
					ok := errors.As(err, &serr)
					// and the error tells us that some other master got to store the round before us
					if ok && serr.Code == "23505" {
						// we then retrieve that round stored by other master
						r, err := d.r.Round.GetByID(ctx, rnd.ID)
						if err != nil {
							cancelFunc()
							return err
						}
						*rnd = *r
						// and set it as our last round
					} else {
						// otherwise this error is likely due to networking issues, so we just exit the process
						cancelFunc()
						return err
					}
				} else {
					// If we succeeded with saving the round, we proceed with scoring it
					go d.Score(time.Duration(cnf.RoundDuration)*time.Second*9/10, rnd)
				}
				// At this point some new round should have been recorded in the database.
				// Retrieve its start time, and calculate durationUntilNextRound based on that start time
				lastRound, err = d.r.Round.GetLastRound(context.Background())
				if err != nil {
					cancelFunc()
					return err
				}
				scoringLoop = time.NewTimer(d.durationUntilNextRound(lastRound, cnf.RoundDuration))
				cancelFunc()
			}
		}
	}
}

func (d *Runner) MasterRunner() (err error) {
	err = d.refreshDsync()
	if err != nil {
		return err
	}
	var scoringLoop *time.Timer

	// Pull new config from database
	cnf, err := d.r.Config.Get(context.TODO())
	if err != nil {
		return err
	}

	lastRound, err := d.r.Round.GetLastRound(context.Background())
	switch {
	case lastRound != nil:
		// if there is a round stored in database, then
		scoringLoop = time.NewTimer(d.durationUntilNextRound(lastRound, cnf.RoundDuration))
	case err != nil && errors.Is(err, gorm.ErrRecordNotFound):
		// if no round exists, then try to score almost as soon as possible
		scoringLoop = time.NewTimer((time.Duration(d.staticConfig.DynamicConfigPullSeconds) * time.Second) / 2)
	case err != nil:
		// Some Other error, that is likely connection/database related
		return err
	}
	// Re-Pull config every DynamicConfigPullSeconds
	configLoop := time.NewTicker(time.Duration(d.staticConfig.DynamicConfigPullSeconds) * time.Second)

	return d.run(configLoop, scoringLoop, cnf, lastRound)
}

func (d *Runner) durationUntilNextRound(rnd *round.Round, roundDuration uint64) time.Duration {
	// Start time of current round + Round Duration - Current time on database
	dur := rnd.Start.Add(time.Duration(roundDuration) * time.Second).Sub(time.Now().Add(d.dsync))
	// if duration is small, then just return minimum duration
	if dur <= 1 {
		return 1
	}
	return dur
}

var ErrUnknownPanic = errors.New("unknown panic")
var ErrPanic = errors.New("panic")

func (d Runner) generateScoredServices(ctx context.Context, teams []*team.Team, hostGroup []*hostgroup.HostGroup, serviceGroups []*servicegroup.ServiceGroup, rnd *round.Round) []*queueing.ScoringData {
	var servicesToBeScored []*queueing.ScoringData
	for t := range teams {
		// Get Child Hosts for a given team
		err := d.db.WithContext(ctx).Model(&teams[t]).Association("Hosts").Find(&teams[t].Hosts)
		if err != nil {
			panic(err)
		}
		hosts := teams[t].Hosts
		for h := range hosts {
			// Get Child Services for a given Host
			err = d.db.WithContext(ctx).Model(&hosts[h]).Association("Services").Find(&hosts[h].Services)
			if err != nil {
				panic(err)
			}
			services := hosts[h].Services
			for s := range services {
				// Get Child Properties for a given Service
				err = d.db.WithContext(ctx).Model(&services[s]).Association("Properties").Find(&services[s].Properties)
				if err != nil {
					panic(err)
				}
				if !*teams[t].Pause {
					var validService bool
					if !*hosts[h].Pause {
						// Get all services, which parent objects are not Paused/Disabled
						if hosts[h].HostGroupID != nil {
							for _, hG := range hostGroup {
								if hG.ID == *hosts[h].HostGroupID && !*(hG.Pause) {
									validService = true
								}
							}
						} else {
							// and set them to be "valid"
							validService = true
						}
					}
					if validService {
						schedule := services[s].RoundUnits
						if services[s].RoundDelay != nil {
							schedule += *(services[s].RoundDelay)
						}
						if !*(services[s].Pause) && rnd.ID%schedule == 0 {
							// If the service is "Valid", and it is scheduled to be run in this round, then
							for _, servGroup := range serviceGroups {
								// Create corresponding ScoringData object, which is to be scoring a second before that rounds deadline.
								if services[s].ServiceGroupID == servGroup.ID && *(servGroup.Enabled) {
									sq := queueing.QService{ID: services[s].ID, Group: servGroup.Name, Name: services[s].Name}
									params := property.PropertiesToMap(services[s].Properties)
									de, _ := ctx.Deadline()
									sd := &queueing.ScoringData{
										Deadline:   de.Add(-time.Second),
										Host:       hosts[h].Address,
										Service:    sq,
										Properties: params,
										RoundID:    rnd.ID,
										MasterTime: time.Now(),
									}
									servicesToBeScored = append(servicesToBeScored, sd)
								}
							}
						}
					}
				}
			}
		}
	}
	return servicesToBeScored
}

func (d Runner) Score(timeout time.Duration, rnd *round.Round) {
	ctx, c := context.WithTimeout(context.Background(), timeout)
	defer c()
	var Note string
	defer func() {
		if x := recover(); x != nil {
			var err error
			switch x := x.(type) {
			case string:
				err = fmt.Errorf("%w: %s", ErrPanic, x)
			case error:
				err = x
			default:
				err = ErrUnknownPanic
			}

			log.Println(err)
			d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("A panic has occurred. Err:%s", err.Error()))
		}
	}()
	log.Printf("Running check for round %d", rnd.ID)
	// Get All teams
	teams, err := d.r.Team.GetAll(ctx)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, "No Teams Detected")
		return
	}
	// Get All Host groups, and Service groups
	hostGroup, _ := d.r.HostGroup.GetAll(ctx)
	serviceGroups, err := d.r.ServiceGroup.GetAll(ctx)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, "No Service Groups Detected")
		return
	}
	servicesToBeScored := d.generateScoredServices(ctx, teams, hostGroup, serviceGroups, rnd)

	// If no services are to be scored, then finalize the round
	if len(servicesToBeScored) == 0 {
		d.finalizeRound(ctx, rnd, Note, "No scorable services detected")
		return
	}
	var chks []*queueing.QCheck
	var nonCriticalErr error
	// Queue the services to be scored
	chks, nonCriticalErr, err = d.q.Send(servicesToBeScored)
	// If there is an error that is not critical, then append it to the log of that round
	if nonCriticalErr != nil {
		Note += nonCriticalErr.Error()
	}
	// If error is critical, terminate the round
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, err.Error())
		return
	}
	checks := d.storeChecks(teams, chks, servicesToBeScored, rnd)

	// Check if we are still on the last round.
	r, _ := d.r.Round.GetLastRound(ctx)
	if r == nil || r.ID != rnd.ID {
		d.finalizeRound(ctx, rnd, Note, "Error while saving checks. Err: a different round started before current round was able to finish. The scores will not be committed")
		return
	}

	// Attempt to store all of the checks performed.
	err = d.r.Check.Store(ctx, checks)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while saving checks. Err: %s", err.Error()))
		return
	}

	// Calculates total score for a given service using an SQL query.
	tp, err := d.r.Report.CountPassedPerService(ctx)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while generating report. Err: %s", err.Error()))
		return
	}

	simpleTeamReport := d.generateSimpleTeamReport(teams, hostGroup, serviceGroups, rnd, tp)

	// Generate new report, and upload it to db
	ch := report.NewReport()
	bt, err := json.Marshal(&report.SimpleReport{Teams: simpleTeamReport, Round: rnd.ID})
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while saving report. Err: %s", err.Error()))
		return
	}

	ch.Cache = string(bt)
	err = d.r.Report.Update(ctx, ch)
	if err != nil {
		if strings.Contains(err.Error(), "split failed while applying backpressure to") {
			Note += "\nPossible solution to error: decrease: gc.ttlseconds for database. "
		}
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while saving report. Err: %s", err.Error()))
		return
	}
	// Notify all of the listening clients that a new report was generated
	pubsub, err := queue.NewMasterStreamPubSub(config.GetQueueConfig())
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while notifying report update. Err: %s", err.Error()))
		return
	}
	pubsub.NotifyTopic(config.GetPubSubConfig().ChannelPrefix + "_report")
	d.finalizeRound(ctx, rnd, Note, "")
}

func (d Runner) generateSimpleTeamReport(teams []*team.Team, hostGroup []*hostgroup.HostGroup, serviceGroups []*servicegroup.ServiceGroup, rnd *round.Round, tp map[uuid.UUID]uint64) map[uuid.UUID]*report.SimpleTeam {
	// Create a minified cumulative report of all previous rounds
	simpleTeamReport := make(map[uuid.UUID]*report.SimpleTeam)
	{
		for _, t := range teams {
			st := &report.SimpleTeam{Name: t.Name, Pause: *t.Pause, Hide: *t.Hide}
			st.Hosts = make(map[uuid.UUID]*report.SimpleHost)
			for _, h := range t.Hosts {
				sh := report.SimpleHost{Address: h.Address, Hide: *h.Hide, Pause: *h.Pause}
				if h.HostGroupID != nil {
					for _, hG := range hostGroup {
						if hG.ID == *h.HostGroupID {
							sh.HostGroup = &report.SimpleHostGroup{Hide: *hG.Hide, Pause: *hG.Pause, ID: *h.HostGroupID, Name: hG.Name}
						}
					}
				}
				sh.Services = make(map[uuid.UUID]*report.SimpleService)
				for _, s := range h.Services {
					points := tp[s.ID] * *s.Weight
					params := map[string]*report.SimpleProperty{}
					for _, p := range s.Properties {
						params[p.Key] = &report.SimpleProperty{Value: *p.Value, Status: p.Status}
					}
					var simpSgr *report.SimpleServiceGroup
					for _, sG := range serviceGroups {
						if sG.ID == s.ServiceGroupID {
							simpSgr = &report.SimpleServiceGroup{ID: sG.ID, Name: sG.Name, Enabled: *sG.Enabled}
						}
					}

					sh.Services[s.ID] = &report.SimpleService{Name: s.Name, DisplayName: s.DisplayName, Hide: *s.Hide, Pause: *s.Pause, Points: points, Properties: params, PointsBoost: *s.PointsBoost, SimpleServiceGroup: simpSgr, Weight: *s.Weight}

					if len(s.Checks) != 0 {
						lastCheck := s.Checks[len(s.Checks)-1]
						if lastCheck.RoundID == rnd.ID {
							sh.Services[s.ID].Check = &report.SimpleCheck{
								Passed: *lastCheck.Passed,
								Log:    lastCheck.Log,
								Err:    lastCheck.Err,
							}
						}
					}
				}
				st.Hosts[h.ID] = &sh
			}
			simpleTeamReport[t.ID] = st
		}
	}
	return simpleTeamReport
}

func (d Runner) storeChecks(teams []*team.Team, chks []*queueing.QCheck, servicesToBeScored []*queueing.ScoringData, rnd *round.Round) []*check.Check {
	var checks []*check.Check
	for _, t := range teams {
		for _, h := range t.Hosts {
			for _, s := range h.Services {
				for i, c := range servicesToBeScored {
					if c.Service.ID == s.ID {
						if chks[i] == nil {
							// If check returned nil (could be due to failing worker, timeout, etc)
							fls := false
							s.Checks = append(s.Checks, &check.Check{Passed: &fls, Log: "", ServiceID: c.Service.ID, RoundID: rnd.ID, Err: "Unable to hear back from the worker that was responsible for performing the check. Make sure worker is able to connect back to scoretrak"})
						} else {
							s.Checks = append(s.Checks, &check.Check{Passed: &chks[i].Passed, Log: chks[i].Log, ServiceID: c.Service.ID, RoundID: rnd.ID, Err: chks[i].Err})
						}
						checks = append(checks, s.Checks[len(s.Checks)-1])
					}
				}
			}
		}
	}
	return checks
}

func (d Runner) finalizeRound(ctx context.Context, rnd *round.Round, note string, errStr string) {
	log.Printf("Note: %s\nError: %s\nRound: %v", note, errStr, rnd)
	now := time.Now().Add(d.dsync)
	rnd.Finish = &now
	rnd.Note = note
	rnd.Err = errStr
	err := d.r.Round.Update(ctx, rnd)
	if err != nil {
		log.Printf("Unable to update round %v", rnd)
	}
}
