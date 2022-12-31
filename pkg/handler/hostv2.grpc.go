package handler

import (
	"context"
	"fmt"
	hostv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host/v2"

	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HostV2Controller struct {
	svc    hostservice.Serv
	client *util.Store
	hostv2.UnimplementedHostServiceServer
}

func (p HostV2Controller) GetByID(ctx context.Context, request *hostv2.HostServiceGetByIDRequest) (*hostv2.HostServiceGetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}

	claim := extractUserClaim(ctx)
	var hst *host.Host
	if claim.Role != user.Black {
		tID, prop, err := teamIDFromHost(ctx, p.client, uid)
		if err != nil {
			return nil, getErrorParser(err)
		}
		if tID.String() != claim.TeamID {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+genericErr,
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
	return &hostv2.HostServiceGetByIDResponse{Host: ConvertHostToHostV2Pb(hst)}, nil
}

func (p HostV2Controller) GetAll(ctx context.Context, _ *hostv2.HostServiceGetAllRequest) (*hostv2.HostServiceGetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	servcspb := make([]*hostv2.Host, 0, len(props))
	for i := range props {
		servcspb = append(servcspb, ConvertHostToHostV2Pb(props[i]))
	}
	return &hostv2.HostServiceGetAllResponse{Hosts: servcspb}, nil
}

func (p HostV2Controller) Delete(ctx context.Context, request *hostv2.HostServiceDeleteRequest) (*hostv2.HostServiceDeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &hostv2.HostServiceDeleteResponse{}, nil
}

func (p HostV2Controller) Store(ctx context.Context, request *hostv2.HostServiceStoreRequest) (*hostv2.HostServiceStoreResponse, error) {
	servcspb := request.GetHosts()
	props := make([]*host.Host, 0, len(servcspb))
	for i := range servcspb {
		hst, err := ConvertHostV2PBtoHost(false, servcspb[i])
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
	return &hostv2.HostServiceStoreResponse{Ids: ids}, nil
}

func (p HostV2Controller) Update(ctx context.Context, request *hostv2.HostServiceUpdateRequest) (*hostv2.HostServiceUpdateResponse, error) {
	hst, err := ConvertHostV2PBtoHost(true, request.GetHost())
	if err != nil {
		return nil, err
	}

	claim := extractUserClaim(ctx)

	if claim.Role != user.Black {
		tID, prop, err := teamIDFromHost(ctx, p.client, hst.ID)
		if err != nil {
			return nil, getErrorParser(err)
		}
		if tID.String() != claim.TeamID || !*prop.EditHost {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+genericErr,
			)
		}
		hst = &host.Host{Address: hst.Address, ID: hst.ID}
	}

	err = p.svc.Update(ctx, hst)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &hostv2.HostServiceUpdateResponse{}, nil
}

func NewHostV2Controller(svc hostservice.Serv, client *util.Store) *HostV2Controller {
	return &HostV2Controller{svc: svc, client: client}
}

func ConvertHostV2PBtoHost(requireID bool, pb *hostv2.Host) (*host.Host, error) {
	var err error
	var id uuid.UUID
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

	var editHost *bool
	if pb.GetEditHost() != nil {
		editHost = &pb.GetEditHost().Value
	}

	var address string
	if pb.GetAddress() != "" {
		address = pb.GetAddress()
	}

	var hostGrpID *uuid.UUID
	if pb.GetHostGroupId() != nil {
		uid, err := uuid.FromString(pb.GetHostGroupId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				unableToParseID+": %v", err,
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
				unableToParseID+": %v", err,
			)
		}
	}

	var addressList *string
	if pb.GetAddressListRange() != nil {
		addressList = &pb.GetAddressListRange().Value
	}

	return &host.Host{
		AddressListRange: addressList,
		ID:               id,
		Address:          address,
		HostGroupID:      hostGrpID,
		TeamID:           teamID,
		Pause:            pause,
		Hide:             hide,
		EditHost:         editHost,
		Services:         nil,
	}, nil
}

func ConvertHostToHostV2Pb(obj *host.Host) *hostv2.Host {
	var hstGrpID *protov1.UUID
	if obj.HostGroupID != nil {
		hstGrpID = &protov1.UUID{Value: obj.HostGroupID.String()}
	}

	var addressList *wrappers.StringValue
	if obj.AddressListRange != nil {
		addressList = &wrappers.StringValue{Value: *obj.AddressListRange}
	}

	return &hostv2.Host{
		Id:               &protov1.UUID{Value: obj.ID.String()},
		Address:          obj.Address,
		HostGroupId:      hstGrpID,
		TeamId:           &protov1.UUID{Value: obj.TeamID.String()},
		Pause:            &wrappers.BoolValue{Value: *obj.Pause},
		Hide:             &wrappers.BoolValue{Value: *obj.Hide},
		EditHost:         &wrappers.BoolValue{Value: *obj.EditHost},
		Services:         nil,
		AddressListRange: addressList,
	}
}
