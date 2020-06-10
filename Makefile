.DEFAULT_GOAL	:= build

#------------------------------------------------------------------------------
# Variables
#------------------------------------------------------------------------------

SHELL 	:= /bin/bash
BINDIR	:= bin
PKG 		:= github.com/envoyproxy/go-control-plane

# Build args for testing
UNAME_S := $(shell uname -s)
TEST_BIN_FLAGS := CGO_ENABLED=1
ifeq ($(UNAME_S), Darwin)
	TEST_BIN_FLAGS += GOOS=darwin GOARCH=amd64
else 
	TEST_BIN_FLAGS += GOOS=linux GOARCH=amd64
endif

.PHONY: build
build:
	@echo "--> building go-control-plane"
	@go build ./pkg/... ./envoy/...

.PHONY: build.test
build.test:
	@echo "--> building go-control-plane test binary"
	@$(TEST_BIN_FLAGS) go build -race -a -tags netgo -ldflags '-w -extldflags "-static"' -o $(BINDIR)/test  pkg/test/main/main.go

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
	@go test ./pkg/...

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

#-----------------
#-- integration
#-----------------
.PHONY: $(BINDIR)/test integration integration.ads integration.ads.tls integration.ads.v3 integration.xds integration.xds.delta integration.xds.v3 integration.rest integration.rest.v3

$(BINDIR)/test:
	echo "Building test binary"
	$(TEST_BIN_FLAGS) go build -race -a -tags netgo -ldflags '-w -extldflags "-static"' -o $@ pkg/test/main/main.go

integration: integration.xds integration.xds.delta integration.xds.v3 integration.ads integration.ads.tls integration.ads.v3 integration.rest integration.rest.v3

integration.ads: $(BINDIR)/test
	env XDS=ads build/integration.sh

integration.ads.tls: $(BINDIR)/test
	env XDS=ads build/integration.sh -tls

integration.ads.v3: $(BINDIR)/test
	env XDS=ads SUFFIX=v3 build/integration.sh

integration.xds: $(BINDIR)/test
	env XDS=xds build/integration.sh

integration.xds.delta: $(BINDIR)/test
	env XDS=delta build/integration.sh

integration.xds.v3: $(BINDIR)/test
	env XDS=xds SUFFIX=v3 build/integration.sh

integration.rest: $(BINDIR)/test
	env XDS=rest build/integration.sh

integration.rest.v3: $(BINDIR)/test
	env XDS=rest SUFFIX=v3 build/integration.sh

