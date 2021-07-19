// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/dynamic_forward_proxy/v4alpha/dns_cache.proto

package envoy_extensions_common_dynamic_forward_proxy_v4alpha

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

	v4alpha "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v4alpha"
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

	_ = v4alpha.Cluster_DnsLookupFamily(0)
)

// Validate checks the field values on DnsCacheCircuitBreakers with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DnsCacheCircuitBreakers) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMaxPendingRequests()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheCircuitBreakersValidationError{
				field:  "MaxPendingRequests",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DnsCacheCircuitBreakersValidationError is the validation error returned by
// DnsCacheCircuitBreakers.Validate if the designated constraints aren't met.
type DnsCacheCircuitBreakersValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsCacheCircuitBreakersValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsCacheCircuitBreakersValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsCacheCircuitBreakersValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsCacheCircuitBreakersValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsCacheCircuitBreakersValidationError) ErrorName() string {
	return "DnsCacheCircuitBreakersValidationError"
}

// Error satisfies the builtin error interface
func (e DnsCacheCircuitBreakersValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsCacheCircuitBreakers.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsCacheCircuitBreakersValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsCacheCircuitBreakersValidationError{}

// Validate checks the field values on DnsCacheConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DnsCacheConfig) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return DnsCacheConfigValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if _, ok := v4alpha.Cluster_DnsLookupFamily_name[int32(m.GetDnsLookupFamily())]; !ok {
		return DnsCacheConfigValidationError{
			field:  "DnsLookupFamily",
			reason: "value must be one of the defined enum values",
		}
	}

	if d := m.GetDnsRefreshRate(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsRefreshRate",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

		if dur < gte {
			return DnsCacheConfigValidationError{
				field:  "DnsRefreshRate",
				reason: "value must be greater than or equal to 1ms",
			}
		}

	}

	if d := m.GetHostTtl(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			return DnsCacheConfigValidationError{
				field:  "HostTtl",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return DnsCacheConfigValidationError{
				field:  "HostTtl",
				reason: "value must be greater than 0s",
			}
		}

	}

	if wrapper := m.GetMaxHosts(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			return DnsCacheConfigValidationError{
				field:  "MaxHosts",
				reason: "value must be greater than 0",
			}
		}

	}

	if v, ok := interface{}(m.GetDnsFailureRefreshRate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsFailureRefreshRate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDnsCacheCircuitBreaker()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsCacheCircuitBreaker",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HiddenEnvoyDeprecatedUseTcpForDnsLookups

	if v, ok := interface{}(m.GetDnsResolutionConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsResolutionConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTypedDnsResolverConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DnsCacheConfigValidationError{
				field:  "TypedDnsResolverConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetPreresolveHostnames() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DnsCacheConfigValidationError{
					field:  fmt.Sprintf("PreresolveHostnames[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if d := m.GetDnsQueryTimeout(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			return DnsCacheConfigValidationError{
				field:  "DnsQueryTimeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return DnsCacheConfigValidationError{
				field:  "DnsQueryTimeout",
				reason: "value must be greater than 0s",
			}
		}

	}

	return nil
}

// DnsCacheConfigValidationError is the validation error returned by
// DnsCacheConfig.Validate if the designated constraints aren't met.
type DnsCacheConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DnsCacheConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DnsCacheConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DnsCacheConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DnsCacheConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DnsCacheConfigValidationError) ErrorName() string { return "DnsCacheConfigValidationError" }

// Error satisfies the builtin error interface
func (e DnsCacheConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDnsCacheConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DnsCacheConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DnsCacheConfigValidationError{}
