package run

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/di"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/queue"
	"ScoreTrak/pkg/queue/queueing"
	"ScoreTrak/pkg/report"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/team"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"math"
	"sync"
	"time"
)

type RepoStore struct {
	Round        round.Repo
	Host         host.Repo
	HostGroup    host_group.Repo
	Service      service.Repo
	ServiceGroup service_group.Repo
	Team         team.Repo
	Check        check.Repo
	Property     property.Repo
	Config       config.Repo
}

func NewRepoStore() RepoStore {
	var hostGroupRepo host_group.Repo
	di.Invoke(func(re host_group.Repo) {
		hostGroupRepo = re
	})
	var hostRepo host.Repo
	di.Invoke(func(re host.Repo) {
		hostRepo = re
	})
	var roundRepo round.Repo
	di.Invoke(func(re round.Repo) {
		roundRepo = re
	})
	var serviceRepo service.Repo
	di.Invoke(func(re service.Repo) {
		serviceRepo = re
	})
	var serviceGroupRepo service_group.Repo
	di.Invoke(func(re service_group.Repo) {
		serviceGroupRepo = re
	})
	var propertyRepo property.Repo
	di.Invoke(func(re property.Repo) {
		propertyRepo = re
	})
	var checkRepo check.Repo
	di.Invoke(func(re check.Repo) {
		checkRepo = re
	})
	var teamRepo team.Repo
	di.Invoke(func(re team.Repo) {
		teamRepo = re
	})
	var configRepo config.Repo
	di.Invoke(func(re config.Repo) {
		configRepo = re
	})

	return RepoStore{
		Round:        roundRepo,
		HostGroup:    hostGroupRepo,
		Host:         hostRepo,
		Service:      serviceRepo,
		ServiceGroup: serviceGroupRepo,
		Property:     propertyRepo,
		Check:        checkRepo,
		Team:         teamRepo,
		Config:       configRepo,
	}
}

type dRunner struct {
	db *gorm.DB
	l  logger.LogInfoFormat
	q  queue.Queue
	r  RepoStore
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
	for res.Next() {
		res.Scan(&tm)
	}
	dsync = -time.Since(tm)

	if float64(time.Second*2) < math.Abs(float64(dsync)) {
		d.l.Error(fmt.Sprintf("time difference between master host, and database host are is large. Please synchronize time\n(The difference should not exceed 2 seconds)\nTime on database:%s\nTime on master:%s", tm.String(), time.Now()))
	}
	res.Close()
}

func (d dRunner) getDsync() time.Duration {
	mud.RLock()
	defer mud.RUnlock()
	return dsync
}

func NewRunner(db *gorm.DB, l logger.LogInfoFormat, q queue.Queue, r RepoStore) *dRunner {
	return &dRunner{
		db: db, l: l, q: q, r: r,
	}
}

func (d *dRunner) MasterRunner() error {
	cnf, err := config.NewDynamicConfig("configs/config.yml")
	if err != nil {
		return err
	}
	err = d.db.Create(cnf).Error
	if err != nil {
		serr, ok := err.(*pq.Error)
		if ok && serr.Code.Name() == "unique_violation" {
			dcc := &config.DynamicConfig{}
			d.db.Take(dcc)
			*cnf = *dcc
		} else {
			return err
		}
	}
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
		return config.MinRoundDuration
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
		serr, ok := err.(*pq.Error)
		if ok && serr.Code.Name() == "unique_violation" {
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

func (d dRunner) Score(rnd round.Round, timeout time.Time) {
	defer func() {
		if err := recover(); err != nil {
			d.l.Error(err)
		}
	}()
	d.l.Info("Running check for round %d", rnd.ID)
	teams, err := d.r.Team.GetAll()
	if err != nil {
		d.finalizeRound(&rnd)
		return
	}
	hostGroup, _ := d.r.HostGroup.GetAll()
	serviceGroups, err := d.r.ServiceGroup.GetAll()
	if err != nil {
		d.finalizeRound(&rnd)
		return
	}
	var sds []*queueing.ScoringData
	for _, t := range teams {
		d.db.Model(&t).Related(&t.Hosts, "Hosts")
		for _, h := range t.Hosts {
			d.db.Model(&h).Related(&h.Services)
			for _, s := range h.Services {
				d.db.Model(&s).Related(&s.Properties)
				if *t.Enabled {
					var validHost bool
					if *h.Enabled {
						if h.HostGroupID != 0 {
							for _, hG := range hostGroup {
								if hG.ID == h.HostGroupID && *(hG.Enabled) {
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
									params := map[string]string{}
									for _, p := range s.Properties {
										params[p.Key] = p.Value
									}
									sd := &queueing.ScoringData{
										Timeout:    timeout,
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
		d.finalizeRound(&rnd)
		return
	}
	var chks []*queueing.QCheck
	chks = d.q.Send(sds)
	// ToDO: Find a better way to handle gorutines (Since they potentially may leak if  d.q.Send(sds) never returns). For Example if a queue dies, or execution takes too long (Say a check that never ends), then we are leaking a gorutine either on a worker side, or on master side.
	var checks []*check.Check
	for _, t := range teams {
		for _, h := range t.Hosts {
			for _, s := range h.Services {
				for _, c := range chks {
					if c.Service.ID == s.ID {
						s.Checks = append(s.Checks, &check.Check{Passed: &c.Passed, Log: c.Log, ServiceID: c.Service.ID, RoundID: rnd.ID, Err: c.Err})
						checks = append(checks, s.Checks...)
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
		schNew.Teams = make(map[uint64]report.SimpleTeam)
		for _, t := range teams {
			st := report.SimpleTeam{}
			st.Hosts = make(map[uint64]report.SimpleHost)
			for _, h := range t.Hosts {
				sh := report.SimpleHost{}
				sh.Services = make(map[uint64]report.SimpleService)
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
						if p.Status != "Hide" {
							params[p.Key] = p.Value
						}
					}
					sh.Services[s.ID] = report.SimpleService{Passed: *s.Checks[0].Passed, Log: s.Checks[0].Log, Err: s.Checks[0].Err, Points: points, Properties: params, PointsBoost: s.PointsBoost}
				}
				st.Hosts[h.ID] = sh
			}
			schNew.Teams[t.ID] = st
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
		return
	}
	err = d.r.Check.StoreMany(checks)
	if err != nil {
		d.l.Error(err)
		return
	}
	d.finalizeRound(&rnd)
}

func (d dRunner) finalizeRound(rnd *round.Round) {
	now := time.Now().Add(d.getDsync())
	rnd.Finish = &now
	err := d.r.Round.Update(rnd)
	if err != nil {
		d.l.Error(err)
	}
}
