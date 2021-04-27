package runner

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	"log"
	"math"
	"strings"
	"time"
)

type dRunner struct {
	db           *gorm.DB
	q            queue.WorkerQueue
	r            *util.Store
	dsync        time.Duration
	staticConfig config.StaticConfig
}

//refreshDsync retrieves current timestamp from the database, and ensures that it is within 2 seconds of range when compared to the host's clock.
func (d dRunner) refreshDsync() error {
	var tm time.Time
	//query current timestamp
	res, err := d.db.Raw("SELECT current_timestamp;").Rows()
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		_ = res.Scan(&tm)
	}
	d.dsync = -time.Since(tm)
	//ensure that drift between database is not larger than DatabaseMaxTimeDriftSeconds
	if float64(time.Duration(d.staticConfig.DatabaseMaxTimeDriftSeconds)*time.Second) < math.Abs(float64(d.dsync)) {
		return fmt.Errorf("time difference between master, and database is too large. Please synchronize time\n(The difference should not exceed 2 seconds)\nTime on database:%s\nTime on master:%s", tm.String(), time.Now())
	}
	return nil
}

func NewRunner(db *gorm.DB, q queue.WorkerQueue, r *util.Store) *dRunner {
	return &dRunner{
		db: db, q: q, r: r,
	}
}

func (d *dRunner) MasterRunner() (err error) {
	err = d.refreshDsync()
	if err != nil {
		return err
	}
	var scoringLoop *time.Ticker

	//Pull new config from database
	cnf, err := d.r.Config.Get(context.TODO())
	if err != nil {
		return err
	}

	lastRound, err := d.r.Round.GetLastRound(context.Background())
	if lastRound != nil {
		//if there is a round stored in database, then
		scoringLoop = time.NewTicker(d.durationUntilNextRound(lastRound, cnf.RoundDuration))
	} else if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		//if no round exists, then try to score almost as soon as possible
		scoringLoop = time.NewTicker(config.MinRoundDuration / 2)
	} else if err != nil {
		//Some Other error, that is likely connection/database related
		return err
	}
	//Re-Pull config every DynamicConfigPullSeconds
	configLoop := time.NewTicker(time.Duration(d.staticConfig.DynamicConfigPullSeconds) * time.Second)

	for {
		select {
		case <-configLoop.C:
			//When config timer kicks in, we re-pull the new config.
			newConfig, err := d.r.Config.Get(context.Background())
			if err != nil {
				return err
			}
			//We then update contents of cnf, with contents of new config
			*cnf = *newConfig
			//Update last known round
			lastRound, err = d.r.Round.GetLastRound(context.Background())
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}
			err = d.refreshDsync()
			if err != nil {
				return err
			}
			//restart scoring loop based on lastRound
			scoringLoop.Stop()
			if lastRound == nil {
				//If no round exists, retry
				scoringLoop = time.NewTicker(config.MinRoundDuration / 2)
			} else {
				scoringLoop = time.NewTicker(d.durationUntilNextRound(lastRound, cnf.RoundDuration))
			}
		case <-scoringLoop.C:
			//When scoring timer kicks in, stop the timer (so we can later assign a new timer)
			scoringLoop.Stop()
			//re-pull config
			dcc, err := d.r.Config.Get(context.Background())
			if err != nil {
				return err
			}
			*cnf = *dcc
			var rnd *round.Round
			//If we are allowed to score, aka competition is enabled
			if *cnf.Enabled {
				//Update last known round
				lastRound, err = d.r.Round.GetLastRound(context.Background())
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						//If there are no known rounds, start with round 1
						rnd = &round.Round{ID: 1}
					} else {
						return err
					}
				} else if lastRound.Finish != nil {
					//if last known round is already finished (aka, currently no, new rounds are elapsing), start the new round with ID of old round incremented
					rnd = &round.Round{ID: lastRound.ID + 1}
				} else if time.Now().After(lastRound.Start.Add(time.Duration(cnf.RoundDuration) * time.Second).Add(config.MinRoundDuration)) {
					//If last round did not finish, and some long time has passed since the start (Likely because master of the previous round has failed to score),
					//then we create a new transaction, that
					err = d.db.Transaction(func(tx *gorm.DB) error {
						delayedLastRound := &round.Round{}
						//gets the last known round
						err := tx.Last(delayedLastRound).Error
						if err != nil {
							return err
						}
						//compares it to the lastRound, and ensures that Finish is still nil (aka, at the time of the transaction, the round is still not complete),
						if lastRound.ID == delayedLastRound.ID && delayedLastRound.Finish == nil {
							now := time.Now()
							//finalizes the round that failed to score
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
					//and attempts to start a new round
					rnd = &round.Round{ID: lastRound.ID + 1}
				} else {
					//else, if there is an elapsing round, but that round is not due to be completed yet, we just wait
					scoringLoop = time.NewTicker(config.MinRoundDuration)
					break
				}
				//After everything is figured out, we are reading to move to the next round
				//Create context that will attempt to finish the round a little than Start time + RoundDuration to account for networking delay.
				ctx, _ := context.WithTimeout(context.Background(), time.Duration(cnf.RoundDuration)*time.Second*9/10)
				//Attempt to Store a new round
				err = d.r.Round.Store(ctx, rnd)
				if err != nil {
					//If attempt to store round failed
					serr, ok := err.(*pgconn.PgError)
					//and the error tells us that some other master got to store the round before us
					if ok && serr.Code == "23505" {
						//we then retreive that round stored by other master
						r, err := d.r.Round.GetByID(ctx, rnd.ID)
						if err != nil {
							return err
						}
						*rnd = *r
						//and set it as our last round
					} else {
						//otherwise this error is likely due to networking issues, so we just exit the process
						return err
					}
				} else {
					//If we succeeded with saving the round, we proceed with scoring it
					go d.Score(ctx, rnd)
				}
				//At this point some new round should have been recorded in the database.
				//Retrieve its start time, and start time, and calculate durationUntilNextRound based on that start time
				lastRound, err = d.r.Round.GetLastRound(context.Background())
				if err != nil {
					return err
				}
				scoringLoop = time.NewTicker(d.durationUntilNextRound(lastRound, cnf.RoundDuration))
			}
		}
	}
}

