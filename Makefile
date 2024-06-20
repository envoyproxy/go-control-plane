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

.PHONY: clean
clean:
	@echo "--> cleaning compiled objects and binaries"
	@go clean -tags netgo -i ./...
	@go mod tidy
	@rm -rf $(BINDIR)
	@rm -rf *.log

# TODO(mattklein123): See the note in TestLinearConcurrentSetWatch() for why we set -parallel here
# This should be removed.
.PHONY: test
test:
	@go test -race -v -timeout 30s -count=1 -parallel 100 ./pkg/...

.PHONY: cover
cover:
	@build/coverage.sh

.PHONY: format
format:
	@goimports -local $(PKG) -w -l pkg

.PHONY: examples
examples:
	@pushd examples/dyplomat && go build ./... && popd

.PHONY: lint
lint:
	@docker run \
		--rm \
		--volume $$(pwd):/src \
		--workdir /src \
		golangci/golangci-lint:latest \
	golangci-lint -v run

#-----------------
#-- integration
#-----------------
.PHONY: $(BINDIR)/test $(BINDIR)/upstream integration integration.ads integration.xds integration.rest integration.xds.mux integration.xds.delta integration.ads.delta

$(BINDIR)/upstream:
	@go build -race -o $@ internal/upstream/main.go

$(BINDIR)/test:
	@echo "Building test binary"
	@go build -race -o $@ pkg/test/main/main.go

integration: integration.xds integration.ads integration.rest integration.xds.mux integration.xds.delta integration.ads.delta

integration.ads: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=ads build/integration.sh

integration.xds: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds build/integration.sh

integration.rest: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=rest build/integration.sh

integration.xds.mux: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds build/integration.sh -mux

integration.xds.delta: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=delta build/integration.sh

integration.ads.delta: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=delta-ads build/integration.sh

#--------------------------------------
#-- example xDS control plane server
#--------------------------------------
.PHONY: $(BINDIR)/example example

$(BINDIR)/example:
	@go build -race -o $@ internal/example/main/main.go

example: $(BINDIR)/example
	@build/example.sh

.PHONY: docker_tests
docker_tests:
	docker build --pull -f Dockerfile.ci . -t gcp_ci && \
	docker run -v $$(pwd):/go-control-plane $$(tty -s && echo "-it" || echo) gcp_ci /bin/bash -c /go-control-plane/build/do_ci.sh

.PHONY: tidy-all
tidy-all:
	go mod tidy
	make -C examples/dyplomat tidy
	make -C xdsmatcher tidy
