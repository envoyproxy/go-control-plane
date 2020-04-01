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

#-----------------
#-- integration
#-----------------
.PHONY: $(BINDIR)/test integration integration.ads integration.ads.v3 integration.xds integration.xds.v3 integration.rest integration.ads.tls

$(BINDIR)/test:
	@go build -race -o $@ pkg/test/main/main.go

integration: integration.xds integration.xds.v3 integration.ads integration.ads.v3 integration.rest integration.ads.tls

integration.ads: $(BINDIR)/test
	env XDS=ads build/integration.sh

integration.ads.v3: $(BINDIR)/test
	env XDS=ads SUFFIX=v3 build/integration.sh

integration.xds: $(BINDIR)/test
	env XDS=xds build/integration.sh

integration.xds.v3: $(BINDIR)/test
	env XDS=xds SUFFIX=v3 build/integration.sh

integration.rest: $(BINDIR)/test
	env XDS=rest build/integration.sh

integration.ads.tls: $(BINDIR)/test
	env XDS=ads build/integration.sh -tls
