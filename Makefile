# Makefile for MCP Workshop

.PHONY: build install clean workshop-build workshop-install

# Build the main MCP server
build:
	go build -o bin/mcp-server .

# Build the workshop tool
workshop-build:
	go build -o bin/workshop ./cmd/workshop

# Install the workshop tool globally
workshop-install:
	go install ./cmd/workshop

# Build both binaries
build-all: build workshop-build

# Install both binaries
install-all: workshop-install

# Clean build artifacts
clean:
	rm -rf bin/

# Run the MCP server
run:
	go run .

# Run tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Run the workshop tool locally (without installing)
workshop:
	go run ./cmd/workshop $(ARGS)

# Examples:
# make workshop ARGS="next"
# make workshop ARGS="status"
# make workshop ARGS="jump 002-mcp-weather-definition"
