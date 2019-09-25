// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/fault/v3alpha/fault.proto

package envoy_config_filter_fault_v3alpha

import (
	fmt "fmt"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
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

type FaultDelay_FaultDelayType int32

const (
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}

var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}

func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_26a4880b6121b9f2, []int{0, 0}
}

type FaultDelay struct {
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	//	*FaultDelay_HeaderDelay_
	FaultDelaySecifier   isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	Percentage           *_type.FractionalPercent        `protobuf:"bytes,4,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *FaultDelay) Reset()         { *m = FaultDelay{} }
func (m *FaultDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()    {}
func (*FaultDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_26a4880b6121b9f2, []int{0}
}

func (m *FaultDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultDelay.Unmarshal(m, b)
}
func (m *FaultDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultDelay.Marshal(b, m, deterministic)
}
func (m *FaultDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay.Merge(m, src)
}
func (m *FaultDelay) XXX_Size() int {
	return xxx_messageInfo_FaultDelay.Size(m)
}
func (m *FaultDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay proto.InternalMessageInfo

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
}

type FaultDelay_FixedDelay struct {
	FixedDelay *duration.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,proto3,oneof"`
}

type FaultDelay_HeaderDelay_ struct {
	HeaderDelay *FaultDelay_HeaderDelay `protobuf:"bytes,5,opt,name=header_delay,json=headerDelay,proto3,oneof"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (*FaultDelay_HeaderDelay_) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetFixedDelay() *duration.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *FaultDelay) GetHeaderDelay() *FaultDelay_HeaderDelay {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_HeaderDelay_); ok {
		return x.HeaderDelay
	}
	return nil
}

func (m *FaultDelay) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultDelay_FixedDelay)(nil),
		(*FaultDelay_HeaderDelay_)(nil),
	}
}

type FaultDelay_HeaderDelay struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultDelay_HeaderDelay) Reset()         { *m = FaultDelay_HeaderDelay{} }
func (m *FaultDelay_HeaderDelay) String() string { return proto.CompactTextString(m) }
func (*FaultDelay_HeaderDelay) ProtoMessage()    {}
func (*FaultDelay_HeaderDelay) Descriptor() ([]byte, []int) {
	return fileDescriptor_26a4880b6121b9f2, []int{0, 0}
}

func (m *FaultDelay_HeaderDelay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Unmarshal(m, b)
}
func (m *FaultDelay_HeaderDelay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Marshal(b, m, deterministic)
}
func (m *FaultDelay_HeaderDelay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultDelay_HeaderDelay.Merge(m, src)
}
func (m *FaultDelay_HeaderDelay) XXX_Size() int {
	return xxx_messageInfo_FaultDelay_HeaderDelay.Size(m)
}
func (m *FaultDelay_HeaderDelay) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultDelay_HeaderDelay.DiscardUnknown(m)
}

var xxx_messageInfo_FaultDelay_HeaderDelay proto.InternalMessageInfo

type FaultRateLimit struct {
	// Types that are valid to be assigned to LimitType:
	//	*FaultRateLimit_FixedLimit_
	//	*FaultRateLimit_HeaderLimit_
	LimitType            isFaultRateLimit_LimitType `protobuf_oneof:"limit_type"`
	Percentage           *_type.FractionalPercent   `protobuf:"bytes,2,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FaultRateLimit) Reset()         { *m = FaultRateLimit{} }
func (m *FaultRateLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit) ProtoMessage()    {}
func (*FaultRateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_26a4880b6121b9f2, []int{1}
}

func (m *FaultRateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit.Merge(m, src)
}
func (m *FaultRateLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit.Size(m)
}
func (m *FaultRateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit proto.InternalMessageInfo

type isFaultRateLimit_LimitType interface {
	isFaultRateLimit_LimitType()
}

type FaultRateLimit_FixedLimit_ struct {
	FixedLimit *FaultRateLimit_FixedLimit `protobuf:"bytes,1,opt,name=fixed_limit,json=fixedLimit,proto3,oneof"`
}

type FaultRateLimit_HeaderLimit_ struct {
	HeaderLimit *FaultRateLimit_HeaderLimit `protobuf:"bytes,3,opt,name=header_limit,json=headerLimit,proto3,oneof"`
}

func (*FaultRateLimit_FixedLimit_) isFaultRateLimit_LimitType() {}

func (*FaultRateLimit_HeaderLimit_) isFaultRateLimit_LimitType() {}

func (m *FaultRateLimit) GetLimitType() isFaultRateLimit_LimitType {
	if m != nil {
		return m.LimitType
	}
	return nil
}

func (m *FaultRateLimit) GetFixedLimit() *FaultRateLimit_FixedLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_FixedLimit_); ok {
		return x.FixedLimit
	}
	return nil
}

