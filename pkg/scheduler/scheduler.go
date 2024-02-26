package scheduler

import (
	"context"
	"log"
	"time"

	elector "github.com/go-co-op/gocron-etcd-elector"
	"github.com/go-co-op/gocron/v2"
	"github.com/scoretrak/scoretrak/pkg/config"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/fx"
)

type Scheduler struct {
	CronScheduler gocron.Scheduler
	elector       *elector.Elector
}

func NewCron(cfg *config.Config) (*Scheduler, error) {
	var options []gocron.SchedulerOption

	if cfg.Scheduler.Elector.Enabled {
		var err error
		var el *elector.Elector

		electorCfg := elector.Config{
			Endpoints:   cfg.Scheduler.Elector.Endpoints,
			DialTimeout: 3 * time.Second,
		}

		el, err = elector.NewElector(context.Background(), electorCfg, elector.WithTTL(10))
		if err != nil {
			panic(err)
		}

		options = append(options, gocron.WithDistributedElector(el))
	}

	s, err := gocron.NewScheduler(options...)
	if err != nil {
		return nil, err
	}
	return &Scheduler{CronScheduler: s}, nil
}

func StartCron(lc fx.Lifecycle, cfg config.Config, s *Scheduler, logger *otelzap.SugaredLogger) {
	if cfg.Scheduler.Elector.Enabled {
		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				log.Printf("etcd key: %s", cfg.Scheduler.Elector.Key)
				err := s.elector.Start(cfg.Scheduler.Elector.Key)
				if err != nil {
					return err
				}
				return nil
			},
			OnStop: func(ctx context.Context) error {
				err := s.elector.Stop()
				if err != nil {
					return err
				}
				return nil
			},
		})
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Infoln("Starting cron")
			s.CronScheduler.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Infoln("Stopping cron")
			err := s.CronScheduler.Shutdown()
			if err != nil {
				return err
			}
			return nil
		},
	})
}
