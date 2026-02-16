module github.com/envoyproxy/go-control-plane/xdsmatcher

go 1.25.0

toolchain go1.25.7

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/cncf/xds/go v0.0.0-20251110193048-8bfbf64dc13e
	github.com/envoyproxy/go-control-plane/envoy v1.36.0
	github.com/stretchr/testify v1.11.1
	google.golang.org/protobuf v1.36.11
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cel.dev/expr v0.24.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.3.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20251029180050-ab9386a59fda // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
