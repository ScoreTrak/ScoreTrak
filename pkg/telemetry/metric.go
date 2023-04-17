package telemetry

import (
	"context"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewMetricExporter() *metric.Exporter {
	exp, _ := otlpmetrichttp.New(
		context.Background(),
	)
	return &exp
}

func NewMeterProvider(r *resource.Resource, exp *metric.Exporter) *metric.MeterProvider {
	mp := metric.NewMeterProvider(
		metric.WithResource(r),
		metric.WithReader(
			metric.NewPeriodicReader(*exp),
		),
	)

	return mp
}

//func RegisterMetricProvider(lc fx.Lifecycle, mp *metric.MeterProvider) {
//	//otel.setMeterProvider(mp)
//
//	lc.Append(fx.Hook{
//		OnStop: func(ctx context.Context) error {
//			log.Println("Stopping meter provider")
//			err := mp.Shutdown(ctx)
//			if err != nil {
//				return err
//			}
//			return nil
//		},
//	})
//}
