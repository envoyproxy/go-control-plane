# Benchmarks

These benchmarks are designed to run off the integration test suite, and provide detailed profiling artifacts generated with pprof.

## Prerequisites
- [Go 1.16+](https://golang.org/dl/)
- [Graphviz](https://graphviz.org/download/)
- [pprof](https://github.com/google/pprof)
- [Docker](https://www.docker.com/)

## Running Locally

To run the benchmarks locally, we take a similar apprach to the integration tests.

### Requirements

* Envoy binary `envoy` available: set `ENVOY` environment variable to the
  location of the binary, or use the default value `/usr/local/bin/envoy`
* `go-control-plane` builds successfully
* Local installation of pprof and graphviz for profiler output

### Steps

To run the benchmark:
```
make benchmark MODE=0
```

There are 5 different modes all corresponding to various profiling outputs:
```go
const (
	PPROF_CPU int = iota
	PPROF_HEAP
	PPROF_MUTEX
	PPROF_BLOCK
	PPROF_GOROUTINE
)
```
To specifiy the mode, use `MODE` environment variable that corresponds to the output you wish to evaluate.

The output file will be stored at the project root with .pprof extension (e.g. `cpu.pprof`).

Run `pprof` tool like so which will open up a shell. From there, you can run command such as `web` to visualize the result:

```bash
go tool pprof cpu.pprof 
(pprof) web
```
For more information, run `help`.

## Running With Docker

To run the benchmarks, we just require the prerequisite of docker and go.

### Steps

To run the benchmarks:

```
make docker_benchmarks
```

This will generate all profile artifacts in the `./benchmarks/reports`. Graphical profile anaylsis is located in `/.benchmarks/pngs`.

For more information on how to interpret these reports/graphs, [click here](https://github.com/google/pprof/blob/master/doc/README.md#graphical-reports)
