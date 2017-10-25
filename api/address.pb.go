// Code generated by protoc-gen-go.
// source: api/address.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}
func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

type Pipe struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *Pipe) Reset()                    { *m = Pipe{} }
func (m *Pipe) String() string            { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()               {}
func (*Pipe) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,enum=envoy.api.v2.SocketAddress_Protocol" json:"protocol,omitempty"`
	Address  string                 `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	ResolverName  string                        `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName" json:"resolver_name,omitempty"`
}

func (m *SocketAddress) Reset()                    { *m = SocketAddress{} }
func (m *SocketAddress) String() string            { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()               {}
func (*SocketAddress) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,oneof"`
}
type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}
func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SocketAddress_OneofMarshaler, _SocketAddress_OneofUnmarshaler, _SocketAddress_OneofSizer, []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

func _SocketAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NamedPort)
	case nil:
	default:
		return fmt.Errorf("SocketAddress.PortSpecifier has unexpected type %T", x)
	}
	return nil
}

func _SocketAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SocketAddress)
	switch tag {
	case 3: // port_specifier.port_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.PortSpecifier = &SocketAddress_PortValue{uint32(x)}
		return true, err
	case 4: // port_specifier.named_port
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PortSpecifier = &SocketAddress_NamedPort{x}
		return true, err
	default:
		return false, nil
	}
}

func _SocketAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.NamedPort)))
		n += len(x.NamedPort)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BindConfig struct {
	SourceAddress *SocketAddress `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
}

func (m *BindConfig) Reset()                    { *m = BindConfig{} }
func (m *BindConfig) String() string            { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()               {}
func (*BindConfig) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address isAddress_Address `protobuf_oneof:"address"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,oneof"`
}
type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}
func (*Address_Pipe) isAddress_Address()          {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Address) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Address_OneofMarshaler, _Address_OneofUnmarshaler, _Address_OneofSizer, []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

func _Address_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SocketAddress); err != nil {
			return err
		}
	case *Address_Pipe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipe); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Address.Address has unexpected type %T", x)
	}
	return nil
}

