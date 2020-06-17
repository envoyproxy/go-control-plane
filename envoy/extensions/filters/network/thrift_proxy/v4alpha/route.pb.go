// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v4alpha/route.proto

package envoy_extensions_filters_network_thrift_proxy_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha1 "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/route/v4alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type RouteConfiguration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3d7ab794fbf7d, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Match                *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3d7ab794fbf7d, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier       isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	Invert               bool                        `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	Headers              []*v4alpha.HeaderMatcher    `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3d7ab794fbf7d, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*v4alpha.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	//	*RouteAction_ClusterHeader
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch        *v4alpha1.Metadata             `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	RateLimits           []*v4alpha.RateLimit           `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	StripServiceName     bool                           `protobuf:"varint,5,opt,name=strip_service_name,json=stripServiceName,proto3" json:"strip_service_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3d7ab794fbf7d, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

type RouteAction_ClusterHeader struct {
	ClusterHeader string `protobuf:"bytes,6,opt,name=cluster_header,json=clusterHeader,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_ClusterHeader) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetClusterHeader() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_ClusterHeader); ok {
		return x.ClusterHeader
	}
	return ""
}

func (m *RouteAction) GetMetadataMatch() *v4alpha1.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*v4alpha.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

func (m *RouteAction) GetStripServiceName() bool {
	if m != nil {
		return m.StripServiceName
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
		(*RouteAction_ClusterHeader)(nil),
	}
}

