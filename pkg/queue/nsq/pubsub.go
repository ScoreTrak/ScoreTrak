package nsq

import (
	"fmt"
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
	ProducerConfig(confp, p.config)
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%s", p.config.NSQ.NSQD.Host, p.config.NSQ.NSQD.Port), confp)
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
		ConsumerConfig(conf, p.config)
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
		addresses := generateNSQLookupdAddresses(p.config.NSQ.NSQLookupd.Hosts, p.config.NSQ.NSQLookupd.Port)
		err = consumer.ConnectToNSQLookupds(addresses)
		if err != nil {
			log.Fatalf("Unable to connect to NSQLookupd instances")
		}
		select {}
	}()
	return n
}

func NewNSQPubSub(config queueing.Config) (*PubSub, error) {
	return &PubSub{config}, nil
}
