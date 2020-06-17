// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/dubbo_proxy/v4alpha/route.proto

package envoy_extensions_filters_network_dubbo_proxy_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/route/v4alpha"
	v4alpha1 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v4alpha"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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
	Interface            string   `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	Group                string   `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Routes               []*Route `protobuf:"bytes,5,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_236b06602011aa4c, []int{0}
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

func (m *RouteConfiguration) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *RouteConfiguration) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *RouteConfiguration) GetVersion() string {
	if m != nil {
		return m.Version
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
	return fileDescriptor_236b06602011aa4c, []int{1}
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
	Method               *MethodMatch             `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Headers              []*v4alpha.HeaderMatcher `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_236b06602011aa4c, []int{2}
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

func (m *RouteMatch) GetMethod() *MethodMatch {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *RouteMatch) GetHeaders() []*v4alpha.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_236b06602011aa4c, []int{3}
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
	WeightedClusters *v4alpha.WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

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

func (m *RouteAction) GetWeightedClusters() *v4alpha.WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type MethodMatch struct {
	Name                 *v4alpha1.StringMatcher                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ParamsMatch          map[uint32]*MethodMatch_ParameterMatchSpecifier `protobuf:"bytes,2,rep,name=params_match,json=paramsMatch,proto3" json:"params_match,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *MethodMatch) Reset()         { *m = MethodMatch{} }
func (m *MethodMatch) String() string { return proto.CompactTextString(m) }
func (*MethodMatch) ProtoMessage()    {}
func (*MethodMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_236b06602011aa4c, []int{4}
}

