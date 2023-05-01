package scheduler

import (
	"context"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
)

func NewCron() *cron.Cron {
	return cron.New(cron.WithSeconds())
}

func StartCron(lc fx.Lifecycle, crn *cron.Cron) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			crn.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			crn.Stop()
			return nil
		},
	})
}
