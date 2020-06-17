// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tls/v3/common.proto

package envoy_extensions_transport_sockets_tls_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type TlsParameters_TlsProtocol int32

const (
	TlsParameters_TLS_AUTO TlsParameters_TlsProtocol = 0
	TlsParameters_TLSv1_0  TlsParameters_TlsProtocol = 1
	TlsParameters_TLSv1_1  TlsParameters_TlsProtocol = 2
	TlsParameters_TLSv1_2  TlsParameters_TlsProtocol = 3
	TlsParameters_TLSv1_3  TlsParameters_TlsProtocol = 4
)

var TlsParameters_TlsProtocol_name = map[int32]string{
	0: "TLS_AUTO",
	1: "TLSv1_0",
	2: "TLSv1_1",
	3: "TLSv1_2",
	4: "TLSv1_3",
}

var TlsParameters_TlsProtocol_value = map[string]int32{
	"TLS_AUTO": 0,
	"TLSv1_0":  1,
	"TLSv1_1":  2,
	"TLSv1_2":  3,
	"TLSv1_3":  4,
}

func (x TlsParameters_TlsProtocol) String() string {
	return proto.EnumName(TlsParameters_TlsProtocol_name, int32(x))
}

func (TlsParameters_TlsProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_09f98a7d9d7974cc, []int{0, 0}
}

type CertificateValidationContext_TrustChainVerification int32

const (
	CertificateValidationContext_VERIFY_TRUST_CHAIN CertificateValidationContext_TrustChainVerification = 0
	CertificateValidationContext_ACCEPT_UNTRUSTED   CertificateValidationContext_TrustChainVerification = 1
)

var CertificateValidationContext_TrustChainVerification_name = map[int32]string{
	0: "VERIFY_TRUST_CHAIN",
	1: "ACCEPT_UNTRUSTED",
}

var CertificateValidationContext_TrustChainVerification_value = map[string]int32{
	"VERIFY_TRUST_CHAIN": 0,
	"ACCEPT_UNTRUSTED":   1,
}

func (x CertificateValidationContext_TrustChainVerification) String() string {
	return proto.EnumName(CertificateValidationContext_TrustChainVerification_name, int32(x))
}

func (CertificateValidationContext_TrustChainVerification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_09f98a7d9d7974cc, []int{4, 0}
}

type TlsParameters struct {
	TlsMinimumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,1,opt,name=tls_minimum_protocol_version,json=tlsMinimumProtocolVersion,proto3,enum=envoy.extensions.transport_sockets.tls.v3.TlsParameters_TlsProtocol" json:"tls_minimum_protocol_version,omitempty"`
	TlsMaximumProtocolVersion TlsParameters_TlsProtocol `protobuf:"varint,2,opt,name=tls_maximum_protocol_version,json=tlsMaximumProtocolVersion,proto3,enum=envoy.extensions.transport_sockets.tls.v3.TlsParameters_TlsProtocol" json:"tls_maximum_protocol_version,omitempty"`
	CipherSuites              []string                  `protobuf:"bytes,3,rep,name=cipher_suites,json=cipherSuites,proto3" json:"cipher_suites,omitempty"`
	EcdhCurves                []string                  `protobuf:"bytes,4,rep,name=ecdh_curves,json=ecdhCurves,proto3" json:"ecdh_curves,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                  `json:"-"`
	XXX_unrecognized          []byte                    `json:"-"`
	XXX_sizecache             int32                     `json:"-"`
}

func (m *TlsParameters) Reset()         { *m = TlsParameters{} }
func (m *TlsParameters) String() string { return proto.CompactTextString(m) }
func (*TlsParameters) ProtoMessage()    {}
func (*TlsParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_09f98a7d9d7974cc, []int{0}
}

func (m *TlsParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsParameters.Unmarshal(m, b)
}
func (m *TlsParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsParameters.Marshal(b, m, deterministic)
}
func (m *TlsParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsParameters.Merge(m, src)
}
func (m *TlsParameters) XXX_Size() int {
	return xxx_messageInfo_TlsParameters.Size(m)
}
func (m *TlsParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsParameters.DiscardUnknown(m)
}

var xxx_messageInfo_TlsParameters proto.InternalMessageInfo

func (m *TlsParameters) GetTlsMinimumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMinimumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetTlsMaximumProtocolVersion() TlsParameters_TlsProtocol {
	if m != nil {
		return m.TlsMaximumProtocolVersion
	}
	return TlsParameters_TLS_AUTO
}

func (m *TlsParameters) GetCipherSuites() []string {
	if m != nil {
		return m.CipherSuites
	}
	return nil
}

func (m *TlsParameters) GetEcdhCurves() []string {
	if m != nil {
		return m.EcdhCurves
	}
	return nil
}

type PrivateKeyProvider struct {
	ProviderName string `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*PrivateKeyProvider_TypedConfig
	//	*PrivateKeyProvider_HiddenEnvoyDeprecatedConfig
	ConfigType           isPrivateKeyProvider_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *PrivateKeyProvider) Reset()         { *m = PrivateKeyProvider{} }
func (m *PrivateKeyProvider) String() string { return proto.CompactTextString(m) }
func (*PrivateKeyProvider) ProtoMessage()    {}
func (*PrivateKeyProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_09f98a7d9d7974cc, []int{1}
}

func (m *PrivateKeyProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrivateKeyProvider.Unmarshal(m, b)
}
func (m *PrivateKeyProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrivateKeyProvider.Marshal(b, m, deterministic)
}
func (m *PrivateKeyProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrivateKeyProvider.Merge(m, src)
}
func (m *PrivateKeyProvider) XXX_Size() int {
	return xxx_messageInfo_PrivateKeyProvider.Size(m)
}
func (m *PrivateKeyProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_PrivateKeyProvider.DiscardUnknown(m)
}

var xxx_messageInfo_PrivateKeyProvider proto.InternalMessageInfo

func (m *PrivateKeyProvider) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

type isPrivateKeyProvider_ConfigType interface {
	isPrivateKeyProvider_ConfigType()
}

type PrivateKeyProvider_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

type PrivateKeyProvider_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

func (*PrivateKeyProvider_TypedConfig) isPrivateKeyProvider_ConfigType() {}

func (*PrivateKeyProvider_HiddenEnvoyDeprecatedConfig) isPrivateKeyProvider_ConfigType() {}

func (m *PrivateKeyProvider) GetConfigType() isPrivateKeyProvider_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *PrivateKeyProvider) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*PrivateKeyProvider_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// Deprecated: Do not use.
func (m *PrivateKeyProvider) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*PrivateKeyProvider_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PrivateKeyProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PrivateKeyProvider_TypedConfig)(nil),
		(*PrivateKeyProvider_HiddenEnvoyDeprecatedConfig)(nil),
	}
}

type TlsCertificate struct {
	CertificateChain           *v3.DataSource      `protobuf:"bytes,1,opt,name=certificate_chain,json=certificateChain,proto3" json:"certificate_chain,omitempty"`
	PrivateKey                 *v3.DataSource      `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	PrivateKeyProvider         *PrivateKeyProvider `protobuf:"bytes,6,opt,name=private_key_provider,json=privateKeyProvider,proto3" json:"private_key_provider,omitempty"`
	Password                   *v3.DataSource      `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	OcspStaple                 *v3.DataSource      `protobuf:"bytes,4,opt,name=ocsp_staple,json=ocspStaple,proto3" json:"ocsp_staple,omitempty"`
	SignedCertificateTimestamp []*v3.DataSource    `protobuf:"bytes,5,rep,name=signed_certificate_timestamp,json=signedCertificateTimestamp,proto3" json:"signed_certificate_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}            `json:"-"`
	XXX_unrecognized           []byte              `json:"-"`
	XXX_sizecache              int32               `json:"-"`
}

func (m *TlsCertificate) Reset()         { *m = TlsCertificate{} }
func (m *TlsCertificate) String() string { return proto.CompactTextString(m) }
func (*TlsCertificate) ProtoMessage()    {}
func (*TlsCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_09f98a7d9d7974cc, []int{2}
}

func (m *TlsCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsCertificate.Unmarshal(m, b)
}
func (m *TlsCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsCertificate.Marshal(b, m, deterministic)
}
func (m *TlsCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsCertificate.Merge(m, src)
}
func (m *TlsCertificate) XXX_Size() int {
	return xxx_messageInfo_TlsCertificate.Size(m)
}
func (m *TlsCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_TlsCertificate proto.InternalMessageInfo

func (m *TlsCertificate) GetCertificateChain() *v3.DataSource {
	if m != nil {
		return m.CertificateChain
	}
	return nil
}

func (m *TlsCertificate) GetPrivateKey() *v3.DataSource {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *TlsCertificate) GetPrivateKeyProvider() *PrivateKeyProvider {
	if m != nil {
		return m.PrivateKeyProvider
	}
	return nil
}

func (m *TlsCertificate) GetPassword() *v3.DataSource {
	if m != nil {
		return m.Password
	}
	return nil
}

func (m *TlsCertificate) GetOcspStaple() *v3.DataSource {
	if m != nil {
		return m.OcspStaple
	}
	return nil
}

func (m *TlsCertificate) GetSignedCertificateTimestamp() []*v3.DataSource {
	if m != nil {
		return m.SignedCertificateTimestamp
	}
	return nil
}

type TlsSessionTicketKeys struct {
	Keys                 []*v3.DataSource `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TlsSessionTicketKeys) Reset()         { *m = TlsSessionTicketKeys{} }
