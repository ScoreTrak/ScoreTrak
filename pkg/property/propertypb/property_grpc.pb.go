// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package propertypb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PropertyServiceClient is the client API for PropertyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PropertyServiceClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetByServiceIDKey(ctx context.Context, in *GetByServiceIDKeyRequest, opts ...grpc.CallOption) (*GetByServiceIDKeyResponse, error)
	GetAllByServiceID(ctx context.Context, in *GetAllByServiceIDRequest, opts ...grpc.CallOption) (*GetAllByServiceIDResponse, error)
}

type propertyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPropertyServiceClient(cc grpc.ClientConnInterface) PropertyServiceClient {
	return &propertyServiceClient{cc}
}

func (c *propertyServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/pkg.property.propertypb.PropertyService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pkg.property.propertypb.PropertyService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) Store(ctx context.Context, in *StoreRequest, opts ...grpc.CallOption) (*StoreResponse, error) {
	out := new(StoreResponse)
	err := c.cc.Invoke(ctx, "/pkg.property.propertypb.PropertyService/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/pkg.property.propertypb.PropertyService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) GetByServiceIDKey(ctx context.Context, in *GetByServiceIDKeyRequest, opts ...grpc.CallOption) (*GetByServiceIDKeyResponse, error) {
	out := new(GetByServiceIDKeyResponse)
	err := c.cc.Invoke(ctx, "/pkg.property.propertypb.PropertyService/GetByServiceIDKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *propertyServiceClient) GetAllByServiceID(ctx context.Context, in *GetAllByServiceIDRequest, opts ...grpc.CallOption) (*GetAllByServiceIDResponse, error) {
	out := new(GetAllByServiceIDResponse)
	err := c.cc.Invoke(ctx, "/pkg.property.propertypb.PropertyService/GetAllByServiceID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PropertyServiceServer is the server API for PropertyService service.
// All implementations should embed UnimplementedPropertyServiceServer
// for forward compatibility
type PropertyServiceServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Store(context.Context, *StoreRequest) (*StoreResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	GetByServiceIDKey(context.Context, *GetByServiceIDKeyRequest) (*GetByServiceIDKeyResponse, error)
	GetAllByServiceID(context.Context, *GetAllByServiceIDRequest) (*GetAllByServiceIDResponse, error)
}

// UnimplementedPropertyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPropertyServiceServer struct {
}

func (*UnimplementedPropertyServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedPropertyServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedPropertyServiceServer) Store(context.Context, *StoreRequest) (*StoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedPropertyServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPropertyServiceServer) GetByServiceIDKey(context.Context, *GetByServiceIDKeyRequest) (*GetByServiceIDKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByServiceIDKey not implemented")
}
func (*UnimplementedPropertyServiceServer) GetAllByServiceID(context.Context, *GetAllByServiceIDRequest) (*GetAllByServiceIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByServiceID not implemented")
}

func RegisterPropertyServiceServer(s *grpc.Server, srv PropertyServiceServer) {
	s.RegisterService(&_PropertyService_serviceDesc, srv)
}

func _PropertyService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.property.propertypb.PropertyService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.property.propertypb.PropertyService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.property.propertypb.PropertyService/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).Store(ctx, req.(*StoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.property.propertypb.PropertyService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_GetByServiceIDKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByServiceIDKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetByServiceIDKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.property.propertypb.PropertyService/GetByServiceIDKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetByServiceIDKey(ctx, req.(*GetByServiceIDKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PropertyService_GetAllByServiceID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByServiceIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PropertyServiceServer).GetAllByServiceID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.property.propertypb.PropertyService/GetAllByServiceID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PropertyServiceServer).GetAllByServiceID(ctx, req.(*GetAllByServiceIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PropertyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.property.propertypb.PropertyService",
	HandlerType: (*PropertyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _PropertyService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PropertyService_Delete_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _PropertyService_Store_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PropertyService_Update_Handler,
		},
		{
			MethodName: "GetByServiceIDKey",
			Handler:    _PropertyService_GetByServiceIDKey_Handler,
		},
		{
			MethodName: "GetAllByServiceID",
			Handler:    _PropertyService_GetAllByServiceID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/property/propertypb/property.proto",
}
