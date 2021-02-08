package run

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
	"sync"
	"time"
)

type dRunner struct {
	db *gorm.DB
	q  queue.WorkerQueue
	r  *util.Store
}

var dsync time.Duration
var mud sync.RWMutex

func (d dRunner) refreshDsync() error {
	mud.Lock()
	defer mud.Unlock()
	var tm time.Time
	res, err := d.db.Raw("SELECT current_timestamp;").Rows()
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		res.Scan(&tm)
	}
	dsync = -time.Since(tm)
	if float64(time.Second*2) < math.Abs(float64(dsync)) {
		return fmt.Errorf("time difference between master, and database is too large. Please synchronize time\n(The difference should not exceed 2 seconds)\nTime on database:%s\nTime on master:%s", tm.String(), time.Now())
	}
	return nil
}
func (d dRunner) getDsync() time.Duration {
	mud.RLock()
	defer mud.RUnlock()
	return dsync
}

func NewRunner(db *gorm.DB, q queue.WorkerQueue, r *util.Store) *dRunner {
	return &dRunner{
		db: db, q: q, r: r,
	}
}

func (d *dRunner) MasterRunner(cnf *config.DynamicConfig) (err error) {
	err = d.refreshDsync()
	if err != nil {
		return err
	}
	var scoringLoop *time.Ticker
	r, _ := d.r.Round.GetLastNonElapsingRound(context.TODO())
	if r != nil {
		scoringLoop = time.NewTicker(d.durationUntilNextRound(r, cnf.RoundDuration))
	} else {
		scoringLoop = time.NewTicker(time.Second)
	}
	configLoop := time.NewTicker(config.MinRoundDuration)

	for {
		select {
		case <-configLoop.C:
			dcc, _ := d.r.Config.Get(context.TODO())
			if !cnf.IsEqual(dcc) {
				*cnf = *dcc
			}
			r, _ := d.r.Round.GetLastRound(context.TODO())
			err = d.refreshDsync()
			if err != nil {
				log.Println(err)
			}
			scoringLoop.Stop()
			scoringLoop = time.NewTicker(d.durationUntilNextRound(r, cnf.RoundDuration))
		case <-scoringLoop.C:
			scoringLoop.Stop()
			dcc, _ := d.r.Config.Get(context.TODO())
			if !cnf.IsEqual(dcc) {
				*cnf = *dcc
			}
			var rnd *round.Round
			if *cnf.Enabled {
				rNoneElapsing, _ := d.r.Round.GetLastNonElapsingRound(context.TODO())
				if rNoneElapsing != nil {
					rnd = &round.Round{ID: rNoneElapsing.ID + 1}
				} else {
					rnd = &round.Round{ID: 1}
				}
				lastRnd, _ := d.r.Round.GetLastRound(context.TODO())
				if lastRnd != nil && rnd.ID <= lastRnd.ID && lastRnd.Finish == nil {
					if time.Now().After(lastRnd.Start.Add(time.Duration(cnf.RoundDuration) * time.Second).Add(time.Second * 3)) {
						d.r.Round.Delete(context.TODO(), lastRnd.ID)
					} //This is indicative of a scoring master dying. In this case we delete the last round
				}
				ctx, _ := context.WithTimeout(context.Background(), (time.Duration(cnf.RoundDuration)*time.Second*9/10)+d.getDsync())
				d.attemptToScore(ctx, rnd)
				time.Sleep(time.Second * 2)
				scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd, cnf.RoundDuration))
			} else {
				scoringLoop = time.NewTicker(config.MinRoundDuration)
			}
		}
	}
}

func (d *dRunner) durationUntilNextRound(rnd *round.Round, RoundDuration uint64) time.Duration {
	if rnd == nil || rnd.ID == 0 {
		return config.MinRoundDuration / 2
	}
	dur := rnd.Start.Add(time.Duration(RoundDuration) * time.Second).Sub(time.Now().Add(d.getDsync()))
	if dur <= 1 {
		return 1
	}
	return dur
}

