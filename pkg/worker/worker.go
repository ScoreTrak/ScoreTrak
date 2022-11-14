package worker

import "C"
import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"go.uber.org/fx"
	"log"
)

func InitWorker(lc fx.Lifecycle, workerQueue queue.WorkerQueue) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting Worker")
			go workerQueue.Receive()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Stopping Worker")
			return nil
		},
	})
}
