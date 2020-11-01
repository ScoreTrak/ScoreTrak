package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_grouppb"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HostGroupController struct {
	svc service.Serv
}

func (p HostGroupController) GetByID(ctx context.Context, request *host_grouppb.GetByIDRequest) (*host_grouppb.GetByIDResponse, error) {
	id := request.GetId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ID was not specified",
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	hst, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &host_grouppb.GetByIDResponse{HostGroup: ConvertHostGroupToHostGroupPb(hst)}, nil
}

func (p HostGroupController) GetAll(ctx context.Context, request *host_grouppb.GetAllRequest) (*host_grouppb.GetAllResponse, error) {
	hostgrps, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var hostgrpspb []*host_grouppb.HostGroup
	for i := range hostgrps {
		hostgrpspb = append(hostgrpspb, ConvertHostGroupToHostGroupPb(hostgrps[i]))
	}
	return &host_grouppb.GetAllResponse{HostGroups: hostgrpspb}, nil
}

func (p HostGroupController) Delete(ctx context.Context, request *host_grouppb.DeleteRequest) (*host_grouppb.DeleteResponse, error) {
	id := request.GetId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ID was not specified",
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &host_grouppb.DeleteResponse{}, nil
}

func (p HostGroupController) Store(ctx context.Context, request *host_grouppb.StoreRequest) (*host_grouppb.StoreResponse, error) {
	servcspb := request.GetHostGroups()
	var props []*host_group.HostGroup
	for i := range servcspb {
		hstgrp, err := ConvertHostGroupPBtoHostGroup(false, servcspb[i])
		if err != nil {
			return nil, err
		}
		props = append(props, hstgrp)
	}
	err := p.svc.Store(ctx, props)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	var ids []*utilpb.UUID
	for i := range props {
		ids = append(ids, &utilpb.UUID{Value: props[i].ID.String()})
	}
	return &host_grouppb.StoreResponse{Ids: ids}, nil
}

func (p HostGroupController) Update(ctx context.Context, request *host_grouppb.UpdateRequest) (*host_grouppb.UpdateResponse, error) {
	hstgrp, err := ConvertHostGroupPBtoHostGroup(true, request.GetHostGroup())
	if err != nil {
		return nil, err
	}
	err = p.svc.Update(ctx, hstgrp)
	if err != nil {

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &host_grouppb.UpdateResponse{}, nil
}

func NewHostGroupController(svc service.Serv) *HostGroupController {
	return &HostGroupController{svc}
}

func ConvertHostGroupPBtoHostGroup(requireID bool, pb *host_grouppb.HostGroup) (*host_group.HostGroup, error) {
	var id uuid.UUID
	var err error
	if pb.GetId() != nil {
		id, err = uuid.FromString(pb.GetId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to parse ID: %v", err,
			)
		}
	} else if requireID {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ID was not specified",
		)
	}
	var enabled *bool
	if pb.GetEnabled() != nil {
		enabled = &pb.GetEnabled().Value
	}
	return &host_group.HostGroup{
		ID:      id,
		Name:    pb.GetName(),
		Enabled: enabled,
		Hosts:   nil,
	}, nil
}

func ConvertHostGroupToHostGroupPb(obj *host_group.HostGroup) *host_grouppb.HostGroup {
	return &host_grouppb.HostGroup{
		Id:      &utilpb.UUID{Value: obj.ID.String()},
		Name:    obj.Name,
		Enabled: &wrappers.BoolValue{Value: *obj.Enabled},
		Hosts:   nil,
	}
}
