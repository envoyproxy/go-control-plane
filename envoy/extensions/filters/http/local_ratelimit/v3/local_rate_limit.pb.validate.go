// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/local_ratelimit/v3/local_rate_limit.proto

package local_ratelimitv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3"
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
	_ = sort.Sort

	_ = v3.XRateLimitHeadersRFCVersion(0)
)

// Validate checks the field values on LocalRateLimit with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LocalRateLimit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalRateLimit with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LocalRateLimitMultiError,
// or nil if none found.
func (m *LocalRateLimit) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalRateLimit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := LocalRateLimitValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTokenBucket()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "TokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "TokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokenBucket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "TokenBucket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFilterEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "FilterEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "FilterEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilterEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "FilterEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFilterEnforced()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "FilterEnforced",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "FilterEnforced",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilterEnforced()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "FilterEnforced",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetRequestHeadersToAddWhenNotEnforced()) > 10 {
		err := LocalRateLimitValidationError{
			field:  "RequestHeadersToAddWhenNotEnforced",
			reason: "value must contain no more than 10 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetRequestHeadersToAddWhenNotEnforced() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalRateLimitValidationError{
						field:  fmt.Sprintf("RequestHeadersToAddWhenNotEnforced[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalRateLimitValidationError{
						field:  fmt.Sprintf("RequestHeadersToAddWhenNotEnforced[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalRateLimitValidationError{
					field:  fmt.Sprintf("RequestHeadersToAddWhenNotEnforced[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(m.GetResponseHeadersToAdd()) > 10 {
		err := LocalRateLimitValidationError{
			field:  "ResponseHeadersToAdd",
			reason: "value must contain no more than 10 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetResponseHeadersToAdd() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalRateLimitValidationError{
						field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalRateLimitValidationError{
						field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalRateLimitValidationError{
					field:  fmt.Sprintf("ResponseHeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetDescriptors() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalRateLimitValidationError{
						field:  fmt.Sprintf("Descriptors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalRateLimitValidationError{
						field:  fmt.Sprintf("Descriptors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalRateLimitValidationError{
					field:  fmt.Sprintf("Descriptors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetStage() > 10 {
		err := LocalRateLimitValidationError{
			field:  "Stage",
			reason: "value must be less than or equal to 10",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for LocalRateLimitPerDownstreamConnection

	if _, ok := v3.XRateLimitHeadersRFCVersion_name[int32(m.GetEnableXRatelimitHeaders())]; !ok {
		err := LocalRateLimitValidationError{
			field:  "EnableXRatelimitHeaders",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := v3.VhRateLimitsOptions_name[int32(m.GetVhRateLimits())]; !ok {
		err := LocalRateLimitValidationError{
			field:  "VhRateLimits",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAlwaysConsumeDefaultTokenBucket()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "AlwaysConsumeDefaultTokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitValidationError{
					field:  "AlwaysConsumeDefaultTokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAlwaysConsumeDefaultTokenBucket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitValidationError{
				field:  "AlwaysConsumeDefaultTokenBucket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LocalRateLimitMultiError(errors)
	}

	return nil
}

// LocalRateLimitMultiError is an error wrapping multiple validation errors
// returned by LocalRateLimit.ValidateAll() if the designated constraints
// aren't met.
type LocalRateLimitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalRateLimitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalRateLimitMultiError) AllErrors() []error { return m }

// LocalRateLimitValidationError is the validation error returned by
// LocalRateLimit.Validate if the designated constraints aren't met.
type LocalRateLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalRateLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalRateLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalRateLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalRateLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalRateLimitValidationError) ErrorName() string { return "LocalRateLimitValidationError" }

// Error satisfies the builtin error interface
func (e LocalRateLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalRateLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalRateLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalRateLimitValidationError{}
