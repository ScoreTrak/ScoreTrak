package master

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/di"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/run"
	"github.com/L1ghtman2k/ScoreTrak/pkg/master/server"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage"
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
	var q queue.Queue
	di.Invoke(func(qu queue.Queue) {
		q = qu
	})
	dr := run.NewRunner(db, l, q, run.NewRepoStore())
	return dr.MasterRunner()

}
