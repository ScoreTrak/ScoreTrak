package handler

import (
	"context"
	"fmt"
	propertyv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/property/v2"

	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PropertyV2Controller struct {
	svc    propertyservice.Serv
	client *util.Store
	propertyv2.UnimplementedPropertyServiceServer
}

func (p PropertyV2Controller) GetAll(ctx context.Context, _ *propertyv2.PropertyServiceGetAllRequest) (*propertyv2.PropertyServiceGetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	propspb := make([]*propertyv2.Property, 0, len(props))
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyV2Pb(props[i]))
	}
	return &propertyv2.PropertyServiceGetAllResponse{Properties: propspb}, nil
}

func (p PropertyV2Controller) Delete(ctx context.Context, request *propertyv2.PropertyServiceDeleteRequest) (*propertyv2.PropertyServiceDeleteResponse, error) {
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
	return &propertyv2.PropertyServiceDeleteResponse{}, nil
}

func (p PropertyV2Controller) Store(ctx context.Context, request *propertyv2.PropertyServiceStoreRequest) (*propertyv2.PropertyServiceStoreResponse, error) {
	propspb := request.GetProperties()
	props := make([]*property.Property, 0, len(propspb))
	for i := range propspb {
		prop, err := ConvertPropertyV2PBtoProperty(propspb[i])
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
	return &propertyv2.PropertyServiceStoreResponse{}, nil
}

func (p PropertyV2Controller) Update(ctx context.Context, request *propertyv2.PropertyServiceUpdateRequest) (*propertyv2.PropertyServiceUpdateResponse, error) {
	pr, err := ConvertPropertyV2PBtoProperty(request.Property)
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
	return &propertyv2.PropertyServiceUpdateResponse{}, nil
}

func (p PropertyV2Controller) GetByServiceIDKey(ctx context.Context, request *propertyv2.PropertyServiceGetByServiceIDKeyRequest) (*propertyv2.PropertyServiceGetByServiceIDKeyResponse, error) {
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
	return &propertyv2.PropertyServiceGetByServiceIDKeyResponse{Property: ConvertPropertyToPropertyV2Pb(chk)}, nil
}

func (p PropertyV2Controller) GetAllByServiceID(ctx context.Context, request *propertyv2.PropertyServiceGetAllByServiceIDRequest) (*propertyv2.PropertyServiceGetAllByServiceIDResponse, error) {
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
	propspb := make([]*propertyv2.Property, 0, len(props))
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyV2Pb(props[i]))
	}
	return &propertyv2.PropertyServiceGetAllByServiceIDResponse{Properties: propspb}, nil
}

func NewPropertyV2Controller(svc propertyservice.Serv, client *util.Store) *PropertyV2Controller {
	return &PropertyV2Controller{svc: svc, client: client}
}

func ConvertPropertyV2PBtoProperty(pb *propertyv2.Property) (*property.Property, error) {
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
	case propertyv2.Status_STATUS_VIEW:
		st = property.View
	case propertyv2.Status_STATUS_EDIT:
		st = property.Edit
	case propertyv2.Status_STATUS_HIDE:
		st = property.Hide
	case propertyv2.Status_STATUS_UNSPECIFIED:
		st = ""
	}

	return &property.Property{
		ServiceID: id,
		Key:       pb.GetKey(),
		Value:     value,
		Status:    st,
	}, nil
}

func ConvertPropertyToPropertyV2Pb(obj *property.Property) *propertyv2.Property {
	var value *wrappers.StringValue
	if obj.Value != nil {
		value = &wrappers.StringValue{Value: *obj.Value}
	}
	var st propertyv2.Status
	switch obj.Status {
	case property.View:
		st = propertyv2.Status_STATUS_VIEW
	case property.Edit:
		st = propertyv2.Status_STATUS_EDIT
	case property.Hide:
		st = propertyv2.Status_STATUS_HIDE
	}

	return &propertyv2.Property{
		ServiceId: &protov1.UUID{Value: obj.ServiceID.String()},
		Key:       obj.Key,
		Value:     value,
		Status:    st,
	}
}
