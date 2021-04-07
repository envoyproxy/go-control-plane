#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

##
## Runs Envoy and the example control plane server.  See
## `internal/example` for the go source.
##

# Envoy start-up command
ENVOY=${ENVOY:-/usr/local/bin/envoy}

# Start envoy: important to keep drain time short
(${ENVOY} -c sample/bootstrap-xds.yaml --drain-time-s 1 -l debug)&
ENVOY_PID=$!

function cleanup() {
  kill ${ENVOY_PID}
}
trap cleanup EXIT

# Run the control plane
bin/example -debug $@
