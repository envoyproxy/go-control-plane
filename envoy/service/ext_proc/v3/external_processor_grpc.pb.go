// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: envoy/service/ext_proc/v3/external_processor.proto

package ext_procv3

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
	ExternalProcessor_Process_FullMethodName = "/envoy.service.ext_proc.v3.ExternalProcessor/Process"
)

// ExternalProcessorClient is the client API for ExternalProcessor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalProcessorClient interface {
	// This begins the bidirectional stream that Envoy will use to
	// give the server control over what the filter does. The actual
	// protocol is described by the ProcessingRequest and ProcessingResponse
	// messages below.
	Process(ctx context.Context, opts ...grpc.CallOption) (ExternalProcessor_ProcessClient, error)
}

type externalProcessorClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalProcessorClient(cc grpc.ClientConnInterface) ExternalProcessorClient {
	return &externalProcessorClient{cc}
}

func (c *externalProcessorClient) Process(ctx context.Context, opts ...grpc.CallOption) (ExternalProcessor_ProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExternalProcessor_ServiceDesc.Streams[0], ExternalProcessor_Process_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &externalProcessorProcessClient{stream}
	return x, nil
}

type ExternalProcessor_ProcessClient interface {
	Send(*ProcessingRequest) error
	Recv() (*ProcessingResponse, error)
	grpc.ClientStream
}

type externalProcessorProcessClient struct {
	grpc.ClientStream
}

func (x *externalProcessorProcessClient) Send(m *ProcessingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *externalProcessorProcessClient) Recv() (*ProcessingResponse, error) {
	m := new(ProcessingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExternalProcessorServer is the server API for ExternalProcessor service.
// All implementations should embed UnimplementedExternalProcessorServer
// for forward compatibility
type ExternalProcessorServer interface {
	// This begins the bidirectional stream that Envoy will use to
	// give the server control over what the filter does. The actual
	// protocol is described by the ProcessingRequest and ProcessingResponse
	// messages below.
	Process(ExternalProcessor_ProcessServer) error
}

// UnimplementedExternalProcessorServer should be embedded to have forward compatible implementations.
type UnimplementedExternalProcessorServer struct {
}

func (UnimplementedExternalProcessorServer) Process(ExternalProcessor_ProcessServer) error {
	return status.Errorf(codes.Unimplemented, "method Process not implemented")
}

// UnsafeExternalProcessorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalProcessorServer will
// result in compilation errors.
type UnsafeExternalProcessorServer interface {
	mustEmbedUnimplementedExternalProcessorServer()
}

func RegisterExternalProcessorServer(s grpc.ServiceRegistrar, srv ExternalProcessorServer) {
	s.RegisterService(&ExternalProcessor_ServiceDesc, srv)
}

func _ExternalProcessor_Process_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ExternalProcessorServer).Process(&externalProcessorProcessServer{stream})
}

type ExternalProcessor_ProcessServer interface {
	Send(*ProcessingResponse) error
	Recv() (*ProcessingRequest, error)
	grpc.ServerStream
}

type externalProcessorProcessServer struct {
	grpc.ServerStream
}

func (x *externalProcessorProcessServer) Send(m *ProcessingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *externalProcessorProcessServer) Recv() (*ProcessingRequest, error) {
	m := new(ProcessingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExternalProcessor_ServiceDesc is the grpc.ServiceDesc for ExternalProcessor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalProcessor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ext_proc.v3.ExternalProcessor",
	HandlerType: (*ExternalProcessorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Process",
			Handler:       _ExternalProcessor_Process_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/ext_proc/v3/external_processor.proto",
}
