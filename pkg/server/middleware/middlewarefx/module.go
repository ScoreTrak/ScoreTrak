package middlewarefx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		middleware.NewCorsMiddleware,
		middleware.NewAuthMiddleware,
	),
	fx.Provide(
		fx.Annotate(
			middleware.NewCorsConstructor,
			fx.ResultTags(`group:"constructors"`),
		),
		fx.Annotate(
			middleware.NewAuthConstructor,
			fx.ResultTags(`group:"constructors"`),
		),
		fx.Annotate(
			middleware.NewMiddlewareChain,
			fx.ParamTags(`group:"constructors"`),
		),
	),
)
