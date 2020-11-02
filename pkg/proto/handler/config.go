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
	return &configpb.GetResponse{DynamicConfig: ConvertDynamicConfigToDynamicConfigPB(cnf)}, nil
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

func ConvertDynamicConfigPBToDynamicConfig(pb *configpb.DynamicConfig) *config.DynamicConfig {
	var enabled *bool
	if pb.GetEnabled() != nil {
		enabled = &pb.GetEnabled().Value
	}
	return &config.DynamicConfig{
		RoundDuration: pb.GetRoundDuration(),
		Enabled:       enabled,
	}
}

func ConvertDynamicConfigToDynamicConfigPB(obj *config.DynamicConfig) *configpb.DynamicConfig {
	return &configpb.DynamicConfig{
		RoundDuration: obj.RoundDuration,
		Enabled:       &wrappers.BoolValue{Value: *obj.Enabled},
	}
}
