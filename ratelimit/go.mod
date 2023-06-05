module github.com/envoyproxy/go-control-plane/ratelimit

go 1.20

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/envoyproxy/go-control-plane/envoy v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/cncf/xds/go v0.0.0-20230428030218-4003588d1b74 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230530153820-e85fd2cbaebc // indirect
)
