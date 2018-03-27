// Code generated by protoc-gen-validate
// source: envoy/service/discovery/v2/hds.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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

	_ = types.DynamicAny{}
)

// Validate checks the field values on Capability with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Capability) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CapabilityValidationError is the validation error returned by
// Capability.Validate if the designated constraints aren't met.
type CapabilityValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CapabilityValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCapability.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CapabilityValidationError{}

// Validate checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckRequestValidationError{
				Field:  "Node",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCapability()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckRequestValidationError{
				Field:  "Capability",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HealthCheckRequestValidationError is the validation error returned by
// HealthCheckRequest.Validate if the designated constraints aren't met.
type HealthCheckRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HealthCheckRequestValidationError{}

// Validate checks the field values on EndpointHealth with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EndpointHealth) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EndpointHealthValidationError{
				Field:  "Endpoint",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for HealthStatus

	return nil
}

// EndpointHealthValidationError is the validation error returned by
// EndpointHealth.Validate if the designated constraints aren't met.
type EndpointHealthValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e EndpointHealthValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointHealth.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = EndpointHealthValidationError{}

// Validate checks the field values on EndpointHealthResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EndpointHealthResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEndpointsHealth() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EndpointHealthResponseValidationError{
					Field:  fmt.Sprintf("EndpointsHealth[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// EndpointHealthResponseValidationError is the validation error returned by
// EndpointHealthResponse.Validate if the designated constraints aren't met.
type EndpointHealthResponseValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e EndpointHealthResponseValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointHealthResponse.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = EndpointHealthResponseValidationError{}

// Validate checks the field values on
// HealthCheckRequestOrEndpointHealthResponse with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HealthCheckRequestOrEndpointHealthResponse) Validate() error {
	if m == nil {
		return nil
	}

	switch m.RequestType.(type) {

	case *HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest:

		if v, ok := interface{}(m.GetHealthCheckRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckRequestOrEndpointHealthResponseValidationError{
					Field:  "HealthCheckRequest",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse:

		if v, ok := interface{}(m.GetEndpointHealthResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckRequestOrEndpointHealthResponseValidationError{
					Field:  "EndpointHealthResponse",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// HealthCheckRequestOrEndpointHealthResponseValidationError is the validation
// error returned by HealthCheckRequestOrEndpointHealthResponse.Validate if
// the designated constraints aren't met.
type HealthCheckRequestOrEndpointHealthResponseValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequestOrEndpointHealthResponse.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HealthCheckRequestOrEndpointHealthResponseValidationError{}

// Validate checks the field values on LocalityEndpoints with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LocalityEndpoints) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityEndpointsValidationError{
				Field:  "Locality",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityEndpointsValidationError{
					Field:  fmt.Sprintf("Endpoints[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// LocalityEndpointsValidationError is the validation error returned by
// LocalityEndpoints.Validate if the designated constraints aren't met.
type LocalityEndpointsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e LocalityEndpointsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityEndpoints.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = LocalityEndpointsValidationError{}

// Validate checks the field values on ClusterHealthCheck with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ClusterName

	for idx, item := range m.GetHealthChecks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterHealthCheckValidationError{
					Field:  fmt.Sprintf("HealthChecks[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterHealthCheckValidationError{
					Field:  fmt.Sprintf("Endpoints[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// ClusterHealthCheckValidationError is the validation error returned by
// ClusterHealthCheck.Validate if the designated constraints aren't met.
type ClusterHealthCheckValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ClusterHealthCheckValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterHealthCheck.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ClusterHealthCheckValidationError{}

// Validate checks the field values on HealthCheckSpecifier with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckSpecifier) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHealthCheck() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckSpecifierValidationError{
					Field:  fmt.Sprintf("HealthCheck[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckSpecifierValidationError{
				Field:  "Interval",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// HealthCheckSpecifierValidationError is the validation error returned by
// HealthCheckSpecifier.Validate if the designated constraints aren't met.
type HealthCheckSpecifierValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e HealthCheckSpecifierValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckSpecifier.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = HealthCheckSpecifierValidationError{}
