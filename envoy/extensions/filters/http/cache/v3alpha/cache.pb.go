// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/cache/v3alpha/cache.proto

package envoy_extensions_filters_http_cache_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type CacheConfig struct {
	TypedConfig          *any.Any                      `protobuf:"bytes,1,opt,name=typed_config,json=typedConfig,proto3" json:"typed_config,omitempty"`
	AllowedVaryHeaders   []*v3.StringMatcher           `protobuf:"bytes,2,rep,name=allowed_vary_headers,json=allowedVaryHeaders,proto3" json:"allowed_vary_headers,omitempty"`
	KeyCreatorParams     *CacheConfig_KeyCreatorParams `protobuf:"bytes,3,opt,name=key_creator_params,json=keyCreatorParams,proto3" json:"key_creator_params,omitempty"`
	MaxBodyBytes         uint32                        `protobuf:"varint,4,opt,name=max_body_bytes,json=maxBodyBytes,proto3" json:"max_body_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CacheConfig) Reset()         { *m = CacheConfig{} }
func (m *CacheConfig) String() string { return proto.CompactTextString(m) }
func (*CacheConfig) ProtoMessage()    {}
func (*CacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e63ce5312e748b, []int{0}
}

func (m *CacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheConfig.Unmarshal(m, b)
}
func (m *CacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheConfig.Marshal(b, m, deterministic)
}
func (m *CacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheConfig.Merge(m, src)
}
func (m *CacheConfig) XXX_Size() int {
	return xxx_messageInfo_CacheConfig.Size(m)
}
func (m *CacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CacheConfig proto.InternalMessageInfo

func (m *CacheConfig) GetTypedConfig() *any.Any {
	if m != nil {
		return m.TypedConfig
	}
	return nil
}

func (m *CacheConfig) GetAllowedVaryHeaders() []*v3.StringMatcher {
	if m != nil {
		return m.AllowedVaryHeaders
	}
	return nil
}

func (m *CacheConfig) GetKeyCreatorParams() *CacheConfig_KeyCreatorParams {
	if m != nil {
		return m.KeyCreatorParams
	}
	return nil
}

func (m *CacheConfig) GetMaxBodyBytes() uint32 {
	if m != nil {
		return m.MaxBodyBytes
	}
	return 0
}

type CacheConfig_KeyCreatorParams struct {
	ExcludeScheme           bool                         `protobuf:"varint,1,opt,name=exclude_scheme,json=excludeScheme,proto3" json:"exclude_scheme,omitempty"`
	ExcludeHost             bool                         `protobuf:"varint,2,opt,name=exclude_host,json=excludeHost,proto3" json:"exclude_host,omitempty"`
	QueryParametersIncluded []*v31.QueryParameterMatcher `protobuf:"bytes,3,rep,name=query_parameters_included,json=queryParametersIncluded,proto3" json:"query_parameters_included,omitempty"`
	QueryParametersExcluded []*v31.QueryParameterMatcher `protobuf:"bytes,4,rep,name=query_parameters_excluded,json=queryParametersExcluded,proto3" json:"query_parameters_excluded,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                     `json:"-"`
	XXX_unrecognized        []byte                       `json:"-"`
	XXX_sizecache           int32                        `json:"-"`
}

func (m *CacheConfig_KeyCreatorParams) Reset()         { *m = CacheConfig_KeyCreatorParams{} }
func (m *CacheConfig_KeyCreatorParams) String() string { return proto.CompactTextString(m) }
func (*CacheConfig_KeyCreatorParams) ProtoMessage()    {}
func (*CacheConfig_KeyCreatorParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_09e63ce5312e748b, []int{0, 0}
}

func (m *CacheConfig_KeyCreatorParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Unmarshal(m, b)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Marshal(b, m, deterministic)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CacheConfig_KeyCreatorParams.Merge(m, src)
}
func (m *CacheConfig_KeyCreatorParams) XXX_Size() int {
	return xxx_messageInfo_CacheConfig_KeyCreatorParams.Size(m)
}
func (m *CacheConfig_KeyCreatorParams) XXX_DiscardUnknown() {
	xxx_messageInfo_CacheConfig_KeyCreatorParams.DiscardUnknown(m)
}

var xxx_messageInfo_CacheConfig_KeyCreatorParams proto.InternalMessageInfo

func (m *CacheConfig_KeyCreatorParams) GetExcludeScheme() bool {
	if m != nil {
		return m.ExcludeScheme
	}
	return false
}

