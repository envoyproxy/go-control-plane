#!/usr/bin/env bash

set -e
set -x

cd /go-control-plane

make create_version
make check_version_dirty
make build
make bin/example
make examples
make test
make integration