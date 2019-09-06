// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v2/rtds.proto

package envoy_service_discovery_v2

import (
	context "context"
	fmt "fmt"
	v2 "github.com/envoyproxy/go-control-plane/v2/envoy/api/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type RtdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RtdsDummy) Reset()         { *m = RtdsDummy{} }
func (m *RtdsDummy) String() string { return proto.CompactTextString(m) }
func (*RtdsDummy) ProtoMessage()    {}
func (*RtdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ef541a8dd8cfbb, []int{0}
}

func (m *RtdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RtdsDummy.Unmarshal(m, b)
}
func (m *RtdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RtdsDummy.Marshal(b, m, deterministic)
}
func (m *RtdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RtdsDummy.Merge(m, src)
}
func (m *RtdsDummy) XXX_Size() int {
	return xxx_messageInfo_RtdsDummy.Size(m)
}
func (m *RtdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_RtdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_RtdsDummy proto.InternalMessageInfo

type Runtime struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Layer                *_struct.Struct `protobuf:"bytes,2,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_60ef541a8dd8cfbb, []int{1}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Runtime) GetLayer() *_struct.Struct {
	if m != nil {
		return m.Layer
	}
	return nil
}

func init() {
	proto.RegisterType((*RtdsDummy)(nil), "envoy.service.discovery.v2.RtdsDummy")
	proto.RegisterType((*Runtime)(nil), "envoy.service.discovery.v2.Runtime")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/rtds.proto", fileDescriptor_60ef541a8dd8cfbb)
}

var fileDescriptor_60ef541a8dd8cfbb = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xbf, 0x4e, 0xe3, 0x30,
	0x18, 0x3f, 0x47, 0xbd, 0xab, 0xea, 0xf6, 0x96, 0x48, 0xa7, 0x56, 0xb9, 0xea, 0x2e, 0x0a, 0x20,
	0x45, 0x48, 0x38, 0x28, 0x6c, 0x1d, 0xa3, 0x8a, 0xb9, 0x4a, 0x05, 0x2b, 0x72, 0x93, 0x8f, 0x62,
	0x29, 0x89, 0x53, 0xdb, 0x89, 0xc8, 0xca, 0xc4, 0xce, 0x33, 0xf1, 0x04, 0xbc, 0x02, 0x4f, 0xc1,
	0x84, 0x6a, 0x37, 0x45, 0x2d, 0x2a, 0x13, 0x9b, 0xed, 0xdf, 0x3f, 0x7f, 0xfe, 0x19, 0x9f, 0x40,
	0x51, 0xf3, 0x26, 0x90, 0x20, 0x6a, 0x96, 0x40, 0x90, 0x32, 0x99, 0xf0, 0x1a, 0x44, 0x13, 0xd4,
	0x61, 0x20, 0x54, 0x2a, 0x49, 0x29, 0xb8, 0xe2, 0xb6, 0xa3, 0x69, 0x64, 0x43, 0x23, 0x5b, 0x1a,
	0xa9, 0x43, 0x67, 0x6c, 0x2c, 0x68, 0xc9, 0xd6, 0xa2, 0x0f, 0x48, 0x2b, 0x9d, 0xf1, 0x92, 0xf3,
	0x65, 0x06, 0x1a, 0xa6, 0x45, 0xc1, 0x15, 0x55, 0x8c, 0x17, 0x72, 0x0f, 0xd5, 0xbb, 0x45, 0x75,
	0x1b, 0x48, 0x25, 0xaa, 0x44, 0x6d, 0xd0, 0x61, 0x4d, 0x33, 0x96, 0x52, 0x05, 0x41, 0xbb, 0x30,
	0x80, 0xd7, 0xc7, 0xbd, 0x58, 0xa5, 0x72, 0x5a, 0xe5, 0x79, 0xe3, 0x5d, 0xe1, 0x6e, 0x5c, 0x15,
	0x8a, 0xe5, 0x60, 0xff, 0xc5, 0x9d, 0x82, 0xe6, 0x30, 0x42, 0x2e, 0xf2, 0x7b, 0x51, 0xf7, 0x2d,
	0xea, 0x08, 0xcb, 0x45, 0xb1, 0x3e, 0xb4, 0xcf, 0xf0, 0xcf, 0x8c, 0x36, 0x20, 0x46, 0x96, 0x8b,
	0xfc, 0x7e, 0x38, 0x24, 0x26, 0x9b, 0xb4, 0xd9, 0x64, 0xae, 0xb3, 0x63, 0xc3, 0x0a, 0x9f, 0x2d,
	0x3c, 0xdc, 0xf8, 0x4e, 0xdb, 0x99, 0xe6, 0x66, 0x7e, 0xfb, 0x1a, 0xff, 0x9e, 0x2b, 0x01, 0x34,
	0x6f, 0x83, 0xff, 0x11, 0xf3, 0x40, 0xb4, 0x64, 0xa4, 0x0e, 0xc9, 0x56, 0x10, 0xc3, 0xaa, 0x02,
	0xa9, 0x9c, 0xff, 0x07, 0x71, 0x59, 0xf2, 0x42, 0x82, 0xf7, 0xc3, 0x47, 0xe7, 0xc8, 0xbe, 0xc1,
	0x83, 0x29, 0x64, 0x8a, 0xb6, 0xb6, 0x47, 0x7b, 0xb2, 0x35, 0xf6, 0xc9, 0xfb, 0xf8, 0x6b, 0xd2,
	0x4e, 0xc0, 0x0a, 0x0f, 0x2e, 0x41, 0x25, 0x77, 0xdf, 0x76, 0x6f, 0xf7, 0xe1, 0xe5, 0xf5, 0xc9,
	0x72, 0xbc, 0x3f, 0x3b, 0xdd, 0x4f, 0x84, 0xf1, 0x9f, 0xa0, 0xd3, 0x68, 0x82, 0x7d, 0xc6, 0x8d,
	0x4d, 0x29, 0xf8, 0x7d, 0x43, 0x0e, 0x7f, 0xa5, 0x48, 0xb7, 0x3a, 0x5b, 0xf7, 0x31, 0x43, 0x8f,
	0x08, 0x2d, 0x7e, 0xe9, 0x6e, 0x2e, 0xde, 0x03, 0x00, 0x00, 0xff, 0xff, 0x51, 0x02, 0x80, 0xf6,
	0xa6, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeDiscoveryServiceClient is the client API for RuntimeDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeDiscoveryServiceClient interface {
	StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error)
	DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error)
	FetchRuntime(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type runtimeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeDiscoveryServiceClient(cc *grpc.ClientConn) RuntimeDiscoveryServiceClient {
	return &runtimeDiscoveryServiceClient{cc}
}

func (c *runtimeDiscoveryServiceClient) StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.RuntimeDiscoveryService/StreamRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceStreamRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_StreamRuntimeClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceStreamRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v2.RuntimeDiscoveryService/DeltaRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceDeltaRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_DeltaRuntimeClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceDeltaRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) FetchRuntime(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v2.RuntimeDiscoveryService/FetchRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeDiscoveryServiceServer is the server API for RuntimeDiscoveryService service.
type RuntimeDiscoveryServiceServer interface {
	StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error
	DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error
	FetchRuntime(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterRuntimeDiscoveryServiceServer(s *grpc.Server, srv RuntimeDiscoveryServiceServer) {
	s.RegisterService(&_RuntimeDiscoveryService_serviceDesc, srv)
}

func _RuntimeDiscoveryService_StreamRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).StreamRuntime(&runtimeDiscoveryServiceStreamRuntimeServer{stream})
}

type RuntimeDiscoveryService_StreamRuntimeServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceStreamRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_DeltaRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).DeltaRuntime(&runtimeDiscoveryServiceDeltaRuntimeServer{stream})
}

type RuntimeDiscoveryService_DeltaRuntimeServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceDeltaRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_FetchRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.RuntimeDiscoveryService/FetchRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.RuntimeDiscoveryService",
	HandlerType: (*RuntimeDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRuntime",
			Handler:    _RuntimeDiscoveryService_FetchRuntime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRuntime",
			Handler:       _RuntimeDiscoveryService_StreamRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRuntime",
			Handler:       _RuntimeDiscoveryService_DeltaRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/rtds.proto",
}
