# Makefile for Cron Expression Parser project

# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOFMT = $(GOCMD) fmt
BINARY_NAME = cron_parser
BINARY_UNIX = $(BINARY_NAME)_unix

# Default target to run all build-related tasks
all: fmt test build

# Build the binary for the current OS
build: 
	$(GOBUILD) -o ./bin/$(BINARY_NAME) ./cmd/cron_parser

# Run tests
test: 
	$(GOTEST) -v ./...

# Format code
fmt:
	$(GOFMT) ./pkg/cron/*.go ./cmd/cron_parser/*.go ./tests/*.go

# Clean up binaries and cached files
clean: 
	$(GOCLEAN)
	rm -f ./bin/$(BINARY_NAME)
	rm -f ./bin/$(BINARY_UNIX)

# Build binary for Unix-like systems
build-unix: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/$(BINARY_UNIX) ./cmd/cron_parser

# Run the application with an example cron expression
run:
	./bin/$(BINARY_NAME) "*/15 0 1,15 * 1-5 /usr/bin/find"

# Install dependencies
mod-tidy:
	$(GOCMD) mod tidy

# Update dependencies
mod-download:
	$(GOCMD) mod download

# Help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all            Run fmt, test and build targets"
	@echo "  build          Build the binary for the current OS"
	@echo "  build-unix     Build the binary for Unix-like systems"
	@echo "  test           Run tests"
	@echo "  fmt            Format the code using go fmt"
	@echo "  clean          Remove binaries and cached files"
	@echo "  run            Run the application with an example input"
	@echo "  mod-tidy       Clean up the go.mod and go.sum files"
	@echo "  mod-download   Download the dependencies"
	@echo "  help           Display this help message"