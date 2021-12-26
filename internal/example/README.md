# Example xDS Server

This is an example of a trivial xDS V3 control plane server.  It serves an Envoy configuration that's roughly equivalent to the one used by the Envoy ["Quick Start"](https://www.envoyproxy.io/docs/envoy/latest/start/start#quick-start-to-run-simple-example) docs: a simple http proxy.  You can run the example using the project top-level Makefile, e.g.:

```
$ make example
```

The Makefile builds the example server and then runs `build/example.sh` which runs both Envoy and the example server.  The example server serves a configuration defined in `internal/example/resource.go`.  If everything works correctly, you should be able to open a browser to [http://localhost:10000](http://localhost:10000) and see Envoy's website.

## Files

* [main/main.go](main/main.go) is the example program entrypoint.  It instantiates the cache and xDS server and runs the xDS server process.
* [resource.go](resource.go) generates a `Snapshot` structure which describes the configuration that the xDS server serves to Envoy.
* [server.go](server.go) runs the xDS control plane server.
* [logger.go](logger.go) implements the `pkg/log/Logger` interface which provides logging services to the cache.
