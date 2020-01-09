// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/common/dynamic_forward_proxy/v2alpha/dns_cache.proto

package envoy_config_common_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
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

type DnsCacheConfig struct {
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DnsLookupFamily      v2.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.api.v2.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	DnsRefreshRate       *duration.Duration         `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	HostTtl              *duration.Duration         `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	MaxHosts             *wrappers.UInt32Value      `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *DnsCacheConfig) Reset()         { *m = DnsCacheConfig{} }
func (m *DnsCacheConfig) String() string { return proto.CompactTextString(m) }
func (*DnsCacheConfig) ProtoMessage()    {}
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d2d9297e0c94cb56, []int{0}
}

func (m *DnsCacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsCacheConfig.Unmarshal(m, b)
}
func (m *DnsCacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsCacheConfig.Marshal(b, m, deterministic)
}
func (m *DnsCacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsCacheConfig.Merge(m, src)
}
func (m *DnsCacheConfig) XXX_Size() int {
	return xxx_messageInfo_DnsCacheConfig.Size(m)
}
func (m *DnsCacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsCacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsCacheConfig proto.InternalMessageInfo

func (m *DnsCacheConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsCacheConfig) GetDnsLookupFamily() v2.Cluster_DnsLookupFamily {
	if m != nil {
		return m.DnsLookupFamily
	}
	return v2.Cluster_AUTO
}

func (m *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if m != nil {
		return m.DnsRefreshRate
	}
	return nil
}

func (m *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if m != nil {
		return m.HostTtl
	}
	return nil
}

func (m *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxHosts
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsCacheConfig)(nil), "envoy.config.common.dynamic_forward_proxy.v2alpha.DnsCacheConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/dynamic_forward_proxy/v2alpha/dns_cache.proto", fileDescriptor_d2d9297e0c94cb56)
}

var fileDescriptor_d2d9297e0c94cb56 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x40, 0xd7, 0xa1, 0x65, 0xbb, 0x41, 0x94, 0x25, 0x17, 0x4a, 0x41, 0x28, 0x42, 0x42, 0xaa,
	0xf6, 0x60, 0x8b, 0x54, 0x88, 0x1b, 0x88, 0xb4, 0x42, 0x20, 0x81, 0xb4, 0x8a, 0x80, 0x0b, 0x87,
	0x68, 0x36, 0x71, 0x5a, 0x0b, 0xc7, 0xb6, 0x6c, 0x27, 0xdb, 0x5e, 0xf9, 0x81, 0xbd, 0xf2, 0x0d,
	0xfc, 0x13, 0x3f, 0xc1, 0x71, 0x0f, 0x08, 0xd9, 0x6e, 0x0f, 0x08, 0x10, 0xec, 0x2d, 0xf6, 0x78,
	0x5e, 0x66, 0xde, 0x4c, 0xfc, 0x82, 0x8a, 0x5e, 0x6e, 0x49, 0x25, 0x45, 0xc3, 0x56, 0xa4, 0x92,
	0x6d, 0x2b, 0x05, 0xa9, 0xb7, 0x02, 0x5a, 0x56, 0x95, 0x8d, 0xd4, 0xe7, 0xa0, 0xeb, 0x52, 0x69,
	0xb9, 0xd9, 0x92, 0x3e, 0x03, 0xae, 0xd6, 0x40, 0x6a, 0x61, 0xca, 0x0a, 0xaa, 0x35, 0xc5, 0x4a,
	0x4b, 0x2b, 0x93, 0xc7, 0x1e, 0x81, 0x03, 0x02, 0x07, 0x04, 0xfe, 0x23, 0x02, 0xef, 0x10, 0xd3,
	0x69, 0xf8, 0x2b, 0x28, 0x46, 0xfa, 0x8c, 0x54, 0xbc, 0x33, 0x96, 0xea, 0x80, 0x9b, 0x3e, 0x58,
	0x49, 0xb9, 0xe2, 0x94, 0xf8, 0xd3, 0x59, 0xd7, 0x90, 0xba, 0xd3, 0x60, 0x99, 0x14, 0x7f, 0x8b,
	0x9f, 0x6b, 0x50, 0x8a, 0x6a, 0xb3, 0x8f, 0x77, 0xb5, 0x02, 0x02, 0x42, 0x48, 0xeb, 0xd3, 0x0c,
	0x69, 0xd9, 0x4a, 0x83, 0xdd, 0x95, 0x3b, 0xbd, 0xd3, 0x03, 0x67, 0x35, 0x58, 0x4a, 0xf6, 0x1f,
	0x21, 0xf0, 0xf0, 0x5b, 0x14, 0x8f, 0x97, 0xc2, 0x2c, 0x5c, 0x6b, 0x0b, 0xdf, 0x4c, 0x72, 0x2f,
	0x1e, 0x08, 0x68, 0xe9, 0x04, 0xa5, 0x68, 0x76, 0x94, 0x1f, 0x5e, 0xe6, 0x03, 0x1d, 0xa5, 0xa8,
	0xf0, 0x97, 0xc9, 0xc7, 0xf8, 0xb6, 0x53, 0xc1, 0xa5, 0xfc, 0xd4, 0xa9, 0xb2, 0x81, 0x96, 0xf1,
	0xed, 0x24, 0x4a, 0xd1, 0x6c, 0x9c, 0x3d, 0xc2, 0xc1, 0x09, 0x28, 0x86, 0xfb, 0x0c, 0x2f, 0x76,
	0x0d, 0x2e, 0x85, 0x79, 0xe3, 0x5f, 0xbf, 0xf4, 0x8f, 0xf3, 0xd1, 0x65, 0x3e, 0xfc, 0x8c, 0xa2,
	0x63, 0x54, 0xdc, 0xaa, 0x7f, 0x0d, 0x25, 0x6f, 0xe3, 0x63, 0x07, 0xd7, 0xb4, 0xd1, 0xd4, 0xac,
	0x4b, 0x57, 0xff, 0xe4, 0x5a, 0x8a, 0x66, 0x37, 0xb2, 0xbb, 0x38, 0x08, 0xc0, 0x7b, 0x01, 0x78,
	0xb9, 0x13, 0xe4, 0x79, 0x5f, 0x51, 0x74, 0x72, 0x50, 0x8c, 0x6b, 0x61, 0x8a, 0x90, 0x5b, 0x80,
	0xa5, 0xc9, 0xb3, 0x78, 0xb4, 0x96, 0xc6, 0x96, 0xd6, 0xf2, 0xc9, 0xe0, 0xff, 0x31, 0x87, 0x2e,
	0xe9, 0x9d, 0xe5, 0x49, 0x1e, 0x1f, 0xb5, 0xb0, 0x29, 0xdd, 0xd1, 0x4c, 0x86, 0x1e, 0x70, 0xff,
	0x37, 0xc0, 0xfb, 0xd7, 0xc2, 0xce, 0xb3, 0x0f, 0xc0, 0x3b, 0xea, 0x5d, 0x9d, 0x44, 0xe9, 0x41,
	0x31, 0x6a, 0x61, 0xf3, 0xca, 0xa5, 0xe5, 0x17, 0xe8, 0xfb, 0x97, 0x1f, 0x17, 0xc3, 0xa7, 0xc9,
	0x93, 0x20, 0x87, 0x6e, 0x2c, 0x15, 0xc6, 0x4d, 0xe8, 0x1f, 0x4b, 0x33, 0xf7, 0x4b, 0x13, 0x3f,
	0x67, 0x32, 0x68, 0x0d, 0xd7, 0x57, 0xde, 0xba, 0xfc, 0xe6, 0x7e, 0xba, 0xa7, 0xae, 0xe2, 0x53,
	0x74, 0x76, 0xdd, 0x97, 0x3e, 0xff, 0x19, 0x00, 0x00, 0xff, 0xff, 0x06, 0x2b, 0x72, 0x90, 0x05,
	0x03, 0x00, 0x00,
}
