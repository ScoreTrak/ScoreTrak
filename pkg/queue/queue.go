package queue

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/none"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
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
