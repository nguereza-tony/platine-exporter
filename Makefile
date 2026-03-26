# 🚀 Platine Exporter Makefile
EXECUTABLE=platine-exporter
WINDOWS=$(EXECUTABLE)_windows_amd64.exe
LINUX=$(EXECUTABLE)_linux_amd64
LINUX_ARM=$(EXECUTABLE)_linux_arm64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always)
CMD_PATH=./cmd/main.go
BUILD_DIR=bin
GO=go

.PHONY: all test clean

all: test build ## Build and run tests

test: ## Run unit tests
	$(GO) test ./...

build: windows linux linux-arm darwin ## Build binaries
	@echo version: $(VERSION)

windows: $(WINDOWS) ## Build for Windows

linux: $(LINUX) ## Build for Linux

linux-arm: $(LINUX_ARM) ## Build for Linux (ARM)

darwin: $(DARWIN) ## Build for Darwin (macOS)

$(WINDOWS):
	env GOOS=windows GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" $(CMD_PATH)

$(LINUX):
	env GOOS=linux GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(LINUX) -ldflags="-s -w -X main.version=$(VERSION)" $(CMD_PATH)

$(LINUX_ARM):
	env GOOS=linux GOARCH=arm64 $(GO) build -o $(BUILD_DIR)/$(LINUX_ARM) -ldflags="-s -w -X main.version=$(VERSION)" $(CMD_PATH)

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 $(GO) build -o $(BUILD_DIR)/$(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" $(CMD_PATH)

clean: ## Remove previous build
	rm -f $(BUILD_DIR)/$(WINDOWS) $(BUILD_DIR)/$(LINUX) $(BUILD_DIR)/$(LINUX_ARM) $(BUILD_DIR)/$(DARWIN)

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
