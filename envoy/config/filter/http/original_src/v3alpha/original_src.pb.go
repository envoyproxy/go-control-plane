// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/original_src/v3alpha/original_src.proto

package envoy_config_filter_http_original_src_v3alpha

import (
	fmt "fmt"
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

type OriginalSrc struct {
	Mark                 uint32   `protobuf:"varint,1,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalSrc) Reset()         { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()    {}
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return fileDescriptor_05c93f7450536988, []int{0}
}

func (m *OriginalSrc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalSrc.Unmarshal(m, b)
}
func (m *OriginalSrc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalSrc.Marshal(b, m, deterministic)
}
func (m *OriginalSrc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalSrc.Merge(m, src)
}
func (m *OriginalSrc) XXX_Size() int {
	return xxx_messageInfo_OriginalSrc.Size(m)
}
func (m *OriginalSrc) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalSrc.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalSrc proto.InternalMessageInfo

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.config.filter.http.original_src.v3alpha.OriginalSrc")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/original_src/v3alpha/original_src.proto", fileDescriptor_05c93f7450536988)
}

var fileDescriptor_05c93f7450536988 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0xcf, 0x2f, 0xca, 0x4c, 0xcf, 0xcc, 0x4b, 0xcc, 0x89, 0x2f, 0x2e,
	0x4a, 0xd6, 0x2f, 0x33, 0x4e, 0xcc, 0x29, 0xc8, 0x48, 0x44, 0x11, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xd2, 0x05, 0x9b, 0xa0, 0x07, 0x31, 0x41, 0x0f, 0x62, 0x82, 0x1e, 0xc8, 0x04, 0x3d,
	0x14, 0xc5, 0x50, 0x13, 0xa4, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61,
	0x0c, 0x88, 0x39, 0x4a, 0x8a, 0x5c, 0xdc, 0xfe, 0x50, 0x0d, 0xc1, 0x45, 0xc9, 0x42, 0x42, 0x5c,
	0x2c, 0xb9, 0x89, 0x45, 0xd9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x60, 0xb6, 0x53, 0x20,
	0x97, 0x75, 0x66, 0xbe, 0x1e, 0xd8, 0xbe, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0x92, 0xac, 0x76,
	0x12, 0x40, 0x32, 0x3f, 0x00, 0x64, 0x67, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x72, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x88, 0x0c, 0xb8, 0x5c, 0x08, 0x01, 0x00, 0x00,
}
