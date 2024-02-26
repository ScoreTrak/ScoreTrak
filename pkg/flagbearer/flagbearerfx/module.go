package flagbearerfx

import (
	"github.com/scoretrak/scoretrak/pkg/flagbearer"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		flagbearer.NewFlagBearer,
	),
	fx.Invoke(
		flagbearer.RegisterFlagBearerCronJob,
	),
)
