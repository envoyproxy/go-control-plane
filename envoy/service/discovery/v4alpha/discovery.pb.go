// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/service/discovery/v4alpha/discovery.proto

package envoy_service_discovery_v4alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v1 "github.com/cncf/udpa/go/udpa/core/v1"
	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/core/v4alpha"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DiscoveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionInfo   string         `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Node          *v4alpha.Node  `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	ResourceNames []string       `protobuf:"bytes,3,rep,name=resource_names,json=resourceNames,proto3" json:"resource_names,omitempty"`
	TypeUrl       string         `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResponseNonce string         `protobuf:"bytes,5,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail   *status.Status `protobuf:"bytes,6,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
}

func (x *DiscoveryRequest) Reset() {
	*x = DiscoveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryRequest) ProtoMessage() {}

func (x *DiscoveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryRequest.ProtoReflect.Descriptor instead.
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_discovery_v4alpha_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *DiscoveryRequest) GetVersionInfo() string {
	if x != nil {
		return x.VersionInfo
	}
	return ""
}

func (x *DiscoveryRequest) GetNode() *v4alpha.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DiscoveryRequest) GetResourceNames() []string {
	if x != nil {
		return x.ResourceNames
	}
	return nil
}

func (x *DiscoveryRequest) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DiscoveryRequest) GetResponseNonce() string {
	if x != nil {
		return x.ResponseNonce
	}
	return ""
}

func (x *DiscoveryRequest) GetErrorDetail() *status.Status {
	if x != nil {
		return x.ErrorDetail
	}
	return nil
}

type DiscoveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VersionInfo  string                `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	Resources    []*any.Any            `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	Canary       bool                  `protobuf:"varint,3,opt,name=canary,proto3" json:"canary,omitempty"`
	TypeUrl      string                `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Nonce        string                `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ControlPlane *v4alpha.ControlPlane `protobuf:"bytes,6,opt,name=control_plane,json=controlPlane,proto3" json:"control_plane,omitempty"`
}

func (x *DiscoveryResponse) Reset() {
	*x = DiscoveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoveryResponse) ProtoMessage() {}

func (x *DiscoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoveryResponse.ProtoReflect.Descriptor instead.
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_discovery_v4alpha_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *DiscoveryResponse) GetVersionInfo() string {
	if x != nil {
		return x.VersionInfo
	}
	return ""
}

func (x *DiscoveryResponse) GetResources() []*any.Any {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *DiscoveryResponse) GetCanary() bool {
	if x != nil {
		return x.Canary
	}
	return false
}

func (x *DiscoveryResponse) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DiscoveryResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *DiscoveryResponse) GetControlPlane() *v4alpha.ControlPlane {
	if x != nil {
		return x.ControlPlane
	}
	return nil
}

type DeltaDiscoveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node                     *v4alpha.Node         `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	TypeUrl                  string                `protobuf:"bytes,2,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	ResourceNamesSubscribe   []string              `protobuf:"bytes,3,rep,name=resource_names_subscribe,json=resourceNamesSubscribe,proto3" json:"resource_names_subscribe,omitempty"`
	UdpaResourcesSubscribe   []*v1.ResourceLocator `protobuf:"bytes,8,rep,name=udpa_resources_subscribe,json=udpaResourcesSubscribe,proto3" json:"udpa_resources_subscribe,omitempty"`
	ResourceNamesUnsubscribe []string              `protobuf:"bytes,4,rep,name=resource_names_unsubscribe,json=resourceNamesUnsubscribe,proto3" json:"resource_names_unsubscribe,omitempty"`
	UdpaResourcesUnsubscribe []*v1.ResourceLocator `protobuf:"bytes,9,rep,name=udpa_resources_unsubscribe,json=udpaResourcesUnsubscribe,proto3" json:"udpa_resources_unsubscribe,omitempty"`
	InitialResourceVersions  map[string]string     `protobuf:"bytes,5,rep,name=initial_resource_versions,json=initialResourceVersions,proto3" json:"initial_resource_versions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponseNonce            string                `protobuf:"bytes,6,opt,name=response_nonce,json=responseNonce,proto3" json:"response_nonce,omitempty"`
	ErrorDetail              *status.Status        `protobuf:"bytes,7,opt,name=error_detail,json=errorDetail,proto3" json:"error_detail,omitempty"`
}

func (x *DeltaDiscoveryRequest) Reset() {
	*x = DeltaDiscoveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaDiscoveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaDiscoveryRequest) ProtoMessage() {}

func (x *DeltaDiscoveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaDiscoveryRequest.ProtoReflect.Descriptor instead.
func (*DeltaDiscoveryRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_discovery_v4alpha_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *DeltaDiscoveryRequest) GetNode() *v4alpha.Node {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DeltaDiscoveryRequest) GetResourceNamesSubscribe() []string {
	if x != nil {
		return x.ResourceNamesSubscribe
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetUdpaResourcesSubscribe() []*v1.ResourceLocator {
	if x != nil {
		return x.UdpaResourcesSubscribe
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetResourceNamesUnsubscribe() []string {
	if x != nil {
		return x.ResourceNamesUnsubscribe
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetUdpaResourcesUnsubscribe() []*v1.ResourceLocator {
	if x != nil {
		return x.UdpaResourcesUnsubscribe
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetInitialResourceVersions() map[string]string {
	if x != nil {
		return x.InitialResourceVersions
	}
	return nil
}

func (x *DeltaDiscoveryRequest) GetResponseNonce() string {
	if x != nil {
		return x.ResponseNonce
	}
	return ""
}

func (x *DeltaDiscoveryRequest) GetErrorDetail() *status.Status {
	if x != nil {
		return x.ErrorDetail
	}
	return nil
}

type DeltaDiscoveryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SystemVersionInfo    string             `protobuf:"bytes,1,opt,name=system_version_info,json=systemVersionInfo,proto3" json:"system_version_info,omitempty"`
	Resources            []*Resource        `protobuf:"bytes,2,rep,name=resources,proto3" json:"resources,omitempty"`
	TypeUrl              string             `protobuf:"bytes,4,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	RemovedResources     []string           `protobuf:"bytes,6,rep,name=removed_resources,json=removedResources,proto3" json:"removed_resources,omitempty"`
	UdpaRemovedResources []*v1.ResourceName `protobuf:"bytes,7,rep,name=udpa_removed_resources,json=udpaRemovedResources,proto3" json:"udpa_removed_resources,omitempty"`
	Nonce                string             `protobuf:"bytes,5,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *DeltaDiscoveryResponse) Reset() {
	*x = DeltaDiscoveryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeltaDiscoveryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeltaDiscoveryResponse) ProtoMessage() {}

func (x *DeltaDiscoveryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeltaDiscoveryResponse.ProtoReflect.Descriptor instead.
func (*DeltaDiscoveryResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_discovery_v4alpha_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *DeltaDiscoveryResponse) GetSystemVersionInfo() string {
	if x != nil {
		return x.SystemVersionInfo
	}
	return ""
}

func (x *DeltaDiscoveryResponse) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *DeltaDiscoveryResponse) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *DeltaDiscoveryResponse) GetRemovedResources() []string {
	if x != nil {
		return x.RemovedResources
	}
	return nil
}

func (x *DeltaDiscoveryResponse) GetUdpaRemovedResources() []*v1.ResourceName {
	if x != nil {
		return x.UdpaRemovedResources
	}
	return nil
}

func (x *DeltaDiscoveryResponse) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to NameSpecifier:
	//	*Resource_Name
	//	*Resource_UdpaResourceName
	NameSpecifier isResource_NameSpecifier `protobuf_oneof:"name_specifier"`
	Aliases       []string                 `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
	Version       string                   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Resource      *any.Any                 `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_envoy_service_discovery_v4alpha_discovery_proto_rawDescGZIP(), []int{4}
}

func (m *Resource) GetNameSpecifier() isResource_NameSpecifier {
	if m != nil {
		return m.NameSpecifier
	}
	return nil
}

func (x *Resource) GetName() string {
	if x, ok := x.GetNameSpecifier().(*Resource_Name); ok {
		return x.Name
	}
	return ""
}

func (x *Resource) GetUdpaResourceName() *v1.ResourceName {
	if x, ok := x.GetNameSpecifier().(*Resource_UdpaResourceName); ok {
		return x.UdpaResourceName
	}
	return nil
}

func (x *Resource) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *Resource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Resource) GetResource() *any.Any {
	if x != nil {
		return x.Resource
	}
	return nil
}

type isResource_NameSpecifier interface {
	isResource_NameSpecifier()
}

type Resource_Name struct {
	Name string `protobuf:"bytes,3,opt,name=name,proto3,oneof"`
}

type Resource_UdpaResourceName struct {
	UdpaResourceName *v1.ResourceName `protobuf:"bytes,5,opt,name=udpa_resource_name,json=udpaResourceName,proto3,oneof"`
}

func (*Resource_Name) isResource_NameSpecifier() {}

func (*Resource_UdpaResourceName) isResource_NameSpecifier() {}

var File_envoy_service_discovery_v4alpha_discovery_proto protoreflect.FileDescriptor

var file_envoy_service_discovery_v4alpha_discovery_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x02, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f,
	0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x3a, 0x32, 0x9a, 0xc5, 0x88, 0x1e, 0x2d, 0x0a, 0x2b, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb6, 0x02, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x32, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x52, 0x0c, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x3a, 0x33, 0x9a, 0xc5, 0x88, 0x1e,
	0x2e, 0x0a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x8a, 0x06, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x34, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x18, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x16, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x57, 0x0a, 0x18, 0x75, 0x64, 0x70, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x16, 0x75, 0x64, 0x70, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x3c, 0x0a, 0x1a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x75,
	0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x55,
	0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x5b, 0x0a, 0x1a, 0x75, 0x64,
	0x70, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x18, 0x75,
	0x64, 0x70, 0x61, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x55, 0x6e, 0x73, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x8f, 0x01, 0x0a, 0x19, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x53, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65,
	0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x17, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4e, 0x6f, 0x6e, 0x63, 0x65,
	0x12, 0x35, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x4a, 0x0a, 0x1c, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfb, 0x02, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x47, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x16, 0x75, 0x64, 0x70, 0x61,
	0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x64, 0x70, 0x61, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x14, 0x75, 0x64, 0x70, 0x61, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x3a, 0x38, 0x9a, 0xc5, 0x88, 0x1e, 0x33, 0x0a, 0x31, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x33, 0x2e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x90, 0x02, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a,
	0x12, 0x75, 0x64, 0x70, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x75, 0x64, 0x70, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x10, 0x75, 0x64, 0x70, 0x61, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x3a,
	0x2a, 0x9a, 0xc5, 0x88, 0x1e, 0x25, 0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x49, 0x0a,
	0x2d, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x34, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x03, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_discovery_v4alpha_discovery_proto_rawDescOnce sync.Once
	file_envoy_service_discovery_v4alpha_discovery_proto_rawDescData = file_envoy_service_discovery_v4alpha_discovery_proto_rawDesc
)

func file_envoy_service_discovery_v4alpha_discovery_proto_rawDescGZIP() []byte {
	file_envoy_service_discovery_v4alpha_discovery_proto_rawDescOnce.Do(func() {
		file_envoy_service_discovery_v4alpha_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_discovery_v4alpha_discovery_proto_rawDescData)
	})
	return file_envoy_service_discovery_v4alpha_discovery_proto_rawDescData
}

var file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_envoy_service_discovery_v4alpha_discovery_proto_goTypes = []interface{}{
	(*DiscoveryRequest)(nil),       // 0: envoy.service.discovery.v4alpha.DiscoveryRequest
	(*DiscoveryResponse)(nil),      // 1: envoy.service.discovery.v4alpha.DiscoveryResponse
	(*DeltaDiscoveryRequest)(nil),  // 2: envoy.service.discovery.v4alpha.DeltaDiscoveryRequest
	(*DeltaDiscoveryResponse)(nil), // 3: envoy.service.discovery.v4alpha.DeltaDiscoveryResponse
	(*Resource)(nil),               // 4: envoy.service.discovery.v4alpha.Resource
	nil,                            // 5: envoy.service.discovery.v4alpha.DeltaDiscoveryRequest.InitialResourceVersionsEntry
	(*v4alpha.Node)(nil),           // 6: envoy.config.core.v4alpha.Node
	(*status.Status)(nil),          // 7: google.rpc.Status
	(*any.Any)(nil),                // 8: google.protobuf.Any
	(*v4alpha.ControlPlane)(nil),   // 9: envoy.config.core.v4alpha.ControlPlane
	(*v1.ResourceLocator)(nil),     // 10: udpa.core.v1.ResourceLocator
	(*v1.ResourceName)(nil),        // 11: udpa.core.v1.ResourceName
}
var file_envoy_service_discovery_v4alpha_discovery_proto_depIdxs = []int32{
	6,  // 0: envoy.service.discovery.v4alpha.DiscoveryRequest.node:type_name -> envoy.config.core.v4alpha.Node
	7,  // 1: envoy.service.discovery.v4alpha.DiscoveryRequest.error_detail:type_name -> google.rpc.Status
	8,  // 2: envoy.service.discovery.v4alpha.DiscoveryResponse.resources:type_name -> google.protobuf.Any
	9,  // 3: envoy.service.discovery.v4alpha.DiscoveryResponse.control_plane:type_name -> envoy.config.core.v4alpha.ControlPlane
	6,  // 4: envoy.service.discovery.v4alpha.DeltaDiscoveryRequest.node:type_name -> envoy.config.core.v4alpha.Node
	10, // 5: envoy.service.discovery.v4alpha.DeltaDiscoveryRequest.udpa_resources_subscribe:type_name -> udpa.core.v1.ResourceLocator
	10, // 6: envoy.service.discovery.v4alpha.DeltaDiscoveryRequest.udpa_resources_unsubscribe:type_name -> udpa.core.v1.ResourceLocator
	5,  // 7: envoy.service.discovery.v4alpha.DeltaDiscoveryRequest.initial_resource_versions:type_name -> envoy.service.discovery.v4alpha.DeltaDiscoveryRequest.InitialResourceVersionsEntry
	7,  // 8: envoy.service.discovery.v4alpha.DeltaDiscoveryRequest.error_detail:type_name -> google.rpc.Status
	4,  // 9: envoy.service.discovery.v4alpha.DeltaDiscoveryResponse.resources:type_name -> envoy.service.discovery.v4alpha.Resource
	11, // 10: envoy.service.discovery.v4alpha.DeltaDiscoveryResponse.udpa_removed_resources:type_name -> udpa.core.v1.ResourceName
	11, // 11: envoy.service.discovery.v4alpha.Resource.udpa_resource_name:type_name -> udpa.core.v1.ResourceName
	8,  // 12: envoy.service.discovery.v4alpha.Resource.resource:type_name -> google.protobuf.Any
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_envoy_service_discovery_v4alpha_discovery_proto_init() }
func file_envoy_service_discovery_v4alpha_discovery_proto_init() {
	if File_envoy_service_discovery_v4alpha_discovery_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoveryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaDiscoveryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeltaDiscoveryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Resource_Name)(nil),
		(*Resource_UdpaResourceName)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_discovery_v4alpha_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_service_discovery_v4alpha_discovery_proto_goTypes,
		DependencyIndexes: file_envoy_service_discovery_v4alpha_discovery_proto_depIdxs,
		MessageInfos:      file_envoy_service_discovery_v4alpha_discovery_proto_msgTypes,
	}.Build()
	File_envoy_service_discovery_v4alpha_discovery_proto = out.File
	file_envoy_service_discovery_v4alpha_discovery_proto_rawDesc = nil
	file_envoy_service_discovery_v4alpha_discovery_proto_goTypes = nil
	file_envoy_service_discovery_v4alpha_discovery_proto_depIdxs = nil
}
