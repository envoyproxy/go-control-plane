// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/api/v2/filter/network/client_ssl_auth.proto

/*
	Package network is a generated protocol buffer package.

	It is generated from these files:
		envoy/api/v2/filter/network/client_ssl_auth.proto
		envoy/api/v2/filter/network/ext_authz.proto
		envoy/api/v2/filter/network/http_connection_manager.proto
		envoy/api/v2/filter/network/mongo_proxy.proto
		envoy/api/v2/filter/network/rate_limit.proto
		envoy/api/v2/filter/network/redis_proxy.proto
		envoy/api/v2/filter/network/tcp_proxy.proto

	It has these top-level messages:
		ClientSSLAuth
		ExtAuthz
		HttpConnectionManager
		Rds
		HttpFilter
		MongoProxy
		RateLimit
		RedisProxy
		TcpProxy
*/
package network

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_api_v2_core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
import google_protobuf2 "github.com/gogo/protobuf/types"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ClientSSLAuth struct {
	// The :ref:`cluster manager <arch_overview_cluster_manager>` cluster that runs
	// the authentication service. The filter will connect to the service every 60s to fetch the list
	// of principals. The service must support the expected :ref:`REST API
	// <config_network_filters_client_ssl_auth_rest_api>`.
	AuthApiCluster string `protobuf:"bytes,1,opt,name=auth_api_cluster,json=authApiCluster,proto3" json:"auth_api_cluster,omitempty"`
	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_client_ssl_auth_stats>`.
	StatPrefix string `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Time in milliseconds between principal refreshes from the
	// authentication service. Default is 60000 (60s). The actual fetch time
	// will be this value plus a random jittered value between
	// 0-refresh_delay_ms milliseconds.
	RefreshDelay *google_protobuf2.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay" json:"refresh_delay,omitempty"`
	// An optional list of IP address and subnet masks that should be white
	// listed for access by the filter. If no list is provided, there is no
	// IP white list.
	IpWhiteList []*envoy_api_v2_core.CidrRange `protobuf:"bytes,4,rep,name=ip_white_list,json=ipWhiteList" json:"ip_white_list,omitempty"`
}

func (m *ClientSSLAuth) Reset()                    { *m = ClientSSLAuth{} }
func (m *ClientSSLAuth) String() string            { return proto.CompactTextString(m) }
func (*ClientSSLAuth) ProtoMessage()               {}
func (*ClientSSLAuth) Descriptor() ([]byte, []int) { return fileDescriptorClientSslAuth, []int{0} }

func (m *ClientSSLAuth) GetAuthApiCluster() string {
	if m != nil {
		return m.AuthApiCluster
	}
	return ""
}

