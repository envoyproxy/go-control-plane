// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: envoy/extensions/filters/network/ext_proc/v3/ext_proc.proto

package ext_procv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Defines how traffic should be handled by the external processor.
type ProcessingMode_DataSendMode int32

const (
	// Send the data to the external processor for processing whenever the data is ready.
	ProcessingMode_STREAMED ProcessingMode_DataSendMode = 0
	// Skip sending the data to the external processor.
	ProcessingMode_SKIP ProcessingMode_DataSendMode = 1
)

// Enum value maps for ProcessingMode_DataSendMode.
var (
	ProcessingMode_DataSendMode_name = map[int32]string{
		0: "STREAMED",
		1: "SKIP",
	}
	ProcessingMode_DataSendMode_value = map[string]int32{
		"STREAMED": 0,
		"SKIP":     1,
	}
)

func (x ProcessingMode_DataSendMode) Enum() *ProcessingMode_DataSendMode {
	p := new(ProcessingMode_DataSendMode)
	*p = x
	return p
}

func (x ProcessingMode_DataSendMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessingMode_DataSendMode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_enumTypes[0].Descriptor()
}

func (ProcessingMode_DataSendMode) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_enumTypes[0]
}

func (x ProcessingMode_DataSendMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessingMode_DataSendMode.Descriptor instead.
func (ProcessingMode_DataSendMode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescGZIP(), []int{1, 0}
}

// The Network External Processing filter allows an external service to process raw TCP/UDP traffic
// in a flexible way using a bidirectional gRPC stream. Unlike the HTTP External Processing filter,
// this filter operates at the L4 (transport) layer, giving access to raw network traffic.
//
// The filter communicates with an external gRPC service that can:
// * Inspect traffic in both directions
// * Modify the network traffic
// * Control connection lifecycle (continue, close, or reset)
//
// By using the filter's processing mode, you can selectively choose which data
// directions to process (read, write or both), allowing for efficient processing.
type NetworkExternalProcessor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The gRPC service that will process network traffic.
	// This service must implement the NetworkExternalProcessor service
	// defined in the proto file /envoy/service/network_ext_proc/v3/external_processor.proto.
	GrpcService *v3.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3" json:"grpc_service,omitempty"`
	// By default, if the gRPC stream cannot be established, or if it is closed
	// prematurely with an error, the filter will fail, leading to the close of connection.
	// With this parameter set to true, however, then if the gRPC stream is prematurely closed
	// or could not be opened, processing continues without error.
	// [#not-implemented-hide:]
	FailureModeAllow bool `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Options for controlling processing behavior.
	// [#not-implemented-hide:]
	ProcessingMode *ProcessingMode `protobuf:"bytes,3,opt,name=processing_mode,json=processingMode,proto3" json:"processing_mode,omitempty"`
	// Specifies the timeout for each individual message sent on the stream and
	// when the filter is running in synchronous mode. Whenever
	// the proxy sends a message on the stream that requires a response, it will
	// reset this timer, and will stop processing and return an error (subject
	// to the processing mode) if the timer expires. Default is 200 ms.
	// [#not-implemented-hide:]
	MessageTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=message_timeout,json=messageTimeout,proto3" json:"message_timeout,omitempty"`
}

func (x *NetworkExternalProcessor) Reset() {
	*x = NetworkExternalProcessor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkExternalProcessor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkExternalProcessor) ProtoMessage() {}

func (x *NetworkExternalProcessor) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkExternalProcessor.ProtoReflect.Descriptor instead.
func (*NetworkExternalProcessor) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkExternalProcessor) GetGrpcService() *v3.GrpcService {
	if x != nil {
		return x.GrpcService
	}
	return nil
}

func (x *NetworkExternalProcessor) GetFailureModeAllow() bool {
	if x != nil {
		return x.FailureModeAllow
	}
	return false
}

func (x *NetworkExternalProcessor) GetProcessingMode() *ProcessingMode {
	if x != nil {
		return x.ProcessingMode
	}
	return nil
}

func (x *NetworkExternalProcessor) GetMessageTimeout() *durationpb.Duration {
	if x != nil {
		return x.MessageTimeout
	}
	return nil
}

