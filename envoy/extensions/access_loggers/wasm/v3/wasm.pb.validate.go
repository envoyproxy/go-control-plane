// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/access_loggers/wasm/v3/wasm.proto

package envoy_extensions_access_loggers_wasm_v3

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

// Validate checks the field values on WasmAccessLog with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *WasmAccessLog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WasmAccessLog with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in WasmAccessLogMultiError, or
// nil if none found.
func (m *WasmAccessLog) ValidateAll() error {
	return m.validate(true)
}

func (m *WasmAccessLog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WasmAccessLogValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WasmAccessLogValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WasmAccessLogValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return WasmAccessLogMultiError(errors)
	}
	return nil
}

// WasmAccessLogMultiError is an error wrapping multiple validation errors
// returned by WasmAccessLog.ValidateAll() if the designated constraints
// aren't met.
type WasmAccessLogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WasmAccessLogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WasmAccessLogMultiError) AllErrors() []error { return m }

// WasmAccessLogValidationError is the validation error returned by
// WasmAccessLog.Validate if the designated constraints aren't met.
type WasmAccessLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WasmAccessLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WasmAccessLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WasmAccessLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WasmAccessLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WasmAccessLogValidationError) ErrorName() string { return "WasmAccessLogValidationError" }

// Error satisfies the builtin error interface
func (e WasmAccessLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWasmAccessLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WasmAccessLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WasmAccessLogValidationError{}
