module github.com/envoyproxy/go-control-plane/contrib

go 1.22.0

toolchain go1.22.5

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78
	github.com/envoyproxy/go-control-plane/envoy v1.32.4
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.5
)

require (
	cel.dev/expr v0.19.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250102185135-69823020774d // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250102185135-69823020774d // indirect
)
