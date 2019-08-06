// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v2/endpoint/load_report.proto

package endpoint

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

// Validate checks the field values on UpstreamLocalityStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpstreamLocalityStats) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetLocality()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return UpstreamLocalityStatsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for TotalSuccessfulRequests

	// no validation rules for TotalRequestsInProgress

	// no validation rules for TotalErrorRequests

	// no validation rules for TotalIssuedRequests

	for idx, item := range m.GetLoadMetricStats() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return UpstreamLocalityStatsValidationError{
						field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetUpstreamEndpointStats() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return UpstreamLocalityStatsValidationError{
						field:  fmt.Sprintf("UpstreamEndpointStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	// no validation rules for Priority

	return nil
}

// UpstreamLocalityStatsValidationError is the validation error returned by
// UpstreamLocalityStats.Validate if the designated constraints aren't met.
type UpstreamLocalityStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamLocalityStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamLocalityStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamLocalityStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamLocalityStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamLocalityStatsValidationError) ErrorName() string {
	return "UpstreamLocalityStatsValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamLocalityStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamLocalityStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamLocalityStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamLocalityStatsValidationError{}

// Validate checks the field values on UpstreamEndpointStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpstreamEndpointStats) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetAddress()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return UpstreamEndpointStatsValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMetadata()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return UpstreamEndpointStatsValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for TotalSuccessfulRequests

	// no validation rules for TotalRequestsInProgress

	// no validation rules for TotalErrorRequests

	// no validation rules for TotalIssuedRequests

	for idx, item := range m.GetLoadMetricStats() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return UpstreamEndpointStatsValidationError{
						field:  fmt.Sprintf("LoadMetricStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// UpstreamEndpointStatsValidationError is the validation error returned by
// UpstreamEndpointStats.Validate if the designated constraints aren't met.
type UpstreamEndpointStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpstreamEndpointStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpstreamEndpointStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpstreamEndpointStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpstreamEndpointStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpstreamEndpointStatsValidationError) ErrorName() string {
	return "UpstreamEndpointStatsValidationError"
}

// Error satisfies the builtin error interface
func (e UpstreamEndpointStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpstreamEndpointStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpstreamEndpointStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpstreamEndpointStatsValidationError{}

// Validate checks the field values on EndpointLoadMetricStats with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *EndpointLoadMetricStats) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MetricName

	// no validation rules for NumRequestsFinishedWithMetric

	// no validation rules for TotalMetricValue

	return nil
}

// EndpointLoadMetricStatsValidationError is the validation error returned by
// EndpointLoadMetricStats.Validate if the designated constraints aren't met.
type EndpointLoadMetricStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointLoadMetricStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointLoadMetricStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointLoadMetricStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointLoadMetricStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointLoadMetricStatsValidationError) ErrorName() string {
	return "EndpointLoadMetricStatsValidationError"
}

// Error satisfies the builtin error interface
func (e EndpointLoadMetricStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointLoadMetricStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointLoadMetricStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointLoadMetricStatsValidationError{}

// Validate checks the field values on ClusterStats with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClusterStats) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetClusterName()) < 1 {
		return ClusterStatsValidationError{
			field:  "ClusterName",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for ClusterServiceName

	if len(m.GetUpstreamLocalityStats()) < 1 {
		return ClusterStatsValidationError{
			field:  "UpstreamLocalityStats",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetUpstreamLocalityStats() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ClusterStatsValidationError{
						field:  fmt.Sprintf("UpstreamLocalityStats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	// no validation rules for TotalDroppedRequests

	for idx, item := range m.GetDroppedRequests() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ClusterStatsValidationError{
						field:  fmt.Sprintf("DroppedRequests[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetLoadReportInterval()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClusterStatsValidationError{
					field:  "LoadReportInterval",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ClusterStatsValidationError is the validation error returned by
// ClusterStats.Validate if the designated constraints aren't met.
type ClusterStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterStatsValidationError) ErrorName() string { return "ClusterStatsValidationError" }

// Error satisfies the builtin error interface
func (e ClusterStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterStatsValidationError{}

// Validate checks the field values on ClusterStats_DroppedRequests with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterStats_DroppedRequests) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCategory()) < 1 {
		return ClusterStats_DroppedRequestsValidationError{
			field:  "Category",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for DroppedCount

	return nil
}

// ClusterStats_DroppedRequestsValidationError is the validation error returned
// by ClusterStats_DroppedRequests.Validate if the designated constraints
// aren't met.
type ClusterStats_DroppedRequestsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterStats_DroppedRequestsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterStats_DroppedRequestsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterStats_DroppedRequestsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterStats_DroppedRequestsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterStats_DroppedRequestsValidationError) ErrorName() string {
	return "ClusterStats_DroppedRequestsValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterStats_DroppedRequestsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterStats_DroppedRequests.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterStats_DroppedRequestsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterStats_DroppedRequestsValidationError{}
