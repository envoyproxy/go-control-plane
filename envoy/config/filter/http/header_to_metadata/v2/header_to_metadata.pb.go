// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/header_to_metadata/v2/header_to_metadata.proto

/*
	Package v2 is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/filter/http/header_to_metadata/v2/header_to_metadata.proto

	It has these top-level messages:
		Config
*/
package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"

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

type Config_ValueType int32

const (
	Config_STRING Config_ValueType = 0
	Config_NUMBER Config_ValueType = 1
)

var Config_ValueType_name = map[int32]string{
	0: "STRING",
	1: "NUMBER",
}
var Config_ValueType_value = map[string]int32{
	"STRING": 0,
	"NUMBER": 1,
}

func (x Config_ValueType) String() string {
	return proto.EnumName(Config_ValueType_name, int32(x))
}
func (Config_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorHeaderToMetadata, []int{0, 0}
}

type Config struct {
	// The list of rules to apply to requests.
	RequestRules []*Config_Rule `protobuf:"bytes,1,rep,name=request_rules,json=requestRules" json:"request_rules,omitempty"`
	// The list of rules to apply to responses.
	ResponseRules []*Config_Rule `protobuf:"bytes,2,rep,name=response_rules,json=responseRules" json:"response_rules,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptorHeaderToMetadata, []int{0} }

func (m *Config) GetRequestRules() []*Config_Rule {
	if m != nil {
		return m.RequestRules
	}
	return nil
}

func (m *Config) GetResponseRules() []*Config_Rule {
	if m != nil {
		return m.ResponseRules
	}
	return nil
}

type Config_KeyValuePair struct {
	// The namespace — if this is empty, the filter's namespace will be used.
	MetadataNamespace string `protobuf:"bytes,1,opt,name=metadata_namespace,json=metadataNamespace,proto3" json:"metadata_namespace,omitempty"`
	// The key to use within the namespace.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// The value to pair with the given key.
	//
	// When used for a `on_header_present` case, if value is non-empty it'll be used
	// instead of the header value. If both are empty, no metadata is added.
	//
	// When used for a `on_header_missing` case, a non-empty value must be provided
	// otherwise no metadata is added.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// The value's type — defaults to string.
	Type Config_ValueType `protobuf:"varint,4,opt,name=type,proto3,enum=envoy.config.filter.http.header_to_metadata.v2.Config_ValueType" json:"type,omitempty"`
}

func (m *Config_KeyValuePair) Reset()         { *m = Config_KeyValuePair{} }
func (m *Config_KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*Config_KeyValuePair) ProtoMessage()    {}
func (*Config_KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptorHeaderToMetadata, []int{0, 0}
}

func (m *Config_KeyValuePair) GetMetadataNamespace() string {
	if m != nil {
		return m.MetadataNamespace
	}
	return ""
}

func (m *Config_KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Config_KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Config_KeyValuePair) GetType() Config_ValueType {
	if m != nil {
		return m.Type
	}
	return Config_STRING
}

// A Rule defines what metadata to apply when a header is present or missing.
type Config_Rule struct {
	// The header that triggers this rule — required.
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// If the header is present, apply this metadata KeyValuePair.
	//
	// If the value in the KeyValuePair is non-empty, it'll be used instead
	// of the header value.
	OnHeaderPresent *Config_KeyValuePair `protobuf:"bytes,2,opt,name=on_header_present,json=onHeaderPresent" json:"on_header_present,omitempty"`
	// If the header is not present, apply this metadata KeyValuePair.
	//
	// The value in the KeyValuePair must be set, since it'll be used in lieu
	// of the missing header value.
	OnHeaderMissing *Config_KeyValuePair `protobuf:"bytes,3,opt,name=on_header_missing,json=onHeaderMissing" json:"on_header_missing,omitempty"`
	// Whether or not to remove the header after a rule is applied.
	//
	// This prevents headers from leaking.
	Remove bool `protobuf:"varint,4,opt,name=remove,proto3" json:"remove,omitempty"`
}

func (m *Config_Rule) Reset()                    { *m = Config_Rule{} }
func (m *Config_Rule) String() string            { return proto.CompactTextString(m) }
func (*Config_Rule) ProtoMessage()               {}
func (*Config_Rule) Descriptor() ([]byte, []int) { return fileDescriptorHeaderToMetadata, []int{0, 1} }

func (m *Config_Rule) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Config_Rule) GetOnHeaderPresent() *Config_KeyValuePair {
	if m != nil {
		return m.OnHeaderPresent
	}
	return nil
}

func (m *Config_Rule) GetOnHeaderMissing() *Config_KeyValuePair {
	if m != nil {
		return m.OnHeaderMissing
	}
	return nil
}

func (m *Config_Rule) GetRemove() bool {
	if m != nil {
		return m.Remove
	}
	return false
}

func init() {
	proto.RegisterType((*Config)(nil), "envoy.config.filter.http.header_to_metadata.v2.Config")
	proto.RegisterType((*Config_KeyValuePair)(nil), "envoy.config.filter.http.header_to_metadata.v2.Config.KeyValuePair")
	proto.RegisterType((*Config_Rule)(nil), "envoy.config.filter.http.header_to_metadata.v2.Config.Rule")
	proto.RegisterEnum("envoy.config.filter.http.header_to_metadata.v2.Config_ValueType", Config_ValueType_name, Config_ValueType_value)
}
func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RequestRules) > 0 {
		for _, msg := range m.RequestRules {
			dAtA[i] = 0xa
			i++
			i = encodeVarintHeaderToMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.ResponseRules) > 0 {
		for _, msg := range m.ResponseRules {
			dAtA[i] = 0x12
			i++
			i = encodeVarintHeaderToMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Config_KeyValuePair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config_KeyValuePair) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.MetadataNamespace) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.MetadataNamespace)))
		i += copy(dAtA[i:], m.MetadataNamespace)
	}
	if len(m.Key) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.Type != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(m.Type))
	}
	return i, nil
}

func (m *Config_Rule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config_Rule) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Header) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(len(m.Header)))
		i += copy(dAtA[i:], m.Header)
	}
	if m.OnHeaderPresent != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(m.OnHeaderPresent.Size()))
		n1, err := m.OnHeaderPresent.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.OnHeaderMissing != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHeaderToMetadata(dAtA, i, uint64(m.OnHeaderMissing.Size()))
		n2, err := m.OnHeaderMissing.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Remove {
		dAtA[i] = 0x20
		i++
		if m.Remove {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeVarintHeaderToMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Config) Size() (n int) {
	var l int
	_ = l
	if len(m.RequestRules) > 0 {
		for _, e := range m.RequestRules {
			l = e.Size()
			n += 1 + l + sovHeaderToMetadata(uint64(l))
		}
	}
	if len(m.ResponseRules) > 0 {
		for _, e := range m.ResponseRules {
			l = e.Size()
			n += 1 + l + sovHeaderToMetadata(uint64(l))
		}
	}
	return n
}

func (m *Config_KeyValuePair) Size() (n int) {
	var l int
	_ = l
	l = len(m.MetadataNamespace)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovHeaderToMetadata(uint64(m.Type))
	}
	return n
}

func (m *Config_Rule) Size() (n int) {
	var l int
	_ = l
	l = len(m.Header)
	if l > 0 {
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.OnHeaderPresent != nil {
		l = m.OnHeaderPresent.Size()
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.OnHeaderMissing != nil {
		l = m.OnHeaderMissing.Size()
		n += 1 + l + sovHeaderToMetadata(uint64(l))
	}
	if m.Remove {
		n += 2
	}
	return n
}

func sovHeaderToMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHeaderToMetadata(x uint64) (n int) {
	return sovHeaderToMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaderToMetadata
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestRules = append(m.RequestRules, &Config_Rule{})
			if err := m.RequestRules[len(m.RequestRules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResponseRules", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResponseRules = append(m.ResponseRules, &Config_Rule{})
			if err := m.ResponseRules[len(m.ResponseRules)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeaderToMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaderToMetadata
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
func (m *Config_KeyValuePair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaderToMetadata
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
			return fmt.Errorf("proto: KeyValuePair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyValuePair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataNamespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataNamespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Config_ValueType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHeaderToMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaderToMetadata
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
func (m *Config_Rule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaderToMetadata
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
			return fmt.Errorf("proto: Rule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Header = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnHeaderPresent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnHeaderPresent == nil {
				m.OnHeaderPresent = &Config_KeyValuePair{}
			}
			if err := m.OnHeaderPresent.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OnHeaderMissing", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
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
				return ErrInvalidLengthHeaderToMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OnHeaderMissing == nil {
				m.OnHeaderMissing = &Config_KeyValuePair{}
			}
			if err := m.OnHeaderMissing.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remove", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderToMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Remove = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipHeaderToMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaderToMetadata
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
func skipHeaderToMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeaderToMetadata
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
					return 0, ErrIntOverflowHeaderToMetadata
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
					return 0, ErrIntOverflowHeaderToMetadata
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
				return 0, ErrInvalidLengthHeaderToMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHeaderToMetadata
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
				next, err := skipHeaderToMetadata(dAtA[start:])
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
	ErrInvalidLengthHeaderToMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeaderToMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/filter/http/header_to_metadata/v2/header_to_metadata.proto", fileDescriptorHeaderToMetadata)
}

var fileDescriptorHeaderToMetadata = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x59, 0x27, 0xb5, 0xc8, 0xf4, 0x83, 0x76, 0xc5, 0x87, 0x15, 0xa4, 0x28, 0x94, 0x4b,
	0x2e, 0xd8, 0x92, 0x39, 0xc2, 0x01, 0xa5, 0x42, 0x05, 0xa1, 0x44, 0xd1, 0x12, 0x38, 0x70, 0x31,
	0xdb, 0x66, 0x9a, 0x5a, 0x38, 0xde, 0x65, 0x77, 0x63, 0xe1, 0xe7, 0xe1, 0x2d, 0x38, 0x71, 0x42,
	0x1c, 0x79, 0x04, 0x94, 0x1b, 0x6f, 0x81, 0x76, 0xd7, 0x46, 0x45, 0xea, 0xa5, 0x85, 0xdb, 0xce,
	0xcc, 0xce, 0xef, 0xff, 0x9f, 0x91, 0x06, 0x8e, 0xb1, 0xac, 0x44, 0x9d, 0x9c, 0x8a, 0xf2, 0x2c,
	0x5f, 0x26, 0x67, 0x79, 0x61, 0x50, 0x25, 0xe7, 0xc6, 0xc8, 0xe4, 0x1c, 0xf9, 0x02, 0x55, 0x66,
	0x44, 0xb6, 0x42, 0xc3, 0x17, 0xdc, 0xf0, 0xa4, 0x4a, 0x2f, 0xc9, 0xc6, 0x52, 0x09, 0x23, 0x68,
	0xec, 0x40, 0xb1, 0x07, 0xc5, 0x1e, 0x14, 0x5b, 0x50, 0x7c, 0x49, 0x4b, 0x95, 0xf6, 0xef, 0x55,
	0xbc, 0xc8, 0x17, 0xdc, 0x60, 0xd2, 0x3e, 0x3c, 0xe8, 0x70, 0xb3, 0x05, 0xe1, 0x91, 0xa3, 0xd0,
	0xf7, 0xb0, 0xab, 0xf0, 0xe3, 0x1a, 0xb5, 0xc9, 0xd4, 0xba, 0x40, 0x1d, 0x91, 0x61, 0x67, 0xb4,
	0x9d, 0x3e, 0xb9, 0xa2, 0x56, 0xec, 0x71, 0x31, 0x5b, 0x17, 0xc8, 0x76, 0x1a, 0xa2, 0x0d, 0x34,
	0x3d, 0x81, 0x3d, 0x85, 0x5a, 0x8a, 0x52, 0x63, 0x23, 0x11, 0xfc, 0xbb, 0xc4, 0x6e, 0x8b, 0x74,
	0x1a, 0xfd, 0x6f, 0x04, 0x76, 0x5e, 0x61, 0xfd, 0x96, 0x17, 0x6b, 0x9c, 0xf1, 0x5c, 0xd1, 0x47,
	0x40, 0xdb, 0xd6, 0xac, 0xe4, 0x2b, 0xd4, 0x92, 0x9f, 0x62, 0x44, 0x86, 0x64, 0xd4, 0x63, 0x07,
	0x6d, 0x65, 0xda, 0x16, 0xe8, 0x7d, 0xe8, 0x7c, 0xc0, 0x3a, 0x0a, 0x6c, 0x7d, 0xdc, 0xfb, 0xf2,
	0xeb, 0x6b, 0xa7, 0xab, 0x82, 0x21, 0x61, 0x36, 0x4b, 0x6f, 0xc3, 0x56, 0x65, 0xc1, 0x51, 0xc7,
	0xb5, 0xfb, 0x80, 0xce, 0xa1, 0x6b, 0x6a, 0x89, 0x51, 0x77, 0x48, 0x46, 0x7b, 0xe9, 0xb3, 0x6b,
	0x0e, 0xe3, 0x1c, 0xcf, 0x6b, 0x89, 0xcc, 0xd1, 0xfa, 0x9f, 0x03, 0xe8, 0xda, 0x91, 0xe8, 0x03,
	0x08, 0x7d, 0xa3, 0x37, 0x7d, 0xd1, 0x54, 0x53, 0xa0, 0x02, 0x0e, 0x44, 0x99, 0x35, 0x78, 0xa9,
	0x50, 0x63, 0x69, 0xdc, 0x08, 0xdb, 0xe9, 0xd1, 0x35, 0xed, 0x5c, 0xdc, 0x21, 0xbb, 0x25, 0xca,
	0x17, 0xee, 0xf3, 0xcc, 0xb3, 0xff, 0x16, 0x5c, 0xe5, 0x5a, 0xe7, 0xe5, 0xd2, 0x2d, 0xe5, 0x7f,
	0x0b, 0x4e, 0x3c, 0x9b, 0xde, 0x85, 0x50, 0xe1, 0x4a, 0x54, 0x7e, 0xcb, 0x37, 0x59, 0x13, 0x1d,
	0x3e, 0x84, 0xde, 0x9f, 0xc5, 0x51, 0x80, 0xf0, 0xf5, 0x9c, 0xbd, 0x9c, 0x1e, 0xef, 0xdf, 0xb0,
	0xef, 0xe9, 0x9b, 0xc9, 0xf8, 0x39, 0xdb, 0x27, 0xe3, 0xe5, 0xf7, 0xcd, 0x80, 0xfc, 0xd8, 0x0c,
	0xc8, 0xcf, 0xcd, 0x80, 0xc0, 0xd3, 0x5c, 0x78, 0x8b, 0x52, 0x89, 0x4f, 0xf5, 0x15, 0xdd, 0x8e,
	0xef, 0x78, 0x5f, 0x73, 0x31, 0x69, 0x92, 0x33, 0x7b, 0x47, 0x33, 0xf2, 0x2e, 0xa8, 0xd2, 0x93,
	0xd0, 0x1d, 0xd5, 0xe3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x22, 0x49, 0x39, 0xe8, 0x03,
	0x00, 0x00,
}
