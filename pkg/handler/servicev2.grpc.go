package handler

import (
	"context"
	"fmt"
	servicev2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service/v2"

	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	service2 "github.com/ScoreTrak/ScoreTrak/pkg/service/serviceservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	checkv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/check/v1"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceV2Controller struct {
	svc    service2.Serv
	client *util.Store
	servicev2.UnimplementedServiceServiceServer
}

func (p ServiceV2Controller) GetByID(ctx context.Context, request *servicev2.ServiceServiceGetByIDRequest) (*servicev2.ServiceServiceGetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}

	claim := extractUserClaim(ctx)

	var serv *service.Service
	if claim.Role != user.Black {
		tID, prop, err := teamIDFromService(ctx, p.client, uid)
		if err != nil {
			return nil, getErrorParser(err)
		}
		if tID.String() != claim.TeamID {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+genericErr,
			)
		}
		serv = prop
	}

	if serv == nil {
		serv, err = p.svc.GetByID(ctx, uid)
		if err != nil {
			return nil, getErrorParser(err)
		}
	}

	return &servicev2.ServiceServiceGetByIDResponse{Service: ConvertServiceToServiceV2Pb(serv)}, nil
}

func (p ServiceV2Controller) TestService(ctx context.Context, request *servicev2.ServiceServiceTestServiceRequest) (*servicev2.ServiceServiceTestServiceResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	chck, err := p.svc.TestService(ctx, uid)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Service failed to perform checking",
		)
	}

	return &servicev2.ServiceServiceTestServiceResponse{Check: &checkv1.Check{
		ServiceId: &protov1.UUID{Value: chck.ServiceID.String()},
		RoundId:   0,
		Log:       chck.Log,
		Err:       chck.Err,
		Passed:    &wrappers.BoolValue{Value: *chck.Passed},
	}}, err
}

func (p ServiceV2Controller) GetAll(ctx context.Context, request *servicev2.ServiceServiceGetAllRequest) (*servicev2.ServiceServiceGetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	servcspb := make([]*servicev2.Service, 0, len(props))
	for i := range props {
		servcspb = append(servcspb, ConvertServiceToServiceV2Pb(props[i]))
	}
	return &servicev2.ServiceServiceGetAllResponse{Services: servcspb}, nil
}

func (p ServiceV2Controller) Delete(ctx context.Context, request *servicev2.ServiceServiceDeleteRequest) (*servicev2.ServiceServiceDeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &servicev2.ServiceServiceDeleteResponse{}, nil
}

func (p ServiceV2Controller) Store(ctx context.Context, request *servicev2.ServiceServiceStoreRequest) (*servicev2.ServiceServiceStoreResponse, error) {
	servcspb := request.GetServices()
	props := make([]*service.Service, 0, len(servcspb))
	for i := range servcspb {
		sr, err := ConvertServiceV2PBtoService(false, servcspb[i])
		if err != nil {
			return nil, err
		}
		if sr.ServiceGroupID == uuid.Nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Service Group ID should not be nil",
			)
		}

		if sr.HostID == uuid.Nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Host ID should not be nil",
			)
		}

		props = append(props, sr)
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
	return &servicev2.ServiceServiceStoreResponse{Ids: ids}, nil
}

func (p ServiceV2Controller) Update(ctx context.Context, request *servicev2.ServiceServiceUpdateRequest) (*servicev2.ServiceServiceUpdateResponse, error) {
	srvpb := request.GetService()
	sr, err := ConvertServiceV2PBtoService(true, srvpb)
	if err != nil {
		return nil, err
	}
	err = p.svc.Update(ctx, sr)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &servicev2.ServiceServiceUpdateResponse{}, nil
}

func NewServiceV2Controller(svc service2.Serv, client *util.Store) *ServiceV2Controller {
	return &ServiceV2Controller{svc: svc, client: client}
}

func ConvertServiceV2PBtoService(requireID bool, pb *servicev2.Service) (*service.Service, error) {
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

	var weight *uint64
	if pb.GetWeight() != nil {
		weight = &pb.GetWeight().Value
	}

	var pointsBoost *uint64
	if pb.GetPointsBoost() != nil {
		pointsBoost = &pb.GetPointsBoost().Value
	}

	var roundDelay *uint64
	if pb.GetRoundDelay() != nil {
		roundDelay = &pb.GetRoundDelay().Value
	}

	var pause *bool
	if pb.GetPause() != nil {
		pause = &pb.GetPause().Value
	}

	var hide *bool
	if pb.GetHide() != nil {
		hide = &pb.GetHide().Value
	}

	var serviceGrpID uuid.UUID
	if pb.GetServiceGroupId() != nil {
		serviceGrpID, err = uuid.FromString(pb.GetServiceGroupId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				unableToParseID+": %v", err,
			)
		}
	}

	var hostID uuid.UUID
	if pb.GetHostId() != nil {
		hostID, err = uuid.FromString(pb.GetHostId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				unableToParseID+": %v", err,
			)
		}
	}

	return &service.Service{
		ID:             id,
		Name:           pb.Name,
		DisplayName:    pb.DisplayName,
		Weight:         weight,
		PointsBoost:    pointsBoost,
		RoundUnits:     pb.RoundUnits,
		RoundDelay:     roundDelay,
		ServiceGroupID: serviceGrpID,
		HostID:         hostID,
		Pause:          pause,
		Hide:           hide,
		Properties:     nil,
		Checks:         nil,
	}, nil
}

func ConvertServiceToServiceV2Pb(obj *service.Service) *servicev2.Service {
	return &servicev2.Service{
		Id:             &protov1.UUID{Value: obj.ID.String()},
		Name:           obj.Name,
		DisplayName:    obj.DisplayName,
		Weight:         &wrappers.UInt64Value{Value: *obj.Weight},
		PointsBoost:    &wrappers.UInt64Value{Value: *obj.PointsBoost},
		RoundUnits:     obj.RoundUnits,
		RoundDelay:     &wrappers.UInt64Value{Value: *obj.RoundDelay},
		ServiceGroupId: &protov1.UUID{Value: obj.ServiceGroupID.String()},
		HostId:         &protov1.UUID{Value: obj.HostID.String()},
		Pause:          &wrappers.BoolValue{Value: *obj.Pause},
		Hide:           &wrappers.BoolValue{Value: *obj.Hide},
		Properties:     nil,
		Checks:         nil,
	}
}
