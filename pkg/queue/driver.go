package queue

import (
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/none"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/nsq"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/queueing"
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
