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
.PHONY: $(BINDIR)/test $(BINDIR)/test-linux docker integration integration.ads integration.xds integration.rest integration.docker

$(BINDIR)/test:
	@echo "--> building test binary"
	@go build -race -o $@ pkg/test/main/main.go

$(BINDIR)/test-linux:
	@echo "--> building Linux test binary"
	@env GOOS=linux GOARCH=amd64 go build -race -o $@ pkg/test/main/main.go

docker:
	@echo "--> building test docker image"
	@docker build -t test .

integration: integration.ads integration.xds integration.rest

integration.ads: $(BINDIR)/test
	env XDS=ads build/integration.sh

integration.xds: $(BINDIR)/test
	env XDS=xds build/integration.sh

integration.rest: $(BINDIR)/test
	env XDS=rest build/integration.sh

integration.docker: docker
	docker run -it -e "XDS=ads" test -debug
	docker run -it -e "XDS=xds" test -debug
	docker run -it -e "XDS=rest" test -debug
	docker run -it -e "XDS=ads" test -debug -tls
	docker run -it -e "XDS=xds" test -debug -tls
	docker run -it -e "XDS=rest" test -debug -tls

#-----------------
#-- code generaion
#-----------------

generate:
	@echo "--> generating pb.go files"
	$(SHELL) build/generate_protos.sh

$(BINDIR):
	@mkdir -p $(BINDIR)
