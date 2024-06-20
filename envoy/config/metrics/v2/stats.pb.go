// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: envoy/config/metrics/v2/stats.proto

package metricsv2

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for pluggable stats sinks.
type StatsSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the stats sink to instantiate. The name must match a supported
	// stats sink. The built-in stats sinks are:
	//
	// * :ref:`envoy.stat_sinks.statsd <envoy_api_msg_config.metrics.v2.StatsdSink>`
	// * :ref:`envoy.stat_sinks.dog_statsd <envoy_api_msg_config.metrics.v2.DogStatsdSink>`
	// * :ref:`envoy.stat_sinks.metrics_service <envoy_api_msg_config.metrics.v2.MetricsServiceConfig>`
	// * :ref:`envoy.stat_sinks.hystrix <envoy_api_msg_config.metrics.v2.HystrixSink>`
	//
	// Sinks optionally support tagged/multiple dimensional metrics.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Stats sink specific configuration which depends on the sink being instantiated. See
	// :ref:`StatsdSink <envoy_api_msg_config.metrics.v2.StatsdSink>` for an example.
	//
	// Types that are assignable to ConfigType:
	//
	//	*StatsSink_Config
	//	*StatsSink_TypedConfig
	ConfigType isStatsSink_ConfigType `protobuf_oneof:"config_type"`
}

func (x *StatsSink) Reset() {
	*x = StatsSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsSink) ProtoMessage() {}

func (x *StatsSink) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsSink.ProtoReflect.Descriptor instead.
func (*StatsSink) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v2_stats_proto_rawDescGZIP(), []int{0}
}

func (x *StatsSink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *StatsSink) GetConfigType() isStatsSink_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Marked as deprecated in envoy/config/metrics/v2/stats.proto.
func (x *StatsSink) GetConfig() *structpb.Struct {
	if x, ok := x.GetConfigType().(*StatsSink_Config); ok {
		return x.Config
	}
	return nil
}

