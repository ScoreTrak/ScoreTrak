package queue

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/golang-queue/nsq"
	"github.com/golang-queue/queue"
)

func NewNsqWorker(cfg *config.Config, logger queue.Logger) *nsq.Worker {
	w := nsq.NewWorker(
		nsq.WithTopic(cfg.Queue.NSQ.Worker.Topic),
		nsq.WithAddr(cfg.Queue.NSQ.Worker.NSQD),
		nsq.WithChannel(cfg.Queue.NSQ.Worker.Channel),
		nsq.WithMaxInFlight(cfg.Queue.NSQ.Worker.MaxInFlight),
		//nsq.
		nsq.WithLogger(logger),
	)

	return w
}
