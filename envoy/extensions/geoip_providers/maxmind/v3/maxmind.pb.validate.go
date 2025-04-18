//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/geoip_providers/maxmind/v3/maxmind.proto

package maxmindv3

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

// Validate checks the field values on MaxMindConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MaxMindConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MaxMindConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MaxMindConfigMultiError, or
// nil if none found.
func (m *MaxMindConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *MaxMindConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_MaxMindConfig_CityDbPath_Pattern.MatchString(m.GetCityDbPath()) {
		err := MaxMindConfigValidationError{
			field:  "CityDbPath",
			reason: "value does not match regex pattern \"^$|^.*\\\\.mmdb$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_MaxMindConfig_AsnDbPath_Pattern.MatchString(m.GetAsnDbPath()) {
		err := MaxMindConfigValidationError{
			field:  "AsnDbPath",
			reason: "value does not match regex pattern \"^$|^.*\\\\.mmdb$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_MaxMindConfig_AnonDbPath_Pattern.MatchString(m.GetAnonDbPath()) {
		err := MaxMindConfigValidationError{
			field:  "AnonDbPath",
			reason: "value does not match regex pattern \"^$|^.*\\\\.mmdb$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_MaxMindConfig_IspDbPath_Pattern.MatchString(m.GetIspDbPath()) {
		err := MaxMindConfigValidationError{
			field:  "IspDbPath",
			reason: "value does not match regex pattern \"^$|^.*\\\\.mmdb$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCommonProviderConfig() == nil {
		err := MaxMindConfigValidationError{
			field:  "CommonProviderConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCommonProviderConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MaxMindConfigValidationError{
					field:  "CommonProviderConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MaxMindConfigValidationError{
					field:  "CommonProviderConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonProviderConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MaxMindConfigValidationError{
				field:  "CommonProviderConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MaxMindConfigMultiError(errors)
	}

	return nil
}

// MaxMindConfigMultiError is an error wrapping multiple validation errors
// returned by MaxMindConfig.ValidateAll() if the designated constraints
// aren't met.
type MaxMindConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MaxMindConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MaxMindConfigMultiError) AllErrors() []error { return m }

// MaxMindConfigValidationError is the validation error returned by
// MaxMindConfig.Validate if the designated constraints aren't met.
type MaxMindConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MaxMindConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MaxMindConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MaxMindConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MaxMindConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MaxMindConfigValidationError) ErrorName() string { return "MaxMindConfigValidationError" }

// Error satisfies the builtin error interface
func (e MaxMindConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMaxMindConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MaxMindConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MaxMindConfigValidationError{}

var _MaxMindConfig_CityDbPath_Pattern = regexp.MustCompile("^$|^.*\\.mmdb$")

var _MaxMindConfig_AsnDbPath_Pattern = regexp.MustCompile("^$|^.*\\.mmdb$")

var _MaxMindConfig_AnonDbPath_Pattern = regexp.MustCompile("^$|^.*\\.mmdb$")

var _MaxMindConfig_IspDbPath_Pattern = regexp.MustCompile("^$|^.*\\.mmdb$")
