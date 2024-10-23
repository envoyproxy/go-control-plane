// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/network/client_ssl_auth/v2/client_ssl_auth.proto

package envoy_config_filter_network_client_ssl_auth_v2

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

// Validate checks the field values on ClientSSLAuth with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClientSSLAuth) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClientSSLAuth with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClientSSLAuthMultiError, or
// nil if none found.
func (m *ClientSSLAuth) ValidateAll() error {
	return m.validate(true)
}

func (m *ClientSSLAuth) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetAuthApiCluster()) < 1 {
		err := ClientSSLAuthValidationError{
			field:  "AuthApiCluster",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetStatPrefix()) < 1 {
		err := ClientSSLAuthValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRefreshDelay()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ClientSSLAuthValidationError{
					field:  "RefreshDelay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ClientSSLAuthValidationError{
					field:  "RefreshDelay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRefreshDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClientSSLAuthValidationError{
				field:  "RefreshDelay",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetIpWhiteList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClientSSLAuthValidationError{
						field:  fmt.Sprintf("IpWhiteList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClientSSLAuthValidationError{
						field:  fmt.Sprintf("IpWhiteList[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientSSLAuthValidationError{
					field:  fmt.Sprintf("IpWhiteList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ClientSSLAuthMultiError(errors)
	}
	return nil
}

// ClientSSLAuthMultiError is an error wrapping multiple validation errors
// returned by ClientSSLAuth.ValidateAll() if the designated constraints
// aren't met.
type ClientSSLAuthMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClientSSLAuthMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClientSSLAuthMultiError) AllErrors() []error { return m }

// ClientSSLAuthValidationError is the validation error returned by
// ClientSSLAuth.Validate if the designated constraints aren't met.
type ClientSSLAuthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClientSSLAuthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClientSSLAuthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClientSSLAuthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClientSSLAuthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClientSSLAuthValidationError) ErrorName() string { return "ClientSSLAuthValidationError" }

// Error satisfies the builtin error interface
func (e ClientSSLAuthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClientSSLAuth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClientSSLAuthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClientSSLAuthValidationError{}
