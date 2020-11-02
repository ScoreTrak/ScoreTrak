package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	service2 "github.com/ScoreTrak/ScoreTrak/pkg/service/service_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicepb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServiceController struct {
	svc service2.Serv
}

func (p ServiceController) GetByID(ctx context.Context, request *servicepb.GetByIDRequest) (*servicepb.GetByIDResponse, error) {
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
	prop, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}

	return &servicepb.GetByIDResponse{Service: ConvertServiceToServicePb(prop)}, nil
}

func (p ServiceController) TestService(ctx context.Context, request *servicepb.TestServiceRequest) (*servicepb.TestServiceResponse, error) {
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
	chck, err := p.svc.TestService(ctx, uid)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Service failed to perform checking",
		)
	}

	return &servicepb.TestServiceResponse{Check: &checkpb.Check{
		ServiceId: &utilpb.UUID{Value: chck.ServiceID.String()},
		RoundId:   0,
		Log:       chck.Log,
		Err:       chck.Err,
		Passed:    &wrappers.BoolValue{Value: *chck.Passed},
	}}, err
}

func (p ServiceController) GetAll(ctx context.Context, request *servicepb.GetAllRequest) (*servicepb.GetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var servcspb []*servicepb.Service
	for i := range props {
		servcspb = append(servcspb, ConvertServiceToServicePb(props[i]))
	}
	return &servicepb.GetAllResponse{Services: servcspb}, nil
}

func (p ServiceController) Delete(ctx context.Context, request *servicepb.DeleteRequest) (*servicepb.DeleteResponse, error) {
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
	return &servicepb.DeleteResponse{}, nil
}

func (p ServiceController) Store(ctx context.Context, request *servicepb.StoreRequest) (*servicepb.StoreResponse, error) {
	servcspb := request.GetServices()
	var props []*service.Service
	for i := range servcspb {
		sr, err := ConvertServicePBtoService(false, servcspb[i])
		if err != nil {
			return nil, err
		}
		if sr.HostID == uuid.Nil || sr.ServiceGroupID == uuid.Nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Host ID And Service ID should not be nil",
			)
		}
		props = append(props, sr)
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
	return &servicepb.StoreResponse{Ids: ids}, nil
}

func (p ServiceController) Update(ctx context.Context, request *servicepb.UpdateRequest) (*servicepb.UpdateResponse, error) {
	srvpb := request.GetService()
	sr, err := ConvertServicePBtoService(true, srvpb)
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
	return &servicepb.UpdateResponse{}, nil
}

func NewServiceController(svc service2.Serv) *ServiceController {
	return &ServiceController{svc}
}

func ConvertServicePBtoService(requireID bool, pb *servicepb.Service) (*service.Service, error) {
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

	var weight *uint64
	if pb.GetWeight() != nil {
		weight = &pb.GetWeight().Value
	}

	var pointsBoost *uint64
	if pb.GetId() != nil {
		pointsBoost = &pb.GetPointsBoost().Value
	}

	var roundDelay *uint64
	if pb.GetRoundDelay() != nil {
		pointsBoost = &pb.GetRoundDelay().Value
	}

	var enabled *bool
	if pb.GetEnabled() != nil {
		enabled = &pb.GetEnabled().Value
	}

	var serviceGrpID uuid.UUID
	if pb.GetServiceGroupId() != nil {
		serviceGrpID, err = uuid.FromString(pb.GetServiceGroupId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to parse ID: %v", err,
			)
		}
	}

	var hostID uuid.UUID
	if pb.GetHostId() != nil {
		hostID, err = uuid.FromString(pb.GetHostId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to parse ID: %v", err,
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
		Enabled:        enabled,
		Properties:     nil,
		Checks:         nil,
	}, nil
}

func ConvertServiceToServicePb(obj *service.Service) *servicepb.Service {
	return &servicepb.Service{
		Id:             &utilpb.UUID{Value: obj.ID.String()},
		Name:           obj.Name,
		DisplayName:    obj.DisplayName,
		Weight:         &wrappers.UInt64Value{Value: *obj.Weight},
		PointsBoost:    &wrappers.UInt64Value{Value: *obj.PointsBoost},
		RoundUnits:     obj.RoundUnits,
		RoundDelay:     &wrappers.UInt64Value{Value: *obj.RoundDelay},
		ServiceGroupId: &utilpb.UUID{Value: obj.ServiceGroupID.String()},
		HostId:         &utilpb.UUID{Value: obj.HostID.String()},
		Enabled:        &wrappers.BoolValue{Value: *obj.Enabled},
		Properties:     nil,
		Checks:         nil,
	}
}
