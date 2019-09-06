// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/discovery.proto

package envoy_api_v2

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/v2/envoy/api/v2/core"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type DiscoveryRequest struct {
	VersionInfo          string         `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Node                 *core.Node     `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	ResourceNames        []string       `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	TypeUrl              string         `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResponseNonce        string         `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail          *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{0}
}

func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(m, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

func (m *DiscoveryRequest) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DiscoveryRequest) GetResourceNames() []string {
	if m != nil {
		return m.ResourceNames
	}
	return nil
}

func (m *DiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	VersionInfo          string             `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Resources            []*any.Any         `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	Canary               bool               `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	TypeUrl              string             `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Nonce                string             `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ControlPlane         *core.ControlPlane `protobuf:"bytes,6,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{1}
}

func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(m, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *DiscoveryResponse) GetResources() []*any.Any {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DiscoveryResponse) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

func (m *DiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *DiscoveryResponse) GetControlPlane() *core.ControlPlane {
	if m != nil {
		return m.ControlPlane
	}
	return nil
}

type DeltaDiscoveryRequest struct {
	Node                     *core.Node        `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	TypeUrl                  string            `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResourceNamesSubscribe   []string          `protobuf:"bytes,3,rep,name=resource_names_subscribe,json=resourceNamesSubscribe,proto3" json:"resource_names_subscribe,omitempty"`
	ResourceNamesUnsubscribe []string          `protobuf:"bytes,4,rep,name=resource_names_unsubscribe,json=resourceNamesUnsubscribe,proto3" json:"resource_names_unsubscribe,omitempty"`
	InitialResourceVersions  map[string]string `protobuf:"bytes,5,rep,name=initial_resource_versions,json=initialResourceVersions,proto3" json:"initial_resource_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponseNonce            string            `protobuf:"bytes,6,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail              *status.Status    `protobuf:"bytes,7,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}          `json:"-"`
	XXX_unrecognized         []byte            `json:"-"`
	XXX_sizecache            int32             `json:"-"`
}

func (m *DeltaDiscoveryRequest) Reset()         { *m = DeltaDiscoveryRequest{} }
func (m *DeltaDiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryRequest) ProtoMessage()    {}
func (*DeltaDiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{2}
}

func (m *DeltaDiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryRequest.Unmarshal(m, b)
}
func (m *DeltaDiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryRequest.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryRequest.Merge(m, src)
}
func (m *DeltaDiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryRequest.Size(m)
}
func (m *DeltaDiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryRequest proto.InternalMessageInfo

func (m *DeltaDiscoveryRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetResourceNamesSubscribe() []string {
	if m != nil {
		return m.ResourceNamesSubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResourceNamesUnsubscribe() []string {
	if m != nil {
		return m.ResourceNamesUnsubscribe
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetInitialResourceVersions() map[string]string {
	if m != nil {
		return m.InitialResourceVersions
	}
	return nil
}

func (m *DeltaDiscoveryRequest) GetResponseNonce() string {
	if m != nil {
		return m.ResponseNonce
	}
	return ""
}

func (m *DeltaDiscoveryRequest) GetErrorDetail() *status.Status {
	if m != nil {
		return m.ErrorDetail
	}
	return nil
}

type DeltaDiscoveryResponse struct {
	SystemVersionInfo    string      `protobuf:"bytes,1,opt,name=system_version_info,json=systemVersionInfo,proto3" json:"system_version_info,omitempty"`
	Resources            []*Resource `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	TypeUrl              string      `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	RemovedResources     []string    `protobuf:"bytes,6,rep,name=removed_resources,json=removedResources,proto3" json:"removed_resources,omitempty"`
	Nonce                string      `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeltaDiscoveryResponse) Reset()         { *m = DeltaDiscoveryResponse{} }
func (m *DeltaDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DeltaDiscoveryResponse) ProtoMessage()    {}
func (*DeltaDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{3}
}

