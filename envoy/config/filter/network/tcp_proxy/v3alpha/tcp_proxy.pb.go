// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_proxy/v3alpha/tcp_proxy.proto

package envoy_config_filter_network_tcp_proxy_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/v2/envoy/api/v3alpha/core"
	v3alpha "github.com/envoyproxy/go-control-plane/v2/envoy/config/filter/accesslog/v3alpha"
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
	AccessLog             []*v3alpha.AccessLog        `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
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
	return fileDescriptor_0cada114fb4da110, []int{0}
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

func (m *TcpProxy) GetAccessLog() []*v3alpha.AccessLog {
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
	return fileDescriptor_0cada114fb4da110, []int{0, 0}
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
	return fileDescriptor_0cada114fb4da110, []int{0, 0, 0}
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
	return fileDescriptor_0cada114fb4da110, []int{0, 1}
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
	return fileDescriptor_0cada114fb4da110, []int{0, 1, 0}
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
	proto.RegisterType((*TcpProxy)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.DeprecatedV1.TCPRoute")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_proxy/v3alpha/tcp_proxy.proto", fileDescriptor_0cada114fb4da110)
}

var fileDescriptor_0cada114fb4da110 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x18, 0xad, 0x9d, 0x34, 0x4d, 0x26, 0x09, 0x6c, 0x67, 0x59, 0xad, 0x09, 0x08, 0x12, 0xc4, 0x45,
	0xb4, 0x08, 0x9b, 0x6d, 0xc5, 0x15, 0xe2, 0x62, 0x9d, 0x8a, 0xdd, 0x4a, 0x1b, 0xc8, 0x5a, 0xdd,
	0xf6, 0xd2, 0x9a, 0xda, 0x13, 0x77, 0x84, 0xed, 0x19, 0xcd, 0x8c, 0x93, 0x54, 0xbc, 0x05, 0x77,
	0x88, 0x37, 0x40, 0xbc, 0x12, 0xef, 0x81, 0x7a, 0x85, 0x3c, 0x33, 0x76, 0xdc, 0xd2, 0xaa, 0x14,
	0xf5, 0x2a, 0xf1, 0xf9, 0xbe, 0x73, 0xce, 0xcc, 0xf7, 0x33, 0xe0, 0x7b, 0x9c, 0xaf, 0xe8, 0xa5,
	0x17, 0xd1, 0x7c, 0x49, 0x12, 0x6f, 0x49, 0x52, 0x89, 0xb9, 0x97, 0x63, 0xb9, 0xa6, 0xfc, 0x67,
	0x4f, 0x46, 0x2c, 0x64, 0x9c, 0x6e, 0x2e, 0xbd, 0xd5, 0x21, 0x4a, 0xd9, 0x05, 0xda, 0x22, 0x2e,
	0xe3, 0x54, 0x52, 0xf8, 0xb5, 0xa2, 0xbb, 0x9a, 0xee, 0x6a, 0xba, 0x6b, 0xe8, 0xee, 0x36, 0xd9,
	0xd0, 0x47, 0xdf, 0xde, 0xe6, 0x86, 0xa2, 0x08, 0x0b, 0x91, 0xd2, 0xa4, 0x76, 0xa9, 0x11, 0xed,
	0x32, 0xfa, 0x52, 0xd3, 0x10, 0x23, 0x75, 0x4a, 0x44, 0x39, 0xf6, 0x50, 0x1c, 0x73, 0x2c, 0x84,
	0xc9, 0x9a, 0xdc, 0x91, 0x75, 0x8e, 0x04, 0x36, 0x29, 0x9f, 0x25, 0x94, 0x26, 0x29, 0xf6, 0xd4,
	0xd7, 0x79, 0xb1, 0xf4, 0xe2, 0x82, 0x23, 0x49, 0x68, 0x7e, 0x57, 0x7c, 0xcd, 0x11, 0x63, 0x98,
	0x57, 0x16, 0xcf, 0x57, 0x28, 0x25, 0x31, 0x92, 0xd8, 0xab, 0xfe, 0xe8, 0xc0, 0x17, 0x7f, 0xf6,
	0x41, 0xf7, 0x24, 0x62, 0x8b, 0xf2, 0xb6, 0x70, 0x0a, 0xfa, 0x42, 0x22, 0x19, 0x32, 0x8e, 0x97,
	0x64, 0xe3, 0x58, 0x63, 0x6b, 0xda, 0xf3, 0xf7, 0xae, 0xfc, 0x36, 0xb7, 0xc7, 0x56, 0x00, 0xca,
	0xd8, 0x42, 0x85, 0xe0, 0x08, 0xec, 0x45, 0x69, 0x21, 0x24, 0xe6, 0x8e, 0x5d, 0x66, 0xbd, 0xd9,
	0x09, 0x2a, 0x00, 0xae, 0xc0, 0xfe, 0x1a, 0x93, 0xe4, 0x42, 0xe2, 0x38, 0x34, 0x98, 0x70, 0xc0,
	0xd8, 0x9a, 0xf6, 0x0f, 0x5e, 0xbb, 0x0f, 0x2a, 0xbb, 0x5b, 0x9d, 0xcc, 0x3d, 0x33, 0x82, 0x33,
	0xad, 0xf7, 0x66, 0x27, 0x78, 0xb2, 0xbe, 0x0e, 0x09, 0xf8, 0x1a, 0x7c, 0x90, 0x61, 0x89, 0x62,
	0x24, 0x51, 0x98, 0x21, 0x19, 0x5d, 0x38, 0x3d, 0x65, 0x3a, 0x36, 0xa6, 0x88, 0x91, 0x5a, 0xb8,
	0xac, 0xaf, 0x3b, 0x37, 0xd9, 0xc1, 0xb0, 0xe2, 0xcd, 0x4b, 0x1a, 0xfc, 0x01, 0x0c, 0x48, 0x9c,
	0xe2, 0x50, 0x92, 0x0c, 0xd3, 0x42, 0x3a, 0x5d, 0x25, 0xf3, 0xb1, 0xab, 0x6b, 0xec, 0x56, 0x35,
	0x76, 0x8f, 0x4c, 0x0f, 0xfc, 0xee, 0x95, 0xbf, 0xfb, 0x87, 0x65, 0xbf, 0xd8, 0x09, 0xfa, 0x25,
	0xf1, 0x44, 0xf3, 0xe0, 0x3b, 0xf0, 0x3c, 0xa6, 0xeb, 0x5c, 0x48, 0x8e, 0x51, 0x16, 0x5e, 0x93,
	0x6c, 0xdd, 0x23, 0x19, 0x3c, 0xdb, 0x32, 0x8f, 0x1b, 0x92, 0x73, 0xf0, 0xac, 0x60, 0xb7, 0x09,
	0xb6, 0xef, 0x13, 0x7c, 0x5a, 0xf1, 0x9a, 0x72, 0x3f, 0x01, 0xa0, 0x47, 0x36, 0x4c, 0x69, 0xe2,
	0xec, 0x8e, 0x5b, 0xd3, 0xfe, 0xc1, 0x37, 0xb7, 0xf6, 0x68, 0x3b, 0xd9, 0x55, 0x09, 0x5f, 0x29,
	0xe4, 0x2d, 0x4d, 0x82, 0x1e, 0xaa, 0xfe, 0xc2, 0x0c, 0x0c, 0x63, 0xcc, 0x38, 0x8e, 0x50, 0xd9,
	0xfd, 0xd5, 0x4b, 0xa7, 0xa3, 0xce, 0x75, 0xf4, 0x7f, 0xfb, 0x7e, 0x54, 0x8b, 0x9d, 0xbe, 0xf4,
	0x6d, 0xc7, 0x0a, 0x06, 0x71, 0x03, 0x81, 0x67, 0xe0, 0xa3, 0x0c, 0x6d, 0xc2, 0x88, 0xe6, 0x39,
	0x8e, 0x64, 0x88, 0xa4, 0xc4, 0x19, 0x93, 0xc2, 0xd9, 0x53, 0xae, 0x9f, 0xfe, 0xab, 0x1a, 0xef,
	0x8f, 0x73, 0x79, 0x78, 0x70, 0x8a, 0xd2, 0x02, 0xab, 0xb9, 0x7e, 0x61, 0x4f, 0xad, 0x00, 0x66,
	0x68, 0x33, 0xd3, 0x0a, 0xaf, 0x8c, 0xc0, 0xe8, 0xf7, 0x16, 0x18, 0x34, 0xbd, 0x21, 0x05, 0x1d,
	0x4e, 0x0b, 0x89, 0x85, 0x63, 0xa9, 0x2a, 0xcd, 0x1f, 0xe3, 0x46, 0xee, 0xc9, 0x6c, 0x11, 0x94,
	0xaa, 0x6a, 0x82, 0x7e, 0xb5, 0xec, 0xae, 0x15, 0x18, 0x9b, 0xd1, 0x6f, 0x36, 0xe8, 0x56, 0x61,
	0x38, 0xd9, 0xae, 0xdb, 0x8d, 0xa5, 0xac, 0xb7, 0xee, 0x1d, 0x78, 0x1a, 0x63, 0x21, 0x49, 0xae,
	0xda, 0x1d, 0x12, 0x16, 0xa6, 0x44, 0x48, 0xc7, 0x56, 0xa7, 0x9d, 0xdc, 0xb5, 0x02, 0x33, 0x12,
	0xf3, 0x00, 0xe5, 0x09, 0x0e, 0xf6, 0x1b, 0xec, 0x63, 0xf6, 0x96, 0x08, 0x09, 0xbf, 0x02, 0x4d,
	0x30, 0x64, 0x94, 0x4b, 0xa1, 0x26, 0xb7, 0x17, 0x3c, 0x69, 0x04, 0x16, 0x25, 0x5e, 0x6e, 0x9f,
	0xa0, 0x05, 0x8f, 0x70, 0x6d, 0xdd, 0xfe, 0xaf, 0xd6, 0x03, 0x4d, 0x34, 0xae, 0x13, 0x60, 0xbe,
	0x8d, 0xe1, 0xae, 0x32, 0xec, 0x6b, 0x4c, 0x79, 0x8d, 0xfe, 0xb2, 0xc0, 0x87, 0x37, 0x5e, 0x04,
	0xf8, 0x0b, 0xe8, 0xd6, 0x8f, 0x8d, 0x6e, 0xd1, 0xfb, 0x47, 0x7a, 0x6c, 0x5c, 0xf3, 0xab, 0xe1,
	0x46, 0xab, 0x6a, 0xc3, 0xd1, 0x1c, 0x0c, 0xaf, 0x25, 0xc1, 0x4f, 0x40, 0x3b, 0x47, 0x19, 0xbe,
	0xd9, 0x2d, 0x05, 0xc2, 0xcf, 0x41, 0x47, 0x3f, 0x5e, 0xea, 0xed, 0x1c, 0x6e, 0x27, 0xd1, 0xc0,
	0xbe, 0x03, 0xf6, 0x8d, 0x74, 0x28, 0x18, 0x8e, 0xc8, 0x92, 0x60, 0x0e, 0x5b, 0x7f, 0xfb, 0x96,
	0xff, 0x23, 0xf8, 0x8e, 0x50, 0x7d, 0x2f, 0x7d, 0xf8, 0x07, 0x5d, 0xd1, 0x1f, 0x56, 0x77, 0x5c,
	0x94, 0x1b, 0xb1, 0xb0, 0xce, 0x3b, 0x6a, 0x35, 0x0e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xb9,
	0x94, 0x91, 0x4b, 0x4e, 0x07, 0x00, 0x00,
}
