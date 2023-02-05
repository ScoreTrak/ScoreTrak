package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/grpc/health/v1/healthv1grpc"
	healthv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/grpc/health/v1"
	"context"
)

type HealthV1Controller struct {
	healthv1grpc.UnimplementedHealthServer
}

func NewHealthV1Controller() *HealthV1Controller {
	return &HealthV1Controller{}
}

func (h HealthV1Controller) Check(ctx context.Context, request *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{Status: healthv1.HealthCheckResponse_SERVING}, nil
}