func (m *FaultRateLimit) GetHeaderLimit() *FaultRateLimit_HeaderLimit {
	if x, ok := m.GetLimitType().(*FaultRateLimit_HeaderLimit_); ok {
		return x.HeaderLimit
	}
	return nil
}

func (m *FaultRateLimit) GetPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FaultRateLimit) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FaultRateLimit_FixedLimit_)(nil),
		(*FaultRateLimit_HeaderLimit_)(nil),
	}
}

type FaultRateLimit_FixedLimit struct {
	LimitKbps            uint64   `protobuf:"varint,1,opt,name=limit_kbps,json=limitKbps,proto3" json:"limit_kbps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_FixedLimit) Reset()         { *m = FaultRateLimit_FixedLimit{} }
func (m *FaultRateLimit_FixedLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_FixedLimit) ProtoMessage()    {}
func (*FaultRateLimit_FixedLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_26a4880b6121b9f2, []int{1, 0}
}

func (m *FaultRateLimit_FixedLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit_FixedLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit_FixedLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_FixedLimit.Merge(m, src)
}
func (m *FaultRateLimit_FixedLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit_FixedLimit.Size(m)
}
func (m *FaultRateLimit_FixedLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_FixedLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_FixedLimit proto.InternalMessageInfo

func (m *FaultRateLimit_FixedLimit) GetLimitKbps() uint64 {
	if m != nil {
		return m.LimitKbps
	}
	return 0
}

type FaultRateLimit_HeaderLimit struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaultRateLimit_HeaderLimit) Reset()         { *m = FaultRateLimit_HeaderLimit{} }
func (m *FaultRateLimit_HeaderLimit) String() string { return proto.CompactTextString(m) }
func (*FaultRateLimit_HeaderLimit) ProtoMessage()    {}
func (*FaultRateLimit_HeaderLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_26a4880b6121b9f2, []int{1, 1}
}

func (m *FaultRateLimit_HeaderLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Unmarshal(m, b)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Marshal(b, m, deterministic)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.Merge(m, src)
}
func (m *FaultRateLimit_HeaderLimit) XXX_Size() int {
	return xxx_messageInfo_FaultRateLimit_HeaderLimit.Size(m)
}
func (m *FaultRateLimit_HeaderLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_FaultRateLimit_HeaderLimit.DiscardUnknown(m)
}

var xxx_messageInfo_FaultRateLimit_HeaderLimit proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.config.filter.fault.v3alpha.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
	proto.RegisterType((*FaultDelay)(nil), "envoy.config.filter.fault.v3alpha.FaultDelay")
	proto.RegisterType((*FaultDelay_HeaderDelay)(nil), "envoy.config.filter.fault.v3alpha.FaultDelay.HeaderDelay")
	proto.RegisterType((*FaultRateLimit)(nil), "envoy.config.filter.fault.v3alpha.FaultRateLimit")
	proto.RegisterType((*FaultRateLimit_FixedLimit)(nil), "envoy.config.filter.fault.v3alpha.FaultRateLimit.FixedLimit")
	proto.RegisterType((*FaultRateLimit_HeaderLimit)(nil), "envoy.config.filter.fault.v3alpha.FaultRateLimit.HeaderLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/fault/v3alpha/fault.proto", fileDescriptor_26a4880b6121b9f2)
}

var fileDescriptor_26a4880b6121b9f2 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0x34, 0xd5, 0xee, 0xab, 0x4a, 0x0d, 0x82, 0xb5, 0x8b, 0xa2, 0x3d, 0xc8, 0x22,
	0x38, 0x03, 0xbb, 0x5e, 0x04, 0x17, 0x61, 0xa8, 0xa1, 0xae, 0x1e, 0x4a, 0xf0, 0xe0, 0xc9, 0x32,
	0x69, 0x26, 0xed, 0x60, 0xec, 0x84, 0x34, 0x2d, 0x9b, 0xaf, 0xe4, 0x87, 0xf1, 0x93, 0xf8, 0x01,
	0x64, 0x4f, 0x32, 0xef, 0x4d, 0x9a, 0xf5, 0xb4, 0xee, 0x6d, 0x1e, 0xf3, 0xde, 0xef, 0xff, 0x9f,
	0xff, 0x3c, 0x78, 0xad, 0x36, 0x7b, 0x53, 0xf3, 0xa5, 0xd9, 0x64, 0x7a, 0xc5, 0x33, 0x9d, 0x57,
	0xaa, 0xe4, 0x99, 0xdc, 0xe5, 0x15, 0xdf, 0x9f, 0xc9, 0xbc, 0x58, 0x4b, 0xaa, 0x58, 0x51, 0x9a,
	0xca, 0x84, 0x2f, 0xb0, 0x9d, 0x51, 0x3b, 0xa3, 0x76, 0x46, 0x0d, 0xae, 0x7d, 0x3c, 0x22, 0x62,
	0x55, 0x17, 0x8a, 0x17, 0xaa, 0x5c, 0xaa, 0x8d, 0x1b, 0x1e, 0x3f, 0x5b, 0x19, 0xb3, 0xca, 0x15,
	0xc7, 0x2a, 0xd9, 0x65, 0x3c, 0xdd, 0x95, 0xb2, 0xd2, 0x66, 0xe3, 0xee, 0x1f, 0xef, 0x65, 0xae,
	0x53, 0x59, 0x29, 0xde, 0x1c, 0xe8, 0x62, 0xf2, 0xcb, 0x07, 0x88, 0xac, 0xc8, 0x54, 0xe5, 0xb2,
	0x0e, 0x23, 0x18, 0x64, 0xfa, 0x52, 0xa5, 0x8b, 0xd4, 0x96, 0xa3, 0xee, 0x73, 0xef, 0x64, 0x70,
	0xfa, 0x84, 0x11, 0x9d, 0x35, 0x74, 0x36, 0x75, 0x74, 0xd1, 0xbf, 0x12, 0xbd, 0x9f, 0x9e, 0xff,
	0xaa, 0x33, 0xeb, 0xc4, 0x80, 0x93, 0xc4, 0xf9, 0x06, 0xf7, 0xd6, 0x4a, 0xa6, 0xaa, 0x74, 0xa0,
	0x1e, 0x82, 0xde, 0xb2, 0x1b, 0xdf, 0xc8, 0x5a, 0x33, 0x6c, 0x86, 0x04, 0x3c, 0xcf, 0x3a, 0xf1,
	0x60, 0xdd, 0x96, 0xe1, 0x39, 0x80, 0x0b, 0x40, 0xae, 0xd4, 0x28, 0x40, 0xfa, 0x53, 0x47, 0xb7,
	0xf1, 0xb0, 0xa8, 0x94, 0x4b, 0xeb, 0x50, 0xe6, 0x73, 0xea, 0x8b, 0xaf, 0x0d, 0x8c, 0xef, 0xc3,
	0xe0, 0x1a, 0x7c, 0x72, 0x0c, 0x0f, 0x5a, 0xd9, 0x2f, 0x75, 0xa1, 0xc2, 0x23, 0xe8, 0x45, 0x1f,
	0xbf, 0x7e, 0x98, 0x0e, 0x3b, 0xe2, 0x18, 0x1e, 0xa1, 0x43, 0x7a, 0xc9, 0x62, 0xab, 0x96, 0x3a,
	0xd3, 0xaa, 0x0c, 0xbb, 0x7f, 0x84, 0x77, 0x11, 0xf4, 0xfd, 0x61, 0xf7, 0x22, 0xe8, 0x7b, 0x43,
	0x3f, 0x0e, 0xac, 0xf0, 0xe4, 0xb7, 0xef, 0x60, 0xb1, 0xac, 0xd4, 0x67, 0xfd, 0x43, 0x57, 0xe1,
	0xa2, 0x09, 0x35, 0xb7, 0xe5, 0xc8, 0x43, 0xb7, 0xef, 0xfe, 0x37, 0x8b, 0x03, 0x87, 0x45, 0x16,
	0x82, 0xc7, 0x43, 0xda, 0x24, 0x90, 0x1c, 0xd2, 0x26, 0x05, 0xfa, 0xb6, 0xf3, 0xdb, 0x2b, 0x50,
	0x28, 0x8d, 0x84, 0x4b, 0x9c, 0x34, 0xfe, 0x4d, 0xdc, 0xbf, 0x6d, 0xe2, 0x6f, 0x00, 0x5a, 0xfb,
	0xe1, 0x4b, 0x00, 0x74, 0xba, 0xf8, 0x9e, 0x14, 0x5b, 0x0c, 0x24, 0x10, 0x77, 0xaf, 0x44, 0x70,
	0xea, 0x9f, 0x78, 0xf1, 0x11, 0x5e, 0x7d, 0x4a, 0x8a, 0x6d, 0xfb, 0x4f, 0x38, 0x26, 0x1e, 0x36,
	0x63, 0x56, 0x10, 0x3f, 0x40, 0xbc, 0x07, 0xae, 0x0d, 0xd9, 0x28, 0x4a, 0x73, 0x59, 0xdf, 0xfc,
	0x66, 0x41, 0xfb, 0x3e, 0xb7, 0xbb, 0x3c, 0xf7, 0x92, 0x3b, 0xb8, 0xd4, 0x67, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x91, 0x8e, 0xfa, 0xae, 0x03, 0x00, 0x00,
}
