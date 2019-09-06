// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto

package envoy_config_filter_http_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	v2alpha "github.com/envoyproxy/go-control-plane/v2/envoy/config/common/dynamic_forward_proxy/v2alpha"
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

type FilterConfig struct {
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto", fileDescriptor_85a2356b260c47da)
}

var fileDescriptor_85a2356b260c47da = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0x89, 0xa0, 0x48, 0x14, 0x19, 0xbb, 0x38, 0x76, 0x12, 0x4f, 0x9e, 0x12, 0x98, 0xe0,
	0x7d, 0x5d, 0xd9, 0xb9, 0xec, 0x0f, 0x84, 0x2c, 0x49, 0x6d, 0xa0, 0xcd, 0x0b, 0x69, 0xac, 0xf6,
	0x07, 0x78, 0xf1, 0x27, 0x7b, 0x92, 0xe4, 0xa9, 0x50, 0x28, 0x0a, 0xbb, 0x05, 0x3e, 0xde, 0xf7,
	0xbe, 0x17, 0x7a, 0x30, 0x6e, 0x80, 0x91, 0x2b, 0x70, 0xb5, 0x7d, 0xe6, 0xb5, 0x6d, 0xa3, 0x09,
	0xbc, 0x89, 0xd1, 0x73, 0x3d, 0x3a, 0xd9, 0x59, 0x25, 0x6a, 0x08, 0xaf, 0x32, 0x68, 0xe1, 0x03,
	0xbc, 0x8d, 0x7c, 0xd8, 0xc8, 0xd6, 0x37, 0x72, 0x9e, 0x32, 0x1f, 0x20, 0xc2, 0xf2, 0x29, 0x3b,
	0x19, 0x3a, 0x19, 0x3a, 0x59, 0x72, 0xb2, 0xf9, 0xa9, 0x6f, 0xe7, 0x7a, 0x3b, 0x69, 0x51, 0xd0,
	0x75, 0xe0, 0xfe, 0xcb, 0x70, 0xbd, 0x50, 0x52, 0x35, 0x06, 0x57, 0xaf, 0x6f, 0x07, 0xd9, 0x5a,
	0x2d, 0xa3, 0xe1, 0x3f, 0x0f, 0x04, 0xf7, 0xef, 0x84, 0x5e, 0xef, 0x73, 0xc9, 0x2e, 0xeb, 0x97,
	0x2f, 0x74, 0xf1, 0x3b, 0x2c, 0x70, 0xe5, 0x8a, 0xdc, 0x91, 0x87, 0xab, 0xcd, 0x96, 0x4d, 0xfa,
	0xb1, 0xe3, 0xef, 0x74, 0x56, 0xba, 0x7e, 0x97, 0x4c, 0x28, 0x2f, 0x2e, 0x3f, 0x8b, 0xf3, 0x0f,
	0x72, 0xb6, 0x20, 0x87, 0x1b, 0x3d, 0x25, 0x47, 0x5a, 0x5a, 0xc0, 0x05, 0x68, 0x38, 0xed, 0xaf,
	0x8a, 0x55, 0x89, 0x78, 0x8f, 0xb4, 0x4a, 0xb0, 0x4a, 0x97, 0x56, 0xe4, 0x78, 0x91, 0x4f, 0x7e,
	0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x07, 0x70, 0x19, 0x00, 0xdc, 0x01, 0x00, 0x00,
}
