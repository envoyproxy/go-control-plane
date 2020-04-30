// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/trace/v3/lightstep.proto

package envoy_config_trace_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type LightstepConfig_PropagationMode int32

const (
	LightstepConfig_ENVOY         LightstepConfig_PropagationMode = 0
	LightstepConfig_LIGHTSTEP     LightstepConfig_PropagationMode = 1
	LightstepConfig_B3            LightstepConfig_PropagationMode = 2
	LightstepConfig_TRACE_CONTEXT LightstepConfig_PropagationMode = 3
)

var LightstepConfig_PropagationMode_name = map[int32]string{
	0: "ENVOY",
	1: "LIGHTSTEP",
	2: "B3",
	3: "TRACE_CONTEXT",
}

var LightstepConfig_PropagationMode_value = map[string]int32{
	"ENVOY":         0,
	"LIGHTSTEP":     1,
	"B3":            2,
	"TRACE_CONTEXT": 3,
}

func (x LightstepConfig_PropagationMode) String() string {
	return proto.EnumName(LightstepConfig_PropagationMode_name, int32(x))
}

func (LightstepConfig_PropagationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fb8137e7d804fcc9, []int{0, 0}
}

type LightstepConfig struct {
	CollectorCluster     string                            `protobuf:"bytes,1,opt,name=collector_cluster,json=collectorCluster,proto3" json:"collector_cluster,omitempty"`
	AccessTokenFile      string                            `protobuf:"bytes,2,opt,name=access_token_file,json=accessTokenFile,proto3" json:"access_token_file,omitempty"`
	PropagationModes     []LightstepConfig_PropagationMode `protobuf:"varint,3,rep,packed,name=propagation_modes,json=propagationModes,proto3,enum=envoy.config.trace.v3.LightstepConfig_PropagationMode" json:"propagation_modes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *LightstepConfig) Reset()         { *m = LightstepConfig{} }
func (m *LightstepConfig) String() string { return proto.CompactTextString(m) }
func (*LightstepConfig) ProtoMessage()    {}
func (*LightstepConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb8137e7d804fcc9, []int{0}
}

func (m *LightstepConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightstepConfig.Unmarshal(m, b)
}
func (m *LightstepConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightstepConfig.Marshal(b, m, deterministic)
}
func (m *LightstepConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightstepConfig.Merge(m, src)
}
func (m *LightstepConfig) XXX_Size() int {
	return xxx_messageInfo_LightstepConfig.Size(m)
}
func (m *LightstepConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LightstepConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LightstepConfig proto.InternalMessageInfo

func (m *LightstepConfig) GetCollectorCluster() string {
	if m != nil {
		return m.CollectorCluster
	}
	return ""
}

func (m *LightstepConfig) GetAccessTokenFile() string {
	if m != nil {
		return m.AccessTokenFile
	}
	return ""
}

func (m *LightstepConfig) GetPropagationModes() []LightstepConfig_PropagationMode {
	if m != nil {
		return m.PropagationModes
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.trace.v3.LightstepConfig_PropagationMode", LightstepConfig_PropagationMode_name, LightstepConfig_PropagationMode_value)
	proto.RegisterType((*LightstepConfig)(nil), "envoy.config.trace.v3.LightstepConfig")
}

func init() {
	proto.RegisterFile("envoy/config/trace/v3/lightstep.proto", fileDescriptor_fb8137e7d804fcc9)
}

var fileDescriptor_fb8137e7d804fcc9 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x6b, 0x13, 0x41,
	0x18, 0x86, 0xdd, 0x8d, 0x49, 0xc9, 0x40, 0x9a, 0xc9, 0x80, 0x18, 0x0a, 0x96, 0x18, 0x11, 0x8b,
	0x94, 0x59, 0xe8, 0x16, 0x0f, 0xde, 0xdc, 0x90, 0xaa, 0x50, 0xdb, 0x10, 0x17, 0xd1, 0xd3, 0x32,
	0x6e, 0xbe, 0x6e, 0x07, 0xa7, 0x33, 0xe3, 0xcc, 0x64, 0x69, 0x6f, 0xe2, 0xc9, 0x9b, 0xe0, 0x49,
	0xfc, 0x29, 0xde, 0x05, 0xaf, 0xfe, 0x09, 0x7f, 0x80, 0xc7, 0x1e, 0xa4, 0x64, 0x27, 0x6d, 0xe9,
	0xb6, 0xb7, 0x85, 0xf7, 0x79, 0xd9, 0x87, 0xef, 0x1d, 0xf4, 0x10, 0x64, 0xa9, 0x4e, 0xa2, 0x5c,
	0xc9, 0x03, 0x5e, 0x44, 0xce, 0xb0, 0x1c, 0xa2, 0x32, 0x8e, 0x04, 0x2f, 0x0e, 0x9d, 0x75, 0xa0,
	0xa9, 0x36, 0xca, 0x29, 0x72, 0xa7, 0xc2, 0xa8, 0xc7, 0x68, 0x85, 0xd1, 0x32, 0x5e, 0x5b, 0x9f,
	0xcf, 0x34, 0x8b, 0x98, 0x94, 0xca, 0x31, 0xc7, 0x95, 0xb4, 0xd1, 0x11, 0x2f, 0x0c, 0x73, 0xe0,
	0x6b, 0x6b, 0xf7, 0xae, 0xe5, 0xd6, 0x31, 0x37, 0xb7, 0xcb, 0xf8, 0xfe, 0xb5, 0xb8, 0x04, 0x63,
	0xb9, 0x92, 0x5c, 0x16, 0x4b, 0xe4, 0x6e, 0xc9, 0x04, 0x9f, 0x31, 0x07, 0xd1, 0xf9, 0x87, 0x0f,
	0x86, 0x7f, 0x43, 0xd4, 0xdd, 0x3d, 0xb7, 0x1c, 0x55, 0x5e, 0x64, 0x1b, 0xf5, 0x72, 0x25, 0x04,
	0xe4, 0x4e, 0x99, 0x2c, 0x17, 0x73, 0xeb, 0xc0, 0xf4, 0x83, 0x41, 0xb0, 0xd1, 0x4e, 0x56, 0x4e,
	0x93, 0xdb, 0x26, 0x1c, 0x04, 0x53, 0x7c, 0x41, 0x8c, 0x3c, 0x40, 0x62, 0xd4, 0x63, 0x79, 0x0e,
	0xd6, 0x66, 0x4e, 0x7d, 0x00, 0x99, 0x1d, 0x70, 0x01, 0xfd, 0xf0, 0x6a, 0xab, 0xeb, 0x89, 0x74,
	0x01, 0xec, 0x70, 0x01, 0xc4, 0xa0, 0x9e, 0x36, 0x4a, 0xb3, 0xa2, 0x12, 0xcf, 0x8e, 0xd4, 0x0c,
	0x6c, 0xbf, 0x31, 0x68, 0x6c, 0xac, 0x6e, 0x3d, 0xa1, 0x37, 0x1e, 0x8b, 0xd6, 0x6c, 0xe9, 0xe4,
	0xb2, 0xff, 0x4a, 0xcd, 0x20, 0xe9, 0x9c, 0x26, 0xe8, 0x5b, 0xb0, 0x32, 0x6c, 0x7e, 0x0e, 0x42,
	0x1c, 0x4c, 0xb1, 0xbe, 0x9a, 0xdb, 0xe1, 0x0e, 0xea, 0xd6, 0x3a, 0xa4, 0x8d, 0x9a, 0xe3, 0xbd,
	0x37, 0xfb, 0xef, 0xf0, 0x2d, 0xd2, 0x41, 0xed, 0xdd, 0x97, 0xcf, 0x5f, 0xa4, 0xaf, 0xd3, 0xf1,
	0x04, 0x07, 0xa4, 0x85, 0xc2, 0x24, 0xc6, 0x21, 0xe9, 0xa1, 0x4e, 0x3a, 0x7d, 0x36, 0x1a, 0x67,
	0xa3, 0xfd, 0xbd, 0x74, 0xfc, 0x36, 0xc5, 0x8d, 0xa7, 0x9b, 0x3f, 0x7e, 0x7d, 0x59, 0x7f, 0xb4,
	0x9c, 0xbe, 0xa6, 0xb9, 0x55, 0xd7, 0x4c, 0x3e, 0xfe, 0xfb, 0xfe, 0xff, 0x6b, 0x73, 0x93, 0x3c,
	0xf6, 0x38, 0x1c, 0x3b, 0x90, 0x8b, 0x8d, 0xac, 0xaf, 0x18, 0x4b, 0x2f, 0x1f, 0x4b, 0xb9, 0xcd,
	0x84, 0x3e, 0x64, 0x3f, 0x3f, 0xfd, 0xfe, 0xd3, 0x0a, 0x71, 0x88, 0x1e, 0x70, 0xe5, 0x8f, 0xa1,
	0x8d, 0x3a, 0x3e, 0xb9, 0xf9, 0x2e, 0xc9, 0xea, 0xc5, 0x1f, 0x27, 0x8b, 0x65, 0x27, 0xc1, 0xfb,
	0x56, 0x35, 0x71, 0x7c, 0x16, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xb0, 0xf9, 0xde, 0x9d, 0x02, 0x00,
	0x00,
}
