// Copyright 2020 Envoyproxy Authors
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

// Package probe provides common queries for envoy stats system.
package probe

const (
	// Values described here:
	// https://www.envoyproxy.io/docs/envoy/v1.15.0/api-v3/admin/v3/server_info.proto#envoy-v3-api-enum-admin-v3-serverinfo-state
	StatServerState = "server.state"

	// Boolean indicating whether listeners have been initialized on the workers.
	StatListenerManagerWorkersStarted = "listener_manager.workers_started"

	// Total number of updates that resulted in a config reload.
	StatLDSConfigReload = "listener_manager.lds.config_reload"
	StatCDSConfigReload = "cluster_manager.cds.config_reload"
)
