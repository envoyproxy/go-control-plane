// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package log provides a logging interface for use in this library.
package log

import "strings"

// Level is an exposed logger level type that users can select
type Level string

const (
	// DEBUG Log Level
	DEBUG Level = "debug"
	// INFO Log Level (little less log pollution from xDS)
	INFO Level = "info"
	// WARN Log Level (show warnigs and potential hazards)
	WARN Level = "warn"
	// ERROR Log Level (only log failures)
	ERROR Level = "error"
)

// ParseLevel will accept a string and return a logging level
func ParseLevel(l string) Level {
	switch strings.TrimSpace(strings.ToLower(l)) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	default:
		return ERROR
	}
}

// Logger interface for reporting informational and warning messages.
type Logger interface {
	// Debugf logs a formatted debugging message.
	Debugf(format string, args ...interface{})

	// Infof logs a formatted informational message.
	Infof(format string, args ...interface{})

	// Warnf logs a formatted warning message.
	Warnf(format string, args ...interface{})

	// Errorf logs a formatted error message.
	Errorf(format string, args ...interface{})
}
