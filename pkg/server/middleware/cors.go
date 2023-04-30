package middleware

import (
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"net/http"
)

type CorsMiddleware struct {
	cfg *cors.Cors
}

func NewCorsMiddleware(cfg *cors.Cors) *CorsMiddleware {
	return &CorsMiddleware{cfg: cfg}
}

func NewCorsConstructor(cm *CorsMiddleware) alice.Constructor {
	return cm.Handler
}
func (c *CorsMiddleware) Handler(next http.Handler) http.Handler {
	return c.cfg.Handler(next)
}
