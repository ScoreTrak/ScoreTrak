package serverfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/server/middleware/middlewarefx"
	"net/http"

	"github.com/ScoreTrak/ScoreTrak/pkg/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	middlewarefx.Module,
	fx.Provide(
		server.NewHandler,
		server.NewEntityServer,
		server.NewServer,
	),
	fx.Invoke(func(*http.Server) {}),
)
