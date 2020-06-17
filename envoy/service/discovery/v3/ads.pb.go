// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v3/ads.proto

package envoy_service_discovery_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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
	return fileDescriptor_1cbfe82b86373f2f, []int{0}
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
	proto.RegisterType((*AdsDummy)(nil), "envoy.service.discovery.v3.AdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v3/ads.proto", fileDescriptor_1cbfe82b86373f2f)
}

var fileDescriptor_1cbfe82b86373f2f = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4a, 0xf3, 0x40,
	0x18, 0xc6, 0xbf, 0xe9, 0xa2, 0x7c, 0xcc, 0x4a, 0xb2, 0xd2, 0x01, 0x05, 0xab, 0x42, 0x29, 0x3a,
	0xa3, 0x0d, 0xb8, 0x70, 0x97, 0x92, 0x03, 0x94, 0xf6, 0x04, 0x63, 0xf2, 0x12, 0x06, 0xcc, 0x4c,
	0x9c, 0x77, 0x32, 0x98, 0x8d, 0xb8, 0xb3, 0x67, 0xf0, 0x00, 0x1e, 0xc2, 0xbd, 0xe0, 0xd6, 0x1b,
	0x49, 0x92, 0xfe, 0x11, 0x6a, 0x83, 0xae, 0xe7, 0xf7, 0xfc, 0xde, 0x67, 0x78, 0xe8, 0x29, 0x68,
	0x6f, 0x2a, 0x81, 0x60, 0xbd, 0x4a, 0x40, 0xa4, 0x0a, 0x13, 0xe3, 0xc1, 0x56, 0xc2, 0x87, 0x42,
	0xa6, 0xc8, 0x0b, 0x6b, 0x9c, 0x09, 0x58, 0x43, 0xf1, 0x25, 0xc5, 0xd7, 0x14, 0xf7, 0x21, 0x1b,
	0x75, 0x18, 0x36, 0x60, 0xe3, 0x61, 0x87, 0x65, 0x5a, 0x48, 0x21, 0xb5, 0x36, 0x4e, 0x3a, 0x65,
	0x34, 0x0a, 0x74, 0xd2, 0x95, 0xcb, 0x33, 0xec, 0x78, 0xeb, 0xd9, 0x83, 0x45, 0x65, 0xb4, 0xd2,
	0x59, 0x8b, 0x0c, 0xae, 0xe9, 0xff, 0x28, 0xc5, 0xb8, 0xcc, 0xf3, 0xea, 0x66, 0xf4, 0xf2, 0xbe,
	0x38, 0x3a, 0xa3, 0x27, 0x3b, 0xcb, 0x8d, 0xf9, 0x8a, 0x1d, 0xbf, 0xf6, 0x28, 0x8b, 0xb2, 0xcc,
	0x42, 0x26, 0x1d, 0xa4, 0xf1, 0x8a, 0x99, 0xb7, 0xa1, 0xe0, 0x91, 0x1e, 0xcc, 0x9d, 0x05, 0x99,
	0x6f, 0x98, 0x19, 0xa0, 0x29, 0x6d, 0x02, 0x18, 0x9c, 0xf3, 0xdd, 0xdf, 0xe7, 0x6b, 0xd5, 0x0c,
	0xee, 0x4b, 0x40, 0xc7, 0x2e, 0x7e, 0x49, 0x63, 0x61, 0x34, 0xc2, 0xe0, 0xdf, 0x90, 0x5c, 0x92,
	0xe0, 0x99, 0xd0, 0xfd, 0x18, 0xee, 0x9c, 0xfc, 0xe9, 0xfe, 0x55, 0xa7, 0xb1, 0x4e, 0x6d, 0x95,
	0x18, 0xff, 0x25, 0xf2, 0xbd, 0xc9, 0x24, 0x7a, 0x7b, 0xfa, 0xf8, 0xec, 0xf7, 0xf6, 0x7a, 0x74,
	0xa8, 0x4c, 0x6b, 0x29, 0xac, 0x79, 0xa8, 0x3a, 0x84, 0x93, 0x7a, 0x92, 0x69, 0x3d, 0xcf, 0x94,
	0x2c, 0x08, 0xb9, 0xed, 0x37, 0x53, 0x85, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0x5e, 0x16,
	0x41, 0x5c, 0x02, 0x00, 0x00,
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
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v3.AggregatedDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *aggregatedDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AggregatedDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v3.AggregatedDiscoveryService/DeltaAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &aggregatedDiscoveryServiceDeltaAggregatedResourcesClient{stream}
	return x, nil
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
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
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AggregatedDiscoveryService_DeltaAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AggregatedDiscoveryServiceServer).DeltaAggregatedResources(&aggregatedDiscoveryServiceDeltaAggregatedResourcesServer{stream})
}

type AggregatedDiscoveryService_DeltaAggregatedResourcesServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AggregatedDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v3.AggregatedDiscoveryService",
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
	Metadata: "envoy/service/discovery/v3/ads.proto",
}
