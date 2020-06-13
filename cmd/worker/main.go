package main

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/queue"
	"fmt"
	"os"
)

func main() {
	if err := config.NewStaticConfig("configs/config.yml"); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
	q, err := queue.NewQueue(config.GetConfig())
	if err != nil {
		panic(err)
	}
	q.Receive()
}
