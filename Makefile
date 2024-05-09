BIN_DIR := $(CURDIR)/bin

GOBIN         = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN         = $(shell go env GOPATH)/bin
endif
GOIMPORTS     = $(GOBIN)/goimports

GOLANGCI_TAG ?= 1.58.1

format: $(GOIMPORTS)
	GO111MODULE=on go list -f '{{.Dir}}' ./... | xargs $(GOIMPORTS) -w

$(GOIMPORTS):
	go install golang.org/x/tools/cmd/goimports@latest

GOLANGCI_BIN := $(shell command -v golangci-lint 2> /dev/null)

install-lint:
	$(info Installing golangci-lint v$(GOLANGCI_TAG))
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v$(GOLANGCI_TAG)

lint:
ifeq ($(wildcard $(GOLANGCI_BIN)),)
	install-lint
	GOLANGCI_BIN := $(GOBIN)/golangci-lint
endif
	$(info Running lint against changed files...)
	$(GOLANGCI_BIN) run \
		--new-from-rev=master \
		--config=.golangci.yml \
		./...

test:
	$(info Running tests...)
	CGO_ENABLED=1 go test -race ./...

BUILD_ENVPARMS ?= CGO_ENABLED=0
BUILD_PATHS ?= ./cmd/linda

build:
	$(info Building...)
	$(BUILD_ENVPARMS) go build -o $(BIN_DIR)/linda $(BUILD_PATHS)

RUN_ENVPARMS ?= CONFIG_PATH=./configs/config-local.yml

run:
	$(info Running...)
	$(BUILD_ENVPARMS) $(RUN_ENVPARMS) go run $(BUILD_PATHS)

precommit: format build test lint

.PHONY: format lint test precommit run build install-lint