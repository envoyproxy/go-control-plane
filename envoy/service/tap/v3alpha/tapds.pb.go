// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/tap/v3alpha/tapds.proto

package envoy_service_tap_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/tap/v3alpha"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TapResource struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config               *v3alpha.TapConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TapResource) Reset()         { *m = TapResource{} }
func (m *TapResource) String() string { return proto.CompactTextString(m) }
func (*TapResource) ProtoMessage()    {}
func (*TapResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba445dccde093552, []int{0}
}

func (m *TapResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TapResource.Unmarshal(m, b)
}
func (m *TapResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TapResource.Marshal(b, m, deterministic)
}
func (m *TapResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TapResource.Merge(m, src)
}
func (m *TapResource) XXX_Size() int {
	return xxx_messageInfo_TapResource.Size(m)
}
func (m *TapResource) XXX_DiscardUnknown() {
	xxx_messageInfo_TapResource.DiscardUnknown(m)
}

var xxx_messageInfo_TapResource proto.InternalMessageInfo

func (m *TapResource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TapResource) GetConfig() *v3alpha.TapConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*TapResource)(nil), "envoy.service.tap.v3alpha.TapResource")
}

func init() {
	proto.RegisterFile("envoy/service/tap/v3alpha/tapds.proto", fileDescriptor_ba445dccde093552)
}

