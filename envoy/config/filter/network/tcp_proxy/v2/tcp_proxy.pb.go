// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto

package envoy_config_filter_network_tcp_proxy_v2

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TcpProxy struct {
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*TcpProxy_Cluster
	//	*TcpProxy_WeightedClusters
	ClusterSpecifier      isTcpProxy_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch         *core.Metadata              `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	IdleTimeout           *duration.Duration          `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	DownstreamIdleTimeout *duration.Duration          `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout,proto3" json:"downstream_idle_timeout,omitempty"`
	UpstreamIdleTimeout   *duration.Duration          `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout,proto3" json:"upstream_idle_timeout,omitempty"`
	AccessLog             []*v2.AccessLog             `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	DeprecatedV1          *TcpProxy_DeprecatedV1      `protobuf:"bytes,6,opt,name=deprecated_v1,json=deprecatedV1,proto3" json:"deprecated_v1,omitempty"` // Deprecated: Do not use.
	MaxConnectAttempts    *wrappers.UInt32Value       `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                    `json:"-"`
	XXX_unrecognized      []byte                      `json:"-"`
	XXX_sizecache         int32                       `json:"-"`
}

func (m *TcpProxy) Reset()         { *m = TcpProxy{} }
func (m *TcpProxy) String() string { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()    {}
func (*TcpProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0}
}

