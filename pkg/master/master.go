package master

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/di"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/master/server"
	"ScoreTrak/pkg/storage"
	"ScoreTrak/pkg/storage/orm"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"time"
)

func Run() error {

	r := server.NewRouter()
	d, err := di.BuildMasterContainer()
	if err != nil {
		return err
	}
	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})
	svr := server.NewServer(r, d, l)
	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		return err
	}

	dc := config.GetConfigCopy()
	db := storage.GetGlobalDB()
	if err := db.Create(dc).Error; err != nil {
		return err
	}
	serr, ok := err.(*pq.Error)
	if ok && serr.Code.Name() == "unique_violation" {
		dcc := config.DynamicConfig{}
		db.Take(&dcc)
		config.UpdateConfig(&dcc)
	} else {
		return err
	}

	err = svr.Start()
	if err != nil {
		return err
	}

	return masterRoutine(db, l)

}

func masterRoutine(db *gorm.DB, l logger.LogInfoFormat) error {
	configLoop := time.NewTicker(config.MinRoundDuration * time.Second)
	rndRepo := orm.NewRoundRepo(db, l)
	cnfRepo := orm.NewConfigRepo(db, l)
	rnd, err := rndRepo.GetLastRound()

	var scoringLoop *time.Ticker

	//if scoring engine doesn't detect any rounds, it will start by giving itself a normal time.duration
	//otherwise it will set the timer to be scheduled at the start of the next round
	if err != nil {
		scoringLoop = time.NewTicker(time.Second * time.Duration(config.GetRoundDuration()))
	} else {
		scoringLoop = time.NewTicker(rnd.Start.Add(time.Duration(config.GetRoundDuration()) * time.Second).Sub(time.Now()))
	}

	for {
		select {
		case <-configLoop.C:
			cnf, _ := cnfRepo.Get()
			if !config.IsEqual(cnf) {
				config.UpdateConfig(cnf)
				scoringLoop = time.NewTicker(time.Second * time.Duration(config.GetRoundDuration()))
			}
		case <-scoringLoop.C:

		}

	}
}

//AttemptToScore should return last round that is being scored
//main loop should reassign/normalize timer based on AttemptToScoreAnswer, and round duration
