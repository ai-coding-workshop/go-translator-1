# Translator Service

A monolithic Go web service that provides English translation using Large Language Models.

## Project Structure

Refer to [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) for details on the directory layout.

## Getting Started

1. Clone the repository
2. Run `go mod tidy` to install dependencies
3. Start the service with `go run cmd/translator/main.go`

The service will start on http://localhost:8080

### Configuration

The service can be configured using either environment variables or a YAML configuration file.

#### Environment Variables
- `PORT` - Server port (default: "8080")
- `OPENAI_ENDPOINT` - OpenAI API endpoint (default: "https://api.openai.com/v1")
- `OPENAI_API_KEY` - OpenAI API key for GPT models
- `ANTHROPIC_ENDPOINT` - Anthropic API endpoint (default: "https://api.anthropic.com/v1")
- `ANTHROPIC_API_KEY` - Anthropic API key for Claude models
- `DEBUG` - Enable debug mode (default: false)
- `TIMEOUT` - Request timeout in seconds (default: 30)

#### YAML Configuration File
You can also use a YAML configuration file:

```bash
go run cmd/translator/main.go -config=config.yaml
```

Example config.yaml:
```yaml
# Translation Service Configuration
server:
  port: "8080"

llm:
  openai_endpoint: "https://idealab.alibaba-inc.com/api/openai/v1"
  openai_key: "your-openai-key"
  anthropic_endpoint: "https://api.anthropic.com/v1"
  anthropic_key: "your-anthropic-key"
  timeout: 30

debug: false
```

Environment variables will override values from the configuration file.

## Development

This project includes a Makefile with common development tasks:

- `make format` - Format code with goimports and gofmt
- `make lint` - Run golangci-lint for static analysis
- `make build` - Build the application
- `make run` - Run the application
- `make clean` - Clean build artifacts
- `make install-tools` - Install development tools

### Pre-commit Hooks

This project uses pre-commit hooks to ensure code quality. Install and set up pre-commit:

1. Install pre-commit: `pip install pre-commit` or `brew install pre-commit`
2. Install the hooks: `pre-commit install`
3. The hooks will now run automatically on each commit

The pre-commit configuration includes:
- Trailing whitespace removal
- End-of-file fixer
- YAML validation
- Large file checks
- Go formatting (gofmt)
- Go imports formatting (goimports)
- Go linting (golangci-lint)
- Go unit tests
