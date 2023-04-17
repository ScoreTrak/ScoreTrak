package telemetry

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

var ServiceName = "scoretrak"

func NewTracerProvider(exp *otlptrace.Exporter, resource *resource.Resource) *trace.TracerProvider {
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(resource),
	)

	return tp
}

//func RegisterTracerProvider(lc fx.Lifecycle, tp *trace.TracerProvider) {
//	otel.SetTracerProvider(tp)
//
//	lc.Append(fx.Hook{
//		OnStop: func(ctx context.Context) error {
//			log.Println("Stopping tracer provider")
//			err := tp.Shutdown(ctx)
//			if err != nil {
//				return err
//			}
//			return nil
//		},
//	})
//}

func NewOtlpGrpcExporter() *otlptrace.Exporter {
	exp, _ := otlptracegrpc.New(
		context.Background(),
	)
	return exp
}

func NewOtlpHttpExporter() *otlptrace.Exporter {
	exp, _ := otlptracehttp.New(
		context.Background(),
	)
	return exp
}
