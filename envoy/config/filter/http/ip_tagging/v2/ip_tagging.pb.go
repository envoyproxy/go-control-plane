// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package envoy_config_filter_http_ip_tagging_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
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
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x4f, 0x6b, 0x14, 0x31,
	0x14, 0x37, 0x63, 0xb7, 0xce, 0x66, 0x4a, 0x59, 0x72, 0x71, 0x59, 0xa4, 0x96, 0x1e, 0xa4, 0xa7,
	0x44, 0x66, 0x95, 0x9e, 0x3c, 0x38, 0x52, 0x70, 0xa1, 0xac, 0xcb, 0x30, 0x07, 0xf1, 0xe0, 0xf2,
	0xec, 0xa4, 0xd3, 0xc8, 0x6c, 0x12, 0x93, 0x74, 0xe8, 0xdc, 0xc4, 0x2f, 0x20, 0x78, 0xf2, 0x23,
	0x7a, 0xf6, 0xe8, 0x41, 0x24, 0xc9, 0xa8, 0x0b, 0x5e, 0xd6, 0x5b, 0x5e, 0xde, 0xfb, 0xfd, 0x7b,
	0x0f, 0x9f, 0x71, 0xd9, 0xa9, 0x9e, 0x5d, 0x2a, 0x79, 0x25, 0x1a, 0x76, 0x25, 0x5a, 0xc7, 0x0d,
	0xbb, 0x76, 0x4e, 0x33, 0xa1, 0xd7, 0x0e, 0x9a, 0x46, 0xc8, 0x86, 0x75, 0xf9, 0x56, 0x45, 0xb5,
	0x51, 0x4e, 0x91, 0x47, 0x01, 0x48, 0x23, 0x90, 0x46, 0x20, 0xf5, 0x40, 0xba, 0x35, 0xda, 0xe5,
	0xb3, 0x87, 0x51, 0x00, 0xb4, 0xf0, 0x34, 0x97, 0xca, 0x70, 0x06, 0x75, 0x6d, 0xb8, 0xb5, 0x91,
	0x68, 0x76, 0x74, 0x53, 0x6b, 0x60, 0x20, 0xa5, 0x72, 0xe0, 0x84, 0x92, 0x96, 0x6d, 0x44, 0x63,
	0xc0, 0xf1, 0xa1, 0x7f, 0xbf, 0x83, 0x56, 0xd4, 0xe0, 0x38, 0xfb, 0xfd, 0x88, 0x8d, 0x93, 0x6f,
	0x09, 0x1e, 0x2f, 0x56, 0x55, 0x94, 0x22, 0xef, 0xf1, 0x81, 0xe1, 0x1f, 0x6e, 0xb8, 0x75, 0x6b,
	0xd7, 0x6b, 0x3e, 0x45, 0xc7, 0xe8, 0xf4, 0x30, 0x7f, 0x46, 0x77, 0xb3, 0x49, 0xff, 0x10, 0xd1,
	0x32, 0xb2, 0x54, 0xbd, 0xe6, 0x45, 0xfa, 0xa3, 0x18, 0x7d, 0x42, 0xc9, 0x04, 0x95, 0x99, 0xf9,
	0xfb, 0x4d, 0xde, 0xe0, 0x7b, 0x11, 0x6d, 0xa7, 0x7b, 0xc7, 0x77, 0x4f, 0xb3, 0xfc, 0xec, 0xff,
	0x65, 0xc2, 0x2b, 0x08, 0x7c, 0x41, 0x49, 0x8a, 0xca, 0x7d, 0xa1, 0x2b, 0x68, 0xec, 0xec, 0x2d,
	0x1e, 0x85, 0x16, 0x39, 0xc2, 0x59, 0xc4, 0xae, 0x25, 0x6c, 0x62, 0x9e, 0x71, 0x39, 0x0e, 0x53,
	0x4b, 0xd8, 0x70, 0xf2, 0x34, 0x98, 0x68, 0x85, 0x75, 0xd3, 0x24, 0x98, 0x78, 0x30, 0x98, 0x00,
	0x2d, 0xbc, 0x94, 0x5f, 0x35, 0x7d, 0x21, 0x6a, 0x53, 0x82, 0x6c, 0xb8, 0xe7, 0xbf, 0x10, 0xd6,
	0x9d, 0xcc, 0x71, 0xb6, 0x95, 0x90, 0xa4, 0x78, 0xaf, 0x78, 0x55, 0xbd, 0x9c, 0xdc, 0x21, 0x07,
	0x38, 0x5d, 0x2c, 0xab, 0xf3, 0x72, 0xf9, 0xfc, 0x62, 0x82, 0x7c, 0x75, 0xfe, 0x7a, 0xa8, 0x92,
	0xe2, 0x23, 0xfa, 0xfe, 0xf5, 0xe7, 0xe7, 0x51, 0x4e, 0x1e, 0x47, 0x09, 0x7e, 0xeb, 0xb8, 0xb4,
	0xfe, 0x58, 0x43, 0x56, 0xfb, 0x6f, 0xd8, 0x39, 0xb4, 0xfa, 0x1a, 0xf0, 0x13, 0xa1, 0xa2, 0x2f,
	0x6d, 0xd4, 0x6d, 0xbf, 0xe3, 0x9e, 0x8a, 0xc3, 0x85, 0x1e, 0x16, 0xb5, 0xf2, 0xb7, 0x5e, 0xa1,
	0x77, 0xfb, 0xe1, 0xe8, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x60, 0xdc, 0x02, 0xb1,
	0x02, 0x00, 0x00,
}
