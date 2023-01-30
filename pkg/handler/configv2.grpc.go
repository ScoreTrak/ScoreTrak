package handler

import (
	"context"
	"encoding/json"
	"fmt"
	configv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/config/v2"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configservice"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConfigV2Controller struct {
	svc configservice.Serv
	configv2.UnimplementedDynamicConfigServiceServer
}

func (p ConfigV2Controller) Get(ctx context.Context, _ *configv2.DynamicConfigServiceGetRequest) (*configv2.DynamicConfigServiceGetResponse, error) {
	cnf, err := p.svc.Get(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &configv2.DynamicConfigServiceGetResponse{DynamicConfig: ConvertDynamicConfigToDynamicConfigV2PB(cnf)}, nil
}

func (p ConfigV2Controller) Update(ctx context.Context, request *configv2.DynamicConfigServiceUpdateRequest) (*configv2.DynamicConfigServiceUpdateResponse, error) {
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
	return &configv2.DynamicConfigServiceUpdateResponse{}, nil
}

func NewConfigV2Controller(svc configservice.Serv) *ConfigV2Controller {
	return &ConfigV2Controller{svc: svc}
}

func ConvertDynamicConfigV2PBToDynamicConfig(pb *configv2.DynamicConfig) *config.DynamicConfig {
	var enabled *bool
	if pb.GetEnabled() != nil {
		enabled = &pb.GetEnabled().Value
	}
	return &config.DynamicConfig{
		RoundDuration: pb.GetRoundDuration(),
		Enabled:       enabled,
	}
}

func ConvertDynamicConfigToDynamicConfigV2PB(obj *config.DynamicConfig) *configv2.DynamicConfig {
	return &configv2.DynamicConfig{
		RoundDuration: obj.RoundDuration,
		Enabled:       &wrappers.BoolValue{Value: *obj.Enabled},
	}
}

func NewStaticConfigV2Controller(svc configservice.StaticServ) *StaticConfigV2Controller {
	return &StaticConfigV2Controller{svc: svc}
}

type StaticConfigV2Controller struct {
	svc configservice.StaticServ
	configv2.UnimplementedStaticConfigServiceServer
}

func (s StaticConfigV2Controller) Get(ctx context.Context, request *configv2.StaticConfigServiceGetRequest) (*configv2.StaticConfigServiceGetResponse, error) {
	staticConfig, err := s.svc.Get()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	ret, err := json.Marshal(staticConfig)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}

	return &configv2.StaticConfigServiceGetResponse{
		StaticConfig: string(ret),
	}, nil
}
