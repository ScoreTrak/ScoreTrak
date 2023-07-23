package eventsfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/events"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		events.NewPubSub,
	),
	fx.Provide(
		fx.Annotate(
			events.NewRouter,
			fx.ParamTags(`group:"handlerEntries"`, `group:"noPublishHandlerEntries"`),
		),
	),
	fx.Invoke(
		events.StartRouter,
	),
)
