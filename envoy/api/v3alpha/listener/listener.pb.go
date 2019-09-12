// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/listener.proto

package envoy_api_v3alpha_listener

import (
	fmt "fmt"
	auth "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/auth"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type FilterChainMatch_ConnectionSourceType int32

const (
	FilterChainMatch_ANY      FilterChainMatch_ConnectionSourceType = 0
	FilterChainMatch_LOCAL    FilterChainMatch_ConnectionSourceType = 1
	FilterChainMatch_EXTERNAL FilterChainMatch_ConnectionSourceType = 2
)

var FilterChainMatch_ConnectionSourceType_name = map[int32]string{
	0: "ANY",
	1: "LOCAL",
	2: "EXTERNAL",
}

var FilterChainMatch_ConnectionSourceType_value = map[string]int32{
	"ANY":      0,
	"LOCAL":    1,
	"EXTERNAL": 2,
}

func (x FilterChainMatch_ConnectionSourceType) String() string {
	return proto.EnumName(FilterChainMatch_ConnectionSourceType_name, int32(x))
}

func (FilterChainMatch_ConnectionSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{1, 0}
}

type Filter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_Config
	//	*Filter_TypedConfig
	ConfigType           isFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Filter_Config) isFilter_ConfigType() {}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Filter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Filter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_Config)(nil),
		(*Filter_TypedConfig)(nil),
	}
}

