// Code generated by protoc-gen-validate
// source: envoy/config/accesslog/v2/file.proto
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

// Validate checks the field values on FileAccessLog with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FileAccessLog) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetPath()) < 1 {
		return FileAccessLogValidationError{
			Field:  "Path",
			Reason: "value length must be at least 1 bytes",
		}
	}

	switch m.AccessLogFormat.(type) {

	case *FileAccessLog_Format:
		// no validation rules for Format

	case *FileAccessLog_JsonFormat:

		if v, ok := interface{}(m.GetJsonFormat()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return FileAccessLogValidationError{
					Field:  "JsonFormat",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// FileAccessLogValidationError is the validation error returned by
// FileAccessLog.Validate if the designated constraints aren't met.
type FileAccessLogValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e FileAccessLogValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFileAccessLog.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = FileAccessLogValidationError{}
