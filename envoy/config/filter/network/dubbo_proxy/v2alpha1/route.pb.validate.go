// Code generated by protoc-gen-validate
// source: envoy/config/filter/network/dubbo_proxy/v2alpha1/route.proto
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

// Validate checks the field values on RouteConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RouteConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Interface

	// no validation rules for Group

	// no validation rules for Version

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					Field:  fmt.Sprintf("Routes[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteConfigurationValidationError is the validation error returned by
// RouteConfiguration.Validate if the designated constraints aren't met.
type RouteConfigurationValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RouteConfigurationValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteConfiguration.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RouteConfigurationValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Route) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMatch()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				Field:  "Match",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRoute()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				Field:  "Route",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RouteValidationError{}

// Validate checks the field values on MethodMatch with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MethodMatch) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetName()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return MethodMatchValidationError{
				Field:  "Name",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for ParamsMatch

	return nil
}

// MethodMatchValidationError is the validation error returned by
// MethodMatch.Validate if the designated constraints aren't met.
type MethodMatchValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e MethodMatchValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMethodMatch.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = MethodMatchValidationError{}

// Validate checks the field values on RouteMatch with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RouteMatch) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMethod()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return RouteMatchValidationError{
				Field:  "Method",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return RouteMatchValidationError{
					Field:  fmt.Sprintf("Headers[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// RouteMatchValidationError is the validation error returned by
// RouteMatch.Validate if the designated constraints aren't met.
type RouteMatchValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RouteMatchValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteMatch.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RouteMatchValidationError{}

// Validate checks the field values on RouteAction with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RouteAction) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ClusterSpecifier.(type) {

	case *RouteAction_Cluster:
		// no validation rules for Cluster

	case *RouteAction_WeightedClusters:

		if v, ok := interface{}(m.GetWeightedClusters()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return RouteActionValidationError{
					Field:  "WeightedClusters",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return RouteActionValidationError{
			Field:  "ClusterSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// RouteActionValidationError is the validation error returned by
// RouteAction.Validate if the designated constraints aren't met.
type RouteActionValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RouteActionValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteAction.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RouteActionValidationError{}

// Validate checks the field values on MethodMatch_ParameterMatchSpecifier with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *MethodMatch_ParameterMatchSpecifier) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ParameterMatchSpecifier.(type) {

	case *MethodMatch_ParameterMatchSpecifier_ExactMatch:
		// no validation rules for ExactMatch

	case *MethodMatch_ParameterMatchSpecifier_RangeMatch:

		if v, ok := interface{}(m.GetRangeMatch()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return MethodMatch_ParameterMatchSpecifierValidationError{
					Field:  "RangeMatch",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// MethodMatch_ParameterMatchSpecifierValidationError is the validation error
// returned by MethodMatch_ParameterMatchSpecifier.Validate if the designated
// constraints aren't met.
type MethodMatch_ParameterMatchSpecifierValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e MethodMatch_ParameterMatchSpecifierValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMethodMatch_ParameterMatchSpecifier.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = MethodMatch_ParameterMatchSpecifierValidationError{}
