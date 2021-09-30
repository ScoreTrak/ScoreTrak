package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	v1 "github.com/ScoreTrak/ScoreTrak/pkg/proto/proto/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func getErrorParser(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Resource Not Found: %v", err))
	}
	return status.Errorf(codes.Internal,
		fmt.Sprintf("Unknown internal error: %v", err))
}

func deleteErrorParser(err error) error {
	if errors.Is(err, &orm.NoRowsAffected{}) {
		return status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Resource Not Found: %v", err),
		)
	} else {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
}

type retrievableID interface {
	GetId() *v1.UUID
}

func extractUUID(r retrievableID) (uuid.UUID, error) {
	id := r.GetId()
	if id == nil {
		return uuid.UUID{}, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return uuid.UUID{}, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}
	return uid, nil
}

func extractUserClaim(ctx context.Context) *auth.UserClaims {
	if val, ok := ctx.Value(auth.KeyClaim).(*auth.UserClaims); ok && val != nil {
		return val
	}
	return nil
}

const (
	noPermissionsTo = "You do not have permissions to "
	genericErr      = "retrieve or update this resource"
	idNotSpecified  = "ID was not specified"
	changingUser    = "change this user"
	unableToParse   = "Unable to parse"
	unableToParseID = unableToParse + " ID"
)

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
