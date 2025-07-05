# Makefile for spotctl

# Variables
BINARY_NAME=spotctl
VERSION?=dev
COMMIT?=$(shell git rev-parse --short HEAD)
DATE?=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS=-ldflags "-X github.com/georgetaylor/spotctl/pkg/version.Version=${VERSION} -X github.com/georgetaylor/spotctl/pkg/version.Commit=${COMMIT} -X github.com/georgetaylor/spotctl/pkg/version.Date=${DATE}"

# Coverage thresholds
COVERAGE_THRESHOLD := 80
CRITICAL_THRESHOLD := 90

# Default target
.PHONY: help
help: ## Show this help message
	@echo 'Usage: make <target>'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build the binary
	@echo "Building ${BINARY_NAME}..."
	go build ${LDFLAGS} -o bin/${BINARY_NAME} .

.PHONY: build-all
build-all: ## Build binaries for all platforms
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-darwin-arm64 .
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-windows-amd64.exe .

.PHONY: install
install: ## Install the binary
	@echo "Installing ${BINARY_NAME}..."
	go install ${LDFLAGS} .

.PHONY: clean
clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	go clean

.PHONY: test
test: ## Run tests
	@echo "Running tests..."
	go test -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint: ## Run linter
	@echo "Running linter..."
	go vet ./...
	gofmt -l .

.PHONY: fmt
fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

.PHONY: mod
mod: ## Download and tidy dependencies
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

.PHONY: run
run: ## Run the application
	@echo "Running ${BINARY_NAME}..."
	go run ${LDFLAGS} . $(ARGS)

.PHONY: dev
dev: fmt lint test build ## Run development workflow (format, lint, test, build)

.PHONY: release
release: clean mod test build-all ## Prepare a release (clean, mod, test, build-all)

.PHONY: coverage
coverage:
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@echo "Generating coverage reports..."
	@# Generate current coverage report
	@echo "# Test Coverage Report\n" > coverage.md
	@echo "Generated on: $$(date)\n" >> coverage.md
	@echo "## Coverage Summary\n" >> coverage.md
	@echo "\`\`\`" >> coverage.md
	@go tool cover -func=coverage.out >> coverage.md
	@echo "\`\`\`\n" >> coverage.md
	@echo "## Coverage Status\n" >> coverage.md
	@echo "| Package | Coverage | Status |" >> coverage.md
	@echo "|---------|----------|--------|" >> coverage.md
	@go tool cover -func=coverage.out | grep -v "total:" | awk -v threshold=$(COVERAGE_THRESHOLD) '{ \
		coverage = $$3; gsub("%", "", coverage); coverageNum = coverage + 0; \
		if (coverageNum == 100) status = "✅"; \
		else if (coverageNum >= threshold) status = "🟡"; \
		else status = "❌"; \
		printf "| %s | %s | %s |\n", $$1, $$3, status \
	}' >> coverage.md
	@echo "\n## Critical Paths Coverage\n" >> coverage.md
	@echo "| File | Coverage | Status |" >> coverage.md
	@echo "|------|----------|--------|" >> coverage.md
	@go tool cover -func=coverage.out | grep -E "client|config|auth" | awk -v threshold=$(COVERAGE_THRESHOLD) '{ \
		coverage = $$3; gsub("%", "", coverage); coverageNum = coverage + 0; \
		if (coverageNum == 100) status = "✅"; \
		else if (coverageNum >= threshold) status = "🟡"; \
		else status = "❌"; \
		printf "| %s | %s | %s |\n", $$1, $$3, status \
	}' >> coverage.md
	@echo "\n## Untested Files\n" >> coverage.md
	@echo "\`\`\`" >> coverage.md
	@go tool cover -func=coverage.out | grep "0.0%" >> coverage.md || true
	@echo "\`\`\`\n" >> coverage.md
	@echo "## Notes\n" >> coverage.md
	@echo "- ✅ 100% coverage" >> coverage.md
	@echo "- 🟡 Meets threshold ($(COVERAGE_THRESHOLD)%) but not perfect" >> coverage.md
	@echo "- ❌ Below threshold ($(COVERAGE_THRESHOLD)%)" >> coverage.md
	@echo "\nHTML report available at: coverage.html"
	@go tool cover -html=coverage.out -o coverage.html

	@# Update coverage history
	@if [ ! -f coverage_history.md ]; then \
		echo "# Coverage History\n" > coverage_history.md; \
		echo "| Date | Coverage | Change | Notes |" >> coverage_history.md; \
		echo "|------|----------|--------|-------|" >> coverage_history.md; \
	fi
	@CURRENT_COVERAGE=$$(go tool cover -func=coverage.out | grep total | awk '{print $$3}'); \
	LAST_COVERAGE=$$(tail -n 1 coverage_history.md 2>/dev/null | awk -F'|' '{print $$3}' | tr -d ' ' || echo "0.0%"); \
	CHANGE=$$(echo "$$CURRENT_COVERAGE $$LAST_COVERAGE" | awk '{ \
		sub(/%/, "", $$1); sub(/%/, "", $$2); \
		diff = $$1 - $$2; \
		if (diff > 0) printf "+%.1f%%", diff; \
		else if (diff < 0) printf "%.1f%%", diff; \
		else printf "0.0%%"; \
	}'); \
	echo "| $$(date +%Y-%m-%d) | $$CURRENT_COVERAGE | $$CHANGE | $$(if [ "$$CHANGE" != "0.0%" ]; then echo "Coverage $$CHANGE"; else echo "No change"; fi) |" >> coverage_history.md

	@echo "Coverage reports generated: coverage.md and coverage.html"
	@echo "Coverage history updated in: coverage_history.md"
