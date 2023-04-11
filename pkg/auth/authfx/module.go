package authfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		auth.NewOryClient,
		auth.NewKratosMiddleware,
		//auth.NewJWTManager,
		// auth.NewAuthInterceptor,
	),
)
