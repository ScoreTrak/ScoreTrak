package scheduler

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/telemetry/cronlogger"
	"github.com/robfig/cron/v3"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx"
)

func NewCron(cfg *config.Config, logger cronlogger.Logger) *cron.Cron {
	return cron.New(cron.WithLogger(logger), cron.WithSeconds(), cron.WithChain(cron.Recover(logger)))
}

func StartCron(lc fx.Lifecycle, c *cron.Cron, logger *otelzap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Infoln("Starting cron")
			c.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Infoln("Stopping cron")
			c.Stop()
			return nil
		},
	})
}
