package master

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/di"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/master/run"
	"ScoreTrak/pkg/master/server"
	"ScoreTrak/pkg/storage"
	"github.com/lib/pq"
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
	dr := run.NewRunner(db, l)
	return dr.MasterRunner(db, l)

}
