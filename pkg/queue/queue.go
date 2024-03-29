package queue

import (
	"errors"

	"github.com/ScoreTrak/ScoreTrak/pkg/queue/none"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/nsq"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
)

// WorkerQueue is an interface that every queue plugin should implement
type WorkerQueue interface {
	// Send is a method that accepts an arraying of queueing.ScoringData, and returns back an array of already performed Checks.
	// Send also returns 2 types of errors: nonCritical, and Critical.
	// As name suggests, nonCritical error could be ignored during checks(for instance, a worker node that is not receiving checks)
	// while Critical error terminates the scoring for that round.
	// (Ex. Also when worker fails to receive the check, except when queueing.Config.NSQ.IgnoreAllScoresIfWorkerFails is set to true)
	Send([]*queueing.ScoringData) (queue []*queueing.QCheck, nonCriticalErr error, CriticalErr error)
	// Receive is an interface used for queues that typically work over the network. Receive receives
	Receive()
	// Acknowledge is a response of a single queueing.QCheck instance sent from the worker back to the master after the check has been completed. The Send aggregates all of the acknowledgements, and combines them into a slice, before returning.
	// Typically implemented for a queue that is over the network
	Acknowledge(queueing.QCheck)
	// Ping is responsible for pinging the workers by sending a simple check to the queue, and getting a response from the network.
	// Typically implemented for over the network queue
	Ping(group *servicegroup.ServiceGroup) error
}

// MasterStreamPubSub is an interface that allows masters to send signals to each other via pub-sub.
type MasterStreamPubSub interface {
	NotifyTopic(topic string)
	ReceiveUpdateFromTopic(topic string) <-chan struct{}
}

const (
	Nsq  = "nsq"
	None = "none"
)

var ErrInvalidPubSub = errors.New("invalid pubsub selected")

func NewMasterStreamPubSub(c queueing.Config) (MasterStreamPubSub, error) {
	switch c.Use {
	case Nsq:
		return nsq.NewNSQPubSub(c)
	case None:
		return none.NewNonePubSub(c)
	default:
		return nil, ErrInvalidPubSub
	}
}

var ErrInvalidQueue = errors.New("invalid queue selected")

func NewWorkerQueue(c queueing.Config) (WorkerQueue, error) {
	switch c.Use {
	case Nsq:
		return nsq.NewNSQWorkerQueue(c)
	case None:
		return none.NewNoneQueue()
	default:
		return nil, ErrInvalidQueue
	}
}