func (d *dRunner) durationUntilNextRound(rnd *round.Round, RoundDuration uint64) time.Duration {
	//Start time of current round + Round Duration - Current time on database
	dur := rnd.Start.Add(time.Duration(RoundDuration) * time.Second).Sub(time.Now().Add(d.dsync))
	//if duration is small, then just return minimum duration
	if dur <= 1 {
		return 1
	}
	return dur
}

func (d dRunner) Score(ctx context.Context, rnd *round.Round) {
	var Note string
	defer func() {
		if x := recover(); x != nil {
			var err error
			switch x := x.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}

			log.Println(err)
			d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("A panic has occured. Err:%s", err.Error()))
		}
	}()
	log.Printf("Running check for round %d", rnd.ID)
	//Get All teams
	teams, err := d.r.Team.GetAll(ctx)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, "No Teams Detected")
		return
	}
	//Get All Host groups, and Service groups
	hostGroup, _ := d.r.HostGroup.GetAll(ctx)
	serviceGroups, err := d.r.ServiceGroup.GetAll(ctx)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, "No Service Groups Detected")
		return
	}
	var servicesToBeScored []*queueing.ScoringData
	for _, t := range teams {
		//Get Child Hosts for a given team
		err = d.db.WithContext(ctx).Model(&t).Association("Hosts").Find(&t.Hosts)
		if err != nil {
			panic(err)
		}
		for _, h := range t.Hosts {
			//Get Child Services for a given Host
			err = d.db.WithContext(ctx).Model(&h).Association("Services").Find(&h.Services)
			if err != nil {
				panic(err)
			}
			for _, s := range h.Services {
				//Get Child Properties for a given Service
				err = d.db.WithContext(ctx).Model(&s).Association("Properties").Find(&s.Properties)
				if err != nil {
					panic(err)
				}
				if !*t.Pause {
					var validService bool
					if !*h.Pause {
						//Get all services, which parent objects are not Paused/Disabled
						if h.HostGroupID != nil {
							for _, hG := range hostGroup {
								if hG.ID == *h.HostGroupID && !*(hG.Pause) {
									validService = true
								}
							}
						} else {
							//and set them to be "valid"
							validService = true
						}
					}
					if validService {
						schedule := s.RoundUnits
						if s.RoundDelay != nil {
							schedule += *(s.RoundDelay)
						}
						if !*(s.Pause) && rnd.ID%schedule == 0 {
							//If the service is "Valid", and it is scheduled to be run in this round, then
							for _, servGroup := range serviceGroups {
								//Create corresponding ScoringData object, which is to be scoring a second before that rounds deadline.
								if s.ServiceGroupID == servGroup.ID && *(servGroup.Enabled) {
									sq := queueing.QService{ID: s.ID, Group: servGroup.Name, Name: s.Name}
									params := property.PropertiesToMap(s.Properties)
									de, _ := ctx.Deadline()
									sd := &queueing.ScoringData{
										Deadline:   de.Add(-time.Second),
										Host:       h.Address,
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

	//If no services are to be scored, then finalize the round
	if len(servicesToBeScored) == 0 {
		d.finalizeRound(ctx, rnd, Note, "No scorable services detected")
		return
	}
	var chks []*queueing.QCheck
	var nonCriticalErr error
	//Queue the services to be scored
	chks, nonCriticalErr, err = d.q.Send(servicesToBeScored)
	//If there is an error that is not critical, then append it to the log of that round
	if nonCriticalErr != nil {
		Note += nonCriticalErr.Error()
	}
	//If error is critical, terminate the round
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, err.Error())
		return
	}
	var checks []*check.Check
	for _, t := range teams {
		for _, h := range t.Hosts {
			for _, s := range h.Services {
				for i, c := range servicesToBeScored {
					if c.Service.ID == s.ID {
						if chks[i] == nil {
							//If check returned nil (could be due to failing worker, timeout, etc)
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

	//Check if we are still on the last round.
	r, _ := d.r.Round.GetLastRound(ctx)
	if r == nil || r.ID != rnd.ID {
		d.finalizeRound(ctx, rnd, Note, "Error while saving checks. Err: a different round started before current round was able to finish. The scores will not be committed")
		return
	}

	//Attempt to store all of the checks performed.
	err = d.r.Check.Store(ctx, checks)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while saving checks. Err: %s", err.Error()))
		return
	}

	//Calculates total score for a given service using an SQL query.
	tp, err := d.r.Report.CountPassedPerService(ctx)
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while generating report. Err: %s", err.Error()))
		return
	}

	//Create a minified cumulative report of all previous rounds
	simpTeams := make(map[uuid.UUID]*report.SimpleTeam)
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
			simpTeams[t.ID] = st
		}
	}
	//Generate new report, and upload it to db
	ch := report.NewReport()
	bt, err := json.Marshal(&report.SimpleReport{Teams: simpTeams, Round: rnd.ID})
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
	//Notify all of the listening clients that a new report was generated
	pubsub, err := queue.NewMasterStreamPubSub(config.GetQueueConfig())
	if err != nil {
		d.finalizeRound(ctx, rnd, Note, fmt.Sprintf("Error while notifying report update. Err: %s", err.Error()))
		return
	}
	pubsub.NotifyTopic(config.GetPubSubConfig().ChannelPrefix + "_report")
	d.finalizeRound(ctx, rnd, Note, "")
}

func (d dRunner) finalizeRound(ctx context.Context, rnd *round.Round, Note string, Error string) {
	log.Printf("Note: %s\nError: %s\nRound: %v", Note, Error, rnd)
	now := time.Now().Add(d.dsync)
	rnd.Finish = &now
	rnd.Note = Note
	rnd.Err = Error
	err := d.r.Round.Update(ctx, rnd)
	if err != nil {
		log.Printf("Unable to update round %v", rnd)
	}
}
