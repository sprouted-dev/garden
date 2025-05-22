# Garden Monorepo Makefile
# Sprouted Ecosystem Build Automation

.PHONY: help build test clean install dev release apps libs tools
.DEFAULT_GOAL := help

# Colors for output
BOLD := \033[1m
GREEN := \033[32m
YELLOW := \033[33m
RESET := \033[0m

help: ## Show this help message
	@echo "$(BOLD)Garden Monorepo Build System$(RESET)"
	@echo "$(YELLOW)Available targets:$(RESET)"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  $(GREEN)%-15s$(RESET) %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development Commands
dev: ## Start development environment
	@echo "$(BOLD)Starting Garden development environment...$(RESET)"
	@$(MAKE) build
	@echo "$(GREEN)✅ Development environment ready$(RESET)"

build: apps libs ## Build all components
	@echo "$(BOLD)Building Garden ecosystem...$(RESET)"

apps: ## Build all applications
	@echo "$(YELLOW)Building applications...$(RESET)"
	@$(MAKE) -C apps/sprout-cli build

libs: ## Build all libraries  
	@echo "$(YELLOW)Building libraries...$(RESET)"
	@cd libs/weather && go build ./...

tools: ## Build development tools
	@echo "$(YELLOW)Building tools...$(RESET)"
	@echo "No tools to build yet"

# Testing
test: test-libs test-apps ## Run all tests
	@echo "$(GREEN)✅ All tests passed$(RESET)"

test-libs: ## Test libraries
	@echo "$(YELLOW)Testing libraries...$(RESET)"
	@cd libs/weather && go test ./...

test-apps: ## Test applications
	@echo "$(YELLOW)Testing applications...$(RESET)"
	@$(MAKE) -C apps/sprout-cli test

# Quality & Linting
lint: ## Run linters
	@echo "$(YELLOW)Running linters...$(RESET)"
	@cd libs/weather && golangci-lint run || echo "golangci-lint not installed, skipping"
	@$(MAKE) -C apps/sprout-cli lint

fmt: ## Format code
	@echo "$(YELLOW)Formatting code...$(RESET)"
	@cd libs/weather && go fmt ./...
	@$(MAKE) -C apps/sprout-cli fmt

# Cleaning
clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(RESET)"
	@$(MAKE) -C apps/sprout-cli clean
	@find . -name "*.log" -delete
	@find . -name ".DS_Store" -delete

# Installation
install: build ## Install tools locally
	@echo "$(YELLOW)Installing Garden tools...$(RESET)"
	@$(MAKE) -C apps/sprout-cli install

# Release Management  
release: ## Build release binaries for all platforms
	@echo "$(BOLD)Building release binaries...$(RESET)"
	@$(MAKE) -C apps/sprout-cli release

release-check: ## Check if ready for release
	@echo "$(YELLOW)Checking release readiness...$(RESET)"
	@$(MAKE) test
	@$(MAKE) lint
	@echo "$(GREEN)✅ Ready for release$(RESET)"

# Documentation
docs: ## Generate documentation
	@echo "$(YELLOW)Generating documentation...$(RESET)"
	@echo "Documentation system coming soon..."

# Weather System Specific
weather-build: ## Build weather system components
	@cd libs/weather && go build ./...
	@$(MAKE) -C apps/sprout-cli build

weather-test: ## Test weather system
	@cd libs/weather && go test -v ./...

weather-demo: install ## Demo weather system
	@echo "$(BOLD)Weather System Demo$(RESET)"
	@echo "Basic weather context:"
	@./apps/sprout-cli/sprout weather || echo "Run 'make install' first"
	@echo ""
	@echo "AI onboarding context:"
	@./apps/sprout-cli/sprout weather --onboard-ai || echo "Run 'make install' first"

# Development Helpers
version: ## Show version information
	@echo "$(BOLD)Garden Ecosystem Version Information$(RESET)"
	@echo "Weather System: $(shell cd libs/weather && go list -f '{{.Version}}' . 2>/dev/null || echo 'dev')"
	@echo "Sprout CLI: $(shell cd apps/sprout-cli && go list -f '{{.Version}}' . 2>/dev/null || echo 'dev')"
	@echo "Git: $(shell git describe --tags --always --dirty 2>/dev/null || echo 'unknown')"

status: ## Show project status
	@echo "$(BOLD)Garden Project Status$(RESET)"
	@echo "🌦️  Weather System: $(shell ./apps/sprout-cli/sprout weather --for-ai >/dev/null 2>&1 && echo 'Active' || echo 'Not configured')"
	@echo "📁 Garden Path: $(PWD)"
	@echo "🌿 Git Branch: $(shell git branch --show-current 2>/dev/null || echo 'unknown')"
	@echo "📊 Git Status: $(shell git status --porcelain | wc -l | xargs echo) uncommitted changes"

# Setup for new developers
setup: ## Set up development environment for new contributors
	@echo "$(BOLD)Setting up Garden development environment...$(RESET)"
	@echo "$(YELLOW)Checking prerequisites...$(RESET)"
	@command -v go >/dev/null 2>&1 || { echo "❌ Go is required but not installed"; exit 1; }
	@command -v git >/dev/null 2>&1 || { echo "❌ Git is required but not installed"; exit 1; }
	@echo "✅ Go version: $(shell go version)"
	@echo "✅ Git version: $(shell git --version)"
	@echo "$(YELLOW)Installing dependencies...$(RESET)"
	@cd libs/weather && go mod tidy
	@cd apps/sprout-cli && go mod tidy
	@echo "$(YELLOW)Building tools...$(RESET)"
	@$(MAKE) build
	@echo "$(YELLOW)Installing weather hooks...$(RESET)"
	@./apps/sprout-cli/sprout weather --install-hooks || echo "Weather hooks will be installed on first run"
	@echo "$(GREEN)✅ Development environment ready!$(RESET)"
	@echo "$(BOLD)Next steps:$(RESET)"
	@echo "  • Run '$(YELLOW)make weather-demo$(RESET)' to see the weather system in action"
	@echo "  • Run '$(YELLOW)make test$(RESET)' to run all tests"
	@echo "  • Run '$(YELLOW)make help$(RESET)' to see all available commands"