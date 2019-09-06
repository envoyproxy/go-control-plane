// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/cluster/v2alpha/outlier_detection_event.proto

package envoy_data_cluster_v2alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type OutlierEjectionType int32

const (
	OutlierEjectionType_CONSECUTIVE_5XX                  OutlierEjectionType = 0
	OutlierEjectionType_CONSECUTIVE_GATEWAY_FAILURE      OutlierEjectionType = 1
	OutlierEjectionType_SUCCESS_RATE                     OutlierEjectionType = 2
	OutlierEjectionType_CONSECUTIVE_LOCAL_ORIGIN_FAILURE OutlierEjectionType = 3
	OutlierEjectionType_SUCCESS_RATE_LOCAL_ORIGIN        OutlierEjectionType = 4
)

var OutlierEjectionType_name = map[int32]string{
	0: "CONSECUTIVE_5XX",
	1: "CONSECUTIVE_GATEWAY_FAILURE",
	2: "SUCCESS_RATE",
	3: "CONSECUTIVE_LOCAL_ORIGIN_FAILURE",
	4: "SUCCESS_RATE_LOCAL_ORIGIN",
}

var OutlierEjectionType_value = map[string]int32{
	"CONSECUTIVE_5XX":                  0,
	"CONSECUTIVE_GATEWAY_FAILURE":      1,
	"SUCCESS_RATE":                     2,
	"CONSECUTIVE_LOCAL_ORIGIN_FAILURE": 3,
	"SUCCESS_RATE_LOCAL_ORIGIN":        4,
}

func (x OutlierEjectionType) String() string {
	return proto.EnumName(OutlierEjectionType_name, int32(x))
}

func (OutlierEjectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{0}
}

type Action int32

const (
	Action_EJECT   Action = 0
	Action_UNEJECT Action = 1
)

var Action_name = map[int32]string{
	0: "EJECT",
	1: "UNEJECT",
}

var Action_value = map[string]int32{
	"EJECT":   0,
	"UNEJECT": 1,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{1}
}

type OutlierDetectionEvent struct {
	Type                OutlierEjectionType   `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.data.cluster.v2alpha.OutlierEjectionType" json:"type,omitempty"`
	Timestamp           *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SecsSinceLastAction *wrappers.UInt64Value `protobuf:"bytes,3,opt,name=secs_since_last_action,json=secsSinceLastAction,proto3" json:"secs_since_last_action,omitempty"`
	ClusterName         string                `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	UpstreamUrl         string                `protobuf:"bytes,5,opt,name=upstream_url,json=upstreamUrl,proto3" json:"upstream_url,omitempty"`
	Action              Action                `protobuf:"varint,6,opt,name=action,proto3,enum=envoy.data.cluster.v2alpha.Action" json:"action,omitempty"`
	NumEjections        uint32                `protobuf:"varint,7,opt,name=num_ejections,json=numEjections,proto3" json:"num_ejections,omitempty"`
	Enforced            bool                  `protobuf:"varint,8,opt,name=enforced,proto3" json:"enforced,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*OutlierDetectionEvent_EjectSuccessRateEvent
	//	*OutlierDetectionEvent_EjectConsecutiveEvent
	Event                isOutlierDetectionEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *OutlierDetectionEvent) Reset()         { *m = OutlierDetectionEvent{} }
func (m *OutlierDetectionEvent) String() string { return proto.CompactTextString(m) }
func (*OutlierDetectionEvent) ProtoMessage()    {}
func (*OutlierDetectionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{0}
}

