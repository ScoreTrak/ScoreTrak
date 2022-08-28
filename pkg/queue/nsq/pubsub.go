package nsq

import (
	"log"

	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/nsqio/go-nsq"
)

type PubSub struct {
	config queueing.Config
}

func (p PubSub) NotifyTopic(topic string) {
	confp := nsq.NewConfig()
	nsqProducerConfig(confp, p.config)
	producer, err := nsq.NewProducer(p.config.NSQ.ProducerNSQD, confp)
	if err != nil {
		log.Panicf("Unable to initialize producer to notify masters using queue. Ensure that the queue is reachable from master. Error Details: %v", err)
	}
	producer.SetLoggerLevel(nsq.LogLevelWarning)
	err = producer.Publish(topic, make([]byte, 1))
	if err != nil {
		log.Panicf("Unable to publish to topic to notify masters. Ensure that the queue is reachable from master. Error Details: %v", err)
	}
	producer.Stop()
}

func (p PubSub) ReceiveUpdateFromTopic(topic string) <-chan struct{} {
	channel := make(chan struct{})
	go func() {
		conf := nsq.NewConfig()
		nsqConsumerConfig(conf, p.config)
		rand, err := queueing.RandomInt()
		if err != nil {
			log.Panicf("unable to generate random number: %v", err)
		}
		consumer, err := nsq.NewConsumer(topic, "master_"+rand, conf)
		if err != nil {
			log.Panicf("Unable to initualize consumer for topic: %s. Error Details: %v", topic, err)
		}
		consumer.SetLoggerLevel(nsq.LogLevelWarning)
		consumer.AddHandler(
			nsq.HandlerFunc(func(m *nsq.Message) error {
				channel <- struct{}{}
				return nil
			}))
		err = connectConsumer(consumer, p.config)
		if err != nil {
			log.Panicf("Unable to establish connection with NSQ")
		}
		select {}
	}()
	return channel
}

func NewNSQPubSub(config queueing.Config) (*PubSub, error) {
	err := validateNSQConfig(config)
	if err != nil {
		return nil, err
	}
	return &PubSub{config}, nil
}
