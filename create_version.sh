#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

MODULES=( '"github.com/envoyproxy/go-control-plane/envoy/api/v2":"github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"'
)

DIRS=( "pkg/cache"
)

for dir in "${DIRS[@]}" ; do
    v2dir="$dir/v2"
    v3dir="$dir/v3"
    mkdir -p $v3dir
    cp -R "$v2dir/" "$v3dir"
    FILES=($(ls -p "$v3dir"))
    for file in "${FILES[@]}" ; do
        path="$v3dir/$file"
        for module in "${MODULES[@]}" ; do
            KEY=${module%%:*}
            VALUE=${module#*:}
            sed -i "s|$KEY|$VALUE|" $path
        done
    done
done