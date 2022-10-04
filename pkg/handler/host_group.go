package handler

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgroupservice"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	host_groupv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host_group/v1"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HostGroupController struct {
	svc hostgroupservice.Serv
	host_groupv1.UnimplementedHostGroupServiceServer
}

func (p HostGroupController) GetByID(ctx context.Context, request *host_groupv1.GetByIDRequest) (*host_groupv1.GetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	hst, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &host_groupv1.GetByIDResponse{HostGroup: ConvertHostGroupToHostGroupPb(hst)}, nil
}

func (p HostGroupController) GetAll(ctx context.Context, request *host_groupv1.GetAllRequest) (*host_groupv1.GetAllResponse, error) {
	hostgrps, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	hostgrpspb := make([]*host_groupv1.HostGroup, 0, len(hostgrps))
	for i := range hostgrps {
		hostgrpspb = append(hostgrpspb, ConvertHostGroupToHostGroupPb(hostgrps[i]))
	}
	return &host_groupv1.GetAllResponse{HostGroups: hostgrpspb}, nil
}

func (p HostGroupController) Delete(ctx context.Context, request *host_groupv1.DeleteRequest) (*host_groupv1.DeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &host_groupv1.DeleteResponse{}, nil
}

func (p HostGroupController) Store(ctx context.Context, request *host_groupv1.StoreRequest) (*host_groupv1.StoreResponse, error) {
	servcspb := request.GetHostGroups()
	props := make([]*hostgroup.HostGroup, 0, len(servcspb))
	for i := range servcspb {
		hstgrp, err := ConvertHostGroupPBtoHostGroup(false, servcspb[i])
		if err != nil {
			return nil, err
		}
		props = append(props, hstgrp)
	}
	if err := p.svc.Store(ctx, props); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	ids := make([]*protov1.UUID, 0, len(props))
	for i := range props {
		ids = append(ids, &protov1.UUID{Value: props[i].ID.String()})
	}
	return &host_groupv1.StoreResponse{Ids: ids}, nil
}

func (p HostGroupController) Update(ctx context.Context, request *host_groupv1.UpdateRequest) (*host_groupv1.UpdateResponse, error) {
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
	return &host_groupv1.UpdateResponse{}, nil
}

func NewHostGroupController(svc hostgroupservice.Serv) *HostGroupController {
	return &HostGroupController{svc: svc}
}

func ConvertHostGroupPBtoHostGroup(requireID bool, pb *host_groupv1.HostGroup) (*hostgroup.HostGroup, error) {
	var id uuid.UUID
	var err error
	if pb.GetId() != nil {
		id, err = uuid.FromString(pb.GetId().GetValue())
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
	var pause *bool
	if pb.GetPause() != nil {
		pause = &pb.GetPause().Value
	}

	var hide *bool
	if pb.GetHide() != nil {
		hide = &pb.GetHide().Value
	}
	return &hostgroup.HostGroup{
		ID:    id,
		Name:  pb.GetName(),
		Pause: pause,
		Hide:  hide,
		Hosts: nil,
	}, nil
}

func ConvertHostGroupToHostGroupPb(obj *hostgroup.HostGroup) *host_groupv1.HostGroup {
	return &host_groupv1.HostGroup{
		Id:    &protov1.UUID{Value: obj.ID.String()},
		Name:  obj.Name,
		Pause: &wrappers.BoolValue{Value: *obj.Pause},
		Hide:  &wrappers.BoolValue{Value: *obj.Hide},
		Hosts: nil,
	}
}
