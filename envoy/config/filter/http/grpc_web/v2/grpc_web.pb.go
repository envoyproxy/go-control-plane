// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_web/v2/grpc_web.proto

package envoy_config_filter_http_grpc_web_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type GrpcWeb struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcWeb) Reset()         { *m = GrpcWeb{} }
func (m *GrpcWeb) String() string { return proto.CompactTextString(m) }
func (*GrpcWeb) ProtoMessage()    {}
func (*GrpcWeb) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fe810f98c50c033, []int{0}
}

func (m *GrpcWeb) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcWeb.Unmarshal(m, b)
}
func (m *GrpcWeb) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcWeb.Marshal(b, m, deterministic)
}
func (m *GrpcWeb) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcWeb.Merge(m, src)
}
func (m *GrpcWeb) XXX_Size() int {
	return xxx_messageInfo_GrpcWeb.Size(m)
}
func (m *GrpcWeb) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcWeb.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcWeb proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GrpcWeb)(nil), "envoy.config.filter.http.grpc_web.v2.GrpcWeb")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_web/v2/grpc_web.proto", fileDescriptor_8fe810f98c50c033)
}

var fileDescriptor_8fe810f98c50c033 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x2f, 0x4f, 0x4d, 0xd2, 0x2f, 0x33,
	0x82, 0xb3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0x9a, 0xf4, 0x20, 0x9a, 0xf4,
	0x20, 0x9a, 0xf4, 0x40, 0x9a, 0xf4, 0xe0, 0x0a, 0xcb, 0x8c, 0xa4, 0xe4, 0x4a, 0x53, 0x0a, 0x12,
	0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x73, 0x33, 0xd3,
	0x8b, 0x12, 0x4b, 0x52, 0x21, 0xa6, 0x28, 0x71, 0x72, 0xb1, 0xbb, 0x17, 0x15, 0x24, 0x87, 0xa7,
	0x26, 0x39, 0x55, 0x7d, 0x9a, 0xf1, 0xaf, 0x9f, 0xd5, 0x40, 0x48, 0x0f, 0x62, 0x70, 0x6a, 0x45,
	0x49, 0x6a, 0x5e, 0x31, 0x48, 0x0b, 0xd4, 0xf0, 0x62, 0x74, 0xd3, 0x8d, 0x13, 0x73, 0x0a, 0x32,
	0x12, 0xb9, 0x8c, 0x32, 0xf3, 0x21, 0x5a, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x88, 0x71, 0x96,
	0x13, 0x0f, 0xd4, 0xd2, 0x00, 0x90, 0x23, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0xae, 0x31, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xc3, 0xa6, 0x81, 0xfe, 0x0a, 0x01, 0x00, 0x00,
}
