package scorerfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/events/entries"
	"github.com/ScoreTrak/ScoreTrak/pkg/events/handlers"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		handlers.NewHostServiceScoreHandler,
	),
	fx.Provide(
		fx.Annotate(
			entries.NewHostServiceScoreNoPublishHandlerEntry,
			fx.ResultTags(`group:"noPublishHandlerEntries"`),
		),
	),
)
