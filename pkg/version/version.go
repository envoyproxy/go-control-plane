// Copyright 2016 IBM Corporation
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

// Package version provides build time version information.
package version

import (
	"time"
)

// Build variable exposes the build-time information
var Build BuildInfo

// BuildInfo provides build information and metadata.
// It is Populated by Makefile at build-time using -ldflags -X.
// Symbol names are reported by 'go tool nm <binary object>".
// Note that unless the cluster of instances is completely homogeneous, different instances can return
// different values.
type BuildInfo struct {
	Version     string    `json:"version"`
	GitRevision string    `json:"revision"`
	Branch      string    `json:"branch"`
	BuildUser   string    `json:"user"`
	BuildDate   time.Time `json:"date"`
	GoVersion   string    `json:"go"`
}

var (
	version     string
	gitRevision string
	branch      string
	buildUser   string
	buildDate   string
	goVersion   string
)

// Unfortunately, it seems that struct fields cannot be assigned directly using ldflags, so we resort to setting
// package scope (local) variables and then assigning their corresponding external values at initialization
func init() {
	if version == "" { // this is most likely running in a test...
		Build.Version = "unknown"
		Build.GitRevision = "unknown"
		Build.Branch = "unknown"
		Build.BuildUser = "unknown"
		Build.BuildDate = time.Time{}
		Build.GoVersion = "unknown"
	} else {
		Build.Version = version
		Build.GitRevision = gitRevision
		Build.Branch = branch
		Build.BuildUser = buildUser
		Build.BuildDate, _ = time.Parse(time.RFC3339, buildDate)
		Build.GoVersion = goVersion
	}
}
