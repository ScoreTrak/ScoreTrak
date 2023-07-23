package events

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func NewGochannelPubsub(cfg *config.Config, logger watermill.LoggerAdapter) (message.Publisher, message.Subscriber, error) {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	return pubSub, pubSub, nil
}