// Options for controlling processing behavior.
// Filter will reject the config if both read and write are SKIP mode.
type ProcessingMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Controls whether inbound (read) data from the client is sent to the external processor.
	// Default: STREAMED
	ProcessRead ProcessingMode_DataSendMode `protobuf:"varint,1,opt,name=process_read,json=processRead,proto3,enum=envoy.extensions.filters.network.ext_proc.v3.ProcessingMode_DataSendMode" json:"process_read,omitempty"`
	// Controls whether outbound (write) data to the client is sent to the external processor.
	// Default: STREAMED
	ProcessWrite ProcessingMode_DataSendMode `protobuf:"varint,2,opt,name=process_write,json=processWrite,proto3,enum=envoy.extensions.filters.network.ext_proc.v3.ProcessingMode_DataSendMode" json:"process_write,omitempty"`
}

func (x *ProcessingMode) Reset() {
	*x = ProcessingMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingMode) ProtoMessage() {}

func (x *ProcessingMode) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingMode.ProtoReflect.Descriptor instead.
func (*ProcessingMode) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessingMode) GetProcessRead() ProcessingMode_DataSendMode {
	if x != nil {
		return x.ProcessRead
	}
	return ProcessingMode_STREAMED
}

func (x *ProcessingMode) GetProcessWrite() ProcessingMode_DataSendMode {
	if x != nil {
		return x.ProcessWrite
	}
	return ProcessingMode_STREAMED
}

var File_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2f, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2f, 0x76, 0x33, 0x2f, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33, 0x1a, 0x27, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02,
	0x0a, 0x18, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x0c, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x65,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x51, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0xaa, 0x01,
	0x07, 0x22, 0x03, 0x08, 0x90, 0x1c, 0x32, 0x00, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x96, 0x02, 0x0a, 0x0e, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x6c, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x49, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x61, 0x64, 0x12, 0x6e, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x49, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0c, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x57, 0x72, 0x69, 0x74, 0x65, 0x22, 0x26, 0x0a, 0x0c, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x52,
	0x45, 0x41, 0x4d, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4b, 0x49, 0x50, 0x10,
	0x01, 0x42, 0xbc, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1,
	0x06, 0x02, 0x08, 0x01, 0x0a, 0x3a, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x2e, 0x76, 0x33,
	0x42, 0x0c, 0x45, 0x78, 0x74, 0x50, 0x72, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72,
	0x6f, 0x63, 0x2f, 0x76, 0x33, 0x3b, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x76, 0x33,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescData = file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDesc
)

func file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescData)
	})
	return file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDescData
}

var file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_goTypes = []interface{}{
	(ProcessingMode_DataSendMode)(0), // 0: envoy.extensions.filters.network.ext_proc.v3.ProcessingMode.DataSendMode
	(*NetworkExternalProcessor)(nil), // 1: envoy.extensions.filters.network.ext_proc.v3.NetworkExternalProcessor
	(*ProcessingMode)(nil),           // 2: envoy.extensions.filters.network.ext_proc.v3.ProcessingMode
	(*v3.GrpcService)(nil),           // 3: envoy.config.core.v3.GrpcService
	(*durationpb.Duration)(nil),      // 4: google.protobuf.Duration
}
var file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_depIdxs = []int32{
	3, // 0: envoy.extensions.filters.network.ext_proc.v3.NetworkExternalProcessor.grpc_service:type_name -> envoy.config.core.v3.GrpcService
	2, // 1: envoy.extensions.filters.network.ext_proc.v3.NetworkExternalProcessor.processing_mode:type_name -> envoy.extensions.filters.network.ext_proc.v3.ProcessingMode
	4, // 2: envoy.extensions.filters.network.ext_proc.v3.NetworkExternalProcessor.message_timeout:type_name -> google.protobuf.Duration
	0, // 3: envoy.extensions.filters.network.ext_proc.v3.ProcessingMode.process_read:type_name -> envoy.extensions.filters.network.ext_proc.v3.ProcessingMode.DataSendMode
	0, // 4: envoy.extensions.filters.network.ext_proc.v3.ProcessingMode.process_write:type_name -> envoy.extensions.filters.network.ext_proc.v3.ProcessingMode.DataSendMode
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_init() }
func file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_init() {
	if File_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkExternalProcessor); i {
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
		file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingMode); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto = out.File
	file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_rawDesc = nil
	file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_goTypes = nil
	file_envoy_extensions_filters_network_ext_proc_v3_ext_proc_proto_depIdxs = nil
}
