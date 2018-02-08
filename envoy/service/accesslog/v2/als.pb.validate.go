// Code generated by protoc-gen-validate
// source: envoy/service/accesslog/v2/als.proto
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

// Validate checks the field values on StreamAccessLogsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamAccessLogsResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamAccessLogsResponseValidationError is the validation error returned by
// StreamAccessLogsResponse.Validate if the designated constraints aren't met.
type StreamAccessLogsResponseValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsResponseValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsResponse.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamAccessLogsResponseValidationError{}

// Validate checks the field values on StreamAccessLogsMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamAccessLogsMessage) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetIdentifier()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return StreamAccessLogsMessageValidationError{
				Field:  "Identifier",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	switch m.LogEntries.(type) {

	case *StreamAccessLogsMessage_HttpLogs:

		if v, ok := interface{}(m.GetHttpLogs()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessageValidationError{
					Field:  "HttpLogs",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *StreamAccessLogsMessage_TcpLogs:

		if v, ok := interface{}(m.GetTcpLogs()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessageValidationError{
					Field:  "TcpLogs",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return StreamAccessLogsMessageValidationError{
			Field:  "LogEntries",
			Reason: "value is required",
		}

	}

	return nil
}

// StreamAccessLogsMessageValidationError is the validation error returned by
// StreamAccessLogsMessage.Validate if the designated constraints aren't met.
type StreamAccessLogsMessageValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessageValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamAccessLogsMessageValidationError{}

// Validate checks the field values on StreamAccessLogsMessage_Identifier with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *StreamAccessLogsMessage_Identifier) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetNode() == nil {
		return StreamAccessLogsMessage_IdentifierValidationError{
			Field:  "Node",
			Reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetNode()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return StreamAccessLogsMessage_IdentifierValidationError{
				Field:  "Node",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if len(m.GetLogName()) < 1 {
		return StreamAccessLogsMessage_IdentifierValidationError{
			Field:  "LogName",
			Reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// StreamAccessLogsMessage_IdentifierValidationError is the validation error
// returned by StreamAccessLogsMessage_Identifier.Validate if the designated
// constraints aren't met.
type StreamAccessLogsMessage_IdentifierValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessage_IdentifierValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage_Identifier.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamAccessLogsMessage_IdentifierValidationError{}

// Validate checks the field values on
// StreamAccessLogsMessage_HTTPAccessLogEntries with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *StreamAccessLogsMessage_HTTPAccessLogEntries) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetLogEntry()) < 1 {
		return StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError{
			Field:  "LogEntry",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetLogEntry() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError{
					Field:  fmt.Sprintf("LogEntry[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError is the
// validation error returned by
// StreamAccessLogsMessage_HTTPAccessLogEntries.Validate if the designated
// constraints aren't met.
type StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage_HTTPAccessLogEntries.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamAccessLogsMessage_HTTPAccessLogEntriesValidationError{}

// Validate checks the field values on
// StreamAccessLogsMessage_TCPAccessLogEntries with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *StreamAccessLogsMessage_TCPAccessLogEntries) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetLogEntry()) < 1 {
		return StreamAccessLogsMessage_TCPAccessLogEntriesValidationError{
			Field:  "LogEntry",
			Reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetLogEntry() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return StreamAccessLogsMessage_TCPAccessLogEntriesValidationError{
					Field:  fmt.Sprintf("LogEntry[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamAccessLogsMessage_TCPAccessLogEntriesValidationError is the validation
// error returned by StreamAccessLogsMessage_TCPAccessLogEntries.Validate if
// the designated constraints aren't met.
type StreamAccessLogsMessage_TCPAccessLogEntriesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e StreamAccessLogsMessage_TCPAccessLogEntriesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamAccessLogsMessage_TCPAccessLogEntries.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = StreamAccessLogsMessage_TCPAccessLogEntriesValidationError{}
