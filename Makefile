EXECUTABLE := drone-git-push
GO ?= go
GOFILES := $(shell find . -name "*.go" -type f)
HAS_GO = $(shell hash $(GO) > /dev/null 2>&1 && echo "GO" || echo "NOGO" )

ifneq ($(shell uname), Darwin)
	EXTLDFLAGS = -extldflags "-static" $(null)
else
	EXTLDFLAGS =
endif

ifeq ($(HAS_GO), GO)
	GOPATH ?= $(shell $(GO) env GOPATH)
	export PATH := $(GOPATH)/bin:$(PATH)

	CGO_EXTRA_CFLAGS := -DSQLITE_MAX_VARIABLE_NUMBER=32766
	CGO_CFLAGS ?= $(shell $(GO) env CGO_CFLAGS) $(CGO_EXTRA_CFLAGS)
endif

ifeq ($(OS), Windows_NT)
	GOFLAGS := -v -buildmode=exe
	EXECUTABLE ?= $(EXECUTABLE).exe
else ifeq ($(OS), Windows)
	GOFLAGS := -v -buildmode=exe
	EXECUTABLE ?= $(EXECUTABLE).exe
else
	GOFLAGS := -v
	EXECUTABLE ?= $(EXECUTABLE)
endif

ifneq ($(DRONE_TAG),)
	VERSION ?= $(DRONE_TAG)
else
	VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
endif

TAGS ?=
LDFLAGS ?= -X 'main.Version=$(VERSION)'

all: build ## Build the project.

test: ## Run tests with coverage.
	@$(GO) test -v -cover -coverprofile coverage.txt ./... && echo "\n==>\033[32m Ok\033[m\n" || exit 1

install: $(GOFILES) ## Install the binary.
	$(GO) install -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)'

build: $(EXECUTABLE) ## Build the binary.

$(EXECUTABLE): $(GOFILES)
	$(GO) build -v -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o $@

build_linux_arm: ## Build for Linux ARM.
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GO) build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm/$(EXECUTABLE)

build_linux_amd64: ## Build for Linux AMD64.
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/amd64/$(EXECUTABLE)

build_linux_arm64: ## Build for Linux ARM64.
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO) build -a -tags '$(TAGS)' -ldflags '$(EXTLDFLAGS)-s -w $(LDFLAGS)' -o release/linux/arm64/$(EXECUTABLE)

coverage: ## Remove main.go from coverage report.
	sed -i '/main.go/d' coverage.txt

clean: ## Clean build artifacts.
	$(GO) clean -x -i ./...
	rm -rf coverage.txt $(EXECUTABLE)

version: ## Print the version.
	@echo $(VERSION)

.PHONY: lint
lint: ## Run golangci-lint
	golangci-lint run -v

.PHONY: fmt
fmt: ## Format code with golangci-lint
	golangci-lint fmt

.PHONY: help
help: ## Print this help message.
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo ""
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
