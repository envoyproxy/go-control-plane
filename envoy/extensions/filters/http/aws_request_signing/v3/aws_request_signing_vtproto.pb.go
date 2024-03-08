//go:build vtprotobuf
// +build vtprotobuf
// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/aws_request_signing/v3/aws_request_signing.proto

package aws_request_signingv3

import (
	durationpb "github.com/planetscale/vtprotobuf/types/known/durationpb"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *AwsRequestSigning_QueryString) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AwsRequestSigning_QueryString) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AwsRequestSigning_QueryString) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.ExpirationTime != nil {
		size, err := (*durationpb.Duration)(m.ExpirationTime).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AwsRequestSigning) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AwsRequestSigning) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AwsRequestSigning) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.QueryString != nil {
		size, err := m.QueryString.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	}
	if m.SigningAlgorithm != 0 {
		i = encodeVarint(dAtA, i, uint64(m.SigningAlgorithm))
		i--
		dAtA[i] = 0x30
	}
	if len(m.MatchExcludedHeaders) > 0 {
		for iNdEx := len(m.MatchExcludedHeaders) - 1; iNdEx >= 0; iNdEx-- {
			if vtmsg, ok := interface{}(m.MatchExcludedHeaders[iNdEx]).(interface {
				MarshalToSizedBufferVTStrict([]byte) (int, error)
			}); ok {
				size, err := vtmsg.MarshalToSizedBufferVTStrict(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarint(dAtA, i, uint64(size))
			} else {
				encoded, err := proto.Marshal(m.MatchExcludedHeaders[iNdEx])
				if err != nil {
					return 0, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = encodeVarint(dAtA, i, uint64(len(encoded)))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.UseUnsignedPayload {
		i--
		if m.UseUnsignedPayload {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.HostRewrite) > 0 {
		i -= len(m.HostRewrite)
		copy(dAtA[i:], m.HostRewrite)
		i = encodeVarint(dAtA, i, uint64(len(m.HostRewrite)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Region) > 0 {
		i -= len(m.Region)
		copy(dAtA[i:], m.Region)
		i = encodeVarint(dAtA, i, uint64(len(m.Region)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarint(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AwsRequestSigningPerRoute) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AwsRequestSigningPerRoute) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *AwsRequestSigningPerRoute) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.StatPrefix) > 0 {
		i -= len(m.StatPrefix)
		copy(dAtA[i:], m.StatPrefix)
		i = encodeVarint(dAtA, i, uint64(len(m.StatPrefix)))
		i--
		dAtA[i] = 0x12
	}
	if m.AwsRequestSigning != nil {
		size, err := m.AwsRequestSigning.MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AwsRequestSigning_QueryString) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExpirationTime != nil {
		l = (*durationpb.Duration)(m.ExpirationTime).SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AwsRequestSigning) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Region)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.HostRewrite)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.UseUnsignedPayload {
		n += 2
	}
	if len(m.MatchExcludedHeaders) > 0 {
		for _, e := range m.MatchExcludedHeaders {
			if size, ok := interface{}(e).(interface {
				SizeVT() int
			}); ok {
				l = size.SizeVT()
			} else {
				l = proto.Size(e)
			}
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.SigningAlgorithm != 0 {
		n += 1 + sov(uint64(m.SigningAlgorithm))
	}
	if m.QueryString != nil {
		l = m.QueryString.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *AwsRequestSigningPerRoute) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AwsRequestSigning != nil {
		l = m.AwsRequestSigning.SizeVT()
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.StatPrefix)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
