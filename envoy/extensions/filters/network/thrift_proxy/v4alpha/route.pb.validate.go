// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/thrift_proxy/v4alpha/route.proto

package envoy_extensions_filters_network_thrift_proxy_v4alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on RouteConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RouteConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					field:  fmt.Sprintf("Routes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteConfigurationValidationError is the validation error returned by
// RouteConfiguration.Validate if the designated constraints aren't met.
type RouteConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteConfigurationValidationError) ErrorName() string {
	return "RouteConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e RouteConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteConfigurationValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Route) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMatch() == nil {
		return RouteValidationError{
			field:  "Match",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetRoute() == nil {
		return RouteValidationError{
			field:  "Route",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetRoute()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Route",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteValidationError) ErrorName() string { return "RouteValidationError" }

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteValidationError{}

// Validate checks the field values on RouteMatch with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RouteMatch) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Invert

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteMatchValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.MatchSpecifier.(type) {

	case *RouteMatch_MethodName:
		// no validation rules for MethodName

	case *RouteMatch_ServiceName:
		// no validation rules for ServiceName

	default:
		return RouteMatchValidationError{
			field:  "MatchSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RouteMatchValidationError is the validation error returned by
// RouteMatch.Validate if the designated constraints aren't met.
type RouteMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteMatchValidationError) ErrorName() string { return "RouteMatchValidationError" }

// Error satisfies the builtin error interface
func (e RouteMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteMatchValidationError{}

// Validate checks the field values on RouteAction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RouteAction) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteActionValidationError{
				field:  "MetadataMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetRateLimits() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteActionValidationError{
					field:  fmt.Sprintf("RateLimits[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for StripServiceName

	for idx, item := range m.GetRequestMirrorPolicies() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteActionValidationError{
					field:  fmt.Sprintf("RequestMirrorPolicies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.ClusterSpecifier.(type) {

	case *RouteAction_Cluster:

		if utf8.RuneCountInString(m.GetCluster()) < 1 {
			return RouteActionValidationError{
				field:  "Cluster",
				reason: "value length must be at least 1 runes",
			}
		}

	case *RouteAction_WeightedClusters:

		if v, ok := interface{}(m.GetWeightedClusters()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteActionValidationError{
					field:  "WeightedClusters",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *RouteAction_ClusterHeader:

		if utf8.RuneCountInString(m.GetClusterHeader()) < 1 {
			return RouteActionValidationError{
				field:  "ClusterHeader",
				reason: "value length must be at least 1 runes",
			}
		}

		if !_RouteAction_ClusterHeader_Pattern.MatchString(m.GetClusterHeader()) {
			return RouteActionValidationError{
				field:  "ClusterHeader",
				reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
			}
		}

	default:
		return RouteActionValidationError{
			field:  "ClusterSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RouteActionValidationError is the validation error returned by
// RouteAction.Validate if the designated constraints aren't met.
type RouteActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteActionValidationError) ErrorName() string { return "RouteActionValidationError" }

// Error satisfies the builtin error interface
func (e RouteActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteActionValidationError{}

var _RouteAction_ClusterHeader_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

// Validate checks the field values on WeightedCluster with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *WeightedCluster) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetClusters()) < 1 {
		return WeightedClusterValidationError{
			field:  "Clusters",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetClusters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WeightedClusterValidationError{
					field:  fmt.Sprintf("Clusters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// WeightedClusterValidationError is the validation error returned by
// WeightedCluster.Validate if the designated constraints aren't met.
type WeightedClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WeightedClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WeightedClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WeightedClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WeightedClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WeightedClusterValidationError) ErrorName() string { return "WeightedClusterValidationError" }

// Error satisfies the builtin error interface
func (e WeightedClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWeightedCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WeightedClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WeightedClusterValidationError{}

// Validate checks the field values on RouteAction_RequestMirrorPolicy with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RouteAction_RequestMirrorPolicy) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetCluster()) < 1 {
		return RouteAction_RequestMirrorPolicyValidationError{
			field:  "Cluster",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetRuntimeFraction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteAction_RequestMirrorPolicyValidationError{
				field:  "RuntimeFraction",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RouteAction_RequestMirrorPolicyValidationError is the validation error
// returned by RouteAction_RequestMirrorPolicy.Validate if the designated
// constraints aren't met.
type RouteAction_RequestMirrorPolicyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteAction_RequestMirrorPolicyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteAction_RequestMirrorPolicyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteAction_RequestMirrorPolicyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteAction_RequestMirrorPolicyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteAction_RequestMirrorPolicyValidationError) ErrorName() string {
	return "RouteAction_RequestMirrorPolicyValidationError"
}

// Error satisfies the builtin error interface
func (e RouteAction_RequestMirrorPolicyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteAction_RequestMirrorPolicy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteAction_RequestMirrorPolicyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteAction_RequestMirrorPolicyValidationError{}

// Validate checks the field values on WeightedCluster_ClusterWeight with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *WeightedCluster_ClusterWeight) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return WeightedCluster_ClusterWeightValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if wrapper := m.GetWeight(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return WeightedCluster_ClusterWeightValidationError{
				field:  "Weight",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WeightedCluster_ClusterWeightValidationError{
				field:  "MetadataMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// WeightedCluster_ClusterWeightValidationError is the validation error
// returned by WeightedCluster_ClusterWeight.Validate if the designated
// constraints aren't met.
type WeightedCluster_ClusterWeightValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WeightedCluster_ClusterWeightValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WeightedCluster_ClusterWeightValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WeightedCluster_ClusterWeightValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WeightedCluster_ClusterWeightValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WeightedCluster_ClusterWeightValidationError) ErrorName() string {
	return "WeightedCluster_ClusterWeightValidationError"
}

// Error satisfies the builtin error interface
func (e WeightedCluster_ClusterWeightValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWeightedCluster_ClusterWeight.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WeightedCluster_ClusterWeightValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WeightedCluster_ClusterWeightValidationError{}