func (x *StatsSink) GetTypedConfig() *anypb.Any {
	if x, ok := x.GetConfigType().(*StatsSink_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

type isStatsSink_ConfigType interface {
	isStatsSink_ConfigType()
}

type StatsSink_Config struct {
	// Deprecated: Marked as deprecated in envoy/config/metrics/v2/stats.proto.
	Config *structpb.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type StatsSink_TypedConfig struct {
	TypedConfig *anypb.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*StatsSink_Config) isStatsSink_ConfigType() {}

func (*StatsSink_TypedConfig) isStatsSink_ConfigType() {}

// Statistics configuration such as tagging.
type StatsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each stat name is iteratively processed through these tag specifiers.
	// When a tag is matched, the first capture group is removed from the name so
	// later :ref:`TagSpecifiers <envoy_api_msg_config.metrics.v2.TagSpecifier>` cannot match that
	// same portion of the match.
	StatsTags []*TagSpecifier `protobuf:"bytes,1,rep,name=stats_tags,json=statsTags,proto3" json:"stats_tags,omitempty"`
	// Use all default tag regexes specified in Envoy. These can be combined with
	// custom tags specified in :ref:`stats_tags
	// <envoy_api_field_config.metrics.v2.StatsConfig.stats_tags>`. They will be processed before
	// the custom tags.
	//
	// .. note::
	//
	//	If any default tags are specified twice, the config will be considered
	//	invalid.
	//
	// See :repo:`well_known_names.h <source/common/config/well_known_names.h>` for a list of the
	// default tags in Envoy.
	//
	// If not provided, the value is assumed to be true.
	UseAllDefaultTags *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=use_all_default_tags,json=useAllDefaultTags,proto3" json:"use_all_default_tags,omitempty"`
	// Inclusion/exclusion matcher for stat name creation. If not provided, all stats are instantiated
	// as normal. Preventing the instantiation of certain families of stats can improve memory
	// performance for Envoys running especially large configs.
	//
	// .. warning::
	//
	//	Excluding stats may affect Envoy's behavior in undocumented ways. See
	//	`issue #8771 <https://github.com/envoyproxy/envoy/issues/8771>`_ for more information.
	//	If any unexpected behavior changes are observed, please open a new issue immediately.
	StatsMatcher *StatsMatcher `protobuf:"bytes,3,opt,name=stats_matcher,json=statsMatcher,proto3" json:"stats_matcher,omitempty"`
}

func (x *StatsConfig) Reset() {
	*x = StatsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsConfig) ProtoMessage() {}

func (x *StatsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsConfig.ProtoReflect.Descriptor instead.
func (*StatsConfig) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v2_stats_proto_rawDescGZIP(), []int{1}
}

func (x *StatsConfig) GetStatsTags() []*TagSpecifier {
	if x != nil {
		return x.StatsTags
	}
	return nil
}

func (x *StatsConfig) GetUseAllDefaultTags() *wrapperspb.BoolValue {
	if x != nil {
		return x.UseAllDefaultTags
	}
	return nil
}

func (x *StatsConfig) GetStatsMatcher() *StatsMatcher {
	if x != nil {
		return x.StatsMatcher
	}
	return nil
}

// Configuration for disabling stat instantiation.
type StatsMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to StatsMatcher:
	//
	//	*StatsMatcher_RejectAll
	//	*StatsMatcher_ExclusionList
	//	*StatsMatcher_InclusionList
	StatsMatcher isStatsMatcher_StatsMatcher `protobuf_oneof:"stats_matcher"`
}

func (x *StatsMatcher) Reset() {
	*x = StatsMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsMatcher) ProtoMessage() {}

func (x *StatsMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsMatcher.ProtoReflect.Descriptor instead.
func (*StatsMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v2_stats_proto_rawDescGZIP(), []int{2}
}

func (m *StatsMatcher) GetStatsMatcher() isStatsMatcher_StatsMatcher {
	if m != nil {
		return m.StatsMatcher
	}
	return nil
}

func (x *StatsMatcher) GetRejectAll() bool {
	if x, ok := x.GetStatsMatcher().(*StatsMatcher_RejectAll); ok {
		return x.RejectAll
	}
	return false
}

func (x *StatsMatcher) GetExclusionList() *matcher.ListStringMatcher {
	if x, ok := x.GetStatsMatcher().(*StatsMatcher_ExclusionList); ok {
		return x.ExclusionList
	}
	return nil
}

func (x *StatsMatcher) GetInclusionList() *matcher.ListStringMatcher {
	if x, ok := x.GetStatsMatcher().(*StatsMatcher_InclusionList); ok {
		return x.InclusionList
	}
	return nil
}

type isStatsMatcher_StatsMatcher interface {
	isStatsMatcher_StatsMatcher()
}

type StatsMatcher_RejectAll struct {
	// If `reject_all` is true, then all stats are disabled. If `reject_all` is false, then all
	// stats are enabled.
	RejectAll bool `protobuf:"varint,1,opt,name=reject_all,json=rejectAll,proto3,oneof"`
}

type StatsMatcher_ExclusionList struct {
	// Exclusive match. All stats are enabled except for those matching one of the supplied
	// StringMatcher protos.
	ExclusionList *matcher.ListStringMatcher `protobuf:"bytes,2,opt,name=exclusion_list,json=exclusionList,proto3,oneof"`
}

type StatsMatcher_InclusionList struct {
	// Inclusive match. No stats are enabled except for those matching one of the supplied
	// StringMatcher protos.
	InclusionList *matcher.ListStringMatcher `protobuf:"bytes,3,opt,name=inclusion_list,json=inclusionList,proto3,oneof"`
}

func (*StatsMatcher_RejectAll) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_ExclusionList) isStatsMatcher_StatsMatcher() {}

func (*StatsMatcher_InclusionList) isStatsMatcher_StatsMatcher() {}

// Designates a tag name and value pair. The value may be either a fixed value
// or a regex providing the value via capture groups. The specified tag will be
// unconditionally set if a fixed value, otherwise it will only be set if one
// or more capture groups in the regex match.
type TagSpecifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Attaches an identifier to the tag values to identify the tag being in the
	// sink. Envoy has a set of default names and regexes to extract dynamic
	// portions of existing stats, which can be found in :repo:`well_known_names.h
	// <source/common/config/well_known_names.h>` in the Envoy repository. If a :ref:`tag_name
	// <envoy_api_field_config.metrics.v2.TagSpecifier.tag_name>` is provided in the config and
	// neither :ref:`regex <envoy_api_field_config.metrics.v2.TagSpecifier.regex>` or
	// :ref:`fixed_value <envoy_api_field_config.metrics.v2.TagSpecifier.fixed_value>` were specified,
	// Envoy will attempt to find that name in its set of defaults and use the accompanying regex.
	//
	// .. note::
	//
	//	It is invalid to specify the same tag name twice in a config.
	TagName string `protobuf:"bytes,1,opt,name=tag_name,json=tagName,proto3" json:"tag_name,omitempty"`
	// Types that are assignable to TagValue:
	//
	//	*TagSpecifier_Regex
	//	*TagSpecifier_FixedValue
	TagValue isTagSpecifier_TagValue `protobuf_oneof:"tag_value"`
}