func (m *ClientSSLAuth) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ClientSSLAuth) GetRefreshDelay() *google_protobuf2.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ClientSSLAuth) GetIpWhiteList() []*envoy_api_v2_core.CidrRange {
	if m != nil {
		return m.IpWhiteList
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientSSLAuth)(nil), "envoy.api.v2.filter.network.ClientSSLAuth")
}
func (m *ClientSSLAuth) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientSSLAuth) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AuthApiCluster) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintClientSslAuth(dAtA, i, uint64(len(m.AuthApiCluster)))
		i += copy(dAtA[i:], m.AuthApiCluster)
	}
	if len(m.StatPrefix) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintClientSslAuth(dAtA, i, uint64(len(m.StatPrefix)))
		i += copy(dAtA[i:], m.StatPrefix)
	}
	if m.RefreshDelay != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintClientSslAuth(dAtA, i, uint64(m.RefreshDelay.Size()))
		n1, err := m.RefreshDelay.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.IpWhiteList) > 0 {
		for _, msg := range m.IpWhiteList {
			dAtA[i] = 0x22
			i++
			i = encodeVarintClientSslAuth(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintClientSslAuth(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ClientSSLAuth) Size() (n int) {
	var l int
	_ = l
	l = len(m.AuthApiCluster)
	if l > 0 {
		n += 1 + l + sovClientSslAuth(uint64(l))
	}
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sovClientSslAuth(uint64(l))
	}
	if m.RefreshDelay != nil {
		l = m.RefreshDelay.Size()
		n += 1 + l + sovClientSslAuth(uint64(l))
	}
	if len(m.IpWhiteList) > 0 {
		for _, e := range m.IpWhiteList {
			l = e.Size()
			n += 1 + l + sovClientSslAuth(uint64(l))
		}
	}
	return n
}

func sovClientSslAuth(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozClientSslAuth(x uint64) (n int) {
	return sovClientSslAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientSSLAuth) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClientSslAuth
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClientSSLAuth: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientSSLAuth: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthApiCluster", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSslAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClientSslAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthApiCluster = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSslAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClientSslAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StatPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshDelay", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSslAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClientSslAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RefreshDelay == nil {
				m.RefreshDelay = &google_protobuf2.Duration{}
			}
			if err := m.RefreshDelay.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpWhiteList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClientSslAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClientSslAuth
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpWhiteList = append(m.IpWhiteList, &envoy_api_v2_core.CidrRange{})
			if err := m.IpWhiteList[len(m.IpWhiteList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClientSslAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthClientSslAuth
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClientSslAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClientSslAuth
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClientSslAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClientSslAuth
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthClientSslAuth
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowClientSslAuth
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipClientSslAuth(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthClientSslAuth = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClientSslAuth   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/api/v2/filter/network/client_ssl_auth.proto", fileDescriptorClientSslAuth)
}

var fileDescriptorClientSslAuth = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0xe5, 0xb6, 0x42, 0x6a, 0xd2, 0x22, 0x94, 0x85, 0x50, 0x50, 0x88, 0x98, 0x2a, 0x06,
	0x5b, 0xa4, 0x3b, 0xa2, 0x97, 0xb1, 0x03, 0x4a, 0x07, 0x46, 0xcb, 0x6d, 0x9c, 0xe6, 0x08, 0x2b,
	0xb6, 0x6c, 0x27, 0xa5, 0xaf, 0xc6, 0xc4, 0xc8, 0xc8, 0x23, 0xa0, 0x6e, 0xac, 0x3c, 0x01, 0xca,
	0xa5, 0x43, 0xd9, 0x6c, 0xfd, 0xdf, 0x67, 0x9f, 0xff, 0x38, 0x0f, 0x3c, 0x2f, 0xe5, 0x9e, 0x30,
	0x05, 0xa4, 0x8c, 0x48, 0x0a, 0xc2, 0x72, 0x4d, 0x72, 0x6e, 0x77, 0x52, 0xbf, 0x92, 0x8d, 0x00,
	0x9e, 0x5b, 0x6a, 0x8c, 0xa0, 0xac, 0xb0, 0x19, 0x56, 0x5a, 0x5a, 0xe9, 0x5d, 0xd7, 0x0a, 0x66,
	0x0a, 0x70, 0x19, 0xe1, 0x46, 0xc1, 0xad, 0x32, 0xba, 0x3d, 0x79, 0x6f, 0x23, 0x35, 0x27, 0x2c,
	0x49, 0x34, 0x37, 0xa6, 0xb1, 0x47, 0xc1, 0x56, 0xca, 0xad, 0xe0, 0xa4, 0xbe, 0xad, 0x8b, 0x94,
	0x24, 0x85, 0x66, 0x16, 0x64, 0xde, 0xe6, 0x97, 0x25, 0x13, 0x90, 0x30, 0xcb, 0xc9, 0xf1, 0xd0,
	0x04, 0x77, 0xbf, 0xc8, 0x19, 0xce, 0xeb, 0x81, 0x56, 0xab, 0xe5, 0xb4, 0xb0, 0x99, 0x37, 0x71,
	0x2e, 0xaa, 0xb1, 0x28, 0x53, 0x40, 0x37, 0xa2, 0x30, 0x96, 0x6b, 0x1f, 0x85, 0x68, 0xdc, 0x9f,
	0xf5, 0xdf, 0x7f, 0x3e, 0xba, 0x3d, 0xdd, 0x09, 0x51, 0x7c, 0x5e, 0x21, 0x53, 0x05, 0xf3, 0x06,
	0xf0, 0xee, 0x1d, 0xd7, 0x58, 0x66, 0xa9, 0xd2, 0x3c, 0x85, 0x37, 0xbf, 0xf3, 0x9f, 0x77, 0xaa,
	0xf4, 0xb9, 0x0e, 0xbd, 0x47, 0x67, 0xa8, 0x79, 0xaa, 0xb9, 0xc9, 0x68, 0xc2, 0x05, 0xdb, 0xfb,
	0xdd, 0x10, 0x8d, 0xdd, 0xe8, 0x0a, 0x37, 0x1d, 0xf0, 0xb1, 0x03, 0x5e, 0xb4, 0x1d, 0xe2, 0x41,
	0xcb, 0x2f, 0x2a, 0xdc, 0x7b, 0x72, 0x86, 0xa0, 0xe8, 0x2e, 0x03, 0xcb, 0xa9, 0x00, 0x63, 0xfd,
	0x5e, 0xd8, 0x1d, 0xbb, 0xd1, 0x0d, 0x3e, 0xd9, 0x60, 0xb5, 0x24, 0x3c, 0x87, 0x44, 0xc7, 0x2c,
	0xdf, 0xf2, 0xd8, 0x05, 0xf5, 0x52, 0x19, 0x4b, 0x30, 0x76, 0x36, 0xf8, 0x3c, 0x04, 0xe8, 0xeb,
	0x10, 0xa0, 0xef, 0x43, 0x80, 0xd6, 0x67, 0xf5, 0x87, 0x93, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x84, 0x60, 0x8e, 0xf2, 0xb5, 0x01, 0x00, 0x00,
}
