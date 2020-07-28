package run

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/queueing"
	"github.com/L1ghtman2k/ScoreTrak/pkg/report"
	"github.com/L1ghtman2k/ScoreTrak/pkg/round"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/util"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
	"math"
	"strings"
	"sync"
	"time"
)

type dRunner struct {
	db *gorm.DB
	l  logger.LogInfoFormat
	q  queue.Queue
	r  util.RepoStore
}

var dsync time.Duration
var mud sync.RWMutex

func (d dRunner) refreshDsync() {
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
		d.l.Error(fmt.Sprintf("time difference between master host, and database host are is large. Please synchronize time\n(The difference should not exceed 2 seconds)\nTime on database:%s\nTime on master:%s", tm.String(), time.Now()))
	}
}
func (d dRunner) getDsync() time.Duration {
	mud.RLock()
	defer mud.RUnlock()
	return dsync
}

func NewRunner(db *gorm.DB, l logger.LogInfoFormat, q queue.Queue, r util.RepoStore) *dRunner {
	return &dRunner{
		db: db, l: l, q: q, r: r,
	}
}

func (d *dRunner) MasterRunner(cnf *config.DynamicConfig) (err error) {
	err = d.db.Create(cnf).Error
	if err != nil {
		serr, ok := err.(*pgconn.PgError)
		if ok && serr.Code == "23505" {
			dcc := &config.DynamicConfig{}
			d.db.Take(dcc)
			*cnf = *dcc
		} else {
			return err
		}
	}
	d.refreshDsync()
	rpr := report.NewReport()
	d.db.Create(rpr)
	var scoringLoop *time.Ticker
	rnd, _ := d.r.Round.GetLastRound()
	if rnd == nil {
		rnd = &round.Round{ID: 0}
		if *(cnf.Enabled) {
			rnd.ID += 1
			d.attemptToScore(rnd, time.Now().Add(d.getDsync()).Add(time.Duration(cnf.RoundDuration)*time.Second*9/10))
		}
	}
	configLoop := time.NewTicker(config.MinRoundDuration)
	scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd, cnf.RoundDuration))
	for {
		select {
		case <-configLoop.C:
			dcc, _ := d.r.Config.Get()
			if !cnf.IsEqual(dcc) {
				*cnf = *dcc
			}
			r, _ := d.r.Round.GetLastRound()
			if r != nil {
				rnd = r
			}
			d.refreshDsync()
			scoringLoop.Stop()
			scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd, cnf.RoundDuration))
		case <-scoringLoop.C:
			scoringLoop.Stop()
			dcc, _ := d.r.Config.Get()
			if !cnf.IsEqual(dcc) {
				*cnf = *dcc
			}
			if *cnf.Enabled {
				rnd = &round.Round{ID: rnd.ID + 1}
				d.attemptToScore(rnd, time.Now().Add(d.getDsync()).Add(time.Duration(cnf.RoundDuration)*time.Second*9/10))
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
func (d *dRunner) attemptToScore(rnd *round.Round, timeout time.Time) {
	err := d.r.Round.Store(rnd)
	if err != nil {
		serr, ok := err.(*pgconn.PgError)
		if ok && serr.Code == "23505" {
			r, _ := d.r.Round.GetByID(rnd.ID)
			if r != nil {
				*rnd = *r
			}
		} else {
			d.l.Error(err)
			panic(err)
		}
	} else {
		go d.Score(*rnd, timeout)
	}
}

func (d dRunner) Score(rnd round.Round, deadline time.Time) {
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
			d.l.Error(err)
			d.finalizeRound(&rnd, Note, fmt.Sprintf("A panic has occured. Err:%s", err.Error()))
		}
	}()
	d.l.Info("Running check for round %d", rnd.ID)
	teams, err := d.r.Team.GetAll()
	if err != nil {
		d.finalizeRound(&rnd, Note, "No Teams Detected")
		return
	}
	hostGroup, _ := d.r.HostGroup.GetAll()
	serviceGroups, err := d.r.ServiceGroup.GetAll()
	if err != nil {
		d.finalizeRound(&rnd, Note, "No Service Groups Detected")
		return
	}
	var sds []*queueing.ScoringData
	for _, t := range teams {
		err = d.db.Model(&t).Association("Hosts").Find(&t.Hosts)
		if err != nil {
			panic(err)
		}
		for _, h := range t.Hosts {
			err = d.db.Model(&h).Association("Services").Find(&h.Services)
			if err != nil {
				panic(err)
			}
			for _, s := range h.Services {
				err = d.db.Model(&s).Association("Properties").Find(&s.Properties)
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
						if *(s.Enabled) && rnd.ID%schedule == 0 {
							for _, servGroup := range serviceGroups {
								if s.ServiceGroupID == servGroup.ID && *(servGroup.Enabled) {
									sq := queueing.QService{ID: s.ID, Group: servGroup.Name, Name: s.Name}
									params := PropertyToMap(s.Properties)
									sd := &queueing.ScoringData{
										Deadline:   deadline,
										Host:       *(h.Address),
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
		d.l.Info("No services are currently scorable. Finalizing the round")
		d.finalizeRound(&rnd, Note, "No scorable services detected")
		return
	}
	var chks []*queueing.QCheck
	var bearableErr error

	chks, bearableErr, err = d.q.Send(sds)
	if bearableErr != nil {
		Note += bearableErr.Error()
		d.l.Error(bearableErr)
	}
	if err != nil {
		d.l.Error(err)
		d.finalizeRound(&rnd, Note, err.Error())
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
							checks = append(checks, s.Checks...)
						}
					}
				}
			}
		}
	}
	err = d.db.Transaction(func(tx *gorm.DB) error {
		r, _ := d.r.Round.GetLastRound()
		if r.ID != rnd.ID {
			return errors.New("A different round started before current round was able to finish. The scores will not be committed")
		}
		ch := report.NewReport()
		if err := tx.Take(&ch).Error; err != nil {
			return err
		}
		schOld := report.SimpleReport{}
		err := json.Unmarshal([]byte(ch.Cache), &schOld)
		if err != nil {
			return err
		}
		schNew := report.SimpleReport{Round: rnd.ID}
		schNew.Teams = make(map[uint64]*report.SimpleTeam)
		for _, t := range teams {
			st := report.SimpleTeam{}
			st.Hosts = make(map[uint64]*report.SimpleHost)
			for _, h := range t.Hosts {
				sh := report.SimpleHost{}
				sh.Services = make(map[uint64]*report.SimpleService)
				for _, s := range h.Services {
					var points uint64
					if rnd.ID != 1 {
						if t1, ok := schOld.Teams[t.ID]; ok {
							if h1, ok := t1.Hosts[h.ID]; ok {
								if s1, ok := h1.Services[s.ID]; ok {
									points += s1.Points
								}
							}
						}
					}
					if len(s.Checks) != 0 && *(s.Checks[0].Passed) {
						points += s.Points
					}
					params := map[string]string{}
					for _, p := range s.Properties {
						params[p.Key] = p.Value
					}

					if len(s.Checks) != 0 {
						sh.Services[s.ID] = &report.SimpleService{Name: s.Name, DisplayName: s.DisplayName, Passed: *s.Checks[0].Passed, Log: s.Checks[0].Log, Err: s.Checks[0].Err, Points: points, Properties: params, PointsBoost: s.PointsBoost}
					} else {
						sh.Services[s.ID] = &report.SimpleService{Name: s.Name, DisplayName: s.DisplayName, Passed: false, Log: "Service was not checked because it was disabled", Err: "", Points: points, Properties: params, PointsBoost: s.PointsBoost}
					}

				}
				st.Hosts[h.ID] = &sh
			}
			schNew.Teams[t.ID] = &st
		}
		bt, err := json.Marshal(&schNew)
		if err != nil {
			return err
		}
		ch.Cache = string(bt)
		err = tx.Model(ch).Updates(report.Report{Cache: ch.Cache}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		d.l.Error(err)
		if strings.Contains(err.Error(), "split failed while applying backpressure to") {
			Note += "\nPossible solution to error: decrease: gc.ttlseconds for database. "
		}
		d.finalizeRound(&rnd, Note, fmt.Sprintf("Error while saving checks. Err: %s", err.Error()))
		return
	}
	err = d.r.Check.StoreMany(checks)
	if err != nil {
		d.l.Error(err)
		d.finalizeRound(&rnd, Note, fmt.Sprintf("Error while saving checks. Err: %s", err.Error()))
		return
	}
	d.finalizeRound(&rnd, Note, "")
}

func (d dRunner) finalizeRound(rnd *round.Round, Note string, Error string) {
	now := time.Now().Add(d.getDsync())
	rnd.Finish = &now
	rnd.Note = Note
	rnd.Err = Error
	err := d.r.Round.Update(rnd)
	if err != nil {
		d.l.Error(err)
	}
}

func PropertyToMap(props []*property.Property) map[string]string {
	params := map[string]string{}
	for _, p := range props {
		params[p.Key] = p.Value
	}
	return params
}
