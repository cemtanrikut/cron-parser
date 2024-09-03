# Cron Expression Parser project

Cron Parser is a Go program that parses and formats cron expressions. This tool helps you understand the details of a given cron expression and presents the output in a readable format.

## Features

- Parses a cron expression.
- Formats and displays parsed fields in a readable format.

## Requirements

- Go (1.17 or newer)

## Installation

1. Install the necessary dependencies for the project:
   ```sh
   go mod tidy
   ```

2. Build the project:
   ```sh
   go build -o cron_parser
   ```

## Usage

To run the program, use the following command in your terminal:
 ```sh
   ./cron_parser "<cron_expression>"
   ```

For example:
 ```sh
   ./cron_parser "0 12 * * 1-5"
   ```

This command will parse the "0 12 * * 1-5" cron expression and print the results to the screen.

# Contributing

If you would like to contribute, please follow these steps:

1. Fork the repository.
2. Make your changes and create a pull request.

## Lisence

This project is licensed under the MIT License.

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
