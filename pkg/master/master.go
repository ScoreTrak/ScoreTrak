package master

import (
	"ScoreTrak/pkg/di"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/master/server"
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

	return svr.Start()
}
