// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/listener/v4alpha/udp_default_writer_config.proto

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
var _udp_default_writer_config_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on UdpDefaultWriterOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UdpDefaultWriterOptions) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UdpDefaultWriterOptionsValidationError is the validation error returned by
// UdpDefaultWriterOptions.Validate if the designated constraints aren't met.
type UdpDefaultWriterOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UdpDefaultWriterOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UdpDefaultWriterOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UdpDefaultWriterOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UdpDefaultWriterOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UdpDefaultWriterOptionsValidationError) ErrorName() string {
	return "UdpDefaultWriterOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e UdpDefaultWriterOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUdpDefaultWriterOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UdpDefaultWriterOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UdpDefaultWriterOptionsValidationError{}