func (m *TlsSessionTicketKeys) String() string { return proto.CompactTextString(m) }
func (*TlsSessionTicketKeys) ProtoMessage()    {}
func (*TlsSessionTicketKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_09f98a7d9d7974cc, []int{3}
}

func (m *TlsSessionTicketKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSessionTicketKeys.Unmarshal(m, b)
}
func (m *TlsSessionTicketKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSessionTicketKeys.Marshal(b, m, deterministic)
}
func (m *TlsSessionTicketKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSessionTicketKeys.Merge(m, src)
}
func (m *TlsSessionTicketKeys) XXX_Size() int {
	return xxx_messageInfo_TlsSessionTicketKeys.Size(m)
}
func (m *TlsSessionTicketKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSessionTicketKeys.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSessionTicketKeys proto.InternalMessageInfo

func (m *TlsSessionTicketKeys) GetKeys() []*v3.DataSource {
	if m != nil {
		return m.Keys
	}
	return nil
}

type CertificateValidationContext struct {
	TrustedCa                                 *v3.DataSource                                      `protobuf:"bytes,1,opt,name=trusted_ca,json=trustedCa,proto3" json:"trusted_ca,omitempty"`
	VerifyCertificateSpki                     []string                                            `protobuf:"bytes,3,rep,name=verify_certificate_spki,json=verifyCertificateSpki,proto3" json:"verify_certificate_spki,omitempty"`
	VerifyCertificateHash                     []string                                            `protobuf:"bytes,2,rep,name=verify_certificate_hash,json=verifyCertificateHash,proto3" json:"verify_certificate_hash,omitempty"`
	MatchSubjectAltNames                      []*v31.StringMatcher                                `protobuf:"bytes,9,rep,name=match_subject_alt_names,json=matchSubjectAltNames,proto3" json:"match_subject_alt_names,omitempty"`
	RequireOcspStaple                         *wrappers.BoolValue                                 `protobuf:"bytes,5,opt,name=require_ocsp_staple,json=requireOcspStaple,proto3" json:"require_ocsp_staple,omitempty"`
	RequireSignedCertificateTimestamp         *wrappers.BoolValue                                 `protobuf:"bytes,6,opt,name=require_signed_certificate_timestamp,json=requireSignedCertificateTimestamp,proto3" json:"require_signed_certificate_timestamp,omitempty"`
	Crl                                       *v3.DataSource                                      `protobuf:"bytes,7,opt,name=crl,proto3" json:"crl,omitempty"`
	AllowExpiredCertificate                   bool                                                `protobuf:"varint,8,opt,name=allow_expired_certificate,json=allowExpiredCertificate,proto3" json:"allow_expired_certificate,omitempty"`
	TrustChainVerification                    CertificateValidationContext_TrustChainVerification `protobuf:"varint,10,opt,name=trust_chain_verification,json=trustChainVerification,proto3,enum=envoy.extensions.transport_sockets.tls.v3.CertificateValidationContext_TrustChainVerification" json:"trust_chain_verification,omitempty"`
	HiddenEnvoyDeprecatedVerifySubjectAltName []string                                            `protobuf:"bytes,4,rep,name=hidden_envoy_deprecated_verify_subject_alt_name,json=hiddenEnvoyDeprecatedVerifySubjectAltName,proto3" json:"hidden_envoy_deprecated_verify_subject_alt_name,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral                      struct{}                                            `json:"-"`
	XXX_unrecognized                          []byte                                              `json:"-"`
	XXX_sizecache                             int32                                               `json:"-"`
}

func (m *CertificateValidationContext) Reset()         { *m = CertificateValidationContext{} }
func (m *CertificateValidationContext) String() string { return proto.CompactTextString(m) }
func (*CertificateValidationContext) ProtoMessage()    {}
func (*CertificateValidationContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_09f98a7d9d7974cc, []int{4}
}

func (m *CertificateValidationContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateValidationContext.Unmarshal(m, b)
}
func (m *CertificateValidationContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateValidationContext.Marshal(b, m, deterministic)
}
func (m *CertificateValidationContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateValidationContext.Merge(m, src)
}
func (m *CertificateValidationContext) XXX_Size() int {
	return xxx_messageInfo_CertificateValidationContext.Size(m)
}
func (m *CertificateValidationContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateValidationContext.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateValidationContext proto.InternalMessageInfo

func (m *CertificateValidationContext) GetTrustedCa() *v3.DataSource {
	if m != nil {
		return m.TrustedCa
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifyCertificateSpki() []string {
	if m != nil {
		return m.VerifyCertificateSpki
	}
	return nil
}

func (m *CertificateValidationContext) GetVerifyCertificateHash() []string {
	if m != nil {
		return m.VerifyCertificateHash
	}
	return nil
}

func (m *CertificateValidationContext) GetMatchSubjectAltNames() []*v31.StringMatcher {
	if m != nil {
		return m.MatchSubjectAltNames
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireOcspStaple() *wrappers.BoolValue {
	if m != nil {
		return m.RequireOcspStaple
	}
	return nil
}

func (m *CertificateValidationContext) GetRequireSignedCertificateTimestamp() *wrappers.BoolValue {
	if m != nil {
		return m.RequireSignedCertificateTimestamp
	}
	return nil
}

func (m *CertificateValidationContext) GetCrl() *v3.DataSource {
	if m != nil {
		return m.Crl
	}
	return nil
}

func (m *CertificateValidationContext) GetAllowExpiredCertificate() bool {
	if m != nil {
		return m.AllowExpiredCertificate
	}
	return false
}

func (m *CertificateValidationContext) GetTrustChainVerification() CertificateValidationContext_TrustChainVerification {
	if m != nil {
		return m.TrustChainVerification
	}
	return CertificateValidationContext_VERIFY_TRUST_CHAIN
}

// Deprecated: Do not use.
func (m *CertificateValidationContext) GetHiddenEnvoyDeprecatedVerifySubjectAltName() []string {
	if m != nil {
		return m.HiddenEnvoyDeprecatedVerifySubjectAltName
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.extensions.transport_sockets.tls.v3.TlsParameters_TlsProtocol", TlsParameters_TlsProtocol_name, TlsParameters_TlsProtocol_value)
	proto.RegisterEnum("envoy.extensions.transport_sockets.tls.v3.CertificateValidationContext_TrustChainVerification", CertificateValidationContext_TrustChainVerification_name, CertificateValidationContext_TrustChainVerification_value)
	proto.RegisterType((*TlsParameters)(nil), "envoy.extensions.transport_sockets.tls.v3.TlsParameters")
	proto.RegisterType((*PrivateKeyProvider)(nil), "envoy.extensions.transport_sockets.tls.v3.PrivateKeyProvider")
	proto.RegisterType((*TlsCertificate)(nil), "envoy.extensions.transport_sockets.tls.v3.TlsCertificate")
	proto.RegisterType((*TlsSessionTicketKeys)(nil), "envoy.extensions.transport_sockets.tls.v3.TlsSessionTicketKeys")
	proto.RegisterType((*CertificateValidationContext)(nil), "envoy.extensions.transport_sockets.tls.v3.CertificateValidationContext")
}

func init() {
	proto.RegisterFile("envoy/extensions/transport_sockets/tls/v3/common.proto", fileDescriptor_09f98a7d9d7974cc)
}

var fileDescriptor_09f98a7d9d7974cc = []byte{
	// 1228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcf, 0x6f, 0x1b, 0xc5,
	0x17, 0xef, 0xda, 0x4e, 0xea, 0x8c, 0x93, 0xca, 0x9d, 0x6f, 0xbe, 0xcd, 0x36, 0x84, 0xd6, 0x35,
	0x55, 0x71, 0x45, 0x59, 0xb7, 0x89, 0x00, 0x29, 0x12, 0x02, 0xaf, 0x93, 0xaa, 0xa5, 0xbf, 0xa2,
	0xf5, 0x36, 0x12, 0x42, 0x62, 0x34, 0x59, 0x4f, 0xe2, 0xc1, 0xeb, 0x9d, 0x65, 0x66, 0x76, 0x1b,
	0xdf, 0x80, 0x13, 0x52, 0x0f, 0x48, 0x3d, 0x72, 0x44, 0xdc, 0x38, 0x71, 0x02, 0x71, 0x47, 0xe2,
	0xca, 0xbf, 0xc2, 0xb1, 0x27, 0x34, 0x33, 0xeb, 0x7a, 0x93, 0x75, 0x5b, 0xf7, 0xc0, 0xcd, 0xb3,
	0xef, 0xbd, 0xcf, 0x7b, 0xf3, 0x99, 0xcf, 0x7b, 0xcf, 0xe0, 0x43, 0x12, 0xa5, 0x6c, 0xdc, 0x26,
	0xc7, 0x92, 0x44, 0x82, 0xb2, 0x48, 0xb4, 0x25, 0xc7, 0x91, 0x88, 0x19, 0x97, 0x48, 0xb0, 0x60,
	0x48, 0xa4, 0x68, 0xcb, 0x50, 0xb4, 0xd3, 0xad, 0x76, 0xc0, 0x46, 0x23, 0x16, 0x39, 0x31, 0x67,
	0x92, 0xc1, 0xeb, 0x3a, 0xce, 0x99, 0xc6, 0x39, 0x85, 0x38, 0x47, 0x86, 0xc2, 0x49, 0xb7, 0xd6,
	0x2f, 0x9b, 0x14, 0x01, 0x8b, 0x0e, 0xe9, 0x51, 0x3b, 0x60, 0x9c, 0x28, 0xb4, 0x03, 0x2c, 0x88,
	0xc1, 0x5a, 0x6f, 0x1a, 0x07, 0x39, 0x8e, 0x49, 0x7b, 0x84, 0x65, 0x30, 0x20, 0x5c, 0x79, 0x08,
	0xc9, 0x69, 0x74, 0x94, 0xf9, 0x5c, 0x3c, 0x62, 0xec, 0x28, 0x24, 0x6d, 0x7d, 0x3a, 0x48, 0x0e,
	0xdb, 0x38, 0x1a, 0x67, 0xa6, 0x8d, 0xd3, 0x26, 0x21, 0x79, 0x12, 0xc8, 0xcc, 0x7a, 0xe9, 0xb4,
	0xf5, 0x09, 0xc7, 0x71, 0x4c, 0xb8, 0xc8, 0xec, 0x8d, 0xa4, 0x1f, 0xe3, 0x36, 0x8e, 0x22, 0x26,
	0xb1, 0xd4, 0x04, 0x08, 0x75, 0x23, 0x49, 0xd3, 0x49, 0x79, 0x6f, 0x17, 0x3d, 0x24, 0x96, 0xc9,
	0x04, 0xe0, 0x4a, 0xc1, 0x9c, 0x12, 0xae, 0x28, 0x99, 0x16, 0xbf, 0x96, 0xe2, 0x90, 0xf6, 0xb1,
	0x24, 0xed, 0xc9, 0x0f, 0x63, 0x68, 0xfe, 0x53, 0x06, 0x2b, 0x7e, 0x28, 0xf6, 0x30, 0xc7, 0x23,
	0x22, 0x09, 0x17, 0xf0, 0x07, 0x0b, 0x6c, 0xc8, 0x50, 0xa0, 0x11, 0x8d, 0xe8, 0x28, 0x19, 0x21,
	0xed, 0x17, 0xb0, 0x10, 0x65, 0xa0, 0xb6, 0xd5, 0xb0, 0x5a, 0xe7, 0x36, 0x77, 0x9c, 0xb9, 0xf9,
	0x77, 0x4e, 0x24, 0xd0, 0xa7, 0x0c, 0xd3, 0xad, 0x3e, 0x77, 0x17, 0xbe, 0xb3, 0x4a, 0x75, 0xcb,
	0xbb, 0x28, 0x43, 0xf1, 0xc0, 0xa4, 0x9c, 0x58, 0xf7, 0x4d, 0xc2, 0x69, 0x45, 0xf8, 0x78, 0x76,
	0x45, 0xa5, 0xff, 0xae, 0x22, 0x93, 0xf2, 0x74, 0x45, 0xef, 0x80, 0x95, 0x80, 0xc6, 0x03, 0xc2,
	0x91, 0x48, 0xa8, 0x24, 0xc2, 0x2e, 0x37, 0xca, 0xad, 0x25, 0x6f, 0xd9, 0x7c, 0xec, 0xe9, 0x6f,
	0xf0, 0x32, 0xa8, 0x91, 0xa0, 0x3f, 0x40, 0x41, 0xc2, 0x53, 0x22, 0xec, 0x8a, 0x76, 0x01, 0xea,
	0x53, 0x57, 0x7f, 0x69, 0x3e, 0x02, 0xb5, 0x5c, 0x66, 0xb8, 0x0c, 0xaa, 0xfe, 0xfd, 0x1e, 0xea,
	0x3c, 0xf6, 0x1f, 0xd5, 0xcf, 0xc0, 0x1a, 0x38, 0xeb, 0xdf, 0xef, 0xa5, 0xb7, 0xd0, 0xcd, 0xba,
	0x35, 0x3d, 0xdc, 0xaa, 0x97, 0xa6, 0x87, 0xcd, 0x7a, 0x79, 0x7a, 0xd8, 0xaa, 0x57, 0xb6, 0xaf,
	0xfd, 0xf8, 0xe7, 0xf7, 0x97, 0xae, 0x00, 0x23, 0x77, 0x07, 0xc7, 0xd4, 0x49, 0x37, 0x1d, 0x9c,
	0xc8, 0xc1, 0xc9, 0xfb, 0x36, 0x7f, 0x29, 0x01, 0xb8, 0xc7, 0x69, 0x8a, 0x25, 0xb9, 0x47, 0xc6,
	0x7b, 0x9c, 0xa5, 0xb4, 0x4f, 0x38, 0xbc, 0x01, 0x56, 0xe2, 0xec, 0x37, 0x8a, 0xf0, 0x88, 0xe8,
	0x97, 0x5e, 0x72, 0xcf, 0x3e, 0x77, 0x2b, 0xbc, 0xd4, 0xb0, 0xbc, 0xe5, 0x89, 0xf5, 0x21, 0x1e,
	0x11, 0xd8, 0x01, 0xcb, 0xaa, 0x5f, 0xfa, 0xc8, 0xb4, 0x95, 0x5d, 0x6e, 0x58, 0xad, 0xda, 0xe6,
	0xaa, 0x63, 0xd4, 0xee, 0x4c, 0xd4, 0xee, 0x74, 0xa2, 0xb1, 0xbb, 0xf8, 0xfb, 0x6f, 0x4f, 0x7f,
	0x2e, 0x59, 0x77, 0xce, 0x78, 0x35, 0x1d, 0xd3, 0xd5, 0x21, 0x70, 0x00, 0x2e, 0x0d, 0x68, 0xbf,
	0x4f, 0x22, 0xa4, 0x2b, 0x46, 0x7d, 0x12, 0x73, 0x12, 0x60, 0x39, 0x05, 0x2d, 0x69, 0xd0, 0xb5,
	0x02, 0x68, 0x4f, 0x37, 0x98, 0x5b, 0x35, 0xb8, 0xb6, 0x42, 0x7e, 0xcb, 0x40, 0xed, 0x2a, 0xa4,
	0x9d, 0x17, 0x40, 0x26, 0xd3, 0xf6, 0x7b, 0x8a, 0x99, 0x6b, 0xe0, 0x6a, 0x91, 0x99, 0x22, 0x0f,
	0xee, 0x0a, 0xa8, 0x99, 0xf4, 0x48, 0x15, 0xdb, 0xfc, 0xb5, 0x02, 0xce, 0xf9, 0xa1, 0xe8, 0x12,
	0x2e, 0xe9, 0x21, 0x55, 0xa8, 0xf0, 0x01, 0x38, 0x1f, 0x4c, 0x8f, 0x28, 0x18, 0x60, 0x6a, 0xfa,
	0xa2, 0xb6, 0xd9, 0xc8, 0x54, 0x68, 0x10, 0x1c, 0x35, 0x6c, 0x94, 0xe0, 0x76, 0xb0, 0xc4, 0x3d,
	0x96, 0xf0, 0x80, 0x78, 0xf5, 0x5c, 0x68, 0x57, 0x45, 0xc2, 0xbb, 0xa0, 0x16, 0x9b, 0x32, 0xd0,
	0x90, 0x8c, 0xb3, 0x4b, 0xbf, 0x16, 0x68, 0xc2, 0xaa, 0x07, 0xe2, 0x17, 0x77, 0x80, 0x0c, 0xac,
	0xe6, 0xa0, 0xd0, 0xe4, 0xc5, 0xec, 0x45, 0x8d, 0xf9, 0xf1, 0x1b, 0xb4, 0x48, 0x91, 0x18, 0x0f,
	0xc6, 0x45, 0xd1, 0xec, 0x80, 0x6a, 0x8c, 0x85, 0x78, 0xc2, 0x78, 0x3f, 0x93, 0xc0, 0xfc, 0x85,
	0xbf, 0x88, 0x84, 0x1d, 0x50, 0x63, 0x81, 0x88, 0x91, 0x90, 0x38, 0x0e, 0x89, 0x5d, 0x99, 0x93,
	0x4a, 0xa0, 0x82, 0x7a, 0x3a, 0x06, 0x1e, 0x80, 0x0d, 0x41, 0x8f, 0x22, 0xa5, 0x9d, 0xdc, 0xd3,
	0x48, 0x3a, 0x22, 0x42, 0xe2, 0x51, 0x6c, 0x2f, 0x34, 0xca, 0x73, 0x61, 0xae, 0x1b, 0x94, 0xdc,
	0x73, 0xfb, 0x13, 0x8c, 0xed, 0x77, 0x95, 0x8c, 0x9a, 0xa0, 0x31, 0xb3, 0xc1, 0x72, 0x11, 0xcd,
	0xa7, 0x16, 0x58, 0xf5, 0x43, 0xd1, 0x23, 0x42, 0xb1, 0xec, 0x53, 0xc5, 0xec, 0x3d, 0x32, 0x16,
	0x70, 0x07, 0x54, 0x86, 0x64, 0x2c, 0x6c, 0x6b, 0xbe, 0x6a, 0xdc, 0x73, 0xcf, 0xdd, 0x85, 0x67,
	0x56, 0xa9, 0x6a, 0x65, 0x94, 0xe9, 0xe8, 0xed, 0xf7, 0x55, 0x1d, 0x2d, 0x70, 0x6d, 0x66, 0x1d,
	0x85, 0xa4, 0xcd, 0x6f, 0xab, 0x60, 0x23, 0x57, 0xdd, 0xbe, 0x59, 0x01, 0x94, 0x45, 0x5d, 0x16,
	0x49, 0x72, 0x2c, 0xe1, 0x27, 0x00, 0x48, 0x9e, 0x08, 0xdd, 0x78, 0x78, 0x6e, 0x21, 0x2f, 0x65,
	0x31, 0x5d, 0x0c, 0x6f, 0x83, 0xb5, 0x94, 0x70, 0x7a, 0x38, 0x3e, 0x41, 0xbe, 0x88, 0x87, 0xd4,
	0x8c, 0x46, 0x75, 0x8f, 0xda, 0x33, 0xab, 0xda, 0x5c, 0xe4, 0x95, 0xc6, 0x8d, 0xd6, 0x0d, 0xef,
	0xff, 0xc6, 0x3d, 0x57, 0x55, 0x2f, 0x1e, 0xd2, 0x97, 0xe0, 0x0c, 0xb0, 0x18, 0xd8, 0xa5, 0x02,
	0xce, 0xa7, 0x2d, 0x34, 0x03, 0xe7, 0x0e, 0x16, 0x03, 0xf8, 0x05, 0x58, 0xd3, 0x7b, 0x1c, 0x89,
	0xe4, 0xe0, 0x2b, 0x12, 0x48, 0x84, 0x43, 0xa9, 0x67, 0x9a, 0xb0, 0x97, 0x34, 0xf3, 0x57, 0xb3,
	0xdb, 0xa9, 0x0e, 0x77, 0xb2, 0x95, 0xaf, 0xae, 0xd7, 0xd3, 0x2b, 0xff, 0x81, 0xf9, 0xe0, 0xad,
	0x6a, 0x4b, 0xcf, 0x60, 0x74, 0x42, 0xa9, 0x06, 0x9f, 0x80, 0x9f, 0x81, 0xff, 0x71, 0xf2, 0x75,
	0x42, 0x39, 0x41, 0x79, 0xd1, 0x2e, 0x68, 0xda, 0xd6, 0x0b, 0xb3, 0xca, 0x65, 0x2c, 0xdc, 0xc7,
	0x61, 0x42, 0xbc, 0xf3, 0x59, 0xd8, 0xa3, 0xa9, 0x6a, 0x87, 0xe0, 0xea, 0x04, 0xeb, 0x95, 0xea,
	0x5d, 0x7c, 0x2d, 0xf8, 0x95, 0x0c, 0xa7, 0xf7, 0x52, 0xf9, 0xc2, 0x4d, 0x50, 0x0e, 0x78, 0x68,
	0x9f, 0x9d, 0xf3, 0x7d, 0x95, 0x33, 0xdc, 0x06, 0x17, 0x71, 0x18, 0xb2, 0x27, 0x88, 0x1c, 0xc7,
	0x94, 0x9f, 0xac, 0xcf, 0xae, 0x36, 0xac, 0x56, 0xd5, 0x5b, 0xd3, 0x0e, 0xbb, 0xc6, 0x9e, 0x1f,
	0x93, 0x3f, 0x59, 0xc0, 0xd6, 0x1a, 0x31, 0x13, 0x12, 0xe9, 0xb7, 0x52, 0x26, 0xb5, 0xb4, 0x81,
	0x5e, 0xda, 0x5f, 0xbe, 0xc1, 0x44, 0x7a, 0x95, 0x84, 0x1d, 0x5f, 0xe5, 0xd1, 0xe3, 0x74, 0x3f,
	0x97, 0x25, 0xb7, 0xce, 0x2f, 0xc8, 0x99, 0x1e, 0xf0, 0x10, 0xb4, 0x5f, 0xb6, 0x84, 0x32, 0x29,
	0x9e, 0xd6, 0x90, 0x59, 0xe5, 0x6e, 0xc9, 0xb6, 0xbc, 0xeb, 0x33, 0x97, 0x8e, 0xc6, 0x1f, 0x9f,
	0x94, 0x4d, 0xf3, 0x36, 0xb8, 0x30, 0xbb, 0x46, 0x78, 0x01, 0xc0, 0xfd, 0x5d, 0xef, 0xee, 0xed,
	0xcf, 0x91, 0xef, 0x3d, 0xee, 0xf9, 0xa8, 0x7b, 0xa7, 0x73, 0xf7, 0x61, 0xfd, 0x0c, 0x5c, 0x05,
	0xf5, 0x4e, 0xb7, 0xbb, 0xbb, 0xe7, 0xa3, 0xc7, 0x0f, 0xb5, 0x65, 0x77, 0xa7, 0x6e, 0x6d, 0x7f,
	0xa0, 0x7a, 0xff, 0x26, 0x70, 0x8a, 0xbd, 0xff, 0x2a, 0x7e, 0xdc, 0xbd, 0x3f, 0xbe, 0xf9, 0xeb,
	0xef, 0xc5, 0x52, 0xbd, 0x04, 0x3e, 0xa2, 0xcc, 0x04, 0xc7, 0x9c, 0x1d, 0x8f, 0xe7, 0xe7, 0xdf,
	0xad, 0x75, 0xf5, 0xff, 0x6f, 0xfd, 0x87, 0x65, 0xcf, 0x3a, 0x58, 0xd4, 0x22, 0xdc, 0xfa, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x3f, 0x4b, 0x2e, 0x3f, 0xc1, 0x0b, 0x00, 0x00,
}
