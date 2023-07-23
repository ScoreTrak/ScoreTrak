package events

import "github.com/ThreeDotsLabs/watermill/message"

type HandlerEntry struct {
	HandlerName    string
	SubscribeTopic string
	Subscriber     message.Subscriber
	PublishTopic   string
	Publisher      message.Publisher
	Handler        message.HandlerFunc
}

func NewHandlerEntry(handlerName string, subscribeTopic string, subscriber message.Subscriber, publishTopic string, publisher message.Publisher, handler message.HandlerFunc) *HandlerEntry {
	return &HandlerEntry{
		HandlerName:    handlerName,
		SubscribeTopic: subscribeTopic,
		Subscriber:     subscriber,
		PublishTopic:   publishTopic,
		Publisher:      publisher,
		Handler:        handler,
	}
}

type NoPublishHandlerEntry struct {
	HandlerName      string
	SubscribeTopic   string
	Subscriber       message.Subscriber
	NoPublishHandler message.NoPublishHandlerFunc
}

func NewNoPublishHandlerEntry(handlerName string, subscribeTopic string, subscriber message.Subscriber, noPublishHandler message.NoPublishHandlerFunc) *NoPublishHandlerEntry {
	return &NoPublishHandlerEntry{
		HandlerName:      handlerName,
		SubscribeTopic:   subscribeTopic,
		Subscriber:       subscriber,
		NoPublishHandler: noPublishHandler,
	}
}
