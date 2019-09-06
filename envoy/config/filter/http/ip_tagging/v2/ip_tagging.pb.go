// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package envoy_config_filter_http_ip_tagging_v2

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/v2/envoy/api/v2/core"
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

type IPTagging_RequestType int32

const (
	IPTagging_BOTH     IPTagging_RequestType = 0
	IPTagging_INTERNAL IPTagging_RequestType = 1
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}

var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}

func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}

type IPTagging struct {
	RequestType          IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType" json:"request_type,omitempty"`
	IpTags               []*IPTagging_IPTag    `protobuf:"bytes,4,rep,name=ip_tags,json=ipTags,proto3" json:"ip_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IPTagging) Reset()         { *m = IPTagging{} }
func (m *IPTagging) String() string { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()    {}
func (*IPTagging) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0}
}

func (m *IPTagging) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging.Unmarshal(m, b)
}
func (m *IPTagging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging.Marshal(b, m, deterministic)
}
func (m *IPTagging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging.Merge(m, src)
}
func (m *IPTagging) XXX_Size() int {
	return xxx_messageInfo_IPTagging.Size(m)
}
func (m *IPTagging) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging proto.InternalMessageInfo

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

type IPTagging_IPTag struct {
	IpTagName            string            `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName,proto3" json:"ip_tag_name,omitempty"`
	IpList               []*core.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList,proto3" json:"ip_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IPTagging_IPTag) Reset()         { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()    {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}

func (m *IPTagging_IPTag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPTagging_IPTag.Unmarshal(m, b)
}
func (m *IPTagging_IPTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPTagging_IPTag.Marshal(b, m, deterministic)
}
func (m *IPTagging_IPTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging_IPTag.Merge(m, src)
}
func (m *IPTagging_IPTag) XXX_Size() int {
	return xxx_messageInfo_IPTagging_IPTag.Size(m)
}
func (m *IPTagging_IPTag) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging_IPTag.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging_IPTag proto.InternalMessageInfo

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*core.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
	proto.RegisterType((*IPTagging)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging.IPTag")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto", fileDescriptor_f4275d0b367744d2)
}

var fileDescriptor_f4275d0b367744d2 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4f, 0x4f, 0xc2, 0x30,
	0x14, 0xb7, 0x13, 0x70, 0x74, 0x84, 0x90, 0x5d, 0x24, 0xc4, 0x28, 0xe1, 0x60, 0x38, 0x75, 0xc9,
	0xd0, 0x70, 0xf2, 0xe0, 0x0c, 0x09, 0x24, 0x04, 0xc9, 0xb2, 0x83, 0xf1, 0x20, 0xa9, 0xac, 0xcc,
	0x1a, 0x58, 0x6b, 0x5b, 0x17, 0x77, 0xf5, 0x23, 0xf8, 0x29, 0xfd, 0x0c, 0x9e, 0x4c, 0x5b, 0x54,
	0x8e, 0x78, 0x7b, 0xaf, 0xef, 0xfd, 0xfe, 0xf5, 0xc1, 0x21, 0xc9, 0x0b, 0x56, 0x06, 0x4b, 0x96,
	0xaf, 0x68, 0x16, 0xac, 0xe8, 0x5a, 0x11, 0x11, 0x3c, 0x29, 0xc5, 0x03, 0xca, 0x17, 0x0a, 0x67,
	0x19, 0xcd, 0xb3, 0xa0, 0x08, 0x77, 0x3a, 0xc4, 0x05, 0x53, 0xcc, 0x3f, 0x37, 0x40, 0x64, 0x81,
	0xc8, 0x02, 0x91, 0x06, 0xa2, 0x9d, 0xd5, 0x22, 0xec, 0x9c, 0x59, 0x01, 0xcc, 0xa9, 0xa6, 0x59,
	0x32, 0x41, 0x02, 0x9c, 0xa6, 0x82, 0x48, 0x69, 0x89, 0x3a, 0xc7, 0x05, 0x5e, 0xd3, 0x14, 0x2b,
	0x12, 0xfc, 0x14, 0x76, 0xd0, 0xfb, 0x74, 0x60, 0x7d, 0x32, 0x4f, 0x2c, 0x95, 0xff, 0x0c, 0x1b,
	0x82, 0xbc, 0xbc, 0x12, 0xa9, 0x16, 0xaa, 0xe4, 0xa4, 0x0d, 0xba, 0xa0, 0xdf, 0x0c, 0xaf, 0xd0,
	0x7e, 0x36, 0xd0, 0x2f, 0x11, 0x8a, 0x2d, 0x4b, 0x52, 0x72, 0x12, 0xb9, 0x5f, 0x51, 0xf5, 0x1d,
	0x38, 0x2d, 0x10, 0x7b, 0xe2, 0xef, 0xd9, 0xbf, 0x87, 0x47, 0x16, 0x2d, 0xdb, 0x95, 0xee, 0x61,
	0xdf, 0x0b, 0x87, 0xff, 0x97, 0x31, 0x95, 0x11, 0xf8, 0x00, 0x8e, 0x0b, 0xe2, 0x1a, 0xe5, 0x09,
	0xce, 0x64, 0xe7, 0x01, 0x56, 0xcd, 0xc8, 0x3f, 0x85, 0x9e, 0xc5, 0x2e, 0x72, 0xbc, 0xb1, 0x79,
	0xea, 0x71, 0xdd, 0x6c, 0xcd, 0xf0, 0x86, 0xf8, 0x97, 0xc6, 0xc4, 0x9a, 0x4a, 0xd5, 0x76, 0x8c,
	0x89, 0x93, 0xad, 0x09, 0xcc, 0xa9, 0x96, 0xd2, 0x5f, 0x89, 0x6e, 0x68, 0x2a, 0x62, 0x9c, 0x67,
	0x44, 0xf3, 0x4f, 0xa9, 0x54, 0xbd, 0x01, 0xf4, 0x76, 0x12, 0xfa, 0x2e, 0xac, 0x44, 0xb7, 0xc9,
	0xb8, 0x75, 0xe0, 0x37, 0xa0, 0x3b, 0x99, 0x25, 0xa3, 0x78, 0x76, 0x3d, 0x6d, 0x01, 0xdd, 0x8d,
	0xee, 0xb6, 0x9d, 0x13, 0x8d, 0xe1, 0x05, 0x65, 0x96, 0x9e, 0x0b, 0xf6, 0x56, 0xee, 0x19, 0x37,
	0x6a, 0x4e, 0xf8, 0x36, 0xef, 0x5c, 0x9f, 0x6c, 0x0e, 0x1e, 0x6b, 0xe6, 0x76, 0x83, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf9, 0xde, 0xf2, 0x64, 0x58, 0x02, 0x00, 0x00,
}
