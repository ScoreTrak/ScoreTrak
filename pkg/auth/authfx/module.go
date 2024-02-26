package authfx

import (
	"github.com/scoretrak/scoretrak/pkg/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		auth.NewOryClient,
	),
)
