package handler

import (
	"context"
	checkv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/check/v2"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckV2Controller struct {
	svc    checkservice.Serv
	client *util.Store
	checkv2.UnimplementedCheckServiceServer
}

func (c *CheckV2Controller) GetAllByRoundID(ctx context.Context, request *checkv2.CheckServiceGetAllByRoundIDRequest) (*checkv2.CheckServiceGetAllByRoundIDResponse, error) {
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
	chkspb := make([]*checkv2.Check, 0, len(chks))
	for i := range chks {
		chkspb = append(
			chkspb,
			ConvertCheckToCheckV2Pb(chks[i]),
		)
	}
	return &checkv2.CheckServiceGetAllByRoundIDResponse{Checks: chkspb}, nil
}

func (c *CheckV2Controller) GetByRoundServiceID(ctx context.Context, request *checkv2.CheckServiceGetByRoundServiceIDRequest) (*checkv2.CheckServiceGetByRoundServiceIDResponse, error) {
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
	return &checkv2.CheckServiceGetByRoundServiceIDResponse{Check: ConvertCheckToCheckV2Pb(chk)}, nil
}

func (c *CheckV2Controller) GetAllByServiceID(ctx context.Context, request *checkv2.CheckServiceGetAllByServiceIDRequest) (*checkv2.CheckServiceGetAllByServiceIDResponse, error) {
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
	chkspb := make([]*checkv2.Check, 0, len(chks))
	for i := range chks {
		chkspb = append(
			chkspb,
			ConvertCheckToCheckV2Pb(chks[i]),
		)
	}
	return &checkv2.CheckServiceGetAllByServiceIDResponse{Checks: chkspb}, nil
}

func NewCheckV2Controller(svc checkservice.Serv, client *util.Store) *CheckV2Controller {
	return &CheckV2Controller{svc: svc, client: client}
}

func ConvertCheckToCheckV2Pb(obj *check.Check) *checkv2.Check {
	return &checkv2.Check{
		ServiceId: &protov1.UUID{Value: obj.ServiceID.String()},
		RoundId:   obj.RoundID,
		Log:       obj.Log,
		Err:       obj.Err,
		Passed:    &wrappers.BoolValue{Value: *obj.Passed},
	}
}

func ConvertCheckV2PBtoCheck(pb *checkv2.Check) (*check.Check, error) {
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
