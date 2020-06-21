package main

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/queue"
	"fmt"
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
