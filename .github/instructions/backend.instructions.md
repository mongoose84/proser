---
applyTo: "**/*.go"
description: "Backend development guidelines for server-side languages"
---
# Backend Development Guidelines

## Technology Stack
- **Language**: Go

## Context Loading
Review [Go module dependencies](../../go.mod) and 
[main application structure](../../main.go) before starting.

## Go-Specific Guidelines
- Follow Go conventions and idioms (effective Go)
- Package names should be lowercase, single words
- Use receiver names consistently
- Prefer composition over inheritance
- Use interfaces to define behavior contracts
- Implement proper resource cleanup with defer

## Deterministic Requirements
- Implement proper error handling and logging
- Use structured logging with appropriate log levels
- Apply dependency injection patterns
- Implement proper HTTP status codes and responses
- Use context.Context for request scoping and cancellation

## Project Code Style
Standard formatting for Go

## Structured Output
Generate code with:
- [ ] Comprehensive error handling with wrapped errors
- [ ] Unit tests with table-driven test patterns
- [ ] Benchmark tests for performance-critical code
- [ ] Proper package/module documentation and examples
- [ ] Integration tests for API endpoints
- [ ] Logging with structured fields
- [ ] Graceful shutdown handling for services
