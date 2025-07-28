# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Project Overview

This is a monolithic Go web service that provides English translation using Large Language Models. The service allows users to input English words or sentences and get translations from multiple selectable LLMs.

## Codebase Structure

The project follows a standard Go project layout:

- `cmd/translator/` - Main application entry point
- `internal/` - Private application code:
  - `handlers/` - HTTP handlers
  - `services/` - Business logic
  - `models/` - Data structures
  - `config/` - Configuration
- `web/` - Frontend assets:
  - `static/` - CSS, JS, images
  - `templates/` - HTML templates
- `configs/` - Configuration files

## Common Development Commands

### Building and Running
- `make build` - Build the application to `bin/translator`
- `make run` - Run the application directly
- `go run cmd/translator/main.go` - Run without building

### Code Quality
- `make format` - Format code with goimports and gofmt
- `make lint` - Run golangci-lint for static analysis
- `make install-tools` - Install development tools

### Pre-commit Hooks
- Install with `pre-commit install`
- Runs automatically on each commit
- Includes formatting, linting, and validation checks

### Dependencies
- `go mod tidy` - Sync dependencies

## Architecture

The application follows a modular architecture:
1. HTTP handlers in `internal/handlers/` process incoming requests
2. Business logic resides in `internal/services/`
3. Data structures are defined in `internal/models/`
4. Frontend templates and static assets are in `web/`

## Development Workflow

1. Make changes to the code
2. Format with `make format`
3. Lint with `make lint`
4. Build and test with `make build` and `make run`
