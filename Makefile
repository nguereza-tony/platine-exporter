# 🚀 Platine Exporter Makefile
EXECUTABLE=platine-exporter
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)
CMD_PATH=./cmd/main.go
BUILD_DIR=bin
GO=go

.PHONY: all test clean

all: test build ## Build and run tests

test: ## Run unit tests
	$(GO) test ./...

build: windows linux darwin ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 $(GO) build -v -o $(BUILD_DIR)/$(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" $(CMD_PATH)

$(LINUX):
	env GOOS=linux GOARCH=amd64 $(GO) build -v -o $(BUILD_DIR)/$(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" $(CMD_PATH)

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 $(GO) build -v -o $(BUILD_DIR)/$(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" $(CMD_PATH)

clean: ## Remove previous build
	rm -f $(BUILD_DIR)/$(WINDOWS) $(BUILD_DIR)/$(LINUX) $(BUILD_DIR)/$(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
