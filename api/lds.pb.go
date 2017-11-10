// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/lds.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf4 "github.com/golang/protobuf/ptypes/struct"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Listener_DrainType int32

const (
	Listener_DEFAULT     Listener_DrainType = 0
	Listener_MODIFY_ONLY Listener_DrainType = 1
)

var Listener_DrainType_name = map[int32]string{
	0: "DEFAULT",
	1: "MODIFY_ONLY",
}
var Listener_DrainType_value = map[string]int32{
	"DEFAULT":     0,
	"MODIFY_ONLY": 1,
}

func (x Listener_DrainType) String() string {
	return proto.EnumName(Listener_DrainType_name, int32(x))
}
func (Listener_DrainType) EnumDescriptor() ([]byte, []int) { return fileDescriptor9, []int{3, 0} }

type Filter struct {
	Name         string                   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Config       *google_protobuf4.Struct `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	DeprecatedV1 *Filter_DeprecatedV1     `protobuf:"bytes,3,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetConfig() *google_protobuf4.Struct {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Filter) GetDeprecatedV1() *Filter_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

type Filter_DeprecatedV1 struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *Filter_DeprecatedV1) Reset()                    { *m = Filter_DeprecatedV1{} }
func (m *Filter_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*Filter_DeprecatedV1) ProtoMessage()               {}
func (*Filter_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0, 0} }

func (m *Filter_DeprecatedV1) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type FilterChainMatch struct {
	SniDomains         []string                       `protobuf:"bytes,1,rep,name=sni_domains,json=sniDomains" json:"sni_domains,omitempty"`
	PrefixRanges       []*CidrRange                   `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges" json:"prefix_ranges,omitempty"`
	AddressSuffix      string                         `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix" json:"address_suffix,omitempty"`
	SuffixLen          *google_protobuf.UInt32Value   `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen" json:"suffix_len,omitempty"`
	SourcePrefixRanges []*CidrRange                   `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges" json:"source_prefix_ranges,omitempty"`
	SourcePorts        []*google_protobuf.UInt32Value `protobuf:"bytes,7,rep,name=source_ports,json=sourcePorts" json:"source_ports,omitempty"`
	DestinationPort    *google_protobuf.UInt32Value   `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort" json:"destination_port,omitempty"`
}

