package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertypb"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PropertyController struct {
	svc service.Serv
}

func (p PropertyController) GetAll(ctx context.Context, request *propertypb.GetAllRequest) (*propertypb.GetAllResponse, error) {
	props, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)

	}
	var propspb []*propertypb.Property
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
			"ID was not specified",
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
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
	var props []*property.Property
	for i := range propspb {
		prop, err := ConvertPropertyPBtoProperty(propspb[i])
		if err != nil {
			return nil, err
		}
		props = append(props, prop)
	}

	err := p.svc.Store(ctx, props)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &propertypb.StoreResponse{}, nil
}

func (p PropertyController) Update(ctx context.Context, request *propertypb.UpdateRequest) (*propertypb.UpdateResponse, error) {
	prop, err := ConvertPropertyPBtoProperty(request.Property)
	if err != nil {
		return nil, err
	}
	err = p.svc.Update(ctx, prop)
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
			"ID was not specified",
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	prop, err := p.svc.GetByServiceIDKey(ctx, uid, request.GetKey())
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &propertypb.GetByServiceIDKeyResponse{Property: ConvertPropertyToPropertyPb(prop)}, nil
}

func (p PropertyController) GetAllByServiceID(ctx context.Context, request *propertypb.GetAllByServiceIDRequest) (*propertypb.GetAllByServiceIDResponse, error) {
	id := request.GetServiceId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ID was not specified",
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	props, err := p.svc.GetAllByServiceID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var propspb []*propertypb.Property
	for i := range props {
		propspb = append(propspb, ConvertPropertyToPropertyPb(props[i]))
	}
	return &propertypb.GetAllByServiceIDResponse{Properties: propspb}, nil
}

func NewPropertyController(svc service.Serv) *PropertyController {
	return &PropertyController{svc}
}

func ConvertPropertyPBtoProperty(pb *propertypb.Property) (*property.Property, error) {
	var id uuid.UUID
	var err error
	if pb.GetServiceId() != nil {
		id, err = uuid.FromString(pb.GetServiceId().GetValue())
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
	if pb.GetStatus() == propertypb.Status_View {
		st = property.View
	} else if pb.GetStatus() == propertypb.Status_Edit {
		st = property.Edit
	} else if pb.GetStatus() == propertypb.Status_Hide {
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
		st = propertypb.Status_View
	} else if obj.Status == property.Edit {
		st = propertypb.Status_Edit
	} else if obj.Status == property.Hide {
		st = propertypb.Status_Hide
	}
	return &propertypb.Property{
		ServiceId: &utilpb.UUID{Value: obj.ServiceID.String()},
		Key:       obj.Key,
		Value:     value,
		Status:    st,
	}
}