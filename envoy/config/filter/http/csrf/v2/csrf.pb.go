// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/csrf/v2/csrf.proto

package envoy_config_filter_http_csrf_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
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

type CsrfPolicy struct {
	FilterEnabled        *core.RuntimeFractionalPercent `protobuf:"bytes,1,opt,name=filter_enabled,json=filterEnabled,proto3" json:"filter_enabled,omitempty"`
	ShadowEnabled        *core.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=shadow_enabled,json=shadowEnabled,proto3" json:"shadow_enabled,omitempty"`
	AdditionalOrigins    []*matcher.StringMatcher       `protobuf:"bytes,3,rep,name=additional_origins,json=additionalOrigins,proto3" json:"additional_origins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *CsrfPolicy) Reset()         { *m = CsrfPolicy{} }
func (m *CsrfPolicy) String() string { return proto.CompactTextString(m) }
func (*CsrfPolicy) ProtoMessage()    {}
func (*CsrfPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9146cdf92353980, []int{0}
}

func (m *CsrfPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsrfPolicy.Unmarshal(m, b)
}
func (m *CsrfPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsrfPolicy.Marshal(b, m, deterministic)
}
func (m *CsrfPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsrfPolicy.Merge(m, src)
}
func (m *CsrfPolicy) XXX_Size() int {
	return xxx_messageInfo_CsrfPolicy.Size(m)
}
func (m *CsrfPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_CsrfPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_CsrfPolicy proto.InternalMessageInfo

func (m *CsrfPolicy) GetFilterEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.FilterEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetShadowEnabled() *core.RuntimeFractionalPercent {
	if m != nil {
		return m.ShadowEnabled
	}
	return nil
}

func (m *CsrfPolicy) GetAdditionalOrigins() []*matcher.StringMatcher {
	if m != nil {
		return m.AdditionalOrigins
	}
	return nil
}

func init() {
	proto.RegisterType((*CsrfPolicy)(nil), "envoy.config.filter.http.csrf.v2.CsrfPolicy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/csrf/v2/csrf.proto", fileDescriptor_a9146cdf92353980)
}

var fileDescriptor_a9146cdf92353980 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x49, 0x4b, 0x1f, 0xef, 0xa5, 0xb4, 0x3c, 0xb3, 0xb1, 0x14, 0xd1, 0xea, 0xaa, 0x58,
	0x99, 0x81, 0xf4, 0x1f, 0x44, 0x74, 0x27, 0x86, 0xb8, 0x15, 0xca, 0x6d, 0x32, 0x49, 0x06, 0xd2,
	0x99, 0x61, 0x66, 0x1a, 0x9b, 0xbf, 0xe0, 0x46, 0x97, 0xfe, 0x4e, 0x97, 0x2e, 0x44, 0x32, 0x37,
	0x55, 0x77, 0xe2, 0x2a, 0x09, 0xf7, 0x9c, 0xef, 0xde, 0x9c, 0xe3, 0x2f, 0x98, 0xa8, 0x65, 0x43,
	0x53, 0x29, 0x72, 0x5e, 0xd0, 0x9c, 0x57, 0x96, 0x69, 0x5a, 0x5a, 0xab, 0x68, 0x6a, 0x74, 0x4e,
	0xeb, 0xd0, 0x3d, 0x89, 0xd2, 0xd2, 0xca, 0x60, 0xe6, 0xc4, 0x04, 0xc5, 0x04, 0xc5, 0xa4, 0x15,
	0x13, 0x27, 0xaa, 0xc3, 0xe9, 0x11, 0xe2, 0x40, 0x71, 0x67, 0x95, 0x9a, 0xd1, 0x35, 0x18, 0x86,
	0xfe, 0xe9, 0x09, 0x4e, 0x6d, 0xa3, 0x18, 0xdd, 0x80, 0x4d, 0x4b, 0xa6, 0xa9, 0xb1, 0x9a, 0x8b,
	0xa2, 0x13, 0x1c, 0x6f, 0x33, 0x05, 0x14, 0x84, 0x90, 0x16, 0x2c, 0x97, 0xc2, 0xd0, 0x0d, 0x2f,
	0x34, 0xd8, 0x3d, 0xe0, 0xb0, 0x86, 0x8a, 0x67, 0x60, 0x19, 0xdd, 0xbf, 0xe0, 0xe0, 0xec, 0xb9,
	0xe7, 0xfb, 0x97, 0x46, 0xe7, 0xb1, 0xac, 0x78, 0xda, 0x04, 0xf7, 0xfe, 0x18, 0xaf, 0x5b, 0x31,
	0x01, 0xeb, 0x8a, 0x65, 0x13, 0x6f, 0xe6, 0xcd, 0x87, 0xe1, 0x82, 0xe0, 0x1f, 0x80, 0xe2, 0xa4,
	0x0e, 0x49, 0x7b, 0x1f, 0x49, 0xb6, 0xc2, 0xf2, 0x0d, 0xbb, 0xd6, 0x90, 0xb6, 0x2b, 0xa1, 0x8a,
	0x99, 0x4e, 0x99, 0xb0, 0xd1, 0xdf, 0xb7, 0x68, 0xf0, 0xe8, 0xf5, 0xfe, 0x7b, 0xc9, 0x08, 0x61,
	0x57, 0xc8, 0x0a, 0x12, 0x7f, 0x6c, 0x4a, 0xc8, 0xe4, 0xc3, 0x27, 0xbd, 0xf7, 0x6b, 0x7a, 0x32,
	0x42, 0xc4, 0x9e, 0x19, 0xfb, 0x01, 0x64, 0x19, 0x47, 0xcd, 0x4a, 0x6a, 0x5e, 0x70, 0x61, 0x26,
	0xfd, 0x59, 0x7f, 0x3e, 0x0c, 0x4f, 0x3b, 0x6e, 0x9b, 0x1b, 0xe9, 0x72, 0x23, 0x77, 0x2e, 0xb7,
	0x1b, 0xfc, 0x4a, 0x0e, 0xbe, 0xcc, 0xb7, 0xe8, 0x8d, 0xe4, 0xeb, 0xcb, 0xfb, 0xd3, 0xe0, 0x22,
	0x38, 0x47, 0x33, 0xdb, 0x59, 0x26, 0x4c, 0x9b, 0x69, 0x57, 0x9c, 0xf9, 0xde, 0xdc, 0x12, 0x2a,
	0x55, 0x82, 0x4f, 0xb8, 0xc4, 0x5d, 0x4a, 0xcb, 0x5d, 0x43, 0x7e, 0xaa, 0x3b, 0xfa, 0xe7, 0x32,
	0x6f, 0x1b, 0x88, 0xbd, 0xf5, 0x1f, 0x57, 0xc5, 0xf2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x24, 0xb2,
	0xd8, 0x82, 0x53, 0x02, 0x00, 0x00,
}