func (m *TcpProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy.Unmarshal(m, b)
}
func (m *TcpProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy.Marshal(b, m, deterministic)
}
func (m *TcpProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy.Merge(m, src)
}
func (m *TcpProxy) XXX_Size() int {
	return xxx_messageInfo_TcpProxy.Size(m)
}
func (m *TcpProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy proto.InternalMessageInfo

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isTcpProxy_ClusterSpecifier interface {
	isTcpProxy_ClusterSpecifier()
}

type TcpProxy_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type TcpProxy_WeightedClusters struct {
	WeightedClusters *TcpProxy_WeightedCluster `protobuf:"bytes,10,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*TcpProxy_Cluster) isTcpProxy_ClusterSpecifier() {}

func (*TcpProxy_WeightedClusters) isTcpProxy_ClusterSpecifier() {}

func (m *TcpProxy) GetClusterSpecifier() isTcpProxy_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *TcpProxy) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *TcpProxy) GetWeightedClusters() *TcpProxy_WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *TcpProxy) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*v2.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

// Deprecated: Do not use.
func (m *TcpProxy) GetDeprecatedV1() *TcpProxy_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TcpProxy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TcpProxy_Cluster)(nil),
		(*TcpProxy_WeightedClusters)(nil),
	}
}

type TcpProxy_DeprecatedV1 struct {
	Routes               []*TcpProxy_DeprecatedV1_TCPRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1) Reset()         { *m = TcpProxy_DeprecatedV1{} }
func (m *TcpProxy_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0}
}

func (m *TcpProxy_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Size(m)
}
func (m *TcpProxy_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1 proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1) GetRoutes() []*TcpProxy_DeprecatedV1_TCPRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

type TcpProxy_DeprecatedV1_TCPRoute struct {
	Cluster              string            `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	DestinationIpList    []*core.CidrRange `protobuf:"bytes,2,rep,name=destination_ip_list,json=destinationIpList,proto3" json:"destination_ip_list,omitempty"`
	DestinationPorts     string            `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	SourceIpList         []*core.CidrRange `protobuf:"bytes,4,rep,name=source_ip_list,json=sourceIpList,proto3" json:"source_ip_list,omitempty"`
	SourcePorts          string            `protobuf:"bytes,5,opt,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Reset()         { *m = TcpProxy_DeprecatedV1_TCPRoute{} }
func (m *TcpProxy_DeprecatedV1_TCPRoute) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1_TCPRoute) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1_TCPRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 0, 0}
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Size(m)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationIpList() []*core.CidrRange {
	if m != nil {
		return m.DestinationIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationPorts() string {
	if m != nil {
		return m.DestinationPorts
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourceIpList() []*core.CidrRange {
	if m != nil {
		return m.SourceIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourcePorts() string {
	if m != nil {
		return m.SourcePorts
	}
	return ""
}

type TcpProxy_WeightedCluster struct {
	Clusters             []*TcpProxy_WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *TcpProxy_WeightedCluster) Reset()         { *m = TcpProxy_WeightedCluster{} }
func (m *TcpProxy_WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1}
}

func (m *TcpProxy_WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Size(m)
}
func (m *TcpProxy_WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster) GetClusters() []*TcpProxy_WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TcpProxy_WeightedCluster_ClusterWeight struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) Reset() {
	*m = TcpProxy_WeightedCluster_ClusterWeight{}
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f6b35dbcbad27ba, []int{0, 1, 0}
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Size(m)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.DeprecatedV1.TCPRoute")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.tcp_proxy.v2.TcpProxy.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_proxy/v2/tcp_proxy.proto", fileDescriptor_1f6b35dbcbad27ba)
}

var fileDescriptor_1f6b35dbcbad27ba = []byte{
	// 768 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4f, 0x6f, 0xdb, 0x36,
	0x14, 0x8f, 0x64, 0xc7, 0xb1, 0x69, 0x7b, 0x6b, 0xd8, 0x15, 0xd5, 0xdc, 0x60, 0x75, 0x77, 0x32,
	0x5a, 0x40, 0x5a, 0x5d, 0x60, 0xd8, 0x6d, 0xa8, 0x1c, 0x0c, 0xf1, 0x10, 0x03, 0x9e, 0x90, 0x25,
	0x47, 0x81, 0x91, 0x68, 0x85, 0x9b, 0x24, 0x72, 0x24, 0x65, 0x3b, 0xf7, 0x1d, 0x07, 0x0c, 0xd8,
	0xc7, 0xd8, 0x27, 0xda, 0x47, 0x19, 0x72, 0x1a, 0x44, 0x52, 0xb6, 0x92, 0x78, 0x48, 0xd0, 0x9c,
	0x24, 0xbe, 0xf7, 0x7e, 0xbf, 0x1f, 0xdf, 0x3f, 0x82, 0xef, 0x70, 0xbe, 0xa4, 0xd7, 0x5e, 0x44,
	0xf3, 0x05, 0x49, 0xbc, 0x05, 0x49, 0x25, 0xe6, 0x5e, 0x8e, 0xe5, 0x8a, 0xf2, 0x5f, 0x3d, 0x19,
	0xb1, 0x90, 0x71, 0xba, 0xbe, 0xf6, 0x96, 0xe3, 0xed, 0xc1, 0x65, 0x9c, 0x4a, 0x0a, 0x47, 0x0a,
	0xe9, 0x6a, 0xa4, 0xab, 0x91, 0xae, 0x41, 0xba, 0xdb, 0xe0, 0xe5, 0x78, 0xf0, 0xcd, 0x2e, 0x0d,
	0x14, 0x45, 0x58, 0x88, 0x94, 0x26, 0x25, 0xf7, 0xe6, 0xa0, 0xb9, 0x07, 0xaf, 0x35, 0x02, 0x31,
	0x52, 0x7a, 0x23, 0xca, 0xb1, 0x87, 0xe2, 0x98, 0x63, 0x21, 0x4c, 0xc0, 0xd1, 0xfd, 0x80, 0x4b,
	0x24, 0xb0, 0xf1, 0x7e, 0x95, 0x50, 0x9a, 0xa4, 0xd8, 0x53, 0xa7, 0xcb, 0x62, 0xe1, 0xc5, 0x05,
	0x47, 0x92, 0xd0, 0xfc, 0xff, 0xfc, 0x2b, 0x8e, 0x18, 0xc3, 0xbc, 0x62, 0x7f, 0xb9, 0x44, 0x29,
	0x89, 0x91, 0xc4, 0x5e, 0xf5, 0xa3, 0x1d, 0x5f, 0xff, 0xde, 0x05, 0xed, 0xb3, 0x88, 0xcd, 0xcb,
	0xcc, 0xe0, 0x08, 0x74, 0x85, 0x44, 0x32, 0x64, 0x1c, 0x2f, 0xc8, 0xda, 0xb1, 0x86, 0xd6, 0xa8,
	0xe3, 0x1f, 0xdc, 0xf8, 0x4d, 0x6e, 0x0f, 0xad, 0x00, 0x94, 0xbe, 0xb9, 0x72, 0xc1, 0x01, 0x38,
	0x88, 0xd2, 0x42, 0x48, 0xcc, 0x1d, 0xbb, 0x8c, 0x3a, 0xd9, 0x0b, 0x2a, 0x03, 0xfc, 0x0d, 0x1c,
	0xae, 0x30, 0x49, 0xae, 0x24, 0x8e, 0x43, 0x63, 0x13, 0x0e, 0x18, 0x5a, 0xa3, 0xee, 0xd8, 0x77,
	0x1f, 0x5b, 0x62, 0xb7, 0xba, 0x94, 0x7b, 0x61, 0xb8, 0x26, 0x9a, 0xea, 0x64, 0x2f, 0x78, 0xb6,
	0xba, 0x6d, 0x12, 0xd0, 0x07, 0x9f, 0x65, 0x58, 0xa2, 0x18, 0x49, 0x14, 0x66, 0x48, 0x46, 0x57,
	0x4e, 0x47, 0xe9, 0xbd, 0x32, 0x7a, 0x88, 0x91, 0x92, 0xb3, 0xac, 0xaa, 0x3b, 0x33, 0x81, 0x41,
	0xbf, 0x82, 0xcc, 0x4a, 0x04, 0xfc, 0x01, 0xf4, 0x48, 0x9c, 0xe2, 0x50, 0x92, 0x0c, 0xd3, 0x42,
	0x3a, 0x6d, 0xc5, 0xf0, 0xa5, 0xab, 0x2b, 0xeb, 0x56, 0x95, 0x75, 0x8f, 0x4d, 0xe5, 0xfd, 0xf6,
	0x8d, 0xbf, 0xff, 0xb7, 0x65, 0xbf, 0xdd, 0x0b, 0xba, 0x25, 0xf0, 0x4c, 0xe3, 0xe0, 0x4f, 0xe0,
	0x65, 0x4c, 0x57, 0xb9, 0x90, 0x1c, 0xa3, 0x2c, 0xbc, 0x45, 0xd9, 0x78, 0x80, 0x32, 0x78, 0xb1,
	0x45, 0x4e, 0x6b, 0x94, 0x33, 0xf0, 0xa2, 0x60, 0xbb, 0x08, 0x9b, 0x0f, 0x11, 0x3e, 0xaf, 0x70,
	0x75, 0xba, 0x1f, 0x01, 0xd0, 0xe3, 0x19, 0xa6, 0x34, 0x71, 0xf6, 0x87, 0x8d, 0x51, 0x77, 0xfc,
	0x6e, 0x67, 0x67, 0xb6, 0x53, 0xbc, 0x1c, 0xbb, 0x1f, 0xd5, 0xe1, 0x94, 0x26, 0x41, 0x07, 0x55,
	0xbf, 0xf0, 0x0a, 0xf4, 0x63, 0xcc, 0x38, 0x8e, 0x50, 0xd9, 0xee, 0xe5, 0x7b, 0xa7, 0xa5, 0xae,
	0xf4, 0xfd, 0x27, 0x34, 0xfa, 0x78, 0xc3, 0x73, 0xfe, 0xde, 0xb7, 0x1d, 0x2b, 0xe8, 0xc5, 0x35,
	0x0b, 0xbc, 0x00, 0x5f, 0x64, 0x68, 0x1d, 0x46, 0x34, 0xcf, 0x71, 0x24, 0x43, 0x24, 0x25, 0xce,
	0x98, 0x14, 0xce, 0x81, 0x12, 0x3c, 0xba, 0x57, 0x83, 0x9f, 0xa7, 0xb9, 0xfc, 0x30, 0x3e, 0x47,
	0x69, 0x81, 0xd5, 0x0c, 0xbf, 0xb5, 0x47, 0x56, 0x00, 0x33, 0xb4, 0x9e, 0x68, 0x86, 0x8f, 0x86,
	0x60, 0xf0, 0x47, 0x03, 0xf4, 0xea, 0xda, 0xf0, 0x17, 0xd0, 0xe2, 0xb4, 0x90, 0x58, 0x38, 0x96,
	0xaa, 0xcd, 0xc9, 0x13, 0x93, 0x71, 0xcf, 0x26, 0xf3, 0xa0, 0x24, 0x54, 0x23, 0xf3, 0x97, 0x65,
	0xb7, 0xad, 0xc0, 0x28, 0x0c, 0xfe, 0xb4, 0x41, 0xbb, 0x72, 0xc3, 0x37, 0xdb, 0xad, 0xba, 0xb3,
	0x7b, 0x9b, 0xe5, 0x3a, 0x05, 0xcf, 0x63, 0x2c, 0x24, 0xc9, 0x55, 0x7f, 0x43, 0xc2, 0xc2, 0x94,
	0x08, 0xe9, 0xd8, 0xea, 0xa2, 0x47, 0x3b, 0xc6, 0x7d, 0x42, 0x62, 0x1e, 0xa0, 0x3c, 0xc1, 0xc1,
	0x61, 0x0d, 0x38, 0x65, 0xa7, 0x44, 0x48, 0xf8, 0x0e, 0xd4, 0x8d, 0x21, 0xa3, 0x5c, 0x0a, 0x35,
	0xa5, 0x9d, 0xe0, 0x59, 0xcd, 0x31, 0x2f, 0xed, 0xe5, 0x92, 0x09, 0x5a, 0xf0, 0x08, 0x6f, 0x54,
	0x9b, 0x8f, 0x50, 0xed, 0x69, 0x8c, 0x11, 0x7c, 0x03, 0xcc, 0xd9, 0x68, 0xed, 0x2b, 0xad, 0xae,
	0xb6, 0x29, 0x99, 0xc1, 0x3f, 0x16, 0xf8, 0xfc, 0xce, 0xce, 0xc3, 0x25, 0x68, 0x6f, 0x5e, 0x12,
	0xdd, 0x93, 0xf9, 0xd3, 0x5f, 0x12, 0xd7, 0x7c, 0xb5, 0xb9, 0xd6, 0x9b, 0x8d, 0xd6, 0x60, 0x06,
	0xfa, 0xb7, 0x82, 0xe0, 0x2b, 0xd0, 0xcc, 0x51, 0x86, 0xef, 0xb6, 0x47, 0x19, 0xe1, 0x6b, 0xd0,
	0xd2, 0x2f, 0x93, 0x7a, 0x13, 0xfb, 0xdb, 0xa9, 0x33, 0x66, 0xdf, 0x01, 0x87, 0x86, 0x3a, 0x14,
	0x0c, 0x47, 0x64, 0x41, 0x30, 0x87, 0x8d, 0x7f, 0x7d, 0xcb, 0x9f, 0x82, 0x6f, 0x09, 0xd5, 0x29,
	0xe9, 0x7b, 0x3f, 0x36, 0x3b, 0xbf, 0x5f, 0xa5, 0x37, 0x2f, 0x07, 0x7f, 0x6e, 0x5d, 0xb6, 0xd4,
	0x06, 0x7c, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xa8, 0x69, 0x2b, 0x08, 0x07, 0x00, 0x00,
}