func (m *DeltaDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeltaDiscoveryResponse.Unmarshal(m, b)
}
func (m *DeltaDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeltaDiscoveryResponse.Marshal(b, m, deterministic)
}
func (m *DeltaDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeltaDiscoveryResponse.Merge(m, src)
}
func (m *DeltaDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DeltaDiscoveryResponse.Size(m)
}
func (m *DeltaDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeltaDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeltaDiscoveryResponse proto.InternalMessageInfo

func (m *DeltaDiscoveryResponse) GetSystemVersionInfo() string {
	if m != nil {
		return m.SystemVersionInfo
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetResources() []*Resource {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetTypeUrl() string {
	if m != nil {
		return m.TypeUrl
	}
	return ""
}

func (m *DeltaDiscoveryResponse) GetRemovedResources() []string {
	if m != nil {
		return m.RemovedResources
	}
	return nil
}

func (m *DeltaDiscoveryResponse) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type Resource struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Aliases              []string `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Resource             *any.Any `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c7365e287e5c035, []int{4}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Resource) GetAliases() []string {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Resource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Resource) GetResource() *any.Any {
	if m != nil {
		return m.Resource
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoveryRequest)(nil), "envoy.api.v2.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "envoy.api.v2.DiscoveryResponse")
	proto.RegisterType((*DeltaDiscoveryRequest)(nil), "envoy.api.v2.DeltaDiscoveryRequest")
	proto.RegisterMapType((map[string]string)(nil), "envoy.api.v2.DeltaDiscoveryRequest.InitialResourceVersionsEntry")
	proto.RegisterType((*DeltaDiscoveryResponse)(nil), "envoy.api.v2.DeltaDiscoveryResponse")
	proto.RegisterType((*Resource)(nil), "envoy.api.v2.Resource")
}

func init() { proto.RegisterFile("envoy/api/v2/discovery.proto", fileDescriptor_2c7365e287e5c035) }

var fileDescriptor_2c7365e287e5c035 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x45, 0xb6, 0xe3, 0xd8, 0x63, 0x27, 0x24, 0xdb, 0xd4, 0x51, 0x4c, 0xa0, 0xae, 0xa1, 0x60,
	0x08, 0xc8, 0xc5, 0x6d, 0x21, 0x94, 0x1e, 0xda, 0xd4, 0x3d, 0xa4, 0x87, 0x10, 0x14, 0x92, 0xab,
	0x58, 0xcb, 0x93, 0x20, 0xaa, 0xec, 0xaa, 0xbb, 0x92, 0xa8, 0xa0, 0xa7, 0xd2, 0xcf, 0xec, 0x8f,
	0xf4, 0xd4, 0xa2, 0xd5, 0xae, 0x2d, 0x25, 0x22, 0xf8, 0xa6, 0xd9, 0x79, 0xfb, 0x76, 0xde, 0xcc,
	0x1b, 0xc1, 0x31, 0xb2, 0x94, 0x67, 0x53, 0x1a, 0x05, 0xd3, 0x74, 0x36, 0x5d, 0x06, 0xd2, 0xe7,
	0x29, 0x8a, 0xcc, 0x89, 0x04, 0x8f, 0x39, 0xe9, 0xab, 0xac, 0x43, 0xa3, 0xc0, 0x49, 0x67, 0xc3,
	0x2a, 0xd6, 0xe7, 0x02, 0xa7, 0x0b, 0x2a, 0xb1, 0xc0, 0x0e, 0x8f, 0xee, 0x38, 0xbf, 0x0b, 0x71,
	0xaa, 0xa2, 0x45, 0x72, 0x3b, 0xa5, 0x4c, 0xd3, 0x0c, 0x0f, 0x75, 0x4a, 0x44, 0xfe, 0x54, 0xc6,
	0x34, 0x4e, 0x64, 0x91, 0x18, 0xff, 0x6a, 0xc0, 0xde, 0xdc, 0xbc, 0xe9, 0xe2, 0xf7, 0x04, 0x65,
	0x4c, 0x5e, 0x42, 0x3f, 0x45, 0x21, 0x03, 0xce, 0xbc, 0x80, 0xdd, 0x72, 0xdb, 0x1a, 0x59, 0x93,
	0xae, 0xdb, 0xd3, 0x67, 0xe7, 0xec, 0x96, 0x93, 0x13, 0x68, 0x31, 0xbe, 0x44, 0xbb, 0x31, 0xb2,
	0x26, 0xbd, 0xd9, 0xa1, 0x53, 0x2e, 0xd3, 0xc9, 0x0b, 0x73, 0x2e, 0xf8, 0x12, 0x5d, 0x05, 0x22,
	0xaf, 0x60, 0x57, 0xa0, 0xe4, 0x89, 0xf0, 0xd1, 0x63, 0xf4, 0x1e, 0xa5, 0xdd, 0x1c, 0x35, 0x27,
	0x5d, 0x77, 0xc7, 0x9c, 0x5e, 0xe4, 0x87, 0xe4, 0x08, 0x3a, 0x71, 0x16, 0xa1, 0x97, 0x88, 0xd0,
	0x6e, 0xa9, 0x27, 0xb7, 0xf3, 0xf8, 0x5a, 0x84, 0x9a, 0x21, 0xe2, 0x4c, 0xa2, 0xc7, 0x38, 0xf3,
	0xd1, 0xde, 0x52, 0x80, 0x1d, 0x73, 0x7a, 0x91, 0x1f, 0x92, 0x77, 0xd0, 0x47, 0x21, 0xb8, 0xf0,
	0x96, 0x18, 0xd3, 0x20, 0xb4, 0xdb, 0xaa, 0x3a, 0xe2, 0x14, 0xea, 0x1d, 0x11, 0xf9, 0xce, 0x95,
	0x52, 0xef, 0xf6, 0x14, 0x6e, 0xae, 0x60, 0xe3, 0xbf, 0x16, 0xec, 0x97, 0x9a, 0x50, 0x30, 0x6e,
	0xd2, 0x85, 0x19, 0x74, 0x8d, 0x04, 0x69, 0x37, 0x46, 0xcd, 0x49, 0x6f, 0x76, 0x60, 0x1e, 0x33,
	0x53, 0x70, 0x3e, 0xb1, 0xcc, 0x5d, 0xc3, 0xc8, 0x00, 0xda, 0x3e, 0x65, 0x54, 0x64, 0x76, 0x73,
	0x64, 0x4d, 0x3a, 0xae, 0x8e, 0x9e, 0x52, 0x7f, 0x00, 0x5b, 0x65, 0xd1, 0x45, 0x40, 0xe6, 0xb0,
	0xe3, 0x73, 0x16, 0x0b, 0x1e, 0x7a, 0x51, 0x48, 0x19, 0x6a, 0xb5, 0x2f, 0x6a, 0x66, 0xf1, 0xb9,
	0xc0, 0x5d, 0xe6, 0x30, 0xb7, 0xef, 0x97, 0xa2, 0xf1, 0xbf, 0x26, 0x3c, 0x9f, 0x63, 0x18, 0xd3,
	0x47, 0x2e, 0x30, 0x23, 0xb6, 0x36, 0x19, 0x71, 0xb9, 0xfa, 0x46, 0xb5, 0xfa, 0x53, 0xb0, 0xab,
	0xd3, 0xf7, 0x64, 0xb2, 0x90, 0xbe, 0x08, 0x16, 0xa8, 0x7d, 0x30, 0xa8, 0xf8, 0xe0, 0xca, 0x64,
	0xc9, 0x07, 0x18, 0x3e, 0xb8, 0x99, 0xb0, 0xf5, 0xdd, 0x96, 0xba, 0x6b, 0x57, 0xee, 0x5e, 0xaf,
	0xf3, 0xe4, 0x27, 0x1c, 0x05, 0x2c, 0x88, 0x03, 0x1a, 0x7a, 0x2b, 0x16, 0x3d, 0x3c, 0x69, 0x6f,
	0xa9, 0x61, 0x7d, 0xac, 0x8a, 0xaa, 0xed, 0x83, 0x73, 0x5e, 0x90, 0xb8, 0x9a, 0xe3, 0x46, 0x53,
	0x7c, 0x61, 0xb1, 0xc8, 0xdc, 0xc3, 0xa0, 0x3e, 0x5b, 0xe3, 0xd8, 0xf6, 0x26, 0x8e, 0xdd, 0xde,
	0xc8, 0xb1, 0xc3, 0xaf, 0x70, 0xfc, 0x54, 0x59, 0x64, 0x0f, 0x9a, 0xdf, 0x30, 0xd3, 0x96, 0xcd,
	0x3f, 0x73, 0x0f, 0xa5, 0x34, 0x4c, 0x50, 0x4f, 0xa7, 0x08, 0xde, 0x37, 0x4e, 0xad, 0xf1, 0x1f,
	0x0b, 0x06, 0x0f, 0x95, 0xeb, 0x15, 0x70, 0xe0, 0x99, 0xcc, 0x64, 0x8c, 0xf7, 0x5e, 0xcd, 0x26,
	0xec, 0x17, 0xa9, 0x9b, 0xd2, 0x3e, 0xbc, 0x7d, 0xbc, 0x0f, 0x83, 0x6a, 0x8b, 0x4d, 0xb9, 0xe5,
	0x8d, 0x78, 0xc2, 0xf9, 0x27, 0xb0, 0x2f, 0xf0, 0x9e, 0xa7, 0xb8, 0xf4, 0xd6, 0xc4, 0x6d, 0x35,
	0xf8, 0x3d, 0x9d, 0x70, 0x57, 0x3c, 0xb5, 0x6b, 0x32, 0xfe, 0x6d, 0x41, 0xc7, 0x60, 0x08, 0x81,
	0x56, 0x6e, 0x24, 0xb5, 0x7a, 0x5d, 0x57, 0x7d, 0x13, 0x1b, 0xb6, 0x69, 0x18, 0x50, 0x89, 0x52,
	0x5b, 0xca, 0x84, 0x79, 0x46, 0xeb, 0xd6, 0x92, 0x4d, 0x48, 0x5e, 0x43, 0xc7, 0xd4, 0xa3, 0x7f,
	0x81, 0xf5, 0x7b, 0xbf, 0x42, 0x9d, 0x39, 0x30, 0x0c, 0x78, 0xd1, 0x8b, 0x48, 0xf0, 0x1f, 0x59,
	0xa5, 0x2d, 0x67, 0xbb, 0xab, 0xde, 0x5f, 0xe6, 0xd7, 0x2f, 0xad, 0x45, 0x5b, 0xf1, 0xbc, 0xf9,
	0x1f, 0x00, 0x00, 0xff, 0xff, 0x97, 0x59, 0x36, 0x09, 0x1f, 0x06, 0x00, 0x00,
}
