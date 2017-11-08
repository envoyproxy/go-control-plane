// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/filter/network/redis_proxy.proto

package network

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf2 "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type RedisProxy struct {
	StatPrefix string                       `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	Cluster    string                       `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Settings   *RedisProxy_ConnPoolSettings `protobuf:"bytes,3,opt,name=settings" json:"settings,omitempty"`
}

func (m *RedisProxy) Reset()                    { *m = RedisProxy{} }
func (m *RedisProxy) String() string            { return proto.CompactTextString(m) }
func (*RedisProxy) ProtoMessage()               {}
func (*RedisProxy) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *RedisProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RedisProxy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy) GetSettings() *RedisProxy_ConnPoolSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type RedisProxy_ConnPoolSettings struct {
	OpTimeout *google_protobuf2.Duration `protobuf:"bytes,1,opt,name=op_timeout,json=opTimeout" json:"op_timeout,omitempty"`
}

func (m *RedisProxy_ConnPoolSettings) Reset()                    { *m = RedisProxy_ConnPoolSettings{} }
func (m *RedisProxy_ConnPoolSettings) String() string            { return proto.CompactTextString(m) }
func (*RedisProxy_ConnPoolSettings) ProtoMessage()               {}
func (*RedisProxy_ConnPoolSettings) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

func (m *RedisProxy_ConnPoolSettings) GetOpTimeout() *google_protobuf2.Duration {
	if m != nil {
		return m.OpTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RedisProxy)(nil), "envoy.api.v2.filter.network.RedisProxy")
	proto.RegisterType((*RedisProxy_ConnPoolSettings)(nil), "envoy.api.v2.filter.network.RedisProxy.ConnPoolSettings")
}

func init() { proto.RegisterFile("api/filter/network/redis_proxy.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x90, 0x80, 0x5e, 0x16, 0xe4, 0x29, 0x14, 0x09, 0x2a, 0xc4, 0xd0, 0xe9, 0x2c,
	0x85, 0xa5, 0x3b, 0x8c, 0x0c, 0x51, 0xe8, 0x1e, 0xa5, 0xf4, 0x12, 0x59, 0x04, 0x9f, 0x65, 0x5f,
	0x4a, 0xfb, 0x9b, 0xf9, 0x13, 0x28, 0x4e, 0x0a, 0x12, 0x43, 0x47, 0xdf, 0xfb, 0xde, 0xf3, 0x7b,
	0xf0, 0x58, 0x3b, 0xa3, 0x1b, 0xd3, 0x09, 0x79, 0x6d, 0x49, 0xbe, 0xd8, 0x7f, 0x68, 0x4f, 0x5b,
	0x13, 0x2a, 0xe7, 0x79, 0x7f, 0x40, 0xe7, 0x59, 0x58, 0xdd, 0x92, 0xdd, 0xf1, 0x01, 0x6b, 0x67,
	0x70, 0x97, 0xe3, 0x88, 0xe3, 0x84, 0xcf, 0xef, 0x5a, 0xe6, 0xb6, 0x23, 0x1d, 0xd1, 0x4d, 0xdf,
	0xe8, 0x6d, 0xef, 0x6b, 0x31, 0x6c, 0x47, 0xf3, 0xc3, 0x77, 0x02, 0x50, 0x0e, 0x91, 0xc5, 0x90,
	0xa8, 0xee, 0x21, 0x0d, 0x52, 0x4b, 0xe5, 0x3c, 0x35, 0x66, 0x9f, 0x25, 0x8b, 0x64, 0x39, 0x2b,
	0x61, 0x38, 0x15, 0xf1, 0xa2, 0x32, 0xb8, 0x7c, 0xef, 0xfa, 0x20, 0xe4, 0xb3, 0xb3, 0x28, 0x1e,
	0x9f, 0x6a, 0x0d, 0x57, 0x81, 0x44, 0x8c, 0x6d, 0x43, 0x76, 0xbe, 0x48, 0x96, 0x69, 0xbe, 0xc2,
	0x13, 0xcd, 0xf0, 0xef, 0x57, 0x7c, 0x66, 0x6b, 0x0b, 0xe6, 0xee, 0x6d, 0xf2, 0x97, 0xbf, 0x49,
	0xf3, 0x57, 0xb8, 0xfe, 0xaf, 0xaa, 0x15, 0x00, 0xbb, 0x4a, 0xcc, 0x27, 0x71, 0x2f, 0xb1, 0x63,
	0x9a, 0xdf, 0xe0, 0x38, 0x14, 0x8f, 0x43, 0xf1, 0x65, 0x1a, 0x5a, 0xce, 0xd8, 0xad, 0x47, 0x76,
	0x73, 0x11, 0xd5, 0xa7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x09, 0x2d, 0xe5, 0x59, 0x01,
	0x00, 0x00,
}
