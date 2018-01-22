#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

root="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

echo "Expecting protoc version >= 3.5.0:"
protoc=$(which protoc)
$protoc --version

echo "Expecting to find sibling data-plane-api repository ..."
pushd ../data-plane-api
  git log -1
popd

paths=(
  "api"
  "api/auth"
  "api/filter"
  "api/filter/accesslog"
  "api/filter/http"
  "api/filter/network"
)

imports=(
  "${root}/../data-plane-api"
  "${root}/vendor/github.com/lyft/protoc-gen-validate"
  "${root}/vendor/github.com/gogo/protobuf"
  "${root}/vendor/github.com/prometheus/client_model"
  "${root}/vendor/istio.io/gogo-genproto/googleapis"
)

protocarg=""
for i in "${imports[@]}"
do
  protocarg+="--proto_path=$i "
done

import="github.com/envoyproxy/go-control-plane/api"
mappings=(
  "google/api/annotations.proto=istio.io/gogo-genproto/googleapis/google/api"
  "google/api/http.proto=istio.io/gogo-genproto/googleapis/google/api"
  "google/rpc/code.proto=istio.io/gogo-genproto/googleapis/google/rpc"
  "google/rpc/error_details.proto=istio.io/gogo-genproto/googleapis/google/rpc"
  "google/rpc/status.proto=istio.io/gogo-genproto/googleapis/google/rpc"
  "google/protobuf/any.proto=github.com/gogo/protobuf/types"
  "google/protobuf/duration.proto=github.com/gogo/protobuf/types"
  "google/protobuf/struct.proto=github.com/gogo/protobuf/types"
  "google/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
  "google/protobuf/wrappers.proto=github.com/gogo/protobuf/types"
  "metrics.proto=github.com/prometheus/client_model/go"
  "gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto"
  "api/address.proto=${import}"
  "api/base.proto=${import}"
  "api/config_source.proto=${import}"
  "api/grpc_service.proto=${import}"
  "api/protocol.proto=${import}"
  "api/rds.proto=${import}"
  "api/sds.proto=${import}"
  "api/rls.proto=${import}"
  "api/filter/fault.proto=${import}/filter"
  "api/filter/accesslog/accesslog.proto=${import}/filter/accesslog"
)

gogoarg="plugins=grpc"

for mapping in "${mappings[@]}"
do
  gogoarg+=",M$mapping"
done

for path in "${paths[@]}"
do
  echo "Generating protos $path ..."
  $protoc ${protocarg} ${root}/../data-plane-api/${path}/*.proto \
    --plugin=protoc-gen-gogofast=${root}/bin/gogofast --gogofast_out=${gogoarg}:.
done

echo "Removing metrics_service.pb.go due to incompatibility with gogo (see https://github.com/prometheus/client_model/issues/15)"
\rm ${root}/api/metrics_service.pb.go