func (m *OutlierDetectionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetectionEvent.Unmarshal(m, b)
}
func (m *OutlierDetectionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetectionEvent.Marshal(b, m, deterministic)
}
func (m *OutlierDetectionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetectionEvent.Merge(m, src)
}
func (m *OutlierDetectionEvent) XXX_Size() int {
	return xxx_messageInfo_OutlierDetectionEvent.Size(m)
}
func (m *OutlierDetectionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetectionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetectionEvent proto.InternalMessageInfo

func (m *OutlierDetectionEvent) GetType() OutlierEjectionType {
	if m != nil {
		return m.Type
	}
	return OutlierEjectionType_CONSECUTIVE_5XX
}

func (m *OutlierDetectionEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *OutlierDetectionEvent) GetSecsSinceLastAction() *wrappers.UInt64Value {
	if m != nil {
		return m.SecsSinceLastAction
	}
	return nil
}

func (m *OutlierDetectionEvent) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *OutlierDetectionEvent) GetUpstreamUrl() string {
	if m != nil {
		return m.UpstreamUrl
	}
	return ""
}

func (m *OutlierDetectionEvent) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_EJECT
}

func (m *OutlierDetectionEvent) GetNumEjections() uint32 {
	if m != nil {
		return m.NumEjections
	}
	return 0
}

func (m *OutlierDetectionEvent) GetEnforced() bool {
	if m != nil {
		return m.Enforced
	}
	return false
}

type isOutlierDetectionEvent_Event interface {
	isOutlierDetectionEvent_Event()
}