var fileDescriptor_ba445dccde093552 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0x41, 0x6b, 0x14, 0x31,
	0x18, 0x35, 0xdb, 0x52, 0x31, 0x3d, 0x54, 0xc6, 0x83, 0x75, 0x95, 0xb2, 0xae, 0x48, 0x17, 0x91,
	0x44, 0xb7, 0xa0, 0xb0, 0xde, 0xd6, 0xe2, 0xb9, 0x6c, 0xf7, 0x2e, 0x9f, 0x33, 0x9f, 0xd3, 0xc0,
	0x4c, 0xbe, 0x98, 0x64, 0x07, 0xf7, 0x24, 0x78, 0xda, 0xbb, 0x78, 0xf1, 0x1f, 0xf8, 0x43, 0x3c,
	0x7a, 0xf2, 0x2f, 0xf8, 0x2b, 0x3c, 0xc9, 0x4c, 0x32, 0xe3, 0x0e, 0x45, 0xa4, 0x17, 0x6f, 0xe1,
	0xcd, 0x7b, 0xef, 0x7b, 0xf3, 0xf2, 0x85, 0x3f, 0x44, 0x5d, 0xd1, 0x5a, 0x3a, 0xb4, 0x95, 0x4a,
	0x51, 0x7a, 0x30, 0xb2, 0x3a, 0x81, 0xc2, 0x5c, 0x40, 0x7d, 0xce, 0x9c, 0x30, 0x96, 0x3c, 0x25,
	0x77, 0x1a, 0x9a, 0x88, 0x34, 0xe1, 0xc1, 0x88, 0x48, 0x1b, 0x46, 0x87, 0x94, 0xf4, 0x5b, 0x95,
	0xf7, 0x0c, 0x52, 0x2a, 0x4b, 0xd2, 0xc1, 0x61, 0x28, 0xfb, 0x83, 0x32, 0xe5, 0x52, 0xaa, 0xd0,
	0xae, 0x3b, 0x76, 0x87, 0x44, 0xc1, 0xbd, 0x9c, 0x28, 0x2f, 0x50, 0x82, 0x51, 0x12, 0xb4, 0x26,
	0x0f, 0x5e, 0x91, 0x8e, 0x81, 0x86, 0xf7, 0x57, 0x99, 0x81, 0x6d, 0x5c, 0x56, 0x68, 0x9d, 0x22,
	0xad, 0x74, 0x1e, 0x29, 0xb7, 0x2b, 0x28, 0x54, 0x06, 0x1e, 0x65, 0x7b, 0x08, 0x1f, 0xc6, 0x9f,
	0x19, 0xdf, 0x5f, 0x82, 0x59, 0xa0, 0xa3, 0x95, 0x4d, 0x31, 0xb9, 0xcb, 0x77, 0x35, 0x94, 0x78,
	0xc8, 0x46, 0x6c, 0x72, 0x63, 0x7e, 0xfd, 0xd7, 0x7c, 0xd7, 0x0e, 0x46, 0x6c, 0xd1, 0x80, 0xc9,
	0x0b, 0xbe, 0x17, 0x7e, 0xed, 0x70, 0x30, 0x62, 0x93, 0xfd, 0xe9, 0x03, 0x11, 0xaa, 0x08, 0xe0,
	0x76, 0x13, 0x62, 0x09, 0xe6, 0x65, 0x83, 0x2e, 0xa2, 0x64, 0xf6, 0xf8, 0xcb, 0xb7, 0xcd, 0xd1,
	0x71, 0x2c, 0xb9, 0xdf, 0xde, 0xb4, 0xd3, 0xb4, 0x39, 0xa6, 0xdf, 0x77, 0xf8, 0xad, 0x25, 0x98,
	0xd3, 0xb6, 0x88, 0xf3, 0xc0, 0x4f, 0x3e, 0xf0, 0x9b, 0xe7, 0xde, 0x22, 0x94, 0xdd, 0x00, 0x97,
	0x3c, 0x15, 0x7d, 0xcf, 0x3f, 0xed, 0xb5, 0x69, 0x3a, 0x9b, 0x05, 0xbe, 0x5b, 0xa1, 0xf3, 0xc3,
	0xe9, 0x55, 0x24, 0xce, 0x90, 0x76, 0x38, 0xbe, 0x36, 0x61, 0x4f, 0x58, 0xb2, 0x61, 0xfc, 0xe0,
	0x14, 0x0b, 0x0f, 0x5b, 0x01, 0x9e, 0xfd, 0xdb, 0xad, 0x56, 0x5c, 0x4a, 0xf1, 0xfc, 0xca, 0xba,
	0x5e, 0x94, 0xaf, 0x8c, 0x1f, 0xbc, 0x42, 0x9f, 0x5e, 0xfc, 0xff, 0x2e, 0xe4, 0xc7, 0x1f, 0x3f,
	0x3f, 0x0d, 0x46, 0xe3, 0xa3, 0xcb, 0x2b, 0x3b, 0xf3, 0x60, 0x5e, 0x87, 0x3b, 0x77, 0x0d, 0x6b,
	0x67, 0xc6, 0x1e, 0xcd, 0x67, 0xfc, 0x58, 0x51, 0x18, 0x64, 0x2c, 0xbd, 0x5f, 0x8b, 0xbf, 0x3e,
	0xa2, 0x39, 0x5f, 0xd6, 0x8f, 0xed, 0xac, 0x5e, 0xcf, 0x33, 0xb6, 0x61, 0xec, 0xcd, 0x5e, 0xb3,
	0xaa, 0x27, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x6d, 0xca, 0x4b, 0xa0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TapDiscoveryServiceClient is the client API for TapDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TapDiscoveryServiceClient interface {
	StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error)
	DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error)
	FetchTapConfigs(ctx context.Context, in *v3alpha1.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha1.DiscoveryResponse, error)
}

type tapDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewTapDiscoveryServiceClient(cc *grpc.ClientConn) TapDiscoveryServiceClient {
	return &tapDiscoveryServiceClient{cc}
}

