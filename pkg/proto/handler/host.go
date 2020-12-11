package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/host_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/role"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HostController struct {
	svc    host_service.Serv
	client *util.Store
}

func (p HostController) GetByID(ctx context.Context, request *hostpb.GetByIDRequest) (*hostpb.GetByIDResponse, error) {
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

	claim := extractUserClaim(ctx)
	var hst *host.Host
	if claim.Role != role.Black {
		tID, prop, err := teamIDFromHost(ctx, p.client, uid)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Unabkle to validate resource. Err: %v", err),
			)
		}
		if tID.String() != claim.TeamID {
			return nil, status.Errorf(
				codes.PermissionDenied,
				fmt.Sprintf("You do not have permissions to retreive or update this resource"),
			)
		}
		hst = prop
	}

	if hst == nil {
		hst, err = p.svc.GetByID(ctx, uid)
		if err != nil {
			return nil, getErrorParser(err)
		}
	}
	return &hostpb.GetByIDResponse{Host: ConvertHostToHostPb(hst)}, nil
}

func (p HostController) GetAll(ctx context.Context, request *hostpb.GetAllRequest) (*hostpb.GetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var servcspb []*hostpb.Host
	for i := range props {
		servcspb = append(servcspb, ConvertHostToHostPb(props[i]))
	}
	return &hostpb.GetAllResponse{Hosts: servcspb}, nil
}

func (p HostController) Delete(ctx context.Context, request *hostpb.DeleteRequest) (*hostpb.DeleteResponse, error) {
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
	return &hostpb.DeleteResponse{}, nil
}

func (p HostController) Store(ctx context.Context, request *hostpb.StoreRequest) (*hostpb.StoreResponse, error) {
	servcspb := request.GetHosts()
	var props []*host.Host
	for i := range servcspb {
		hst, err := ConvertHostPBtoHost(false, servcspb[i])
		if err != nil {
			return nil, err
		}
		if hst.TeamID == uuid.Nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Team ID should not be nil",
			)
		}
		props = append(props, hst)
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
	return &hostpb.StoreResponse{Ids: ids}, nil
}

func (p HostController) Update(ctx context.Context, request *hostpb.UpdateRequest) (*hostpb.UpdateResponse, error) {
	hst, err := ConvertHostPBtoHost(true, request.GetHost())
	if err != nil {
		return nil, err
	}

	claim := extractUserClaim(ctx)

	if claim.Role != role.Black {
		tID, prop, err := teamIDFromHost(ctx, p.client, hst.ID)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Unabkle to validate resource. Err: %v", err),
			)
		}
		if tID.String() != claim.TeamID || !*prop.EditHost {
			return nil, status.Errorf(
				codes.PermissionDenied,
				fmt.Sprintf("You do not have permissions to retreive or update this resource"),
			)
		}
		hst = &host.Host{Address: prop.Address, ID: hst.ID}
	}

	err = p.svc.Update(ctx, hst)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &hostpb.UpdateResponse{}, nil
}

func NewHostController(svc host_service.Serv, client *util.Store) *HostController {
	return &HostController{svc, client}
}

func ConvertHostPBtoHost(requireID bool, pb *hostpb.Host) (*host.Host, error) {
	var err error
	var id uuid.UUID
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

	var editHost *bool
	if pb.GetEditHost() != nil {
		editHost = &pb.GetEditHost().Value
	}

	var address *string
	if pb.GetAddress() != "" {
		addr := pb.GetAddress()
		address = &addr
	}

	var hostGrpID *uuid.UUID
	if pb.GetHostGroupId() != nil {
		uid, err := uuid.FromString(pb.GetHostGroupId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to parse ID: %v", err,
			)
		}
		hostGrpID = &uid
	}

	var teamID uuid.UUID
	if pb.GetTeamId() != nil {
		teamID, err = uuid.FromString(pb.GetTeamId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to parse ID: %v", err,
			)
		}
	}
	return &host.Host{
		ID:          id,
		Address:     address,
		HostGroupID: hostGrpID,
		TeamID:      teamID,
		Enabled:     enabled,
		EditHost:    editHost,
		Services:    nil,
	}, nil
}

func ConvertHostToHostPb(obj *host.Host) *hostpb.Host {

	var hstGrpID *utilpb.UUID
	if obj.HostGroupID != nil {
		hstGrpID = &utilpb.UUID{Value: obj.HostGroupID.String()}
	}

	return &hostpb.Host{
		Id:          &utilpb.UUID{Value: obj.ID.String()},
		Address:     *obj.Address,
		HostGroupId: hstGrpID,
		TeamId:      &utilpb.UUID{Value: obj.TeamID.String()},
		Enabled:     &wrappers.BoolValue{Value: *obj.Enabled},
		EditHost:    &wrappers.BoolValue{Value: *obj.EditHost},
		Services:    nil,
	}
}
