package main

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue"
	"os"
)

func main() {
	if err := config.NewStaticConfig("configs/config.yml"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}

	l, err := logger.NewLogger(config.GetStaticConfig())
	if err != nil {
		panic(err)
	}
	q, err := queue.NewQueue(config.GetConfig(), l)
	if err != nil {
		panic(err)
	}
	q.Receive()
}
