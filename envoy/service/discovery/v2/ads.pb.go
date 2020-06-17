// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/ads.proto

package envoy_service_discovery_v2

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdsDummy) Reset()         { *m = AdsDummy{} }
func (m *AdsDummy) String() string { return proto.CompactTextString(m) }
func (*AdsDummy) ProtoMessage()    {}
func (*AdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_187fd5dcc2dab695, []int{0}
}

func (m *AdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdsDummy.Unmarshal(m, b)
}
func (m *AdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdsDummy.Marshal(b, m, deterministic)
}
func (m *AdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdsDummy.Merge(m, src)
}
func (m *AdsDummy) XXX_Size() int {
	return xxx_messageInfo_AdsDummy.Size(m)
}
func (m *AdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_AdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_AdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdsDummy)(nil), "envoy.service.discovery.v2.AdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/ads.proto", fileDescriptor_187fd5dcc2dab695)
}

var fileDescriptor_187fd5dcc2dab695 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0x31, 0x8b, 0x0a, 0x79, 0x85, 0xb2, 0x82, 0x88, 0x1f, 0xa9, 0x74, 0xd1, 0x95, 0x8d,
	0xc2, 0x09, 0x52, 0xe5, 0x00, 0x55, 0x7b, 0x82, 0x69, 0x32, 0x8a, 0x2c, 0x48, 0xc6, 0x78, 0x6c,
	0x8b, 0xec, 0x58, 0x72, 0x26, 0x4e, 0xc0, 0x96, 0x83, 0x70, 0x07, 0x94, 0xa4, 0xb4, 0x08, 0x28,
	0xeb, 0xf9, 0xde, 0x7b, 0xa3, 0x4f, 0xce, 0xb0, 0x8d, 0xd4, 0x69, 0x46, 0x17, 0x4d, 0x89, 0xba,
	0x32, 0x5c, 0x52, 0x44, 0xd7, 0xe9, 0x98, 0x69, 0xa8, 0x58, 0x59, 0x47, 0x9e, 0x92, 0x74, 0xa0,
	0xd4, 0x96, 0x52, 0x3b, 0x4a, 0xc5, 0x2c, 0xbd, 0x18, 0x1b, 0xc0, 0x9a, 0x3e, 0xb3, 0x3f, 0x0d,
	0xc9, 0xf4, 0x32, 0x54, 0x16, 0x34, 0xb4, 0x2d, 0x79, 0xf0, 0x86, 0x5a, 0xd6, 0xec, 0xc1, 0x87,
	0x6d, 0xf1, 0x54, 0xca, 0x93, 0xbc, 0xe2, 0x22, 0x34, 0x4d, 0x97, 0x7d, 0x08, 0x99, 0xe6, 0x75,
	0xed, 0xb0, 0x06, 0x8f, 0x55, 0xf1, 0x55, 0xb4, 0x1e, 0x47, 0x93, 0x8d, 0x3c, 0x5f, 0x7b, 0x87,
	0xd0, 0xec, 0x99, 0x15, 0x32, 0x05, 0x57, 0x22, 0x27, 0x57, 0x6a, 0xfc, 0x10, 0xac, 0x51, 0x31,
	0x53, 0xbb, 0xf0, 0x0a, 0x1f, 0x03, 0xb2, 0x4f, 0xaf, 0x0f, 0xde, 0xd9, 0x52, 0xcb, 0x38, 0x3d,
	0x9a, 0x8b, 0x5b, 0x91, 0xdc, 0xcb, 0xb3, 0x02, 0x1f, 0x3c, 0xfc, 0x35, 0x71, 0xf3, 0xa3, 0xa2,
	0xe7, 0x7e, 0xed, 0xcc, 0xfe, 0x87, 0xbe, 0x8f, 0x2d, 0xf2, 0xd7, 0xe7, 0xb7, 0xf7, 0xc9, 0xf1,
	0xa9, 0x90, 0x73, 0x43, 0x63, 0xce, 0x3a, 0x7a, 0xea, 0xd4, 0x61, 0xd9, 0x8b, 0xde, 0xd6, 0xb2,
	0x37, 0xb7, 0x14, 0x2f, 0x42, 0x6c, 0x26, 0x83, 0xc5, 0xbb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf0, 0xa8, 0xc5, 0xb7, 0xc6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AggregatedDiscoveryServiceClient is the client API for AggregatedDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AggregatedDiscoveryServiceClient interface {
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error)
	DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error)
}

type aggregatedDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewAggregatedDiscoveryServiceClient(cc *grpc.ClientConn) AggregatedDiscoveryServiceClient {
	return &aggregatedDiscoveryServiceClient{cc}
}

func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aggregatedDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v2.AggregatedDiscoveryService/DeltaAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceDeltaAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AggregatedDiscoveryServiceServer is the server API for AggregatedDiscoveryService service.
type AggregatedDiscoveryServiceServer interface {
	StreamAggregatedResources(AggregatedDiscoveryService_StreamAggregatedResourcesServer) error
	DeltaAggregatedResources(AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error
}

// UnimplementedAggregatedDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAggregatedDiscoveryServiceServer struct {
}

func (*UnimplementedAggregatedDiscoveryServiceServer) StreamAggregatedResources(srv AggregatedDiscoveryService_StreamAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAggregatedResources not implemented")
}
func (*UnimplementedAggregatedDiscoveryServiceServer) DeltaAggregatedResources(srv AggregatedDiscoveryService_DeltaAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaAggregatedResources not implemented")
}

func RegisterAggregatedDiscoveryServiceServer(s *grpc.Server, srv AggregatedDiscoveryServiceServer) {
	s.RegisterService(&_AggregatedDiscoveryService_serviceDesc, srv)
}

func _AggregatedDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).StreamAggregatedResources(&aggregatedDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AggregatedDiscoveryService_DeltaAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).DeltaAggregatedResources(&aggregatedDiscoveryServiceDeltaAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.AggregatedDiscoveryService",
	HandlerType: (*AggregatedDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _AggregatedDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaAggregatedResources",
			Handler:       _AggregatedDiscoveryService_DeltaAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/ads.proto",
}
