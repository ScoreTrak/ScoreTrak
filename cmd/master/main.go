package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/cmd/master/run"
	"github.com/L1ghtman2k/ScoreTrak/cmd/master/server/gorilla"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/di"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage"
	"github.com/L1ghtman2k/ScoreTrak/pkg/storage/util"
	"os"
)

func main() {
	path := flag.String("config", "configs/config.yml", "Please enter a path to config file")
	encodedConfig := flag.String("encoded-config", "", "Please enter encoded config")
	flag.Parse()
	if *encodedConfig != "" {
		dec, err := base64.StdEncoding.DecodeString(*encodedConfig)
		handleErr(err)
		*path = "config.yml"
		f, err := os.Create(*path)
		handleErr(err)
		defer f.Close()
		_, err = f.Write(dec)
		handleErr(err)
		handleErr(f.Sync())
	} else if !configExists(*path) {
		handleErr(errors.New("you need to provide config"))
	}
	handleErr(config.NewStaticConfig(*path))
	cnf, err := config.NewDynamicConfig(*path)
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
	handleErr(svr.Start())
	var q queue.Queue
	di.Invoke(func(qu queue.Queue) {
		q = qu
	})
	dr := run.NewRunner(db, l, q, util.NewRepoStore())
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

func configExists(f string) bool {
	file, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}
