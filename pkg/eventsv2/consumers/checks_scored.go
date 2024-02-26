package consumers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/internal/entities/check"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/messages"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/streams"
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx"
	"time"
)

const (
	CHECK_SCORED_CONSUMER_NAME = "saver"
)

type ChecksScoredConsumer struct {
	dbClient *entities.Client
	scorer   *scorer.Scorer
	logger   *otelzap.SugaredLogger
	js       jetstream.JetStream
	consumer jetstream.Consumer
}

func NewChecksScoredConsumer(ctx context.Context, js jetstream.JetStream, entitiesClient *entities.Client, scorer *scorer.Scorer, logger *otelzap.SugaredLogger) (*ChecksScoredConsumer, error) {
	jsctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	consumer, err := js.CreateOrUpdateConsumer(jsctx, streams.STREAM_CHECKS, jetstream.ConsumerConfig{
		Name:    CHECK_SCORED_CONSUMER_NAME,
		Durable: CHECK_SCORED_CONSUMER_NAME,
		//AckPolicy: jetstream.AckExplicitPolicy,
		MaxDeliver:    10,
		FilterSubject: fmt.Sprintf("%s.*.scored", streams.STREAM_CHECKS),
	})
	if err != nil {
		return nil, err
	}

	return &ChecksScoredConsumer{
		dbClient: entitiesClient,
		scorer:   scorer,
		logger:   logger,
		consumer: consumer,
		js:       js,
	}, nil
}

func (c *ChecksScoredConsumer) Consume(msg jetstream.Msg) {
	ctx := context.Background()

	payload := messages.ChecksScoredMessage{}
	err := json.Unmarshal(msg.Data(), &payload)
	if err != nil {
		c.logger.Error(err)
		err := msg.Nak()
		if err != nil {
			c.logger.Error(err)
		}
		return
	}

	c.logger.Infow("Saving queue message", "check_id", payload.CheckID)

	chkUpdate := c.dbClient.Check.UpdateOneID(payload.CheckID).SetOutcomeStatus(check.OutcomeStatus(payload.Outcome.Status)).SetProgressStatus(check.ProgressStatusFinished)
	if payload.Outcome.Error != nil {
		chkUpdate.SetError(payload.Outcome.Error.Error())
	}

	chk, err := chkUpdate.Save(ctx)
	if err != nil {
		c.logger.Error(err)
		err := msg.Nak()
		if err != nil {
			c.logger.Error(err)
		}
		return
	}
	c.logger.Infoln(chk.OutcomeStatus)

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

func (c *ChecksScoredConsumer) Start(ctx context.Context) (jetstream.ConsumeContext, error) {
	cc, err := c.consumer.Consume(c.Consume)
	if err != nil {
		return nil, err
	}

	return cc, nil
}

func RegisterChecksScoredConsumer(ctx context.Context, lc fx.Lifecycle, csc *ChecksScoredConsumer) {
	var cc jetstream.ConsumeContext

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			consumeContext, err := csc.Start(ctx)
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
