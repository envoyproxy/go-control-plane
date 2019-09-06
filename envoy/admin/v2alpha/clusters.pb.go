// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v2alpha/clusters.proto

package envoy_admin_v2alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/v2/envoy/api/v2/core"
	_type "github.com/envoyproxy/go-control-plane/v2/envoy/type"
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

type Clusters struct {
	ClusterStatuses      []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Clusters) Reset()         { *m = Clusters{} }
func (m *Clusters) String() string { return proto.CompactTextString(m) }
func (*Clusters) ProtoMessage()    {}
func (*Clusters) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{0}
}

func (m *Clusters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Clusters.Unmarshal(m, b)
}
func (m *Clusters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Clusters.Marshal(b, m, deterministic)
}
func (m *Clusters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Clusters.Merge(m, src)
}
func (m *Clusters) XXX_Size() int {
	return xxx_messageInfo_Clusters.Size(m)
}
func (m *Clusters) XXX_DiscardUnknown() {
	xxx_messageInfo_Clusters.DiscardUnknown(m)
}

var xxx_messageInfo_Clusters proto.InternalMessageInfo

func (m *Clusters) GetClusterStatuses() []*ClusterStatus {
	if m != nil {
		return m.ClusterStatuses
	}
	return nil
}

type ClusterStatus struct {
	Name                                    string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	AddedViaApi                             bool           `protobuf:"varint,2,opt,name=added_via_api,json=addedViaApi,proto3" json:"added_via_api,omitempty"`
	SuccessRateEjectionThreshold            *_type.Percent `protobuf:"bytes,3,opt,name=success_rate_ejection_threshold,json=successRateEjectionThreshold,proto3" json:"success_rate_ejection_threshold,omitempty"`
	HostStatuses                            []*HostStatus  `protobuf:"bytes,4,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
	LocalOriginSuccessRateEjectionThreshold *_type.Percent `protobuf:"bytes,5,opt,name=local_origin_success_rate_ejection_threshold,json=localOriginSuccessRateEjectionThreshold,proto3" json:"local_origin_success_rate_ejection_threshold,omitempty"`
	XXX_NoUnkeyedLiteral                    struct{}       `json:"-"`
	XXX_unrecognized                        []byte         `json:"-"`
	XXX_sizecache                           int32          `json:"-"`
}

func (m *ClusterStatus) Reset()         { *m = ClusterStatus{} }
func (m *ClusterStatus) String() string { return proto.CompactTextString(m) }
func (*ClusterStatus) ProtoMessage()    {}
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{1}
}

func (m *ClusterStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterStatus.Unmarshal(m, b)
}
func (m *ClusterStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterStatus.Marshal(b, m, deterministic)
}
func (m *ClusterStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterStatus.Merge(m, src)
}
func (m *ClusterStatus) XXX_Size() int {
	return xxx_messageInfo_ClusterStatus.Size(m)
}
func (m *ClusterStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterStatus proto.InternalMessageInfo

func (m *ClusterStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ClusterStatus) GetAddedViaApi() bool {
	if m != nil {
		return m.AddedViaApi
	}
	return false
}

func (m *ClusterStatus) GetSuccessRateEjectionThreshold() *_type.Percent {
	if m != nil {
		return m.SuccessRateEjectionThreshold
	}
	return nil
}

func (m *ClusterStatus) GetHostStatuses() []*HostStatus {
	if m != nil {
		return m.HostStatuses
	}
	return nil
}

func (m *ClusterStatus) GetLocalOriginSuccessRateEjectionThreshold() *_type.Percent {
	if m != nil {
		return m.LocalOriginSuccessRateEjectionThreshold
	}
	return nil
}

type HostStatus struct {
	Address                *core.Address     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stats                  []*SimpleMetric   `protobuf:"bytes,2,rep,name=stats,proto3" json:"stats,omitempty"`
	HealthStatus           *HostHealthStatus `protobuf:"bytes,3,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	SuccessRate            *_type.Percent    `protobuf:"bytes,4,opt,name=success_rate,json=successRate,proto3" json:"success_rate,omitempty"`
	Weight                 uint32            `protobuf:"varint,5,opt,name=weight,proto3" json:"weight,omitempty"`
	Hostname               string            `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Priority               uint32            `protobuf:"varint,7,opt,name=priority,proto3" json:"priority,omitempty"`
	LocalOriginSuccessRate *_type.Percent    `protobuf:"bytes,8,opt,name=local_origin_success_rate,json=localOriginSuccessRate,proto3" json:"local_origin_success_rate,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *HostStatus) Reset()         { *m = HostStatus{} }
func (m *HostStatus) String() string { return proto.CompactTextString(m) }
func (*HostStatus) ProtoMessage()    {}
func (*HostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{2}
}

func (m *HostStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostStatus.Unmarshal(m, b)
}
func (m *HostStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostStatus.Marshal(b, m, deterministic)
}
func (m *HostStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostStatus.Merge(m, src)
}
func (m *HostStatus) XXX_Size() int {
	return xxx_messageInfo_HostStatus.Size(m)
}
func (m *HostStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostStatus proto.InternalMessageInfo

func (m *HostStatus) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *HostStatus) GetStats() []*SimpleMetric {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *HostStatus) GetHealthStatus() *HostHealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return nil
}

