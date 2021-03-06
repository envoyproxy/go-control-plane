# End-to-end testing script

This directory contains a basic end-to-end testing script.
The script sets up a configuration cache, stands up a configuration server,
and starts up Envoy with the server as either ADS or xDS discovery option. The
configuration is periodically refreshed with new routes and new clusters. In
parallel, the test sends echo requests one after another through Envoy,
exercising the pushed configuration.

## Requirements

* Envoy binary `envoy` available: set `ENVOY` environment variable to the
  location of the binary, or use the default value `/usr/local/bin/envoy`
* `go-control-plane` builds successfully

## Steps

To run the script with a single ADS server:

    make integration.ads

To run the script with a single server configured as different xDS servers:

    make integration.xds

To run the script with a single server configured to use `Fetch` through HTTP:

    make integration.rest

You should see runs of configuration push events and request batch reports. The
test executes batches of requests to exercise multiple listeners, routes, and
clusters, and records the number of successful and failed requests. The test is
successful if at least one batch passes through all requests (e.g. Envoy
eventually converges to use the latest pushed configuration) for each run.

## Customizing the test driver

You can run ```bin/test -help``` to get a list of the cli flags that
the test program accepts.  There are also comments in ```main.go```.

## Using the pprof profiler

One customization is to run the go language profiler [pprof](https://github.com/DataDog/go-profiler-notes/blob/main/pprof.md). See also <https://golang.org/pkg/runtime/pprof/>.

The profiler is normally off because it adds overhead to the tests. You can turn 
it on with the command line option `--pprof`. There is an environment variable 
`PPROF` for the `make` commands shown above. For example:

    (export PPROF=true; make integration.xds)

The test will then write files of the form `block_profile_xds.pb.gz`. The files
get written to the root of the project, in the same place as the envoy logs.

You can use `go tool pprof bin/test <file name>` to analyze the profile data. 