type WeightedCluster struct {
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3d7ab794fbf7d, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	MetadataMatch        *v4alpha1.Metadata    `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3d7ab794fbf7d, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *v4alpha1.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.extensions.filters.network.thrift_proxy.v4alpha.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/thrift_proxy/v4alpha/route.proto", fileDescriptor_3ac3d7ab794fbf7d)
}

var fileDescriptor_3ac3d7ab794fbf7d = []byte{
	// 803 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x4f, 0x1b, 0x47,
	0x14, 0x66, 0xd7, 0xd8, 0xb8, 0xcf, 0x05, 0xcc, 0x1c, 0xa8, 0x45, 0x5b, 0x64, 0xa0, 0x95, 0xdc,
	0xaa, 0xda, 0x6d, 0xa1, 0xad, 0x2a, 0x17, 0xaa, 0xb2, 0x56, 0x88, 0x83, 0x02, 0x42, 0x1b, 0x48,
	0x0e, 0x39, 0xac, 0xc6, 0xeb, 0xb1, 0xbd, 0x8a, 0xbd, 0xb3, 0x9a, 0x9d, 0xb5, 0x21, 0xa7, 0x1c,
	0xa3, 0x48, 0xc9, 0x21, 0xc7, 0xfc, 0x29, 0xb9, 0x47, 0x8a, 0x72, 0xcb, 0x5f, 0x91, 0x7f, 0x21,
	0x42, 0x39, 0x44, 0xf3, 0x63, 0xfd, 0x03, 0x24, 0x24, 0x0c, 0x27, 0x8f, 0x67, 0xde, 0xfb, 0xde,
	0xf7, 0xbe, 0xef, 0xcd, 0x2c, 0xec, 0x92, 0xb0, 0x4f, 0xcf, 0x6c, 0x72, 0xca, 0x49, 0x18, 0x07,
	0x34, 0x8c, 0xed, 0x56, 0xd0, 0xe5, 0x84, 0xc5, 0x76, 0x48, 0xf8, 0x80, 0xb2, 0x27, 0x36, 0xef,
	0xb0, 0xa0, 0xc5, 0xbd, 0x88, 0xd1, 0xd3, 0x33, 0xbb, 0xff, 0x27, 0xee, 0x46, 0x1d, 0x6c, 0x33,
	0x9a, 0x70, 0x62, 0x45, 0x8c, 0x72, 0x8a, 0xfe, 0x92, 0x10, 0xd6, 0x08, 0xc2, 0xd2, 0x10, 0x96,
	0x86, 0xb0, 0xc6, 0x21, 0x2c, 0x0d, 0xb1, 0xf2, 0x93, 0xaa, 0xec, 0xd3, 0xb0, 0x15, 0xb4, 0x6d,
	0x9f, 0x32, 0x32, 0x44, 0x6f, 0xe0, 0x58, 0x83, 0xaf, 0xfc, 0x31, 0x11, 0x25, 0xcb, 0x4e, 0x92,
	0xf0, 0x7c, 0xda, 0x8b, 0x68, 0x48, 0x42, 0x1e, 0xeb, 0x94, 0xd5, 0x36, 0xa5, 0xed, 0x2e, 0xb1,
	0xe5, 0xbf, 0x46, 0xd2, 0xb2, 0x07, 0x0c, 0x47, 0x91, 0xe0, 0xa3, 0xce, 0x7f, 0x4c, 0x9a, 0x11,
	0xb6, 0x71, 0x18, 0x52, 0x8e, 0xb9, 0x6c, 0x39, 0xe6, 0x98, 0x27, 0xe9, 0xf1, 0xda, 0xa5, 0xe3,
	0x3e, 0x61, 0xa2, 0xaf, 0x20, 0x6c, 0xeb, 0x90, 0xef, 0xfa, 0xb8, 0x1b, 0x34, 0xb1, 0xa0, 0xa2,
	0x17, 0xea, 0x60, 0xfd, 0x83, 0x01, 0xc8, 0x15, 0xac, 0x6a, 0x92, 0x6f, 0xc2, 0x24, 0x02, 0x42,
	0x30, 0x1b, 0xe2, 0x1e, 0x29, 0x19, 0x65, 0xa3, 0xf2, 0x8d, 0x2b, 0xd7, 0xe8, 0x18, 0x72, 0x92,
	0x7f, 0x5c, 0x32, 0xcb, 0x99, 0x4a, 0x61, 0x73, 0xdb, 0x9a, 0x4a, 0x46, 0x4b, 0x96, 0x73, 0x35,
	0x56, 0x75, 0xff, 0xcd, 0xbb, 0xe7, 0xab, 0x77, 0xa0, 0x76, 0x4d, 0xac, 0x2d, 0xeb, 0x32, 0xeb,
	0xf5, 0x97, 0x26, 0x64, 0xe5, 0x36, 0xc2, 0x90, 0xed, 0x61, 0xee, 0x77, 0x64, 0x03, 0x85, 0xcd,
	0xdd, 0x9b, 0x50, 0x3d, 0x10, 0x40, 0x4e, 0xfe, 0xdc, 0xc9, 0xbe, 0x30, 0xcc, 0xa2, 0xe1, 0x2a,
	0x64, 0xd4, 0x80, 0xac, 0x6c, 0xa1, 0x64, 0xca, 0x12, 0xce, 0x4d, 0x4a, 0xec, 0xfa, 0x82, 0xff,
	0x78, 0x0d, 0x09, 0x5d, 0xdd, 0x11, 0xe2, 0xfc, 0x03, 0x7f, 0x4f, 0x27, 0x8e, 0xd0, 0x03, 0x46,
	0x2d, 0xa0, 0x35, 0x28, 0xf4, 0x08, 0xef, 0xd0, 0xa6, 0x37, 0xf2, 0xb6, 0x3e, 0xe3, 0x82, 0xda,
	0x3c, 0x14, 0x1e, 0x6f, 0xc0, 0xb7, 0x31, 0x61, 0xfd, 0xc0, 0x27, 0x2a, 0xc6, 0xd4, 0x31, 0x05,
	0xbd, 0x2b, 0x83, 0x96, 0x21, 0x17, 0x84, 0x7d, 0xc2, 0x78, 0x29, 0x53, 0x36, 0x2a, 0x79, 0x57,
	0xff, 0x43, 0x35, 0x98, 0xeb, 0x10, 0xdc, 0x24, 0x2c, 0x2e, 0xcd, 0xca, 0x09, 0xf9, 0x45, 0x6b,
	0xa2, 0xee, 0x82, 0xa5, 0xae, 0x60, 0xda, 0x78, 0x5d, 0x86, 0x4a, 0x66, 0x84, 0xb9, 0x69, 0x66,
	0xd5, 0x11, 0x2d, 0xef, 0xc0, 0xbf, 0xd3, 0xb5, 0xac, 0xbc, 0x5a, 0x86, 0x45, 0xe9, 0x91, 0x17,
	0x47, 0xc4, 0x0f, 0x5a, 0x01, 0x61, 0x28, 0xf3, 0xd9, 0x31, 0xd6, 0xbf, 0x64, 0xa0, 0x30, 0xa6,
	0x37, 0xda, 0x80, 0x39, 0xbf, 0x9b, 0xc4, 0x9c, 0x30, 0x25, 0x86, 0x33, 0x77, 0xee, 0xcc, 0x32,
	0xb3, 0x6c, 0xd4, 0x67, 0xdc, 0xf4, 0x04, 0x25, 0xb0, 0x34, 0x20, 0x41, 0xbb, 0xc3, 0x49, 0xd3,
	0xd3, 0x7b, 0xb1, 0xf6, 0x7c, 0x6f, 0x4a, 0xcf, 0x1f, 0x69, 0xbc, 0x9a, 0x82, 0xab, 0xcf, 0xb8,
	0xc5, 0xc1, 0xe4, 0x56, 0x8c, 0x7e, 0x87, 0x05, 0x5d, 0xcd, 0x53, 0xd2, 0x94, 0x72, 0x17, 0x29,
	0xce, 0xeb, 0x00, 0xa5, 0x24, 0xda, 0x87, 0x85, 0x1e, 0xe1, 0xb8, 0x89, 0x39, 0xf6, 0xd4, 0xf0,
	0x67, 0x24, 0xcb, 0x8d, 0x49, 0x17, 0xc4, 0xbb, 0x35, 0x64, 0x72, 0xa0, 0x13, 0xdc, 0xf9, 0x34,
	0x55, 0x8d, 0xca, 0x1e, 0x14, 0x18, 0xe6, 0xc4, 0xeb, 0x06, 0xbd, 0x80, 0xa7, 0x76, 0xfe, 0x7c,
	0x95, 0x9d, 0x2e, 0xe6, 0xe4, 0xbe, 0x88, 0x76, 0x81, 0xa5, 0xcb, 0x18, 0xfd, 0x06, 0x28, 0xe6,
	0x2c, 0x88, 0xbc, 0x89, 0xa9, 0xca, 0xca, 0xb1, 0x29, 0xca, 0x93, 0x07, 0xa3, 0xc1, 0xaa, 0xd6,
	0x84, 0xf7, 0xff, 0xc1, 0xf6, 0x74, 0xde, 0xeb, 0x4b, 0x54, 0x82, 0xa5, 0x54, 0xb8, 0x0b, 0xf6,
	0x7f, 0xca, 0xc0, 0xe2, 0x05, 0xe9, 0xd1, 0x53, 0xc8, 0x0f, 0x4d, 0x35, 0x64, 0x97, 0xc7, 0xb7,
	0x63, 0xaa, 0xa5, 0x7f, 0xd5, 0xb6, 0xbc, 0xda, 0xaf, 0x0d, 0x33, 0x6f, 0xb8, 0xc3, 0x7a, 0x2b,
	0xaf, 0x4c, 0x98, 0x9f, 0x88, 0x42, 0xdf, 0x8f, 0x3f, 0xbb, 0x43, 0xab, 0xf5, 0xfb, 0xbb, 0x03,
	0x39, 0x35, 0x25, 0x7a, 0xfa, 0x7e, 0xb0, 0xd4, 0x67, 0xc3, 0x4a, 0x3f, 0x1b, 0xd6, 0xc9, 0xbd,
	0x90, 0x6f, 0x6d, 0x3e, 0xc4, 0xdd, 0x84, 0xc8, 0xe4, 0x5f, 0xcd, 0x8a, 0xe1, 0xea, 0xa4, 0xdb,
	0x1c, 0x8f, 0xea, 0x89, 0x30, 0xea, 0x08, 0x0e, 0xaf, 0x6d, 0xd4, 0x95, 0x22, 0x55, 0xef, 0x0a,
	0x58, 0x07, 0xfe, 0xbf, 0x29, 0xac, 0xf3, 0xf8, 0xed, 0xb3, 0xf7, 0x1f, 0x73, 0x66, 0x31, 0x03,
	0xb5, 0x80, 0xaa, 0xfe, 0x54, 0xec, 0x54, 0xd6, 0x3a, 0xea, 0x11, 0x3d, 0x12, 0x32, 0x1f, 0x19,
	0x8d, 0x9c, 0xd4, 0x7b, 0xeb, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xe1, 0x38, 0x39, 0x89,
	0x08, 0x00, 0x00,
}
