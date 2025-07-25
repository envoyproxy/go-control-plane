//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/tap/v3/tap.proto

package tapv3

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
)

// Validate checks the field values on Tap with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Tap) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tap with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TapMultiError, or nil if none found.
func (m *Tap) ValidateAll() error {
	return m.validate(true)
}

func (m *Tap) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCommonConfig() == nil {
		err := TapValidationError{
			field:  "CommonConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCommonConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TapValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TapValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TapValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTransportSocket() == nil {
		err := TapValidationError{
			field:  "TransportSocket",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTransportSocket()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TapValidationError{
					field:  "TransportSocket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TapValidationError{
					field:  "TransportSocket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTransportSocket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TapValidationError{
				field:  "TransportSocket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSocketTapConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TapValidationError{
					field:  "SocketTapConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TapValidationError{
					field:  "SocketTapConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSocketTapConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TapValidationError{
				field:  "SocketTapConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TapMultiError(errors)
	}

	return nil
}

// TapMultiError is an error wrapping multiple validation errors returned by
// Tap.ValidateAll() if the designated constraints aren't met.
type TapMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TapMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TapMultiError) AllErrors() []error { return m }

// TapValidationError is the validation error returned by Tap.Validate if the
// designated constraints aren't met.
type TapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TapValidationError) ErrorName() string { return "TapValidationError" }

// Error satisfies the builtin error interface
func (e TapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TapValidationError{}

// Validate checks the field values on SocketTapConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SocketTapConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SocketTapConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SocketTapConfigMultiError, or nil if none found.
func (m *SocketTapConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *SocketTapConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SetConnectionPerEvent

	// no validation rules for StatsPrefix

	if len(errors) > 0 {
		return SocketTapConfigMultiError(errors)
	}

	return nil
}

// SocketTapConfigMultiError is an error wrapping multiple validation errors
// returned by SocketTapConfig.ValidateAll() if the designated constraints
// aren't met.
type SocketTapConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SocketTapConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SocketTapConfigMultiError) AllErrors() []error { return m }

// SocketTapConfigValidationError is the validation error returned by
// SocketTapConfig.Validate if the designated constraints aren't met.
type SocketTapConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketTapConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketTapConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketTapConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketTapConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketTapConfigValidationError) ErrorName() string { return "SocketTapConfigValidationError" }

// Error satisfies the builtin error interface
func (e SocketTapConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketTapConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketTapConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketTapConfigValidationError{}
