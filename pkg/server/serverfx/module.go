package serverfx

import (
	"github.com/scoretrak/scoretrak/pkg/server"
	"github.com/scoretrak/scoretrak/pkg/server/handler"
	"github.com/scoretrak/scoretrak/pkg/server/middleware"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		middleware.NewCorsConfig,
		middleware.NewCorsMiddleware,
		middleware.NewUserMiddleware,
		server.NewApiTokenSecurityHandler,
		server.NewApiServer,
		handler.NewHandler,
		server.NewServer,
	),
	fx.Invoke(server.StartServer),
)
