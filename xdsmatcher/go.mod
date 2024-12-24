module github.com/envoyproxy/go-control-plane/xdsmatcher

go 1.22.8

replace (
	github.com/envoyproxy/go-control-plane/envoy => ../envoy
	// Used to resolve import issues related to go-control-plane package split (https://github.com/envoyproxy/go-control-plane/issues/1074)
	github.com/envoyproxy/go-control-plane@v0.13.2 => ../
)

require (
	github.com/cncf/xds/go v0.0.0-20240723142845-024c85f92f20
	github.com/envoyproxy/go-control-plane v0.13.2
	github.com/envoyproxy/go-control-plane/envoy v1.32.2
	github.com/stretchr/testify v1.10.0
	google.golang.org/protobuf v1.36.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cel.dev/expr v0.16.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.67.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
