package run

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/di/repo"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
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
	r  repo.Store
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
		d.l.Error(fmt.Sprintf("time difference between master host, and database host is too large. Please synchronize time\n(The difference should not exceed 2 seconds)\nTime on database:%s\nTime on master:%s", tm.String(), time.Now()))
	}
}
func (d dRunner) getDsync() time.Duration {
	mud.RLock()
	defer mud.RUnlock()
	return dsync
}

func NewRunner(db *gorm.DB, l logger.LogInfoFormat, q queue.Queue, r repo.Store) *dRunner {
	return &dRunner{
		db: db, l: l, q: q, r: r,
	}
}

func (d *dRunner) MasterRunner(cnf *config.DynamicConfig) (err error) {
	d.refreshDsync()
	var scoringLoop *time.Ticker
	r, _ := d.r.Round.GetLastNonElapsingRound()
	if r != nil {
		scoringLoop = time.NewTicker(d.durationUntilNextRound(r, cnf.RoundDuration))
	} else {
		scoringLoop = time.NewTicker(time.Second)
	}
	configLoop := time.NewTicker(config.MinRoundDuration)

	for {
		select {
		case <-configLoop.C:
			dcc, _ := d.r.Config.Get()
			if !cnf.IsEqual(dcc) {
				*cnf = *dcc
			}
			r, _ := d.r.Round.GetLastRound()
			d.refreshDsync()
			scoringLoop.Stop()
			scoringLoop = time.NewTicker(d.durationUntilNextRound(r, cnf.RoundDuration))
		case <-scoringLoop.C:
			scoringLoop.Stop()
			dcc, _ := d.r.Config.Get()
			if !cnf.IsEqual(dcc) {
				*cnf = *dcc
			}
			rnd := &round.Round{}
			if *cnf.Enabled {
				rNoneElapsing, _ := d.r.Round.GetLastNonElapsingRound()
				if rNoneElapsing != nil {
					rnd = &round.Round{ID: rNoneElapsing.ID + 1}
				} else {
					rnd = &round.Round{ID: 1}
				}

				lastRnd, _ := d.r.Round.GetLastRound()

				if lastRnd != nil && rnd.ID <= lastRnd.ID && lastRnd.Finish == nil {
					if time.Now().After(lastRnd.Start.Add(time.Duration(cnf.RoundDuration) * time.Second).Add(time.Second * 3)) {
						d.r.Round.Delete(lastRnd.ID)
					} //This is indicative of a scoring master dying. In this case we delete the last round
				}
				d.attemptToScore(rnd, time.Now().Add(d.getDsync()).Add(time.Duration(cnf.RoundDuration)*time.Second*9/10))
				scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd, cnf.RoundDuration))
			} else {
				scoringLoop = time.NewTicker(config.MinRoundDuration)
			}
		}
	}
}

func (d *dRunner) durationUntilNextRound(rnd *round.Round, RoundDuration uint) time.Duration {
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
				err = d.db.Model(&s).Association("Checks").Find(&s.Checks)
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
						}
						checks = append(checks, s.Checks[len(s.Checks)-1])
					}
				}
			}
		}
	}

	err = d.r.Check.Store(checks)
	if err != nil {
		d.l.Error(err)
		d.finalizeRound(&rnd, Note, fmt.Sprintf("Error while saving checks. Err: %s", err.Error()))
		return
	}
	r, _ := d.r.Round.GetLastRound()
	if r == nil || r.ID != rnd.ID {
		err = errors.New("A different round started before current round was able to finish. The scores will not be committed")
		d.l.Error(err)
		d.finalizeRound(&rnd, Note, fmt.Sprintf("Error while saving checks. Err: %s", err.Error()))
		return
	}
	reportServ := report.NewReportCalculator(d.r.Report)
	rep, err := reportServ.RecalculateReport(teams, hostGroup, serviceGroups, rnd)
	ch := report.NewReport()
	bt, err := json.Marshal(&rep)
	if err != nil {
		d.l.Error(err)
		d.finalizeRound(&rnd, Note, fmt.Sprintf("Error while saving report. Err: %s", err.Error()))
		return
	}
	ch.Cache = string(bt)
	err = d.r.Report.Update(ch)
	if err != nil {
		d.l.Error(err)
		if strings.Contains(err.Error(), "split failed while applying backpressure to") {
			Note += "\nPossible solution to error: decrease: gc.ttlseconds for database. "
		}
		d.finalizeRound(&rnd, Note, fmt.Sprintf("Error while saving report. Err: %s", err.Error()))
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
