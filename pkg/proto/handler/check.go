package handler

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckController struct {
	svc check_service.Serv
}

func (c *CheckController) GetAllByRoundID(ctx context.Context, request *checkpb.GetAllByRoundIDRequest) (*checkpb.GetAllByRoundIDResponse, error) {
	roundID := request.GetRoundId()
	if roundID == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round ID was not specified",
		)
	}

	chks, err := c.svc.GetAllByRoundID(ctx, roundID)
	if err != nil {
		return nil, getErrorParser(err)

	}
	var chkspb []*checkpb.Check
	for i := range chks {
		chkspb = append(
			chkspb,
			ConvertCheckToCheckPb(chks[i]),
		)
	}
	return &checkpb.GetAllByRoundIDResponse{Checks: chkspb}, nil
}

func (c *CheckController) GetByRoundServiceID(ctx context.Context, request *checkpb.GetByRoundServiceIDRequest) (*checkpb.GetByRoundServiceIDResponse, error) {
	serviceID := request.GetServiceId()
	if serviceID == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Service ID was not specified",
		)
	}
	roundID := request.GetRoundId()
	if roundID == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round ID was not specified",
		)
	}
	uid, err := uuid.FromString(serviceID.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	chk, err := c.svc.GetByRoundServiceID(ctx, roundID, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &checkpb.GetByRoundServiceIDResponse{Checks: ConvertCheckToCheckPb(chk)}, nil
}

func (c *CheckController) GetAllByServiceID(ctx context.Context, request *checkpb.GetAllByServiceIDRequest) (*checkpb.GetAllByServiceIDResponse, error) {
	serviceID := request.GetServiceId()
	if serviceID == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Service ID was not specified",
		)
	}
	uid, err := uuid.FromString(serviceID.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	chks, err := c.svc.GetAllByServiceID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var chkspb []*checkpb.Check
	for i := range chks {
		chkspb = append(
			chkspb,
			ConvertCheckToCheckPb(chks[i]),
		)
	}
	return &checkpb.GetAllByServiceIDResponse{Checks: chkspb}, nil
}

func NewCheckController(svc check_service.Serv) *CheckController {
	return &CheckController{svc}
}

func ConvertCheckToCheckPb(obj *check.Check) *checkpb.Check {
	return &checkpb.Check{
		ServiceId: &utilpb.UUID{Value: obj.ServiceID.String()},
		RoundId:   obj.RoundID,
		Log:       obj.Log,
		Err:       obj.Err,
		Passed:    &wrappers.BoolValue{Value: *obj.Passed},
	}
}

func ConvertCheckPBtoCheck(pb *checkpb.Check) (*check.Check, error) {
	var sID uuid.UUID
	var err error
	if pb.GetServiceId() != nil {
		sID, err = uuid.FromString(pb.GetServiceId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to parse ID: %v", err,
			)
		}
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Service ID was not specified",
		)
	}
	if pb.GetRoundId() == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round ID was not specified",
		)
	}
	var passed *bool
	if pb.GetPassed() != nil {
		passed = &pb.GetPassed().Value
	}
	return &check.Check{
		ServiceID: sID,
		RoundID:   pb.GetRoundId(),
		Log:       pb.Log,
		Err:       pb.Err,
		Passed:    passed,
	}, nil
}
