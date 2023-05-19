#!/usr/bin/env bash

set -e
set -x

go install golang.org/x/tools/cmd/goimports@latest
go get -u github.com/google/pprof
apt update && apt install -y graphviz

cd /go-control-plane

NUM_DIFF_LINES=`/go/bin/goimports -local github.com/envoyproxy/go-control-plane -d pkg | wc -l`
if [[ $NUM_DIFF_LINES > 0 ]]; then
  echo "Failed format check. Run make format"
  exit 1
fi

make benchmark MODE=0
make benchmark MODE=1
make benchmark MODE=2
make benchmark MODE=3
make benchmark MODE=4

mkdir -p benchmarks/reports
mkdir -p benchmarks/pngs

# generate our consumable pprof profiles
pprof -text bin/test cpu.pprof > benchmarks/reports/cpu_text.txt
pprof -tree bin/test cpu.pprof > benchmarks/reports/cpu_tree.txt
pprof -traces bin/test cpu.pprof > benchmarks/reports/cpu_trace.txt

pprof -text bin/test goroutine.pprof > benchmarks/reports/goroutine_text.txt
pprof -tree bin/test goroutine.pprof > benchmarks/reports/goroutine_tree.txt
pprof -traces bin/test goroutine.pprof > benchmarks/reports/goroutine_trace.txt

pprof -text bin/test block.pprof > benchmarks/reports/block_text.txt
pprof -tree bin/test block.pprof > benchmarks/reports/block_tree.txt
pprof -traces bin/test block.pprof > benchmarks/reports/block_trace.txt

pprof -text bin/test mutex.pprof > benchmarks/reports/mutex_text.txt
pprof -tree bin/test mutex.pprof > benchmarks/reports/mutex_tree.txt
pprof -traces bin/test mutex.pprof > benchmarks/reports/mutex_trace.txt

pprof -text bin/test mem.pprof > benchmarks/reports/mem_text.txt
pprof -tree bin/test mem.pprof > benchmarks/reports/mem_tree.txt
pprof -traces bin/test mem.pprof > benchmarks/reports/mem_trace.txt

# generate some cool pprof graphs
pprof -png bin/test cpu.pprof > benchmarks/pngs/cpu_graph.png
pprof -png bin/test goroutine.pprof > benchmarks/pngs/goroutine_graph.png
pprof -png bin/test block.pprof > benchmarks/pngs/block_graph.png
pprof -png bin/test mutex.pprof > benchmarks/pngs/mutex_graph.png
pprof -png bin/test mem.pprof > benchmarks/pngs/mem_graph.png
