package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configservice"
	"github.com/golang/protobuf/ptypes/wrappers"
	configv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/config/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConfigV1Controller struct {
	svc configservice.Serv
	configv1.UnimplementedDynamicConfigServiceServer
}

func (p ConfigV1Controller) Get(ctx context.Context, _ *configv1.GetRequest) (*configv1.GetResponse, error) {
	cnf, err := p.svc.Get(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &configv1.GetResponse{DynamicConfig: ConvertDynamicConfigToDynamicConfigV1PB(cnf)}, nil
}

func (p ConfigV1Controller) Update(ctx context.Context, request *configv1.UpdateRequest) (*configv1.UpdateResponse, error) {
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
	return &configv1.UpdateResponse{}, nil
}

func NewConfigV1Controller(svc configservice.Serv) *ConfigV1Controller {
	return &ConfigV1Controller{svc: svc}
}

func ConvertDynamicConfigV1PBToDynamicConfig(pb *configv1.DynamicConfig) *config.DynamicConfig {
	var enabled *bool
	if pb.GetEnabled() != nil {
		enabled = &pb.GetEnabled().Value
	}
	return &config.DynamicConfig{
		RoundDuration: pb.GetRoundDuration(),
		Enabled:       enabled,
	}
}

func ConvertDynamicConfigToDynamicConfigV1PB(obj *config.DynamicConfig) *configv1.DynamicConfig {
	return &configv1.DynamicConfig{
		RoundDuration: obj.RoundDuration,
		Enabled:       &wrappers.BoolValue{Value: *obj.Enabled},
	}
}

func NewStaticConfigV1Controller(svc configservice.StaticServ) *StaticConfigV1Controller {
	return &StaticConfigV1Controller{svc: svc}
}

type StaticConfigV1Controller struct {
	svc configservice.StaticServ
	configv1.UnimplementedStaticConfigServiceServer
}

func (s StaticConfigV1Controller) Get(ctx context.Context, request *configv1.GetStaticConfigRequest) (*configv1.GetStaticConfigResponse, error) {
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

	return &configv1.GetStaticConfigResponse{
		StaticConfig: string(ret),
	}, nil
}
