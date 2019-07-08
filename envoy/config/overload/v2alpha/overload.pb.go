// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/overload/v2alpha/overload.proto

/*
	Package v2alpha is a generated protocol buffer package.

	It is generated from these files:
		envoy/config/overload/v2alpha/overload.proto

	It has these top-level messages:
		ResourceMonitor
		ThresholdTrigger
		Trigger
		OverloadAction
		OverloadManager
*/
package v2alpha

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/types"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"
import _ "github.com/envoyproxy/protoc-gen-validate/validate"

import binary "encoding/binary"

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

type ResourceMonitor struct {
	// The name of the resource monitor to instantiate. Must match a registered
	// resource monitor type. The built-in resource monitors are:
	//
	// * :ref:`envoy.resource_monitors.fixed_heap
	//   <envoy_api_msg_config.resource_monitor.fixed_heap.v2alpha.FixedHeapConfig>`
	// * :ref:`envoy.resource_monitors.injected_resource
	//   <envoy_api_msg_config.resource_monitor.injected_resource.v2alpha.InjectedResourceConfig>`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Configuration for the resource monitor being instantiated.
	//
	// Types that are valid to be assigned to ConfigType:
	//	*ResourceMonitor_Config
	//	*ResourceMonitor_TypedConfig
	ConfigType isResourceMonitor_ConfigType `protobuf_oneof:"config_type"`
}

func (m *ResourceMonitor) Reset()                    { *m = ResourceMonitor{} }
func (m *ResourceMonitor) String() string            { return proto.CompactTextString(m) }
func (*ResourceMonitor) ProtoMessage()               {}
func (*ResourceMonitor) Descriptor() ([]byte, []int) { return fileDescriptorOverload, []int{0} }

type isResourceMonitor_ConfigType interface {
	isResourceMonitor_ConfigType()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ResourceMonitor_Config struct {
	Config *google_protobuf2.Struct `protobuf:"bytes,2,opt,name=config,oneof"`
}
type ResourceMonitor_TypedConfig struct {
	TypedConfig *google_protobuf.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,oneof"`
}

func (*ResourceMonitor_Config) isResourceMonitor_ConfigType()      {}
func (*ResourceMonitor_TypedConfig) isResourceMonitor_ConfigType() {}

func (m *ResourceMonitor) GetConfigType() isResourceMonitor_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ResourceMonitor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResourceMonitor) GetConfig() *google_protobuf2.Struct {
	if x, ok := m.GetConfigType().(*ResourceMonitor_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ResourceMonitor) GetTypedConfig() *google_protobuf.Any {
	if x, ok := m.GetConfigType().(*ResourceMonitor_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ResourceMonitor) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ResourceMonitor_OneofMarshaler, _ResourceMonitor_OneofUnmarshaler, _ResourceMonitor_OneofSizer, []interface{}{
		(*ResourceMonitor_Config)(nil),
		(*ResourceMonitor_TypedConfig)(nil),
	}
}

func _ResourceMonitor_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ResourceMonitor)
	// config_type
	switch x := m.ConfigType.(type) {
	case *ResourceMonitor_Config:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Config); err != nil {
			return err
		}
	case *ResourceMonitor_TypedConfig:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TypedConfig); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ResourceMonitor.ConfigType has unexpected type %T", x)
	}
	return nil
}

