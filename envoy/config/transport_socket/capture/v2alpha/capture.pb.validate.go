// Code generated by protoc-gen-validate
// source: envoy/config/transport_socket/capture/v2alpha/capture.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on FileSink with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *FileSink) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PathPrefix

	// no validation rules for Format

	return nil
}

// FileSinkValidationError is the validation error returned by
// FileSink.Validate if the designated constraints aren't met.
type FileSinkValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FileSinkValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileSink.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FileSinkValidationError{}

// Validate checks the field values on Capture with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Capture) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTransportSocket()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return CaptureValidationError{
				Field:  "TransportSocket",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	switch m.SinkSelector.(type) {

	case *Capture_FileSink:

		if v, ok := interface{}(m.GetFileSink()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return CaptureValidationError{
					Field:  "FileSink",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// CaptureValidationError is the validation error returned by Capture.Validate
// if the designated constraints aren't met.
type CaptureValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CaptureValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCapture.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CaptureValidationError{}
