// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/mongo_proxy/v3/mongo_proxy.proto

package envoy_extensions_filters_network_mongo_proxy_v3

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
var _mongo_proxy_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on MongoProxy with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *MongoProxy) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		return MongoProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
	}

	// no validation rules for AccessLog

	if v, ok := interface{}(m.GetDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MongoProxyValidationError{
				field:  "Delay",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for EmitDynamicMetadata

	return nil
}

// MongoProxyValidationError is the validation error returned by
// MongoProxy.Validate if the designated constraints aren't met.
type MongoProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MongoProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MongoProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MongoProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MongoProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MongoProxyValidationError) ErrorName() string { return "MongoProxyValidationError" }

// Error satisfies the builtin error interface
func (e MongoProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMongoProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MongoProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MongoProxyValidationError{}
