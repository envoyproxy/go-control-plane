module github.com/envoyproxy/go-control-plane/ratelimit

go 1.23.0

toolchain go1.23.6

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/envoyproxy/go-control-plane/envoy v1.32.4
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/cncf/xds/go v0.0.0-20250501225837-2ac532fd4443 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250528174236-200df99c418a // indirect
)
