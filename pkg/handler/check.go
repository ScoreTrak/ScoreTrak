package handler

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_service"
	checkpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/check/v1"
	utilpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/proto/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckController struct {
	svc    check_service.Serv
	client *util.Store
	checkpb.UnimplementedCheckServiceServer
}

func (c *CheckController) GetAllByRoundID(ctx context.Context, request *checkpb.GetAllByRoundIDRequest) (*checkpb.GetAllByRoundIDResponse, error) {
	roundID := request.GetRoundId()
	if roundID == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round"+idNotSpecified,
		)
	}

	chks, err := c.svc.GetAllByRoundID(ctx, roundID)
	if err != nil {
		return nil, getErrorParser(err)
	}
	chkspb := make([]*checkpb.Check, 0, len(chks))
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
			"Service"+idNotSpecified,
		)
	}
	roundID := request.GetRoundId()
	if roundID == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round"+idNotSpecified,
		)
	}
	uid, err := uuid.FromString(serviceID.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}

	claim := extractUserClaim(ctx)

	var chk *check.Check
	if claim.Role != user.Black {
		tID, prop, err := teamIDFromCheck(ctx, c.client, roundID, uid)
		if err != nil {
			return nil, getErrorParser(err)
		}
		if tID.String() != claim.TeamID {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+genericErr,
			)
		}
		chk = prop
	}

	if chk == nil {
		chk, err = c.svc.GetByRoundServiceID(ctx, roundID, uid)
	}
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &checkpb.GetByRoundServiceIDResponse{Check: ConvertCheckToCheckPb(chk)}, nil
}

func (c *CheckController) GetAllByServiceID(ctx context.Context, request *checkpb.GetAllByServiceIDRequest) (*checkpb.GetAllByServiceIDResponse, error) {
	serviceID := request.GetServiceId()
	if serviceID == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Service"+idNotSpecified,
		)
	}
	uid, err := uuid.FromString(serviceID.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}

	claim := extractUserClaim(ctx)

	if claim.Role != user.Black {
		tID, _, err := teamIDFromService(ctx, c.client, uid)
		if err != nil {
			return nil, getErrorParser(err)
		}
		if tID.String() != claim.TeamID {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+genericErr,
			)
		}
	}

	chks, err := c.svc.GetAllByServiceID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	chkspb := make([]*checkpb.Check, 0, len(chks))
	for i := range chks {
		chkspb = append(
			chkspb,
			ConvertCheckToCheckPb(chks[i]),
		)
	}
	return &checkpb.GetAllByServiceIDResponse{Checks: chkspb}, nil
}

func NewCheckController(svc check_service.Serv, client *util.Store) *CheckController {
	return &CheckController{svc: svc, client: client}
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
				unableToParseID+": %v", err,
			)
		}
	} else {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Service"+idNotSpecified,
		)
	}
	if pb.GetRoundId() == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round"+idNotSpecified,
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
