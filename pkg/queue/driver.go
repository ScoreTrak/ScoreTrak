package queue

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/queue/nsq"
	"ScoreTrak/pkg/queue/queueing"
	"errors"
)

type Queue interface {
	Send([]queueing.ScoringData) []queueing.QCheck
	Receive() queueing.ScoringData
	Acknowledge(queueing.QCheck)
}

func NewQueue(c *config.StaticConfig) (Queue, error) {
	if c.Queue.Use == "nsq" {
		return nsq.NewNSQQueue(c)
	}
	return nil, errors.New("not supported queue")
}
