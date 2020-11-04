package queue

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/none"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
)

type WorkerQueue interface {
	Send([]*queueing.ScoringData) (queue []*queueing.QCheck, bearable error, terminatable error)
	Receive()
	Acknowledge(queueing.QCheck)
	Ping(group *service_group.ServiceGroup) error
}

type MasterStreamPubSub interface {
	NotifyTopic(topic string)
	ReceiveUpdateFromTopic(topic string) <-chan struct{}
}

func NewMasterStreamPubSub(c queueing.Config) (MasterStreamPubSub, error) {
	if c.Use == "nsq" {
		return nsq.NewNSQPubSub(c)
	} else if c.Use == "none" {
		panic("implement me")
	} else {
		return nil, errors.New("invalid pubsub selected")
	}
}

func NewWorkerQueue(c queueing.Config) (WorkerQueue, error) {
	if c.Use == "nsq" {
		return nsq.NewNSQWorkerQueue(c)
	} else if c.Use == "none" {
		return none.NewNoneQueue()
	} else {
		return nil, errors.New("invalid queue selected")
	}
}
