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
	if !configExists(*path) {
		fmt.Fprintf(os.Stderr, "You need to provide config!")
		os.Exit(-1)
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

func configExists(f string) bool {
	file, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !file.IsDir()
}
