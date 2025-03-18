//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/network_ext_proc/v3/network_external_processor.proto

package network_ext_procv3

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

// Validate checks the field values on Data with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Data with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DataMultiError, or nil if none found.
func (m *Data) ValidateAll() error {
	return m.validate(true)
}

func (m *Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Data

	// no validation rules for EndOfStream

	if len(errors) > 0 {
		return DataMultiError(errors)
	}

	return nil
}

// DataMultiError is an error wrapping multiple validation errors returned by
// Data.ValidateAll() if the designated constraints aren't met.
type DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DataMultiError) AllErrors() []error { return m }

// DataValidationError is the validation error returned by Data.Validate if the
// designated constraints aren't met.
type DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataValidationError) ErrorName() string { return "DataValidationError" }

// Error satisfies the builtin error interface
func (e DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataValidationError{}

// Validate checks the field values on ProcessingRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProcessingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProcessingRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProcessingRequestMultiError, or nil if none found.
func (m *ProcessingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ProcessingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetReadData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProcessingRequestValidationError{
					field:  "ReadData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProcessingRequestValidationError{
					field:  "ReadData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReadData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingRequestValidationError{
				field:  "ReadData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWriteData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProcessingRequestValidationError{
					field:  "WriteData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProcessingRequestValidationError{
					field:  "WriteData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWriteData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingRequestValidationError{
				field:  "WriteData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProcessingRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProcessingRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProcessingRequestMultiError(errors)
	}

	return nil
}

// ProcessingRequestMultiError is an error wrapping multiple validation errors
// returned by ProcessingRequest.ValidateAll() if the designated constraints
// aren't met.
type ProcessingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProcessingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProcessingRequestMultiError) AllErrors() []error { return m }

// ProcessingRequestValidationError is the validation error returned by
// ProcessingRequest.Validate if the designated constraints aren't met.
type ProcessingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessingRequestValidationError) ErrorName() string {
	return "ProcessingRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProcessingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessingRequestValidationError{}

// Validate checks the field values on ProcessingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProcessingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProcessingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProcessingResponseMultiError, or nil if none found.
func (m *ProcessingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProcessingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetReadData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProcessingResponseValidationError{
					field:  "ReadData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProcessingResponseValidationError{
					field:  "ReadData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReadData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingResponseValidationError{
				field:  "ReadData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWriteData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProcessingResponseValidationError{
					field:  "WriteData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProcessingResponseValidationError{
					field:  "WriteData",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWriteData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingResponseValidationError{
				field:  "WriteData",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DataProcessingStatus

	// no validation rules for ConnectionStatus

	if all {
		switch v := interface{}(m.GetDynamicMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProcessingResponseValidationError{
					field:  "DynamicMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProcessingResponseValidationError{
					field:  "DynamicMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDynamicMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessingResponseValidationError{
				field:  "DynamicMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProcessingResponseMultiError(errors)
	}

	return nil
}

// ProcessingResponseMultiError is an error wrapping multiple validation errors
// returned by ProcessingResponse.ValidateAll() if the designated constraints
// aren't met.
type ProcessingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProcessingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProcessingResponseMultiError) AllErrors() []error { return m }

// ProcessingResponseValidationError is the validation error returned by
// ProcessingResponse.Validate if the designated constraints aren't met.
type ProcessingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessingResponseValidationError) ErrorName() string {
	return "ProcessingResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProcessingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessingResponseValidationError{}
