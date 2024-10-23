// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/http/original_ip_detection/custom_header/v3/custom_header.proto

package envoy_extensions_http_original_ip_detection_custom_header_v3

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

// Validate checks the field values on CustomHeaderConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CustomHeaderConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CustomHeaderConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CustomHeaderConfigMultiError, or nil if none found.
func (m *CustomHeaderConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *CustomHeaderConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetHeaderName()) < 1 {
		err := CustomHeaderConfigValidationError{
			field:  "HeaderName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_CustomHeaderConfig_HeaderName_Pattern.MatchString(m.GetHeaderName()) {
		err := CustomHeaderConfigValidationError{
			field:  "HeaderName",
			reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for AllowExtensionToSetAddressAsTrusted

	if all {
		switch v := interface{}(m.GetRejectWithStatus()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CustomHeaderConfigValidationError{
					field:  "RejectWithStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CustomHeaderConfigValidationError{
					field:  "RejectWithStatus",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRejectWithStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CustomHeaderConfigValidationError{
				field:  "RejectWithStatus",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CustomHeaderConfigMultiError(errors)
	}
	return nil
}

// CustomHeaderConfigMultiError is an error wrapping multiple validation errors
// returned by CustomHeaderConfig.ValidateAll() if the designated constraints
// aren't met.
type CustomHeaderConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomHeaderConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomHeaderConfigMultiError) AllErrors() []error { return m }

// CustomHeaderConfigValidationError is the validation error returned by
// CustomHeaderConfig.Validate if the designated constraints aren't met.
type CustomHeaderConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomHeaderConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomHeaderConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomHeaderConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomHeaderConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomHeaderConfigValidationError) ErrorName() string {
	return "CustomHeaderConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CustomHeaderConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomHeaderConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomHeaderConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomHeaderConfigValidationError{}

var _CustomHeaderConfig_HeaderName_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")
