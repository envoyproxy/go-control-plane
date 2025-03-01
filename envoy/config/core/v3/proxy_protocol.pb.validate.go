//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/proxy_protocol.proto

package corev3

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

// Validate checks the field values on ProxyProtocolPassThroughTLVs with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProxyProtocolPassThroughTLVs) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProxyProtocolPassThroughTLVs with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProxyProtocolPassThroughTLVsMultiError, or nil if none found.
func (m *ProxyProtocolPassThroughTLVs) ValidateAll() error {
	return m.validate(true)
}

func (m *ProxyProtocolPassThroughTLVs) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MatchType

	for idx, item := range m.GetTlvType() {
		_, _ = idx, item

		if item >= 256 {
			err := ProxyProtocolPassThroughTLVsValidationError{
				field:  fmt.Sprintf("TlvType[%v]", idx),
				reason: "value must be less than 256",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return ProxyProtocolPassThroughTLVsMultiError(errors)
	}

	return nil
}

// ProxyProtocolPassThroughTLVsMultiError is an error wrapping multiple
// validation errors returned by ProxyProtocolPassThroughTLVs.ValidateAll() if
// the designated constraints aren't met.
type ProxyProtocolPassThroughTLVsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProxyProtocolPassThroughTLVsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProxyProtocolPassThroughTLVsMultiError) AllErrors() []error { return m }

// ProxyProtocolPassThroughTLVsValidationError is the validation error returned
// by ProxyProtocolPassThroughTLVs.Validate if the designated constraints
// aren't met.
type ProxyProtocolPassThroughTLVsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocolPassThroughTLVsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocolPassThroughTLVsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocolPassThroughTLVsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocolPassThroughTLVsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocolPassThroughTLVsValidationError) ErrorName() string {
	return "ProxyProtocolPassThroughTLVsValidationError"
}

// Error satisfies the builtin error interface
func (e ProxyProtocolPassThroughTLVsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocolPassThroughTLVs.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocolPassThroughTLVsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocolPassThroughTLVsValidationError{}

// Validate checks the field values on TlvEntry with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TlvEntry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TlvEntry with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TlvEntryMultiError, or nil
// if none found.
func (m *TlvEntry) ValidateAll() error {
	return m.validate(true)
}

func (m *TlvEntry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetType() >= 256 {
		err := TlvEntryValidationError{
			field:  "Type",
			reason: "value must be less than 256",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetValue()) < 1 {
		err := TlvEntryValidationError{
			field:  "Value",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TlvEntryMultiError(errors)
	}

	return nil
}

// TlvEntryMultiError is an error wrapping multiple validation errors returned
// by TlvEntry.ValidateAll() if the designated constraints aren't met.
type TlvEntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TlvEntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TlvEntryMultiError) AllErrors() []error { return m }

// TlvEntryValidationError is the validation error returned by
// TlvEntry.Validate if the designated constraints aren't met.
type TlvEntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TlvEntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TlvEntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TlvEntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TlvEntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TlvEntryValidationError) ErrorName() string { return "TlvEntryValidationError" }

// Error satisfies the builtin error interface
func (e TlvEntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTlvEntry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TlvEntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TlvEntryValidationError{}

// Validate checks the field values on ProxyProtocolConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProxyProtocolConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProxyProtocolConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProxyProtocolConfigMultiError, or nil if none found.
func (m *ProxyProtocolConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *ProxyProtocolConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	if all {
		switch v := interface{}(m.GetPassThroughTlvs()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProxyProtocolConfigValidationError{
					field:  "PassThroughTlvs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProxyProtocolConfigValidationError{
					field:  "PassThroughTlvs",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPassThroughTlvs()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProxyProtocolConfigValidationError{
				field:  "PassThroughTlvs",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetAddedTlvs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProxyProtocolConfigValidationError{
						field:  fmt.Sprintf("AddedTlvs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProxyProtocolConfigValidationError{
						field:  fmt.Sprintf("AddedTlvs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProxyProtocolConfigValidationError{
					field:  fmt.Sprintf("AddedTlvs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ProxyProtocolConfigMultiError(errors)
	}

	return nil
}

// ProxyProtocolConfigMultiError is an error wrapping multiple validation
// errors returned by ProxyProtocolConfig.ValidateAll() if the designated
// constraints aren't met.
type ProxyProtocolConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProxyProtocolConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProxyProtocolConfigMultiError) AllErrors() []error { return m }

// ProxyProtocolConfigValidationError is the validation error returned by
// ProxyProtocolConfig.Validate if the designated constraints aren't met.
type ProxyProtocolConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProxyProtocolConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProxyProtocolConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProxyProtocolConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProxyProtocolConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProxyProtocolConfigValidationError) ErrorName() string {
	return "ProxyProtocolConfigValidationError"
}

// Error satisfies the builtin error interface
func (e ProxyProtocolConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProxyProtocolConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProxyProtocolConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProxyProtocolConfigValidationError{}

// Validate checks the field values on PerHostConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PerHostConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PerHostConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PerHostConfigMultiError, or
// nil if none found.
func (m *PerHostConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *PerHostConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAddedTlvs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PerHostConfigValidationError{
						field:  fmt.Sprintf("AddedTlvs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PerHostConfigValidationError{
						field:  fmt.Sprintf("AddedTlvs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PerHostConfigValidationError{
					field:  fmt.Sprintf("AddedTlvs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PerHostConfigMultiError(errors)
	}

	return nil
}

// PerHostConfigMultiError is an error wrapping multiple validation errors
// returned by PerHostConfig.ValidateAll() if the designated constraints
// aren't met.
type PerHostConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PerHostConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PerHostConfigMultiError) AllErrors() []error { return m }

// PerHostConfigValidationError is the validation error returned by
// PerHostConfig.Validate if the designated constraints aren't met.
type PerHostConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PerHostConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PerHostConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PerHostConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PerHostConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PerHostConfigValidationError) ErrorName() string { return "PerHostConfigValidationError" }

// Error satisfies the builtin error interface
func (e PerHostConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerHostConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PerHostConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PerHostConfigValidationError{}
