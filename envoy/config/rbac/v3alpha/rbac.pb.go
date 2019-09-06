// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v3alpha/rbac.proto

package envoy_config_rbac_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/v2/envoy/api/v3alpha/core"
	route "github.com/envoyproxy/go-control-plane/v2/envoy/api/v3alpha/route"
	matcher "github.com/envoyproxy/go-control-plane/v2/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RBAC_Action int32

const (
	RBAC_ALLOW RBAC_Action = 0
	RBAC_DENY  RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}

func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{0, 0}
}

type RBAC struct {
	Action               RBAC_Action        `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.rbac.v3alpha.RBAC_Action" json:"action,omitempty"`
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type Policy struct {
	Permissions          []*Permission  `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Principals           []*Principal   `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	Condition            *v1alpha1.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetCondition() *v1alpha1.Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

type Permission struct {
	// Types that are valid to be assigned to Rule:
	//	*Permission_AndRules
	//	*Permission_OrRules
	//	*Permission_Any
	//	*Permission_Header
	//	*Permission_DestinationIp
	//	*Permission_DestinationPort
	//	*Permission_Metadata
	//	*Permission_NotRule
	//	*Permission_RequestedServerName
	Rule                 isPermission_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{2}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

type isPermission_Rule interface {
	isPermission_Rule()
}

type Permission_AndRules struct {
	AndRules *Permission_Set `protobuf:"bytes,1,opt,name=and_rules,json=andRules,proto3,oneof"`
}

type Permission_OrRules struct {
	OrRules *Permission_Set `protobuf:"bytes,2,opt,name=or_rules,json=orRules,proto3,oneof"`
}

type Permission_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Permission_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type Permission_DestinationIp struct {
	DestinationIp *core.CidrRange `protobuf:"bytes,5,opt,name=destination_ip,json=destinationIp,proto3,oneof"`
}

type Permission_DestinationPort struct {
	DestinationPort uint32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,proto3,oneof"`
}

type Permission_Metadata struct {
	Metadata *matcher.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Permission_NotRule struct {
	NotRule *Permission `protobuf:"bytes,8,opt,name=not_rule,json=notRule,proto3,oneof"`
}

type Permission_RequestedServerName struct {
	RequestedServerName *matcher.StringMatcher `protobuf:"bytes,9,opt,name=requested_server_name,json=requestedServerName,proto3,oneof"`
}

func (*Permission_AndRules) isPermission_Rule() {}

func (*Permission_OrRules) isPermission_Rule() {}

func (*Permission_Any) isPermission_Rule() {}

func (*Permission_Header) isPermission_Rule() {}

func (*Permission_DestinationIp) isPermission_Rule() {}

func (*Permission_DestinationPort) isPermission_Rule() {}

func (*Permission_Metadata) isPermission_Rule() {}

func (*Permission_NotRule) isPermission_Rule() {}

func (*Permission_RequestedServerName) isPermission_Rule() {}

func (m *Permission) GetRule() isPermission_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Permission) GetAndRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_AndRules); ok {
		return x.AndRules
	}
	return nil
}

func (m *Permission) GetOrRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_OrRules); ok {
		return x.OrRules
	}
	return nil
}

func (m *Permission) GetAny() bool {
	if x, ok := m.GetRule().(*Permission_Any); ok {
		return x.Any
	}
	return false
}

func (m *Permission) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetRule().(*Permission_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission) GetDestinationIp() *core.CidrRange {
	if x, ok := m.GetRule().(*Permission_DestinationIp); ok {
		return x.DestinationIp
	}
	return nil
}

func (m *Permission) GetDestinationPort() uint32 {
	if x, ok := m.GetRule().(*Permission_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *Permission) GetMetadata() *matcher.MetadataMatcher {
	if x, ok := m.GetRule().(*Permission_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Permission) GetNotRule() *Permission {
	if x, ok := m.GetRule().(*Permission_NotRule); ok {
		return x.NotRule
	}
	return nil
}

func (m *Permission) GetRequestedServerName() *matcher.StringMatcher {
	if x, ok := m.GetRule().(*Permission_RequestedServerName); ok {
		return x.RequestedServerName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Permission) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Permission_AndRules)(nil),
		(*Permission_OrRules)(nil),
		(*Permission_Any)(nil),
		(*Permission_Header)(nil),
		(*Permission_DestinationIp)(nil),
		(*Permission_DestinationPort)(nil),
		(*Permission_Metadata)(nil),
		(*Permission_NotRule)(nil),
		(*Permission_RequestedServerName)(nil),
	}
}

