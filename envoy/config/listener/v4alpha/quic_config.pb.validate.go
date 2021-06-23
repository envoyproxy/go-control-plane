// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/listener/v4alpha/quic_config.proto

package envoy_config_listener_v4alpha

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
)

// Validate checks the field values on QuicProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *QuicProtocolOptions) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetQuicProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "QuicProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "IdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCryptoHandshakeTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "CryptoHandshakeTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetPacketsToReadToConnectionCountRatio(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			return QuicProtocolOptionsValidationError{
				field:  "PacketsToReadToConnectionCountRatio",
				reason: "value must be greater than or equal to 1",
			}
		}

	}

	if v, ok := interface{}(m.GetCryptoStreamConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "CryptoStreamConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetProofSourceConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "ProofSourceConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// QuicProtocolOptionsValidationError is the validation error returned by
// QuicProtocolOptions.Validate if the designated constraints aren't met.
type QuicProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuicProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuicProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuicProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuicProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuicProtocolOptionsValidationError) ErrorName() string {
	return "QuicProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e QuicProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuicProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuicProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuicProtocolOptionsValidationError{}
