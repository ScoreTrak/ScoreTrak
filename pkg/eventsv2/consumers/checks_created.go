package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/messages"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/streams"
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx"
	"time"
)

const (
	CHECK_CREATED_CONSUMER_NAME = "scorer"
)

type ChecksCreatedConsumer struct {
	dbClient *entities.Client
	scorer   *scorer.Scorer
	logger   *otelzap.SugaredLogger
	js       jetstream.JetStream
	consumer jetstream.Consumer
}

func NewChecksCreatedConsumer(ctx context.Context, js jetstream.JetStream, entitiesClient *entities.Client, scorer *scorer.Scorer, logger *otelzap.SugaredLogger) (*ChecksCreatedConsumer, error) {
	jsctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	consumer, err := js.CreateOrUpdateConsumer(jsctx, streams.STREAM_CHECKS, jetstream.ConsumerConfig{
		Name:    CHECK_CREATED_CONSUMER_NAME,
		Durable: CHECK_CREATED_CONSUMER_NAME,
		//AckPolicy: jetstream.AckExplicitPolicy,
		MaxDeliver:    3,
		FilterSubject: fmt.Sprintf("%s.*.created", streams.STREAM_CHECKS),
	})
	if err != nil {
		return nil, err
	}

	return &ChecksCreatedConsumer{
		dbClient: entitiesClient,
		scorer:   scorer,
		logger:   logger,
		consumer: consumer,
		js:       js,
	}, nil
}

func (c *ChecksCreatedConsumer) Consume(msg jetstream.Msg) {
	ctx := context.Background()

	payload := messages.ChecksCreatedMessage{}
	err := json.Unmarshal(msg.Data(), &payload)
	if err != nil {
		c.logger.Error(err)
		err := msg.Nak()
		if err != nil {
			c.logger.Error(err)
		}
		return
	}

	c.logger.Infow("Scoring host service", "host_service_id", payload.HostServiceID)
	outcome := c.scorer.Score(ctx, payload.ServiceType, payload.Properties)
	if outcome.Error != nil {
		c.logger.Error(outcome.Error)
		err := msg.Nak()
		if err != nil {
			c.logger.Error(err)
		}
		return
	}

	chkScoredMsg := messages.ChecksScoredMessage{
		Outcome: outcome,
		CheckID: payload.CheckID,
	}
	chkScoredMsgBytes, err := json.Marshal(chkScoredMsg)
	if err != nil {
		c.logger.Error(err)
		err := msg.Nak()
		if err != nil {
			c.logger.Error(err)
		}
		return
	}

	_, err = c.js.PublishMsg(ctx, &nats.Msg{
		Data:    chkScoredMsgBytes,
		Subject: fmt.Sprintf("%s.%s.scored", streams.STREAM_CHECKS, payload.TeamID),
	})
	if err != nil {
		c.logger.Error(err)
		err := msg.Nak()
		if err != nil {
			c.logger.Error(err)
		}
		return
	}

	err = msg.Ack()
	if err != nil {
		c.logger.Error(err)
		err := msg.Nak()
		if err != nil {
			c.logger.Error(err)
		}
		return
	}
}

func (c *ChecksCreatedConsumer) Start(ctx context.Context) (jetstream.ConsumeContext, error) {
	cc, err := c.consumer.Consume(c.Consume)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func RegisterChecksCreatedConsumer(ctx context.Context, lc fx.Lifecycle, ccc *ChecksCreatedConsumer) {
	var cc jetstream.ConsumeContext

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			consumeContext, err := ccc.Start(ctx)
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
