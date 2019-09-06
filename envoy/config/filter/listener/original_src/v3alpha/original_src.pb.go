// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/listener/original_src/v3alpha/original_src.proto

package envoy_config_filter_listener_original_src_v3alpha

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
	BindPort             bool     `protobuf:"varint,1,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	Mark                 uint32   `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalSrc) Reset()         { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()    {}
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff8467f0be045f8d, []int{0}
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

func (m *OriginalSrc) GetBindPort() bool {
	if m != nil {
		return m.BindPort
	}
	return false
}

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.config.filter.listener.original_src.v3alpha.OriginalSrc")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/original_src/v3alpha/original_src.proto", fileDescriptor_ff8467f0be045f8d)
}

var fileDescriptor_ff8467f0be045f8d = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xce, 0xc1, 0x4b, 0x86, 0x30,
	0x18, 0xc7, 0x71, 0x16, 0x11, 0x6f, 0x8b, 0x20, 0x76, 0x49, 0xea, 0x22, 0x9d, 0x3c, 0x6d, 0x84,
	0xf7, 0x02, 0xe9, 0x9e, 0x18, 0x9d, 0x65, 0xea, 0xb4, 0x87, 0xd6, 0x9e, 0xf1, 0x38, 0x24, 0xff,
	0xfb, 0x70, 0x2a, 0xe4, 0xf1, 0xbd, 0x3d, 0x3c, 0xf0, 0xfb, 0xf0, 0xe5, 0x6f, 0xc6, 0x4d, 0x38,
	0xab, 0x16, 0x5d, 0x0f, 0x83, 0xea, 0xc1, 0x06, 0x43, 0xca, 0xc2, 0x18, 0x8c, 0x33, 0xa4, 0x90,
	0x60, 0x00, 0xa7, 0x6d, 0x3d, 0x52, 0xab, 0xa6, 0x5c, 0x5b, 0xff, 0xa5, 0x0f, 0x4f, 0xe9, 0x09,
	0x03, 0x8a, 0xe7, 0xa8, 0xc8, 0x55, 0x91, 0xab, 0x22, 0x77, 0x45, 0x1e, 0x06, 0x9b, 0xf2, 0x70,
	0x3f, 0x69, 0x0b, 0x9d, 0x0e, 0x46, 0xed, 0xc7, 0x6a, 0x3d, 0xbd, 0xf0, 0x9b, 0xf7, 0x6d, 0xf0,
	0x41, 0xad, 0x78, 0xe4, 0xd7, 0x0d, 0xb8, 0xae, 0xf6, 0x48, 0x21, 0x61, 0x29, 0xcb, 0x4e, 0xd5,
	0x69, 0x79, 0x94, 0x48, 0x41, 0x08, 0x7e, 0xf9, 0xa3, 0xe9, 0x3b, 0xb9, 0x48, 0x59, 0x76, 0x5b,
	0xc5, 0xbb, 0xf8, 0xe4, 0xaf, 0x80, 0x32, 0x06, 0x79, 0xc2, 0xdf, 0x59, 0x9e, 0xdd, 0x56, 0xdc,
	0xfd, 0x0b, 0x28, 0x97, 0xa8, 0x92, 0x35, 0x57, 0xb1, 0x2e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x80, 0xb9, 0x74, 0xbd, 0x31, 0x01, 0x00, 0x00,
}
