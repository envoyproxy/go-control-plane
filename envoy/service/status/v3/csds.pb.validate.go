// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/status/v3/csds.proto

package envoy_service_status_v3

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

	"github.com/golang/protobuf/ptypes"

	v3 "github.com/envoyproxy/go-control-plane/envoy/admin/v3"
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
	_ = ptypes.DynamicAny{}

	_ = v3.ClientResourceStatus(0)
)

// Validate checks the field values on ClientStatusRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClientStatusRequest) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNodeMatchers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientStatusRequestValidationError{
					field:  fmt.Sprintf("NodeMatchers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientStatusRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ClientStatusRequestValidationError is the validation error returned by
// ClientStatusRequest.Validate if the designated constraints aren't met.
type ClientStatusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientStatusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientStatusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientStatusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientStatusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientStatusRequestValidationError) ErrorName() string {
	return "ClientStatusRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ClientStatusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientStatusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientStatusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientStatusRequestValidationError{}

// Validate checks the field values on PerXdsConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PerXdsConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for ClientStatus

	switch m.PerXdsConfig.(type) {

	case *PerXdsConfig_ListenerConfig:

		if v, ok := interface{}(m.GetListenerConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PerXdsConfigValidationError{
					field:  "ListenerConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PerXdsConfig_ClusterConfig:

		if v, ok := interface{}(m.GetClusterConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PerXdsConfigValidationError{
					field:  "ClusterConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PerXdsConfig_RouteConfig:

		if v, ok := interface{}(m.GetRouteConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PerXdsConfigValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PerXdsConfig_ScopedRouteConfig:

		if v, ok := interface{}(m.GetScopedRouteConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PerXdsConfigValidationError{
					field:  "ScopedRouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *PerXdsConfig_EndpointConfig:

		if v, ok := interface{}(m.GetEndpointConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PerXdsConfigValidationError{
					field:  "EndpointConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PerXdsConfigValidationError is the validation error returned by
// PerXdsConfig.Validate if the designated constraints aren't met.
type PerXdsConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PerXdsConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PerXdsConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PerXdsConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PerXdsConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PerXdsConfigValidationError) ErrorName() string { return "PerXdsConfigValidationError" }

// Error satisfies the builtin error interface
func (e PerXdsConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerXdsConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PerXdsConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PerXdsConfigValidationError{}

// Validate checks the field values on ClientConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClientConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientConfigValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetXdsConfig() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientConfigValidationError{
					field:  fmt.Sprintf("XdsConfig[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetGenericXdsConfigs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientConfigValidationError{
					field:  fmt.Sprintf("GenericXdsConfigs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ClientConfigValidationError is the validation error returned by
// ClientConfig.Validate if the designated constraints aren't met.
type ClientConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientConfigValidationError) ErrorName() string { return "ClientConfigValidationError" }

// Error satisfies the builtin error interface
func (e ClientConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientConfigValidationError{}

// Validate checks the field values on ClientStatusResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClientStatusResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetConfig() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientStatusResponseValidationError{
					field:  fmt.Sprintf("Config[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ClientStatusResponseValidationError is the validation error returned by
// ClientStatusResponse.Validate if the designated constraints aren't met.
type ClientStatusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientStatusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientStatusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientStatusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientStatusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientStatusResponseValidationError) ErrorName() string {
	return "ClientStatusResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ClientStatusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientStatusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientStatusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientStatusResponseValidationError{}

// Validate checks the field values on ClientConfig_GenericXdsConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClientConfig_GenericXdsConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TypeUrl

	// no validation rules for Name

	// no validation rules for VersionInfo

	if v, ok := interface{}(m.GetXdsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientConfig_GenericXdsConfigValidationError{
				field:  "XdsConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLastUpdated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientConfig_GenericXdsConfigValidationError{
				field:  "LastUpdated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ConfigStatus

	// no validation rules for ClientStatus

	if v, ok := interface{}(m.GetErrorState()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientConfig_GenericXdsConfigValidationError{
				field:  "ErrorState",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsStaticResource

	return nil
}

// ClientConfig_GenericXdsConfigValidationError is the validation error
// returned by ClientConfig_GenericXdsConfig.Validate if the designated
// constraints aren't met.
type ClientConfig_GenericXdsConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientConfig_GenericXdsConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientConfig_GenericXdsConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientConfig_GenericXdsConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientConfig_GenericXdsConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientConfig_GenericXdsConfigValidationError) ErrorName() string {
	return "ClientConfig_GenericXdsConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ClientConfig_GenericXdsConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientConfig_GenericXdsConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientConfig_GenericXdsConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientConfig_GenericXdsConfigValidationError{}
