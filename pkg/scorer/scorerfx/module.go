package scorerfx

import (
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		scorer.NewScorer,
	// handlers.NewHostServiceScoreHandler,
	),
	fx.Provide(
	// fx.Annotate(
	// 	entries.NewHostServiceScoreNoPublishHandlerEntry,
	// 	fx.ResultTags(`group:"noPublishHandlerEntries"`),
	// ),
	),
)
