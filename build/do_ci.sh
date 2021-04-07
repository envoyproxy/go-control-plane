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

make build
make bin/example
make examples
make test
make integration