func _ResourceMonitor_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ResourceMonitor)
	switch tag {
	case 2: // config_type.config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf2.Struct)
		err := b.DecodeMessage(msg)
		m.ConfigType = &ResourceMonitor_Config{msg}
		return true, err
	case 3: // config_type.typed_config
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf.Any)
		err := b.DecodeMessage(msg)
		m.ConfigType = &ResourceMonitor_TypedConfig{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ResourceMonitor_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ResourceMonitor)
	// config_type
	switch x := m.ConfigType.(type) {
	case *ResourceMonitor_Config:
		s := proto.Size(x.Config)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ResourceMonitor_TypedConfig:
		s := proto.Size(x.TypedConfig)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ThresholdTrigger struct {
	// If the resource pressure is greater than or equal to this value, the trigger
	// will fire.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ThresholdTrigger) Reset()                    { *m = ThresholdTrigger{} }
func (m *ThresholdTrigger) String() string            { return proto.CompactTextString(m) }
func (*ThresholdTrigger) ProtoMessage()               {}
func (*ThresholdTrigger) Descriptor() ([]byte, []int) { return fileDescriptorOverload, []int{1} }

func (m *ThresholdTrigger) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Trigger struct {
	// The name of the resource this is a trigger for.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to TriggerOneof:
	//	*Trigger_Threshold
	TriggerOneof isTrigger_TriggerOneof `protobuf_oneof:"trigger_oneof"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptorOverload, []int{2} }

type isTrigger_TriggerOneof interface {
	isTrigger_TriggerOneof()
	MarshalTo([]byte) (int, error)
	Size() int
}

type Trigger_Threshold struct {
	Threshold *ThresholdTrigger `protobuf:"bytes,2,opt,name=threshold,oneof"`
}

func (*Trigger_Threshold) isTrigger_TriggerOneof() {}

func (m *Trigger) GetTriggerOneof() isTrigger_TriggerOneof {
	if m != nil {
		return m.TriggerOneof
	}
	return nil
}

func (m *Trigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Trigger) GetThreshold() *ThresholdTrigger {
	if x, ok := m.GetTriggerOneof().(*Trigger_Threshold); ok {
		return x.Threshold
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Trigger) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Trigger_OneofMarshaler, _Trigger_OneofUnmarshaler, _Trigger_OneofSizer, []interface{}{
		(*Trigger_Threshold)(nil),
	}
}

func _Trigger_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Trigger)
	// trigger_oneof
	switch x := m.TriggerOneof.(type) {
	case *Trigger_Threshold:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Threshold); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Trigger.TriggerOneof has unexpected type %T", x)
	}
	return nil
}

func _Trigger_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Trigger)
	switch tag {
	case 2: // trigger_oneof.threshold
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ThresholdTrigger)
		err := b.DecodeMessage(msg)
		m.TriggerOneof = &Trigger_Threshold{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Trigger_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Trigger)
	// trigger_oneof
	switch x := m.TriggerOneof.(type) {
	case *Trigger_Threshold:
		s := proto.Size(x.Threshold)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type OverloadAction struct {
	// The name of the overload action. This is just a well-known string that listeners can
	// use for registering callbacks. Custom overload actions should be named using reverse
	// DNS to ensure uniqueness.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A set of triggers for this action. If any of these triggers fire the overload action
	// is activated. Listeners are notified when the overload action transitions from
	// inactivated to activated, or vice versa.
	Triggers []*Trigger `protobuf:"bytes,2,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *OverloadAction) Reset()                    { *m = OverloadAction{} }
func (m *OverloadAction) String() string            { return proto.CompactTextString(m) }
func (*OverloadAction) ProtoMessage()               {}
func (*OverloadAction) Descriptor() ([]byte, []int) { return fileDescriptorOverload, []int{3} }

func (m *OverloadAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OverloadAction) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type OverloadManager struct {
	// The interval for refreshing resource usage.
	RefreshInterval *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval" json:"refresh_interval,omitempty"`
	// The set of resources to monitor.
	ResourceMonitors []*ResourceMonitor `protobuf:"bytes,2,rep,name=resource_monitors,json=resourceMonitors" json:"resource_monitors,omitempty"`
	// The set of overload actions.
	Actions []*OverloadAction `protobuf:"bytes,3,rep,name=actions" json:"actions,omitempty"`
}

func (m *OverloadManager) Reset()                    { *m = OverloadManager{} }
func (m *OverloadManager) String() string            { return proto.CompactTextString(m) }
func (*OverloadManager) ProtoMessage()               {}
func (*OverloadManager) Descriptor() ([]byte, []int) { return fileDescriptorOverload, []int{4} }

func (m *OverloadManager) GetRefreshInterval() *google_protobuf1.Duration {
	if m != nil {
		return m.RefreshInterval
	}
	return nil
}

func (m *OverloadManager) GetResourceMonitors() []*ResourceMonitor {
	if m != nil {
		return m.ResourceMonitors
	}
	return nil
}

func (m *OverloadManager) GetActions() []*OverloadAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceMonitor)(nil), "envoy.config.overload.v2alpha.ResourceMonitor")
	proto.RegisterType((*ThresholdTrigger)(nil), "envoy.config.overload.v2alpha.ThresholdTrigger")
	proto.RegisterType((*Trigger)(nil), "envoy.config.overload.v2alpha.Trigger")
	proto.RegisterType((*OverloadAction)(nil), "envoy.config.overload.v2alpha.OverloadAction")
	proto.RegisterType((*OverloadManager)(nil), "envoy.config.overload.v2alpha.OverloadManager")
}
func (m *ResourceMonitor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceMonitor) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOverload(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.ConfigType != nil {
		nn1, err := m.ConfigType.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *ResourceMonitor_Config) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Config != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOverload(dAtA, i, uint64(m.Config.Size()))
		n2, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func (m *ResourceMonitor_TypedConfig) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.TypedConfig != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintOverload(dAtA, i, uint64(m.TypedConfig.Size()))
		n3, err := m.TypedConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}
