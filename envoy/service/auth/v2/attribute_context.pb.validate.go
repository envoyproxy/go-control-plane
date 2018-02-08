// Code generated by protoc-gen-validate
// source: envoy/service/auth/v2/attribute_context.proto
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

// Validate checks the field values on AttributeContext with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AttributeContext) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSource()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				Field:  "Source",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDestination()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				Field:  "Destination",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRequest()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				Field:  "Request",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for ContextExtensions

	return nil
}

// AttributeContextValidationError is the validation error returned by
// AttributeContext.Validate if the designated constraints aren't met.
type AttributeContextValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AttributeContextValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AttributeContextValidationError{}

// Validate checks the field values on AttributeContext_Peer with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AttributeContext_Peer) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAddress()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_PeerValidationError{
				Field:  "Address",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for Service

	// no validation rules for Labels

	// no validation rules for Principal

	return nil
}

// AttributeContext_PeerValidationError is the validation error returned by
// AttributeContext_Peer.Validate if the designated constraints aren't met.
type AttributeContext_PeerValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AttributeContext_PeerValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Peer.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AttributeContext_PeerValidationError{}

// Validate checks the field values on AttributeContext_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AttributeContext_Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTime()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_RequestValidationError{
				Field:  "Time",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHttp()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_RequestValidationError{
				Field:  "Http",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// AttributeContext_RequestValidationError is the validation error returned by
// AttributeContext_Request.Validate if the designated constraints aren't met.
type AttributeContext_RequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AttributeContext_RequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Request.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AttributeContext_RequestValidationError{}

// Validate checks the field values on AttributeContext_HttpRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AttributeContext_HttpRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Method

	// no validation rules for Headers

	// no validation rules for Path

	// no validation rules for Host

	// no validation rules for Scheme

	// no validation rules for Query

	// no validation rules for Fragment

	// no validation rules for Size

	// no validation rules for Protocol

	return nil
}

// AttributeContext_HttpRequestValidationError is the validation error returned
// by AttributeContext_HttpRequest.Validate if the designated constraints
// aren't met.
type AttributeContext_HttpRequestValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AttributeContext_HttpRequestValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_HttpRequest.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AttributeContext_HttpRequestValidationError{}
