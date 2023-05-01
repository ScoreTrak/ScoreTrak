package queuefx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			queue.NewWorker,
			fx.ResultTags(`group:"queueOptions"`),
		),
		fx.Annotate(
			queue.NewQueue,
			fx.ParamTags(`group:"queueOptions"`),
		),
	),
	fx.Invoke(
		queue.StartQueue,
	),
)