func (m *ThresholdTrigger) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ThresholdTrigger) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x9
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i += 8
	}
	return i, nil
}

func (m *Trigger) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Trigger) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOverload(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.TriggerOneof != nil {
		nn4, err := m.TriggerOneof.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn4
	}
	return i, nil
}

func (m *Trigger_Threshold) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Threshold != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintOverload(dAtA, i, uint64(m.Threshold.Size()))
		n5, err := m.Threshold.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}
func (m *OverloadAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OverloadAction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOverload(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Triggers) > 0 {
		for _, msg := range m.Triggers {
			dAtA[i] = 0x12
			i++
			i = encodeVarintOverload(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *OverloadManager) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OverloadManager) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RefreshInterval != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintOverload(dAtA, i, uint64(m.RefreshInterval.Size()))
		n6, err := m.RefreshInterval.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.ResourceMonitors) > 0 {
		for _, msg := range m.ResourceMonitors {
			dAtA[i] = 0x12
			i++
			i = encodeVarintOverload(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Actions) > 0 {
		for _, msg := range m.Actions {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintOverload(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintOverload(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ResourceMonitor) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOverload(uint64(l))
	}
	if m.ConfigType != nil {
		n += m.ConfigType.Size()
	}
	return n
}

func (m *ResourceMonitor_Config) Size() (n int) {
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovOverload(uint64(l))
	}
	return n
}
func (m *ResourceMonitor_TypedConfig) Size() (n int) {
	var l int
	_ = l
	if m.TypedConfig != nil {
		l = m.TypedConfig.Size()
		n += 1 + l + sovOverload(uint64(l))
	}
	return n
}
func (m *ThresholdTrigger) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	return n
}

func (m *Trigger) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOverload(uint64(l))
	}
	if m.TriggerOneof != nil {
		n += m.TriggerOneof.Size()
	}
	return n
}

func (m *Trigger_Threshold) Size() (n int) {
	var l int
	_ = l
	if m.Threshold != nil {
		l = m.Threshold.Size()
		n += 1 + l + sovOverload(uint64(l))
	}
	return n
}
func (m *OverloadAction) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovOverload(uint64(l))
	}
	if len(m.Triggers) > 0 {
		for _, e := range m.Triggers {
			l = e.Size()
			n += 1 + l + sovOverload(uint64(l))
		}
	}
	return n
}

