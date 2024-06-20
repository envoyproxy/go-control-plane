// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: envoy/service/trace/v3/trace_service.proto

package tracev3

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
	TraceService_StreamTraces_FullMethodName = "/envoy.service.trace.v3.TraceService/StreamTraces"
)

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceServiceClient interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error)
}

type traceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceServiceClient(cc grpc.ClientConnInterface) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &TraceService_ServiceDesc.Streams[0], TraceService_StreamTraces_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceStreamTracesClient{stream}
	return x, nil
}

type TraceService_StreamTracesClient interface {
	Send(*StreamTracesMessage) error
	CloseAndRecv() (*StreamTracesResponse, error)
	grpc.ClientStream
}

type traceServiceStreamTracesClient struct {
	grpc.ClientStream
}

func (x *traceServiceStreamTracesClient) Send(m *StreamTracesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceStreamTracesClient) CloseAndRecv() (*StreamTracesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
// All implementations should embed UnimplementedTraceServiceServer
// for forward compatibility
type TraceServiceServer interface {
	// Envoy will connect and send StreamTracesMessage messages forever. It does
	// not expect any response to be sent as nothing would be done in the case
	// of failure.
	StreamTraces(TraceService_StreamTracesServer) error
}

// UnimplementedTraceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (UnimplementedTraceServiceServer) StreamTraces(TraceService_StreamTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTraces not implemented")
}

// UnsafeTraceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraceServiceServer will
// result in compilation errors.
type UnsafeTraceServiceServer interface {
	mustEmbedUnimplementedTraceServiceServer()
}

func RegisterTraceServiceServer(s grpc.ServiceRegistrar, srv TraceServiceServer) {
	s.RegisterService(&TraceService_ServiceDesc, srv)
}

func _TraceService_StreamTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).StreamTraces(&traceServiceStreamTracesServer{stream})
}

type TraceService_StreamTracesServer interface {
	SendAndClose(*StreamTracesResponse) error
	Recv() (*StreamTracesMessage, error)
	grpc.ServerStream
}

type traceServiceStreamTracesServer struct {
	grpc.ServerStream
}

func (x *traceServiceStreamTracesServer) SendAndClose(m *StreamTracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceStreamTracesServer) Recv() (*StreamTracesMessage, error) {
	m := new(StreamTracesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceService_ServiceDesc is the grpc.ServiceDesc for TraceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.trace.v3.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTraces",
			Handler:       _TraceService_StreamTraces_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/trace/v3/trace_service.proto",
}
