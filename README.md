# Translator Service

A monolithic Go web service that provides English translation using Large Language Models.

## Project Structure

Refer to [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) for details on the directory layout.

## Getting Started

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Start the service with `go run cmd/translator/main.go`

The service will start on http://localhost:8080

## Development

This project includes a Makefile with common development tasks:

- `make format` - Format code with goimports and gofmt
- `make lint` - Run golangci-lint for static analysis
- `make build` - Build the application
- `make run` - Run the application
- `make clean` - Clean build artifacts
- `make install-tools` - Install development tools