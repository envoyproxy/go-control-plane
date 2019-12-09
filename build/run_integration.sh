#!/usr/bin/env bash

# execute the integration script in a container for controlled env
docker run -e XDS=$XDS -v $(pwd):/go-control-plane -w /go-control-plane envoyproxy/envoy:latest ./build/integration.sh
