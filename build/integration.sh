#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

##
## Integration testing script for the control plane library against the Envoy binary.
## This is a wrapper around the test app `pkg/test/main` that spawns/kills Envoy.
##

# Management server type. Valid values are "ads", "xds", "rest"
XDS=${XDS:-ads}

#Represents V3 api version
V3=${V3:-}

# Number of RTDS layers.
if [ "$XDS" = "ads" ]; then
  RUNTIMES=2
else
  RUNTIMES=1
fi

TEST_CLI_PARAM="--xds=${XDS} --runtimes=${RUNTIMES} -debug"
if [ -n "${V3}" ]; then
  TEST_CLI_PARAM="${TEST_CLI_PARAM} --resourceAPIVersion=${V3}"
fi


(bin/test ${TEST_CLI_PARAM} "$@")&
SERVER_PID=$!

# Envoy start-up command
ENVOY=${ENVOY:-/usr/local/bin/envoy}
ENVOY_LOG="envoy.${XDS}${V3}.log"
echo Envoy log: ${ENVOY_LOG}

# Start envoy: important to keep drain time short
(${ENVOY} -c sample/bootstrap-${XDS}${V3}.yaml --drain-time-s 1 -l debug 2> ${ENVOY_LOG})&
ENVOY_PID=$!

function cleanup() {
  kill ${ENVOY_PID}
}
trap cleanup EXIT

wait ${SERVER_PID}
