package nsq

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"github.com/nsqio/go-nsq"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type PubSub struct {
	config queueing.Config
}

func (p PubSub) NotifyTopic(topic string) {
	confp := nsq.NewConfig()
	nsqProducerConfig(confp, p.config)
	producer, err := nsq.NewProducer(p.config.NSQ.ProducerNSQD, confp)
	if err != nil {
		log.Fatalf("Unable to initialize producer to notify masters using queue. Ensure that the queue is reachable from master. Error Details: %v", err)
	}
	err = producer.Publish(topic, make([]byte, 1))
	if err != nil {
		log.Fatalf("Unable to publish to topic to notify masters. Ensure that the queue is reachable from master. Error Details: %v", err)
	}
	producer.Stop()
}

func (p PubSub) ReceiveUpdateFromTopic(topic string) <-chan struct{} {
	n := make(chan struct{})
	go func() {
		conf := nsq.NewConfig()
		nsqConsumerConfig(conf, p.config)
		consumer, err := nsq.NewConsumer(topic, "master_"+strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int()), conf)
		if err != nil {
			log.Fatalf("Unable to initualize consumer for topic: %s. Error Details: %v", topic, err)
		}
		consumer.SetLoggerLevel(nsq.LogLevelError)
		consumer.AddHandler(
			nsq.HandlerFunc(func(m *nsq.Message) error {
				n <- struct{}{}
				return nil
			}))
		err = connectConsumer(consumer, p.config)
		if err != nil {
			log.Fatalf("Unable to establish connection with NSQ")
		}
		select {}
	}()
	return n
}

func NewNSQPubSub(config queueing.Config) (*PubSub, error) {
	err := validateNSQConfig(config)
	if err != nil {
		return nil, err
	}
	return &PubSub{config}, nil
}
