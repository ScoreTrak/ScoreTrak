package schedulerfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/scheduler"
	"github.com/ScoreTrak/ScoreTrak/pkg/scheduler/job"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		scheduler.NewCron,
	),
	fx.Invoke(
		scheduler.StartCron,
	),
	fx.Provide(
		job.NewRoundStarterJob,
	),
	fx.Invoke(
		job.RegisterRoundStarterJob,
	),
)
