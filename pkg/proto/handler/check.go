package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/check_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/role"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckController struct {
	svc    check_service.Serv
	client *util.Store
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

	claim := extractUserClaim(ctx)

	var chk *check.Check
	if claim.Role != role.Black {
		tID, prop, err := teamIDFromCheck(ctx, c.client, roundID, uid)
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
		chk = prop
	}

	if chk == nil {
		chk, err = c.svc.GetByRoundServiceID(ctx, roundID, uid)
	}
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

	claim := extractUserClaim(ctx)

	if claim.Role != role.Black {
		tID, _, err := teamIDFromService(ctx, c.client, uid)
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

func NewCheckController(svc check_service.Serv, client *util.Store) *CheckController {
	return &CheckController{svc, client}
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

func teamIDFromProperty(ctx context.Context, c *util.Store, serviceID uuid.UUID, key string) (teamID uuid.UUID, property *property.Property, err error) {
	property, err = c.Property.GetByServiceIDKey(ctx, serviceID, key)
	if err != nil || property == nil {
		return
	}
	teamID, _, err = teamIDFromService(ctx, c, property.ServiceID)
	return
}

func teamIDFromCheck(ctx context.Context, c *util.Store, roundID uint64, serviceID uuid.UUID) (teamID uuid.UUID, check *check.Check, err error) {
	check, err = c.Check.GetByRoundServiceID(ctx, roundID, serviceID)
	if err != nil || check == nil {
		return
	}
	teamID, _, err = teamIDFromService(ctx, c, check.ServiceID)
	return
}

func teamIDFromService(ctx context.Context, c *util.Store, serviceID uuid.UUID) (teamID uuid.UUID, service *service.Service, err error) {
	service, err = c.Service.GetByID(ctx, serviceID)
	if err != nil || service == nil {
		return
	}
	teamID, _, err = teamIDFromHost(ctx, c, service.HostID)
	return
}

func teamIDFromHost(ctx context.Context, c *util.Store, hostID uuid.UUID) (teamID uuid.UUID, host *host.Host, err error) {
	host, err = c.Host.GetByID(ctx, hostID)
	if err != nil || host == nil {
		return
	}
	return host.TeamID, host, err
}
