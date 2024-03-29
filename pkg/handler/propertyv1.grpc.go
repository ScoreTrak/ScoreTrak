package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/property/v1/propertyv1grpc"
	propertyv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/property/v1"
	protov1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/proto/v1"
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PropertyV1Controller struct {
	svc    propertyservice.Serv
	client *util.Store
	propertyv1grpc.UnimplementedPropertyServiceServer
}

func (p PropertyV1Controller) GetAll(ctx context.Context, _ *propertyv1.GetAllRequest) (*propertyv1.GetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	propspb := make([]*propertyv1.Property, 0, len(props))
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyV1Pb(props[i]))
	}
	return &propertyv1.GetAllResponse{Properties: propspb}, nil
}

func (p PropertyV1Controller) Delete(ctx context.Context, request *propertyv1.DeleteRequest) (*propertyv1.DeleteResponse, error) {
	id := request.GetServiceId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}
	if request.GetKey() == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Key was not specified",
		)
	}
	err = p.svc.Delete(ctx, uid, request.GetKey())
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &propertyv1.DeleteResponse{}, nil
}

func (p PropertyV1Controller) Store(ctx context.Context, request *propertyv1.StoreRequest) (*propertyv1.StoreResponse, error) {
	propspb := request.GetProperties()
	props := make([]*property.Property, 0, len(propspb))
	for i := range propspb {
		prop, err := ConvertPropertyV1PBtoProperty(propspb[i])
		if err != nil {
			return nil, err
		}
		props = append(props, prop)
	}

	if err := p.svc.Store(ctx, props); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &propertyv1.StoreResponse{}, nil
}

func (p PropertyV1Controller) Update(ctx context.Context, request *propertyv1.UpdateRequest) (*propertyv1.UpdateResponse, error) {
	pr, err := ConvertPropertyV1PBtoProperty(request.Property)
	if err != nil {
		return nil, err
	}

	claim := extractUserClaim(ctx)
	if claim.Role != user.Black {
		tID, prop, err := teamIDFromProperty(ctx, p.client, pr.ServiceID, pr.Key)
		if err != nil {
			return nil, getErrorParser(err)
		}
		if tID.String() != claim.TeamID || prop.Status != property.Edit {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+genericErr,
			)
		}
		pr = &property.Property{Value: pr.Value, ServiceID: prop.ServiceID, Key: prop.Key}
	}

	err = p.svc.Update(ctx, pr)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &propertyv1.UpdateResponse{}, nil
}

func (p PropertyV1Controller) GetByServiceIDKey(ctx context.Context, request *propertyv1.GetByServiceIDKeyRequest) (*propertyv1.GetByServiceIDKeyResponse, error) {
	id := request.GetServiceId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}

	claim := extractUserClaim(ctx)

	var chk *property.Property
	if claim.Role != user.Black {
		tID, prop, err := teamIDFromProperty(ctx, p.client, uid, request.Key)
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
		chk, err = p.svc.GetByServiceIDKey(ctx, uid, request.GetKey())
		if err != nil {
			return nil, getErrorParser(err)
		}
	}
	return &propertyv1.GetByServiceIDKeyResponse{Property: ConvertPropertyToPropertyV1Pb(chk)}, nil
}

func (p PropertyV1Controller) GetAllByServiceID(ctx context.Context, request *propertyv1.GetAllByServiceIDRequest) (*propertyv1.GetAllByServiceIDResponse, error) {
	ID := request.GetServiceId()
	if ID == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}
	uid, err := uuid.FromString(ID.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}

	claim := extractUserClaim(ctx)

	if claim.Role != user.Black {
		tID, _, err := teamIDFromService(ctx, p.client, uid)
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

	props, err := p.svc.GetAllByServiceID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	propspb := make([]*propertyv1.Property, 0, len(props))
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyV1Pb(props[i]))
	}
	return &propertyv1.GetAllByServiceIDResponse{Properties: propspb}, nil
}

func NewPropertyV1Controller(svc propertyservice.Serv, client *util.Store) *PropertyV1Controller {
	return &PropertyV1Controller{svc: svc, client: client}
}

func ConvertPropertyV1PBtoProperty(pb *propertyv1.Property) (*property.Property, error) {
	var id uuid.UUID
	var err error
	if pb.GetServiceId() != nil {
		id, err = uuid.FromString(pb.GetServiceId().GetValue())
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
	if pb.GetKey() == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Key was not specified",
		)
	}
	var value *string
	if pb.GetValue() != nil {
		value = &pb.GetValue().Value
	}
	var st string

	switch pb.GetStatus() {
	case propertyv1.Status_STATUS_VIEW:
		st = property.View
	case propertyv1.Status_STATUS_EDIT:
		st = property.Edit
	case propertyv1.Status_STATUS_HIDE:
		st = property.Hide
	case propertyv1.Status_STATUS_UNSPECIFIED:
		st = ""
	}

	return &property.Property{
		ServiceID: id,
		Key:       pb.GetKey(),
		Value:     value,
		Status:    st,
	}, nil
}

func ConvertPropertyToPropertyV1Pb(obj *property.Property) *propertyv1.Property {
	var value *wrappers.StringValue
	if obj.Value != nil {
		value = &wrappers.StringValue{Value: *obj.Value}
	}
	var st propertyv1.Status
	switch obj.Status {
	case property.View:
		st = propertyv1.Status_STATUS_VIEW
	case property.Edit:
		st = propertyv1.Status_STATUS_EDIT
	case property.Hide:
		st = propertyv1.Status_STATUS_HIDE
	}

	return &propertyv1.Property{
		ServiceId: &protov1.UUID{Value: obj.ServiceID.String()},
		Key:       obj.Key,
		Value:     value,
		Status:    st,
	}
}
