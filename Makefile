.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 		:= github.com/envoyproxy/go-control-plane

.PHONY: build
build:
	@echo "--> building go-control-plane"
	@go build ./pkg/... ./envoy/...

.PHONY: clean
clean:
	@echo "--> cleaning compiled objects and binaries"
	@go clean -tags netgo -i ./...
	@go mod tidy
	@rm -rf $(BINDIR)
	@rm -rf *.log

.PHONY: test
test:
	@echo "--> testing go-control-plane"
	@go test -count=1 ./pkg/...

.PHONY: cover
cover:
	@echo "--> checking coverage"
	@build/coverage.sh

.PHONY: format
format:
	@echo "--> formatting"
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
.PHONY: $(BINDIR)/test $(BINDIR)/upstream integration integration.ads integration.ads.tls integration.ads.delta integration.ads.delta.v3 integration.ads.v3 integration.xds integration.xds.delta integration.xds.v3 integration.rest integration.rest.v3 integration.xds.mux.v3

$(BINDIR)/upstream:
	@go build -race -o $@ internal/upstream/main.go

$(BINDIR)/test:
	@echo "Building test binary"
	@go build -race -a -tags netgo -ldflags '-w -extldflags "-static"' -o $@ pkg/test/main/main.go

integration: integration.xds integration.xds.delta integration.xds.delta.v3 integration.xds.v3 integration.ads integration.ads.tls integration.ads.v3 integration.ads.delta integration.ads.delta.v3 integration.rest integration.rest.v3

integration.ads: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads build/integration.sh

integration.ads.tls: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads build/integration.sh -tls

integration.ads.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads SUFFIX=v3 build/integration.sh

integration.ads.delta: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads-delta build/integration.sh

integration.ads.delta.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads-delta SUFFIX=v3 build/integration.sh

integration.xds: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds build/integration.sh

integration.xds.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds SUFFIX=v3 build/integration.sh

integration.xds.delta: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=delta build/integration.sh

integration.xds.delta.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=delta SUFFIX=v3 build/integration.sh

integration.rest: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=rest build/integration.sh

integration.rest.v3: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=rest SUFFIX=v3 build/integration.sh

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
