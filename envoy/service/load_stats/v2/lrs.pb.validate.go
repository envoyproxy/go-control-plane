// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/load_stats/v2/lrs.proto

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

// Validate checks the field values on LoadStatsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LoadStatsRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoadStatsRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetClusterStats() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LoadStatsRequestValidationError{
					field:  fmt.Sprintf("ClusterStats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LoadStatsRequestValidationError is the validation error returned by
// LoadStatsRequest.Validate if the designated constraints aren't met.
type LoadStatsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoadStatsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoadStatsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoadStatsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoadStatsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoadStatsRequestValidationError) ErrorName() string { return "LoadStatsRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoadStatsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoadStatsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoadStatsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoadStatsRequestValidationError{}

// Validate checks the field values on LoadStatsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LoadStatsResponse) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetClusters()) < 1 {
		return LoadStatsResponseValidationError{
			field:  "Clusters",
			reason: "value must contain at least 1 item(s)",
		}
	}

	if v, ok := interface{}(m.GetLoadReportingInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoadStatsResponseValidationError{
				field:  "LoadReportingInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ReportEndpointGranularity

	return nil
}

// LoadStatsResponseValidationError is the validation error returned by
// LoadStatsResponse.Validate if the designated constraints aren't met.
type LoadStatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoadStatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoadStatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoadStatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoadStatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoadStatsResponseValidationError) ErrorName() string {
	return "LoadStatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LoadStatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoadStatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoadStatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoadStatsResponseValidationError{}
