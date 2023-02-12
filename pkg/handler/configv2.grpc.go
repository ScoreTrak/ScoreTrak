package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/config/v2/configv2grpc"
	configv2 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/config/v2"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configservice"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ConfigV2Controller struct {
	svc             configservice.Serv
	staticConfigSvc configservice.StaticServ
	configv2grpc.UnimplementedDynamicConfigServiceServer
}

func (p ConfigV2Controller) Get(ctx context.Context, _ *configv2.DynamicConfigServiceGetRequest) (*configv2.DynamicConfigServiceGetResponse, error) {
	cnf, err := p.svc.Get(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &configv2.DynamicConfigServiceGetResponse{DynamicConfig: ConvertDynamicConfigToDynamicConfigV2PB(cnf)}, nil
}

var ErrRoundDurationLargerThanMinRoundDuration = errors.New("round Duration should not be larger than MinRoundDuration")

func (p ConfigV2Controller) Update(ctx context.Context, request *configv2.DynamicConfigServiceUpdateRequest) (*configv2.DynamicConfigServiceUpdateResponse, error) {
	tmspb := request.GetDynamicConfig()
	var enabled *bool
	if tmspb.GetEnabled() != nil {
		enabled = &tmspb.GetEnabled().Value
	}

	// Get Static Config
	staticConfig, err := p.staticConfigSvc.Get()
	minTimeoutDuration := time.Duration(staticConfig.MinTimeoutDuration) * time.Second

	if tmspb.RoundDuration != 0 && tmspb.RoundDuration < uint64(minTimeoutDuration.Seconds()) {
		return nil, fmt.Errorf("%w, MinRoundDuration: %d", ErrRoundDurationLargerThanMinRoundDuration, uint64(minTimeoutDuration.Seconds()))
	}

	err = p.svc.Update(ctx, &config.DynamicConfig{
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

func NewConfigV2Controller(svc configservice.Serv, staticConfigSvc configservice.StaticServ) *ConfigV2Controller {
	return &ConfigV2Controller{svc: svc, staticConfigSvc: staticConfigSvc}
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
	configv2grpc.UnimplementedStaticConfigServiceServer
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
