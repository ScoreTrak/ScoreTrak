package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ScoreTrak/ScoreTrak/cmd/master/server"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	cutil "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/di"
	diutil "github.com/ScoreTrak/ScoreTrak/pkg/di/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/runner"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage"
	sutil "github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func main() {
	flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.String("encoded-config", "", "Please enter encoded config")
	flag.Parse()
	path, err := cutil.ConfigFlagParser()
	if err != nil {
		handleErr(err)
	}
	handleErr(config.NewStaticConfig(path))
	cnf, err := config.NewDynamicConfig(path)
	if err != nil {
		handleErr(err)
	}

	staticConfig := config.GetStaticConfig()

	if !staticConfig.Prod {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	d, err := di.BuildMasterContainer()
	handleErr(err)
	handleErr(SetupDB(d))
	db := storage.GetGlobalDB()
	handleErr(sutil.CreateAllTables(db))
	err = sutil.LoadConfig(db, cnf)
	if err != nil {
		handleErr(err)
	}
	err = sutil.LoadReport(db)
	if err != nil {
		handleErr(err)
	}

	store := diutil.NewStore()

	var q queue.WorkerQueue
	di.Invoke(func(qu queue.WorkerQueue) {
		q = qu
	})
	dr := runner.NewRunner(db, q, store, staticConfig)
	go func() {
		handleErr(dr.MasterRunner())
	}()
	handleErr(server.Start(staticConfig, d, db))
}
func handleErr(err error) {
	if err != nil {
		log.Panicf("%v", err)
	} else {
		return
	}
}

func SetupDB(cont *dig.Container) error {
	var db *gorm.DB
	err := cont.Invoke(func(d *gorm.DB) {
		db = d
	})
	if err != nil {
		return err
	}
	var tm time.Time
	res, err := db.Raw("SELECT current_timestamp;").Rows()
	if err != nil {
		panic(err)
	}
	if res.Err() != nil {
		panic(err)
	}
	defer func(res *sql.Rows) {
		err := res.Close()
		if err != nil {
			log.Fatalln(fmt.Errorf("unable to close the database connection properly: %w", err))
		}
	}(res)
	for res.Next() {
		err := res.Scan(&tm)
		if err != nil {
			return err
		}
	}
	err = sutil.DatabaseOutOfSync(tm)
	if err != nil {
		return err
	}
	return nil
}
