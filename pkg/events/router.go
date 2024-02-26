package events

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"go.uber.org/fx"
)

func NewRouter(handlerEntries []*HandlerEntry, noPublishHandlerEntries []*NoPublishHandlerEntry, logger watermill.LoggerAdapter) (*message.Router, error) {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		middleware.Recoverer,
		//	middleware.CorrelationID,
	)

	for _, handlerEntry := range handlerEntries {
		logger.Info(fmt.Sprintf("Adding handler %s", handlerEntry.HandlerName), watermill.LogFields{})
		router.AddHandler(handlerEntry.HandlerName, handlerEntry.SubscribeTopic, handlerEntry.Subscriber, handlerEntry.PublishTopic, handlerEntry.Publisher, handlerEntry.Handler)
	}

	for _, noPublishHandlerEntry := range noPublishHandlerEntries {
		logger.Info(fmt.Sprintf("Adding non publishing handler %s", noPublishHandlerEntry.HandlerName), watermill.LogFields{})
		router.AddNoPublisherHandler(noPublishHandlerEntry.HandlerName, noPublishHandlerEntry.SubscribeTopic, noPublishHandlerEntry.Subscriber, noPublishHandlerEntry.NoPublishHandler)
	}

	return router, nil
}

func StartRouter(lc fx.Lifecycle, router *message.Router) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := router.Run(context.Background())
				if err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return router.Close()
		},
	})
}
