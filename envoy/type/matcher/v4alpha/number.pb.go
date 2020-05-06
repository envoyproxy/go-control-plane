// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v4alpha/number.proto

package envoy_type_matcher_v4alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type DoubleMatcher struct {
	// Types that are valid to be assigned to MatchPattern:
	//	*DoubleMatcher_Range
	//	*DoubleMatcher_Exact
	MatchPattern         isDoubleMatcher_MatchPattern `protobuf_oneof:"match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *DoubleMatcher) Reset()         { *m = DoubleMatcher{} }
func (m *DoubleMatcher) String() string { return proto.CompactTextString(m) }
func (*DoubleMatcher) ProtoMessage()    {}
func (*DoubleMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_01a24a128e699b79, []int{0}
}

func (m *DoubleMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DoubleMatcher.Unmarshal(m, b)
}
func (m *DoubleMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DoubleMatcher.Marshal(b, m, deterministic)
}
func (m *DoubleMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DoubleMatcher.Merge(m, src)
}
func (m *DoubleMatcher) XXX_Size() int {
	return xxx_messageInfo_DoubleMatcher.Size(m)
}
func (m *DoubleMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_DoubleMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_DoubleMatcher proto.InternalMessageInfo

type isDoubleMatcher_MatchPattern interface {
	isDoubleMatcher_MatchPattern()
}

type DoubleMatcher_Range struct {
	Range *v3.DoubleRange `protobuf:"bytes,1,opt,name=range,proto3,oneof"`
}

type DoubleMatcher_Exact struct {
	Exact float64 `protobuf:"fixed64,2,opt,name=exact,proto3,oneof"`
}

func (*DoubleMatcher_Range) isDoubleMatcher_MatchPattern() {}

func (*DoubleMatcher_Exact) isDoubleMatcher_MatchPattern() {}

func (m *DoubleMatcher) GetMatchPattern() isDoubleMatcher_MatchPattern {
	if m != nil {
		return m.MatchPattern
	}
	return nil
}

func (m *DoubleMatcher) GetRange() *v3.DoubleRange {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Range); ok {
		return x.Range
	}
	return nil
}

func (m *DoubleMatcher) GetExact() float64 {
	if x, ok := m.GetMatchPattern().(*DoubleMatcher_Exact); ok {
		return x.Exact
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DoubleMatcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DoubleMatcher_Range)(nil),
		(*DoubleMatcher_Exact)(nil),
	}
}

func init() {
	proto.RegisterType((*DoubleMatcher)(nil), "envoy.type.matcher.v4alpha.DoubleMatcher")
}

func init() {
	proto.RegisterFile("envoy/type/matcher/v4alpha/number.proto", fileDescriptor_01a24a128e699b79)
}

var fileDescriptor_01a24a128e699b79 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x4d, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0xd2, 0x2f,
	0x33, 0x49, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0xcf, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x02, 0x2b, 0xd4, 0x03, 0x29, 0xd4, 0x83, 0x2a, 0xd4, 0x83, 0x2a,
	0x94, 0x92, 0x44, 0x32, 0xa4, 0xcc, 0x58, 0xbf, 0x28, 0x31, 0x2f, 0x3d, 0x15, 0xa2, 0x4d, 0x4a,
	0xb6, 0x34, 0xa5, 0x20, 0x51, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf,
	0x58, 0xbf, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x18, 0x2a, 0xad, 0x88, 0x21, 0x5d, 0x96, 0x5a, 0x54,
	0x9c, 0x99, 0x9f, 0x97, 0x99, 0x97, 0x0e, 0x55, 0x22, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58,
	0x92, 0xaa, 0x0f, 0x63, 0x40, 0x24, 0x94, 0xe6, 0x32, 0x72, 0xf1, 0xba, 0xe4, 0x97, 0x26, 0xe5,
	0xa4, 0xfa, 0x42, 0xdc, 0x23, 0x64, 0xc4, 0xc5, 0x0a, 0xb6, 0x5b, 0x82, 0x51, 0x81, 0x51, 0x83,
	0xdb, 0x48, 0x4a, 0x0f, 0xc9, 0xcd, 0x65, 0xc6, 0x7a, 0x10, 0xc5, 0x41, 0x20, 0x15, 0x1e, 0x0c,
	0x41, 0x10, 0xa5, 0x42, 0x62, 0x5c, 0xac, 0xa9, 0x15, 0x89, 0xc9, 0x25, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x8c, 0x20, 0x71, 0x30, 0xd7, 0x4a, 0x6b, 0xd6, 0xd1, 0x0e, 0x39, 0x55, 0x2e, 0x65, 0x6c,
	0xde, 0x86, 0x19, 0x05, 0xb5, 0xd7, 0x49, 0x84, 0x8b, 0x17, 0x2c, 0x17, 0x5f, 0x90, 0x58, 0x52,
	0x92, 0x5a, 0x94, 0x27, 0xc4, 0xfc, 0xc3, 0x89, 0xd1, 0xc9, 0x71, 0x57, 0xc3, 0x89, 0x8b, 0x6c,
	0x4c, 0x02, 0xcc, 0x5c, 0x1a, 0x99, 0xf9, 0x10, 0xa7, 0x14, 0x14, 0xe5, 0x57, 0x54, 0xea, 0xe1,
	0x0e, 0x49, 0x27, 0x6e, 0x3f, 0x70, 0x98, 0x07, 0x80, 0x3c, 0x18, 0xc0, 0x98, 0xc4, 0x06, 0xf6,
	0xa9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x58, 0x15, 0xab, 0xa6, 0x01, 0x00, 0x00,
}
