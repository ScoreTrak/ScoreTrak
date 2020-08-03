package main

import (
	"flag"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/run"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/server/gorilla"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	cutil "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/di"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	sutil "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"os"
)

func main() {
	flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.String("encoded-config", "", "Please enter encoded config")
	flag.Parse()
	path, err := cutil.ConfigFlagParser()
	handleErr(config.NewStaticConfig(path))
	cnf, err := config.NewDynamicConfig(path)
	if err != nil {
		handleErr(err)
	}
	r := gorilla.NewRouter()
	d, err := di.BuildMasterContainer()
	handleErr(err)
	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})
	svr := gorilla.NewServer(r, d, l)
	svr.MapRoutes()
	handleErr(svr.SetupDB())
	db := storage.GetGlobalDB()

	dataPreload := os.Getenv("DATA_PRELOAD_ONLY")
	if dataPreload == "TRUE" {
		handleErr(sutil.CreateAllTables(db))
		sutil.DataPreload(db)
		return
	}

	handleErr(svr.Start())
	var q queue.Queue
	di.Invoke(func(qu queue.Queue) {
		q = qu
	})
	dr := run.NewRunner(db, l, q, sutil.NewRepoStore())
	var count int64
	db.Table("config").Count(&count)
	err = sutil.LoadConfig(db, cnf)
	if err != nil {
		handleErr(err)
	}
	err = sutil.LoadReport(db)
	if err != nil {
		handleErr(err)
	}
	handleErr(dr.MasterRunner(cnf))
}
func handleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	} else {
		return
	}
}
