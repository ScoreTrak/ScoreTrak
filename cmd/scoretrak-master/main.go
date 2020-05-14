package main

import (
	"ScoreTrak/pkg/di"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/master/server"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	r := server.NewRouter()

	d := di.BuildContainer()

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
