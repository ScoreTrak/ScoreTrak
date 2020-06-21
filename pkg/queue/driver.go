package queue

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/logger"
	"ScoreTrak/pkg/queue/none"
	"ScoreTrak/pkg/queue/nsq"
	"ScoreTrak/pkg/queue/queueing"
)

type Queue interface {
	Send([]*queueing.ScoringData) (queue []*queueing.QCheck, bearable error, terminatable error)
	Receive()
	Acknowledge(queueing.QCheck)
}

func NewQueue(c *config.StaticConfig, l logger.LogInfoFormat) (Queue, error) {
	if c.Queue.Use == "nsq" {
		return nsq.NewNSQQueue(l)
	} else {
		return none.NewNoneQueue(l)
	}
}
