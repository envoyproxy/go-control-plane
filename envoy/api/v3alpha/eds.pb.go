// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/eds.proto

package envoy_api_v3alpha

import (
	context "context"
	fmt "fmt"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/endpoint"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type ClusterLoadAssignment struct {
	ClusterName          string                          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*endpoint.Endpoint   `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy   `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_35014a248c9ef3f8, []int{0}
}

func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(m, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads           []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor  *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	EndpointStaleAfter      *duration.Duration                           `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	DisableOverprovisioning bool                                         `protobuf:"varint,5,opt,name=disable_overprovisioning,json=disableOverprovisioning,proto3" json:"disable_overprovisioning,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                     `json:"-"`
	XXX_unrecognized        []byte                                       `json:"-"`
	XXX_sizecache           int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_35014a248c9ef3f8, []int{0, 1}
}

func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *duration.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetDisableOverprovisioning() bool {
	if m != nil {
		return m.DisableOverprovisioning
	}
	return false
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                   `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *_type.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_35014a248c9ef3f8, []int{0, 1, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *_type.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v3alpha.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*endpoint.Endpoint)(nil), "envoy.api.v3alpha.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v3alpha.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v3alpha.ClusterLoadAssignment.Policy.DropOverload")
}

func init() { proto.RegisterFile("envoy/api/v3alpha/eds.proto", fileDescriptor_35014a248c9ef3f8) }

var fileDescriptor_35014a248c9ef3f8 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4d, 0x53, 0x13, 0x31,
	0x18, 0xc7, 0xc9, 0xb6, 0x40, 0x09, 0x58, 0x30, 0xbe, 0xb0, 0x56, 0xd4, 0xf2, 0x32, 0x63, 0xe9,
	0x61, 0xcb, 0xc0, 0x45, 0x3b, 0x5c, 0xa8, 0xa5, 0x8e, 0x0e, 0x4a, 0x67, 0x19, 0xf5, 0x66, 0x27,
	0xdd, 0x0d, 0x25, 0xc3, 0x36, 0x89, 0x49, 0x5a, 0xdd, 0xf1, 0xe6, 0x78, 0xf0, 0xac, 0xdf, 0xc2,
	0xaf, 0xe3, 0x57, 0xe0, 0x13, 0x78, 0xe4, 0xe4, 0xec, 0x6b, 0xa1, 0x5d, 0x18, 0x3d, 0x78, 0xcb,
	0xee, 0xff, 0x79, 0x7e, 0xcf, 0x6b, 0x02, 0xef, 0x13, 0x36, 0xe4, 0x7e, 0x0d, 0x0b, 0x5a, 0x1b,
	0xee, 0x60, 0x4f, 0x9c, 0xe0, 0x1a, 0x71, 0x95, 0x25, 0x24, 0xd7, 0x1c, 0xdd, 0x0c, 0x45, 0x0b,
	0x0b, 0x6a, 0xc5, 0x62, 0x69, 0x75, 0xd2, 0xde, 0xa5, 0xca, 0xe1, 0x43, 0x22, 0xfd, 0xc8, 0xab,
	0xb4, 0x99, 0x81, 0x64, 0xae, 0xe0, 0x94, 0xe9, 0xf4, 0x10, 0x9b, 0x9a, 0x91, 0xa9, 0xf6, 0x05,
	0xa9, 0x09, 0x22, 0x1d, 0x92, 0x2a, 0x2b, 0x3d, 0xce, 0x7b, 0x1e, 0x09, 0x29, 0x98, 0x31, 0xae,
	0xb1, 0xa6, 0x9c, 0xc5, 0x89, 0x95, 0x96, 0x87, 0xd8, 0xa3, 0x2e, 0xd6, 0xa4, 0x96, 0x1c, 0x62,
	0xe1, 0x61, 0xec, 0x16, 0x7e, 0x75, 0x07, 0xc7, 0xb5, 0x8f, 0x12, 0x0b, 0x41, 0xa4, 0xba, 0x4a,
	0x77, 0x07, 0x32, 0x24, 0x47, 0xfa, 0xda, 0xf7, 0x59, 0x78, 0xe7, 0x99, 0x37, 0x50, 0x9a, 0xc8,
	0x03, 0x8e, 0xdd, 0x3d, 0xa5, 0x68, 0x8f, 0xf5, 0x09, 0xd3, 0xa8, 0x0a, 0x17, 0x9c, 0x48, 0xe8,
	0x30, 0xdc, 0x27, 0x26, 0x28, 0x83, 0xca, 0x5c, 0x63, 0xf6, 0xbc, 0x91, 0x97, 0x46, 0x19, 0xd8,
	0xf3, 0xb1, 0xf8, 0x1a, 0xf7, 0x09, 0x7a, 0x05, 0xe7, 0x92, 0x42, 0x95, 0x69, 0x94, 0x73, 0x95,
	0xf9, 0xed, 0x9a, 0x35, 0xd1, 0x4b, 0x2b, 0x6d, 0xc6, 0x01, 0x77, 0xb0, 0x47, 0xb5, 0x7f, 0xd0,
	0xdd, 0x4f, 0xdc, 0xec, 0x11, 0x01, 0x11, 0xb8, 0x18, 0x84, 0x74, 0x3b, 0x23, 0xe8, 0x74, 0x08,
	0xdd, 0xcd, 0x80, 0x66, 0x66, 0x6f, 0x05, 0x69, 0xb9, 0x29, 0x7c, 0x9f, 0x69, 0xe9, 0xdb, 0x45,
	0x76, 0xe9, 0x27, 0x7a, 0x0e, 0x67, 0x04, 0xf7, 0xa8, 0xe3, 0x9b, 0xf9, 0x32, 0xb8, 0x22, 0xe5,
	0x6c, 0x7a, 0x3b, 0x74, 0xb3, 0x63, 0xf7, 0x52, 0x0f, 0xde, 0xca, 0x88, 0x87, 0x96, 0x60, 0xee,
	0x94, 0xf8, 0x51, 0xe3, 0xec, 0xe0, 0x88, 0xea, 0x70, 0x7a, 0x88, 0xbd, 0x01, 0x31, 0x8d, 0x30,
	0xe0, 0xc6, 0x75, 0x3d, 0x4a, 0x60, 0x76, 0xe4, 0x52, 0x37, 0x9e, 0x80, 0xd2, 0x59, 0x0e, 0xce,
	0x44, 0xb1, 0x91, 0x03, 0x8b, 0xae, 0xe4, 0xa2, 0x13, 0x2c, 0xa2, 0xc7, 0xb1, 0x9b, 0xf4, 0x7d,
	0xf7, 0x1f, 0x8b, 0xb0, 0x9a, 0x92, 0x8b, 0xc3, 0x18, 0x62, 0xdf, 0x70, 0x2f, 0x7c, 0x29, 0xf4,
	0x1e, 0x2e, 0x07, 0x7c, 0x21, 0xf9, 0x90, 0x2a, 0xca, 0x19, 0x65, 0xbd, 0xce, 0x31, 0x76, 0x34,
	0x97, 0x66, 0x2e, 0xac, 0x60, 0xc5, 0x8a, 0xf6, 0xcb, 0x4a, 0xf6, 0xcb, 0x7a, 0xf3, 0x82, 0xe9,
	0x9d, 0xed, 0xb7, 0x41, 0xca, 0xe1, 0xb2, 0x54, 0x8d, 0xf2, 0x94, 0x7d, 0x77, 0x9c, 0xd2, 0x0a,
	0x21, 0xe8, 0x1d, 0xbc, 0x9d, 0xd4, 0xdb, 0x51, 0x1a, 0x7b, 0xa4, 0x83, 0x8f, 0x35, 0x91, 0xf1,
	0x3c, 0xee, 0x4d, 0xc0, 0x9b, 0xf1, 0xf2, 0x36, 0xe0, 0x79, 0x63, 0xf6, 0x27, 0xc8, 0x57, 0x8d,
	0xc2, 0x94, 0x8d, 0x12, 0xc4, 0x51, 0x40, 0xd8, 0x0b, 0x00, 0xe8, 0x29, 0x34, 0x5d, 0xaa, 0x70,
	0xd7, 0x23, 0x9d, 0xf1, 0xd0, 0xe6, 0x74, 0x19, 0x54, 0x0a, 0xf6, 0x72, 0xac, 0x1f, 0x8e, 0xc9,
	0xa5, 0xcf, 0x70, 0xe1, 0x62, 0x4b, 0xd0, 0x3a, 0x2c, 0x38, 0x58, 0x93, 0x1e, 0x97, 0xfe, 0xf8,
	0x1d, 0x48, 0x05, 0xd4, 0x82, 0x8b, 0xe1, 0x34, 0xe2, 0x3b, 0x8d, 0x7b, 0xc9, 0x88, 0x1f, 0xc4,
	0xe3, 0x08, 0x6e, 0xbc, 0xd5, 0x92, 0xd8, 0x09, 0xd2, 0xc7, 0x5e, 0x3b, 0xb2, 0xb3, 0xc3, 0x19,
	0xb6, 0x53, 0xa7, 0x97, 0xf9, 0x02, 0x58, 0x32, 0xb6, 0x7f, 0x1b, 0xd0, 0x4c, 0xc6, 0xdf, 0x4c,
	0x1e, 0x9b, 0x23, 0x22, 0x87, 0xd4, 0x21, 0xa8, 0x0b, 0x17, 0x8f, 0xb4, 0x24, 0xb8, 0x3f, 0x5a,
	0xe4, 0xf5, 0x8c, 0x99, 0xa7, 0x7e, 0x36, 0xf9, 0x30, 0x20, 0x4a, 0x97, 0x36, 0xae, 0x37, 0x52,
	0x82, 0x33, 0x45, 0xd6, 0xa6, 0x2a, 0x60, 0x0b, 0xa0, 0x53, 0x58, 0x6c, 0x12, 0x4f, 0xe3, 0x51,
	0x88, 0x4a, 0x96, 0x77, 0x60, 0x32, 0x11, 0x67, 0xf3, 0x2f, 0x2c, 0x2f, 0x05, 0xfb, 0x0a, 0x60,
	0xb1, 0x45, 0xb4, 0x73, 0xf2, 0x5f, 0x0a, 0x7a, 0xfc, 0xe5, 0xd7, 0xd9, 0x0f, 0x63, 0x75, 0x6d,
	0x65, 0xf2, 0xf9, 0xae, 0xa7, 0xcf, 0x4b, 0x1d, 0x54, 0x1b, 0x5b, 0xf0, 0x11, 0xe5, 0x11, 0x52,
	0x48, 0xfe, 0xc9, 0x9f, 0xa4, 0x37, 0x0a, 0xfb, 0xae, 0x6a, 0x07, 0xbb, 0xd8, 0x06, 0xdf, 0x00,
	0xe8, 0xce, 0x84, 0x7b, 0xb9, 0xf3, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x36, 0x10, 0x09, 0x53,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v3alpha.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v3alpha.EndpointDiscoveryService/DeltaEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceDeltaEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_DeltaEndpointsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceDeltaEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v3alpha.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	DeltaEndpoints(EndpointDiscoveryService_DeltaEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_DeltaEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).DeltaEndpoints(&endpointDiscoveryServiceDeltaEndpointsServer{stream})
}

type EndpointDiscoveryService_DeltaEndpointsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceDeltaEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v3alpha.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v3alpha.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaEndpoints",
			Handler:       _EndpointDiscoveryService_DeltaEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v3alpha/eds.proto",
}
