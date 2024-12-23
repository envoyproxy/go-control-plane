.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 	:= github.com/envoyproxy/go-control-plane

include ./Makefile.common

.PHONY: build
build:
	@make -C pkg common/build
	@make -C envoy common/build
	@make -C ratelimit common/build
	@make -C xdsmatcher common/build 

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
	@scripts/coverage.sh

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
	env XDS=ads scripts/integration.sh

integration.xds: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds scripts/integration.sh

integration.rest: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=rest scripts/integration.sh

integration.xds.mux: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=xds scripts/integration.sh -mux

integration.xds.delta: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=delta scripts/integration.sh

integration.ads.delta: $(BINDIR)/test $(BINDIR)/upstream
	env XDS=delta-ads scripts/integration.sh

#--------------------------------------
#-- example xDS control plane server
#--------------------------------------
.PHONY: $(BINDIR)/example example

$(BINDIR)/example:
	@go build -race -o $@ internal/example/main/main.go

example: $(BINDIR)/example
	@scripts/example.sh

.PHONY: docker_tests
docker_tests:
	docker build --pull -f Dockerfile.ci . -t gcp_ci && \
	docker run -v $$(pwd):/go-control-plane $$(tty -s && echo "-it" || echo) gcp_ci /bin/bash -c /go-control-plane/scripts/do_ci.sh

.PHONY: tidy-all
tidy-all: common/tidy
	make -C contrib common/tidy
	make -C envoy common/tidy
	make -C examples/dyplomat common/tidy
	make -C ratelimit common/tidy
	make -C xdsmatcher common/tidy

.PHONY: multimod/verify
multimod/verify: $(MULTIMOD)
	@echo "Validating versions.yaml"
	$(MULTIMOD) verify

.PHONY: multimod/prerelease
multimod/prerelease: $(MULTIMOD)
	$(MULTIMOD) prerelease -s=true -b=false -v ./versions.yaml -m ${MODSET}
	$(MAKE) tidy-all

COMMIT?=HEAD
REMOTE?=git@github.com:envoyproxy/go-control-plane.git
.PHONY: multimod/tag
multimod/tag: $(MULTIMOD)
	$(MULTIMOD) verify
	$(MULTIMOD) tag -m ${MODSET} -c ${COMMIT} --print-tags

COMMIT?=HEAD
REMOTE?=git@github.com:envoyproxy/go-control-plane.git
.PHONY: push-tags
multimod/push-tags: $(MULTIMOD)
	$(MULTIMOD) verify
	set -e; for tag in `$(MULTIMOD) tag -m ${MODSET} -c ${COMMIT} --print-tags | grep -v "Using" `; do \
		echo "pushing tag $${tag}"; \
		git push ${REMOTE} $${tag}; \
	done;
