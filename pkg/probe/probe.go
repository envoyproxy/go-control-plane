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

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/prometheus/common/expfmt"
)

const (
	// Values described here:
	// https://www.envoyproxy.io/docs/envoy/v1.15.0/api-v3/admin/v3/server_info.proto#envoy-v3-api-enum-admin-v3-serverinfo-state
	// LIVE is value 0.
	StatServerState = "server.state"

	// Boolean indicating whether listeners have been initialized on the workers.
	StatListenerManagerWorkersStarted = "listener_manager.workers_started"

	// Total number of updates that resulted in a config reload.
	StatLDSConfigReload = "listener_manager.lds.config_reload"
	StatCDSConfigReload = "cluster_manager.cds.config_reload"
)

var (
	// General stats URI in prometheus format
	StatsURI = "/stats/prometheus"
	// Stats endpoint for reading readiness status (usedonly because some stats default to 0)
	ReadinessURI = fmt.Sprintf("%s?usedonly&filter=(%s|%s)",
		StatsURI,
		strings.ReplaceAll(StatServerState, ".", "\\."),
		strings.ReplaceAll(StatListenerManagerWorkersStarted, ".", "\\."))

	promRegex = regexp.MustCompile("[^a-zA-Z0-9_]")
)

// Converts a stat name to prometheus compatible name
func PromStatName(s string) string {
	return "envoy_" + promRegex.ReplaceAllString(s, "_")
}

// IsReady accesses the admin stats endpoint to check whether the server is initialized.
func IsReady(client *http.Client, host string) (bool, error) {
	url := fmt.Sprintf("http://%s%s", host, ReadinessURI)
	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	metrics, err := (&expfmt.TextParser{}).TextToMetricFamilies(resp.Body)
	if err != nil {
		return false, err
	}
	// server state must be 0
	if state, exists := metrics[PromStatName(StatServerState)]; !exists {
		return false, nil
	} else if len(state.Metric) != 1 || state.Metric[0].Gauge == nil ||
		state.Metric[0].Gauge.Value == nil || *state.Metric[0].Gauge.Value != 0 {
		return false, nil
	}
	// workers started must be 1
	if started, exists := metrics[PromStatName(StatListenerManagerWorkersStarted)]; !exists {
		return false, nil
	} else if len(started.Metric) != 1 || started.Metric[0].Gauge == nil ||
		started.Metric[0].Gauge.Value == nil || *started.Metric[0].Gauge.Value != 1 {
		return false, nil
	}
	return true, nil
}
