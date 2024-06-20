// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: envoy/service/status/v3/csds.proto

package statusv3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ClientStatusDiscoveryService_StreamClientStatus_FullMethodName = "/envoy.service.status.v3.ClientStatusDiscoveryService/StreamClientStatus"
	ClientStatusDiscoveryService_FetchClientStatus_FullMethodName  = "/envoy.service.status.v3.ClientStatusDiscoveryService/FetchClientStatus"
)

// ClientStatusDiscoveryServiceClient is the client API for ClientStatusDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientStatusDiscoveryServiceClient interface {
	StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (ClientStatusDiscoveryService_StreamClientStatusClient, error)
	FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error)
}

type clientStatusDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStatusDiscoveryServiceClient(cc grpc.ClientConnInterface) ClientStatusDiscoveryServiceClient {
	return &clientStatusDiscoveryServiceClient{cc}
}

func (c *clientStatusDiscoveryServiceClient) StreamClientStatus(ctx context.Context, opts ...grpc.CallOption) (ClientStatusDiscoveryService_StreamClientStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientStatusDiscoveryService_ServiceDesc.Streams[0], ClientStatusDiscoveryService_StreamClientStatus_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStatusDiscoveryServiceStreamClientStatusClient{stream}
	return x, nil
}

type ClientStatusDiscoveryService_StreamClientStatusClient interface {
	Send(*ClientStatusRequest) error
	Recv() (*ClientStatusResponse, error)
	grpc.ClientStream
}

type clientStatusDiscoveryServiceStreamClientStatusClient struct {
	grpc.ClientStream
}

func (x *clientStatusDiscoveryServiceStreamClientStatusClient) Send(m *ClientStatusRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStatusDiscoveryServiceStreamClientStatusClient) Recv() (*ClientStatusResponse, error) {
	m := new(ClientStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientStatusDiscoveryServiceClient) FetchClientStatus(ctx context.Context, in *ClientStatusRequest, opts ...grpc.CallOption) (*ClientStatusResponse, error) {
	out := new(ClientStatusResponse)
	err := c.cc.Invoke(ctx, ClientStatusDiscoveryService_FetchClientStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientStatusDiscoveryServiceServer is the server API for ClientStatusDiscoveryService service.
// All implementations should embed UnimplementedClientStatusDiscoveryServiceServer
// for forward compatibility
type ClientStatusDiscoveryServiceServer interface {
	StreamClientStatus(ClientStatusDiscoveryService_StreamClientStatusServer) error
	FetchClientStatus(context.Context, *ClientStatusRequest) (*ClientStatusResponse, error)
}

// UnimplementedClientStatusDiscoveryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClientStatusDiscoveryServiceServer struct {
}

func (UnimplementedClientStatusDiscoveryServiceServer) StreamClientStatus(ClientStatusDiscoveryService_StreamClientStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClientStatus not implemented")
}
func (UnimplementedClientStatusDiscoveryServiceServer) FetchClientStatus(context.Context, *ClientStatusRequest) (*ClientStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClientStatus not implemented")
}

// UnsafeClientStatusDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStatusDiscoveryServiceServer will
// result in compilation errors.
type UnsafeClientStatusDiscoveryServiceServer interface {
	mustEmbedUnimplementedClientStatusDiscoveryServiceServer()
}

func RegisterClientStatusDiscoveryServiceServer(s grpc.ServiceRegistrar, srv ClientStatusDiscoveryServiceServer) {
	s.RegisterService(&ClientStatusDiscoveryService_ServiceDesc, srv)
}

func _ClientStatusDiscoveryService_StreamClientStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStatusDiscoveryServiceServer).StreamClientStatus(&clientStatusDiscoveryServiceStreamClientStatusServer{stream})
}

type ClientStatusDiscoveryService_StreamClientStatusServer interface {
	Send(*ClientStatusResponse) error
	Recv() (*ClientStatusRequest, error)
	grpc.ServerStream
}

type clientStatusDiscoveryServiceStreamClientStatusServer struct {
	grpc.ServerStream
}

func (x *clientStatusDiscoveryServiceStreamClientStatusServer) Send(m *ClientStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStatusDiscoveryServiceStreamClientStatusServer) Recv() (*ClientStatusRequest, error) {
	m := new(ClientStatusRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientStatusDiscoveryService_FetchClientStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClientStatusDiscoveryService_FetchClientStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientStatusDiscoveryServiceServer).FetchClientStatus(ctx, req.(*ClientStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientStatusDiscoveryService_ServiceDesc is the grpc.ServiceDesc for ClientStatusDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStatusDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.status.v3.ClientStatusDiscoveryService",
	HandlerType: (*ClientStatusDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClientStatus",
			Handler:    _ClientStatusDiscoveryService_FetchClientStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClientStatus",
			Handler:       _ClientStatusDiscoveryService_StreamClientStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/status/v3/csds.proto",
}
