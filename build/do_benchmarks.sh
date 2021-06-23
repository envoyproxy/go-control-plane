#!/usr/bin/env bash

set -e
set -x

go install golang.org/x/tools/cmd/goimports@latest

cd /go-control-plane

NUM_DIFF_LINES=`/go/bin/goimports -local github.com/envoyproxy/go-control-plane -d pkg | wc -l`
if [[ $NUM_DIFF_LINES > 0 ]]; then
  echo "Failed format check. Run make format"
  exit 1
fi

make artifacts

# Run our two benchmarks as integration tests
# While these are running we will need to scrape the pprof server and pull the metrics out of that
make integration.xds
make integration.xds.delta