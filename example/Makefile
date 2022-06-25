.DEFAULT_GOAL := help

.PHONY: help
help: ## Display this help message
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  build        Build the application"
	@echo "  run          Run the application"
	@echo "  test         Run the tests"
	@echo "  clean        Clean up the build artifacts"

.PHONY: build
build: ## Build the application
	go build -o myapp cmd/myapp/main.go

.PHONY: run
run: ## Run the application
	go run cmd/myapp/main.go

.PHONY: test
test: ## Run the tests
	go test ./...

.PHONY: clean
clean: ## Clean up the build artifacts
	rm -f myapp