#!/bin/bash

set -o nounset
set -o errexit
set -o pipefail

DIRS=(  "pkg/cache"
        "pkg/server"
        "pkg/server/rest"
        "pkg/server/sotw"
        "pkg/test/resource"
        "pkg/test"
)
