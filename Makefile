GOMOD := $(shell go list -m)
BUILD_DIR ?= $(CURDIR)/build
MOCKS_DIR = $(CURDIR)/testutil/mock

# Install & build

mkdirs:
	@mkdir -p $(BUILD_DIR)
	@mkdir -p $(MOCKS_DIR)

go-mod:
	go mod tidy
	go mod verify
	go mod download

lint:
	@echo "SKIPPED"
	#go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=10m

build: mkdirs
	scripts/build

.PHONY: mkdirs go-mod lint build

# Testing

PACKAGES_UNIT=$(shell go list ./x/epochs/... ./x/intergamm/... ./x/qbank/... ./x/qoracle/... | grep -E -v "simapp|e2e" | grep -E -v "x/qoracle/client/cli")

mocks: mkdirs
	mockgen -package=mock -destination=./testutil/mock/ica_mocks.go $(GOMOD)/x/intergamm/types ICAControllerKeeper

test:
	go test -mod=readonly -v $(PACKAGES_UNIT)

test-path:
	go test -mod=readonly -v $(path)

test-ibc-transfer:
	go test -mod=readonly -v -timeout 99999s demos/ibc-test-framework/ibc_transfer_test.go

test-cover: mkdirs
	go test -mod=readonly -timeout 30m -coverprofile=$(BUILD_DIR)/coverage.txt -covermode=atomic $(PACKAGES_UNIT)

test-simulation:
	ignite chain simulate -v

.PHONY: mocks $(MOCKS_DIR) test test-path test-cover test-simulation

# Documentation

docs-gen:
	scripts/gen_grpc_doc

docs-serve:
	scripts/serve_doc_docker

.PHONY: docs-gen docs-serve

# Run targets

run:
	scripts/run

run-silent:
	scripts/run > q.log 2>&1

.PHONY: run run-silent
