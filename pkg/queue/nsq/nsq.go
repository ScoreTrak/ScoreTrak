package nsq

import (
	"ScoreTrak/pkg/config"
	"ScoreTrak/pkg/queue/queueing"
	"fmt"
	"github.com/nsqio/go-nsq"
)

type NSQ struct {
	producer *nsq.Producer
	consumer *nsq.Consumer
}

func NewNSQProducer(c *config.StaticConfig) (*nsq.Producer, error) {
	conf := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQD.Host, c.Queue.NSQ.NSQD.Port), conf)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func NewNSQConsumer(c *config.StaticConfig) (*nsq.Consumer, error) {
	conf := nsq.NewConfig()
	conf.MaxInFlight = c.Queue.NSQ.MaxInFlight
	consumer, err := nsq.NewConsumer(c.Queue.NSQ.Topic, c.Queue.NSQ.Channel, conf)
	if err != nil {
		return nil, err
	}
	// Set the Handler for messages received by this Consumer. Can be called multiple times.
	// See also AddConcurrentHandlers.
	consumer.ChangeMaxInFlight(2)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(func(m *nsq.Message) error { return nil }), 2)
	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%s", c.Queue.NSQ.NSQLookupd.Host, c.Queue.NSQ.NSQLookupd.Port))
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func (n NSQ) Send(sd []*queueing.ScoringData) []queueing.QCheck {
	if n.producer == nil {
		panic("You should not call send if producer is not defined!")
	}
	return []queueing.QCheck{}
}

func (n NSQ) Receive() queueing.ScoringData {
	if n.consumer == nil {
		panic("You should not call receive if consumer is not defined!")
	}
	return queueing.ScoringData{}
}

func (n NSQ) Acknowledge(queueing.QCheck) {
	if n.consumer == nil {
		panic("You should not call acknowledge if consumer is not defined!")
	}
}

func NewNSQQueue(c *config.StaticConfig) (*NSQ, error) {
	producer, err := NewNSQProducer(c)

	if err != nil {
		return nil, err
	}

	consumer, err := NewNSQConsumer(c)

	if err != nil {
		return nil, err
	}

	return &NSQ{
		producer,
		consumer,
	}, nil

}
