package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/host/v1/hostv1grpc"
	hostv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/host/v1"
	protov1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/proto/v1"
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HostV1Controller struct {
	svc    hostservice.Serv
	client *util.Store
	hostv1grpc.UnimplementedHostServiceServer
}

func (p HostV1Controller) GetByID(ctx context.Context, request *hostv1.GetByIDRequest) (*hostv1.GetByIDResponse, error) {
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
	return &hostv1.GetByIDResponse{Host: ConvertHostToHostV1Pb(hst)}, nil
}

func (p HostV1Controller) GetAll(ctx context.Context, _ *hostv1.GetAllRequest) (*hostv1.GetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	servcspb := make([]*hostv1.Host, 0, len(props))
	for i := range props {
		servcspb = append(servcspb, ConvertHostToHostV1Pb(props[i]))
	}
	return &hostv1.GetAllResponse{Hosts: servcspb}, nil
}

func (p HostV1Controller) Delete(ctx context.Context, request *hostv1.DeleteRequest) (*hostv1.DeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &hostv1.DeleteResponse{}, nil
}

func (p HostV1Controller) Store(ctx context.Context, request *hostv1.StoreRequest) (*hostv1.StoreResponse, error) {
	servcspb := request.GetHosts()
	props := make([]*host.Host, 0, len(servcspb))
	for i := range servcspb {
		hst, err := ConvertHostV1PBtoHost(false, servcspb[i])
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
	return &hostv1.StoreResponse{Ids: ids}, nil
}

func (p HostV1Controller) Update(ctx context.Context, request *hostv1.UpdateRequest) (*hostv1.UpdateResponse, error) {
	hst, err := ConvertHostV1PBtoHost(true, request.GetHost())
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
	return &hostv1.UpdateResponse{}, nil
}

func NewHostV1Controller(svc hostservice.Serv, client *util.Store) *HostV1Controller {
	return &HostV1Controller{svc: svc, client: client}
}

func ConvertHostV1PBtoHost(requireID bool, pb *hostv1.Host) (*host.Host, error) {
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

func ConvertHostToHostV1Pb(obj *host.Host) *hostv1.Host {
	var hstGrpID *protov1.UUID
	if obj.HostGroupID != nil {
		hstGrpID = &protov1.UUID{Value: obj.HostGroupID.String()}
	}

	var addressList *wrappers.StringValue
	if obj.AddressListRange != nil {
		addressList = &wrappers.StringValue{Value: *obj.AddressListRange}
	}

	return &hostv1.Host{
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
