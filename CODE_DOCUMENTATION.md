# Code Documentation

This document provides detailed documentation of the Translation Service codebase, including architecture, key components, and implementation details.

## Table of Contents
- [Architecture Overview](#architecture-overview)
- [Project Structure](#project-structure)
- [Core Components](#core-components)
  - [Models](#models)
  - [Configuration](#configuration)
  - [Services](#services)
  - [Handlers](#handlers)
- [Frontend Components](#frontend-components)
- [Testing](#testing)
- [Error Handling](#error-handling)

## Architecture Overview

The Translation Service follows a modular monolithic architecture with clear separation of concerns:

1. **Models** - Data structures used throughout the application
2. **Configuration** - Configuration management and loading
3. **Services** - Business logic including translation and validation
4. **Handlers** - HTTP request handling and routing
5. **Frontend** - HTML templates, CSS, and JavaScript

The service uses an adapter pattern for LLM integration, allowing support for multiple providers (OpenAI, Anthropic) through a common interface.

## Project Structure

```
translator-service/
├── cmd/
│   └── translator/
│       └── main.go          # Application entry point
├── internal/
│   ├── models/              # Data structures
│   ├── config/              # Configuration management
│   ├── services/            # Business logic
│   ├── handlers/            # HTTP handlers
│   └── testutils/           # Testing utilities
├── web/
│   ├── templates/           # HTML templates
│   └── static/              # CSS, JS, images
├── configs/                 # Configuration files
├── docs/                    # Documentation files
└── go.mod                   # Go module definition
```

## Core Components

### Models

The models package defines the core data structures used in the application.

**TranslationRequest** - Represents a translation request:
- `Text` (string) - The text to be translated
- `Model` (string) - The LLM model to use for translation

**TranslationResponse** - Represents a translation response:
- `Original` (string) - The original text
- `Translation` (string) - The translated text
- `Model` (string) - The model used for translation

### Configuration

The config package handles loading and managing application configuration from multiple sources:
- Command-line flags
- Environment variables
- YAML configuration files

The configuration follows a hierarchy where command-line flags override environment variables, which override configuration file values.

Key configuration options:
- Server settings (port)
- LLM settings (endpoints, API keys, timeout)
- Debug mode

### Services

The services package contains the core business logic.

#### TranslatorService

The main service that orchestrates translation requests:
- Validates input using ValidationService
- Routes requests to appropriate LLM translators
- Handles timeouts and errors
- Returns formatted responses

#### ValidationService

Handles input validation:
- Ensures text is provided and not empty
- Validates that the selected model is supported
- Sanitizes input data

#### LLM Translators

Implementation of the Translator interface for different LLM providers:
- **OpenAITranslator** - For OpenAI GPT models
- **AnthropicTranslator** - For Anthropic Claude models

Each translator implements:
- `Name()` - Returns the translator name
- `SupportsModel(model string)` - Checks if the translator supports a model
- `Translate(ctx context.Context, req *TranslationRequest)` - Performs the translation

### Handlers

The handlers package contains HTTP handlers for both web interface and REST API.

#### HomeHandler

Serves the main web page with the translation form.

#### TranslateHandler

Processes translation requests from the web form:
- Parses form data
- Validates input
- Calls TranslatorService
- Renders results in HTML template

#### APIHandler

REST API endpoint for programmatic access:
- Parses JSON requests
- Validates input
- Calls TranslatorService
- Returns JSON responses

## Frontend Components

### HTML Templates

Located in `web/templates/`:
- `home.html` - Main translation form
- `result.html` - Translation results display

### Static Assets

Located in `web/static/`:
- CSS styling in `css/style.css`
- JavaScript functionality in `js/main.js`

The frontend provides a responsive web interface with:
- Text input field
- Model selection dropdown
- Real-time validation
- Clear error messaging

## Testing

The service includes comprehensive tests:

### Unit Tests

- Config package tests
- Services tests (translation, validation)
- Handlers tests
- Test utilities and mocks

### Integration Tests

Full service integration tests covering:
- Configuration loading
- Translation service functionality
- HTTP endpoint testing

### Test Utilities

Reusable testing components in `internal/testutils/`:
- Mock implementations
- Helper functions for creating test requests
- Assertion utilities

## Error Handling

The service implements comprehensive error handling:

### Input Validation

- Client-side validation in JavaScript
- Server-side validation in ValidationService
- User-friendly error messages

### LLM API Errors

- Timeout handling with context
- Service unavailability detection
- Graceful degradation

### HTTP Error Responses

- Proper HTTP status codes
- Structured JSON error responses for API
- User-friendly HTML error pages for web interface