func (x *TagSpecifier) Reset() {
	*x = TagSpecifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagSpecifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagSpecifier) ProtoMessage() {}

func (x *TagSpecifier) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagSpecifier.ProtoReflect.Descriptor instead.
func (*TagSpecifier) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v2_stats_proto_rawDescGZIP(), []int{3}
}

func (x *TagSpecifier) GetTagName() string {
	if x != nil {
		return x.TagName
	}
	return ""
}

func (m *TagSpecifier) GetTagValue() isTagSpecifier_TagValue {
	if m != nil {
		return m.TagValue
	}
	return nil
}

func (x *TagSpecifier) GetRegex() string {
	if x, ok := x.GetTagValue().(*TagSpecifier_Regex); ok {
		return x.Regex
	}
	return ""
}

func (x *TagSpecifier) GetFixedValue() string {
	if x, ok := x.GetTagValue().(*TagSpecifier_FixedValue); ok {
		return x.FixedValue
	}
	return ""
}

type isTagSpecifier_TagValue interface {
	isTagSpecifier_TagValue()
}

type TagSpecifier_Regex struct {
	// Designates a tag to strip from the tag extracted name and provide as a named
	// tag value for all statistics. This will only occur if any part of the name
	// matches the regex provided with one or more capture groups.
	//
	// The first capture group identifies the portion of the name to remove. The
	// second capture group (which will normally be nested inside the first) will
	// designate the value of the tag for the statistic. If no second capture
	// group is provided, the first will also be used to set the value of the tag.
	// All other capture groups will be ignored.
	//
	// Example 1. a stat name “cluster.foo_cluster.upstream_rq_timeout“ and
	// one tag specifier:
	//
	// .. code-block:: json
	//
	//	{
	//	  "tag_name": "envoy.cluster_name",
	//	  "regex": "^cluster\\.((.+?)\\.)"
	//	}
	//
	// Note that the regex will remove “foo_cluster.“ making the tag extracted
	// name “cluster.upstream_rq_timeout“ and the tag value for
	// “envoy.cluster_name“ will be “foo_cluster“ (note: there will be no
	// “.“ character because of the second capture group).
	//
	// Example 2. a stat name
	// “http.connection_manager_1.user_agent.ios.downstream_cx_total“ and two
	// tag specifiers:
	//
	// .. code-block:: json
	//
	//	[
	//	  {
	//	    "tag_name": "envoy.http_user_agent",
	//	    "regex": "^http(?=\\.).*?\\.user_agent\\.((.+?)\\.)\\w+?$"
	//	  },
	//	  {
	//	    "tag_name": "envoy.http_conn_manager_prefix",
	//	    "regex": "^http\\.((.*?)\\.)"
	//	  }
	//	]
	//
	// The two regexes of the specifiers will be processed in the definition order.
	//
	// The first regex will remove “ios.“, leaving the tag extracted name
	// “http.connection_manager_1.user_agent.downstream_cx_total“. The tag
	// “envoy.http_user_agent“ will be added with tag value “ios“.
	//
	// The second regex will remove “connection_manager_1.“ from the tag
	// extracted name produced by the first regex
	// “http.connection_manager_1.user_agent.downstream_cx_total“, leaving
	// “http.user_agent.downstream_cx_total“ as the tag extracted name. The tag
	// “envoy.http_conn_manager_prefix“ will be added with the tag value
	// “connection_manager_1“.
	Regex string `protobuf:"bytes,2,opt,name=regex,proto3,oneof"`
}

type TagSpecifier_FixedValue struct {
	// Specifies a fixed tag value for the “tag_name“.
	FixedValue string `protobuf:"bytes,3,opt,name=fixed_value,json=fixedValue,proto3,oneof"`
}

func (*TagSpecifier_Regex) isTagSpecifier_TagValue() {}

func (*TagSpecifier_FixedValue) isTagSpecifier_TagValue() {}

