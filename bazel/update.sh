#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -x

BAZEL="bazel"
ENVOY_API="envoy_api"

# Generate go sources
GO_LIBRARIES=`$BAZEL query "attr(name, go_default_library, @${ENVOY_API}//...)" | tr '\n' ' '`
BAZEL_BIN=`$BAZEL info bazel-bin`
$BAZEL build $GO_LIBRARIES

PACKAGES=`echo $GO_LIBRARIES | sed "s/[^ ]*\/\/\([^ ]*\):go_default_library/\1/g"`

# Move generated API files to the sources
for package in $PACKAGES; do
  mkdir -p "$package"
  rsync -a --chmod=775 --include='*.pb.go' --exclude='*' \
    "$BAZEL_BIN/external/$ENVOY_API/$package/go_default_library/$package/" "$package/"
done

# Rewrite the local imports
gofmt -w -r '"api"->"github.com/envoyproxy/go-control-plane/api"' api/



