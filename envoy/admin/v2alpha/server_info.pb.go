// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: envoy/admin/v2alpha/server_info.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
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

type ServerInfo_State int32

const (
	// Server is live and serving traffic.
	ServerInfo_LIVE ServerInfo_State = 0
	// Server is draining listeners in response to external health checks failing.
	ServerInfo_DRAINING ServerInfo_State = 1
	// Server has not yet completed cluster manager initialization.
	ServerInfo_PRE_INITIALIZING ServerInfo_State = 2
	// Server is running the cluster manager initialization callbacks (e.g., RDS).
	ServerInfo_INITIALIZING ServerInfo_State = 3
)

// Enum value maps for ServerInfo_State.
var (
	ServerInfo_State_name = map[int32]string{
		0: "LIVE",
		1: "DRAINING",
		2: "PRE_INITIALIZING",
		3: "INITIALIZING",
	}
	ServerInfo_State_value = map[string]int32{
		"LIVE":             0,
		"DRAINING":         1,
		"PRE_INITIALIZING": 2,
		"INITIALIZING":     3,
	}
)

func (x ServerInfo_State) Enum() *ServerInfo_State {
	p := new(ServerInfo_State)
	*p = x
	return p
}

func (x ServerInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_admin_v2alpha_server_info_proto_enumTypes[0].Descriptor()
}

func (ServerInfo_State) Type() protoreflect.EnumType {
	return &file_envoy_admin_v2alpha_server_info_proto_enumTypes[0]
}