func _Address_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Address)
	switch tag {
	case 1: // address.socket_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SocketAddress)
		err := b.DecodeMessage(msg)
		m.Address = &Address_SocketAddress{msg}
		return true, err
	case 2: // address.pipe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pipe)
		err := b.DecodeMessage(msg)
		m.Address = &Address_Pipe{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Address_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		s := proto.Size(x.SocketAddress)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Address_Pipe:
		s := proto.Size(x.Pipe)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CidrRange struct {
	AddressPrefix string                       `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix" json:"address_prefix,omitempty"`
	PrefixLen     *google_protobuf.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen" json:"prefix_len,omitempty"`
}

func (m *CidrRange) Reset()                    { *m = CidrRange{} }
func (m *CidrRange) String() string            { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()               {}
func (*CidrRange) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *google_protobuf.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterType((*Pipe)(nil), "envoy.api.v2.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.api.v2.SocketAddress")
	proto.RegisterType((*BindConfig)(nil), "envoy.api.v2.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.api.v2.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.api.v2.CidrRange")
	proto.RegisterEnum("envoy.api.v2.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
}

func init() { proto.RegisterFile("api/address.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0x93, 0x36, 0xb0, 0xcd, 0xb4, 0x89, 0x16, 0x9f, 0xa2, 0x52, 0x41, 0x15, 0x40, 0xda,
	0x93, 0x57, 0x4a, 0x8f, 0x5c, 0x20, 0xdb, 0xc3, 0x22, 0x21, 0x14, 0x05, 0xca, 0x35, 0x72, 0x93,
	0x49, 0xb0, 0x48, 0x6d, 0xcb, 0xc9, 0x06, 0xb8, 0x22, 0x3e, 0x38, 0xb2, 0xe3, 0x54, 0xed, 0x05,
	0x89, 0x9b, 0xfd, 0xfc, 0x9b, 0x3f, 0x7e, 0x33, 0xf0, 0x8c, 0x29, 0xbe, 0x65, 0x4d, 0xa3, 0x71,
	0x18, 0xa8, 0xd2, 0x72, 0x94, 0xe4, 0x0c, 0xc5, 0x24, 0x7f, 0x51, 0xa6, 0x38, 0x9d, 0xb2, 0xf3,
	0x17, 0x9d, 0x94, 0x5d, 0x8f, 0x5b, 0xfb, 0x76, 0x7b, 0x68, 0xb7, 0x3f, 0x34, 0x53, 0x0a, 0xb5,
	0xa3, 0xd3, 0x73, 0x08, 0x0a, 0xae, 0x90, 0x10, 0x08, 0x14, 0x1b, 0xbf, 0x25, 0xfe, 0xa5, 0xbf,
	0x09, 0x4b, 0x7b, 0x4e, 0x7f, 0x1f, 0x41, 0xf4, 0x59, 0xd6, 0xdf, 0x71, 0x7c, 0x3f, 0x57, 0x20,
	0xef, 0xe0, 0xc4, 0x86, 0xd5, 0xb2, 0xb7, 0x64, 0x9c, 0xbd, 0xa6, 0x0f, 0xcb, 0xd1, 0x47, 0x38,
	0x2d, 0x1c, 0x5b, 0xde, 0x47, 0x91, 0x04, 0x56, 0xae, 0xdd, 0xe4, 0xc8, 0x96, 0x5a, 0xae, 0xe4,
	0x25, 0x80, 0x92, 0x7a, 0xac, 0x26, 0xd6, 0x1f, 0x30, 0x39, 0xbe, 0xf4, 0x37, 0xd1, 0xde, 0x2b,
	0x43, 0xa3, 0x7d, 0x35, 0x92, 0x01, 0x04, 0xbb, 0xc3, 0xa6, 0x32, 0x52, 0x12, 0x98, 0x68, 0x03,
	0x58, 0xad, 0x90, 0x7a, 0x24, 0xaf, 0x20, 0xd2, 0x38, 0xc8, 0x7e, 0x42, 0x5d, 0x19, 0x35, 0x79,
	0x62, 0x2b, 0x9c, 0x2d, 0xe2, 0x27, 0x76, 0x87, 0xe9, 0x05, 0x9c, 0x2c, 0x6d, 0x91, 0x15, 0x1c,
	0x7f, 0xd9, 0x15, 0x6b, 0xcf, 0x1c, 0x6e, 0xae, 0x8b, 0xb5, 0x9f, 0xaf, 0x21, 0xb6, 0x4d, 0x0c,
	0x0a, 0x6b, 0xde, 0x72, 0xd4, 0x69, 0x01, 0x90, 0x73, 0xd1, 0xec, 0xa4, 0x68, 0x79, 0x47, 0x72,
	0x88, 0x07, 0x79, 0xd0, 0x35, 0x56, 0xcb, 0x2f, 0x8c, 0x0d, 0xa7, 0xd9, 0xf3, 0x7f, 0xd8, 0x50,
	0x46, 0x73, 0x88, 0xbb, 0xa6, 0x7f, 0x7c, 0x58, 0x2d, 0x86, 0x5e, 0x9b, 0x7c, 0x86, 0xfd, 0x8f,
	0x7c, 0x7b, 0xcf, 0x64, 0x7c, 0x38, 0x96, 0x0d, 0x04, 0x8a, 0x2b, 0xb4, 0x8e, 0x9e, 0x66, 0xe4,
	0x71, 0xac, 0x19, 0xef, 0xde, 0x2b, 0x2d, 0x91, 0x87, 0xf7, 0xf6, 0xa7, 0x12, 0xc2, 0x1d, 0x6f,
	0x74, 0xc9, 0x44, 0x87, 0xe4, 0x0d, 0xc4, 0x4e, 0xaf, 0x94, 0xc6, 0x96, 0xff, 0x74, 0x8b, 0x10,
	0x39, 0xb5, 0xb0, 0x22, 0x79, 0x0b, 0x30, 0x3f, 0x57, 0x3d, 0x0a, 0x57, 0xee, 0x82, 0xce, 0x2b,
	0x46, 0x97, 0x15, 0xa3, 0x37, 0x1f, 0xc4, 0x78, 0x95, 0xd9, 0xa1, 0x95, 0xe1, 0xcc, 0x7f, 0x44,
	0x71, 0xfb, 0xd4, 0x02, 0x57, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x42, 0xbb, 0xa6, 0x24, 0xb4,
	0x02, 0x00, 0x00,
}