func (m *CacheConfig_KeyCreatorParams) GetExcludeHost() bool {
	if m != nil {
		return m.ExcludeHost
	}
	return false
}

func (m *CacheConfig_KeyCreatorParams) GetQueryParametersIncluded() []*v31.QueryParameterMatcher {
	if m != nil {
		return m.QueryParametersIncluded
	}
	return nil
}

func (m *CacheConfig_KeyCreatorParams) GetQueryParametersExcluded() []*v31.QueryParameterMatcher {
	if m != nil {
		return m.QueryParametersExcluded
	}
	return nil
}

func init() {
	proto.RegisterType((*CacheConfig)(nil), "envoy.extensions.filters.http.cache.v3alpha.CacheConfig")
	proto.RegisterType((*CacheConfig_KeyCreatorParams)(nil), "envoy.extensions.filters.http.cache.v3alpha.CacheConfig.KeyCreatorParams")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/cache/v3alpha/cache.proto", fileDescriptor_09e63ce5312e748b)
}

var fileDescriptor_09e63ce5312e748b = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0x91, 0x9c, 0x1a, 0xb3, 0x76, 0x82, 0x11, 0x81, 0x28, 0x86, 0x16, 0x27, 0xa4, 0x60,
	0x68, 0x58, 0x81, 0x5d, 0x28, 0xc9, 0x2d, 0x32, 0x81, 0xa4, 0xa5, 0xe0, 0x3a, 0x10, 0xe8, 0x49,
	0xac, 0xa5, 0xb5, 0x25, 0x22, 0xed, 0x2a, 0xbb, 0x2b, 0xc5, 0x7b, 0x6b, 0x6f, 0x7d, 0x86, 0x1e,
	0xfb, 0x18, 0xbd, 0x17, 0x7a, 0xed, 0xeb, 0x94, 0x1e, 0xca, 0xfe, 0x31, 0xa9, 0xd3, 0x42, 0x31,
	0xe4, 0x24, 0x69, 0x66, 0xbe, 0xdf, 0xce, 0x37, 0xa3, 0x05, 0xaf, 0x30, 0xa9, 0xa9, 0x0c, 0xf0,
	0x52, 0x60, 0xc2, 0x33, 0x4a, 0x78, 0x30, 0xcf, 0x72, 0x81, 0x19, 0x0f, 0x52, 0x21, 0xca, 0x20,
	0x46, 0x71, 0x8a, 0x83, 0x7a, 0x84, 0xf2, 0x32, 0x45, 0xe6, 0x0b, 0x96, 0x8c, 0x0a, 0xea, 0xbd,
	0xd0, 0x42, 0x78, 0x2f, 0x84, 0x56, 0x08, 0x95, 0x10, 0x9a, 0x52, 0x2b, 0xec, 0x1d, 0x9b, 0x53,
	0x62, 0x4a, 0xe6, 0xd9, 0x22, 0x60, 0xb4, 0x12, 0x0a, 0x6a, 0x5e, 0xa2, 0x98, 0x16, 0x25, 0x25,
	0x98, 0x08, 0x6e, 0xd0, 0xbd, 0x43, 0x53, 0x2d, 0x64, 0x89, 0x83, 0x02, 0x89, 0x38, 0xc5, 0x4c,
	0x55, 0x73, 0xc1, 0x32, 0xb2, 0xb0, 0x35, 0xfb, 0x0b, 0x4a, 0x17, 0x39, 0x0e, 0xf4, 0xd7, 0xac,
	0x9a, 0x07, 0x88, 0x48, 0x9b, 0x7a, 0x5a, 0x25, 0x25, 0x0a, 0x10, 0x21, 0x54, 0x20, 0xa1, 0x2d,
	0x71, 0x81, 0x44, 0xb5, 0xa2, 0x1f, 0xfc, 0x95, 0xae, 0x31, 0x53, 0x0e, 0xee, 0xe1, 0x7b, 0x35,
	0xca, 0xb3, 0x04, 0xa9, 0x26, 0xed, 0x8b, 0x49, 0x1c, 0x7e, 0x6c, 0x82, 0xf6, 0x58, 0x39, 0x1b,
	0x6b, 0x27, 0xde, 0x19, 0xe8, 0xa8, 0x2e, 0x93, 0xc8, 0x38, 0xf3, 0x9d, 0xbe, 0x33, 0x68, 0x0f,
	0x77, 0xa1, 0x69, 0x0e, 0xae, 0x9a, 0x83, 0x67, 0x44, 0x86, 0xad, 0x9f, 0xe1, 0x93, 0x2f, 0x8e,
	0xdb, 0x72, 0xa6, 0x6d, 0xad, 0xb1, 0x88, 0x6b, 0xb0, 0x8b, 0xf2, 0x9c, 0xde, 0xe1, 0x24, 0xaa,
	0x11, 0x93, 0x51, 0x8a, 0x51, 0x82, 0x19, 0xf7, 0xdd, 0x7e, 0x63, 0xd0, 0x1e, 0x1e, 0x41, 0x33,
	0x66, 0xa5, 0x80, 0x76, 0x16, 0xb0, 0x1e, 0xc1, 0x2b, 0x3d, 0x8b, 0xb7, 0x26, 0x30, 0xf5, 0x2c,
	0xe1, 0x1a, 0x31, 0x79, 0x61, 0xf4, 0xde, 0x1d, 0xf0, 0x6e, 0xb0, 0x8c, 0x62, 0x86, 0x91, 0xa0,
	0x2c, 0x2a, 0x11, 0x43, 0x05, 0xf7, 0x1b, 0xba, 0xc1, 0x4b, 0xb8, 0xc1, 0xf2, 0xe0, 0x1f, 0x86,
	0xe1, 0x1b, 0x2c, 0xc7, 0x86, 0x38, 0xd1, 0xc0, 0x69, 0xf7, 0xe6, 0x41, 0xc4, 0x3b, 0x02, 0x3b,
	0x05, 0x5a, 0x46, 0x33, 0x9a, 0xc8, 0x68, 0x26, 0x05, 0xe6, 0xfe, 0x56, 0xdf, 0x19, 0x6c, 0x4f,
	0x3b, 0x05, 0x5a, 0x86, 0x34, 0x91, 0xa1, 0x8a, 0xf5, 0x7e, 0xb9, 0xa0, 0xfb, 0x10, 0xe6, 0x3d,
	0x07, 0x3b, 0x78, 0x19, 0xe7, 0x55, 0x82, 0x23, 0x1e, 0xa7, 0xb8, 0xc0, 0x7a, 0xa0, 0xad, 0xe9,
	0xb6, 0x8d, 0x5e, 0xe9, 0xa0, 0x77, 0x00, 0x3a, 0xab, 0xb2, 0x94, 0x72, 0xe1, 0xbb, 0xba, 0xa8,
	0x6d, 0x63, 0x17, 0x94, 0x0b, 0x2f, 0x05, 0xfb, 0xb7, 0x15, 0x66, 0xd2, 0xf8, 0xc6, 0xca, 0x5a,
	0x94, 0x11, 0x9d, 0x4f, 0xfc, 0x86, 0x1e, 0xed, 0xb1, 0x1d, 0x82, 0x59, 0x1d, 0xd4, 0xff, 0xa2,
	0x1a, 0xed, 0x3b, 0xa5, 0x9b, 0xac, 0x64, 0xab, 0x11, 0xef, 0xdd, 0xae, 0x85, 0xf9, 0xa5, 0x85,
	0xfd, 0xf3, 0x24, 0xdb, 0x49, 0xe2, 0x6f, 0x3d, 0xc2, 0x49, 0xe7, 0x16, 0x76, 0xfa, 0xfa, 0xf3,
	0xb7, 0x4f, 0xcf, 0xce, 0xc1, 0x78, 0x0d, 0x66, 0xf6, 0xb6, 0xb6, 0xb6, 0xe1, 0xff, 0xd7, 0x76,
	0x7a, 0xa2, 0x58, 0x2f, 0xc1, 0x70, 0x73, 0x56, 0xf8, 0xfe, 0xeb, 0x87, 0xef, 0x3f, 0x9a, 0x6e,
	0xcb, 0x31, 0xcf, 0xae, 0x0b, 0x4e, 0x32, 0x6a, 0x1c, 0x96, 0x8c, 0x2e, 0xe5, 0x26, 0xff, 0x56,
	0x08, 0x34, 0x79, 0xa2, 0xee, 0xc9, 0xc4, 0x99, 0x35, 0xf5, 0x85, 0x19, 0xfd, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0xda, 0xe7, 0xfd, 0x95, 0x04, 0x00, 0x00,
}
