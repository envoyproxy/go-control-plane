.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 		:= github.com/envoyproxy/go-control-plane

.PHONY: build
build:
	@go build ./pkg/... ./envoy/...

.PHONY: test
test:
	@go test ./pkg/...

.PHONY: cover
cover:
	@build/coverage.sh

.PHONY: format
format:
	@goimports -local $(PKG) -w -l pkg

.PHONY: create_version
create_version:
	./scripts/create_version.sh

.PHONY: check_version_dirty
check_version_dirty:
	./scripts/check_version_dirty.sh

.PHONY: examples
examples:
	@pushd examples/dyplomat && go build ./... && popd

#-----------------
#-- integration
#-----------------
.PHONY: $(BINDIR)/test $(BINDIR)/upstream integration integration.ads integration.ads.v3 integration.xds integration.xds.v3 integration.rest integration.rest.v3 integration.ads.tls integration.xds.mux.v3

$(BINDIR)/upstream:
	@go build -race -o $@ internal/upstream/main.go

$(BINDIR)/test:
	@go build -race -o $@ pkg/test/main/main.go

integration: integration.xds integration.xds.v3 integration.ads integration.ads.v3 integration.rest integration.rest.v3 integration.ads.tls

integration.ads: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads build/integration.sh

integration.ads.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads SUFFIX=v3 build/integration.sh

integration.xds: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds build/integration.sh

integration.xds.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds SUFFIX=v3 build/integration.sh

integration.rest: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=rest build/integration.sh

integration.rest.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=rest SUFFIX=v3 build/integration.sh

integration.ads.tls: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads build/integration.sh -tls

integration.xds.mux.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds SUFFIX=v3 build/integration.sh -mux

#--------------------------------------
#-- example xDS control plane server
#--------------------------------------
.PHONY: $(BINDIR)/example example

$(BINDIR)/example:
	@go build -race -o $@ internal/example/main/main.go

example: $(BINDIR)/example
	@build/example.sh
