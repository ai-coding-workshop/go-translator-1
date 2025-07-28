# Deployment Guide

This document provides instructions for deploying the Translation Service in various environments.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Local Development Setup](#local-development-setup)
- [Production Deployment](#production-deployment)
  - [Building the Application](#building-the-application)
  - [Running the Service](#running-the-service)
  - [Process Management](#process-management)
- [Docker Deployment](#docker-deployment)
- [Environment Configuration](#environment-configuration)
- [Security Considerations](#security-considerations)
- [Monitoring and Logging](#monitoring-and-logging)

## Prerequisites

- Go 1.21 or higher
- Git
- Make (for using Makefile commands)
- Access to LLM API keys (OpenAI, Anthropic)

## Local Development Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd translator-service
```

2. Install dependencies:
```bash
go mod tidy
```

3. Set up configuration:
```bash
cp config.yaml.example config.yaml
# Edit config.yaml with your API keys
```

4. Run the service:
```bash
go run cmd/translator/main.go
```

The service will be available at http://localhost:8080

## Production Deployment

### Building the Application

Use the Makefile to build the application:

```bash
make build
```

This will create a binary in the `bin/` directory.

For cross-compilation to different platforms:

```bash
# Build for Linux
GOOS=linux GOARCH=amd64 make build

# Build for Windows
GOOS=windows GOARCH=amd64 make build
```

### Running the Service

After building, you can run the service directly:

```bash
./bin/translator
```

Or with a specific configuration file:

```bash
./bin/translator -config /path/to/config.yaml
```

### Process Management

For production environments, use a process manager like systemd, supervisor, or pm2.

Example systemd service file (`/etc/systemd/system/translator.service`):

```ini
[Unit]
Description=Translation Service
After=network.target

[Service]
Type=simple
User=translator
WorkingDirectory=/opt/translator-service
ExecStart=/opt/translator-service/bin/translator -config /opt/translator-service/config.yaml
Restart=always
RestartSec=10
Environment=PORT=8080

[Install]
WantedBy=multi-user.target
```

Enable and start the service:

```bash
sudo systemctl enable translator.service
sudo systemctl start translator.service
```

## Docker Deployment

The service includes a Dockerfile for containerized deployment.

### Building the Docker Image

```bash
docker build -t translator-service .
```

### Running with Docker

```bash
docker run -d \
  --name translator \
  -p 8080:8080 \
  -e OPENAI_API_KEY=your-openai-key \
  -e ANTHROPIC_API_KEY=your-anthropic-key \
  translator-service
```

### Docker Compose

The project includes a `docker-compose.yml` file for easier deployment:

```yaml
version: '3.8'

services:
  translator:
    build: .
    ports:
      - "8080:8080"
    environment:
      - PORT=8080
      # Uncomment and set your API keys when running in production
      # - OPENAI_API_KEY=your-openai-api-key
      # - ANTHROPIC_API_KEY=your-anthropic-api-key
    volumes:
      # Mount config file if you want to use a custom configuration
      # - ./config.yaml:/app/config.yaml
    restart: unless-stopped
```

Run with Docker Compose:

```bash
docker-compose up -d
```

## Environment Configuration

### Environment Variables

Set these environment variables in production:

```bash
# Server configuration
PORT=8080

# LLM API keys (keep these secure)
OPENAI_API_KEY=your-openai-api-key
ANTHROPIC_API_KEY=your-anthropic-api-key

# API endpoints (optional, use defaults if not specified)
OPENAI_ENDPOINT=https://api.openai.com/v1
ANTHROPIC_ENDPOINT=https://api.anthropic.com/v1

# Timeout settings
TIMEOUT=30

# Debug mode (disable in production)
DEBUG=false
```

### Configuration File

Alternatively, use a YAML configuration file:

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

## Security Considerations

### API Key Management

1. **Never commit API keys** to version control
2. Use environment variables or secure configuration management
3. Rotate API keys regularly
4. Monitor API usage for unusual activity

### HTTPS

In production, always serve the service over HTTPS. Use a reverse proxy like nginx or Apache for SSL termination:

Example nginx configuration:

```nginx
server {
    listen 443 ssl;
    server_name your-domain.com;

    ssl_certificate /path/to/certificate.crt;
    ssl_certificate_key /path/to/private.key;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### Rate Limiting

Implement rate limiting at the reverse proxy or load balancer level to prevent abuse.

### Input Validation

The service includes built-in input validation, but additional validation at the network level can provide extra security.

## Monitoring and Logging

### Application Logging

The service logs to stdout/stderr. Configure your process manager or container platform to capture and store these logs.

Log levels:
- INFO - Normal operation
- ERROR - Error conditions
- DEBUG - Debug information (only when DEBUG=true)

### Health Checks

The service includes a health check endpoint at `/health`:

```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "ok",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

### Metrics

For production deployments, consider adding application metrics using Prometheus or similar monitoring solutions.

Example Prometheus configuration:

```yaml
scrape_configs:
  - job_name: 'translator-service'
    static_configs:
      - targets: ['localhost:8080']
```

### Alerting

Set up alerts for:
- High error rates
- Service downtime
- High latency
- API quota limits