func (m *MethodMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch.Unmarshal(m, b)
}
func (m *MethodMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch.Marshal(b, m, deterministic)
}
func (m *MethodMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch.Merge(m, src)
}
func (m *MethodMatch) XXX_Size() int {
	return xxx_messageInfo_MethodMatch.Size(m)
}
func (m *MethodMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch proto.InternalMessageInfo

func (m *MethodMatch) GetName() *v4alpha1.StringMatcher {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MethodMatch) GetParamsMatch() map[uint32]*MethodMatch_ParameterMatchSpecifier {
	if m != nil {
		return m.ParamsMatch
	}
	return nil
}

type MethodMatch_ParameterMatchSpecifier struct {
	// Types that are valid to be assigned to ParameterMatchSpecifier:
	//	*MethodMatch_ParameterMatchSpecifier_ExactMatch
	//	*MethodMatch_ParameterMatchSpecifier_RangeMatch
	ParameterMatchSpecifier isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier `protobuf_oneof:"parameter_match_specifier"`
	XXX_NoUnkeyedLiteral    struct{}                                                      `json:"-"`
	XXX_unrecognized        []byte                                                        `json:"-"`
	XXX_sizecache           int32                                                         `json:"-"`
}

func (m *MethodMatch_ParameterMatchSpecifier) Reset()         { *m = MethodMatch_ParameterMatchSpecifier{} }
func (m *MethodMatch_ParameterMatchSpecifier) String() string { return proto.CompactTextString(m) }
func (*MethodMatch_ParameterMatchSpecifier) ProtoMessage()    {}
func (*MethodMatch_ParameterMatchSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_236b06602011aa4c, []int{4, 0}
}

func (m *MethodMatch_ParameterMatchSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Unmarshal(m, b)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Marshal(b, m, deterministic)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Merge(m, src)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_Size() int {
	return xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.Size(m)
}
func (m *MethodMatch_ParameterMatchSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MethodMatch_ParameterMatchSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_MethodMatch_ParameterMatchSpecifier proto.InternalMessageInfo

type isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier interface {
	isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier()
}

type MethodMatch_ParameterMatchSpecifier_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,3,opt,name=exact_match,json=exactMatch,proto3,oneof"`
}

type MethodMatch_ParameterMatchSpecifier_RangeMatch struct {
	RangeMatch *v3.Int64Range `protobuf:"bytes,4,opt,name=range_match,json=rangeMatch,proto3,oneof"`
}

func (*MethodMatch_ParameterMatchSpecifier_ExactMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (*MethodMatch_ParameterMatchSpecifier_RangeMatch) isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier() {
}

func (m *MethodMatch_ParameterMatchSpecifier) GetParameterMatchSpecifier() isMethodMatch_ParameterMatchSpecifier_ParameterMatchSpecifier {
	if m != nil {
		return m.ParameterMatchSpecifier
	}
	return nil
}

func (m *MethodMatch_ParameterMatchSpecifier) GetExactMatch() string {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *MethodMatch_ParameterMatchSpecifier) GetRangeMatch() *v3.Int64Range {
	if x, ok := m.GetParameterMatchSpecifier().(*MethodMatch_ParameterMatchSpecifier_RangeMatch); ok {
		return x.RangeMatch
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MethodMatch_ParameterMatchSpecifier) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MethodMatch_ParameterMatchSpecifier_ExactMatch)(nil),
		(*MethodMatch_ParameterMatchSpecifier_RangeMatch)(nil),
	}
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.extensions.filters.network.dubbo_proxy.v4alpha.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.extensions.filters.network.dubbo_proxy.v4alpha.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.extensions.filters.network.dubbo_proxy.v4alpha.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.extensions.filters.network.dubbo_proxy.v4alpha.RouteAction")
	proto.RegisterType((*MethodMatch)(nil), "envoy.extensions.filters.network.dubbo_proxy.v4alpha.MethodMatch")
	proto.RegisterMapType((map[uint32]*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.extensions.filters.network.dubbo_proxy.v4alpha.MethodMatch.ParamsMatchEntry")
	proto.RegisterType((*MethodMatch_ParameterMatchSpecifier)(nil), "envoy.extensions.filters.network.dubbo_proxy.v4alpha.MethodMatch.ParameterMatchSpecifier")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/dubbo_proxy/v4alpha/route.proto", fileDescriptor_236b06602011aa4c)
}

var fileDescriptor_236b06602011aa4c = []byte{
	// 770 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x4f, 0xdb, 0x4a,
	0x10, 0xc7, 0x09, 0x0e, 0x8f, 0xf1, 0x7b, 0x52, 0x58, 0x3d, 0x09, 0x93, 0xf7, 0x47, 0x90, 0xcb,
	0xe3, 0xe9, 0x49, 0xb6, 0x1e, 0xa1, 0x7f, 0x14, 0x40, 0x02, 0xd3, 0x4a, 0x70, 0x40, 0x42, 0x8e,
	0xaa, 0x0a, 0x2e, 0xe9, 0xe2, 0x6c, 0x12, 0x8b, 0xc4, 0x6b, 0xad, 0xd7, 0x81, 0xdc, 0x7a, 0xac,
	0xf8, 0x00, 0x95, 0xe8, 0x47, 0xe9, 0xbd, 0x52, 0xaf, 0xfd, 0x0c, 0xfd, 0x00, 0x3d, 0xf4, 0x52,
	0xf5, 0x54, 0x79, 0x76, 0x4d, 0x4c, 0x23, 0x55, 0x22, 0xf4, 0xb6, 0x9e, 0x9d, 0xf9, 0xcd, 0xef,
	0x37, 0xb3, 0x33, 0x86, 0x5d, 0x16, 0x8d, 0xf8, 0xd8, 0x65, 0x97, 0x92, 0x45, 0x49, 0xc8, 0xa3,
	0xc4, 0xed, 0x86, 0x03, 0xc9, 0x44, 0xe2, 0x46, 0x4c, 0x5e, 0x70, 0x71, 0xee, 0x76, 0xd2, 0xb3,
	0x33, 0xde, 0x8e, 0x05, 0xbf, 0x1c, 0xbb, 0xa3, 0x4d, 0x3a, 0x88, 0xfb, 0xd4, 0x15, 0x3c, 0x95,
	0xcc, 0x89, 0x05, 0x97, 0x9c, 0x6c, 0x22, 0x82, 0x33, 0x41, 0x70, 0x34, 0x82, 0xa3, 0x11, 0x9c,
	0x02, 0x82, 0xa3, 0x11, 0x6a, 0xff, 0xab, 0xbc, 0x01, 0x8f, 0xba, 0x61, 0x4f, 0xe1, 0xdd, 0x46,
	0x6f, 0x07, 0x7c, 0x18, 0xf3, 0x88, 0x45, 0x32, 0x51, 0x89, 0x6a, 0xff, 0xa8, 0x10, 0x39, 0x8e,
	0x99, 0x3b, 0xa4, 0x32, 0xe8, 0x33, 0x71, 0x13, 0x92, 0x48, 0x11, 0x46, 0x3d, 0xed, 0xb8, 0x52,
	0x70, 0x1c, 0x35, 0x5c, 0x41, 0xa3, 0x9e, 0x26, 0x5b, 0xfb, 0x2b, 0xed, 0xc4, 0xd4, 0xa5, 0x51,
	0xc4, 0x25, 0x95, 0x28, 0x37, 0x91, 0x54, 0xa6, 0x79, 0x8a, 0xb5, 0xa9, 0xeb, 0x11, 0x13, 0x99,
	0xa8, 0x09, 0xf8, 0xf2, 0x88, 0x0e, 0xc2, 0x0e, 0xcd, 0xe8, 0xea, 0x83, 0xba, 0xa8, 0xbf, 0x2e,
	0x01, 0xf1, 0x33, 0xe6, 0xfb, 0xa8, 0x29, 0x15, 0x88, 0x40, 0x08, 0xcc, 0x47, 0x74, 0xc8, 0x6c,
	0x63, 0xd5, 0x58, 0x5f, 0xf4, 0xf1, 0x4c, 0xfe, 0x84, 0xc5, 0x30, 0x92, 0x4c, 0x74, 0x69, 0xc0,
	0xec, 0x12, 0x5e, 0x4c, 0x0c, 0xe4, 0x77, 0x30, 0x7b, 0x82, 0xa7, 0xb1, 0x5d, 0xc6, 0x1b, 0xf5,
	0x41, 0x6c, 0x58, 0xd0, 0x5c, 0xec, 0x79, 0xb4, 0xe7, 0x9f, 0xa4, 0x05, 0x15, 0xac, 0x58, 0x62,
	0x9b, 0xab, 0xe5, 0x75, 0x6b, 0x63, 0xcb, 0x99, 0xa5, 0x23, 0x0e, 0x72, 0xf7, 0x35, 0x54, 0xf3,
	0xf0, 0xcd, 0xbb, 0x57, 0x7f, 0x3f, 0x01, 0xef, 0x6e, 0x50, 0x0d, 0x67, 0xba, 0x02, 0xf5, 0xab,
	0x12, 0x98, 0x68, 0x26, 0x2f, 0xc0, 0xc4, 0xc6, 0x61, 0x31, 0xac, 0x8d, 0xdd, 0x7b, 0x10, 0x3d,
	0xca, 0x70, 0xbc, 0x5f, 0xbe, 0x7a, 0xe6, 0x95, 0x51, 0xaa, 0x1a, 0xbe, 0x02, 0x26, 0x14, 0x4c,
	0x14, 0x80, 0x55, 0xb5, 0x36, 0xf6, 0xee, 0x91, 0x61, 0x2f, 0xc8, 0xd8, 0x17, 0x53, 0x20, 0x72,
	0x73, 0x3b, 0xab, 0xcc, 0x23, 0x78, 0x30, 0x53, 0x65, 0xea, 0x9f, 0x0c, 0x80, 0x89, 0x00, 0x72,
	0x02, 0x95, 0x21, 0x93, 0x7d, 0xde, 0xd1, 0x25, 0x99, 0x91, 0xf0, 0x11, 0x62, 0x20, 0xa4, 0xaf,
	0x01, 0xc9, 0x3e, 0x2c, 0xf4, 0x19, 0xed, 0x30, 0x91, 0xd8, 0x25, 0x7c, 0x17, 0xff, 0x6a, 0x6c,
	0x35, 0x73, 0x8e, 0x9a, 0xe1, 0x1c, 0xe1, 0x00, 0x5d, 0x8f, 0xd4, 0x58, 0xf9, 0x79, 0x64, 0x73,
	0x2f, 0x13, 0xbb, 0x0d, 0xcd, 0x99, 0xc4, 0x22, 0x5a, 0xfd, 0xa3, 0x01, 0x56, 0xa1, 0xa0, 0xa4,
	0x06, 0x0b, 0xc1, 0x20, 0x4d, 0x24, 0x13, 0x6a, 0x26, 0x0e, 0xe6, 0xfc, 0xdc, 0x40, 0x4e, 0x61,
	0xe9, 0x82, 0x85, 0xbd, 0xbe, 0x64, 0x9d, 0xb6, 0xb6, 0x25, 0xba, 0x95, 0xff, 0xfd, 0x88, 0xfd,
	0x73, 0x1d, 0xb4, 0xaf, 0x62, 0x0e, 0xe6, 0xfc, 0xea, 0xc5, 0x6d, 0x53, 0xd2, 0xf4, 0x32, 0x29,
	0x3b, 0xb0, 0x35, 0x93, 0x14, 0xfd, 0x18, 0x6c, 0x58, 0xd2, 0xb4, 0xda, 0x49, 0xcc, 0x82, 0xb0,
	0x1b, 0x32, 0x41, 0xca, 0x5f, 0x3c, 0xa3, 0x7e, 0x6d, 0x82, 0x55, 0xe8, 0x02, 0xd9, 0x29, 0x8c,
	0xfd, 0xa4, 0xf4, 0xd9, 0x4a, 0x72, 0xf4, 0xee, 0xba, 0x21, 0xdf, 0xc2, 0xdd, 0x95, 0x97, 0x5e,
	0x6d, 0x88, 0x14, 0x7e, 0x8d, 0xa9, 0xa0, 0xc3, 0xa4, 0xad, 0x06, 0x46, 0x75, 0xd0, 0xbf, 0xf7,
	0xeb, 0x70, 0x8e, 0x11, 0x15, 0xcf, 0x4f, 0x23, 0x29, 0xc6, 0xbe, 0x15, 0x4f, 0x2c, 0xb5, 0xcf,
	0x06, 0x2c, 0xa3, 0x07, 0x93, 0xfa, 0x31, 0xb4, 0x6e, 0x64, 0xae, 0x81, 0xc5, 0x2e, 0x69, 0x20,
	0x35, 0xa3, 0xb2, 0xee, 0x1d, 0xa0, 0x51, 0x89, 0xde, 0x06, 0x0b, 0x97, 0xad, 0x76, 0x99, 0x47,
	0xed, 0x2b, 0x45, 0xed, 0xa3, 0x86, 0x73, 0x18, 0xc9, 0x87, 0x9b, 0x7e, 0xe6, 0x96, 0x45, 0xa3,
	0x3f, 0x46, 0x37, 0x4f, 0xb3, 0x06, 0x3d, 0x83, 0xd6, 0x5d, 0x1b, 0x34, 0x25, 0x6f, 0x9a, 0xbc,
	0xf7, 0x07, 0xac, 0xc4, 0xf9, 0x95, 0x62, 0x37, 0x69, 0x60, 0xed, 0xda, 0x80, 0xea, 0xf7, 0x75,
	0x21, 0x55, 0x28, 0x9f, 0xb3, 0x31, 0xf6, 0xef, 0x37, 0x3f, 0x3b, 0x12, 0x0e, 0xe6, 0x88, 0x0e,
	0xd2, 0x7c, 0xb7, 0x9c, 0xfc, 0xa4, 0x66, 0x4c, 0xb3, 0xf5, 0x55, 0x9e, 0x66, 0xe9, 0xb1, 0x31,
	0xf3, 0xab, 0x2d, 0xa4, 0xf1, 0x4e, 0xdf, 0xbe, 0x7c, 0xff, 0xa1, 0x52, 0xaa, 0x96, 0xc1, 0x0b,
	0xb9, 0x62, 0xad, 0xfc, 0x66, 0x11, 0xe0, 0xa9, 0xf5, 0x75, 0x9c, 0xfd, 0xf3, 0x8e, 0x8d, 0xb3,
	0x0a, 0xfe, 0xfc, 0x1a, 0xdf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xc6, 0x85, 0x31, 0x48, 0x08,
	0x00, 0x00,
}
