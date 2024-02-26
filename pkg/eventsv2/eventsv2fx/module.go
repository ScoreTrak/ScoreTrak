package eventsv2fx

import (
	"github.com/scoretrak/scoretrak/pkg/eventsv2"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/consumers"
	"github.com/scoretrak/scoretrak/pkg/eventsv2/streams"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		eventsv2.NewNats,
	),

	// Streams
	fx.Provide(
		streams.NewCheckStream,
		streams.NewPrintStream,
	),

	// Consumers
	fx.Provide(
		consumers.NewPrintConsumer,
		consumers.NewChecksCreatedConsumer,
		consumers.NewChecksScoredConsumer,
	),
)
