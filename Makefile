.PHONY: clean run build install dep test lint format docker gql tools tools-golangci-lint tools-gotestsum
export PATH := $(abspath bin/):${PATH}
BIN_NAME					?= codegen
DEFAULT_INSTALL_DIR			:= $(go env GOPATH)/bin
DEFAULT_ARCH				:= $(shell go env GOARCH)
DEFAULT_GOOS				:= $(shell go env GOOS)
ARCH						?= $(DEFAULT_ARCH)
GOOS						?= $(DEFAULT_GOOS)
INSTALL_DIR					?= $(DEFAULT_INSTALL_DIR)
.DEFAULT_GOAL := run

VERSION   := $(shell git describe --tags || echo "v0.0.0")
VER_CUT   := $(shell echo $(VERSION) | cut -c2-)

# Dependency versions
GOLANGCI_VERSION   = v1.57.2

build:
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(ARCH) \
		go build -o bin/$(BIN_NAME) ./cmd/$(BIN_NAME)

run: build
	@./bin/$(BIN_NAME)

all: clean target

clean:
	@rm -rf bin
	
install: build
	@install -d $(INSTALL_DIR)
	@rm -f $(INSTALL_DIR)/$(BIN_NAME)
	@cp bin/* $(INSTALL_DIR)/

dep: 
	@go mod tidy

test:
	@go test ./...

lint:
	@golangci-lint version
	@golangci-lint run

format:
	@golangci-lint run --fix

migration: build
	./bin/codegen -output=./pkg/migrations -package=migrations -generators=migration -migration.file-name="${name}"

tools-golangci-lint:
	@mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINARY=golangci-lint bash -s -- ${GOLANGCI_VERSION}

tools: tools-golangci-lint