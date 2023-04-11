package server

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCorsMiddleware(c config.Config) gin.HandlerFunc {
	ccfg := cors.DefaultConfig()
	ccfg.AllowAllOrigins = true

	return cors.New(ccfg)
}
