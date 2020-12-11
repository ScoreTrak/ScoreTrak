package none

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/queue/queueing"
	"time"
)

type PubSub struct {
	config  queueing.Config
	chanSub map[string]chan struct{}
}

var pb *PubSub

func (p PubSub) NotifyTopic(topic string) {
	if _, ok := p.chanSub[topic]; !ok {
		p.chanSub[topic] = make(chan struct{})
	}
	p.chanSub[topic] <- struct{}{}
}

func (p PubSub) ReceiveUpdateFromTopic(topic string) <-chan struct{} {
	n := make(chan struct{})
	go func() {
		for {
			if _, ok := p.chanSub[topic]; ok {
				break
			}
			time.Sleep(time.Second)
		}
		for {
			n <- <-p.chanSub[topic]
		}
	}()
	return n
}

func NewNonePubSub(config queueing.Config) (*PubSub, error) {
	if pb == nil {
		pb = &PubSub{config, make(map[string]chan struct{})}
	}
	return pb, nil
}
