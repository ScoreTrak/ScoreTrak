package consumers

import (
	"context"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/streams"
	"go.uber.org/fx"
	"log"
	"time"
)

const ()

type PrintConsumer struct {
	consumer jetstream.Consumer
}

func NewPrintConsumer(ctx context.Context, js jetstream.JetStream) (*PrintConsumer, error) {
	jsctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	consumer, err := js.CreateOrUpdateConsumer(jsctx, streams.STREAM_PRINT, jetstream.ConsumerConfig{
		Name:    "printer",
		Durable: "printer",
	})
	if err != nil {
		return nil, err
	}

	return &PrintConsumer{consumer: consumer}, nil
}

func (c *PrintConsumer) Consume(msg jetstream.Msg) {
	log.Println(string(msg.Data()))
}

func (c *PrintConsumer) Start(ctx context.Context) (jetstream.ConsumeContext, error) {
	cc, err := c.consumer.Consume(c.Consume)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func RegisterPrintConsumer(ctx context.Context, lc fx.Lifecycle, pc *PrintConsumer) {
	var cc jetstream.ConsumeContext

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			consumeContext, err := pc.Start(ctx)
			if err != nil {
				return err
			}
			cc = consumeContext
			return nil
		},
		OnStop: func(ctx context.Context) error {
			cc.Stop()
			return nil
		},
	})
}
