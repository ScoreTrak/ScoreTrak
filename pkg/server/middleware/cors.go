package middleware

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"net/http"
)

func NewCorsConstructor(cm *CorsMiddleware) alice.Constructor {
	return cm.Handler
}

type CorsMiddleware struct {
	cfg *cors.Cors
}

func NewCorsConfig(c *config.Config) *cors.Cors {
	crsCfg := c.Server.Cors
	crs := cors.New(cors.Options{
		AllowedOrigins:       crsCfg.AllowedOrigins,
		AllowedMethods:       crsCfg.AllowedMethods,
		AllowedHeaders:       crsCfg.AllowedHeaders,
		ExposedHeaders:       crsCfg.ExposedHeaders,
		MaxAge:               crsCfg.MaxAge,
		AllowCredentials:     crsCfg.AllowCredentials,
		AllowPrivateNetwork:  crsCfg.AllowPrivateNetwork,
		OptionsPassthrough:   crsCfg.OptionsPassthrough,
		OptionsSuccessStatus: crsCfg.OptionsSuccessStatus,
		Debug:                crsCfg.Debug,
	})
	return crs
}

func NewCorsMiddleware(cfg *cors.Cors) *CorsMiddleware {
	return &CorsMiddleware{cfg: cfg}
}

func (c *CorsMiddleware) Handler(next http.Handler) http.Handler {
	return c.cfg.Handler(next)
}
