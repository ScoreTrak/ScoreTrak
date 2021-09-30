package handler

import (
	"context"
	"fmt"

	utilpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/proto/v1"
	service_grouppb "github.com/ScoreTrak/ScoreTrak/pkg/proto/service_group/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceGroupController struct {
	svc servicegroupservice.Serv
	service_grouppb.UnimplementedServiceGroupServiceServer
}

func (p ServiceGroupController) Redeploy(ctx context.Context, request *service_grouppb.RedeployRequest) (*service_grouppb.RedeployResponse, error) {
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
	return &service_grouppb.RedeployResponse{}, err
}

func (p ServiceGroupController) GetByID(ctx context.Context, request *service_grouppb.GetByIDRequest) (*service_grouppb.GetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	servgrp, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &service_grouppb.GetByIDResponse{ServiceGroup: ConvertServiceGroupToServiceGroupPb(servgrp)}, nil
}

func (p ServiceGroupController) GetAll(ctx context.Context, request *service_grouppb.GetAllRequest) (*service_grouppb.GetAllResponse, error) {
	servgrps, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	servcspb := make([]*service_grouppb.ServiceGroup, 0, len(servgrps))

	for i := range servgrps {
		servcspb = append(servcspb, ConvertServiceGroupToServiceGroupPb(servgrps[i]))
	}
	return &service_grouppb.GetAllResponse{ServiceGroups: servcspb}, nil
}

func (p ServiceGroupController) Delete(ctx context.Context, request *service_grouppb.DeleteRequest) (*service_grouppb.DeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &service_grouppb.DeleteResponse{}, nil
}

func (p ServiceGroupController) Store(ctx context.Context, request *service_grouppb.StoreRequest) (*service_grouppb.StoreResponse, error) {
	servcspb := request.GetServiceGroup()
	svg, err := ConvertServiceGroupPBtoServiceGroup(false, servcspb)
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
	return &service_grouppb.StoreResponse{Id: &utilpb.UUID{Value: svg.ID.String()}}, nil
}

func (p ServiceGroupController) Update(ctx context.Context, request *service_grouppb.UpdateRequest) (*service_grouppb.UpdateResponse, error) {
	svgrp := request.GetServiceGroup()
	svg, err := ConvertServiceGroupPBtoServiceGroup(true, svgrp)
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
	return &service_grouppb.UpdateResponse{}, nil
}

func NewServiceGroupController(svc servicegroupservice.Serv) *ServiceGroupController {
	return &ServiceGroupController{svc: svc}
}

func ConvertServiceGroupPBtoServiceGroup(requireID bool, sg *service_grouppb.ServiceGroup) (*servicegroup.ServiceGroup, error) {
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

func ConvertServiceGroupToServiceGroupPb(obj *servicegroup.ServiceGroup) *service_grouppb.ServiceGroup {
	return &service_grouppb.ServiceGroup{
		Id:          &utilpb.UUID{Value: obj.ID.String()},
		Name:        obj.Name,
		DisplayName: obj.DisplayName,
		Enabled:     &wrappers.BoolValue{Value: *obj.Enabled},
		SkipHelper:  obj.SkipHelper,
		Label:       obj.Label,
		Services:    nil,
	}
}
