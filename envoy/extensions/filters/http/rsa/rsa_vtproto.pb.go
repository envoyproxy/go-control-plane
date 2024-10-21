//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: envoy/extensions/filters/http/rsa/rsa.proto

package rsav3

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *RsaFilterConfig) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *RsaFilterConfig) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *RsaFilterConfig) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if len(m.RsaApis) > 0 {
		for iNdEx := len(m.RsaApis) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.RsaApis[iNdEx])
			copy(dAtA[i:], m.RsaApis[iNdEx])
			i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.RsaApis[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PublicKeyPath) > 0 {
		i -= len(m.PublicKeyPath)
		copy(dAtA[i:], m.PublicKeyPath)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.PublicKeyPath)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PrivateKeyPath) > 0 {
		i -= len(m.PrivateKeyPath)
		copy(dAtA[i:], m.PrivateKeyPath)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.PrivateKeyPath)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RsaFilterConfig) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PrivateKeyPath)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	l = len(m.PublicKeyPath)
	if l > 0 {
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if len(m.RsaApis) > 0 {
		for _, s := range m.RsaApis {
			l = len(s)
			n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}
