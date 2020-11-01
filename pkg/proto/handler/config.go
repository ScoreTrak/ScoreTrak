package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/service"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConfigController struct {
	svc service.Serv
}

func (p ConfigController) Get(ctx context.Context, request *configpb.GetRequest) (*configpb.GetResponse, error) {
	cnf, err := p.svc.Get(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &configpb.GetResponse{DynamicConfig: &configpb.DynamicConfig{
		RoundDuration: cnf.RoundDuration,
		Enabled:       &wrappers.BoolValue{Value: *cnf.Enabled},
	}}, nil
}

func (p ConfigController) Update(ctx context.Context, request *configpb.UpdateRequest) (*configpb.UpdateResponse, error) {
	tmspb := request.GetDynamicConfig()
	var enabled *bool
	if tmspb.GetEnabled() != nil {
		enabled = &tmspb.GetEnabled().Value
	}
	err := p.svc.Update(ctx, &config.DynamicConfig{
		RoundDuration: tmspb.RoundDuration,
		Enabled:       enabled,
	})
	if err != nil {

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &configpb.UpdateResponse{}, nil
}

func NewConfigController(svc service.Serv) *ConfigController {
	return &ConfigController{svc}
}