type FilterChainMatch struct {
	DestinationPort      *wrappers.UInt32Value                 `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	PrefixRanges         []*core.CidrRange                     `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	AddressSuffix        string                                `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	SuffixLen            *wrappers.UInt32Value                 `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	SourceType           FilterChainMatch_ConnectionSourceType `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=envoy.api.v3alpha.listener.FilterChainMatch_ConnectionSourceType" json:"source_type,omitempty"`
	SourcePrefixRanges   []*core.CidrRange                     `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	SourcePorts          []uint32                              `protobuf:"varint,7,rep,packed,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	ServerNames          []string                              `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	TransportProtocol    string                                `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	ApplicationProtocols []string                              `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourceType() FilterChainMatch_ConnectionSourceType {
	if m != nil {
		return m.SourceType
	}
	return FilterChainMatch_ANY
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []uint32 {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

type FilterChain struct {
	FilterChainMatch     *FilterChainMatch          `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	TlsContext           *auth.DownstreamTlsContext `protobuf:"bytes,2,opt,name=tls_context,json=tlsContext,proto3" json:"tls_context,omitempty"`
	Filters              []*Filter                  `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	UseProxyProto        *wrappers.BoolValue        `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	Metadata             *core.Metadata             `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	TransportSocket      *core.TransportSocket      `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	Name                 string                     `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetTlsContext() *auth.DownstreamTlsContext {
	if m != nil {
		return m.TlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func (m *FilterChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListenerFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_Config
	//	*ListenerFilter_TypedConfig
	ConfigType           isListenerFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{3}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ListenerFilter_Config) isListenerFilter_ConfigType() {}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ListenerFilter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ListenerFilter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilter_Config)(nil),
		(*ListenerFilter_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v3alpha.listener.FilterChainMatch_ConnectionSourceType", FilterChainMatch_ConnectionSourceType_name, FilterChainMatch_ConnectionSourceType_value)
	proto.RegisterType((*Filter)(nil), "envoy.api.v3alpha.listener.Filter")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v3alpha.listener.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v3alpha.listener.FilterChain")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.api.v3alpha.listener.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/listener/listener.proto", fileDescriptor_907f943165b7183f)
}

var fileDescriptor_907f943165b7183f = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xcd, 0x6e, 0x23, 0x45,
	0x10, 0xce, 0xd8, 0x8e, 0x7f, 0x6a, 0xec, 0x64, 0x68, 0x05, 0x65, 0x30, 0x2b, 0xe4, 0x58, 0x20,
	0x0c, 0x82, 0xb1, 0xb0, 0x2f, 0x60, 0xf6, 0xe2, 0x31, 0x59, 0x2d, 0x2b, 0x27, 0x58, 0x63, 0x13,
	0x01, 0x97, 0x51, 0x67, 0xdc, 0x4e, 0x5a, 0x8c, 0xbb, 0x47, 0xdd, 0xed, 0x6c, 0x7c, 0xe5, 0x11,
	0x78, 0x0a, 0xb4, 0x4f, 0xb5, 0x67, 0x1e, 0x21, 0x97, 0x45, 0xd3, 0xd3, 0xe3, 0x64, 0xe3, 0x0d,
	0x01, 0x0e, 0xdc, 0xba, 0xab, 0xbe, 0xfa, 0xaa, 0xba, 0xbe, 0xaa, 0x86, 0xcf, 0x08, 0xbb, 0xe2,
	0xeb, 0x2e, 0x4e, 0x68, 0xf7, 0xaa, 0x8f, 0xe3, 0xe4, 0x12, 0x77, 0x63, 0x2a, 0x15, 0x61, 0x44,
	0x6c, 0x0e, 0x5e, 0x22, 0xb8, 0xe2, 0xa8, 0xa9, 0xa1, 0x1e, 0x4e, 0xa8, 0x67, 0xa0, 0x5e, 0x8e,
	0x68, 0x7e, 0xbc, 0x4d, 0x13, 0x71, 0x41, 0xba, 0x78, 0x3e, 0x17, 0x44, 0xca, 0x8c, 0xa1, 0x79,
	0xb4, 0x8d, 0xc2, 0x2b, 0x75, 0xd9, 0x8d, 0x88, 0x50, 0x0f, 0x43, 0x34, 0xd1, 0x39, 0x96, 0xc4,
	0x40, 0x3e, 0xb8, 0xe0, 0xfc, 0x22, 0x26, 0x5d, 0x7d, 0x3b, 0x5f, 0x2d, 0xba, 0x98, 0xad, 0x8d,
	0xeb, 0xc9, 0x7d, 0x97, 0x54, 0x62, 0x15, 0xe5, 0xdc, 0x1f, 0xdd, 0xf7, 0xbe, 0x14, 0x38, 0x49,
	0x88, 0xc8, 0xcb, 0x3b, 0xbc, 0xc2, 0x31, 0x9d, 0x63, 0x45, 0xba, 0xf9, 0x21, 0x73, 0xb4, 0xff,
	0xb0, 0xa0, 0xfc, 0x8c, 0xc6, 0x8a, 0x08, 0xf4, 0x21, 0x94, 0x18, 0x5e, 0x12, 0xd7, 0x6a, 0x59,
	0x9d, 0x9a, 0x5f, 0xb9, 0xf1, 0x4b, 0xa2, 0xd0, 0xb2, 0x02, 0x6d, 0x44, 0x5f, 0x41, 0x39, 0xe2,
	0x6c, 0x41, 0x2f, 0xdc, 0x42, 0xcb, 0xea, 0xd8, 0xbd, 0x43, 0x2f, 0xcb, 0xe8, 0xe5, 0x19, 0xbd,
	0xa9, 0xae, 0xe7, 0xf9, 0x4e, 0x60, 0x80, 0xe8, 0x1b, 0xa8, 0xab, 0x75, 0x42, 0xe6, 0xa1, 0x09,
	0x2c, 0xe9, 0xc0, 0x83, 0xad, 0xc0, 0x21, 0x5b, 0x3f, 0xdf, 0x09, 0x6c, 0x8d, 0x1d, 0x69, 0xa8,
	0xdf, 0x00, 0x3b, 0x0b, 0x0a, 0x53, 0xeb, 0x8b, 0x52, 0xb5, 0xe8, 0x94, 0xda, 0xaf, 0x77, 0xc1,
	0xc9, 0x4a, 0x1d, 0x5d, 0x62, 0xca, 0x4e, 0xb0, 0x8a, 0x2e, 0xd1, 0x0c, 0x9c, 0x39, 0x91, 0x8a,
	0x32, 0xac, 0x28, 0x67, 0x61, 0xc2, 0x85, 0x72, 0xab, 0x3a, 0xd1, 0x93, 0xad, 0x44, 0x3f, 0x7e,
	0xcf, 0x54, 0xbf, 0x77, 0x86, 0xe3, 0x15, 0xf1, 0xed, 0x1b, 0xbf, 0xfa, 0x79, 0xd9, 0x7d, 0xf3,
	0xa6, 0xd8, 0xb1, 0x82, 0xfd, 0x3b, 0x14, 0x13, 0x2e, 0x14, 0x7a, 0x06, 0x8d, 0x44, 0x90, 0x05,
	0xbd, 0x0e, 0x05, 0x66, 0x17, 0x44, 0xba, 0xc5, 0x56, 0xb1, 0x63, 0xf7, 0x8e, 0xbc, 0xed, 0x39,
	0x49, 0x25, 0xf4, 0x46, 0x74, 0x2e, 0x82, 0x14, 0x19, 0xd4, 0xb3, 0x38, 0x7d, 0x91, 0xe8, 0x13,
	0xd8, 0x33, 0x63, 0x12, 0xca, 0xd5, 0x62, 0x41, 0xaf, 0x75, 0x13, 0x6a, 0x41, 0xc3, 0x58, 0xa7,
	0xda, 0x88, 0xbe, 0x05, 0xc8, 0xdc, 0x61, 0x4c, 0x98, 0xbb, 0xfb, 0x78, 0xf9, 0x41, 0x2d, 0xc3,
	0x8f, 0x09, 0x43, 0x31, 0xd8, 0x92, 0xaf, 0x44, 0x44, 0x74, 0xaf, 0xdc, 0x7a, 0xcb, 0xea, 0xec,
	0xf5, 0x86, 0xde, 0xc3, 0x13, 0xed, 0xdd, 0x6f, 0xa2, 0x37, 0xe2, 0x8c, 0x91, 0x28, 0x7d, 0xfd,
	0x54, 0x33, 0xcd, 0xd6, 0x09, 0xf1, 0xab, 0x37, 0xfe, 0xee, 0x6f, 0x56, 0xc1, 0xb1, 0x02, 0x90,
	0x1b, 0x2b, 0x9a, 0xc2, 0x81, 0xc9, 0xf6, 0x76, 0x83, 0xca, 0xff, 0xb4, 0x41, 0x28, 0x0b, 0x9f,
	0xdc, 0x6d, 0x53, 0x1f, 0xea, 0x39, 0x29, 0x17, 0x4a, 0xba, 0x95, 0x56, 0xb1, 0xd3, 0xf0, 0x9d,
	0x1b, 0xbf, 0xf1, 0xbb, 0x05, 0xed, 0x5b, 0x9d, 0xcc, 0x43, 0x53, 0x89, 0x24, 0x3a, 0x82, 0xba,
	0x24, 0xe2, 0x8a, 0x88, 0x30, 0x1d, 0x50, 0xe9, 0xda, 0xad, 0x62, 0xa7, 0x16, 0xd8, 0x99, 0xed,
	0x34, 0x35, 0xa1, 0x2f, 0x01, 0x29, 0x81, 0x99, 0x4c, 0x59, 0x43, 0xdd, 0xc7, 0x88, 0xc7, 0x6e,
	0x4d, 0x4b, 0xf0, 0xde, 0xc6, 0x33, 0x31, 0x0e, 0xd4, 0x87, 0xf7, 0x71, 0x92, 0xc4, 0x34, 0x32,
	0xb3, 0x64, 0xec, 0xd2, 0x05, 0x4d, 0x7d, 0x70, 0xc7, 0x99, 0xc7, 0xc8, 0xf6, 0xd7, 0x70, 0xf0,
	0xae, 0xf6, 0xa1, 0x0a, 0x14, 0x87, 0xa7, 0x3f, 0x3b, 0x3b, 0xa8, 0x06, 0xbb, 0xe3, 0x1f, 0x46,
	0xc3, 0xb1, 0x63, 0xa1, 0x3a, 0x54, 0x8f, 0x7f, 0x9a, 0x1d, 0x07, 0xa7, 0xc3, 0xb1, 0x53, 0x78,
	0x51, 0xaa, 0x5a, 0x4e, 0x21, 0xb0, 0x25, 0xa3, 0xe1, 0x9c, 0x2f, 0x31, 0x65, 0xb2, 0xfd, 0xba,
	0x08, 0xf6, 0x1d, 0x75, 0xd0, 0x2f, 0x80, 0x16, 0xfa, 0x1a, 0x46, 0xe9, 0x3d, 0x5c, 0xa6, 0x72,
	0xe9, 0x05, 0xb5, 0x7b, 0x5f, 0xfc, 0x1b, 0x89, 0x03, 0x67, 0x71, 0x7f, 0x73, 0x4e, 0xc0, 0x56,
	0xb1, 0x4c, 0x97, 0x53, 0x91, 0x6b, 0x65, 0xd6, 0xfa, 0x5d, 0xa4, 0xe9, 0x3f, 0xe6, 0x7d, 0xc7,
	0x5f, 0x32, 0xa9, 0x04, 0xc1, 0xcb, 0x59, 0x2c, 0x47, 0x59, 0x4c, 0x00, 0x6a, 0x73, 0x46, 0x4f,
	0xa1, 0x92, 0xa5, 0xc8, 0x97, 0xa5, 0xfd, 0x78, 0x7d, 0x41, 0x1e, 0x82, 0x7c, 0xd8, 0x5f, 0xc9,
	0x74, 0xa6, 0xf8, 0xf5, 0x3a, 0x6b, 0xbc, 0xf9, 0x2e, 0x9a, 0x5b, 0x6b, 0xe0, 0x73, 0x1e, 0x67,
	0x4b, 0xd0, 0x58, 0x49, 0x32, 0x49, 0x23, 0xb4, 0x1a, 0xe8, 0x29, 0x54, 0x97, 0x44, 0xe1, 0x39,
	0x56, 0xd8, 0xec, 0x50, 0xeb, 0xa1, 0x71, 0x3c, 0x31, 0xb8, 0x60, 0x13, 0x81, 0x02, 0x70, 0x6e,
	0x67, 0x45, 0xf2, 0xe8, 0x57, 0xa2, 0xdc, 0xb2, 0x66, 0xf9, 0xf4, 0x21, 0x96, 0x59, 0x8e, 0x9f,
	0x6a, 0x78, 0xb0, 0xaf, 0xde, 0x36, 0x20, 0x64, 0x7e, 0xd4, 0x8a, 0x9e, 0x38, 0x7d, 0x6e, 0xbf,
	0xb2, 0x60, 0x6f, 0x6c, 0xda, 0xf0, 0x3f, 0x7d, 0xbc, 0xc5, 0xff, 0xfa, 0xf1, 0xfa, 0x02, 0x3a,
	0x94, 0x67, 0xcf, 0xd7, 0xd2, 0xfc, 0x8d, 0xa4, 0x7e, 0x23, 0x7f, 0x95, 0x56, 0x63, 0x62, 0xbd,
	0x2a, 0x1c, 0x1e, 0x6b, 0xf4, 0x30, 0xa1, 0xde, 0x59, 0xcf, 0xcb, 0xdd, 0xa7, 0xd3, 0x3f, 0x0b,
	0x4d, 0xed, 0x19, 0x0c, 0x86, 0x09, 0x1d, 0x0c, 0xce, 0x7a, 0x83, 0xc1, 0xad, 0xf3, 0xbc, 0xac,
	0xeb, 0xeb, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x09, 0x3a, 0xe1, 0x1d, 0xbf, 0x07, 0x00, 0x00,
}