func (x ServerInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerInfo_State.Descriptor instead.
func (ServerInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_server_info_proto_rawDescGZIP(), []int{0, 0}
}

type CommandLineOptions_IpVersion int32

const (
	CommandLineOptions_v4 CommandLineOptions_IpVersion = 0
	CommandLineOptions_v6 CommandLineOptions_IpVersion = 1
)

// Enum value maps for CommandLineOptions_IpVersion.
var (
	CommandLineOptions_IpVersion_name = map[int32]string{
		0: "v4",
		1: "v6",
	}
	CommandLineOptions_IpVersion_value = map[string]int32{
		"v4": 0,
		"v6": 1,
	}
)

func (x CommandLineOptions_IpVersion) Enum() *CommandLineOptions_IpVersion {
	p := new(CommandLineOptions_IpVersion)
	*p = x
	return p
}

func (x CommandLineOptions_IpVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandLineOptions_IpVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_admin_v2alpha_server_info_proto_enumTypes[1].Descriptor()
}

func (CommandLineOptions_IpVersion) Type() protoreflect.EnumType {
	return &file_envoy_admin_v2alpha_server_info_proto_enumTypes[1]
}

func (x CommandLineOptions_IpVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandLineOptions_IpVersion.Descriptor instead.
func (CommandLineOptions_IpVersion) EnumDescriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_server_info_proto_rawDescGZIP(), []int{1, 0}
}

type CommandLineOptions_Mode int32

const (
	// Validate configs and then serve traffic normally.
	CommandLineOptions_Serve CommandLineOptions_Mode = 0
	// Validate configs and exit.
	CommandLineOptions_Validate CommandLineOptions_Mode = 1
	// Completely load and initialize the config, and then exit without running the listener loop.
	CommandLineOptions_InitOnly CommandLineOptions_Mode = 2
)

// Enum value maps for CommandLineOptions_Mode.
var (
	CommandLineOptions_Mode_name = map[int32]string{
		0: "Serve",
		1: "Validate",
		2: "InitOnly",
	}
	CommandLineOptions_Mode_value = map[string]int32{
		"Serve":    0,
		"Validate": 1,
		"InitOnly": 2,
	}
)

func (x CommandLineOptions_Mode) Enum() *CommandLineOptions_Mode {
	p := new(CommandLineOptions_Mode)
	*p = x
	return p
}

func (x CommandLineOptions_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandLineOptions_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_admin_v2alpha_server_info_proto_enumTypes[2].Descriptor()
}

func (CommandLineOptions_Mode) Type() protoreflect.EnumType {
	return &file_envoy_admin_v2alpha_server_info_proto_enumTypes[2]
}

func (x CommandLineOptions_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandLineOptions_Mode.Descriptor instead.
func (CommandLineOptions_Mode) EnumDescriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_server_info_proto_rawDescGZIP(), []int{1, 1}
}

// Proto representation of the value returned by /server_info, containing
// server version/server status information.
// [#next-free-field: 7]
type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Server version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// State of the server.
	State ServerInfo_State `protobuf:"varint,2,opt,name=state,proto3,enum=envoy.admin.v2alpha.ServerInfo_State" json:"state,omitempty"`
	// Uptime since current epoch was started.
	UptimeCurrentEpoch *durationpb.Duration `protobuf:"bytes,3,opt,name=uptime_current_epoch,json=uptimeCurrentEpoch,proto3" json:"uptime_current_epoch,omitempty"`
	// Uptime since the start of the first epoch.
	UptimeAllEpochs *durationpb.Duration `protobuf:"bytes,4,opt,name=uptime_all_epochs,json=uptimeAllEpochs,proto3" json:"uptime_all_epochs,omitempty"`
	// Hot restart version.
	HotRestartVersion string `protobuf:"bytes,5,opt,name=hot_restart_version,json=hotRestartVersion,proto3" json:"hot_restart_version,omitempty"`
	// Command line options the server is currently running with.
	CommandLineOptions *CommandLineOptions `protobuf:"bytes,6,opt,name=command_line_options,json=commandLineOptions,proto3" json:"command_line_options,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v2alpha_server_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v2alpha_server_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_server_info_proto_rawDescGZIP(), []int{0}
}

func (x *ServerInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ServerInfo) GetState() ServerInfo_State {
	if x != nil {
		return x.State
	}
	return ServerInfo_LIVE
}

func (x *ServerInfo) GetUptimeCurrentEpoch() *durationpb.Duration {
	if x != nil {
		return x.UptimeCurrentEpoch
	}
	return nil
}

func (x *ServerInfo) GetUptimeAllEpochs() *durationpb.Duration {
	if x != nil {
		return x.UptimeAllEpochs
	}
	return nil
}

func (x *ServerInfo) GetHotRestartVersion() string {
	if x != nil {
		return x.HotRestartVersion
	}
	return ""
}

func (x *ServerInfo) GetCommandLineOptions() *CommandLineOptions {
	if x != nil {
		return x.CommandLineOptions
	}
	return nil
}

// [#next-free-field: 29]
type CommandLineOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// See :option:`--base-id` for details.
	BaseId uint64 `protobuf:"varint,1,opt,name=base_id,json=baseId,proto3" json:"base_id,omitempty"`
	// See :option:`--concurrency` for details.
	Concurrency uint32 `protobuf:"varint,2,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// See :option:`--config-path` for details.
	ConfigPath string `protobuf:"bytes,3,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`
	// See :option:`--config-yaml` for details.
	ConfigYaml string `protobuf:"bytes,4,opt,name=config_yaml,json=configYaml,proto3" json:"config_yaml,omitempty"`
	// See :option:`--allow-unknown-static-fields` for details.
	AllowUnknownStaticFields bool `protobuf:"varint,5,opt,name=allow_unknown_static_fields,json=allowUnknownStaticFields,proto3" json:"allow_unknown_static_fields,omitempty"`
	// See :option:`--reject-unknown-dynamic-fields` for details.
	RejectUnknownDynamicFields bool `protobuf:"varint,26,opt,name=reject_unknown_dynamic_fields,json=rejectUnknownDynamicFields,proto3" json:"reject_unknown_dynamic_fields,omitempty"`
	// See :option:`--admin-address-path` for details.
	AdminAddressPath string `protobuf:"bytes,6,opt,name=admin_address_path,json=adminAddressPath,proto3" json:"admin_address_path,omitempty"`
	// See :option:`--local-address-ip-version` for details.
	LocalAddressIpVersion CommandLineOptions_IpVersion `protobuf:"varint,7,opt,name=local_address_ip_version,json=localAddressIpVersion,proto3,enum=envoy.admin.v2alpha.CommandLineOptions_IpVersion" json:"local_address_ip_version,omitempty"`
	// See :option:`--log-level` for details.
	LogLevel string `protobuf:"bytes,8,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	// See :option:`--component-log-level` for details.
	ComponentLogLevel string `protobuf:"bytes,9,opt,name=component_log_level,json=componentLogLevel,proto3" json:"component_log_level,omitempty"`
	// See :option:`--log-format` for details.
	LogFormat string `protobuf:"bytes,10,opt,name=log_format,json=logFormat,proto3" json:"log_format,omitempty"`
	// See :option:`--log-format-escaped` for details.
	LogFormatEscaped bool `protobuf:"varint,27,opt,name=log_format_escaped,json=logFormatEscaped,proto3" json:"log_format_escaped,omitempty"`
	// See :option:`--log-path` for details.
	LogPath string `protobuf:"bytes,11,opt,name=log_path,json=logPath,proto3" json:"log_path,omitempty"`
	// See :option:`--service-cluster` for details.
	ServiceCluster string `protobuf:"bytes,13,opt,name=service_cluster,json=serviceCluster,proto3" json:"service_cluster,omitempty"`
	// See :option:`--service-node` for details.
	ServiceNode string `protobuf:"bytes,14,opt,name=service_node,json=serviceNode,proto3" json:"service_node,omitempty"`
	// See :option:`--service-zone` for details.
	ServiceZone string `protobuf:"bytes,15,opt,name=service_zone,json=serviceZone,proto3" json:"service_zone,omitempty"`
	// See :option:`--file-flush-interval-msec` for details.
	FileFlushInterval *durationpb.Duration `protobuf:"bytes,16,opt,name=file_flush_interval,json=fileFlushInterval,proto3" json:"file_flush_interval,omitempty"`
	// See :option:`--drain-time-s` for details.
	DrainTime *durationpb.Duration `protobuf:"bytes,17,opt,name=drain_time,json=drainTime,proto3" json:"drain_time,omitempty"`
	// See :option:`--parent-shutdown-time-s` for details.
	ParentShutdownTime *durationpb.Duration `protobuf:"bytes,18,opt,name=parent_shutdown_time,json=parentShutdownTime,proto3" json:"parent_shutdown_time,omitempty"`
	// See :option:`--mode` for details.
	Mode CommandLineOptions_Mode `protobuf:"varint,19,opt,name=mode,proto3,enum=envoy.admin.v2alpha.CommandLineOptions_Mode" json:"mode,omitempty"`
	// max_stats and max_obj_name_len are now unused and have no effect.
	//
	// Deprecated: Marked as deprecated in envoy/admin/v2alpha/server_info.proto.
	MaxStats uint64 `protobuf:"varint,20,opt,name=max_stats,json=maxStats,proto3" json:"max_stats,omitempty"`
	// Deprecated: Marked as deprecated in envoy/admin/v2alpha/server_info.proto.
	MaxObjNameLen uint64 `protobuf:"varint,21,opt,name=max_obj_name_len,json=maxObjNameLen,proto3" json:"max_obj_name_len,omitempty"`
	// See :option:`--disable-hot-restart` for details.
	DisableHotRestart bool `protobuf:"varint,22,opt,name=disable_hot_restart,json=disableHotRestart,proto3" json:"disable_hot_restart,omitempty"`
	// See :option:`--enable-mutex-tracing` for details.
	EnableMutexTracing bool `protobuf:"varint,23,opt,name=enable_mutex_tracing,json=enableMutexTracing,proto3" json:"enable_mutex_tracing,omitempty"`
	// See :option:`--restart-epoch` for details.
	RestartEpoch uint32 `protobuf:"varint,24,opt,name=restart_epoch,json=restartEpoch,proto3" json:"restart_epoch,omitempty"`
	// See :option:`--cpuset-threads` for details.
	CpusetThreads bool `protobuf:"varint,25,opt,name=cpuset_threads,json=cpusetThreads,proto3" json:"cpuset_threads,omitempty"`
	// See :option:`--disable-extensions` for details.
	DisabledExtensions []string `protobuf:"bytes,28,rep,name=disabled_extensions,json=disabledExtensions,proto3" json:"disabled_extensions,omitempty"`
}

func (x *CommandLineOptions) Reset() {
	*x = CommandLineOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_admin_v2alpha_server_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandLineOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandLineOptions) ProtoMessage() {}

func (x *CommandLineOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_admin_v2alpha_server_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandLineOptions.ProtoReflect.Descriptor instead.
func (*CommandLineOptions) Descriptor() ([]byte, []int) {
	return file_envoy_admin_v2alpha_server_info_proto_rawDescGZIP(), []int{1}
}

func (x *CommandLineOptions) GetBaseId() uint64 {
	if x != nil {
		return x.BaseId
	}
	return 0
}

func (x *CommandLineOptions) GetConcurrency() uint32 {
	if x != nil {
		return x.Concurrency
	}
	return 0
}

func (x *CommandLineOptions) GetConfigPath() string {
	if x != nil {
		return x.ConfigPath
	}
	return ""
}

func (x *CommandLineOptions) GetConfigYaml() string {
	if x != nil {
		return x.ConfigYaml
	}
	return ""
}

func (x *CommandLineOptions) GetAllowUnknownStaticFields() bool {
	if x != nil {
		return x.AllowUnknownStaticFields
	}
	return false
}

func (x *CommandLineOptions) GetRejectUnknownDynamicFields() bool {
	if x != nil {
		return x.RejectUnknownDynamicFields
	}
	return false
}

func (x *CommandLineOptions) GetAdminAddressPath() string {
	if x != nil {
		return x.AdminAddressPath
	}
	return ""
}

func (x *CommandLineOptions) GetLocalAddressIpVersion() CommandLineOptions_IpVersion {
	if x != nil {
		return x.LocalAddressIpVersion
	}
	return CommandLineOptions_v4
}

func (x *CommandLineOptions) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *CommandLineOptions) GetComponentLogLevel() string {
	if x != nil {
		return x.ComponentLogLevel
	}
	return ""
}

func (x *CommandLineOptions) GetLogFormat() string {
	if x != nil {
		return x.LogFormat
	}
	return ""
}

func (x *CommandLineOptions) GetLogFormatEscaped() bool {
	if x != nil {
		return x.LogFormatEscaped
	}
	return false
}

func (x *CommandLineOptions) GetLogPath() string {
	if x != nil {
		return x.LogPath
	}
	return ""
}

func (x *CommandLineOptions) GetServiceCluster() string {
	if x != nil {
		return x.ServiceCluster
	}
	return ""
}

func (x *CommandLineOptions) GetServiceNode() string {
	if x != nil {
		return x.ServiceNode
	}
	return ""
}

func (x *CommandLineOptions) GetServiceZone() string {
	if x != nil {
		return x.ServiceZone
	}
	return ""
}

func (x *CommandLineOptions) GetFileFlushInterval() *durationpb.Duration {
	if x != nil {
		return x.FileFlushInterval
	}
	return nil
}

func (x *CommandLineOptions) GetDrainTime() *durationpb.Duration {
	if x != nil {
		return x.DrainTime
	}
	return nil
}

func (x *CommandLineOptions) GetParentShutdownTime() *durationpb.Duration {
	if x != nil {
		return x.ParentShutdownTime
	}
	return nil
}

func (x *CommandLineOptions) GetMode() CommandLineOptions_Mode {
	if x != nil {
		return x.Mode
	}
	return CommandLineOptions_Serve
}

// Deprecated: Marked as deprecated in envoy/admin/v2alpha/server_info.proto.
func (x *CommandLineOptions) GetMaxStats() uint64 {
	if x != nil {
		return x.MaxStats
	}
	return 0
}

// Deprecated: Marked as deprecated in envoy/admin/v2alpha/server_info.proto.
func (x *CommandLineOptions) GetMaxObjNameLen() uint64 {
	if x != nil {
		return x.MaxObjNameLen
	}
	return 0
}

func (x *CommandLineOptions) GetDisableHotRestart() bool {
	if x != nil {
		return x.DisableHotRestart
	}
	return false
}

func (x *CommandLineOptions) GetEnableMutexTracing() bool {
	if x != nil {
		return x.EnableMutexTracing
	}
	return false
}

func (x *CommandLineOptions) GetRestartEpoch() uint32 {
	if x != nil {
		return x.RestartEpoch
	}
	return 0
}

func (x *CommandLineOptions) GetCpusetThreads() bool {
	if x != nil {
		return x.CpusetThreads
	}
	return false
}

func (x *CommandLineOptions) GetDisabledExtensions() []string {
	if x != nil {
		return x.DisabledExtensions
	}
	return nil
}

var File_envoy_admin_v2alpha_server_info_proto protoreflect.FileDescriptor

var file_envoy_admin_v2alpha_server_info_proto_rawDesc = []byte{
	0x0a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x32,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xcb, 0x03, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x12, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x45, 0x70,
	0x6f, 0x63, 0x68, 0x12, 0x45, 0x0a, 0x11, 0x75, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x75, 0x70, 0x74, 0x69, 0x6d,
	0x65, 0x41, 0x6c, 0x6c, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x6f,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x59, 0x0a, 0x14, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x47, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x52, 0x41, 0x49,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x45, 0x5f, 0x49, 0x4e,
	0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x22, 0xf0,
	0x0a, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x79, 0x61, 0x6d, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61,
	0x6d, 0x6c, 0x12, 0x3d, 0x0a, 0x1b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x6a, 0x0a, 0x18, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x70,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x6f, 0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x6f,
	0x67, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x65, 0x73, 0x63, 0x61, 0x70, 0x65, 0x64,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6c, 0x6f, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x45, 0x73, 0x63, 0x61, 0x70, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5a, 0x6f,
	0x6e, 0x65, 0x12, 0x49, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x66, 0x6c, 0x75, 0x73, 0x68,
	0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11, 0x66, 0x69, 0x6c, 0x65,
	0x46, 0x6c, 0x75, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x38, 0x0a,
	0x0a, 0x64, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x72,
	0x61, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4c, 0x69, 0x6e, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xb8, 0xee, 0xf2, 0xd2, 0x05,
	0x01, 0x18, 0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x31, 0x0a,
	0x10, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x62, 0x6a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x65,
	0x6e, 0x18, 0x15, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xb8, 0xee, 0xf2, 0xd2, 0x05, 0x01, 0x18,
	0x01, 0x52, 0x0d, 0x6d, 0x61, 0x78, 0x4f, 0x62, 0x6a, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x65, 0x6e,
	0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x6f, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64,
	0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x75, 0x74, 0x65, 0x78,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x75, 0x74, 0x65, 0x78, 0x54, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x18, 0x18, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x70, 0x75, 0x73, 0x65,
	0x74, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x63, 0x70, 0x75, 0x73, 0x65, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x2f,
	0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x1b, 0x0a, 0x09, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02,
	0x76, 0x34, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x76, 0x36, 0x10, 0x01, 0x22, 0x2d, 0x0a, 0x04,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x65, 0x72, 0x76, 0x65, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0c, 0x0a,
	0x08, 0x49, 0x6e, 0x69, 0x74, 0x4f, 0x6e, 0x6c, 0x79, 0x10, 0x02, 0x4a, 0x04, 0x08, 0x0c, 0x10,
	0x0d, 0x42, 0x7a, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x0a, 0x21, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x0f,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_admin_v2alpha_server_info_proto_rawDescOnce sync.Once
	file_envoy_admin_v2alpha_server_info_proto_rawDescData = file_envoy_admin_v2alpha_server_info_proto_rawDesc
)

func file_envoy_admin_v2alpha_server_info_proto_rawDescGZIP() []byte {
	file_envoy_admin_v2alpha_server_info_proto_rawDescOnce.Do(func() {
		file_envoy_admin_v2alpha_server_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_admin_v2alpha_server_info_proto_rawDescData)
	})
	return file_envoy_admin_v2alpha_server_info_proto_rawDescData
}

var file_envoy_admin_v2alpha_server_info_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_envoy_admin_v2alpha_server_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_admin_v2alpha_server_info_proto_goTypes = []interface{}{
	(ServerInfo_State)(0),             // 0: envoy.admin.v2alpha.ServerInfo.State
	(CommandLineOptions_IpVersion)(0), // 1: envoy.admin.v2alpha.CommandLineOptions.IpVersion
	(CommandLineOptions_Mode)(0),      // 2: envoy.admin.v2alpha.CommandLineOptions.Mode
	(*ServerInfo)(nil),                // 3: envoy.admin.v2alpha.ServerInfo
	(*CommandLineOptions)(nil),        // 4: envoy.admin.v2alpha.CommandLineOptions
	(*durationpb.Duration)(nil),       // 5: google.protobuf.Duration
}
var file_envoy_admin_v2alpha_server_info_proto_depIdxs = []int32{
	0, // 0: envoy.admin.v2alpha.ServerInfo.state:type_name -> envoy.admin.v2alpha.ServerInfo.State
	5, // 1: envoy.admin.v2alpha.ServerInfo.uptime_current_epoch:type_name -> google.protobuf.Duration
	5, // 2: envoy.admin.v2alpha.ServerInfo.uptime_all_epochs:type_name -> google.protobuf.Duration
	4, // 3: envoy.admin.v2alpha.ServerInfo.command_line_options:type_name -> envoy.admin.v2alpha.CommandLineOptions
	1, // 4: envoy.admin.v2alpha.CommandLineOptions.local_address_ip_version:type_name -> envoy.admin.v2alpha.CommandLineOptions.IpVersion
	5, // 5: envoy.admin.v2alpha.CommandLineOptions.file_flush_interval:type_name -> google.protobuf.Duration
	5, // 6: envoy.admin.v2alpha.CommandLineOptions.drain_time:type_name -> google.protobuf.Duration
	5, // 7: envoy.admin.v2alpha.CommandLineOptions.parent_shutdown_time:type_name -> google.protobuf.Duration
	2, // 8: envoy.admin.v2alpha.CommandLineOptions.mode:type_name -> envoy.admin.v2alpha.CommandLineOptions.Mode
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_envoy_admin_v2alpha_server_info_proto_init() }
func file_envoy_admin_v2alpha_server_info_proto_init() {
	if File_envoy_admin_v2alpha_server_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_admin_v2alpha_server_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_envoy_admin_v2alpha_server_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandLineOptions); i {
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
			RawDescriptor: file_envoy_admin_v2alpha_server_info_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_admin_v2alpha_server_info_proto_goTypes,
		DependencyIndexes: file_envoy_admin_v2alpha_server_info_proto_depIdxs,
		EnumInfos:         file_envoy_admin_v2alpha_server_info_proto_enumTypes,
		MessageInfos:      file_envoy_admin_v2alpha_server_info_proto_msgTypes,
	}.Build()
	File_envoy_admin_v2alpha_server_info_proto = out.File
	file_envoy_admin_v2alpha_server_info_proto_rawDesc = nil
	file_envoy_admin_v2alpha_server_info_proto_goTypes = nil
	file_envoy_admin_v2alpha_server_info_proto_depIdxs = nil
}
