# Makefile for Cron Expression Parser project

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