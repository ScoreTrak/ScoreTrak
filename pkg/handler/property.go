package handler

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	propertyv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/property/v1"
	protov1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PropertyController struct {
	svc    propertyservice.Serv
	client *util.Store
	propertyv1.UnimplementedPropertyServiceServer
}

func (p PropertyController) GetAll(ctx context.Context, _ *propertyv1.GetAllRequest) (*propertyv1.GetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	propspb := make([]*propertyv1.Property, 0, len(props))
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyPb(props[i]))
	}
	return &propertyv1.GetAllResponse{Properties: propspb}, nil
}

func (p PropertyController) Delete(ctx context.Context, request *propertyv1.DeleteRequest) (*propertyv1.DeleteResponse, error) {
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

func (p PropertyController) Store(ctx context.Context, request *propertyv1.StoreRequest) (*propertyv1.StoreResponse, error) {
	propspb := request.GetProperties()
	props := make([]*property.Property, 0, len(propspb))
	for i := range propspb {
		prop, err := ConvertPropertyPBtoProperty(propspb[i])
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

func (p PropertyController) Update(ctx context.Context, request *propertyv1.UpdateRequest) (*propertyv1.UpdateResponse, error) {
	pr, err := ConvertPropertyPBtoProperty(request.Property)
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

func (p PropertyController) GetByServiceIDKey(ctx context.Context, request *propertyv1.GetByServiceIDKeyRequest) (*propertyv1.GetByServiceIDKeyResponse, error) {
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
	return &propertyv1.GetByServiceIDKeyResponse{Property: ConvertPropertyToPropertyPb(chk)}, nil
}

func (p PropertyController) GetAllByServiceID(ctx context.Context, request *propertyv1.GetAllByServiceIDRequest) (*propertyv1.GetAllByServiceIDResponse, error) {
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
		propspb = append(propspb, ConvertPropertyToPropertyPb(props[i]))
	}
	return &propertyv1.GetAllByServiceIDResponse{Properties: propspb}, nil
}

func NewPropertyController(svc propertyservice.Serv, client *util.Store) *PropertyController {
	return &PropertyController{svc: svc, client: client}
}

func ConvertPropertyPBtoProperty(pb *propertyv1.Property) (*property.Property, error) {
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

func ConvertPropertyToPropertyPb(obj *property.Property) *propertyv1.Property {
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
