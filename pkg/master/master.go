package master

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/di"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/master/run"
	"ScoreTrak/pkg/master/server"
	"ScoreTrak/pkg/queue"
	"ScoreTrak/pkg/storage"
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
	db := storage.GetGlobalDB()
	err = svr.Start()
	if err != nil {
		return err
	}

	q, err := queue.NewQueue(config.GetStaticConfig(), l)

	if err != nil {
		return err
	}

	dr := run.NewRunner(db, l, q, run.NewRepoStore())
	return dr.MasterRunner()

}
