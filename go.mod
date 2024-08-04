module github.com/envoyproxy/go-control-plane

go 1.21

replace (
	github.com/envoyproxy/go-control-plane/envoy => ./envoy
	github.com/envoyproxy/go-control-plane/ratelimit => ./ratelimit
)

require (
	github.com/envoyproxy/go-control-plane/envoy v0.0.0-00010101000000-000000000000
	github.com/envoyproxy/go-control-plane/ratelimit v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.9.0
	go.uber.org/goleak v1.3.0
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cel.dev/expr v0.15.0 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cncf/xds/go v0.0.0-20240423153145-555b57ec207b // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
