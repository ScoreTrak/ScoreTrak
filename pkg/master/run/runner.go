package run

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/storage/orm"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"time"
)

type drunner struct {
	db *gorm.DB
	l  logger.LogInfoFormat
}

func NewRunner(db *gorm.DB, l logger.LogInfoFormat) *drunner {
	return &drunner{
		db: db, l: l,
	}
}

func (d *drunner) MasterRunner(db *gorm.DB, l logger.LogInfoFormat) error {
	configLoop := time.NewTicker(config.MinRoundDuration)
	cnfRepo := orm.NewConfigRepo(db, l)
	var scoringLoop *time.Ticker
	rndRepo := orm.NewRoundRepo(d.db, d.l)
	rnd, err := rndRepo.GetElapsingRound()
	if err != nil {
		rnd, err = rndRepo.GetLastRound()
		if err != nil {
			rnd = d.attemptToScore(nil)
		}
	}
	scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))

	for {
		select {
		case <-configLoop.C:
			cnf, _ := cnfRepo.Get()
			if !config.IsEqual(cnf) {
				config.UpdateConfig(cnf)
				rnd, err = rndRepo.GetElapsingRound()
				if err != nil {
					rnd, err = rndRepo.GetLastRound()
					if err != nil {
						rnd = nil
					}
				}
				scoringLoop.Stop()
				scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))
			}
		case <-scoringLoop.C:
			rnd = d.attemptToScore(rnd)
			scoringLoop.Stop()
			scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))
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
func (d *drunner) attemptToScore(rnd *round.Round) (elapsingRound *round.Round) {
	if !config.GetEnabled() {
		return nil
	}
	r := round.Round{}
	if rnd == nil {
		r.ID = 1
	} else {
		r.ID = rnd.ID + 1
	}
	rndRepo := orm.NewRoundRepo(d.db, d.l)
	err := rndRepo.Store(&r)
	if err != nil {
		serr, ok := err.(*pq.Error)
		if ok && serr.Code.Name() == "unique_violation" {
			return &r
		}
		d.l.Error(err)
		panic(err)
	}
	go d.Score(r)
	return &r
}

func (d drunner) Score(rnd round.Round) {
	//Perform Scoring
	//Save
}
