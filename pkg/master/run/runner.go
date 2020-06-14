package run

import (
	"ScoreTrak/pkg/check"
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/host"
	"ScoreTrak/pkg/host_group"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/property"
	"ScoreTrak/pkg/queue"
	"ScoreTrak/pkg/queue/queueing"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/service"
	"ScoreTrak/pkg/service_group"
	"ScoreTrak/pkg/team"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
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

type drunner struct {
	db *gorm.DB
	l  logger.LogInfoFormat
	q  queue.Queue
	r  RepoStore
}

func NewRunner(db *gorm.DB, l logger.LogInfoFormat, q queue.Queue, r RepoStore) *drunner {
	return &drunner{
		db: db, l: l, q: q, r: r,
	}
}

func (d *drunner) MasterRunner() error {
	var scoringLoop *time.Ticker
	rnd, _ := d.r.Round.GetLastRound()
	configLoop := time.NewTicker(config.MinRoundDuration)

	if rnd == nil {
		rnd = &round.Round{ID: 1}
		d.attemptToScore(rnd, time.Now().Add(time.Duration(config.GetRoundDuration())*time.Second*8/10))
	}
	scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))

	for {
		select {
		case <-configLoop.C:
			cnf, _ := d.r.Config.Get()
			if !config.IsEqual(cnf) {
				config.UpdateConfig(cnf)
			}
			scoringLoop.Stop()
			rnd, _ = d.r.Round.GetLastRound()
			scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))
		case <-scoringLoop.C:
			scoringLoop.Stop()
			cnf, _ := d.r.Config.Get()
			if !config.IsEqual(cnf) {
				break
			}
			if config.GetEnabled() {
				rnd = &round.Round{ID: rnd.ID + 1}
				d.attemptToScore(rnd, time.Now().Add(time.Duration(config.GetRoundDuration())*time.Second*8/10))
				scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))
			} else {
				scoringLoop = time.NewTicker(config.MinRoundDuration)
			}
		}
	}
}

func (d *drunner) durationUntilNextRound(rnd *round.Round) time.Duration {
	if rnd == nil {
		return config.MinRoundDuration
	}
	dur := rnd.Start.Add(time.Duration(config.GetRoundDuration()) * time.Second).Sub(time.Now())
	if dur <= 1 {
		return 1
	}
	return dur
}

//Runs check in the background as a gorutine.
func (d *drunner) attemptToScore(rnd *round.Round, timeout time.Time) {
	err := d.r.Round.Store(rnd)
	if err != nil {
		serr, ok := err.(*pq.Error)
		if ok && serr.Code.Name() == "unique_violation" {
			r, _ := d.r.Round.GetByID(rnd.ID)
			*rnd = *r
		} else {
			d.l.Error(err)
			panic(err)
		}
	} else {
		go d.Score(*rnd, timeout)
	}
}

func (d drunner) Score(rnd round.Round, timeout time.Time) {
	//TODO: TERMINATION BASED ON TIMEOUT
	teams, err := d.r.Team.GetAll()
	if err != nil {
		d.finalizeRound(&rnd)
		return
	}
	hostGroup, _ := d.r.HostGroup.GetAll()
	var hosts []*host.Host
	for _, t := range teams {
		if *(t.Enabled) {
			var hsts []*host.Host
			d.db.Model(&t).Related(&hsts)
			for _, h := range hsts {
				if h.HostGroupID != 0 {
					for _, hG := range hostGroup {
						if hG.ID == h.HostGroupID && *(hG.Enabled) {
							hosts = append(hosts, h)
						}
					}
				} else {
					hosts = append(hosts, h)
				}
			}
		}
	}
	if len(hosts) == 0 {
		d.finalizeRound(&rnd)
		return
	}
	serviceGroups, _ := d.r.ServiceGroup.GetAll()
	var services []*service.Service
	for _, h := range hosts {
		if *(h.Enabled) {
			var serv []*service.Service
			d.db.Model(&h).Related(&serv)
			for _, s := range serv {
				schedule := s.RoundUnits
				if s.RoundDelay != nil {
					schedule += *(s.RoundDelay)
				}
				if *(s.Enabled) && rnd.ID%schedule == 0 {
					for _, servGroup := range serviceGroups {
						if s.ServiceGroupID == servGroup.ID && *(servGroup.Enabled) {
							services = append(services, s)
						}
					}
				}
			}
		}
	}
	if len(services) == 0 {
		d.finalizeRound(&rnd)
		return
	}
	var sds []*queueing.ScoringData
	for _, s := range services {
		var prop []*property.Property
		d.db.Model(&s).Related(&prop)
		var servGroupName string
		for _, servGroup := range serviceGroups {
			if s.ServiceGroupID == servGroup.ID {
				servGroupName = servGroup.Name
			}
		}
		var hst string
		for _, h := range hosts {
			if s.HostID == h.ID {
				hst = *(h.Address)
			}
		}
		sq := queueing.QService{ID: s.ID, Group: servGroupName, Name: s.Name}
		params := map[string]string{}
		for _, p := range prop {
			params[p.Key] = p.Value
		}
		sd := &queueing.ScoringData{
			Timeout:    timeout,
			Host:       hst,
			Service:    sq,
			Properties: params,
			RoundID:    rnd.ID,
		}
		sds = append(sds, sd)
	}

	var chks []*queueing.QCheck
	completed := make(chan bool, 1)

	go func() {
		chks = d.q.Send(sds)
		completed <- true
	}()

	// Listen on our channel AND a timeout channel - which ever happens first.
	select {
	case <-completed:
		break
	case <-time.After(time.Until(timeout)):
		d.finalizeRound(&rnd)
	}

	r, _ := d.r.Round.GetLastRound()
	if r.ID != rnd.ID {
		d.finalizeRound(&rnd)
		return
	}
	var checks []*check.Check
	for _, c := range chks {
		checks = append(checks, &check.Check{Passed: &c.Passed, Log: c.Log, ServiceID: c.Service.ID, RoundID: rnd.ID})
	}
	err = d.r.Check.StoreMany(checks)
	if err != nil {
		d.l.Error(err)
	}
	d.finalizeRound(&rnd)
}

func (d drunner) finalizeRound(rnd *round.Round) {
	now := time.Now()
	rnd.Finish = &now
	err := d.r.Round.Update(rnd)
	if err != nil {
		d.l.Error(err)
	}
}
