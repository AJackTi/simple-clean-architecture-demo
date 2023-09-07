tidy:
	go mod tidy

# ==============================================================================
# Linters https://golangci-lint.run/usage/install/
run-linter:
	@echo Starting linters
	golangci-lint run ./...