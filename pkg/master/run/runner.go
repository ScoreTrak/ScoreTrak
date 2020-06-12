package run

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/round"
	"ScoreTrak/pkg/storage/orm"
	"github.com/jinzhu/gorm"
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

func (d drunner) MasterRunner(db *gorm.DB, l logger.LogInfoFormat) error {
	configLoop := time.NewTicker(config.MinRoundDuration)
	cnfRepo := orm.NewConfigRepo(db, l)
	var scoringLoop *time.Ticker
	rndRepo := orm.NewRoundRepo(d.db, d.l)
	rnd, err := rndRepo.GetElapsingRound()
	if err != nil {
		rnd, err = rndRepo.GetLastRound()
		if err != nil {
			rnd = d.attemptToScore()
		}
	}
	scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))

	for {
		select {
		case <-configLoop.C:
			cnf, _ := cnfRepo.Get()
			if !config.IsEqual(cnf) {
				config.UpdateConfig(cnf)
				rnd, err := rndRepo.GetElapsingRound()
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
			rnd = d.attemptToScore()
			scoringLoop.Stop()
			scoringLoop = time.NewTicker(d.durationUntilNextRound(rnd))
		}

	}
}

func (d drunner) durationUntilNextRound(rnd *round.Round) time.Duration {
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
func (d drunner) attemptToScore() (elapsingRound *round.Round) {
	return nil
}

//AttemptToScore should return last round that is being scored
//main loop should reassign/normalize timer based on AttemptToScoreAnswer, and round duration

//CREATE CLEAR TIMELINE