type Permission_Set struct {
	Rules                []*Permission `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Permission_Set) Reset()         { *m = Permission_Set{} }
func (m *Permission_Set) String() string { return proto.CompactTextString(m) }
func (*Permission_Set) ProtoMessage()    {}
func (*Permission_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{2, 0}
}

func (m *Permission_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Set.Unmarshal(m, b)
}
func (m *Permission_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Set.Marshal(b, m, deterministic)
}
func (m *Permission_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Set.Merge(m, src)
}
func (m *Permission_Set) XXX_Size() int {
	return xxx_messageInfo_Permission_Set.Size(m)
}
func (m *Permission_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Set proto.InternalMessageInfo

func (m *Permission_Set) GetRules() []*Permission {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Principal struct {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds
	//	*Principal_OrIds
	//	*Principal_Any
	//	*Principal_Authenticated_
	//	*Principal_SourceIp
	//	*Principal_Header
	//	*Principal_Metadata
	//	*Principal_NotId
	Identifier           isPrincipal_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{3}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

type isPrincipal_Identifier interface {
	isPrincipal_Identifier()
}

type Principal_AndIds struct {
	AndIds *Principal_Set `protobuf:"bytes,1,opt,name=and_ids,json=andIds,proto3,oneof"`
}

type Principal_OrIds struct {
	OrIds *Principal_Set `protobuf:"bytes,2,opt,name=or_ids,json=orIds,proto3,oneof"`
}

type Principal_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Principal_Authenticated_ struct {
	Authenticated *Principal_Authenticated `protobuf:"bytes,4,opt,name=authenticated,proto3,oneof"`
}

type Principal_SourceIp struct {
	SourceIp *core.CidrRange `protobuf:"bytes,5,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

type Principal_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,6,opt,name=header,proto3,oneof"`
}

type Principal_Metadata struct {
	Metadata *matcher.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Principal_NotId struct {
	NotId *Principal `protobuf:"bytes,8,opt,name=not_id,json=notId,proto3,oneof"`
}

func (*Principal_AndIds) isPrincipal_Identifier() {}

func (*Principal_OrIds) isPrincipal_Identifier() {}

func (*Principal_Any) isPrincipal_Identifier() {}

func (*Principal_Authenticated_) isPrincipal_Identifier() {}

func (*Principal_SourceIp) isPrincipal_Identifier() {}

func (*Principal_Header) isPrincipal_Identifier() {}

func (*Principal_Metadata) isPrincipal_Identifier() {}

func (*Principal_NotId) isPrincipal_Identifier() {}

func (m *Principal) GetIdentifier() isPrincipal_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Principal) GetAndIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_AndIds); ok {
		return x.AndIds
	}
	return nil
}

func (m *Principal) GetOrIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_OrIds); ok {
		return x.OrIds
	}
	return nil
}

func (m *Principal) GetAny() bool {
	if x, ok := m.GetIdentifier().(*Principal_Any); ok {
		return x.Any
	}
	return false
}

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if x, ok := m.GetIdentifier().(*Principal_Authenticated_); ok {
		return x.Authenticated
	}
	return nil
}

func (m *Principal) GetSourceIp() *core.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_SourceIp); ok {
		return x.SourceIp
	}
	return nil
}

func (m *Principal) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Principal) GetMetadata() *matcher.MetadataMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Principal) GetNotId() *Principal {
	if x, ok := m.GetIdentifier().(*Principal_NotId); ok {
		return x.NotId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Principal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Principal_AndIds)(nil),
		(*Principal_OrIds)(nil),
		(*Principal_Any)(nil),
		(*Principal_Authenticated_)(nil),
		(*Principal_SourceIp)(nil),
		(*Principal_Header)(nil),
		(*Principal_Metadata)(nil),
		(*Principal_NotId)(nil),
	}
}

type Principal_Set struct {
	Ids                  []*Principal `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Principal_Set) Reset()         { *m = Principal_Set{} }
func (m *Principal_Set) String() string { return proto.CompactTextString(m) }
func (*Principal_Set) ProtoMessage()    {}
func (*Principal_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{3, 0}
}

func (m *Principal_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Set.Unmarshal(m, b)
}
func (m *Principal_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Set.Marshal(b, m, deterministic)
}
func (m *Principal_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Set.Merge(m, src)
}
func (m *Principal_Set) XXX_Size() int {
	return xxx_messageInfo_Principal_Set.Size(m)
}
func (m *Principal_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Set proto.InternalMessageInfo

func (m *Principal_Set) GetIds() []*Principal {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Principal_Authenticated struct {
	PrincipalName        *matcher.StringMatcher `protobuf:"bytes,2,opt,name=principal_name,json=principalName,proto3" json:"principal_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{3, 1}
}

