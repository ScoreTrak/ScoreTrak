package schedulerfx

import (
	"github.com/scoretrak/scoretrak/pkg/scheduler"
	"github.com/scoretrak/scoretrak/pkg/scheduler/job"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// Cron
	fx.Provide(
		scheduler.NewCron,
		scheduler.NewElector,
	),
	fx.Invoke(
		scheduler.StartCron,
		scheduler.StartElector,
	),

	// Ping Job

	fx.Provide(
		job.NewPingJob,
	),
	fx.Invoke(
		job.RegisterPingJob,
	),

	// RoundStarterJob
	fx.Provide(
		job.NewRoundStarterJob,
	),
	fx.Invoke(
		job.RegisterRoundStarterJob,
	),

	// RoundFinisherJob
)
