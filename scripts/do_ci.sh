#!/usr/bin/env bash

set -e
set -x

# Needed to avoid issues with go version stamping in CI build
git config --global --add safe.directory /go-control-plane

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

make -C xdsmatcher test
# TODO(snowp): Output coverage in CI

make -C examples/dyplomat test
