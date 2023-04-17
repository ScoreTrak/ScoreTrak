package configfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.NewScoreTrakConfig,
	),
)
