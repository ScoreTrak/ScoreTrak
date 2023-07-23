package telemetry

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/version"
	"github.com/oklog/ulid/v2"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/semconv/v1.4.0"
)

const ServiceName = "scoretrak"

func NewResource() *resource.Resource {
	rs, _ := resource.New(
		context.Background(),
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithOS(),
		resource.WithContainer(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(ServiceName),
			semconv.ServiceVersionKey.String(version.Version),
			semconv.ServiceInstanceIDKey.String(ulid.Make().String()),
		),
	)

	return rs
}
