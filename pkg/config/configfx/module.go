package configfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			config.NewViperConfig,
			fx.ParamTags(`name:"cfgFile"`),
		),
		config.NewScoreTrakConfig,
	),
)
