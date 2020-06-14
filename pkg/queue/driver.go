package queue

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/queue/none"
	"ScoreTrak/pkg/queue/nsq"
	"ScoreTrak/pkg/queue/queueing"
)

type Queue interface {
	Send([]*queueing.ScoringData) []*queueing.QCheck
	Receive()
	Acknowledge(queueing.QCheck)
}

func NewQueue(c *config.StaticConfig) (Queue, error) {
	if c.Queue.Use == "nsq" {
		return nsq.NewNSQQueue()
	} else {
		return none.NewNoneQueue()
	}
}
