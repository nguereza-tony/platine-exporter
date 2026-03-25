# 🚀 Platine Exporter Makefile

APP_NAME=platine-exporter
CMD_PATH=./cmd
BUILD_DIR=bin

GO=go

# Default target

.PHONY: all
all: build

# 🔨 Build binary

.PHONY: build
build:
$(GO) build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)

# ▶️ Run (dev)

.PHONY: run
run:
$(GO) run $(CMD_PATH)

# 🧪 Run with custom flags

.PHONY: run-dev
run-dev:
$(GO) run $(CMD_PATH) -log=app.log -addr=:8080 -workers=4

# 🧹 Clean build artifacts

.PHONY: clean
clean:
rm -rf $(BUILD_DIR)

# 📦 Install binary globally

.PHONY: install
install:
$(GO) install $(CMD_PATH)

# 🔍 Format code

.PHONY: fmt
fmt:
$(GO) fmt ./...

# 🔎 Lint (basic)

.PHONY: lint
lint:
$(GO) vet ./...

# 🧪 Test

.PHONY: test
test:
$(GO) test ./...

# 📊 Coverage

.PHONY: coverage
coverage:
$(GO) test -cover ./...

# 📈 Benchmark

.PHONY: bench
bench:
$(GO) test -bench=. ./...

# 📦 Tidy dependencies

.PHONY: tidy
tidy:
$(GO) mod tidy

# 🔄 Download dependencies

.PHONY: deps
deps:
$(GO) mod download

# 🐳 Docker build (optional)

.PHONY: docker-build
docker-build:
docker build -t $(APP_NAME) .

# 🚀 Full pipeline

.PHONY: all-in-one
all-in-one: tidy fmt lint test build
