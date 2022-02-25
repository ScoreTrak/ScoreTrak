package handler

import (
	"context"

	healthv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/grpc/health/v1"
)

type HealthController struct {
	healthv1.UnimplementedHealthServer
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h HealthController) Check(ctx context.Context, request *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{Status: healthv1.HealthCheckResponse_SERVING}, nil
}
