module github.com/envoyproxy/go-control-plane/envoy

go 1.22

// Used to resolve import issues related to go-control-plane package split (https://github.com/envoyproxy/go-control-plane/issues/1074)
replace github.com/envoyproxy/go-control-plane@v0.13.4 => ../

require (
	github.com/cncf/xds/go v0.0.0-20240723142845-024c85f92f20
	github.com/envoyproxy/go-control-plane v0.13.4
	github.com/envoyproxy/protoc-gen-validate v1.1.0
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	github.com/prometheus/client_model v0.6.0
	go.opentelemetry.io/proto/otlp v1.0.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142
	google.golang.org/grpc v1.67.1
	google.golang.org/protobuf v1.35.2
)

require (
	cel.dev/expr v0.16.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
