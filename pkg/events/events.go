package events

import (
	"errors"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/scoretrak/scoretrak/pkg/config"
)

const (
	PUBSUB_GOCHANNEL = "gochannel"
	PUBSUB_NATS      = "nats"
)

var ERROR_NO_PUBSUB_CHOSEN = errors.New("No pubsub chosen")

func NewPubSub(cfg *config.Config, logger watermill.LoggerAdapter) (message.Publisher, message.Subscriber, error) {
	switch cfg.Queue.Use {
	case PUBSUB_GOCHANNEL:
		return NewGochannelPubsub(cfg, logger)
	case PUBSUB_NATS:
		return NewNatsPubSub(cfg, logger)
	default:
		return nil, nil, ERROR_NO_PUBSUB_CHOSEN
	}
}
