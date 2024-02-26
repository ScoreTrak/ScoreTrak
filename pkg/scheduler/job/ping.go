package job

import (
	"context"
	"github.com/go-co-op/gocron/v2"
	"github.com/scoretrak/scoretrak/pkg/config"
	"github.com/scoretrak/scoretrak/pkg/scheduler"
	"go.uber.org/fx"
	"log"
	"time"
)

const (
	PING_JOB_TAG  = "PING"
	PING_JOB_NAME = "PING"
)

type PingJob struct {
}

func NewPingJob() *PingJob {
	return &PingJob{}
}

func RegisterPingJob(lc fx.Lifecycle, cfg *config.Config, pj *PingJob, s *scheduler.Scheduler) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			job, err := s.CronScheduler.NewJob(
				gocron.DurationJob(
					time.Second*time.Duration(cfg.Scheduler.Jobs.Ping.FrequencySecond),
				),
				gocron.NewTask(pj.task),
				gocron.WithTags(PING_JOB_TAG),
				gocron.WithName(PING_JOB_NAME),
				gocron.WithSingletonMode(gocron.LimitModeReschedule),
				//gocron.WithSingletonMode(gocron.LimitModeWait),

			)
			if err != nil {
				return err
			}
			log.Printf("job started id: %s", job.ID())
			return nil
		},
		OnStop: func(ctx context.Context) error {
			s.CronScheduler.RemoveByTags(PING_JOB_TAG)
			return nil
		},
	})
}

func (j *PingJob) task() {
	log.Println("ping")
}
