#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail
shopt -s nullglob

root="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/.."
xds="${root}/data-plane-api"

imports=(
  ${xds}
  "${root}/vendor/github.com/envoyproxy/protoc-gen-validate"
  "${root}/vendor/github.com/gogo/protobuf"
  "${root}/vendor/github.com/gogo/googleapis"
  "${root}/vendor/istio.io/gogo-genproto/prometheus"
  "${root}/vendor/istio.io/gogo-genproto"
)

protoc="protoc"
protocarg=""
for i in "${imports[@]}"
do
  protocarg+="--proto_path=$i "
done

mappings=(
  "google/api/annotations.proto=github.com/gogo/googleapis/google/api"
  "google/api/http.proto=github.com/gogo/googleapis/google/api"
  "google/rpc/code.proto=github.com/gogo/googleapis/google/rpc"
  "google/rpc/error_details.proto=github.com/gogo/googleapis/google/rpc"
  "google/rpc/status.proto=github.com/gogo/googleapis/google/rpc"
  "google/protobuf/any.proto=github.com/gogo/protobuf/types"
  "google/protobuf/duration.proto=github.com/gogo/protobuf/types"
  "google/protobuf/empty.proto=github.com/gogo/protobuf/types"
  "google/protobuf/struct.proto=github.com/gogo/protobuf/types"
  "google/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
  "google/protobuf/wrappers.proto=github.com/gogo/protobuf/types"
  "gogoproto/gogo.proto=github.com/gogo/protobuf/gogoproto"
  "opencensus/proto/trace/v1/trace.proto=istio.io/gogo-genproto/opencensus/proto/trace/v1"
  "opencensus/proto/trace/v1/trace_config.proto=istio.io/gogo-genproto/opencensus/proto/trace/v1"
  "metrics.proto=istio.io/gogo-genproto/prometheus"
  "validate/validate.proto=github.com/envoyproxy/protoc-gen-validate/validate"
)

gogoarg="plugins=grpc"

# assign importmap for canonical protos
for mapping in "${mappings[@]}"
do
  gogoarg+=",M$mapping"
done

# assign importmap for all referenced protos in data-plane-api
for path in $(find ${xds}/envoy -type d)
do
  path_protos=(${path}/*.proto)
  if [[ ${#path_protos[@]} > 0 ]]
  then
    for path_proto in "${path_protos[@]}"
    do
      mapping=${path_proto##${xds}/}=github.com/envoyproxy/go-control-plane/${path##${xds}/}
      gogoarg+=",M$mapping"
    done
  fi
done

for path in $(find ${xds}/envoy -type d)
do
  path_protos=(${path}/*.proto)
  if [[ ${#path_protos[@]} > 0 ]]
  then
    echo "Generating protos ${path} ..."
    $protoc ${protocarg} ${path}/*.proto \
      --plugin=protoc-gen-gogofast=${root}/bin/gogofast --gogofast_out=${gogoarg}:. \
      --plugin=protoc-gen-validate=${root}/bin/validate --validate_out="lang=gogo:."
  fi
done
