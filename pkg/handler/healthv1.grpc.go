package handler

import (
	"context"

	healthv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/grpc/health/v1"
)

type HealthV1Controller struct {
	healthv1.UnimplementedHealthServer
}

func NewHealthV1Controller() *HealthV1Controller {
	return &HealthV1Controller{}
}

func (h HealthV1Controller) Check(ctx context.Context, request *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{Status: healthv1.HealthCheckResponse_SERVING}, nil
}
