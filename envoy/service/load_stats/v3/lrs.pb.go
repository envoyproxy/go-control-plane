// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/load_stats/v3/lrs.proto

package envoy_service_load_stats_v3

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type LoadStatsRequest struct {
	Node                 *v3.Node            `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClusterStats         []*v31.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a87f04ee5f8fca3, []int{0}
}

func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsRequest.Unmarshal(m, b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
}
func (m *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(m, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return xxx_messageInfo_LoadStatsRequest.Size(m)
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *v3.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*v31.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

type LoadStatsResponse struct {
	Clusters                  []string           `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	SendAllClusters           bool               `protobuf:"varint,4,opt,name=send_all_clusters,json=sendAllClusters,proto3" json:"send_all_clusters,omitempty"`
	LoadReportingInterval     *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	ReportEndpointGranularity bool               `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a87f04ee5f8fca3, []int{1}
}

func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsResponse.Unmarshal(m, b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
}
func (m *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(m, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return xxx_messageInfo_LoadStatsResponse.Size(m)
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetSendAllClusters() bool {
	if m != nil {
		return m.SendAllClusters
	}
	return false
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func (m *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if m != nil {
		return m.ReportEndpointGranularity
	}
	return false
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v3.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v3.LoadStatsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v3/lrs.proto", fileDescriptor_4a87f04ee5f8fca3)
}

var fileDescriptor_4a87f04ee5f8fca3 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x71, 0x3b, 0x4d, 0xc5, 0x03, 0x6d, 0x8b, 0x40, 0xeb, 0x32, 0x31, 0x4a, 0x25, 0x50,
	0x99, 0x98, 0x8d, 0x1a, 0x4e, 0x3b, 0x20, 0xd1, 0x81, 0x10, 0xa2, 0x42, 0x23, 0xfd, 0x00, 0x91,
	0x9b, 0x78, 0x91, 0x25, 0xe3, 0x17, 0x6c, 0x27, 0xa2, 0x37, 0x8e, 0x3b, 0xf1, 0x01, 0x38, 0xf3,
	0x29, 0x38, 0x83, 0xc4, 0x95, 0x6f, 0x84, 0x1c, 0x27, 0x5d, 0xc7, 0xb4, 0x69, 0xb7, 0xd8, 0xef,
	0xf7, 0x7f, 0xfe, 0xbf, 0xff, 0x0b, 0x7e, 0xcc, 0x55, 0x05, 0x0b, 0x6a, 0xb8, 0xae, 0x44, 0xca,
	0xa9, 0x04, 0x96, 0x25, 0xc6, 0x32, 0x6b, 0x68, 0x15, 0x51, 0xa9, 0x0d, 0x29, 0x34, 0x58, 0x08,
	0xf6, 0x6a, 0x8c, 0x34, 0x18, 0x39, 0xc7, 0x48, 0x15, 0x85, 0x0f, 0x7d, 0x8f, 0x14, 0xd4, 0xa9,
	0xc8, 0x69, 0x0a, 0x9a, 0x3b, 0xf1, 0x9c, 0x19, 0xee, 0xd5, 0xe1, 0xc1, 0x05, 0x80, 0xab, 0xac,
	0x00, 0xa1, 0x6c, 0xfd, 0x82, 0x6b, 0xa4, 0x79, 0x01, 0xda, 0x36, 0xec, 0x7e, 0x0e, 0x90, 0x4b,
	0x4e, 0xeb, 0xd3, 0xbc, 0x3c, 0xa5, 0x59, 0xa9, 0x99, 0x15, 0xa0, 0x9a, 0xfa, 0x83, 0x32, 0x2b,
	0x18, 0x65, 0x4a, 0x81, 0xad, 0xaf, 0x0d, 0x75, 0x3e, 0xca, 0xc6, 0x68, 0xf8, 0xe8, 0x52, 0xb9,
	0xe2, 0xda, 0x08, 0x50, 0x42, 0xe5, 0x0d, 0xb2, 0x53, 0x31, 0x29, 0x32, 0x66, 0x39, 0x6d, 0x3f,
	0x7c, 0x61, 0xf8, 0x0b, 0xe1, 0xad, 0x29, 0xb0, 0x6c, 0xe6, 0x06, 0x8b, 0xf9, 0xe7, 0x92, 0x1b,
	0x1b, 0x10, 0xbc, 0xa6, 0x20, 0xe3, 0x7d, 0x34, 0x40, 0xa3, 0x8d, 0x71, 0x48, 0x7c, 0x10, 0x7e,
	0x14, 0xe2, 0x66, 0x25, 0x55, 0x44, 0x3e, 0x40, 0xc6, 0xe3, 0x9a, 0x0b, 0xde, 0xe3, 0xbb, 0xa9,
	0x2c, 0x8d, 0xe5, 0xda, 0x07, 0xd4, 0xef, 0x0c, 0xba, 0xa3, 0x8d, 0xf1, 0x93, 0x8b, 0xc2, 0x36,
	0x03, 0x27, 0x3e, 0xf6, 0xb8, 0x7f, 0xf5, 0x4e, 0xba, 0x72, 0x3a, 0x8a, 0xbe, 0xff, 0x3e, 0xdb,
	0x27, 0xf8, 0xd9, 0xd5, 0xe9, 0x8f, 0xc9, 0xff, 0x8e, 0x87, 0x3f, 0x3a, 0x78, 0x7b, 0xe5, 0xd2,
	0x14, 0xa0, 0x0c, 0x0f, 0x42, 0xdc, 0x6b, 0x5a, 0x9b, 0x3e, 0x1a, 0x74, 0x47, 0xb7, 0xe3, 0xe5,
	0x39, 0x38, 0xc0, 0xdb, 0x86, 0xab, 0x2c, 0x61, 0x52, 0x26, 0x4b, 0x68, 0x6d, 0x80, 0x46, 0xbd,
	0x78, 0xd3, 0x15, 0x5e, 0x49, 0x79, 0xdc, 0xb2, 0x1f, 0xf1, 0xce, 0xca, 0xd2, 0x84, 0xca, 0x13,
	0xa1, 0x2c, 0xd7, 0x15, 0x93, 0xfd, 0x4e, 0x1d, 0xd1, 0x2e, 0xf1, 0x1b, 0x24, 0xed, 0x06, 0xc9,
	0xeb, 0x66, 0x83, 0xf1, 0x7d, 0xa7, 0x8c, 0x5b, 0xe1, 0xbb, 0x46, 0x17, 0xbc, 0xc4, 0x7b, 0xbe,
	0x5b, 0xd2, 0xc6, 0x92, 0xe4, 0x9a, 0xa9, 0x52, 0x32, 0x2d, 0xec, 0xa2, 0xdf, 0xad, 0x8d, 0xec,
	0x7a, 0xe4, 0x4d, 0x43, 0xbc, 0x3d, 0x07, 0x8e, 0x5e, 0xb8, 0x94, 0x28, 0x3e, 0xbc, 0x61, 0x4a,
	0x3e, 0x90, 0xf1, 0x37, 0x84, 0xef, 0x4d, 0x57, 0xfd, 0xcc, 0xbc, 0x30, 0xa8, 0xf0, 0xe6, 0xcc,
	0x6a, 0xce, 0x3e, 0x2d, 0x35, 0xc1, 0x21, 0xb9, 0xe6, 0xff, 0xbf, 0xb4, 0x81, 0x90, 0xdc, 0x14,
	0xf7, 0x56, 0x86, 0xb7, 0x46, 0xe8, 0x39, 0x9a, 0x4c, 0x7e, 0x7e, 0xfd, 0xf3, 0x77, 0xbd, 0xb3,
	0xd5, 0xc1, 0x4f, 0x05, 0xf8, 0x0e, 0x85, 0x86, 0x2f, 0x8b, 0xeb, 0x9a, 0x4d, 0x7a, 0x53, 0x6d,
	0x4e, 0x5c, 0xd0, 0x27, 0xe8, 0x0c, 0xa1, 0xf9, 0x7a, 0x1d, 0x7a, 0xf4, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x8c, 0xd6, 0x6d, 0x18, 0xd7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v3.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

// UnimplementedLoadReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadReportingServiceServer struct {
}

func (*UnimplementedLoadReportingServiceServer) StreamLoadStats(srv LoadReportingService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v3.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v3/lrs.proto",
}
