package handler

import (
	"context"
	"fmt"
	host_groupv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host_group/v2"

	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgroupservice"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HostGroupV2Controller struct {
	svc hostgroupservice.Serv
	host_groupv2.UnimplementedHostGroupServiceServer
}

func (p HostGroupV2Controller) GetByID(ctx context.Context, request *host_groupv2.HostGroupServiceGetByIDRequest) (*host_groupv2.HostGroupServiceGetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	hst, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &host_groupv2.HostGroupServiceGetByIDResponse{HostGroup: ConvertHostGroupToHostGroupV2Pb(hst)}, nil
}

func (p HostGroupV2Controller) GetAll(ctx context.Context, request *host_groupv2.HostGroupServiceGetAllRequest) (*host_groupv2.HostGroupServiceGetAllResponse, error) {
	hostgrps, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	hostgrpspb := make([]*host_groupv2.HostGroup, 0, len(hostgrps))
	for i := range hostgrps {
		hostgrpspb = append(hostgrpspb, ConvertHostGroupToHostGroupV2Pb(hostgrps[i]))
	}
	return &host_groupv2.HostGroupServiceGetAllResponse{HostGroups: hostgrpspb}, nil
}

func (p HostGroupV2Controller) Delete(ctx context.Context, request *host_groupv2.HostGroupServiceDeleteRequest) (*host_groupv2.HostGroupServiceDeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &host_groupv2.HostGroupServiceDeleteResponse{}, nil
}

func (p HostGroupV2Controller) Store(ctx context.Context, request *host_groupv2.HostGroupServiceStoreRequest) (*host_groupv2.HostGroupServiceStoreResponse, error) {
	servcspb := request.GetHostGroups()
	props := make([]*hostgroup.HostGroup, 0, len(servcspb))
	for i := range servcspb {
		hstgrp, err := ConvertHostGroupV2PBtoHostGroup(false, servcspb[i])
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
	return &host_groupv2.HostGroupServiceStoreResponse{Ids: ids}, nil
}

func (p HostGroupV2Controller) Update(ctx context.Context, request *host_groupv2.HostGroupServiceUpdateRequest) (*host_groupv2.HostGroupServiceUpdateResponse, error) {
	hstgrp, err := ConvertHostGroupV2PBtoHostGroup(true, request.GetHostGroup())
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
	return &host_groupv2.HostGroupServiceUpdateResponse{}, nil
}

func NewHostGroupV2Controller(svc hostgroupservice.Serv) *HostGroupV2Controller {
	return &HostGroupV2Controller{svc: svc}
}

func ConvertHostGroupV2PBtoHostGroup(requireID bool, pb *host_groupv2.HostGroup) (*hostgroup.HostGroup, error) {
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

func ConvertHostGroupToHostGroupV2Pb(obj *hostgroup.HostGroup) *host_groupv2.HostGroup {
	return &host_groupv2.HostGroup{
		Id:    &protov1.UUID{Value: obj.ID.String()},
		Name:  obj.Name,
		Pause: &wrappers.BoolValue{Value: *obj.Pause},
		Hide:  &wrappers.BoolValue{Value: *obj.Hide},
		Hosts: nil,
	}
}