//Runs check in the background as a gorutine.
func (d *dRunner) attemptToScore(ctx context.Context, rnd *round.Round) {
	err := d.r.Round.Store(ctx, rnd)
	if err != nil {
		serr, ok := err.(*pgconn.PgError)
		if ok && serr.Code == "23505" {
			r, _ := d.r.Round.GetByID(ctx, rnd.ID)
			if r != nil {
				*rnd = *r
			}
		} else {
			log.Println(err)
			panic(err)
		}
	} else {
		go d.Score(ctx, *rnd)
	}
}

func (d dRunner) Score(ctx context.Context, rnd round.Round) {
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
			d.finalizeRound(ctx, &rnd, Note, fmt.Sprintf("A panic has occured. Err:%s", err.Error()))
		}
	}()
	log.Printf("Running check for round %d", rnd.ID)
	teams, err := d.r.Team.GetAll(ctx)
	if err != nil {
		d.finalizeRound(ctx, &rnd, Note, "No Teams Detected")
		return
	}
	hostGroup, _ := d.r.HostGroup.GetAll(ctx)
	serviceGroups, err := d.r.ServiceGroup.GetAll(ctx)
	if err != nil {
		d.finalizeRound(ctx, &rnd, Note, "No Service Groups Detected")
		return
	}
	var sds []*queueing.ScoringData
	for _, t := range teams {
		err = d.db.WithContext(ctx).Model(&t).Association("Hosts").Find(&t.Hosts)
		if err != nil {
			panic(err)
		}
		for _, h := range t.Hosts {
			err = d.db.WithContext(ctx).Model(&h).Association("Services").Find(&h.Services)
			if err != nil {
				panic(err)
			}
			for _, s := range h.Services {
				err = d.db.WithContext(ctx).Model(&s).Association("Properties").Find(&s.Properties)
				if err != nil {
					panic(err)
				}
				if *t.Enabled {
					var validHost bool
					if *h.Enabled {
						if h.HostGroupID != nil {
							for _, hG := range hostGroup {
								if hG.ID == *h.HostGroupID && *(hG.Enabled) {
									validHost = true
								}
							}
						} else {
							validHost = true
						}
					}
					if validHost {
						schedule := s.RoundUnits
						if s.RoundDelay != nil {
							schedule += *(s.RoundDelay)
						}
						if *(s.Enabled) && rnd.ID%schedule == 0 { //Todo: Fix: Add unscheduled checks into report. (Looking into report:30 To perhaps utilize Pause to indicate skipped service) This way scoreboard does not flicked when check is skipped (Or handle this on frontend by locally caching. This should also fix an issue with changing colors on ranking page
							for _, servGroup := range serviceGroups {
								if s.ServiceGroupID == servGroup.ID && *(servGroup.Enabled) {
									sq := queueing.QService{ID: s.ID, Group: servGroup.Name, Name: s.Name}
									params := PropertyToMap(s.Properties)
									de, _ := ctx.Deadline()
									sd := &queueing.ScoringData{
										Deadline:   de.Add(-time.Second),
										Host:       h.Address,
										Service:    sq,
										Properties: params,
										RoundID:    rnd.ID,
									}
									sds = append(sds, sd)
								}
							}
						}
					}
				}
			}
		}
	}
	if len(sds) == 0 {
		d.finalizeRound(ctx, &rnd, Note, "No scorable services detected")
		return
	}
	var chks []*queueing.QCheck
	var bearableErr error

	chks, bearableErr, err = d.q.Send(sds)
	if bearableErr != nil {
		Note += bearableErr.Error()
	}
	if err != nil {
		d.finalizeRound(ctx, &rnd, Note, err.Error())
		return
	}
	var checks []*check.Check
	for _, t := range teams {
		for _, h := range t.Hosts {
			for _, s := range h.Services {
				for i, c := range sds {
					if c.Service.ID == s.ID {
						if chks[i] == nil {
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

	err = d.r.Check.Store(ctx, checks)
	if err != nil {
		d.finalizeRound(ctx, &rnd, Note, fmt.Sprintf("Error while saving checks. Err: %s", err.Error()))
		return
	}
	r, _ := d.r.Round.GetLastRound(ctx)
	if r == nil || r.ID != rnd.ID {
		d.finalizeRound(ctx, &rnd, Note, "Error while saving checks. Err: a different round started before current round was able to finish. The scores will not be committed")
		return
	}

	tp, err := d.r.Report.CountPassedPerService(ctx)
	if err != nil {
		d.finalizeRound(ctx, &rnd, Note, fmt.Sprintf("Error while generating report. Err: %s", err.Error()))
		return
	}

	simpTeams := make(map[uuid.UUID]*report.SimpleTeam)
	{
		for _, t := range teams {
			st := &report.SimpleTeam{Name: t.Name, Enabled: *t.Enabled, Hidden: *t.Hidden}
			st.Hosts = make(map[uuid.UUID]*report.SimpleHost)
			for _, h := range t.Hosts {
				sh := report.SimpleHost{Address: h.Address, Enabled: *h.Enabled}
				if h.HostGroupID != nil {
					for _, hG := range hostGroup {
						if hG.ID == *h.HostGroupID {
							sh.HostGroup = &report.SimpleHostGroup{Enabled: *hG.Enabled, ID: *h.HostGroupID, Name: hG.Name}
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
					if len(s.Checks) != 0 {
						lastCheck := s.Checks[len(s.Checks)-1]
						sh.Services[s.ID] = &report.SimpleService{Name: s.Name, DisplayName: s.DisplayName, Enabled: *s.Enabled, Points: points, Properties: params, PointsBoost: *s.PointsBoost, SimpleServiceGroup: simpSgr, Weight: *s.Weight}
						if lastCheck.RoundID == rnd.ID {
							sh.Services[s.ID].Passed = *lastCheck.Passed
							sh.Services[s.ID].Log = lastCheck.Log
							sh.Services[s.ID].Err = lastCheck.Err
						} else {
							sh.Services[s.ID].Passed = false
							sh.Services[s.ID].Log = "Service was not checked because it was disabled"
							sh.Services[s.ID].Err = ""
						}
					}
				}
				st.Hosts[h.ID] = &sh
			}
			simpTeams[t.ID] = st
		}
	}

	ch := report.NewReport()
	bt, err := json.Marshal(&report.SimpleReport{Teams: simpTeams, Round: rnd.ID})
	if err != nil {
		d.finalizeRound(ctx, &rnd, Note, fmt.Sprintf("Error while saving report. Err: %s", err.Error()))
		return
	}
	ch.Cache = string(bt)
	err = d.r.Report.Update(ctx, ch)
	if err != nil {
		if strings.Contains(err.Error(), "split failed while applying backpressure to") {
			Note += "\nPossible solution to error: decrease: gc.ttlseconds for database. "
		}
		d.finalizeRound(ctx, &rnd, Note, fmt.Sprintf("Error while saving report. Err: %s", err.Error()))
		return
	}
	pubsub, err := queue.NewMasterStreamPubSub(config.GetQueueConfig())
	if err != nil {
		d.finalizeRound(ctx, &rnd, Note, fmt.Sprintf("Error while notifying report update. Err: %s", err.Error()))
		return
	}
	pubsub.NotifyTopic(config.GetPubSubConfig().ChannelPrefix + "_report")
	d.finalizeRound(ctx, &rnd, Note, "")
}

func (d dRunner) finalizeRound(ctx context.Context, rnd *round.Round, Note string, Error string) {
	log.Printf("Note: %s\nError: %s\nRound: %v", Note, Error, rnd)
	now := time.Now().Add(d.getDsync())
	rnd.Finish = &now
	rnd.Note = Note
	rnd.Err = Error
	err := d.r.Round.Update(ctx, rnd)
	if err != nil {
		log.Printf("Unable to update round %v", rnd)
	}
}

func PropertyToMap(props []*property.Property) map[string]string {
	params := map[string]string{}
	for _, p := range props {
		params[p.Key] = *p.Value
	}
	return params
}
