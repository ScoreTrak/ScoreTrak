package configfx

import (
	"github.com/scoretrak/scoretrak/pkg/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.NewScoreTrakConfig,
	),
)
