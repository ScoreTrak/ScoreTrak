package server

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/server/middleware"
	"github.com/go-chi/chi/v5"
	middleware2 "github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/alice"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"net/http"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/ogent"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, c *config.Config, um *middleware.UserMiddleware, cm *middleware.CorsMiddleware, entityServer *ogent.Server, logger *otelzap.SugaredLogger) (*http.Server, error) {
	r := chi.NewRouter()
	r.Use(middleware2.Logger)
	r.Use(middleware2.Recoverer)
	r.Use()

	var middlewares alice.Chain
	if !c.Server.Cors.Enabled {
		middlewares = alice.New(middleware.NewUserConstructor(um))
	} else {
		middlewares = alice.New(middleware.NewCorsConstructor(cm), middleware.NewUserConstructor(um))
	}

	handler := middlewares.Then(entityServer)

	r.Mount("/", handler)
	//r.Mount("/api", handler)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.Server.Address, c.Server.Port),
		Handler: r,
	}

	return srv, nil
}

func StartServer(lc fx.Lifecycle, c *config.Config, srv *http.Server, logger *otelzap.SugaredLogger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Infoln("Starting Server")
			if c.Server.TLS.CertFile != "" && c.Server.TLS.KeyFile != "" {
				go srv.ListenAndServeTLS(c.Server.TLS.CertFile, c.Server.TLS.KeyFile)
			} else {
				go srv.ListenAndServe()
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
}