// Stats configuration proto schema for built-in *envoy.stat_sinks.statsd* sink. This sink does not support
// tagged metrics.
// [#extension: envoy.stat_sinks.statsd]
type StatsdSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to StatsdSpecifier:
	//
	//	*StatsdSink_Address
	//	*StatsdSink_TcpClusterName
	StatsdSpecifier isStatsdSink_StatsdSpecifier `protobuf_oneof:"statsd_specifier"`
	// Optional custom prefix for StatsdSink. If
	// specified, this will override the default prefix.
	// For example:
	//
	// .. code-block:: json
	//
	//	{
	//	  "prefix" : "envoy-prod"
	//	}
	//
	// will change emitted stats to
	//
	// .. code-block:: cpp
	//
	//	envoy-prod.test_counter:1|c
	//	envoy-prod.test_timer:5|ms
	//
	// Note that the default prefix, "envoy", will be used if a prefix is not
	// specified.
	//
	// Stats with default prefix:
	//
	// .. code-block:: cpp
	//
	//	envoy.test_counter:1|c
	//	envoy.test_timer:5|ms
	Prefix string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *StatsdSink) Reset() {
	*x = StatsdSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsdSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsdSink) ProtoMessage() {}

func (x *StatsdSink) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsdSink.ProtoReflect.Descriptor instead.
func (*StatsdSink) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v2_stats_proto_rawDescGZIP(), []int{4}
}

func (m *StatsdSink) GetStatsdSpecifier() isStatsdSink_StatsdSpecifier {
	if m != nil {
		return m.StatsdSpecifier
	}
	return nil
}

func (x *StatsdSink) GetAddress() *core.Address {
	if x, ok := x.GetStatsdSpecifier().(*StatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (x *StatsdSink) GetTcpClusterName() string {
	if x, ok := x.GetStatsdSpecifier().(*StatsdSink_TcpClusterName); ok {
		return x.TcpClusterName
	}
	return ""
}

func (x *StatsdSink) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type isStatsdSink_StatsdSpecifier interface {
	isStatsdSink_StatsdSpecifier()
}

type StatsdSink_Address struct {
	// The UDP address of a running `statsd <https://github.com/etsy/statsd>`_
	// compliant listener. If specified, statistics will be flushed to this
	// address.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type StatsdSink_TcpClusterName struct {
	// The name of a cluster that is running a TCP `statsd
	// <https://github.com/etsy/statsd>`_ compliant listener. If specified,
	// Envoy will connect to this cluster to flush statistics.
	TcpClusterName string `protobuf:"bytes,2,opt,name=tcp_cluster_name,json=tcpClusterName,proto3,oneof"`
}

func (*StatsdSink_Address) isStatsdSink_StatsdSpecifier() {}

func (*StatsdSink_TcpClusterName) isStatsdSink_StatsdSpecifier() {}

// Stats configuration proto schema for built-in *envoy.stat_sinks.dog_statsd* sink.
// The sink emits stats with `DogStatsD <https://docs.datadoghq.com/guides/dogstatsd/>`_
// compatible tags. Tags are configurable via :ref:`StatsConfig
// <envoy_api_msg_config.metrics.v2.StatsConfig>`.
// [#extension: envoy.stat_sinks.dog_statsd]
type DogStatsdSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to DogStatsdSpecifier:
	//
	//	*DogStatsdSink_Address
	DogStatsdSpecifier isDogStatsdSink_DogStatsdSpecifier `protobuf_oneof:"dog_statsd_specifier"`
	// Optional custom metric name prefix. See :ref:`StatsdSink's prefix field
	// <envoy_api_field_config.metrics.v2.StatsdSink.prefix>` for more details.
	Prefix string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *DogStatsdSink) Reset() {
	*x = DogStatsdSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DogStatsdSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DogStatsdSink) ProtoMessage() {}

func (x *DogStatsdSink) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DogStatsdSink.ProtoReflect.Descriptor instead.
func (*DogStatsdSink) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v2_stats_proto_rawDescGZIP(), []int{5}
}

func (m *DogStatsdSink) GetDogStatsdSpecifier() isDogStatsdSink_DogStatsdSpecifier {
	if m != nil {
		return m.DogStatsdSpecifier
	}
	return nil
}

func (x *DogStatsdSink) GetAddress() *core.Address {
	if x, ok := x.GetDogStatsdSpecifier().(*DogStatsdSink_Address); ok {
		return x.Address
	}
	return nil
}

func (x *DogStatsdSink) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type isDogStatsdSink_DogStatsdSpecifier interface {
	isDogStatsdSink_DogStatsdSpecifier()
}

