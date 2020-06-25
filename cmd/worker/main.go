package main

import (
	"flag"
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"os"
)

func main() {
	path := flag.String("config", "configs/config.yml", "Please enter a path to config file")
	flag.Parse()
	if *path != "configs/config.yml" {
		_, err := os.Stat(*path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Provided config was not found")
			os.Exit(-1)
		}
	}
	if err := config.NewStaticConfig(*path); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

	l, err := logger.NewLogger(config.GetStaticConfig())
	if err != nil {
		panic(err)
	}
	q, err := queue.NewQueue(config.GetStaticConfig(), l)
	if err != nil {
		panic(err)
	}
	q.Receive()
}
