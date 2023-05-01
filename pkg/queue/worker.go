package queue

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/job"
	"github.com/golang-queue/nsq"
	"github.com/golang-queue/queue"
)

const (
	QUEUE_RING     = "ring"
	QUEUE_NSQ      = "nsq"
	QUEUE_NATS     = "nats"
	QUEUE_REDIS    = "redis"
	QUEUE_RABBITMQ = "rabbitmq"
)

func NewNsqWorker(cfg *config.Config) queue.Option {
	w := nsq.NewWorker(
		nsq.WithTopic(cfg.Queue.NSQ.Worker.Topic),
		nsq.WithAddr(cfg.Queue.NSQ.Worker.NSQD),
		nsq.WithChannel(cfg.Queue.NSQ.Worker.Channel),
		nsq.WithMaxInFlight(cfg.Queue.NSQ.Worker.MaxInFlight),
		//nsq.WithLogger(logger),
		nsq.WithRunFunc(job.Ping),
	)

	return queue.WithWorker(w)
}

func NewRingWorker(cfg *config.Config) queue.Option {
	return queue.WithFn(job.Ping)
}

func NewWorker(cfg *config.Config) (queue.Option, error) {
	switch cfg.Queue.Use {
	case QUEUE_RING:
		return NewRingWorker(cfg), nil
	case QUEUE_NSQ:
		return NewNsqWorker(cfg), nil
	default:
		return nil, nil
	}
}