type DogStatsdSink_Address struct {
	// The UDP address of a running DogStatsD compliant listener. If specified,
	// statistics will be flushed to this address.
	Address *core.Address `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

func (*DogStatsdSink_Address) isDogStatsdSink_DogStatsdSpecifier() {}

// Stats configuration proto schema for built-in *envoy.stat_sinks.hystrix* sink.
// The sink emits stats in `text/event-stream
// <https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events>`_
// formatted stream for use by `Hystrix dashboard
// <https://github.com/Netflix-Skunkworks/hystrix-dashboard/wiki>`_.
//
// Note that only a single HystrixSink should be configured.
//
// Streaming is started through an admin endpoint :http:get:`/hystrix_event_stream`.
// [#extension: envoy.stat_sinks.hystrix]
type HystrixSink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of buckets the rolling statistical window is divided into.
	//
	// Each time the sink is flushed, all relevant Envoy statistics are sampled and
	// added to the rolling window (removing the oldest samples in the window
	// in the process). The sink then outputs the aggregate statistics across the
	// current rolling window to the event stream(s).
	//
	// rolling_window(ms) = stats_flush_interval(ms) * num_of_buckets
	//
	// More detailed explanation can be found in `Hystrix wiki
	// <https://github.com/Netflix/Hystrix/wiki/Metrics-and-Monitoring#hystrixrollingnumber>`_.
	NumBuckets int64 `protobuf:"varint,1,opt,name=num_buckets,json=numBuckets,proto3" json:"num_buckets,omitempty"`
}

func (x *HystrixSink) Reset() {
	*x = HystrixSink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HystrixSink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HystrixSink) ProtoMessage() {}

func (x *HystrixSink) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_metrics_v2_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HystrixSink.ProtoReflect.Descriptor instead.
func (*HystrixSink) Descriptor() ([]byte, []int) {
	return file_envoy_config_metrics_v2_stats_proto_rawDescGZIP(), []int{6}
}

func (x *HystrixSink) GetNumBuckets() int64 {
	if x != nil {
		return x.NumBuckets
	}
	return 0
}

var File_envoy_config_metrics_v2_stats_proto protoreflect.FileDescriptor

var file_envoy_config_metrics_v2_stats_proto_rawDesc = []byte{
	0x0a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1f,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x69, 0x6e, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x02, 0x18, 0x01,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x0c, 0x74, 0x79,
	0x70, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x54, 0x61, 0x67, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x73, 0x54, 0x61, 0x67, 0x73, 0x12, 0x4b, 0x0a, 0x14, 0x75, 0x73,
	0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x75, 0x73, 0x65, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x4a, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x22, 0xe5, 0x01, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0a, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x61,
	0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09, 0x72, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x4e, 0x0a, 0x0e, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4e, 0x0a, 0x0e, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x14, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x7b, 0x0a, 0x0c, 0x54,
	0x61, 0x67, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x61, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x61, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x28, 0x80, 0x08, 0x48,
	0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0b, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0a, 0x66, 0x69, 0x78, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x74,
	0x61, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x64, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x36, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2a, 0x0a, 0x10, 0x74, 0x63, 0x70, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0e, 0x74, 0x63, 0x70,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x42, 0x17, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x82, 0x01, 0x0a,
	0x0d, 0x44, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x73, 0x64, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x36,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0x1b,
	0x0a, 0x14, 0x64, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x64, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x4a, 0x04, 0x08, 0x02, 0x10,
	0x03, 0x22, 0x2e, 0x0a, 0x0b, 0x48, 0x79, 0x73, 0x74, 0x72, 0x69, 0x78, 0x53, 0x69, 0x6e, 0x6b,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x73, 0x42, 0x87, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x0a, 0x25, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x32, 0x42, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x32, 0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_metrics_v2_stats_proto_rawDescOnce sync.Once
	file_envoy_config_metrics_v2_stats_proto_rawDescData = file_envoy_config_metrics_v2_stats_proto_rawDesc
)

func file_envoy_config_metrics_v2_stats_proto_rawDescGZIP() []byte {
	file_envoy_config_metrics_v2_stats_proto_rawDescOnce.Do(func() {
		file_envoy_config_metrics_v2_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_metrics_v2_stats_proto_rawDescData)
	})
	return file_envoy_config_metrics_v2_stats_proto_rawDescData
}

