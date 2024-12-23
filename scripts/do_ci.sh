#!/usr/bin/env bash

set -e
set -x

# Needed to avoid issues with go version stamping in CI build
git config --global --add safe.directory /go-control-plane

cd /go-control-plane

make build
make bin/example
make examples
make test
make integration

make -C xdsmatcher common/test
# TODO(snowp): Output coverage in CI

make -C examples/dyplomat common/test
