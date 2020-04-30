// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto

package envoy_config_filter_thrift_rate_limit_v2alpha1

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2 "github.com/envoyproxy/go-control-plane/envoy/config/ratelimit/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	Domain               string                     `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Stage                uint32                     `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	FailureModeDeny      bool                       `protobuf:"varint,4,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	RateLimitService     *v2.RateLimitServiceConfig `protobuf:"bytes,5,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_961acdee13c1bd42, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.thrift.rate_limit.v2alpha1.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/thrift/rate_limit/v2alpha1/rate_limit.proto", fileDescriptor_961acdee13c1bd42)
}

var fileDescriptor_961acdee13c1bd42 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x6f, 0xd4, 0x30,
	0x18, 0x86, 0xe5, 0xd0, 0xbb, 0xb6, 0x46, 0xc0, 0xe1, 0x85, 0x50, 0xa9, 0x10, 0xc1, 0x12, 0x75,
	0xb0, 0xd5, 0xbb, 0x15, 0x09, 0x29, 0xdc, 0x08, 0x52, 0x15, 0x7e, 0x40, 0xe4, 0x36, 0x5f, 0x52,
	0x8b, 0xc4, 0x8e, 0x9c, 0x2f, 0xa1, 0xd9, 0x98, 0x59, 0x18, 0x61, 0xe6, 0x67, 0xf0, 0x0b, 0x58,
	0xf9, 0x2b, 0x8c, 0x1d, 0x10, 0x4a, 0x9c, 0x84, 0x3b, 0x98, 0xd8, 0x92, 0xef, 0x7d, 0xf3, 0x7c,
	0x7e, 0x5f, 0x87, 0xbe, 0x04, 0xdd, 0x9a, 0x4e, 0x5c, 0x19, 0x9d, 0xa9, 0x5c, 0x64, 0xaa, 0x40,
	0xb0, 0x02, 0xaf, 0xad, 0xca, 0x50, 0x58, 0x89, 0x90, 0x14, 0xaa, 0x54, 0x28, 0xda, 0xb5, 0x2c,
	0xaa, 0x6b, 0x79, 0xbe, 0x33, 0xe3, 0x95, 0x35, 0x68, 0x18, 0x1f, 0x00, 0xdc, 0x01, 0xb8, 0x03,
	0x70, 0x07, 0xe0, 0x3b, 0xe6, 0x09, 0x70, 0xf2, 0x7c, 0x6f, 0x61, 0xef, 0x98, 0x36, 0x08, 0x5b,
	0xd4, 0x0e, 0x7a, 0xf2, 0x24, 0x37, 0x26, 0x2f, 0x40, 0x0c, 0x6f, 0x97, 0x4d, 0x26, 0xd2, 0xc6,
	0x4a, 0x54, 0x46, 0x4f, 0x7a, 0x93, 0x56, 0x52, 0x48, 0xad, 0x0d, 0x0e, 0xe3, 0x5a, 0x94, 0x2a,
	0xef, 0x59, 0xa3, 0x7e, 0xfa, 0x8f, 0x5e, 0xa3, 0xc4, 0x66, 0xc2, 0x3f, 0x6a, 0x65, 0xa1, 0x52,
	0x89, 0x20, 0xa6, 0x07, 0x27, 0x3c, 0xfb, 0xec, 0xd1, 0xe3, 0x58, 0x22, 0xbc, 0xee, 0x8f, 0xc4,
	0x9e, 0xd2, 0x65, 0x6a, 0x4a, 0xa9, 0xb4, 0x4f, 0x02, 0x12, 0x1e, 0x47, 0x87, 0xb7, 0xd1, 0x81,
	0xf5, 0x02, 0x12, 0x8f, 0x63, 0x76, 0x4a, 0x17, 0x35, 0xca, 0x1c, 0x7c, 0x2f, 0x20, 0xe1, 0xbd,
	0x41, 0x3f, 0xf3, 0x7c, 0x1a, 0xbb, 0x29, 0xdb, 0xd0, 0x43, 0x54, 0x25, 0x98, 0x06, 0xfd, 0x3b,
	0x01, 0x09, 0xef, 0xae, 0x1f, 0x73, 0x97, 0x8b, 0x4f, 0xb9, 0xf8, 0x76, 0xcc, 0x15, 0x4f, 0x4e,
	0x76, 0x46, 0x1f, 0x66, 0x52, 0x15, 0x8d, 0x85, 0xa4, 0x34, 0x29, 0x24, 0x29, 0xe8, 0xce, 0x3f,
	0x08, 0x48, 0x78, 0x14, 0x3f, 0x18, 0x85, 0x37, 0x26, 0x85, 0x2d, 0xe8, 0x8e, 0x29, 0xca, 0xfe,
	0x54, 0x9c, 0xd4, 0x60, 0x5b, 0x75, 0x05, 0xfe, 0x62, 0xd8, 0x75, 0xbe, 0x7f, 0x31, 0x73, 0xd1,
	0xbc, 0x5d, 0xf3, 0x39, 0xe2, 0x5b, 0xf7, 0xc9, 0xab, 0xc1, 0x13, 0x1d, 0xdd, 0x46, 0x8b, 0x8f,
	0xc4, 0x5b, 0x91, 0x78, 0x65, 0xff, 0x72, 0x44, 0x5f, 0xc9, 0xcf, 0x2f, 0xbf, 0x3e, 0x2d, 0xb6,
	0x2c, 0x72, 0x58, 0xb8, 0x41, 0xd0, 0x75, 0x5f, 0xed, 0x78, 0xe7, 0x35, 0xd7, 0x80, 0xef, 0x8d,
	0x7d, 0x37, 0x5e, 0x7e, 0x52, 0x59, 0x73, 0xd3, 0xcd, 0xe2, 0xce, 0xfe, 0xcd, 0xb7, 0x0f, 0xdf,
	0x7f, 0x2c, 0xbd, 0x15, 0xa1, 0x2f, 0x94, 0x71, 0xa7, 0x74, 0xde, 0xff, 0xfb, 0x93, 0xa2, 0xfb,
	0x73, 0x8c, 0x8b, 0xbe, 0xce, 0x0b, 0x72, 0xb9, 0x1c, 0x7a, 0xdd, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x79, 0xde, 0xf6, 0xb1, 0xd5, 0x02, 0x00, 0x00,
}
