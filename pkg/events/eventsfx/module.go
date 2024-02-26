package eventsfx

import (
	"github.com/scoretrak/scoretrak/pkg/events"
	"github.com/scoretrak/scoretrak/pkg/events/entries"
	"github.com/scoretrak/scoretrak/pkg/events/handlers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		events.NewPubSub,
		handlers.NewHostServiceScoreHandler,
		handlers.NewCheckSaveHandler,
	),
	fx.Provide(
		fx.Annotate(
			events.NewRouter,
			fx.ParamTags(`group:"handlerEntries"`, `group:"noPublishHandlerEntries"`),
		),
		fx.Annotate(
			entries.NewHostServiceScoreHandlerEntry,
			fx.ResultTags(`group:"handlerEntries"`),
		),
		fx.Annotate(
			entries.NewCheckSaveNoPublishHandlerEntry,
			fx.ResultTags(`group:"noPublishHandlerEntries"`),
		),
	),
	fx.Invoke(
		events.StartRouter,
	),
)
