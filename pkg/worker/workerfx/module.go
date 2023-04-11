package workerfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/worker"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		queue.NewWorkerQueue,
		worker.InitWorker,
	),
)
