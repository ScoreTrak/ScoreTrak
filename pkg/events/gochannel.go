package events

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/scoretrak/scoretrak/pkg/config"
)

func NewGochannelPubsub(cfg *config.Config, logger watermill.LoggerAdapter) (message.Publisher, message.Subscriber, error) {
	pubSub := gochannel.NewGoChannel(gochannel.Config{
		OutputChannelBuffer:            cfg.Queue.GoChannel.OutputChannelBuffer,
		Persistent:                     cfg.Queue.GoChannel.Persistent,
		BlockPublishUntilSubscriberAck: cfg.Queue.GoChannel.BlockPublishUntilSubscriberAck,
	}, logger)
	return pubSub, pubSub, nil
}
