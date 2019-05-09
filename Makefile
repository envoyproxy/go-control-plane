.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 		:= github.com/envoyproxy/go-control-plane

# Pure Go sources (not vendored and not generated)
GOFILES		= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GODIRS		= $(shell go list -f '{{.Dir}}' ./... \
						| grep -vFf <(go list -f '{{.Dir}}' ./vendor/...))

.PHONY: build
build:
	@echo "--> building"
	@go build ./...

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
	@goimports -local $(PKG) -w -l $(GOFILES)

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
	@echo "--> generating pb.go files"
	$(SHELL) build/generate_protos.sh

$(BINDIR):
	@mkdir -p $(BINDIR)
