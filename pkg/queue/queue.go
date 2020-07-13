package queue

import (
	"errors"
	"github.com/L1ghtman2k/ScoreTrak/pkg/config"
	"github.com/L1ghtman2k/ScoreTrak/pkg/logger"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/none"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/nsq"
	"github.com/L1ghtman2k/ScoreTrak/pkg/queue/queueing"
	"github.com/L1ghtman2k/ScoreTrak/pkg/service_group"
)

type Queue interface {
	Send([]*queueing.ScoringData) (queue []*queueing.QCheck, bearable error, terminatable error)
	Receive()
	Acknowledge(queueing.QCheck)
	Ping(group *service_group.ServiceGroup) error
}

func NewQueue(c config.StaticConfig, l logger.LogInfoFormat) (Queue, error) {
	if c.Queue.Use == "nsq" {
		return nsq.NewNSQQueue(l)
	} else if c.Queue.Use == "none" {
		return none.NewNoneQueue(l)
	} else {
		return nil, errors.New("invalid queue selected")
	}
}
