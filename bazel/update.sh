#!/bin/bash
# Updates the API dependency SHA to master and uses bazel to generate code and copy it into the source tree

set -o errexit
set -o nounset
set -o pipefail
set -x

BAZEL="bazel"
ENVOY_API="envoy_api"

# Update the SHA on API repo
ENVOY_API_SHA=`git ls-remote https://github.com/envoyproxy/data-plane-api master | awk '{print $1}'`
if [[ "$OSTYPE" == "darwin"* ]]; then
  sed_params=(-i "")
else
  sed_params=(-i)
fi
sed "${sed_params[@]}" \
  -e "s|^ENVOY_API_SHA = .*$|ENVOY_API_SHA = \"$ENVOY_API_SHA\"|" \
  bazel/repositories.bzl

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
for package in $PACKAGES; do
  gofmt -w -r "\"$package\"->\"github.com/envoyproxy/go-control-plane/$package\"" api/
done