var file_envoy_config_metrics_v2_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_envoy_config_metrics_v2_stats_proto_goTypes = []interface{}{
	(*StatsSink)(nil),                 // 0: envoy.config.metrics.v2.StatsSink
	(*StatsConfig)(nil),               // 1: envoy.config.metrics.v2.StatsConfig
	(*StatsMatcher)(nil),              // 2: envoy.config.metrics.v2.StatsMatcher
	(*TagSpecifier)(nil),              // 3: envoy.config.metrics.v2.TagSpecifier
	(*StatsdSink)(nil),                // 4: envoy.config.metrics.v2.StatsdSink
	(*DogStatsdSink)(nil),             // 5: envoy.config.metrics.v2.DogStatsdSink
	(*HystrixSink)(nil),               // 6: envoy.config.metrics.v2.HystrixSink
	(*structpb.Struct)(nil),           // 7: google.protobuf.Struct
	(*anypb.Any)(nil),                 // 8: google.protobuf.Any
	(*wrapperspb.BoolValue)(nil),      // 9: google.protobuf.BoolValue
	(*matcher.ListStringMatcher)(nil), // 10: envoy.type.matcher.ListStringMatcher
	(*core.Address)(nil),              // 11: envoy.api.v2.core.Address
}
var file_envoy_config_metrics_v2_stats_proto_depIdxs = []int32{
	7,  // 0: envoy.config.metrics.v2.StatsSink.config:type_name -> google.protobuf.Struct
	8,  // 1: envoy.config.metrics.v2.StatsSink.typed_config:type_name -> google.protobuf.Any
	3,  // 2: envoy.config.metrics.v2.StatsConfig.stats_tags:type_name -> envoy.config.metrics.v2.TagSpecifier
	9,  // 3: envoy.config.metrics.v2.StatsConfig.use_all_default_tags:type_name -> google.protobuf.BoolValue
	2,  // 4: envoy.config.metrics.v2.StatsConfig.stats_matcher:type_name -> envoy.config.metrics.v2.StatsMatcher
	10, // 5: envoy.config.metrics.v2.StatsMatcher.exclusion_list:type_name -> envoy.type.matcher.ListStringMatcher
	10, // 6: envoy.config.metrics.v2.StatsMatcher.inclusion_list:type_name -> envoy.type.matcher.ListStringMatcher
	11, // 7: envoy.config.metrics.v2.StatsdSink.address:type_name -> envoy.api.v2.core.Address
	11, // 8: envoy.config.metrics.v2.DogStatsdSink.address:type_name -> envoy.api.v2.core.Address
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_envoy_config_metrics_v2_stats_proto_init() }
func file_envoy_config_metrics_v2_stats_proto_init() {
	if File_envoy_config_metrics_v2_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_metrics_v2_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsSink); i {
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
		file_envoy_config_metrics_v2_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsConfig); i {
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
		file_envoy_config_metrics_v2_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsMatcher); i {
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
		file_envoy_config_metrics_v2_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagSpecifier); i {
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
		file_envoy_config_metrics_v2_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsdSink); i {
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
		file_envoy_config_metrics_v2_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DogStatsdSink); i {
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
		file_envoy_config_metrics_v2_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HystrixSink); i {
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
	file_envoy_config_metrics_v2_stats_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StatsSink_Config)(nil),
		(*StatsSink_TypedConfig)(nil),
	}
	file_envoy_config_metrics_v2_stats_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*StatsMatcher_RejectAll)(nil),
		(*StatsMatcher_ExclusionList)(nil),
		(*StatsMatcher_InclusionList)(nil),
	}
	file_envoy_config_metrics_v2_stats_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TagSpecifier_Regex)(nil),
		(*TagSpecifier_FixedValue)(nil),
	}
	file_envoy_config_metrics_v2_stats_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*StatsdSink_Address)(nil),
		(*StatsdSink_TcpClusterName)(nil),
	}
	file_envoy_config_metrics_v2_stats_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*DogStatsdSink_Address)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_metrics_v2_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_metrics_v2_stats_proto_goTypes,
		DependencyIndexes: file_envoy_config_metrics_v2_stats_proto_depIdxs,
		MessageInfos:      file_envoy_config_metrics_v2_stats_proto_msgTypes,
	}.Build()
	File_envoy_config_metrics_v2_stats_proto = out.File
	file_envoy_config_metrics_v2_stats_proto_rawDesc = nil
	file_envoy_config_metrics_v2_stats_proto_goTypes = nil
	file_envoy_config_metrics_v2_stats_proto_depIdxs = nil
}
