package connect

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/server"
	"go.uber.org/fx"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func NewConnectServer(paths []string, handlers []http.Handler) *http.ServeMux {
	fmt.Printf("Length of paths %d", len(paths))
	fmt.Printf("Length of handlers %d", len(handlers))
	mux := http.NewServeMux()
	return mux
}

func InitConnectServer(lc fx.Lifecycle, config server.Config, mux *http.ServeMux) {
	address := fmt.Sprintf("%s:%s", config.Address, config.Port)
	tlsEnabled := config.TLS.CertFile != "" || config.TLS.KeyFile != ""

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Starting connect server")
			go func() {
				if tlsEnabled {
					http.ListenAndServeTLS(address, config.TLS.CertFile, config.TLS.KeyFile, h2c.NewHandler(mux, &http2.Server{}))
				} else {
					http.ListenAndServe(address, h2c.NewHandler(mux, &http2.Server{}))
				}
			}()
			return nil
		},
	})
}