type OutlierDetectionEvent_EjectSuccessRateEvent struct {
	EjectSuccessRateEvent *OutlierEjectSuccessRate `protobuf:"bytes,9,opt,name=eject_success_rate_event,json=ejectSuccessRateEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectConsecutiveEvent struct {
	EjectConsecutiveEvent *OutlierEjectConsecutive `protobuf:"bytes,10,opt,name=eject_consecutive_event,json=ejectConsecutiveEvent,proto3,oneof"`
}

func (*OutlierDetectionEvent_EjectSuccessRateEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectConsecutiveEvent) isOutlierDetectionEvent_Event() {}

func (m *OutlierDetectionEvent) GetEvent() isOutlierDetectionEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectSuccessRateEvent() *OutlierEjectSuccessRate {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectSuccessRateEvent); ok {
		return x.EjectSuccessRateEvent
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectConsecutiveEvent() *OutlierEjectConsecutive {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectConsecutiveEvent); ok {
		return x.EjectConsecutiveEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutlierDetectionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutlierDetectionEvent_EjectSuccessRateEvent)(nil),
		(*OutlierDetectionEvent_EjectConsecutiveEvent)(nil),
	}
}

type OutlierEjectSuccessRate struct {
	HostSuccessRate                     uint32   `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
	ClusterAverageSuccessRate           uint32   `protobuf:"varint,2,opt,name=cluster_average_success_rate,json=clusterAverageSuccessRate,proto3" json:"cluster_average_success_rate,omitempty"`
	ClusterSuccessRateEjectionThreshold uint32   `protobuf:"varint,3,opt,name=cluster_success_rate_ejection_threshold,json=clusterSuccessRateEjectionThreshold,proto3" json:"cluster_success_rate_ejection_threshold,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *OutlierEjectSuccessRate) Reset()         { *m = OutlierEjectSuccessRate{} }
func (m *OutlierEjectSuccessRate) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectSuccessRate) ProtoMessage()    {}
func (*OutlierEjectSuccessRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{1}
}

func (m *OutlierEjectSuccessRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectSuccessRate.Unmarshal(m, b)
}
func (m *OutlierEjectSuccessRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectSuccessRate.Marshal(b, m, deterministic)
}
func (m *OutlierEjectSuccessRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectSuccessRate.Merge(m, src)
}
func (m *OutlierEjectSuccessRate) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectSuccessRate.Size(m)
}
func (m *OutlierEjectSuccessRate) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectSuccessRate.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectSuccessRate proto.InternalMessageInfo

func (m *OutlierEjectSuccessRate) GetHostSuccessRate() uint32 {
	if m != nil {
		return m.HostSuccessRate
	}
	return 0
}

func (m *OutlierEjectSuccessRate) GetClusterAverageSuccessRate() uint32 {
	if m != nil {
		return m.ClusterAverageSuccessRate
	}
	return 0
}

func (m *OutlierEjectSuccessRate) GetClusterSuccessRateEjectionThreshold() uint32 {
	if m != nil {
		return m.ClusterSuccessRateEjectionThreshold
	}
	return 0
}

type OutlierEjectConsecutive struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutlierEjectConsecutive) Reset()         { *m = OutlierEjectConsecutive{} }
func (m *OutlierEjectConsecutive) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectConsecutive) ProtoMessage()    {}
func (*OutlierEjectConsecutive) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{2}
}

func (m *OutlierEjectConsecutive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectConsecutive.Unmarshal(m, b)
}
func (m *OutlierEjectConsecutive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectConsecutive.Marshal(b, m, deterministic)
}
func (m *OutlierEjectConsecutive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectConsecutive.Merge(m, src)
}
func (m *OutlierEjectConsecutive) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectConsecutive.Size(m)
}
func (m *OutlierEjectConsecutive) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectConsecutive.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectConsecutive proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.data.cluster.v2alpha.OutlierEjectionType", OutlierEjectionType_name, OutlierEjectionType_value)
	proto.RegisterEnum("envoy.data.cluster.v2alpha.Action", Action_name, Action_value)
	proto.RegisterType((*OutlierDetectionEvent)(nil), "envoy.data.cluster.v2alpha.OutlierDetectionEvent")
	proto.RegisterType((*OutlierEjectSuccessRate)(nil), "envoy.data.cluster.v2alpha.OutlierEjectSuccessRate")
	proto.RegisterType((*OutlierEjectConsecutive)(nil), "envoy.data.cluster.v2alpha.OutlierEjectConsecutive")
}

func init() {
	proto.RegisterFile("envoy/data/cluster/v2alpha/outlier_detection_event.proto", fileDescriptor_5e03c92c55863094)
}

var fileDescriptor_5e03c92c55863094 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x52, 0xda, 0x40,
	0x14, 0xc6, 0x0d, 0x02, 0xc2, 0x8a, 0x95, 0xae, 0x63, 0x8d, 0xd4, 0xd6, 0x0c, 0x76, 0xa6, 0x0c,
	0x17, 0xc9, 0x8c, 0xb6, 0x1d, 0x6f, 0x09, 0xa6, 0x4a, 0x87, 0x82, 0x0d, 0x60, 0xed, 0x45, 0x67,
	0x67, 0x0d, 0x47, 0xa1, 0x93, 0x7f, 0x93, 0xdd, 0xd0, 0x72, 0xdb, 0x47, 0xe9, 0x8b, 0xf4, 0xb1,
	0xda, 0xf1, 0xaa, 0xc3, 0x26, 0xc1, 0x60, 0xd5, 0x69, 0xef, 0x60, 0xf7, 0xfb, 0x7d, 0xe7, 0xcb,
	0xd9, 0xb3, 0x8b, 0x0e, 0xc1, 0x9d, 0x78, 0x53, 0x6d, 0x48, 0x39, 0xd5, 0x2c, 0x3b, 0x64, 0x1c,
	0x02, 0x6d, 0xb2, 0x4f, 0x6d, 0x7f, 0x44, 0x35, 0x2f, 0xe4, 0xf6, 0x18, 0x02, 0x32, 0x04, 0x0e,
	0x16, 0x1f, 0x7b, 0x2e, 0x81, 0x09, 0xb8, 0x5c, 0xf5, 0x03, 0x8f, 0x7b, 0xb8, 0x22, 0x48, 0x75,
	0x46, 0xaa, 0x31, 0xa9, 0xc6, 0x64, 0x65, 0xf7, 0xca, 0xf3, 0xae, 0x6c, 0xd0, 0x84, 0xf2, 0x22,
	0xbc, 0xd4, 0xf8, 0xd8, 0x01, 0xc6, 0xa9, 0xe3, 0x47, 0x70, 0xe5, 0xf9, 0x6d, 0xc1, 0xd7, 0x80,
	0xfa, 0x3e, 0x04, 0x2c, 0xde, 0xdf, 0x9a, 0x50, 0x7b, 0x3c, 0xa4, 0x1c, 0xb4, 0xe4, 0x47, 0xb4,
	0x51, 0xfd, 0x99, 0x43, 0x9b, 0xdd, 0x28, 0xd7, 0x51, 0x12, 0xcb, 0x98, 0xa5, 0xc2, 0xef, 0x51,
	0x96, 0x4f, 0x7d, 0x90, 0x25, 0x45, 0xaa, 0x3d, 0xda, 0xd7, 0xd4, 0xfb, 0xe3, 0xa9, 0xb1, 0x81,
	0xf1, 0x25, 0xe2, 0xfb, 0x53, 0x1f, 0xf4, 0xc2, 0xb5, 0x9e, 0xfb, 0x2e, 0x65, 0xca, 0x92, 0x29,
	0x6c, 0xf0, 0x21, 0x2a, 0xce, 0x43, 0xcb, 0x19, 0x45, 0xaa, 0xad, 0xee, 0x57, 0xd4, 0x28, 0xb5,
	0x9a, 0xa4, 0x56, 0xfb, 0x89, 0xc2, 0xbc, 0x11, 0xe3, 0x0f, 0xe8, 0x09, 0x03, 0x8b, 0x11, 0x36,
	0x76, 0x2d, 0x20, 0x36, 0x65, 0x9c, 0x50, 0x51, 0x47, 0x5e, 0x16, 0x36, 0x3b, 0x7f, 0xd9, 0x0c,
	0x5a, 0x2e, 0x7f, 0xf3, 0xea, 0x8c, 0xda, 0x21, 0x98, 0x1b, 0x33, 0xb6, 0x37, 0x43, 0xdb, 0x94,
	0xf1, 0x86, 0x00, 0x71, 0x1d, 0x95, 0xe2, 0x6f, 0x20, 0x2e, 0x75, 0x40, 0xce, 0x2a, 0x52, 0xad,
	0xa8, 0xaf, 0x5c, 0xeb, 0xd9, 0x20, 0xa3, 0x48, 0xe6, 0x6a, 0xbc, 0xd9, 0xa1, 0x0e, 0xcc, 0xb4,
	0xa1, 0xcf, 0x78, 0x00, 0xd4, 0x21, 0x61, 0x60, 0xcb, 0xb9, 0x5b, 0xda, 0x64, 0x73, 0x10, 0xd8,
	0xf8, 0x08, 0xe5, 0xe3, 0x68, 0x79, 0xd1, 0xb5, 0xea, 0x43, 0x5d, 0x8b, 0xb2, 0xa4, 0x1a, 0x15,
	0xb3, 0x78, 0x0f, 0xad, 0xb9, 0xa1, 0x43, 0x20, 0x6e, 0x27, 0x93, 0x57, 0x14, 0xa9, 0xb6, 0x66,
	0x96, 0xdc, 0xd0, 0x49, 0x5a, 0xcc, 0x70, 0x05, 0x15, 0xc0, 0xbd, 0xf4, 0x02, 0x0b, 0x86, 0x72,
	0x41, 0x91, 0x6a, 0x05, 0x73, 0xfe, 0x1f, 0xbb, 0x48, 0x16, 0x30, 0x61, 0xa1, 0x65, 0x01, 0x63,
	0x24, 0xa0, 0x1c, 0xa2, 0x61, 0x93, 0x8b, 0xa2, 0x67, 0x07, 0xff, 0x7a, 0x9c, 0xbd, 0xc8, 0xc1,
	0xa4, 0x1c, 0x4e, 0x96, 0xcc, 0x4d, 0xb8, 0xb5, 0x16, 0x8d, 0x8a, 0x83, 0xb6, 0xa2, 0x7a, 0x96,
	0xe7, 0x32, 0xb0, 0x42, 0x3e, 0x9e, 0x24, 0xe5, 0xd0, 0xff, 0x95, 0x6b, 0xde, 0x18, 0xcc, 0xcb,
	0xa5, 0xd6, 0x44, 0x39, 0xbd, 0x84, 0x72, 0xc2, 0x1c, 0x2f, 0xff, 0xd6, 0xa5, 0xea, 0x2f, 0x09,
	0x6d, 0xdd, 0x93, 0x18, 0x1f, 0xa0, 0xc7, 0x23, 0x8f, 0x2d, 0xf6, 0x41, 0x0c, 0xf4, 0x9a, 0x38,
	0xc0, 0x7a, 0x46, 0x1e, 0x9a, 0xeb, 0x33, 0x45, 0x1a, 0x3a, 0x41, 0x3b, 0xc9, 0x70, 0xd0, 0x09,
	0x04, 0xf4, 0x0a, 0x16, 0xf9, 0xcc, 0x22, 0xbf, 0x1d, 0x8b, 0x1b, 0x91, 0x36, 0xed, 0xf4, 0x19,
	0xbd, 0x4c, 0x9c, 0x16, 0x4f, 0x22, 0x3e, 0x45, 0xc2, 0x47, 0x01, 0xb0, 0x91, 0x67, 0x0f, 0xc5,
	0x28, 0xa7, 0x4c, 0xf7, 0x62, 0x2e, 0xdd, 0xea, 0xe4, 0x76, 0x25, 0x4c, 0x75, 0x7b, 0xf1, 0xc3,
	0x53, 0x7d, 0xaa, 0xff, 0x90, 0xd0, 0xc6, 0x1d, 0xb7, 0x12, 0x6f, 0xa0, 0xf5, 0x66, 0xb7, 0xd3,
	0x33, 0x9a, 0x83, 0x7e, 0xeb, 0xcc, 0x20, 0xaf, 0xcf, 0xcf, 0xcb, 0x4b, 0x78, 0x17, 0x3d, 0x4d,
	0x2f, 0x1e, 0x37, 0xfa, 0xc6, 0xc7, 0xc6, 0x27, 0xf2, 0xb6, 0xd1, 0x6a, 0x0f, 0x4c, 0xa3, 0x2c,
	0xe1, 0x32, 0x2a, 0xf5, 0x06, 0xcd, 0xa6, 0xd1, 0xeb, 0x11, 0xb3, 0xd1, 0x37, 0xca, 0x19, 0xfc,
	0x02, 0x29, 0x69, 0xa4, 0xdd, 0x6d, 0x36, 0xda, 0xa4, 0x6b, 0xb6, 0x8e, 0x5b, 0x9d, 0x39, 0xb7,
	0x8c, 0x9f, 0xa1, 0xed, 0x34, 0xb7, 0x20, 0x2b, 0x67, 0xeb, 0x0a, 0xca, 0xc7, 0xf7, 0xb1, 0x88,
	0x72, 0xc6, 0x3b, 0xa3, 0xd9, 0x2f, 0x2f, 0xe1, 0x55, 0xb4, 0x32, 0xe8, 0x44, 0x7f, 0x24, 0xfd,
	0x04, 0xd5, 0xc6, 0x5e, 0x34, 0x3b, 0x7e, 0xe0, 0x7d, 0x9b, 0x3e, 0x30, 0x46, 0x7a, 0xe5, 0xce,
	0x67, 0xec, 0x74, 0xf6, 0x26, 0x9c, 0x4a, 0x17, 0x79, 0xf1, 0x38, 0x1c, 0xfc, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xcd, 0x97, 0x62, 0xc7, 0xa0, 0x05, 0x00, 0x00,
}
