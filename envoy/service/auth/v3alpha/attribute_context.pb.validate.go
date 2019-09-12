// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/auth/v3alpha/attribute_context.proto

package envoy_service_auth_v3alpha

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
var _attribute_context_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AttributeContext with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AttributeContext) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDestination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Destination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ContextExtensions

	if v, ok := interface{}(m.GetMetadataContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContextValidationError{
				field:  "MetadataContext",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AttributeContextValidationError is the validation error returned by
// AttributeContext.Validate if the designated constraints aren't met.
type AttributeContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContextValidationError) ErrorName() string { return "AttributeContextValidationError" }

// Error satisfies the builtin error interface
func (e AttributeContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContextValidationError{}

// Validate checks the field values on AttributeContext_Peer with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AttributeContext_Peer) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_PeerValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
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
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_PeerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_PeerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_PeerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_PeerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_PeerValidationError) ErrorName() string {
	return "AttributeContext_PeerValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_PeerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Peer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_PeerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_PeerValidationError{}

// Validate checks the field values on AttributeContext_Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AttributeContext_Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_RequestValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHttp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AttributeContext_RequestValidationError{
				field:  "Http",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AttributeContext_RequestValidationError is the validation error returned by
// AttributeContext_Request.Validate if the designated constraints aren't met.
type AttributeContext_RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_RequestValidationError) ErrorName() string {
	return "AttributeContext_RequestValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_RequestValidationError{}

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

	// no validation rules for Body

	return nil
}

// AttributeContext_HttpRequestValidationError is the validation error returned
// by AttributeContext_HttpRequest.Validate if the designated constraints
// aren't met.
type AttributeContext_HttpRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AttributeContext_HttpRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AttributeContext_HttpRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AttributeContext_HttpRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AttributeContext_HttpRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AttributeContext_HttpRequestValidationError) ErrorName() string {
	return "AttributeContext_HttpRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AttributeContext_HttpRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAttributeContext_HttpRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AttributeContext_HttpRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AttributeContext_HttpRequestValidationError{}
