package queue

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/none"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
)

//WorkerQueue is an interface that every queue plugin should implement
type WorkerQueue interface {
	//Send is a method that accepts an arraying of queueing.ScoringData, and returns back an array of already performed Checks.
	//Send also returns 2 types of errors: bearable, and terminatable. As name suggests, bearable error could be ignored during checks(for instance, a worker node that is not receiving checks)
	//while terminatable error terminates the scoring for that round. (Ex. Also when worker fails to receive the check, except when queueing.Config.NSQ.IgnoreAllScoresIfWorkerFails is set to true)
	Send([]*queueing.ScoringData) (queue []*queueing.QCheck, bearable error, terminatable error)
	//Receive is an interface used for queues that typically work over the network. Receive receives
	Receive()
	//Acknowledge is a response of a single queueing.QCheck instance sent from the worker back to the master after the check has been completed. The Send aggregates all of the acknowledgements, and combines them into a slice, before returning.
	//Typically implemented for a queue that is over the network
	Acknowledge(queueing.QCheck)
	//Ping is responsible for pinging the workers by sending a simple check to the queue, and getting a response from the network. Typically implemented for over the network queue
	Ping(group *service_group.ServiceGroup) error
}

//MasterStreamPubSub is an interface that allows masters to send signals to each other via pub-sub.
type MasterStreamPubSub interface {
	NotifyTopic(topic string)
	ReceiveUpdateFromTopic(topic string) <-chan struct{}
}

func NewMasterStreamPubSub(c queueing.Config) (MasterStreamPubSub, error) {
	if c.Use == "nsq" {
		return nsq.NewNSQPubSub(c)
	} else if c.Use == "none" {
		return none.NewNonePubSub(c)
	} else {
		return nil, errors.New("invalid pub-sub selected")
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
