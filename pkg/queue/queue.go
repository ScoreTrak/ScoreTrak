package queue

import (
	"errors"
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

func NewQueue(c queueing.Config, l logger.LogInfoFormat) (Queue, error) {
	if c.Use == "nsq" {
		return nsq.NewNSQQueue(l, c)
	} else if c.Use == "none" {
		return none.NewNoneQueue(l)
	} else {
		return nil, errors.New("invalid queue selected")
	}
}
