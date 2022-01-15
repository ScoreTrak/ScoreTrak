package handler

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	checkv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/check/v1"
	protov1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckController struct {
	svc    checkservice.Serv
	client *util.Store
	checkv1.UnimplementedCheckServiceServer
}

func (c *CheckController) GetAllByRoundID(ctx context.Context, request *checkv1.GetAllByRoundIDRequest) (*checkv1.GetAllByRoundIDResponse, error) {
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
	chkspb := make([]*checkv1.Check, 0, len(chks))
	for i := range chks {
		chkspb = append(
			chkspb,
			ConvertCheckToCheckPb(chks[i]),
		)
	}
	return &checkv1.GetAllByRoundIDResponse{Checks: chkspb}, nil
}

func (c *CheckController) GetByRoundServiceID(ctx context.Context, request *checkv1.GetByRoundServiceIDRequest) (*checkv1.GetByRoundServiceIDResponse, error) {
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
	return &checkv1.GetByRoundServiceIDResponse{Check: ConvertCheckToCheckPb(chk)}, nil
}

func (c *CheckController) GetAllByServiceID(ctx context.Context, request *checkv1.GetAllByServiceIDRequest) (*checkv1.GetAllByServiceIDResponse, error) {
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
	chkspb := make([]*checkv1.Check, 0, len(chks))
	for i := range chks {
		chkspb = append(
			chkspb,
			ConvertCheckToCheckPb(chks[i]),
		)
	}
	return &checkv1.GetAllByServiceIDResponse{Checks: chkspb}, nil
}

func NewCheckController(svc checkservice.Serv, client *util.Store) *CheckController {
	return &CheckController{svc: svc, client: client}
}

func ConvertCheckToCheckPb(obj *check.Check) *checkv1.Check {
	return &checkv1.Check{
		ServiceId: &protov1.UUID{Value: obj.ServiceID.String()},
		RoundId:   obj.RoundID,
		Log:       obj.Log,
		Err:       obj.Err,
		Passed:    &wrappers.BoolValue{Value: *obj.Passed},
	}
}

func ConvertCheckPBtoCheck(pb *checkv1.Check) (*check.Check, error) {
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
