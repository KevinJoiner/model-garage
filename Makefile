.PHONY: clean run build install dep test lint format docker gql tools tools-golangci-lint tools-gotestsum
# Set the bin path
PATHINSTBIN = $(abspath ./bin)
export PATH := $(PATHINSTBIN):$(PATH)

BIN_NAME					?= codegen
DEFAULT_INSTALL_DIR			:= $(go env GOPATH)/bin
DEFAULT_ARCH				:= $(shell go env GOARCH)
DEFAULT_GOOS				:= $(shell go env GOOS)
ARCH						?= $(DEFAULT_ARCH)
GOOS						?= $(DEFAULT_GOOS)
INSTALL_DIR					?= $(DEFAULT_INSTALL_DIR)
.DEFAULT_GOAL := run

PLATFORMS=linux/amd64 linux/386 linux/arm linux/arm64 \
         windows/amd64 windows/386 \
         darwin/amd64 darwin/arm64

# Get the app version from the git tag and commit
GIT_COMMIT := $(shell git rev-parse --short HEAD)
TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
TAG_COMMIT := $(shell git rev-list -n 1 $(TAG))
VERSION := $(TAG)
ifneq ($(TAG_COMMIT), $(shell git rev-parse HEAD))
	VERSION := $(TAG)-$(GIT_COMMIT)
endif

# Dependency versions
GOLANGCI_VERSION := latest
CLICKHOUSE_INFRA_VERSION := $(shell go list -m -f '{{.Version}}' github.com/DIMO-Network/clickhouse-infra)

build:
	@CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(ARCH) \
		go build -ldflags "-X 'github.com/DIMO-Network/model-garage/pkg/version.version=$(VERSION)'" \
		-o bin/$(BIN_NAME) ./cmd/$(BIN_NAME)

build-all-ruptela-convert:
	mkdir -p bin
	$(foreach platform,$(PLATFORMS),\
		$(eval OS := $(word 1,$(subst /, ,$(platform)))) \
		$(eval ARCH := $(word 2,$(subst /, ,$(platform)))) \
		GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/ruptela-convert-$(OS)-$(ARCH)$(if $(filter windows,$(OS)),.exe,) ./pkg/ruptela/cmd/ruptela-convert/ || exit 1; \
	)

run: build
	@./bin/$(BIN_NAME)

all: clean target

clean:
	@rm -rf $(PATHINSTBIN)
	
install: build
	@install -d $(INSTALL_DIR)
	@rm -f $(INSTALL_DIR)/$(BIN_NAME)
	@cp $(PATHINSTBIN)/* $(INSTALL_DIR)/

dep: 
	@go mod tidy

test:
	@go test ./...

lint:
	@golangci-lint version
	@golangci-lint run --timeout=5m

format:
	@golangci-lint run --fix

migration:
	migration -output=./pkg/migrations -package=migrations -filename="${name}"

tools-golangci-lint:
	@mkdir -p $(PATHINSTBIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINARY=golangci-lint bash -s -- ${GOLANGCI_VERSION}

tools-migration: ## Install migration tool
	@mkdir -p $(PATHINSTBIN)
	GOBIN=$(PATHINSTBIN) go install github.com/DIMO-Network/clickhouse-infra/cmd/migration@${CLICKHOUSE_INFRA_VERSION}

tools: tools-golangci-lint

clickhouse:
	go run ./cmd/clickhouse-container

generate: generate-nativestatus generate-ruptela # Generate all files for the repository
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/vss/vehicle-structs.go -custom.template-file=./internal/generator/vehicle.tmpl -custom.format=true

generate-nativestatus: # Generate all files for nativestatus
	go run ./cmd/codegen -convert.package=nativestatus -generators=convert -convert.output-file=./pkg/nativestatus/vehicle-convert-funcs_gen.go
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/nativestatus/vehicle-v1-convert_gen.go -custom.template-file=./pkg/nativestatus/convertv1.tmpl -custom.format=true
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/nativestatus/vehicle-v2-convert_gen.go -custom.template-file=./pkg/nativestatus/convertv2.tmpl -custom.format=true

generate-ruptela: # Generate all files for ruptela
	go run ./cmd/codegen -convert.package=ruptela -generators=convert -convert.output-file=./pkg/ruptela/vehicle-convert-funcs_gen.go -definitions=./pkg/ruptela/schema/ruptela-definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/ruptela/vehicle-v1-convert_gen.go -custom.template-file=./pkg/ruptela/codegen/convert-status.tmpl -custom.format=true -definitions=./pkg/ruptela/schema/ruptela-definitions.yaml
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/ruptela/vehicle-location-convert_gen.go -custom.template-file=./pkg/ruptela/codegen/convert-location.tmpl -custom.format=true -definitions=./pkg/ruptela/schema/ruptela-definitions.yaml
	go run ./pkg/ruptela/codegen 
	go run ./cmd/codegen -generators=custom -custom.output-file=./pkg/ruptela/vehicle-convert-replace_gen.go -custom.template-file=./pkg/ruptela/codegen/convert-replace.tmpl -custom.format=true -definitions=./pkg/ruptela/schema/ruptela-definitions.yaml
