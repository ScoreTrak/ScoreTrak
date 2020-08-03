package main

import (
	"flag"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	cutil "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"os"
)

func main() {
	flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.String("encoded-config", "", "Please enter encoded config")
	flag.Parse()
	path, err := cutil.ConfigFlagParser()
	handleErr(config.NewStaticConfig(path))
	l, err := logger.NewLogger(config.GetStaticConfig().Logger)
	handleErr(err)
	q, err := queue.NewQueue(config.GetStaticConfig().Queue, l)
	handleErr(err)
	q.Receive()
}

func handleErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	} else {
		return
	}
}
