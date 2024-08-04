module github.com/envoyproxy/go-control-plane/ratelimit

go 1.21

toolchain go1.22.0

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/envoyproxy/go-control-plane/envoy v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/cncf/xds/go v0.0.0-20240423153145-555b57ec207b // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
)