func (c *tapDiscoveryServiceClient) StreamTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_StreamTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[0], "/envoy.service.tap.v3alpha.TapDiscoveryService/StreamTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceStreamTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_StreamTapConfigsClient interface {
	Send(*v3alpha1.DiscoveryRequest) error
	Recv() (*v3alpha1.DiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceStreamTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Send(m *v3alpha1.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsClient) Recv() (*v3alpha1.DiscoveryResponse, error) {
	m := new(v3alpha1.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) DeltaTapConfigs(ctx context.Context, opts ...grpc.CallOption) (TapDiscoveryService_DeltaTapConfigsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TapDiscoveryService_serviceDesc.Streams[1], "/envoy.service.tap.v3alpha.TapDiscoveryService/DeltaTapConfigs", opts...)
	if err != nil {
		return nil, err
	}
	x := &tapDiscoveryServiceDeltaTapConfigsClient{stream}
	return x, nil
}

type TapDiscoveryService_DeltaTapConfigsClient interface {
	Send(*v3alpha1.DeltaDiscoveryRequest) error
	Recv() (*v3alpha1.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type tapDiscoveryServiceDeltaTapConfigsClient struct {
	grpc.ClientStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Send(m *v3alpha1.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsClient) Recv() (*v3alpha1.DeltaDiscoveryResponse, error) {
	m := new(v3alpha1.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tapDiscoveryServiceClient) FetchTapConfigs(ctx context.Context, in *v3alpha1.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha1.DiscoveryResponse, error) {
	out := new(v3alpha1.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.tap.v3alpha.TapDiscoveryService/FetchTapConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDiscoveryServiceServer is the server API for TapDiscoveryService service.
type TapDiscoveryServiceServer interface {
	StreamTapConfigs(TapDiscoveryService_StreamTapConfigsServer) error
	DeltaTapConfigs(TapDiscoveryService_DeltaTapConfigsServer) error
	FetchTapConfigs(context.Context, *v3alpha1.DiscoveryRequest) (*v3alpha1.DiscoveryResponse, error)
}

// UnimplementedTapDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTapDiscoveryServiceServer struct {
}

func (*UnimplementedTapDiscoveryServiceServer) StreamTapConfigs(srv TapDiscoveryService_StreamTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) DeltaTapConfigs(srv TapDiscoveryService_DeltaTapConfigsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaTapConfigs not implemented")
}
func (*UnimplementedTapDiscoveryServiceServer) FetchTapConfigs(ctx context.Context, req *v3alpha1.DiscoveryRequest) (*v3alpha1.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchTapConfigs not implemented")
}

func RegisterTapDiscoveryServiceServer(s *grpc.Server, srv TapDiscoveryServiceServer) {
	s.RegisterService(&_TapDiscoveryService_serviceDesc, srv)
}

func _TapDiscoveryService_StreamTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).StreamTapConfigs(&tapDiscoveryServiceStreamTapConfigsServer{stream})
}

type TapDiscoveryService_StreamTapConfigsServer interface {
	Send(*v3alpha1.DiscoveryResponse) error
	Recv() (*v3alpha1.DiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceStreamTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Send(m *v3alpha1.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceStreamTapConfigsServer) Recv() (*v3alpha1.DiscoveryRequest, error) {
	m := new(v3alpha1.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_DeltaTapConfigs_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TapDiscoveryServiceServer).DeltaTapConfigs(&tapDiscoveryServiceDeltaTapConfigsServer{stream})
}

type TapDiscoveryService_DeltaTapConfigsServer interface {
	Send(*v3alpha1.DeltaDiscoveryResponse) error
	Recv() (*v3alpha1.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type tapDiscoveryServiceDeltaTapConfigsServer struct {
	grpc.ServerStream
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Send(m *v3alpha1.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tapDiscoveryServiceDeltaTapConfigsServer) Recv() (*v3alpha1.DeltaDiscoveryRequest, error) {
	m := new(v3alpha1.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TapDiscoveryService_FetchTapConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3alpha1.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.tap.v3alpha.TapDiscoveryService/FetchTapConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDiscoveryServiceServer).FetchTapConfigs(ctx, req.(*v3alpha1.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TapDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.tap.v3alpha.TapDiscoveryService",
	HandlerType: (*TapDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchTapConfigs",
			Handler:    _TapDiscoveryService_FetchTapConfigs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTapConfigs",
			Handler:       _TapDiscoveryService_StreamTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaTapConfigs",
			Handler:       _TapDiscoveryService_DeltaTapConfigs_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/tap/v3alpha/tapds.proto",
}
