#!/usr/bin/env bash

set -e

NUM_DIFF_LINES=`${GOIMPORTS} -local ${PKG} -d pkg | wc -l`
if [[ $NUM_DIFF_LINES > 0 ]]; then
  echo "Failed format check. Run make format"
  exit 1
fi