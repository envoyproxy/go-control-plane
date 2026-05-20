module github.com/envoyproxy/go-control-plane/envoy

go 1.25.0

toolchain go1.25.7

// Used to resolve import issues related to go-control-plane package split (https://github.com/envoyproxy/go-control-plane/issues/1074)
replace github.com/envoyproxy/go-control-plane@v0.13.4 => ../

require (
	github.com/cncf/xds/go v0.0.0-20260202195803-dba9d589def2
	github.com/envoyproxy/go-control-plane v0.14.0
	github.com/envoyproxy/protoc-gen-validate v1.3.3
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	github.com/prometheus/client_model v0.6.2
	go.opentelemetry.io/proto/otlp v1.10.0
	google.golang.org/genproto/googleapis/api v0.0.0-20260226221140-a57be14db171
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171
	google.golang.org/grpc v1.81.1
	google.golang.org/protobuf v1.36.11
)

require (
	cel.dev/expr v0.25.1 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.34.0 // indirect
)
