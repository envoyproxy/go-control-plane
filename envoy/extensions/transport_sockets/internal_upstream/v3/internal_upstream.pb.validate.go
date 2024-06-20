//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/transport_sockets/internal_upstream/v3/internal_upstream.proto

package internal_upstreamv3

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

// Validate checks the field values on InternalUpstreamTransport with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InternalUpstreamTransport) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InternalUpstreamTransport with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InternalUpstreamTransportMultiError, or nil if none found.
func (m *InternalUpstreamTransport) ValidateAll() error {
	return m.validate(true)
}

func (m *InternalUpstreamTransport) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPassthroughMetadata() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InternalUpstreamTransportValidationError{
						field:  fmt.Sprintf("PassthroughMetadata[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InternalUpstreamTransportValidationError{
						field:  fmt.Sprintf("PassthroughMetadata[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InternalUpstreamTransportValidationError{
					field:  fmt.Sprintf("PassthroughMetadata[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetTransportSocket() == nil {
		err := InternalUpstreamTransportValidationError{
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
				errors = append(errors, InternalUpstreamTransportValidationError{
					field:  "TransportSocket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalUpstreamTransportValidationError{
					field:  "TransportSocket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTransportSocket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalUpstreamTransportValidationError{
				field:  "TransportSocket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return InternalUpstreamTransportMultiError(errors)
	}

	return nil
}

// InternalUpstreamTransportMultiError is an error wrapping multiple validation
// errors returned by InternalUpstreamTransport.ValidateAll() if the
// designated constraints aren't met.
type InternalUpstreamTransportMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InternalUpstreamTransportMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InternalUpstreamTransportMultiError) AllErrors() []error { return m }

// InternalUpstreamTransportValidationError is the validation error returned by
// InternalUpstreamTransport.Validate if the designated constraints aren't met.
type InternalUpstreamTransportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InternalUpstreamTransportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InternalUpstreamTransportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InternalUpstreamTransportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InternalUpstreamTransportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InternalUpstreamTransportValidationError) ErrorName() string {
	return "InternalUpstreamTransportValidationError"
}

// Error satisfies the builtin error interface
func (e InternalUpstreamTransportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInternalUpstreamTransport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InternalUpstreamTransportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InternalUpstreamTransportValidationError{}

// Validate checks the field values on
// InternalUpstreamTransport_MetadataValueSource with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InternalUpstreamTransport_MetadataValueSource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// InternalUpstreamTransport_MetadataValueSource with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// InternalUpstreamTransport_MetadataValueSourceMultiError, or nil if none found.
func (m *InternalUpstreamTransport_MetadataValueSource) ValidateAll() error {
	return m.validate(true)
}

func (m *InternalUpstreamTransport_MetadataValueSource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetKind() == nil {
		err := InternalUpstreamTransport_MetadataValueSourceValidationError{
			field:  "Kind",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetKind()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalUpstreamTransport_MetadataValueSourceValidationError{
					field:  "Kind",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalUpstreamTransport_MetadataValueSourceValidationError{
					field:  "Kind",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKind()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalUpstreamTransport_MetadataValueSourceValidationError{
				field:  "Kind",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := InternalUpstreamTransport_MetadataValueSourceValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return InternalUpstreamTransport_MetadataValueSourceMultiError(errors)
	}

	return nil
}

// InternalUpstreamTransport_MetadataValueSourceMultiError is an error wrapping
// multiple validation errors returned by
// InternalUpstreamTransport_MetadataValueSource.ValidateAll() if the
// designated constraints aren't met.
type InternalUpstreamTransport_MetadataValueSourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InternalUpstreamTransport_MetadataValueSourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InternalUpstreamTransport_MetadataValueSourceMultiError) AllErrors() []error { return m }

// InternalUpstreamTransport_MetadataValueSourceValidationError is the
// validation error returned by
// InternalUpstreamTransport_MetadataValueSource.Validate if the designated
// constraints aren't met.
type InternalUpstreamTransport_MetadataValueSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InternalUpstreamTransport_MetadataValueSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InternalUpstreamTransport_MetadataValueSourceValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e InternalUpstreamTransport_MetadataValueSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InternalUpstreamTransport_MetadataValueSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InternalUpstreamTransport_MetadataValueSourceValidationError) ErrorName() string {
	return "InternalUpstreamTransport_MetadataValueSourceValidationError"
}

// Error satisfies the builtin error interface
func (e InternalUpstreamTransport_MetadataValueSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInternalUpstreamTransport_MetadataValueSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InternalUpstreamTransport_MetadataValueSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InternalUpstreamTransport_MetadataValueSourceValidationError{}
