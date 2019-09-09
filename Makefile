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

#-----------------
#-- integration
#-----------------
.PHONY: $(BINDIR)/test integration integration.ads integration.xds integration.rest integration.ads.tls

$(BINDIR)/test:
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
