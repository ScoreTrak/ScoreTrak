package main

import (
	"flag"
	"fmt"
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
	"log"
	"math"
	"time"
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
		log.Fatalf("%v", err)
	} else {
		return
	}
}

func SetupDB(cont *dig.Container) error {
	var db *gorm.DB
	cont.Invoke(func(d *gorm.DB) {
		db = d
	})
	var tm time.Time
	res, err := db.Raw("SELECT current_timestamp;").Rows()
	if err != nil {
		panic(err)
	}
	defer res.Close()
	for res.Next() {
		res.Scan(&tm)
	}
	timeDiff := time.Since(tm)
	if float64(time.Second*2) < math.Abs(float64(timeDiff)) {
		panic(fmt.Errorf("time difference between master host, and database host are is large. Please synchronize time\n(The difference should not exceed 2 seconds)\nTime on database:%s\nTime on master:%s", tm.String(), time.Now()))
	}
	return nil
}
