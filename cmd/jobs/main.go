package main

import (
	"flag"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/cmd/master/server/gorilla"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	cutil "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/di"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	sutil "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/util"
	"os"
)

func main() {
	flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.String("encoded-config", "", "Please enter encoded config")
	skipBootstrap := flag.Bool("skip-bootstrap", false, "Specify this flag if you want to skip the setup of tables, report, and config(This operation is idempotent)")
	preloadData := flag.Bool("preload-data", false, "Specify this flag if you want to preload sample data into the database")
	flag.Parse()
	path, err := cutil.ConfigFlagParser()
	handleErr(config.NewStaticConfig(path))
	cnf, err := config.NewDynamicConfig(path)
	if err != nil {
		handleErr(err)
	}
	d, err := di.BuildMasterContainer()
	handleErr(err)
	var l logger.LogInfoFormat
	di.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})
	svr := gorilla.NewServer(nil, d, l)
	handleErr(svr.SetupDB())
	db := storage.GetGlobalDB()
	if !*skipBootstrap {
		handleErr(svr.LoadTables(db))
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
	}
	if *preloadData {
		sutil.DataPreload(db)
	}
}
func handleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	} else {
		return
	}
}
