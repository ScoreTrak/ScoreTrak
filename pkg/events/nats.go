package events

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/v2/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	wotelfloss "github.com/dentech-floss/watermill-opentelemetry-go-extra/pkg/opentelemetry"
	nc "github.com/nats-io/nats.go"
	wotel "github.com/voi-oss/watermill-opentelemetry/pkg/opentelemetry"
	"log"
	"strings"
)

// TODO: create global pub sub and all consuming sub

func NewNatsPublisher(cfg *config.Config, logger watermill.LoggerAdapter) (message.Publisher, error) {
	qCfg := cfg.Queue.NATS
	natsPublisher, err := nats.NewPublisher(nats.PublisherConfig{
		URL: qCfg.Url,
		JetStream: nats.JetStreamConfig{
			Disabled:       qCfg.Jetstream.Disabled,
			AutoProvision:  qCfg.Jetstream.AutoProvision,
			AckAsync:       false,
			TrackMsgId:     true,
			ConnectOptions: []nc.JSOpt{},
			PublishOptions: []nc.PubOpt{},
		},
	}, logger)
	if err != nil {
		return nil, err
	}
	tracePropagatingPublisherDecorator := wotelfloss.NewTracePropagatingPublisherDecorator(natsPublisher)
	return wotel.NewNamedPublisherDecorator("operation.Publish", tracePropagatingPublisherDecorator), nil
}

func NewNatsSubscriber(cfg *config.Config, logger watermill.LoggerAdapter) (message.Subscriber, error) {
	qCfg := cfg.Queue.NATS
	natsSubscriber, err := nats.NewSubscriber(nats.SubscriberConfig{
		URL:              qCfg.Url,
		SubscribersCount: qCfg.SubscriberCount,
		QueueGroupPrefix: qCfg.QueueGroupPrefix,
		JetStream: nats.JetStreamConfig{
			Disabled: qCfg.Jetstream.Disabled,
			//DurablePrefix:    qCfg.Jetstream.DurablePrefix,
			AutoProvision:    qCfg.Jetstream.AutoProvision,
			AckAsync:         false,
			TrackMsgId:       true,
			ConnectOptions:   []nc.JSOpt{},
			SubscribeOptions: []nc.SubOpt{},
			DurableCalculator: func(s string, s2 string) string {
				parts := strings.Split(s2, ".")
				if len(parts) > 1 {
					afterDot := parts[1]
					log.Printf("Durable Name is %s", afterDot)
					return afterDot
				} else {
					return s
				}
			},
		},
	}, logger)
	if err != nil {
		return nil, err
	}
	return natsSubscriber, nil
}

func NewNatsPubSub(cfg *config.Config, logger watermill.LoggerAdapter) (message.Publisher, message.Subscriber, error) {
	pub, err := NewNatsPublisher(cfg, logger)
	if err != nil {
		return nil, nil, err
	}
	sub, err := NewNatsSubscriber(cfg, logger)
	if err != nil {
		return nil, nil, err
	}
	return pub, sub, nil
}
