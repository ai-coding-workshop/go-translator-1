# Configuration Guide

This document explains how to configure the Translation Service.

## Configuration Sources

The service can be configured through three sources with the following precedence (highest to lowest):

1. Command-line flags
2. Environment variables
3. Configuration file (YAML)

## Configuration File

The service looks for a `config.yaml` file in the current directory. You can specify a different location using the `-config` command-line flag.

Example configuration file:

```yaml
# Translation Service Configuration
server:
  port: "8088"

llm:
  openai_endpoint: "https://api.openai.com/v1"
  openai_key: "your-openai-api-key"
  anthropic_endpoint: "https://api.anthropic.com/v1"
  anthropic_key: "your-anthropic-api-key"
  timeout: 30

debug: false
```

## Environment Variables

All configuration options can be set using environment variables:

| Environment Variable | Corresponding Config Field | Description |
|---------------------|----------------------------|-------------|
| PORT | server.port | Server port to listen on |
| OPENAI_ENDPOINT | llm.openai_endpoint | OpenAI API endpoint |
| OPENAI_API_KEY | llm.openai_key | OpenAI API key |
| ANTHROPIC_ENDPOINT | llm.anthropic_endpoint | Anthropic API endpoint |
| ANTHROPIC_API_KEY | llm.anthropic_key | Anthropic API key |
| DEBUG | debug | Enable debug mode (true/false) |
| TIMEOUT | llm.timeout | API timeout in seconds |

## Security Best Practices

1. **Never commit API keys** to version control. Use environment variables or a separate config file that is not committed.
2. **Use environment variables** for sensitive information in production environments.
3. **Validate configuration** at startup to ensure all required values are present.

## Command-line Flags

| Flag | Description |
|------|-------------|
| -config | Path to configuration file |

Example usage:
```bash
./translator -config /path/to/config.yaml
```
