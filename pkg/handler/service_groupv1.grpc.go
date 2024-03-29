package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/service_group/v1/service_groupv1grpc"
	protov1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/proto/v1"
	service_groupv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/service_group/v1"
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceGroupV1Controller struct {
	svc servicegroupservice.Serv
	service_groupv1grpc.UnimplementedServiceGroupServiceServer
}

func (p ServiceGroupV1Controller) Redeploy(ctx context.Context, request *service_groupv1.RedeployRequest) (*service_groupv1.RedeployResponse, error) {
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
	return &service_groupv1.RedeployResponse{}, err
}

func (p ServiceGroupV1Controller) GetByID(ctx context.Context, request *service_groupv1.GetByIDRequest) (*service_groupv1.GetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	servgrp, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &service_groupv1.GetByIDResponse{ServiceGroup: ConvertServiceGroupToServiceGroupV1Pb(servgrp)}, nil
}

func (p ServiceGroupV1Controller) GetAll(ctx context.Context, request *service_groupv1.GetAllRequest) (*service_groupv1.GetAllResponse, error) {
	servgrps, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	servcspb := make([]*service_groupv1.ServiceGroup, 0, len(servgrps))

	for i := range servgrps {
		servcspb = append(servcspb, ConvertServiceGroupToServiceGroupV1Pb(servgrps[i]))
	}
	return &service_groupv1.GetAllResponse{ServiceGroups: servcspb}, nil
}

func (p ServiceGroupV1Controller) Delete(ctx context.Context, request *service_groupv1.DeleteRequest) (*service_groupv1.DeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &service_groupv1.DeleteResponse{}, nil
}

func (p ServiceGroupV1Controller) Store(ctx context.Context, request *service_groupv1.StoreRequest) (*service_groupv1.StoreResponse, error) {
	servcspb := request.GetServiceGroup()
	svg, err := ConvertServiceGroupV1PBtoServiceGroup(false, servcspb)
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
	return &service_groupv1.StoreResponse{Id: &protov1.UUID{Value: svg.ID.String()}}, nil
}

func (p ServiceGroupV1Controller) Update(ctx context.Context, request *service_groupv1.UpdateRequest) (*service_groupv1.UpdateResponse, error) {
	svgrp := request.GetServiceGroup()
	svg, err := ConvertServiceGroupV1PBtoServiceGroup(true, svgrp)
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
	return &service_groupv1.UpdateResponse{}, nil
}

func NewServiceGroupV1Controller(svc servicegroupservice.Serv) *ServiceGroupV1Controller {
	return &ServiceGroupV1Controller{svc: svc}
}

func ConvertServiceGroupV1PBtoServiceGroup(requireID bool, sg *service_groupv1.ServiceGroup) (*servicegroup.ServiceGroup, error) {
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

func ConvertServiceGroupToServiceGroupV1Pb(obj *servicegroup.ServiceGroup) *service_groupv1.ServiceGroup {
	return &service_groupv1.ServiceGroup{
		Id:          &protov1.UUID{Value: obj.ID.String()},
		Name:        obj.Name,
		DisplayName: obj.DisplayName,
		Enabled:     &wrappers.BoolValue{Value: *obj.Enabled},
		SkipHelper:  obj.SkipHelper,
		Label:       obj.Label,
		Services:    nil,
	}
}
