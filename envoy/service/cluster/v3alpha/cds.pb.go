// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/cluster/v3alpha/cds.proto

package envoy_service_cluster_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3alpha"
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

type CdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CdsDummy) Reset()         { *m = CdsDummy{} }
func (m *CdsDummy) String() string { return proto.CompactTextString(m) }
func (*CdsDummy) ProtoMessage()    {}
func (*CdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e2d75d1d4b158d, []int{0}
}

func (m *CdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CdsDummy.Unmarshal(m, b)
}
func (m *CdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CdsDummy.Marshal(b, m, deterministic)
}
func (m *CdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CdsDummy.Merge(m, src)
}
func (m *CdsDummy) XXX_Size() int {
	return xxx_messageInfo_CdsDummy.Size(m)
}
func (m *CdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_CdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_CdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CdsDummy)(nil), "envoy.service.cluster.v3alpha.CdsDummy")
}

func init() {
	proto.RegisterFile("envoy/service/cluster/v3alpha/cds.proto", fileDescriptor_64e2d75d1d4b158d)
}

var fileDescriptor_64e2d75d1d4b158d = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x4a, 0x23, 0x31,
	0x1c, 0xc7, 0x37, 0x5b, 0xd8, 0x5d, 0x02, 0xdd, 0xc3, 0xc0, 0x52, 0x98, 0x6d, 0x97, 0xdd, 0xb2,
	0xec, 0x16, 0x2d, 0x89, 0xb6, 0xa0, 0x50, 0x3c, 0xb5, 0xc5, 0x73, 0xb1, 0x4f, 0x10, 0x67, 0x7e,
	0x4e, 0x03, 0xd3, 0x64, 0x4c, 0x32, 0x83, 0x73, 0xf0, 0xe2, 0x41, 0x8a, 0x57, 0xc1, 0x83, 0xe7,
	0xbe, 0x8a, 0x4f, 0xe0, 0x2b, 0xf8, 0x06, 0x82, 0x67, 0x69, 0x93, 0x8e, 0xad, 0x05, 0xa5, 0x17,
	0xaf, 0xc9, 0x27, 0xdf, 0x3f, 0xe1, 0x8b, 0xff, 0x83, 0xc8, 0x64, 0x4e, 0x35, 0xa8, 0x8c, 0x07,
	0x40, 0x83, 0x38, 0xd5, 0x06, 0x14, 0xcd, 0xda, 0x2c, 0x4e, 0x46, 0x8c, 0x06, 0xa1, 0x26, 0x89,
	0x92, 0x46, 0x7a, 0xb5, 0x39, 0x48, 0x1c, 0x48, 0x1c, 0x48, 0x1c, 0xe8, 0xd3, 0x55, 0x9d, 0x90,
	0xeb, 0x40, 0x66, 0xa0, 0xf2, 0x42, 0xa9, 0x38, 0xb1, 0x7a, 0x7e, 0x35, 0x92, 0x32, 0x8a, 0x81,
	0xb2, 0x84, 0x53, 0x26, 0x84, 0x34, 0xcc, 0x70, 0x29, 0x9c, 0x9b, 0xff, 0x27, 0x0d, 0x13, 0xb6,
	0x7c, 0x4e, 0x33, 0x50, 0x9a, 0x4b, 0xc1, 0x45, 0xe4, 0x90, 0xdf, 0xd6, 0x71, 0x99, 0x51, 0xa0,
	0x65, 0xaa, 0x02, 0xb0, 0x44, 0xbd, 0x81, 0xbf, 0xf5, 0x42, 0xdd, 0x4f, 0xc7, 0xe3, 0xbc, 0x53,
	0xbd, 0xbd, 0x9b, 0xfc, 0xaa, 0xe0, 0x1f, 0xb6, 0x05, 0x4b, 0x38, 0xc9, 0x5a, 0x64, 0x71, 0xdb,
	0x7a, 0x2a, 0xe1, 0x4a, 0xcf, 0x36, 0xea, 0x2f, 0x72, 0x0e, 0x6d, 0x15, 0xef, 0x1c, 0x7f, 0x1f,
	0x1a, 0x05, 0x6c, 0xec, 0x00, 0xed, 0xed, 0x92, 0xd5, 0xbf, 0x78, 0xa9, 0xe6, 0xca, 0x92, 0x42,
	0xe4, 0x08, 0x4e, 0x53, 0xd0, 0xc6, 0x6f, 0x6d, 0xf2, 0x44, 0x27, 0x52, 0x68, 0xa8, 0x7f, 0x6a,
	0xa0, 0x1d, 0xe4, 0x5d, 0x22, 0x5c, 0xee, 0x43, 0x6c, 0x58, 0x61, 0xbf, 0xf7, 0xbe, 0xd6, 0x8c,
	0x5f, 0xcb, 0xb0, 0xbf, 0xf1, 0xbb, 0x95, 0x20, 0x53, 0x84, 0xcb, 0x87, 0x60, 0x82, 0xd1, 0x47,
	0xff, 0x43, 0xf3, 0xe2, 0xfe, 0xe1, 0xfa, 0x73, 0xad, 0xfe, 0x73, 0x7d, 0x4b, 0x1d, 0x37, 0x43,
	0x3d, 0x47, 0x4a, 0x1d, 0xb4, 0xe5, 0x37, 0xaf, 0xa6, 0x37, 0x8f, 0x5f, 0xff, 0xe1, 0xbf, 0xd6,
	0x28, 0x90, 0xe2, 0x84, 0x47, 0xaf, 0xe7, 0x4a, 0x5c, 0x87, 0xee, 0x01, 0xde, 0xe6, 0xd2, 0x66,
	0x4a, 0x94, 0x3c, 0xcb, 0xc9, 0x9b, 0x2b, 0xef, 0xce, 0xf6, 0x34, 0x98, 0x6d, 0x6b, 0x80, 0x26,
	0x08, 0x1d, 0x7f, 0x99, 0xef, 0xac, 0xfd, 0x1c, 0x00, 0x00, 0xff, 0xff, 0x86, 0xd4, 0x6b, 0xf6,
	0x45, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterDiscoveryServiceClient is the client API for ClusterDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterDiscoveryServiceClient interface {
	StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error)
	DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error)
	FetchClusters(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error)
}

type clusterDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterDiscoveryServiceClient(cc *grpc.ClientConn) ClusterDiscoveryServiceClient {
	return &clusterDiscoveryServiceClient{cc}
}

func (c *clusterDiscoveryServiceClient) StreamClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_StreamClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[0], "/envoy.service.cluster.v3alpha.ClusterDiscoveryService/StreamClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceStreamClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_StreamClustersClient interface {
	Send(*v3alpha.DiscoveryRequest) error
	Recv() (*v3alpha.DiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceStreamClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceStreamClustersClient) Send(m *v3alpha.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersClient) Recv() (*v3alpha.DiscoveryResponse, error) {
	m := new(v3alpha.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) DeltaClusters(ctx context.Context, opts ...grpc.CallOption) (ClusterDiscoveryService_DeltaClustersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClusterDiscoveryService_serviceDesc.Streams[1], "/envoy.service.cluster.v3alpha.ClusterDiscoveryService/DeltaClusters", opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterDiscoveryServiceDeltaClustersClient{stream}
	return x, nil
}

type ClusterDiscoveryService_DeltaClustersClient interface {
	Send(*v3alpha.DeltaDiscoveryRequest) error
	Recv() (*v3alpha.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type clusterDiscoveryServiceDeltaClustersClient struct {
	grpc.ClientStream
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Send(m *v3alpha.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersClient) Recv() (*v3alpha.DeltaDiscoveryResponse, error) {
	m := new(v3alpha.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterDiscoveryServiceClient) FetchClusters(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error) {
	out := new(v3alpha.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.cluster.v3alpha.ClusterDiscoveryService/FetchClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterDiscoveryServiceServer is the server API for ClusterDiscoveryService service.
type ClusterDiscoveryServiceServer interface {
	StreamClusters(ClusterDiscoveryService_StreamClustersServer) error
	DeltaClusters(ClusterDiscoveryService_DeltaClustersServer) error
	FetchClusters(context.Context, *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error)
}

// UnimplementedClusterDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClusterDiscoveryServiceServer struct {
}

func (*UnimplementedClusterDiscoveryServiceServer) StreamClusters(srv ClusterDiscoveryService_StreamClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) DeltaClusters(srv ClusterDiscoveryService_DeltaClustersServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaClusters not implemented")
}
func (*UnimplementedClusterDiscoveryServiceServer) FetchClusters(ctx context.Context, req *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchClusters not implemented")
}

func RegisterClusterDiscoveryServiceServer(s *grpc.Server, srv ClusterDiscoveryServiceServer) {
	s.RegisterService(&_ClusterDiscoveryService_serviceDesc, srv)
}

func _ClusterDiscoveryService_StreamClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).StreamClusters(&clusterDiscoveryServiceStreamClustersServer{stream})
}

type ClusterDiscoveryService_StreamClustersServer interface {
	Send(*v3alpha.DiscoveryResponse) error
	Recv() (*v3alpha.DiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceStreamClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceStreamClustersServer) Send(m *v3alpha.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceStreamClustersServer) Recv() (*v3alpha.DiscoveryRequest, error) {
	m := new(v3alpha.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_DeltaClusters_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClusterDiscoveryServiceServer).DeltaClusters(&clusterDiscoveryServiceDeltaClustersServer{stream})
}

type ClusterDiscoveryService_DeltaClustersServer interface {
	Send(*v3alpha.DeltaDiscoveryResponse) error
	Recv() (*v3alpha.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type clusterDiscoveryServiceDeltaClustersServer struct {
	grpc.ServerStream
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Send(m *v3alpha.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clusterDiscoveryServiceDeltaClustersServer) Recv() (*v3alpha.DeltaDiscoveryRequest, error) {
	m := new(v3alpha.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClusterDiscoveryService_FetchClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3alpha.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.cluster.v3alpha.ClusterDiscoveryService/FetchClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterDiscoveryServiceServer).FetchClusters(ctx, req.(*v3alpha.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.cluster.v3alpha.ClusterDiscoveryService",
	HandlerType: (*ClusterDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchClusters",
			Handler:    _ClusterDiscoveryService_FetchClusters_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamClusters",
			Handler:       _ClusterDiscoveryService_StreamClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaClusters",
			Handler:       _ClusterDiscoveryService_DeltaClusters_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/cluster/v3alpha/cds.proto",
}