func (m *FilterChainMatch) Reset()                    { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string            { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()               {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{1} }

func (m *FilterChainMatch) GetSniDomains() []string {
	if m != nil {
		return m.SniDomains
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *google_protobuf.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []*google_protobuf.UInt32Value {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetDestinationPort() *google_protobuf.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

type FilterChain struct {
	FilterChainMatch *FilterChainMatch          `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch" json:"filter_chain_match,omitempty"`
	TlsContext       *DownstreamTlsContext      `protobuf:"bytes,2,opt,name=tls_context,json=tlsContext" json:"tls_context,omitempty"`
	Filters          []*Filter                  `protobuf:"bytes,3,rep,name=filters" json:"filters,omitempty"`
	UseProxyProto    *google_protobuf.BoolValue `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto" json:"use_proxy_proto,omitempty"`
	Metadata         *Metadata                  `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
	TransportSocket  *TransportSocket           `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket" json:"transport_socket,omitempty"`
}

func (m *FilterChain) Reset()                    { *m = FilterChain{} }
func (m *FilterChain) String() string            { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()               {}
func (*FilterChain) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{2} }

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetTlsContext() *DownstreamTlsContext {
	if m != nil {
		return m.TlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *google_protobuf.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

type Listener struct {
	Name                          string                       `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Address                       *Address                     `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	FilterChains                  []*FilterChain               `protobuf:"bytes,3,rep,name=filter_chains,json=filterChains" json:"filter_chains,omitempty"`
	UseOriginalDst                *google_protobuf.BoolValue   `protobuf:"bytes,4,opt,name=use_original_dst,json=useOriginalDst" json:"use_original_dst,omitempty"`
	PerConnectionBufferLimitBytes *google_protobuf.UInt32Value `protobuf:"bytes,5,opt,name=per_connection_buffer_limit_bytes,json=perConnectionBufferLimitBytes" json:"per_connection_buffer_limit_bytes,omitempty"`
	Metadata                      *Metadata                    `protobuf:"bytes,6,opt,name=metadata" json:"metadata,omitempty"`
	DeprecatedV1                  *Listener_DeprecatedV1       `protobuf:"bytes,7,opt,name=deprecated_v1,json=deprecatedV1" json:"deprecated_v1,omitempty"`
	DrainType                     Listener_DrainType           `protobuf:"varint,8,opt,name=drain_type,json=drainType,enum=envoy.api.v2.Listener_DrainType" json:"drain_type,omitempty"`
}

func (m *Listener) Reset()                    { *m = Listener{} }
func (m *Listener) String() string            { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()               {}
func (*Listener) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3} }

func (m *Listener) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Listener) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Listener) GetFilterChains() []*FilterChain {
	if m != nil {
		return m.FilterChains
	}
	return nil
}

func (m *Listener) GetUseOriginalDst() *google_protobuf.BoolValue {
	if m != nil {
		return m.UseOriginalDst
	}
	return nil
}

func (m *Listener) GetPerConnectionBufferLimitBytes() *google_protobuf.UInt32Value {
	if m != nil {
		return m.PerConnectionBufferLimitBytes
	}
	return nil
}

func (m *Listener) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Listener) GetDeprecatedV1() *Listener_DeprecatedV1 {
	if m != nil {
		return m.DeprecatedV1
	}
	return nil
}

func (m *Listener) GetDrainType() Listener_DrainType {
	if m != nil {
		return m.DrainType
	}
	return Listener_DEFAULT
}

type Listener_DeprecatedV1 struct {
	BindToPort *google_protobuf.BoolValue `protobuf:"bytes,1,opt,name=bind_to_port,json=bindToPort" json:"bind_to_port,omitempty"`
}

func (m *Listener_DeprecatedV1) Reset()                    { *m = Listener_DeprecatedV1{} }
func (m *Listener_DeprecatedV1) String() string            { return proto.CompactTextString(m) }
func (*Listener_DeprecatedV1) ProtoMessage()               {}
func (*Listener_DeprecatedV1) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{3, 0} }

func (m *Listener_DeprecatedV1) GetBindToPort() *google_protobuf.BoolValue {
	if m != nil {
		return m.BindToPort
	}
	return nil
}

func init() {
	proto.RegisterType((*Filter)(nil), "envoy.api.v2.Filter")
	proto.RegisterType((*Filter_DeprecatedV1)(nil), "envoy.api.v2.Filter.DeprecatedV1")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v2.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v2.FilterChain")
	proto.RegisterType((*Listener)(nil), "envoy.api.v2.Listener")
	proto.RegisterType((*Listener_DeprecatedV1)(nil), "envoy.api.v2.Listener.DeprecatedV1")
	proto.RegisterEnum("envoy.api.v2.Listener_DrainType", Listener_DrainType_name, Listener_DrainType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ListenerDiscoveryService service

type ListenerDiscoveryServiceClient interface {
	StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error)
	FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type listenerDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewListenerDiscoveryServiceClient(cc *grpc.ClientConn) ListenerDiscoveryServiceClient {
	return &listenerDiscoveryServiceClient{cc}
}

func (c *listenerDiscoveryServiceClient) StreamListeners(ctx context.Context, opts ...grpc.CallOption) (ListenerDiscoveryService_StreamListenersClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ListenerDiscoveryService_serviceDesc.Streams[0], c.cc, "/envoy.api.v2.ListenerDiscoveryService/StreamListeners", opts...)
	if err != nil {
		return nil, err
	}
	x := &listenerDiscoveryServiceStreamListenersClient{stream}
	return x, nil
}

type ListenerDiscoveryService_StreamListenersClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type listenerDiscoveryServiceStreamListenersClient struct {
	grpc.ClientStream
}

func (x *listenerDiscoveryServiceStreamListenersClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *listenerDiscoveryServiceClient) FetchListeners(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/envoy.api.v2.ListenerDiscoveryService/FetchListeners", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ListenerDiscoveryService service

type ListenerDiscoveryServiceServer interface {
	StreamListeners(ListenerDiscoveryService_StreamListenersServer) error
	FetchListeners(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

func RegisterListenerDiscoveryServiceServer(s *grpc.Server, srv ListenerDiscoveryServiceServer) {
	s.RegisterService(&_ListenerDiscoveryService_serviceDesc, srv)
}

func _ListenerDiscoveryService_StreamListeners_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ListenerDiscoveryServiceServer).StreamListeners(&listenerDiscoveryServiceStreamListenersServer{stream})
}

type ListenerDiscoveryService_StreamListenersServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type listenerDiscoveryServiceStreamListenersServer struct {
	grpc.ServerStream
}

func (x *listenerDiscoveryServiceStreamListenersServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *listenerDiscoveryServiceStreamListenersServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ListenerDiscoveryService_FetchListeners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v2.ListenerDiscoveryService/FetchListeners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ListenerDiscoveryServiceServer).FetchListeners(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ListenerDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v2.ListenerDiscoveryService",
	HandlerType: (*ListenerDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchListeners",
			Handler:    _ListenerDiscoveryService_FetchListeners_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamListeners",
			Handler:       _ListenerDiscoveryService_StreamListeners_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/lds.proto",
}

func init() { proto.RegisterFile("api/lds.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x92, 0xdb, 0x44,
	0x10, 0x8e, 0xe2, 0x60, 0xaf, 0xdb, 0xbf, 0x0c, 0x81, 0x08, 0xd7, 0x26, 0x71, 0x4c, 0x51, 0xb5,
	0x70, 0x90, 0x89, 0x72, 0x0b, 0x29, 0x52, 0x59, 0x1b, 0xc3, 0x56, 0x79, 0xd9, 0x2d, 0xd9, 0x49,
	0x91, 0x93, 0x6a, 0x2c, 0x8d, 0xbc, 0x53, 0xc8, 0x33, 0x62, 0x66, 0xe4, 0xac, 0xaf, 0x3c, 0x02,
	0xbc, 0x05, 0x6f, 0xc0, 0x73, 0xf0, 0x0a, 0xdc, 0x38, 0x73, 0xa7, 0x66, 0x24, 0x39, 0x96, 0x93,
	0x5d, 0xf6, 0xc0, 0x6d, 0x7a, 0xfa, 0xeb, 0x6f, 0x5a, 0xfd, 0x75, 0xb7, 0xa0, 0x85, 0x13, 0x3a,
	0x8c, 0x43, 0xe9, 0x24, 0x82, 0x2b, 0x8e, 0x9a, 0x84, 0xad, 0xf9, 0xc6, 0xc1, 0x09, 0x75, 0xd6,
	0x6e, 0xef, 0x43, 0xed, 0xc4, 0x61, 0x28, 0x88, 0xcc, 0x01, 0xbd, 0xb6, 0xbe, 0x5a, 0x60, 0x49,
	0x72, 0xfb, 0x23, 0x6d, 0x87, 0x54, 0x06, 0x7c, 0x4d, 0xc4, 0x26, 0xbf, 0x34, 0xa4, 0xb2, 0x20,
	0xed, 0x1d, 0x2e, 0x39, 0x5f, 0xc6, 0x64, 0x68, 0xd8, 0x18, 0xe3, 0x0a, 0x2b, 0xca, 0xd9, 0xbe,
	0xd7, 0x58, 0x8b, 0x34, 0x1a, 0x4a, 0x25, 0xd2, 0x40, 0xe5, 0xde, 0x07, 0xfb, 0xde, 0x37, 0x02,
	0x27, 0x09, 0x11, 0x79, 0xf4, 0xe0, 0x0f, 0x0b, 0xaa, 0x13, 0x1a, 0x2b, 0x22, 0x10, 0x82, 0x3b,
	0x0c, 0xaf, 0x88, 0x6d, 0xf5, 0xad, 0xa3, 0xba, 0x67, 0xce, 0x68, 0x08, 0xd5, 0x80, 0xb3, 0x88,
	0x2e, 0xed, 0xdb, 0x7d, 0xeb, 0xa8, 0xe1, 0xde, 0x73, 0x32, 0x3e, 0xa7, 0xe0, 0x73, 0x66, 0xe6,
	0x35, 0x2f, 0x87, 0xa1, 0x09, 0xb4, 0x42, 0x92, 0x08, 0x12, 0x60, 0x45, 0x42, 0x7f, 0xfd, 0xd8,
	0xae, 0x98, 0xb8, 0x47, 0xce, 0x6e, 0x61, 0x9c, 0xec, 0x45, 0x67, 0xbc, 0x45, 0xbe, 0x7a, 0xec,
	0x35, 0xc3, 0x1d, 0xab, 0x37, 0x80, 0xe6, 0xae, 0x57, 0x27, 0xa7, 0x36, 0xc9, 0x36, 0x39, 0x7d,
	0x1e, 0xfc, 0x5e, 0x81, 0x6e, 0xc6, 0x34, 0xba, 0xc0, 0x94, 0x9d, 0x62, 0x15, 0x5c, 0xa0, 0x87,
	0xd0, 0x90, 0x8c, 0xfa, 0x21, 0x5f, 0x61, 0xca, 0xa4, 0x6d, 0xf5, 0x2b, 0x47, 0x75, 0x0f, 0x24,
	0xa3, 0xe3, 0xec, 0x06, 0x3d, 0x83, 0x56, 0x22, 0x48, 0x44, 0x2f, 0x7d, 0x81, 0xd9, 0x92, 0x48,
	0xbb, 0xd2, 0xaf, 0x98, 0x2f, 0x2b, 0x65, 0x38, 0xa2, 0xa1, 0xf0, 0xb4, 0xdf, 0x6b, 0x66, 0x68,
	0x63, 0x48, 0xf4, 0x39, 0xb4, 0x73, 0x41, 0x7d, 0x99, 0x46, 0x11, 0xbd, 0xb4, 0xef, 0x98, 0x8c,
	0x5a, 0xf9, 0xed, 0xcc, 0x5c, 0xa2, 0xaf, 0x01, 0x32, 0xb7, 0x1f, 0x13, 0x66, 0x7f, 0x60, 0x6a,
	0x70, 0xf8, 0x4e, 0xed, 0x5e, 0x9e, 0x30, 0xf5, 0xc4, 0x7d, 0x85, 0xe3, 0x94, 0x78, 0xf5, 0x0c,
	0x3f, 0x25, 0x0c, 0x9d, 0xc0, 0x5d, 0xc9, 0x53, 0x11, 0x10, 0xbf, 0x9c, 0x68, 0xf5, 0xfa, 0x44,
	0x51, 0x16, 0x74, 0xbe, 0x9b, 0xee, 0x73, 0x68, 0x16, 0x54, 0x5c, 0x28, 0x69, 0xd7, 0x0c, 0xc5,
	0xf5, 0x99, 0x34, 0x72, 0x1e, 0x1d, 0x80, 0xbe, 0x83, 0x6e, 0x48, 0xa4, 0xa2, 0xcc, 0xf4, 0x9c,
	0x61, 0xb1, 0x0f, 0x6e, 0xf0, 0x39, 0x9d, 0x9d, 0x28, 0xcd, 0x34, 0xf8, 0xb5, 0x02, 0x8d, 0x1d,
	0xb1, 0xd0, 0x14, 0x50, 0x64, 0x4c, 0x3f, 0xd0, 0xb6, 0xbf, 0xd2, 0xea, 0x19, 0x79, 0x1b, 0xee,
	0x83, 0xf7, 0x75, 0xcb, 0x5b, 0x8d, 0xbd, 0x6e, 0xb4, 0xaf, 0xfa, 0x08, 0x1a, 0x2a, 0x96, 0x7e,
	0xc0, 0x99, 0x22, 0x97, 0x2a, 0x6f, 0xd6, 0x41, 0x99, 0x66, 0xcc, 0xdf, 0x30, 0xa9, 0x04, 0xc1,
	0xab, 0x79, 0x2c, 0x47, 0x19, 0xd2, 0x03, 0xb5, 0x3d, 0x23, 0x07, 0x6a, 0x19, 0x71, 0xd1, 0x13,
	0x77, 0xdf, 0x97, 0x87, 0x57, 0x80, 0xd0, 0x31, 0x74, 0x52, 0xa9, 0x45, 0xe2, 0x97, 0x1b, 0xdf,
	0x54, 0xc1, 0x34, 0x43, 0xc3, 0xed, 0xbd, 0x53, 0x9a, 0x63, 0xce, 0xe3, 0xac, 0x30, 0xad, 0x54,
	0x92, 0x73, 0x1d, 0x71, 0x6e, 0x16, 0x86, 0x0b, 0x07, 0x2b, 0xa2, 0x70, 0x88, 0x15, 0xce, 0xdb,
	0xe4, 0x93, 0xf2, 0xa3, 0xa7, 0xb9, 0xd7, 0xdb, 0xe2, 0xd0, 0xf7, 0xd0, 0x55, 0x02, 0x33, 0xa9,
	0xc5, 0xf0, 0x25, 0x0f, 0x7e, 0x22, 0xca, 0xae, 0x9a, 0xd8, 0xfb, 0xe5, 0xd8, 0x79, 0x81, 0x9a,
	0x19, 0x90, 0xd7, 0x51, 0xe5, 0x8b, 0xc1, 0x3f, 0x77, 0xe0, 0x60, 0x4a, 0xa5, 0x22, 0xec, 0xca,
	0xf9, 0xaf, 0xe5, 0x8d, 0x9d, 0xd7, 0xf4, 0xe3, 0xf2, 0x0b, 0x2f, 0x32, 0xa7, 0x57, 0xa0, 0xd0,
	0x37, 0xd0, 0xda, 0x95, 0xb5, 0xa8, 0xe4, 0xa7, 0x57, 0x2a, 0xea, 0x35, 0x77, 0xc4, 0x94, 0x68,
	0x0c, 0x5d, 0x5d, 0x53, 0x2e, 0xe8, 0x92, 0x32, 0x1c, 0xfb, 0xa1, 0x54, 0x37, 0x28, 0x6a, 0x3b,
	0x95, 0xe4, 0x2c, 0x0f, 0x19, 0x4b, 0x85, 0x22, 0x78, 0x94, 0xe8, 0x14, 0x38, 0x63, 0x24, 0x30,
	0x8d, 0xbb, 0x48, 0xa3, 0x88, 0x08, 0x3f, 0xa6, 0x2b, 0xaa, 0xfc, 0xc5, 0x46, 0x11, 0x79, 0xa3,
	0xa9, 0xbc, 0x9f, 0x10, 0x31, 0xda, 0xb2, 0x1c, 0x1b, 0x92, 0xa9, 0xe6, 0x38, 0xd6, 0x14, 0x25,
	0xf5, 0xaa, 0x37, 0x56, 0x6f, 0x6f, 0x43, 0xd6, 0x4c, 0xe0, 0x67, 0xe5, 0xc0, 0x42, 0x95, 0x6b,
	0x76, 0x24, 0x7a, 0x0e, 0x10, 0x0a, 0x3d, 0x3b, 0x66, 0x33, 0xea, 0xa9, 0x6c, 0xbb, 0xfd, 0xab,
	0x68, 0x34, 0x70, 0xbe, 0x49, 0x88, 0x57, 0x0f, 0x8b, 0x63, 0x6f, 0xba, 0xb7, 0x64, 0x9f, 0x41,
	0x73, 0x41, 0x59, 0xe8, 0x2b, 0x9e, 0x0d, 0xba, 0xf5, 0x9f, 0x85, 0x07, 0x8d, 0x9f, 0x73, 0x33,
	0xe1, 0x5f, 0x40, 0x7d, 0xfb, 0x0a, 0x6a, 0x40, 0x6d, 0xfc, 0xed, 0xe4, 0xc5, 0xcb, 0xe9, 0xbc,
	0x7b, 0x0b, 0x75, 0xa0, 0x71, 0x7a, 0x36, 0x3e, 0x99, 0xbc, 0xf6, 0xcf, 0x7e, 0x98, 0xbe, 0xee,
	0x5a, 0xee, 0xdf, 0x16, 0xd8, 0x45, 0x6a, 0xe3, 0xe2, 0xe7, 0x37, 0x23, 0x62, 0x4d, 0x03, 0x82,
	0x7e, 0x84, 0xce, 0xcc, 0x8c, 0x69, 0x81, 0x90, 0x68, 0x6f, 0x21, 0x6c, 0x43, 0x3c, 0xf2, 0x73,
	0x4a, 0xa4, 0xea, 0x3d, 0xbc, 0xd2, 0x2f, 0x13, 0xce, 0x24, 0x19, 0xdc, 0x3a, 0xb2, 0xbe, 0xb2,
	0x50, 0x0a, 0xed, 0x09, 0x51, 0xc1, 0xc5, 0xff, 0x48, 0x3c, 0xf8, 0xe5, 0xcf, 0xbf, 0x7e, 0xbb,
	0x7d, 0x38, 0xb8, 0x37, 0x5c, 0xbb, 0x6f, 0xff, 0xe3, 0x4f, 0xe3, 0xe2, 0x85, 0xa7, 0xd6, 0x97,
	0x8b, 0xaa, 0x29, 0xdc, 0x93, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x12, 0xdf, 0x86, 0x2c,
	0x08, 0x00, 0x00,
}