func (m *OverloadManager) Size() (n int) {
	var l int
	_ = l
	if m.RefreshInterval != nil {
		l = m.RefreshInterval.Size()
		n += 1 + l + sovOverload(uint64(l))
	}
	if len(m.ResourceMonitors) > 0 {
		for _, e := range m.ResourceMonitors {
			l = e.Size()
			n += 1 + l + sovOverload(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovOverload(uint64(l))
		}
	}
	return n
}

func sovOverload(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozOverload(x uint64) (n int) {
	return sovOverload(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceMonitor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOverload
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
			return fmt.Errorf("proto: ResourceMonitor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceMonitor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &google_protobuf2.Struct{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigType = &ResourceMonitor_Config{v}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypedConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &google_protobuf.Any{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.ConfigType = &ResourceMonitor_TypedConfig{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOverload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOverload
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
func (m *ThresholdTrigger) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOverload
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
			return fmt.Errorf("proto: ThresholdTrigger: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ThresholdTrigger: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipOverload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOverload
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
func (m *Trigger) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOverload
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
			return fmt.Errorf("proto: Trigger: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Trigger: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ThresholdTrigger{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.TriggerOneof = &Trigger_Threshold{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOverload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOverload
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
func (m *OverloadAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOverload
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
			return fmt.Errorf("proto: OverloadAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OverloadAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Triggers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Triggers = append(m.Triggers, &Trigger{})
			if err := m.Triggers[len(m.Triggers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOverload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOverload
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
func (m *OverloadManager) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOverload
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
			return fmt.Errorf("proto: OverloadManager: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OverloadManager: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefreshInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RefreshInterval == nil {
				m.RefreshInterval = &google_protobuf1.Duration{}
			}
			if err := m.RefreshInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResourceMonitors", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResourceMonitors = append(m.ResourceMonitors, &ResourceMonitor{})
			if err := m.ResourceMonitors[len(m.ResourceMonitors)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOverload
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
				return ErrInvalidLengthOverload
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, &OverloadAction{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOverload(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOverload
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
func skipOverload(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOverload
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
					return 0, ErrIntOverflowOverload
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
					return 0, ErrIntOverflowOverload
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
				return 0, ErrInvalidLengthOverload
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowOverload
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
				next, err := skipOverload(dAtA[start:])
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
	ErrInvalidLengthOverload = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOverload   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/config/overload/v2alpha/overload.proto", fileDescriptorOverload)
}

var fileDescriptorOverload = []byte{
	// 518 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xcf, 0xd9, 0x6d, 0xd3, 0x9c, 0x09, 0x09, 0xa7, 0x8a, 0x26, 0x15, 0x8d, 0x22, 0x0f, 0x28,
	0x08, 0xb0, 0x45, 0x98, 0x98, 0x50, 0xdd, 0x4a, 0x84, 0xa1, 0xb4, 0x72, 0x3b, 0xb1, 0x58, 0xd7,
	0xf8, 0xe2, 0x58, 0x72, 0xef, 0xa2, 0xf3, 0xd9, 0xc2, 0x13, 0x33, 0x23, 0x7c, 0x0f, 0x76, 0xc4,
	0xd4, 0x91, 0x91, 0x8f, 0x80, 0xb2, 0x75, 0xe3, 0x23, 0x20, 0xdf, 0x9d, 0x83, 0x92, 0x4a, 0xb5,
	0x27, 0xeb, 0xfd, 0xfe, 0xbc, 0xdf, 0xbd, 0xf7, 0xe0, 0x0b, 0x42, 0x73, 0x56, 0xb8, 0x53, 0x46,
	0x67, 0x71, 0xe4, 0xb2, 0x9c, 0xf0, 0x84, 0xe1, 0xd0, 0xcd, 0xc7, 0x38, 0x59, 0xcc, 0xf1, 0xaa,
	0xe0, 0x2c, 0x38, 0x13, 0x0c, 0x1d, 0x4a, 0xb6, 0xa3, 0xd8, 0xce, 0x0a, 0xd4, 0xec, 0x83, 0x7e,
	0xc4, 0x58, 0x94, 0x10, 0x57, 0x92, 0xaf, 0xb2, 0x99, 0x8b, 0x69, 0xa1, 0x94, 0x07, 0x83, 0x4d,
	0x28, 0xcc, 0x38, 0x16, 0x31, 0xa3, 0x1a, 0x7f, 0xb2, 0x89, 0xa7, 0x82, 0x67, 0x53, 0xa1, 0xd1,
	0xfd, 0x1c, 0x27, 0x71, 0x88, 0x05, 0x71, 0xab, 0x1f, 0x05, 0xd8, 0xdf, 0x01, 0xec, 0xf8, 0x24,
	0x65, 0x19, 0x9f, 0x92, 0x53, 0x46, 0x63, 0xc1, 0x38, 0x3a, 0x84, 0x5b, 0x14, 0x5f, 0x93, 0x1e,
	0x18, 0x82, 0x51, 0xcb, 0x6b, 0xfd, 0xbc, 0xbd, 0x31, 0xb7, 0xb8, 0x31, 0x04, 0xbe, 0x2c, 0xa3,
	0x57, 0x70, 0x47, 0xe5, 0xef, 0x19, 0x43, 0x30, 0xb2, 0xc6, 0xfb, 0x8e, 0x6a, 0xed, 0x54, 0xad,
	0x9d, 0x0b, 0xd9, 0x7a, 0xd2, 0xf0, 0x35, 0x11, 0xbd, 0x81, 0x0f, 0x44, 0xb1, 0x20, 0x61, 0xa0,
	0x85, 0xa6, 0x14, 0xee, 0xdd, 0x11, 0x1e, 0xd1, 0x62, 0xd2, 0xf0, 0x2d, 0xc9, 0x3d, 0x96, 0x54,
	0xaf, 0x0d, 0x2d, 0x25, 0x0a, 0xca, 0xaa, 0x7d, 0x0c, 0xbb, 0x97, 0x73, 0x4e, 0xd2, 0x39, 0x4b,
	0xc2, 0x4b, 0x1e, 0x47, 0x11, 0xe1, 0xc8, 0x85, 0xdb, 0x39, 0x4e, 0x32, 0x15, 0x18, 0x78, 0xfd,
	0x32, 0xf0, 0x1e, 0x42, 0xfd, 0x86, 0xfc, 0xfe, 0xbe, 0x7d, 0xd6, 0xd0, 0x9f, 0xaf, 0x78, 0xf6,
	0x57, 0x00, 0x9b, 0x95, 0xb8, 0xe6, 0xb1, 0x67, 0xb0, 0x25, 0xaa, 0x7e, 0xfa, 0xbd, 0xae, 0x73,
	0xef, 0x12, 0x9d, 0xcd, 0x7c, 0x93, 0x86, 0xff, 0xdf, 0xc3, 0x7b, 0x0c, 0xdb, 0x42, 0xd5, 0x03,
	0x46, 0x09, 0x9b, 0xa1, 0xed, 0x1f, 0xb7, 0x37, 0x26, 0xb0, 0x3f, 0xc3, 0x87, 0x67, 0xda, 0xe9,
	0x68, 0x5a, 0xee, 0xb5, 0x2e, 0xd9, 0x07, 0xb8, 0xab, 0x8d, 0xd2, 0x9e, 0x31, 0x34, 0x47, 0xd6,
	0xf8, 0x69, 0x5d, 0x30, 0x45, 0xf7, 0x60, 0x69, 0xb5, 0xfd, 0x0d, 0x18, 0xbb, 0xc0, 0x5f, 0x79,
	0xd8, 0x5f, 0x0c, 0xd8, 0xa9, 0x12, 0x9c, 0x62, 0x8a, 0xcb, 0xe1, 0x9c, 0xc0, 0x2e, 0x27, 0xb3,
	0x32, 0x7a, 0x10, 0x53, 0x41, 0x78, 0x8e, 0x13, 0x19, 0xc7, 0x1a, 0xf7, 0xef, 0xec, 0xee, 0x44,
	0xdf, 0xa3, 0xdf, 0xd1, 0x92, 0xf7, 0x5a, 0x81, 0x22, 0xf8, 0x88, 0xeb, 0x13, 0x0b, 0xae, 0xd5,
	0x8d, 0x55, 0x91, 0x9d, 0x9a, 0xc8, 0x1b, 0xa7, 0xb9, 0x16, 0xbd, 0xcb, 0xd7, 0xc1, 0x14, 0xbd,
	0x83, 0x4d, 0x2c, 0x67, 0x97, 0xf6, 0x4c, 0x69, 0xff, 0xb2, 0xc6, 0x7e, 0x7d, 0xe2, 0x7e, 0xa5,
	0xf6, 0x2e, 0x7e, 0x2d, 0x07, 0xe0, 0xf7, 0x72, 0x00, 0xfe, 0x2c, 0x07, 0x00, 0x3e, 0x8f, 0x99,
	0xf2, 0x59, 0x70, 0xf6, 0xa9, 0xb8, 0xdf, 0xd2, 0x6b, 0x57, 0x9e, 0xe7, 0xe5, 0x60, 0xce, 0xc1,
	0xc7, 0xa6, 0x46, 0xae, 0x76, 0xe4, 0xa8, 0x5e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x60, 0x67,
	0x99, 0xb6, 0x32, 0x04, 0x00, 0x00,
}
