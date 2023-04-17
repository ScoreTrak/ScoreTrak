package middlewarefx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		middleware.NewKratosMiddleware,
	),
)
