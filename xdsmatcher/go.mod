module github.com/envoyproxy/go-control-plane/xdsmatcher

go 1.23.0

toolchain go1.23.6

replace github.com/envoyproxy/go-control-plane/envoy => ../envoy

require (
	github.com/cncf/xds/go v0.0.0-20241223141626-cff3c89139a3
	github.com/envoyproxy/go-control-plane/envoy v1.32.4
	github.com/stretchr/testify v1.10.0
	google.golang.org/protobuf v1.36.5
	gopkg.in/yaml.v2 v2.4.0
)

require (
	cel.dev/expr v0.19.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
