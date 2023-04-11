package serverfx

import (
	"net/http"

	"github.com/ScoreTrak/ScoreTrak/pkg/server"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		server.NewEntityServer,
		server.NewServer,
	),
	fx.Invoke(func(*http.Server) {}),
)
