//go:build vtprotobuf
// +build vtprotobuf

// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// source: contrib/envoy/extensions/compression/qatzstd/compressor/v3alpha/qatzstd.proto

package v3alpha

import (
	protohelpers "github.com/planetscale/vtprotobuf/protohelpers"
	wrapperspb "github.com/planetscale/vtprotobuf/types/known/wrapperspb"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *Qatzstd) MarshalVTStrict() (dAtA []byte, err error) {
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

func (m *Qatzstd) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *Qatzstd) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
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
	if m.QatZstdFallbackThreshold != nil {
		size, err := (*wrapperspb.UInt32Value)(m.QatZstdFallbackThreshold).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x3a
	}
	if m.EnableQatZstd {
		i--
		if m.EnableQatZstd {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.ChunkSize != nil {
		size, err := (*wrapperspb.UInt32Value)(m.ChunkSize).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0x2a
	}
	if m.Strategy != 0 {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(m.Strategy))
		i--
		dAtA[i] = 0x18
	}
	if m.EnableChecksum {
		i--
		if m.EnableChecksum {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.CompressionLevel != nil {
		size, err := (*wrapperspb.UInt32Value)(m.CompressionLevel).MarshalToSizedBufferVTStrict(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = protohelpers.EncodeVarint(dAtA, i, uint64(size))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Qatzstd) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CompressionLevel != nil {
		l = (*wrapperspb.UInt32Value)(m.CompressionLevel).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.EnableChecksum {
		n += 2
	}
	if m.Strategy != 0 {
		n += 1 + protohelpers.SizeOfVarint(uint64(m.Strategy))
	}
	if m.ChunkSize != nil {
		l = (*wrapperspb.UInt32Value)(m.ChunkSize).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.EnableQatZstd {
		n += 2
	}
	if m.QatZstdFallbackThreshold != nil {
		l = (*wrapperspb.UInt32Value)(m.QatZstdFallbackThreshold).SizeVT()
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}
