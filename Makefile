.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 		:= github.com/envoyproxy/go-control-plane

.PHONY: build
build:
	@echo "--> building"
	@go build ./pkg/... ./envoy/...

.PHONY: clean
clean:
	@echo "--> cleaning compiled objects and binaries"
	@go clean -tags netgo -i ./...
	@rm -rf $(BINDIR)/*

.PHONY: test
test:
	@echo "--> running unit tests"
	@go test ./pkg/...

.PHONY: cover
cover:
	@echo "--> running coverage tests"
	@build/coverage.sh

.PHONY: format
format:
	@echo "--> formatting code with 'goimports' tool"
	@goimports -local $(PKG) -w -l pkg envoy

#-----------------
#-- integration
#-----------------
.PHONY: $(BINDIR)/test integration integration.ads integration.xds integration.rest integration.ads.tls

$(BINDIR)/test:
	@echo "--> building test binary"
	@go build -race -o $@ pkg/test/main/main.go

integration: integration.xds integration.ads integration.rest integration.ads.tls

integration.ads: $(BINDIR)/test
	env XDS=ads build/integration.sh

integration.xds: $(BINDIR)/test
	env XDS=xds build/integration.sh

integration.rest: $(BINDIR)/test
	env XDS=rest build/integration.sh

integration.ads.tls: $(BINDIR)/test
	env XDS=ads build/integration.sh -tls

#-----------------
#-- code generaion
#-----------------

generate:
	@echo "--> vendoring protobufs"
	@go mod vendor
	@echo "--> generating pb.go files"
	$(SHELL) build/generate_protos.sh

$(BINDIR):
	@mkdir -p $(BINDIR)

.PHONY: generate-patch
generate-patch:
	@echo "--> patching generated code due to issue with protoc-gen-validate"
	find envoy -type f -print0 |\
		xargs -0 sed -i 's#"envoy/api/v2/core"#"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"#g'
