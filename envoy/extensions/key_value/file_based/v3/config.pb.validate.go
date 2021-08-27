// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/key_value/file_based/v3/config.proto

package envoy_extensions_key_value_file_based_v3

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

// Validate checks the field values on FileBasedKeyValueStoreConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FileBasedKeyValueStoreConfig) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetFilename()) < 1 {
		return FileBasedKeyValueStoreConfigValidationError{
			field:  "Filename",
			reason: "value length must be at least 1 runes",
		}
	}

	if v, ok := interface{}(m.GetFlushInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FileBasedKeyValueStoreConfigValidationError{
				field:  "FlushInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// FileBasedKeyValueStoreConfigValidationError is the validation error returned
// by FileBasedKeyValueStoreConfig.Validate if the designated constraints
// aren't met.
type FileBasedKeyValueStoreConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FileBasedKeyValueStoreConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FileBasedKeyValueStoreConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FileBasedKeyValueStoreConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FileBasedKeyValueStoreConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FileBasedKeyValueStoreConfigValidationError) ErrorName() string {
	return "FileBasedKeyValueStoreConfigValidationError"
}

// Error satisfies the builtin error interface
func (e FileBasedKeyValueStoreConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileBasedKeyValueStoreConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FileBasedKeyValueStoreConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FileBasedKeyValueStoreConfigValidationError{}
