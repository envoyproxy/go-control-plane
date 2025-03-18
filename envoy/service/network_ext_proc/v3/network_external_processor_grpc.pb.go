// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: envoy/service/network_ext_proc/v3/network_external_processor.proto

package network_ext_procv3

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
	NetworkExternalProcessor_Process_FullMethodName = "/envoy.service.network_ext_proc.v3.NetworkExternalProcessor/Process"
)

// NetworkExternalProcessorClient is the client API for NetworkExternalProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkExternalProcessorClient interface {
	// Process establishes a bidirectional stream between Envoy and the external
	// processing server. Envoy sends ProcessingRequests containing network data
	// and the server responds with ProcessingResponses containing processing
	// decisions and potentially modified data.
	//
	// The server should handle processing timeout properly to avoid blocking
	// network traffic for extended periods. Any uncaught exceptions will
	// be treated as errors and will terminate the stream.
	//
	// Implementation note: The server should process requests in the order
	// they are received to maintain proper sequencing of network traffic.
	Process(ctx context.Context, opts ...grpc.CallOption) (NetworkExternalProcessor_ProcessClient, error)
}

type networkExternalProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkExternalProcessorClient(cc grpc.ClientConnInterface) NetworkExternalProcessorClient {
	return &networkExternalProcessorClient{cc}
}

func (c *networkExternalProcessorClient) Process(ctx context.Context, opts ...grpc.CallOption) (NetworkExternalProcessor_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &NetworkExternalProcessor_ServiceDesc.Streams[0], NetworkExternalProcessor_Process_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &networkExternalProcessorProcessClient{stream}
	return x, nil
}

type NetworkExternalProcessor_ProcessClient interface {
	Send(*ProcessingRequest) error
	Recv() (*ProcessingResponse, error)
	grpc.ClientStream
}

type networkExternalProcessorProcessClient struct {
	grpc.ClientStream
}

func (x *networkExternalProcessorProcessClient) Send(m *ProcessingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *networkExternalProcessorProcessClient) Recv() (*ProcessingResponse, error) {
	m := new(ProcessingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkExternalProcessorServer is the server API for NetworkExternalProcessor service.
// All implementations should embed UnimplementedNetworkExternalProcessorServer
// for forward compatibility
type NetworkExternalProcessorServer interface {
	// Process establishes a bidirectional stream between Envoy and the external
	// processing server. Envoy sends ProcessingRequests containing network data
	// and the server responds with ProcessingResponses containing processing
	// decisions and potentially modified data.
	//
	// The server should handle processing timeout properly to avoid blocking
	// network traffic for extended periods. Any uncaught exceptions will
	// be treated as errors and will terminate the stream.
	//
	// Implementation note: The server should process requests in the order
	// they are received to maintain proper sequencing of network traffic.
	Process(NetworkExternalProcessor_ProcessServer) error
}

// UnimplementedNetworkExternalProcessorServer should be embedded to have forward compatible implementations.
type UnimplementedNetworkExternalProcessorServer struct {
}

func (UnimplementedNetworkExternalProcessorServer) Process(NetworkExternalProcessor_ProcessServer) error {
	return status.Errorf(codes.Unimplemented, "method Process not implemented")
}

// UnsafeNetworkExternalProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkExternalProcessorServer will
// result in compilation errors.
type UnsafeNetworkExternalProcessorServer interface {
	mustEmbedUnimplementedNetworkExternalProcessorServer()
}

func RegisterNetworkExternalProcessorServer(s grpc.ServiceRegistrar, srv NetworkExternalProcessorServer) {
	s.RegisterService(&NetworkExternalProcessor_ServiceDesc, srv)
}

func _NetworkExternalProcessor_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NetworkExternalProcessorServer).Process(&networkExternalProcessorProcessServer{stream})
}

type NetworkExternalProcessor_ProcessServer interface {
	Send(*ProcessingResponse) error
	Recv() (*ProcessingRequest, error)
	grpc.ServerStream
}

type networkExternalProcessorProcessServer struct {
	grpc.ServerStream
}

func (x *networkExternalProcessorProcessServer) Send(m *ProcessingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *networkExternalProcessorProcessServer) Recv() (*ProcessingRequest, error) {
	m := new(ProcessingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetworkExternalProcessor_ServiceDesc is the grpc.ServiceDesc for NetworkExternalProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkExternalProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.network_ext_proc.v3.NetworkExternalProcessor",
	HandlerType: (*NetworkExternalProcessorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _NetworkExternalProcessor_Process_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/network_ext_proc/v3/network_external_processor.proto",
}
