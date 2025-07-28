# Translation Service User Guide

This document provides a comprehensive guide for using and understanding the Translation Service.

## Table of Contents
- [Overview](#overview)
- [Features](#features)
- [System Requirements](#system-requirements)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
  - [Web Interface](#web-interface)
  - [REST API](#rest-api)
- [Supported Models](#supported-models)
- [Error Handling](#error-handling)
- [Performance](#performance)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)

## Overview

The Translation Service is a monolithic Go web application that provides English translation capabilities using various Large Language Models (LLMs) from different providers. It offers both a user-friendly web interface and a powerful REST API for programmatic access.

## Features

- **Multiple LLM Support** - Translate using GPT models from OpenAI or Claude models from Anthropic
- **Dual Interface** - Web interface for human users and REST API for programmatic access
- **Robust Error Handling** - Comprehensive error handling with user-friendly messages
- **Configurable** - Flexible configuration through environment variables, YAML files, or command-line flags
- **Secure** - Secure handling of API keys and credentials
- **Containerized** - Docker support for easy deployment
- **Well-Tested** - Comprehensive test suite with unit and integration tests

## System Requirements

- **Go 1.21 or higher** (for building from source)
- **Docker** (optional, for containerized deployment)
- **API Keys** from one or more LLM providers:
  - OpenAI API key for GPT models
  - Anthropic API key for Claude models

## Installation

### Option 1: From Source

1. Clone the repository:
```bash
git clone <repository-url>
cd translator-service
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
make build
```

### Option 2: Using Docker

1. Build the Docker image:
```bash
docker build -t translator-service .
```

2. Or use Docker Compose:
```bash
docker-compose up -d
```

## Configuration

The service can be configured through multiple methods with the following precedence:
1. Command-line flags
2. Environment variables
3. Configuration file

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | 8080 |
| `OPENAI_ENDPOINT` | OpenAI API endpoint | https://api.openai.com/v1 |
| `OPENAI_API_KEY` | OpenAI API key | (required) |
| `ANTHROPIC_ENDPOINT` | Anthropic API endpoint | https://api.anthropic.com/v1 |
| `ANTHROPIC_API_KEY` | Anthropic API key | (required) |
| `DEBUG` | Enable debug mode | false |
| `TIMEOUT` | Request timeout in seconds | 30 |

### Configuration File

Create a `config.yaml` file:

```yaml
server:
  port: "8080"

llm:
  openai_endpoint: "https://api.openai.com/v1"
  openai_key: "your-openai-key"
  anthropic_endpoint: "https://api.anthropic.com/v1"
  anthropic_key: "your-anthropic-key"
  timeout: 30

debug: false
```

### Command-line Flags

| Flag | Description |
|------|-------------|
| `-config` | Path to configuration file |

## Usage

### Web Interface

1. Start the service:
```bash
go run cmd/translator/main.go
```

2. Open your browser to http://localhost:8080

3. Enter text to translate and select a model from the dropdown

4. Click "Translate" to get the translation

### REST API

The service provides a REST API for programmatic access:

#### Endpoint: POST /api/translate

**Request:**
```json
{
  "text": "string",
  "model": "string"
}
```

**Response (Success):**
```json
{
  "original": "string",
  "translation": "string",
  "model": "string"
}
```

**Response (Error):**
```json
{
  "error": true,
  "message": "string",
  "details": "string"
}
```

**Example:**
```bash
curl -X POST http://localhost:8080/api/translate \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Hello, world!",
    "model": "gpt-3.5"
  }'
```

## Supported Models

The service supports the following models:

| Provider | Model ID | Description |
|----------|----------|-------------|
| OpenAI | `gpt-4` | GPT-4 model |
| OpenAI | `gpt-3.5` | GPT-3.5 model |
| Anthropic | `claude-3-opus` | Claude 3 Opus model |
| Anthropic | `claude-3-sonnet` | Claude 3 Sonnet model |
| Anthropic | `claude-3-haiku` | Claude 3 Haiku model |

## Error Handling

The service provides comprehensive error handling with appropriate HTTP status codes:

| Status Code | Description |
|-------------|-------------|
| 200 | Success |
| 400 | Bad Request (invalid input) |
| 405 | Method Not Allowed |
| 408 | Request Timeout |
| 503 | Service Unavailable |
| 500 | Internal Server Error |

## Performance

The service is designed for high performance and includes:

- Concurrent request handling
- Connection pooling for HTTP clients
- Proper timeout handling
- Efficient memory usage
- Built-in health check endpoint

For performance optimization details, see [OPTIMIZATION.md](OPTIMIZATION.md).

## Troubleshooting

### Common Issues

1. **Service fails to start**
   - Check that required environment variables are set
   - Verify API keys are valid
   - Check port availability

2. **Translations fail**
   - Verify API keys have proper permissions
   - Check network connectivity to LLM providers
   - Review timeout settings

3. **Slow responses**
   - Check LLM provider status
   - Review timeout settings
   - Consider caching strategies

### Logs

The service logs to stdout/stderr. In debug mode, additional information is logged.

### Health Check

Verify the service is running:
```bash
curl http://localhost:8080/health
```

## Contributing

We welcome contributions to the Translation Service. Please see our contributing guidelines for more information.

### Development Setup

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for your changes
5. Run the test suite
6. Submit a pull request

### Testing

Run all tests:
```bash
make test
```

Run benchmarks:
```bash
make bench
```

Format code:
```bash
make format
```

Run linters:
```bash
make lint
```
