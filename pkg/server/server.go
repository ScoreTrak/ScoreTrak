package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	middleware2 "github.com/go-chi/chi/v5/middleware"
	api_stub "github.com/scoretrak/scoretrak/internal/api-stub"
	"github.com/scoretrak/scoretrak/pkg/server/middleware"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"net/http"

	"github.com/scoretrak/scoretrak/pkg/config"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, c *config.Config, s *api_stub.Server, um *middleware.UserMiddleware, cm *middleware.CorsMiddleware, logger *otelzap.SugaredLogger) (*http.Server, error) {
	r := chi.NewRouter()
	r.Use(middleware2.Logger)
	r.Use(middleware2.Recoverer)

	//var middlewares alice.Chain
	//if !c.Server.Cors.Enabled {
	//	middlewares = alice.New(middleware.NewUserConstructor(um))
	//} else {
	//	middlewares = alice.New(middleware.NewCorsConstructor(cm), middleware.NewUserConstructor(um))
	//}
	//
	//handler := middlewares.Then(r)

	r.Mount("/", s)
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
