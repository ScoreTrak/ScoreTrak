package server

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net"
	"net/http"

	"github.com/ScoreTrak/ScoreTrak/internal/entities/ogent"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"go.uber.org/fx"
)

func NewServer(lc fx.Lifecycle, c *config.Config, entityServer *ogent.Server) (*http.Server, error) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.Server.Address, c.Server.Port),
		Handler: cors.AllowAll().Handler(entityServer),
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			log.Println("Starting Server")
			if c.Server.TLS.CertFile != "" && c.Server.TLS.KeyFile != "" {
				go http.ServeTLS(ln, srv.Handler, c.Server.TLS.CertFile, c.Server.TLS.KeyFile)
			} else {
				go http.Serve(ln, srv.Handler)
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv, nil
}
