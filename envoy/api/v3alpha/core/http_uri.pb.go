// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/core/http_uri.proto

package envoy_api_v3alpha_core

import (
	fmt "fmt"
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

type HttpUri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType     isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a09f7c0d394634e, []int{0}
}

func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpUri.Unmarshal(m, b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return xxx_messageInfo_HttpUri.Size(m)
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.api.v3alpha.core.HttpUri")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/core/http_uri.proto", fileDescriptor_6a09f7c0d394634e)
}

var fileDescriptor_6a09f7c0d394634e = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0x9b, 0x99, 0xef, 0x73, 0x34, 0xba, 0x90, 0x2c, 0x74, 0xa6, 0x0b, 0x19, 0xfc, 0x81,
	0xae, 0x12, 0x68, 0xc1, 0x0b, 0x08, 0x0a, 0x5d, 0x96, 0x01, 0xd7, 0x25, 0x6d, 0x63, 0x1b, 0x98,
	0xf6, 0x84, 0xd3, 0x93, 0xc1, 0xb9, 0x25, 0xbd, 0x41, 0xe9, 0x4a, 0xe6, 0x6f, 0x21, 0xb8, 0x4a,
	0xc8, 0xfb, 0x1c, 0xf2, 0x9c, 0x97, 0x3f, 0xd9, 0x43, 0x05, 0xb5, 0x32, 0xde, 0xa9, 0x6a, 0x66,
	0x4a, 0xbf, 0x33, 0x6a, 0x0d, 0x68, 0xd5, 0x8e, 0xc8, 0x2f, 0x03, 0x3a, 0xe9, 0x11, 0x08, 0xc4,
	0x4d, 0x8b, 0x49, 0xe3, 0x9d, 0xec, 0x31, 0xd9, 0x60, 0xe3, 0xbb, 0x2d, 0xc0, 0xb6, 0xb4, 0xaa,
	0xa5, 0x56, 0xe1, 0x5d, 0x6d, 0x02, 0x1a, 0x72, 0x70, 0xe8, 0xe6, 0xc6, 0xb7, 0x95, 0x29, 0xdd,
	0xc6, 0x90, 0x55, 0xc3, 0xa5, 0x0b, 0xee, 0xbf, 0x18, 0x4f, 0xe6, 0x44, 0xfe, 0x0d, 0x9d, 0xc8,
	0x78, 0x1c, 0xd0, 0xa5, 0x2c, 0x67, 0x93, 0x0b, 0x9d, 0x9c, 0xf4, 0x3f, 0x8c, 0x72, 0x56, 0x34,
	0x6f, 0xe2, 0x81, 0x27, 0xeb, 0x32, 0x1c, 0xc9, 0x62, 0x1a, 0xfd, 0x8a, 0xe7, 0xa3, 0x62, 0x48,
	0xc4, 0x2b, 0x4f, 0xc8, 0xed, 0x2d, 0x04, 0x4a, 0xe3, 0x9c, 0x4d, 0x2e, 0xa7, 0x99, 0xec, 0xb4,
	0xe4, 0xa0, 0x25, 0x5f, 0x7a, 0x2d, 0x7d, 0x7d, 0xd2, 0xff, 0x3f, 0x59, 0x34, 0x1d, 0x75, 0xe7,
	0x39, 0x2b, 0x86, 0x59, 0x9d, 0x71, 0xd1, 0x6d, 0xed, 0x8f, 0x84, 0xd6, 0xec, 0x97, 0x54, 0x7b,
	0x2b, 0xe2, 0x6f, 0xcd, 0xf4, 0x33, 0x7f, 0x74, 0x20, 0xdb, 0x0e, 0x3c, 0xc2, 0x47, 0x2d, 0xff,
	0xae, 0x43, 0x5f, 0xf5, 0x2b, 0x2d, 0x9a, 0x7f, 0x17, 0x6c, 0x75, 0xd6, 0x0a, 0xcc, 0x7e, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xaa, 0xaf, 0x0d, 0x4f, 0x66, 0x01, 0x00, 0x00,
}
