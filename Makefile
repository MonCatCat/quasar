BUILD_DIR ?= $(CURDIR)/build

mkdirs = mkdir -p $(BUILD_DIR)

# Install & build

go-mod:
	go mod tidy
	go mod verify
	go mod download

lint:
	@echo "SKIPPED"
	#go run github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout=10m

build:
	scripts/build

# Testing

PACKAGES_UNIT=$(shell go list ./x/epochs/... ./x/intergamm/... ./x/qbank/... ./x/qoracle/... | grep -E -v "simapp|e2e")

test:
	go test -mod=readonly -v $(PACKAGES_UNIT)

test-cover:
	${mkdirs}
	go test -mod=readonly -timeout 30m -coverprofile=$(BUILD_DIR)/coverage.txt -covermode=atomic $(PACKAGES_UNIT)

test-simulation:
	ignite chain simulate -v

# Documentation

docs-gen:
	scripts/gen_grpc_doc

docs-serve:
	scripts/serve_doc_docker

# Run targets

run:
	scripts/run

run-silent:
	scripts/run > q.log 2>&1

.PHONY: go.mod build test test-simulation docs-gen docs-serve run run-silent