func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (m *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(m, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetPrincipalName() *matcher.StringMatcher {
	if m != nil {
		return m.PrincipalName
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.rbac.v3alpha.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v3alpha.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v3alpha.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v3alpha.Policy")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v3alpha.Permission")
	proto.RegisterType((*Permission_Set)(nil), "envoy.config.rbac.v3alpha.Permission.Set")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v3alpha.Principal")
	proto.RegisterType((*Principal_Set)(nil), "envoy.config.rbac.v3alpha.Principal.Set")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v3alpha.Principal.Authenticated")
}

func init() {
	proto.RegisterFile("envoy/config/rbac/v3alpha/rbac.proto", fileDescriptor_eb7d22ec66d90520)
}

var fileDescriptor_eb7d22ec66d90520 = []byte{
	// 896 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0xe3, 0x44,
	0x1b, 0x8d, 0x9d, 0xc4, 0x75, 0xbe, 0x28, 0xfd, 0xf3, 0x0f, 0x42, 0x98, 0x20, 0xa0, 0x9b, 0xdd,
	0x85, 0x80, 0x84, 0xad, 0xcd, 0x4a, 0x0b, 0x42, 0x80, 0x1a, 0x97, 0x42, 0xb2, 0xea, 0x96, 0xe2,
	0x5e, 0xac, 0xe0, 0x82, 0x6a, 0xea, 0x99, 0x4d, 0x07, 0x92, 0x19, 0x33, 0x9e, 0x44, 0xcd, 0x5b,
	0x20, 0x9e, 0x87, 0xd7, 0xe0, 0x21, 0x78, 0x03, 0xd4, 0x9b, 0x45, 0x33, 0x63, 0xa7, 0x8e, 0x60,
	0x4b, 0x5a, 0x71, 0x13, 0x4d, 0xec, 0x73, 0xce, 0xf7, 0xcd, 0x37, 0xe7, 0x8c, 0xe1, 0x01, 0xe5,
	0x4b, 0xb1, 0x8a, 0x52, 0xc1, 0x5f, 0xb0, 0x69, 0x24, 0xcf, 0x71, 0x1a, 0x2d, 0x1f, 0xe3, 0x59,
	0x76, 0x81, 0xcd, 0x9f, 0x30, 0x93, 0x42, 0x09, 0xf4, 0xa6, 0x41, 0x85, 0x16, 0x15, 0x9a, 0x17,
	0x05, 0xaa, 0xf7, 0xc6, 0x12, 0xcf, 0x18, 0xc1, 0x8a, 0x46, 0xe5, 0xc2, 0x72, 0x7a, 0x85, 0x32,
	0xce, 0xd8, 0x5a, 0x31, 0x15, 0x92, 0x46, 0x98, 0x10, 0x49, 0xf3, 0xbc, 0x40, 0xdd, 0xff, 0x3b,
	0x4a, 0x8a, 0x85, 0xa2, 0xf6, 0xb7, 0x00, 0xdd, 0xb3, 0x20, 0xb5, 0xca, 0x68, 0x34, 0xc7, 0x2a,
	0xbd, 0xa0, 0x32, 0x9a, 0x53, 0x85, 0x09, 0x56, 0xb8, 0x80, 0xbc, 0xfb, 0x0f, 0x90, 0x5c, 0x49,
	0xc6, 0xa7, 0x05, 0xe0, 0xe1, 0x54, 0x88, 0xe9, 0x8c, 0x9a, 0x4a, 0xf4, 0x32, 0x93, 0xd1, 0xf2,
	0x91, 0x29, 0xf7, 0x28, 0xca, 0x57, 0x5c, 0xe1, 0x4b, 0x0b, 0xeb, 0xff, 0xe2, 0x42, 0x23, 0x89,
	0x47, 0x07, 0xe8, 0x0b, 0xf0, 0x70, 0xaa, 0x98, 0xe0, 0x81, 0xb3, 0xe7, 0x0c, 0x76, 0x87, 0xef,
	0x85, 0xaf, 0x9c, 0x41, 0xa8, 0x09, 0xe1, 0xc8, 0xa0, 0x93, 0x82, 0x85, 0x26, 0xe0, 0x67, 0x62,
	0xc6, 0x52, 0x46, 0xf3, 0xc0, 0xdd, 0xab, 0x0f, 0xda, 0xc3, 0x8f, 0xfe, 0x4d, 0xe1, 0xa4, 0xc0,
	0x1f, 0x72, 0x25, 0x57, 0xc9, 0x9a, 0xde, 0xfb, 0x01, 0x3a, 0x1b, 0xaf, 0x50, 0x17, 0xea, 0x3f,
	0xd1, 0x95, 0x69, 0xac, 0x95, 0xe8, 0x25, 0xfa, 0x18, 0x9a, 0x4b, 0x3c, 0x5b, 0xd0, 0xc0, 0xdd,
	0x73, 0x06, 0xed, 0xe1, 0xbd, 0x1b, 0x4a, 0x19, 0xa9, 0x55, 0x62, 0xf1, 0x9f, 0xba, 0x9f, 0x38,
	0xfd, 0xb7, 0xc1, 0xb3, 0xcd, 0xa3, 0x16, 0x34, 0x47, 0x47, 0x47, 0xdf, 0x3c, 0xef, 0xd6, 0x90,
	0x0f, 0x8d, 0x2f, 0x0f, 0x8f, 0xbf, 0xeb, 0x3a, 0xfd, 0x3f, 0x1c, 0xf0, 0x2c, 0x09, 0x7d, 0x0b,
	0xed, 0x8c, 0xca, 0x39, 0xcb, 0x73, 0x26, 0x78, 0x1e, 0x38, 0x66, 0x5f, 0x0f, 0x6f, 0x2a, 0xb6,
	0x46, 0xc7, 0xfe, 0x55, 0xdc, 0xfc, 0xd5, 0x71, 0x7d, 0x27, 0xa9, 0x6a, 0xa0, 0x63, 0x80, 0x4c,
	0x32, 0x9e, 0xb2, 0x0c, 0xcf, 0xca, 0x49, 0x3d, 0xb8, 0x49, 0xb1, 0x04, 0x57, 0x04, 0x2b, 0x0a,
	0xe8, 0x33, 0x68, 0xa5, 0x82, 0x13, 0x66, 0x8e, 0xae, 0x6e, 0xa6, 0xf1, 0x4e, 0x68, 0xcf, 0x3e,
	0xc4, 0x19, 0x0b, 0xf5, 0xd9, 0x87, 0xe5, 0xd9, 0x87, 0x87, 0x97, 0x99, 0x4c, 0xae, 0x09, 0xfd,
	0xdf, 0x9a, 0x00, 0xd7, 0x3d, 0xa3, 0x31, 0xb4, 0x30, 0x27, 0x67, 0x72, 0x31, 0xa3, 0xb9, 0x19,
	0x77, 0x7b, 0xf8, 0xc1, 0x56, 0xbb, 0x0d, 0x4f, 0xa9, 0x1a, 0xd7, 0x12, 0x1f, 0x73, 0x92, 0x68,
	0x32, 0xfa, 0x0a, 0x7c, 0x21, 0x0b, 0x21, 0xf7, 0xf6, 0x42, 0x3b, 0x42, 0x5a, 0x9d, 0xb7, 0xa0,
	0x8e, 0xf9, 0xca, 0x6c, 0xcc, 0x8f, 0x77, 0xae, 0xe2, 0xc6, 0x8f, 0xae, 0xef, 0x8c, 0x6b, 0x89,
	0x7e, 0x8a, 0xf6, 0xc1, 0xbb, 0xa0, 0x98, 0x50, 0x19, 0x34, 0x4c, 0x89, 0xd2, 0xb3, 0x7a, 0xdf,
	0xa5, 0xb4, 0xcd, 0xd5, 0xd8, 0xc0, 0x9e, 0xd9, 0xa0, 0x8c, 0x6b, 0x49, 0xc1, 0x43, 0x4f, 0x61,
	0x97, 0xd0, 0x5c, 0x31, 0x8e, 0xf5, 0x38, 0xce, 0x58, 0x16, 0x34, 0x37, 0x0c, 0x55, 0x55, 0xd2,
	0x69, 0x0e, 0x0f, 0x18, 0x91, 0x09, 0xe6, 0x53, 0x3a, 0xae, 0x25, 0x9d, 0x0a, 0x75, 0x92, 0xa1,
	0x27, 0xd0, 0xad, 0x6a, 0x65, 0x42, 0xaa, 0xc0, 0xdb, 0x73, 0x06, 0x9d, 0xb8, 0x75, 0x15, 0x7b,
	0x1f, 0x36, 0x82, 0x97, 0x2f, 0xeb, 0xe3, 0x5a, 0xf2, 0xbf, 0x0a, 0xe8, 0x44, 0x48, 0x85, 0x46,
	0xe0, 0x97, 0xe1, 0x0e, 0x76, 0x4c, 0xf5, 0xfb, 0x45, 0x75, 0x9d, 0xee, 0xb0, 0x48, 0x77, 0xf8,
	0xac, 0xc0, 0x5c, 0x6f, 0x62, 0x4d, 0x43, 0x31, 0xf8, 0x5c, 0x28, 0x33, 0xee, 0xc0, 0x37, 0x12,
	0xdb, 0x99, 0x54, 0x4f, 0x9a, 0x0b, 0xa5, 0x47, 0x8d, 0x9e, 0xc3, 0xeb, 0x92, 0xfe, 0xbc, 0xa0,
	0xb9, 0xa2, 0xe4, 0x2c, 0xa7, 0x72, 0x49, 0xe5, 0x19, 0xc7, 0x73, 0x1a, 0xb4, 0x36, 0x26, 0xb2,
	0xd1, 0xd3, 0xa9, 0xb9, 0x71, 0xae, 0x3b, 0x7a, 0x6d, 0xad, 0x70, 0x6a, 0x04, 0x8e, 0xf1, 0x9c,
	0xf6, 0x8e, 0xa0, 0x7e, 0x4a, 0x15, 0x3a, 0x84, 0x66, 0xe9, 0xab, 0x3b, 0xa5, 0xc8, 0xb2, 0xe3,
	0x36, 0x34, 0xf4, 0x02, 0xd5, 0xff, 0x8c, 0x9d, 0xfe, 0xef, 0x4d, 0x68, 0xad, 0x03, 0x82, 0x0e,
	0x60, 0x47, 0xbb, 0x97, 0x91, 0xd2, 0xbb, 0x83, 0x6d, 0x72, 0x55, 0x38, 0xce, 0xc3, 0x9c, 0x4c,
	0x48, 0x8e, 0x46, 0xe0, 0x09, 0x69, 0x34, 0xdc, 0x5b, 0x6b, 0x34, 0x85, 0xd4, 0x12, 0x37, 0x7a,
	0xf6, 0x7b, 0xe8, 0xe0, 0x85, 0xba, 0xa0, 0x5c, 0xb1, 0x14, 0x2b, 0x4a, 0x0a, 0xeb, 0x0e, 0xb7,
	0x2a, 0x33, 0xaa, 0x32, 0xb5, 0x03, 0x37, 0xa4, 0xd0, 0x3e, 0xb4, 0x72, 0xb1, 0x90, 0x29, 0xbd,
	0xa5, 0x91, 0x7d, 0xcb, 0x9a, 0x64, 0x95, 0x44, 0x79, 0x77, 0x4c, 0xd4, 0x7f, 0xe0, 0xe6, 0xcf,
	0xc1, 0xd3, 0x6e, 0x66, 0xa4, 0xf0, 0xf2, 0x56, 0xd7, 0xa3, 0x1e, 0x3f, 0x17, 0x6a, 0x42, 0x7a,
	0x5f, 0x5b, 0xbf, 0xed, 0x43, 0xdd, 0x3a, 0xe1, 0x2e, 0x37, 0xac, 0xa6, 0xf6, 0x28, 0x74, 0x36,
	0x06, 0x8e, 0xc6, 0xb0, 0xbb, 0xbe, 0x79, 0x6d, 0x36, 0xdc, 0x2d, 0xb3, 0x91, 0x74, 0xd6, 0x44,
	0x9d, 0x89, 0xa7, 0x0d, 0xdf, 0xe9, 0xba, 0x49, 0x43, 0x6b, 0xc4, 0xff, 0x07, 0x60, 0x44, 0x17,
	0x79, 0xc1, 0xa8, 0x34, 0xbe, 0x8e, 0x9f, 0xc0, 0xfb, 0x4c, 0x58, 0xd1, 0x4c, 0x8a, 0xcb, 0xd5,
	0xab, 0xbb, 0x8f, 0x5b, 0xc9, 0x39, 0x4e, 0x4f, 0xf4, 0xb7, 0xfc, 0xc4, 0x39, 0xf7, 0xcc, 0x47,
	0xfd, 0xf1, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x6a, 0x48, 0xfe, 0xe6, 0x08, 0x00, 0x00,
}
