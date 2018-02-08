// Code generated by protoc-gen-validate
// source: envoy/api/v2/ratelimit/ratelimit.proto
// DO NOT EDIT!!!

package ratelimit

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

// Validate checks the field values on RateLimitDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitDescriptor) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetEntries()) < 1 {
		return RateLimitDescriptorValidationError{
			Field:  "Entries",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return RateLimitDescriptorValidationError{
					Field:  fmt.Sprintf("Entries[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// RateLimitDescriptorValidationError is the validation error returned by
// RateLimitDescriptor.Validate if the designated constraints aren't met.
type RateLimitDescriptorValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptorValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RateLimitDescriptorValidationError{}

// Validate checks the field values on RateLimitDescriptor_Entry with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RateLimitDescriptor_Entry) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetKey()) < 1 {
		return RateLimitDescriptor_EntryValidationError{
			Field:  "Key",
			Reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetValue()) < 1 {
		return RateLimitDescriptor_EntryValidationError{
			Field:  "Value",
			Reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// RateLimitDescriptor_EntryValidationError is the validation error returned by
// RateLimitDescriptor_Entry.Validate if the designated constraints aren't met.
type RateLimitDescriptor_EntryValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptor_EntryValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor_Entry.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RateLimitDescriptor_EntryValidationError{}
