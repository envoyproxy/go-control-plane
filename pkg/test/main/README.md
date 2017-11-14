# End-to-end testing script

This directory contains a basic end-to-end testing script.
The script sets up a configuration cache, stands up a configuration server,
and starts up Envoy with the server as either ADS or xDS discovery option. The
configuration is periodically refreshed with new routes and new clusters. In
parallel, the test sends echo requests one after another through Envoy,
exercising the pushed configuration.

## Requirements

* Envoy binary `envoy` available on the path
* `go-control-plane` builds successfully

## Steps

To run the script with a single ADS server:

    go run pkg/test/main/main.go --logtostderr -v 10 --ads=true

To run the script with EDS, CDS, RDS, and LDS servers:

    go run pkg/test/main/main.go --logtostderr -v 10 --ads=false

You should see configuration push events logged as well as individual requests
reports (`OK` or `ERROR` depending whether the proxy is ready or not to accept requests).
