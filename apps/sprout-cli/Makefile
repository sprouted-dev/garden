# Sprout CLI Makefile
# Cross-platform build automation for the Sprout CLI tool

BINARY_NAME := sprout
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo 'dev')
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo 'unknown')

# Go build settings
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GO_BUILD_FLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.Commit=$(COMMIT)"

# Output directories
BUILD_DIR := build
DIST_DIR := dist

# Colors for output
BOLD := \033[1m
GREEN := \033[32m
YELLOW := \033[33m
BLUE := \033[34m
RESET := \033[0m

.PHONY: help build build-all test clean install release fmt lint version
.DEFAULT_GOAL := help

help: ## Show this help message
	@echo "$(BOLD)Sprout CLI Build System$(RESET)"
	@echo "$(YELLOW)Available targets:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development builds
build: ## Build for current platform
	@echo "$(YELLOW)Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)...$(RESET)"
	@mkdir -p $(BUILD_DIR)
	@go build $(GO_BUILD_FLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .
	@echo "$(GREEN)✅ Built: $(BUILD_DIR)/$(BINARY_NAME)$(RESET)"

build-debug: ## Build with debug symbols
	@echo "$(YELLOW)Building $(BINARY_NAME) with debug info...$(RESET)"
	@mkdir -p $(BUILD_DIR)
	@go build -gcflags="all=-N -l" -o $(BUILD_DIR)/$(BINARY_NAME)-debug .
	@echo "$(GREEN)✅ Built: $(BUILD_DIR)/$(BINARY_NAME)-debug$(RESET)"

# Cross-platform release builds
release: clean test ## Build release binaries for all platforms
	@echo "$(BOLD)Building release binaries...$(RESET)"
	@mkdir -p $(DIST_DIR)
	@$(MAKE) build-linux-amd64
	@$(MAKE) build-linux-arm64
	@$(MAKE) build-darwin-amd64
	@$(MAKE) build-darwin-arm64
	@$(MAKE) build-windows-amd64
	@$(MAKE) create-checksums
	@echo "$(GREEN)✅ Release binaries built in $(DIST_DIR)/$(RESET)"
	@ls -la $(DIST_DIR)/

build-linux-amd64: ## Build for Linux AMD64
	@echo "$(BLUE)Building for Linux AMD64...$(RESET)"
	@GOOS=linux GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64 .

build-linux-arm64: ## Build for Linux ARM64
	@echo "$(BLUE)Building for Linux ARM64...$(RESET)"
	@GOOS=linux GOARCH=arm64 go build $(GO_BUILD_FLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm64 .

build-darwin-amd64: ## Build for macOS AMD64
	@echo "$(BLUE)Building for macOS AMD64...$(RESET)"
	@GOOS=darwin GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64 .

build-darwin-arm64: ## Build for macOS ARM64 (Apple Silicon)
	@echo "$(BLUE)Building for macOS ARM64...$(RESET)"
	@GOOS=darwin GOARCH=arm64 go build $(GO_BUILD_FLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64 .

build-windows-amd64: ## Build for Windows AMD64
	@echo "$(BLUE)Building for Windows AMD64...$(RESET)"
	@GOOS=windows GOARCH=amd64 go build $(GO_BUILD_FLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe .

create-checksums: ## Create SHA256 checksums for release binaries
	@echo "$(YELLOW)Creating checksums...$(RESET)"
	@cd $(DIST_DIR) && shasum -a 256 * > checksums.txt
	@echo "$(GREEN)✅ Checksums created$(RESET)"

# Testing
test: ## Run tests
	@echo "$(YELLOW)Running tests...$(RESET)"
	@go test -v ./...
	@echo "$(GREEN)✅ Tests passed$(RESET)"

test-coverage: ## Run tests with coverage
	@echo "$(YELLOW)Running tests with coverage...$(RESET)"
	@go test -v -cover -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✅ Coverage report: coverage.html$(RESET)"

# Code quality
fmt: ## Format code
	@echo "$(YELLOW)Formatting code...$(RESET)"
	@go fmt ./...

lint: ## Run linter
	@echo "$(YELLOW)Running linter...$(RESET)"
	@golangci-lint run || echo "$(YELLOW)⚠️  golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest$(RESET)"

vet: ## Run go vet
	@echo "$(YELLOW)Running go vet...$(RESET)"
	@go vet ./...

# Dependencies
deps: ## Download dependencies
	@echo "$(YELLOW)Downloading dependencies...$(RESET)"
	@go mod download
	@go mod tidy

deps-update: ## Update dependencies
	@echo "$(YELLOW)Updating dependencies...$(RESET)"
	@go get -u ./...
	@go mod tidy

# Installation
install: build ## Install binary to GOPATH/bin
	@echo "$(YELLOW)Installing $(BINARY_NAME)...$(RESET)"
	@go install $(GO_BUILD_FLAGS) .
	@echo "$(GREEN)✅ Installed $(BINARY_NAME) to $(shell go env GOPATH)/bin$(RESET)"

install-local: build ## Install binary to /usr/local/bin (requires sudo)
	@echo "$(YELLOW)Installing $(BINARY_NAME) to /usr/local/bin...$(RESET)"
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "$(GREEN)✅ Installed $(BINARY_NAME) to /usr/local/bin$(RESET)"

uninstall: ## Remove installed binary
	@echo "$(YELLOW)Removing installed $(BINARY_NAME)...$(RESET)"
	@rm -f $(shell go env GOPATH)/bin/$(BINARY_NAME)
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "$(GREEN)✅ Uninstalled $(BINARY_NAME)$(RESET)"

# Cleaning
clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(RESET)"
	@rm -rf $(BUILD_DIR) $(DIST_DIR)
	@rm -f $(BINARY_NAME) coverage.out coverage.html
	@echo "$(GREEN)✅ Cleaned$(RESET)"

# Information
version: ## Show version information
	@echo "$(BOLD)Sprout CLI Version Information$(RESET)"
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Commit: $(COMMIT)"
	@echo "Go Version: $(shell go version)"
	@echo "Platform: $(GOOS)/$(GOARCH)"

size: build ## Show binary size
	@echo "$(BOLD)Binary Size Information$(RESET)"
	@ls -lh $(BUILD_DIR)/$(BINARY_NAME) | awk '{print "Size: " $$5}'
	@file $(BUILD_DIR)/$(BINARY_NAME)

# Development helpers
run: build ## Build and run with sample args
	@echo "$(YELLOW)Running $(BINARY_NAME) weather...$(RESET)"
	@./$(BUILD_DIR)/$(BINARY_NAME) weather

run-onboard: build ## Build and run AI onboarding demo
	@echo "$(YELLOW)Running AI onboarding demo...$(RESET)"
	@./$(BUILD_DIR)/$(BINARY_NAME) weather --onboard-ai

# Release helpers
package: release ## Create distribution packages
	@echo "$(BOLD)Creating distribution packages...$(RESET)"
	@mkdir -p $(DIST_DIR)/packages
	@cd $(DIST_DIR) && tar czf packages/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz $(BINARY_NAME)-linux-amd64
	@cd $(DIST_DIR) && tar czf packages/$(BINARY_NAME)-$(VERSION)-linux-arm64.tar.gz $(BINARY_NAME)-linux-arm64
	@cd $(DIST_DIR) && tar czf packages/$(BINARY_NAME)-$(VERSION)-darwin-amd64.tar.gz $(BINARY_NAME)-darwin-amd64
	@cd $(DIST_DIR) && tar czf packages/$(BINARY_NAME)-$(VERSION)-darwin-arm64.tar.gz $(BINARY_NAME)-darwin-arm64
	@cd $(DIST_DIR) && zip packages/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BINARY_NAME)-windows-amd64.exe
	@echo "$(GREEN)✅ Distribution packages created in $(DIST_DIR)/packages/$(RESET)"

# Quick development workflow
dev: ## Quick development build and test
	@$(MAKE) fmt
	@$(MAKE) build
	@$(MAKE) test
	@echo "$(GREEN)✅ Development cycle complete$(RESET)"