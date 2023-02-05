package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/service_group/v2/service_groupv2grpc"
	protov1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/proto/v1"
	service_groupv2 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/service_group/v2"
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceGroupV2Controller struct {
	svc servicegroupservice.Serv
	service_groupv2grpc.UnimplementedServiceGroupServiceServer
}

func (p ServiceGroupV2Controller) Redeploy(ctx context.Context, request *service_groupv2.ServiceGroupServiceRedeployRequest) (*service_groupv2.ServiceGroupServiceRedeployResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Redeploy(ctx, uid)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Failed to deploy workers: %v",
			err,
		)
	}
	return &service_groupv2.ServiceGroupServiceRedeployResponse{}, err
}

func (p ServiceGroupV2Controller) GetByID(ctx context.Context, request *service_groupv2.ServiceGroupServiceGetByIDRequest) (*service_groupv2.ServiceGroupServiceGetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	servgrp, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &service_groupv2.ServiceGroupServiceGetByIDResponse{ServiceGroup: ConvertServiceGroupToServiceGroupV2Pb(servgrp)}, nil
}

func (p ServiceGroupV2Controller) GetAll(ctx context.Context, request *service_groupv2.ServiceGroupServiceGetAllRequest) (*service_groupv2.ServiceGroupServiceGetAllResponse, error) {
	servgrps, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	servcspb := make([]*service_groupv2.ServiceGroup, 0, len(servgrps))

	for i := range servgrps {
		servcspb = append(servcspb, ConvertServiceGroupToServiceGroupV2Pb(servgrps[i]))
	}
	return &service_groupv2.ServiceGroupServiceGetAllResponse{ServiceGroups: servcspb}, nil
}

func (p ServiceGroupV2Controller) Delete(ctx context.Context, request *service_groupv2.ServiceGroupServiceDeleteRequest) (*service_groupv2.ServiceGroupServiceDeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &service_groupv2.ServiceGroupServiceDeleteResponse{}, nil
}

func (p ServiceGroupV2Controller) Store(ctx context.Context, request *service_groupv2.ServiceGroupServiceStoreRequest) (*service_groupv2.ServiceGroupServiceStoreResponse, error) {
	servcspb := request.GetServiceGroup()
	svg, err := ConvertServiceGroupV2PBtoServiceGroup(false, servcspb)
	if err != nil {
		return nil, err
	}
	err = p.svc.Store(ctx, svg)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &service_groupv2.ServiceGroupServiceStoreResponse{Id: &protov1.UUID{Value: svg.ID.String()}}, nil
}

func (p ServiceGroupV2Controller) Update(ctx context.Context, request *service_groupv2.ServiceGroupServiceUpdateRequest) (*service_groupv2.ServiceGroupServiceUpdateResponse, error) {
	svgrp := request.GetServiceGroup()
	svg, err := ConvertServiceGroupV2PBtoServiceGroup(true, svgrp)
	if err != nil {
		return nil, err
	}
	err = p.svc.Update(ctx, svg)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &service_groupv2.ServiceGroupServiceUpdateResponse{}, nil
}

func NewServiceGroupV2Controller(svc servicegroupservice.Serv) *ServiceGroupV2Controller {
	return &ServiceGroupV2Controller{svc: svc}
}

func ConvertServiceGroupV2PBtoServiceGroup(requireID bool, sg *service_groupv2.ServiceGroup) (*servicegroup.ServiceGroup, error) {
	var id uuid.UUID
	var err error
	if sg.GetId() != nil {
		id, err = uuid.FromString(sg.GetId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				unableToParseID+": %v", err,
			)
		}
	} else if requireID {
		return nil, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}
	var enabled *bool
	if sg.GetEnabled() != nil {
		enabled = &sg.GetEnabled().Value
	}
	return &servicegroup.ServiceGroup{
		ID:          id,
		Name:        sg.Name,
		DisplayName: sg.DisplayName,
		Enabled:     enabled,
		SkipHelper:  sg.SkipHelper,
		Label:       sg.Label,
		Services:    nil,
	}, nil
}

func ConvertServiceGroupToServiceGroupV2Pb(obj *servicegroup.ServiceGroup) *service_groupv2.ServiceGroup {
	return &service_groupv2.ServiceGroup{
		Id:          &protov1.UUID{Value: obj.ID.String()},
		Name:        obj.Name,
		DisplayName: obj.DisplayName,
		Enabled:     &wrappers.BoolValue{Value: *obj.Enabled},
		SkipHelper:  obj.SkipHelper,
		Label:       obj.Label,
		Services:    nil,
	}
}
