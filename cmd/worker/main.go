package main

import (
	"flag"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	cutil "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"log"

	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
)

func main() {
	flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.String("encoded-config", "", "Please enter encoded config")
	flag.Parse()
	path, err := cutil.ConfigFlagParser()
	handleErr(err)

	handleErr(config.NewStaticConfig(path))
	staticConfig := config.GetStaticConfig()
	if !staticConfig.Prod {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
	q, err := queue.NewWorkerQueue(config.GetStaticConfig().Queue)
	handleErr(err)
	q.Receive()
}

func handleErr(err error) {
	if err != nil {
		log.Panicf("%v", err)
	} else {
		return
	}
}
