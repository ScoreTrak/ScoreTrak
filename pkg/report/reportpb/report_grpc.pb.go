// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package reportpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportServiceClient is the client API for ReportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (ReportService_GetClient, error)
}

type reportServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceClient(cc grpc.ClientConnInterface) ReportServiceClient {
	return &reportServiceClient{cc}
}

func (c *reportServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (ReportService_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ReportService_serviceDesc.Streams[0], "/pkg.report.reportpb.ReportService/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &reportServiceGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReportService_GetClient interface {
	Recv() (*GetResponse, error)
	grpc.ClientStream
}

type reportServiceGetClient struct {
	grpc.ClientStream
}

func (x *reportServiceGetClient) Recv() (*GetResponse, error) {
	m := new(GetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReportServiceServer is the server API for ReportService service.
// All implementations should embed UnimplementedReportServiceServer
// for forward compatibility
type ReportServiceServer interface {
	Get(*GetRequest, ReportService_GetServer) error
}

// UnimplementedReportServiceServer should be embedded to have forward compatible implementations.
type UnimplementedReportServiceServer struct {
}

func (*UnimplementedReportServiceServer) Get(*GetRequest, ReportService_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterReportServiceServer(s *grpc.Server, srv ReportServiceServer) {
	s.RegisterService(&_ReportService_serviceDesc, srv)
}

func _ReportService_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReportServiceServer).Get(m, &reportServiceGetServer{stream})
}

type ReportService_GetServer interface {
	Send(*GetResponse) error
	grpc.ServerStream
}

type reportServiceGetServer struct {
	grpc.ServerStream
}

func (x *reportServiceGetServer) Send(m *GetResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ReportService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.report.reportpb.ReportService",
	HandlerType: (*ReportServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _ReportService_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/report/reportpb/report.proto",
}
