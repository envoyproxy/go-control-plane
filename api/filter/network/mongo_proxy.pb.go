// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/filter/network/mongo_proxy.proto

package network

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_filter "github.com/envoyproxy/go-control-plane/api/filter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MongoProxy struct {
	StatPrefix string                          `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix" json:"stat_prefix,omitempty"`
	AccessLog  string                          `protobuf:"bytes,2,opt,name=access_log,json=accessLog" json:"access_log,omitempty"`
	Delay      *envoy_api_v2_filter.FaultDelay `protobuf:"bytes,3,opt,name=delay" json:"delay,omitempty"`
}

func (m *MongoProxy) Reset()                    { *m = MongoProxy{} }
func (m *MongoProxy) String() string            { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()               {}
func (*MongoProxy) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *MongoProxy) GetDelay() *envoy_api_v2_filter.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.api.v2.filter.network.MongoProxy")
}

func init() { proto.RegisterFile("api/filter/network/mongo_proxy.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x49, 0x2c, 0xc8, 0xd4,
	0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0xcf,
	0xcd, 0xcf, 0x4b, 0xcf, 0x8f, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x4e, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4, 0x4b, 0x2c, 0xc8, 0xd4, 0x2b, 0x33, 0xd2, 0x83,
	0x28, 0xd7, 0x83, 0x2a, 0x97, 0x12, 0x43, 0x32, 0x22, 0x2d, 0xb1, 0x34, 0xa7, 0x04, 0xa2, 0x49,
	0xa9, 0x99, 0x91, 0x8b, 0xcb, 0x17, 0x64, 0x54, 0x00, 0xc8, 0x24, 0x21, 0x79, 0x2e, 0xee, 0xe2,
	0x92, 0xc4, 0x92, 0xf8, 0x82, 0xa2, 0xd4, 0xb4, 0xcc, 0x0a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce,
	0x20, 0x2e, 0x90, 0x50, 0x00, 0x58, 0x44, 0x48, 0x96, 0x8b, 0x2b, 0x31, 0x39, 0x39, 0xb5, 0xb8,
	0x38, 0x3e, 0x27, 0x3f, 0x5d, 0x82, 0x09, 0x2c, 0xcf, 0x09, 0x11, 0xf1, 0xc9, 0x4f, 0x17, 0x32,
	0xe5, 0x62, 0x4d, 0x49, 0xcd, 0x49, 0xac, 0x94, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd7,
	0xc3, 0xe6, 0x26, 0x37, 0x90, 0xfd, 0x2e, 0x20, 0x65, 0x41, 0x10, 0xd5, 0x49, 0x6c, 0x60, 0xc7,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd9, 0xff, 0xf5, 0xcc, 0xe9, 0x00, 0x00, 0x00,
}
