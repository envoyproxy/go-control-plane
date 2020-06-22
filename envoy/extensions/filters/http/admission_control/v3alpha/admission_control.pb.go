// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.10.1
// source: envoy/extensions/filters/http/admission_control/v3alpha/admission_control.proto

package envoy_extensions_filters_http_admission_control_v3alpha

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AdmissionControl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *v3.RuntimeFeatureFlag `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Types that are assignable to EvaluationCriteria:
	//	*AdmissionControl_SuccessCriteria_
	EvaluationCriteria    isAdmissionControl_EvaluationCriteria `protobuf_oneof:"evaluation_criteria"`
	SamplingWindow        *duration.Duration                    `protobuf:"bytes,3,opt,name=sampling_window,json=samplingWindow,proto3" json:"sampling_window,omitempty"`
	AggressionCoefficient *v3.RuntimeDouble                     `protobuf:"bytes,4,opt,name=aggression_coefficient,json=aggressionCoefficient,proto3" json:"aggression_coefficient,omitempty"`
}

func (x *AdmissionControl) Reset() {
	*x = AdmissionControl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionControl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionControl) ProtoMessage() {}

func (x *AdmissionControl) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionControl.ProtoReflect.Descriptor instead.
func (*AdmissionControl) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescGZIP(), []int{0}
}

func (x *AdmissionControl) GetEnabled() *v3.RuntimeFeatureFlag {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (m *AdmissionControl) GetEvaluationCriteria() isAdmissionControl_EvaluationCriteria {
	if m != nil {
		return m.EvaluationCriteria
	}
	return nil
}

func (x *AdmissionControl) GetSuccessCriteria() *AdmissionControl_SuccessCriteria {
	if x, ok := x.GetEvaluationCriteria().(*AdmissionControl_SuccessCriteria_); ok {
		return x.SuccessCriteria
	}
	return nil
}

func (x *AdmissionControl) GetSamplingWindow() *duration.Duration {
	if x != nil {
		return x.SamplingWindow
	}
	return nil
}

func (x *AdmissionControl) GetAggressionCoefficient() *v3.RuntimeDouble {
	if x != nil {
		return x.AggressionCoefficient
	}
	return nil
}

type isAdmissionControl_EvaluationCriteria interface {
	isAdmissionControl_EvaluationCriteria()
}

type AdmissionControl_SuccessCriteria_ struct {
	SuccessCriteria *AdmissionControl_SuccessCriteria `protobuf:"bytes,2,opt,name=success_criteria,json=successCriteria,proto3,oneof"`
}

func (*AdmissionControl_SuccessCriteria_) isAdmissionControl_EvaluationCriteria() {}

type AdmissionControl_SuccessCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpCriteria *AdmissionControl_SuccessCriteria_HttpCriteria `protobuf:"bytes,1,opt,name=http_criteria,json=httpCriteria,proto3" json:"http_criteria,omitempty"`
	GrpcCriteria *AdmissionControl_SuccessCriteria_GrpcCriteria `protobuf:"bytes,2,opt,name=grpc_criteria,json=grpcCriteria,proto3" json:"grpc_criteria,omitempty"`
}

func (x *AdmissionControl_SuccessCriteria) Reset() {
	*x = AdmissionControl_SuccessCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionControl_SuccessCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionControl_SuccessCriteria) ProtoMessage() {}

func (x *AdmissionControl_SuccessCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionControl_SuccessCriteria.ProtoReflect.Descriptor instead.
func (*AdmissionControl_SuccessCriteria) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescGZIP(), []int{0, 0}
}

func (x *AdmissionControl_SuccessCriteria) GetHttpCriteria() *AdmissionControl_SuccessCriteria_HttpCriteria {
	if x != nil {
		return x.HttpCriteria
	}
	return nil
}

func (x *AdmissionControl_SuccessCriteria) GetGrpcCriteria() *AdmissionControl_SuccessCriteria_GrpcCriteria {
	if x != nil {
		return x.GrpcCriteria
	}
	return nil
}

type AdmissionControl_SuccessCriteria_HttpCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpSuccessStatus []*v31.Int32Range `protobuf:"bytes,1,rep,name=http_success_status,json=httpSuccessStatus,proto3" json:"http_success_status,omitempty"`
}

func (x *AdmissionControl_SuccessCriteria_HttpCriteria) Reset() {
	*x = AdmissionControl_SuccessCriteria_HttpCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionControl_SuccessCriteria_HttpCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionControl_SuccessCriteria_HttpCriteria) ProtoMessage() {}

func (x *AdmissionControl_SuccessCriteria_HttpCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionControl_SuccessCriteria_HttpCriteria.ProtoReflect.Descriptor instead.
func (*AdmissionControl_SuccessCriteria_HttpCriteria) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *AdmissionControl_SuccessCriteria_HttpCriteria) GetHttpSuccessStatus() []*v31.Int32Range {
	if x != nil {
		return x.HttpSuccessStatus
	}
	return nil
}

type AdmissionControl_SuccessCriteria_GrpcCriteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrpcSuccessStatus []uint32 `protobuf:"varint,1,rep,packed,name=grpc_success_status,json=grpcSuccessStatus,proto3" json:"grpc_success_status,omitempty"`
}

func (x *AdmissionControl_SuccessCriteria_GrpcCriteria) Reset() {
	*x = AdmissionControl_SuccessCriteria_GrpcCriteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdmissionControl_SuccessCriteria_GrpcCriteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdmissionControl_SuccessCriteria_GrpcCriteria) ProtoMessage() {}

func (x *AdmissionControl_SuccessCriteria_GrpcCriteria) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdmissionControl_SuccessCriteria_GrpcCriteria.ProtoReflect.Descriptor instead.
func (*AdmissionControl_SuccessCriteria_GrpcCriteria) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *AdmissionControl_SuccessCriteria_GrpcCriteria) GetGrpcSuccessStatus() []uint32 {
	if x != nil {
		return x.GrpcSuccessStatus
	}
	return nil
}

var File_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDesc = []byte{
	0x0a, 0x4f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2f, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x37, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75,
	0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x06, 0x0a, 0x10, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x42, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x86,
	0x01, 0x0a, 0x10, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x59, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x42, 0x0a, 0x0f, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x69, 0x6e, 0x67, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x69, 0x6e, 0x67, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x5a, 0x0a, 0x16, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x52, 0x15, 0x61, 0x67, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x65, 0x66,
	0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0xdc, 0x03, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x8b, 0x01, 0x0a, 0x0d,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x66, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x64,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x48,
	0x74, 0x74, 0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x0c, 0x68, 0x74, 0x74,
	0x70, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x8b, 0x01, 0x0a, 0x0d, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x66, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x64, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x2e, 0x47, 0x72, 0x70,
	0x63, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x52, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x1a, 0x63, 0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x43,
	0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x53, 0x0a, 0x13, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x11, 0x68, 0x74, 0x74, 0x70, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x48, 0x0a, 0x0c,
	0x47, 0x72, 0x70, 0x63, 0x43, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x38, 0x0a, 0x13,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x11, 0x67, 0x72, 0x70, 0x63, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x1a, 0x0a, 0x13, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x03, 0xf8,
	0x42, 0x01, 0x42, 0x70, 0x0a, 0x45, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x33, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x15, 0x41, 0x64, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x08, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescData = file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDesc
)

func file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDescData
}

var file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_goTypes = []interface{}{
	(*AdmissionControl)(nil),                              // 0: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl
	(*AdmissionControl_SuccessCriteria)(nil),              // 1: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria
	(*AdmissionControl_SuccessCriteria_HttpCriteria)(nil), // 2: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria.HttpCriteria
	(*AdmissionControl_SuccessCriteria_GrpcCriteria)(nil), // 3: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria.GrpcCriteria
	(*v3.RuntimeFeatureFlag)(nil),                         // 4: envoy.config.core.v3.RuntimeFeatureFlag
	(*duration.Duration)(nil),                             // 5: google.protobuf.Duration
	(*v3.RuntimeDouble)(nil),                              // 6: envoy.config.core.v3.RuntimeDouble
	(*v31.Int32Range)(nil),                                // 7: envoy.type.v3.Int32Range
}
var file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_depIdxs = []int32{
	4, // 0: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.enabled:type_name -> envoy.config.core.v3.RuntimeFeatureFlag
	1, // 1: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.success_criteria:type_name -> envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria
	5, // 2: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.sampling_window:type_name -> google.protobuf.Duration
	6, // 3: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.aggression_coefficient:type_name -> envoy.config.core.v3.RuntimeDouble
	2, // 4: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria.http_criteria:type_name -> envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria.HttpCriteria
	3, // 5: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria.grpc_criteria:type_name -> envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria.GrpcCriteria
	7, // 6: envoy.extensions.filters.http.admission_control.v3alpha.AdmissionControl.SuccessCriteria.HttpCriteria.http_success_status:type_name -> envoy.type.v3.Int32Range
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_init()
}
func file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_init() {
	if File_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionControl); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionControl_SuccessCriteria); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionControl_SuccessCriteria_HttpCriteria); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdmissionControl_SuccessCriteria_GrpcCriteria); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AdmissionControl_SuccessCriteria_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto = out.File
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_rawDesc = nil
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_goTypes = nil
	file_envoy_extensions_filters_http_admission_control_v3alpha_admission_control_proto_depIdxs = nil
}