func (m *HostStatus) GetSuccessRate() *_type.Percent {
	if m != nil {
		return m.SuccessRate
	}
	return nil
}

func (m *HostStatus) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *HostStatus) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *HostStatus) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *HostStatus) GetLocalOriginSuccessRate() *_type.Percent {
	if m != nil {
		return m.LocalOriginSuccessRate
	}
	return nil
}

type HostHealthStatus struct {
	FailedActiveHealthCheck   bool              `protobuf:"varint,1,opt,name=failed_active_health_check,json=failedActiveHealthCheck,proto3" json:"failed_active_health_check,omitempty"`
	FailedOutlierCheck        bool              `protobuf:"varint,2,opt,name=failed_outlier_check,json=failedOutlierCheck,proto3" json:"failed_outlier_check,omitempty"`
	FailedActiveDegradedCheck bool              `protobuf:"varint,4,opt,name=failed_active_degraded_check,json=failedActiveDegradedCheck,proto3" json:"failed_active_degraded_check,omitempty"`
	PendingDynamicRemoval     bool              `protobuf:"varint,5,opt,name=pending_dynamic_removal,json=pendingDynamicRemoval,proto3" json:"pending_dynamic_removal,omitempty"`
	PendingActiveHc           bool              `protobuf:"varint,6,opt,name=pending_active_hc,json=pendingActiveHc,proto3" json:"pending_active_hc,omitempty"`
	EdsHealthStatus           core.HealthStatus `protobuf:"varint,3,opt,name=eds_health_status,json=edsHealthStatus,proto3,enum=envoy.api.v2.core.HealthStatus" json:"eds_health_status,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}          `json:"-"`
	XXX_unrecognized          []byte            `json:"-"`
	XXX_sizecache             int32             `json:"-"`
}

func (m *HostHealthStatus) Reset()         { *m = HostHealthStatus{} }
func (m *HostHealthStatus) String() string { return proto.CompactTextString(m) }
func (*HostHealthStatus) ProtoMessage()    {}
func (*HostHealthStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6251a3a957f478b, []int{3}
}

func (m *HostHealthStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HostHealthStatus.Unmarshal(m, b)
}
func (m *HostHealthStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HostHealthStatus.Marshal(b, m, deterministic)
}
func (m *HostHealthStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostHealthStatus.Merge(m, src)
}
func (m *HostHealthStatus) XXX_Size() int {
	return xxx_messageInfo_HostHealthStatus.Size(m)
}
func (m *HostHealthStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HostHealthStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HostHealthStatus proto.InternalMessageInfo

func (m *HostHealthStatus) GetFailedActiveHealthCheck() bool {
	if m != nil {
		return m.FailedActiveHealthCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedOutlierCheck() bool {
	if m != nil {
		return m.FailedOutlierCheck
	}
	return false
}

func (m *HostHealthStatus) GetFailedActiveDegradedCheck() bool {
	if m != nil {
		return m.FailedActiveDegradedCheck
	}
	return false
}

func (m *HostHealthStatus) GetPendingDynamicRemoval() bool {
	if m != nil {
		return m.PendingDynamicRemoval
	}
	return false
}

func (m *HostHealthStatus) GetPendingActiveHc() bool {
	if m != nil {
		return m.PendingActiveHc
	}
	return false
}

func (m *HostHealthStatus) GetEdsHealthStatus() core.HealthStatus {
	if m != nil {
		return m.EdsHealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func init() {
	proto.RegisterType((*Clusters)(nil), "envoy.admin.v2alpha.Clusters")
	proto.RegisterType((*ClusterStatus)(nil), "envoy.admin.v2alpha.ClusterStatus")
	proto.RegisterType((*HostStatus)(nil), "envoy.admin.v2alpha.HostStatus")
	proto.RegisterType((*HostHealthStatus)(nil), "envoy.admin.v2alpha.HostHealthStatus")
}

func init() { proto.RegisterFile("envoy/admin/v2alpha/clusters.proto", fileDescriptor_c6251a3a957f478b) }

var fileDescriptor_c6251a3a957f478b = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xef, 0x6e, 0xd3, 0x3e,
	0x14, 0x55, 0xbb, 0x6e, 0xeb, 0xcf, 0x5d, 0x7f, 0xdb, 0x3c, 0xd8, 0xb2, 0x6a, 0xd2, 0xba, 0x08,
	0x44, 0x85, 0x50, 0x82, 0xca, 0x34, 0x3e, 0xf0, 0x01, 0xed, 0x0f, 0xd2, 0x04, 0x1a, 0x9b, 0x32,
	0x84, 0x04, 0x5f, 0x2c, 0xe3, 0x5c, 0x1a, 0x43, 0x1a, 0x47, 0xb6, 0x1b, 0xe8, 0xd3, 0xf0, 0x64,
	0x3c, 0x05, 0x2f, 0x80, 0x62, 0x3b, 0x5d, 0xc6, 0x5a, 0xf8, 0x96, 0x9b, 0x7b, 0x4e, 0x72, 0xce,
	0xbd, 0xc7, 0x46, 0x3e, 0x64, 0x85, 0x98, 0x86, 0x34, 0x1e, 0xf3, 0x2c, 0x2c, 0x86, 0x34, 0xcd,
	0x13, 0x1a, 0xb2, 0x74, 0xa2, 0x34, 0x48, 0x15, 0xe4, 0x52, 0x68, 0x81, 0xb7, 0x0c, 0x26, 0x30,
	0x98, 0xc0, 0x61, 0x7a, 0x07, 0xf3, 0x88, 0x63, 0xd0, 0x92, 0x33, 0xc7, 0xeb, 0xed, 0x3b, 0x48,
	0xce, 0xc3, 0x62, 0x18, 0x32, 0x21, 0x21, 0xa4, 0x71, 0x2c, 0x41, 0x55, 0x80, 0x07, 0x77, 0x01,
	0x09, 0xd0, 0x54, 0x27, 0x84, 0x25, 0xc0, 0xbe, 0x3a, 0x94, 0x67, 0x51, 0x7a, 0x9a, 0x43, 0x98,
	0x83, 0x64, 0x90, 0x69, 0xdb, 0xf1, 0x3f, 0xa0, 0xf6, 0xa9, 0x93, 0x8a, 0x2f, 0xd0, 0x86, 0x93,
	0x4d, 0x94, 0xa6, 0x7a, 0xa2, 0x40, 0x79, 0x8d, 0xfe, 0xd2, 0xa0, 0x33, 0xf4, 0x83, 0x39, 0xfa,
	0x03, 0x47, 0xbc, 0x36, 0xd8, 0x68, 0x9d, 0xd5, 0x4b, 0x50, 0xfe, 0xcf, 0x26, 0xea, 0xde, 0x82,
	0x60, 0x8c, 0x5a, 0x19, 0x1d, 0x83, 0xd7, 0xe8, 0x37, 0x06, 0xff, 0x45, 0xe6, 0x19, 0xfb, 0xa8,
	0x4b, 0xe3, 0x18, 0x62, 0x52, 0x70, 0x4a, 0x68, 0xce, 0xbd, 0x66, 0xbf, 0x31, 0x68, 0x47, 0x1d,
	0xf3, 0xf2, 0x3d, 0xa7, 0xc7, 0x39, 0xc7, 0x1f, 0xd1, 0xbe, 0x9a, 0x30, 0x06, 0x4a, 0x11, 0x49,
	0x35, 0x10, 0xf8, 0x02, 0x4c, 0x73, 0x91, 0x11, 0x9d, 0x48, 0x50, 0x89, 0x48, 0x63, 0x6f, 0xa9,
	0xdf, 0x18, 0x74, 0x86, 0x5b, 0x4e, 0x67, 0x69, 0x34, 0xb8, 0xb2, 0x46, 0xa3, 0x3d, 0xc7, 0x8d,
	0xa8, 0x86, 0x57, 0x8e, 0xf9, 0xae, 0x22, 0xe2, 0x33, 0xd4, 0x4d, 0x84, 0xd2, 0x37, 0x8e, 0x5b,
	0xc6, 0xf1, 0xfe, 0x5c, 0xc7, 0xe7, 0x42, 0x69, 0x67, 0x77, 0x2d, 0x99, 0x3d, 0x83, 0xc2, 0x12,
	0x3d, 0x49, 0x05, 0xa3, 0x29, 0x11, 0x92, 0x8f, 0x78, 0x46, 0xfe, 0x25, 0x77, 0x79, 0xb1, 0xdc,
	0x47, 0xe6, 0x43, 0x97, 0xe6, 0x3b, 0xd7, 0x7f, 0x51, 0xee, 0xff, 0x58, 0x42, 0xe8, 0x46, 0x10,
	0x3e, 0x44, 0xab, 0x2e, 0x1a, 0x66, 0xbe, 0x9d, 0x61, 0xaf, 0xb2, 0x90, 0xf3, 0xa0, 0x18, 0x06,
	0x65, 0x36, 0x82, 0x63, 0x8b, 0x88, 0x2a, 0x28, 0x7e, 0x8e, 0x96, 0x4b, 0xe7, 0xca, 0x6b, 0x1a,
	0xdb, 0x07, 0x73, 0x6d, 0x5f, 0xf3, 0x71, 0x9e, 0xc2, 0x85, 0x49, 0x66, 0x64, 0xf1, 0xf8, 0x35,
	0xea, 0xba, 0xa0, 0xd9, 0xc9, 0xb9, 0x0d, 0x3c, 0x5c, 0x38, 0xb7, 0x73, 0x83, 0x9e, 0x4d, 0xaf,
	0x56, 0xe1, 0x23, 0xb4, 0x56, 0x1f, 0x98, 0xd7, 0x5a, 0x3c, 0x9d, 0x4e, 0x6d, 0x99, 0x78, 0x1b,
	0xad, 0x7c, 0x03, 0x3e, 0x4a, 0xb4, 0x99, 0x67, 0x37, 0x72, 0x15, 0xee, 0xa1, 0x76, 0xb9, 0x1d,
	0x93, 0xb5, 0x15, 0x93, 0xb5, 0x59, 0x5d, 0xf6, 0x72, 0xc9, 0x85, 0xe4, 0x7a, 0xea, 0xad, 0x1a,
	0xd6, 0xac, 0xc6, 0x6f, 0xd1, 0xee, 0xc2, 0x2d, 0x7a, 0xed, 0xc5, 0xa2, 0xb6, 0xe7, 0xaf, 0xcc,
	0xff, 0xd5, 0x44, 0x1b, 0x7f, 0x5a, 0xc7, 0x2f, 0x50, 0xef, 0x33, 0xe5, 0x29, 0xc4, 0x84, 0x32,
	0xcd, 0x0b, 0x20, 0xf5, 0xf3, 0x6a, 0x56, 0xd7, 0x8e, 0x76, 0x2c, 0xe2, 0xd8, 0x00, 0x2c, 0xfb,
	0xb4, 0x6c, 0xe3, 0xa7, 0xe8, 0x9e, 0x23, 0x8b, 0x89, 0x4e, 0x39, 0x48, 0x47, 0xb3, 0x87, 0x06,
	0xdb, 0xde, 0xa5, 0x6d, 0x59, 0xc6, 0x4b, 0xb4, 0x77, 0xfb, 0x77, 0x31, 0x8c, 0x24, 0x2d, 0x0f,
	0x9c, 0x65, 0xb6, 0x0c, 0x73, 0xb7, 0xfe, 0xc3, 0x33, 0x87, 0xb0, 0x1f, 0x38, 0x42, 0x3b, 0x39,
	0x64, 0x31, 0xcf, 0x46, 0x24, 0x9e, 0x66, 0x74, 0xcc, 0x19, 0x91, 0x30, 0x16, 0x05, 0x4d, 0xcd,
	0xd4, 0xdb, 0xd1, 0x7d, 0xd7, 0x3e, 0xb3, 0xdd, 0xc8, 0x36, 0xf1, 0x63, 0xb4, 0x59, 0xf1, 0x2a,
	0xa3, 0xcc, 0x6c, 0xa3, 0x1d, 0xad, 0xbb, 0x86, 0xf3, 0xc7, 0xf0, 0x1b, 0xb4, 0x09, 0xb1, 0x22,
	0x77, 0x03, 0xf5, 0xff, 0xcd, 0x41, 0xac, 0xa5, 0xf8, 0x56, 0x94, 0xd6, 0x21, 0x56, 0xf5, 0x17,
	0x27, 0x87, 0xe8, 0x80, 0x0b, 0xcb, 0xca, 0xa5, 0xf8, 0x3e, 0x9d, 0x97, 0xc8, 0x93, 0xea, 0x66,
	0x52, 0x57, 0xe5, 0x35, 0x78, 0xd5, 0xf8, 0xb4, 0x62, 0xee, 0xc3, 0x67, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x95, 0x47, 0x7a, 0xc5, 0xce, 0x05, 0x00, 0x00,
}
