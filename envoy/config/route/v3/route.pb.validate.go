// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/route/v3/route.proto

package envoy_config_route_v3

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

	for idx, item := range m.GetVirtualHosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					field:  fmt.Sprintf("VirtualHosts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetVhds()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteConfigurationValidationError{
				field:  "Vhds",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetInternalOnlyHeaders() {
		_, _ = idx, item

		if !_RouteConfiguration_InternalOnlyHeaders_Pattern.MatchString(item) {
			return RouteConfigurationValidationError{
				field:  fmt.Sprintf("InternalOnlyHeaders[%v]", idx),
				reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
			}
		}

	}

	if len(m.GetResponseHeadersToAdd()) > 1000 {
		return RouteConfigurationValidationError{
			field:  "ResponseHeadersToAdd",
			reason: "value must contain no more than 1000 item(s)",
		}
	}

	for idx, item := range m.GetResponseHeadersToAdd() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetResponseHeadersToRemove() {
		_, _ = idx, item

		if !_RouteConfiguration_ResponseHeadersToRemove_Pattern.MatchString(item) {
			return RouteConfigurationValidationError{
				field:  fmt.Sprintf("ResponseHeadersToRemove[%v]", idx),
				reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
			}
		}

	}

	if len(m.GetRequestHeadersToAdd()) > 1000 {
		return RouteConfigurationValidationError{
			field:  "RequestHeadersToAdd",
			reason: "value must contain no more than 1000 item(s)",
		}
	}

	for idx, item := range m.GetRequestHeadersToAdd() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					field:  fmt.Sprintf("RequestHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRequestHeadersToRemove() {
		_, _ = idx, item

		if !_RouteConfiguration_RequestHeadersToRemove_Pattern.MatchString(item) {
			return RouteConfigurationValidationError{
				field:  fmt.Sprintf("RequestHeadersToRemove[%v]", idx),
				reason: "value does not match regex pattern \"^[^\\x00\\n\\r]*$\"",
			}
		}

	}

	// no validation rules for MostSpecificHeaderMutationsWins

	if v, ok := interface{}(m.GetValidateClusters()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteConfigurationValidationError{
				field:  "ValidateClusters",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxDirectResponseBodySizeBytes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteConfigurationValidationError{
				field:  "MaxDirectResponseBodySizeBytes",
				reason: "embedded message failed validation",
				cause:  err,
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

var _RouteConfiguration_InternalOnlyHeaders_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

var _RouteConfiguration_ResponseHeadersToRemove_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

var _RouteConfiguration_RequestHeadersToRemove_Pattern = regexp.MustCompile("^[^\x00\n\r]*$")

// Validate checks the field values on Vhds with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Vhds) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetConfigSource() == nil {
		return VhdsValidationError{
			field:  "ConfigSource",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetConfigSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VhdsValidationError{
				field:  "ConfigSource",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// VhdsValidationError is the validation error returned by Vhds.Validate if the
// designated constraints aren't met.
type VhdsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VhdsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VhdsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VhdsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VhdsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VhdsValidationError) ErrorName() string { return "VhdsValidationError" }

// Error satisfies the builtin error interface
func (e VhdsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVhds.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VhdsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VhdsValidationError{}
