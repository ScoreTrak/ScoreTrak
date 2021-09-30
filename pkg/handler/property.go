package handler

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/property_service"
	propertypb "github.com/ScoreTrak/ScoreTrak/pkg/proto/property/v1"
	utilpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/proto/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PropertyController struct {
	svc    property_service.Serv
	client *util.Store
	propertypb.UnimplementedPropertyServiceServer
}

func (p PropertyController) GetAll(ctx context.Context, request *propertypb.GetAllRequest) (*propertypb.GetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	propspb := make([]*propertypb.Property, 0, len(props))
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyPb(props[i]))
	}
	return &propertypb.GetAllResponse{Properties: propspb}, nil
}

func (p PropertyController) Delete(ctx context.Context, request *propertypb.DeleteRequest) (*propertypb.DeleteResponse, error) {
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
	return &propertypb.DeleteResponse{}, nil
}

func (p PropertyController) Store(ctx context.Context, request *propertypb.StoreRequest) (*propertypb.StoreResponse, error) {
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
	return &propertypb.StoreResponse{}, nil
}

func (p PropertyController) Update(ctx context.Context, request *propertypb.UpdateRequest) (*propertypb.UpdateResponse, error) {
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
	return &propertypb.UpdateResponse{}, nil
}

func (p PropertyController) GetByServiceIDKey(ctx context.Context, request *propertypb.GetByServiceIDKeyRequest) (*propertypb.GetByServiceIDKeyResponse, error) {
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
	return &propertypb.GetByServiceIDKeyResponse{Property: ConvertPropertyToPropertyPb(chk)}, nil
}

func (p PropertyController) GetAllByServiceID(ctx context.Context, request *propertypb.GetAllByServiceIDRequest) (*propertypb.GetAllByServiceIDResponse, error) {
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
	propspb := make([]*propertypb.Property, 0, len(props))
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyPb(props[i]))
	}
	return &propertypb.GetAllByServiceIDResponse{Properties: propspb}, nil
}

func NewPropertyController(svc property_service.Serv, client *util.Store) *PropertyController {
	return &PropertyController{svc: svc, client: client}
}

func ConvertPropertyPBtoProperty(pb *propertypb.Property) (*property.Property, error) {
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

	if pb.GetStatus() == propertypb.Status_STATUS_VIEW {
		st = property.View
	} else if pb.GetStatus() == propertypb.Status_STATUS_EDIT {
		st = property.Edit
	} else if pb.GetStatus() == propertypb.Status_STATUS_HIDE {
		st = property.Hide
	}

	return &property.Property{
		ServiceID: id,
		Key:       pb.GetKey(),
		Value:     value,
		Status:    st,
	}, nil
}

func ConvertPropertyToPropertyPb(obj *property.Property) *propertypb.Property {
	var value *wrappers.StringValue
	if obj.Value != nil {
		value = &wrappers.StringValue{Value: *obj.Value}
	}
	var st propertypb.Status
	if obj.Status == property.View {
		st = propertypb.Status_STATUS_VIEW
	} else if obj.Status == property.Edit {
		st = propertypb.Status_STATUS_EDIT
	} else if obj.Status == property.Hide {
		st = propertypb.Status_STATUS_HIDE
	}
	return &propertypb.Property{
		ServiceId: &utilpb.UUID{Value: obj.ServiceID.String()},
		Key:       obj.Key,
		Value:     value,
		Status:    st,
	}
}
