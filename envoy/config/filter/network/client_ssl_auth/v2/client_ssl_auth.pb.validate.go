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
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _client_ssl_auth_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ClientSSLAuth with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClientSSLAuth) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetAuthApiCluster()) < 1 {
		return ClientSSLAuthValidationError{
			field:  "AuthApiCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetStatPrefix()) < 1 {
		return ClientSSLAuthValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	if v, ok := interface{}(m.GetRefreshDelay()).(interface{ Validate() error }); ok {
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

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClientSSLAuthValidationError{
					field:  fmt.Sprintf("IpWhiteList[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

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
