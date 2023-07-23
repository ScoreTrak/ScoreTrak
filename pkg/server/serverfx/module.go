package serverfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/server"
	"github.com/ScoreTrak/ScoreTrak/pkg/server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		middleware.NewCorsConfig,
		middleware.NewCorsMiddleware,
		middleware.NewUserMiddleware,
		server.NewHandler,
		server.NewOgentServer,
		server.NewServer,
	),
	fx.Invoke(server.StartServer),
)
