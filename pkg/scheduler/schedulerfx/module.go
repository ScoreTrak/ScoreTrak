package schedulerfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/scheduler"
	"go.uber.org/fx"
)

var Module = fx.Options(
	// fx.Provide(scheduler.NewRunner),
	// fx.Invoke(scheduler.InitRunner),
	fx.Provide(
		scheduler.NewCron,
		scheduler.NewScheduler,
	),
	fx.Invoke(
		scheduler.StartCron,
		scheduler.RegisterScheduler,
	),
)
