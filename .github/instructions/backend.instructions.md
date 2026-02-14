---
applyTo: "**/*.go"
description: "Go backend development guidelines with context engineering"
---
# Go Backend Development Guidelines

## Context Loading
Review [module dependencies](../../go.mod) and
[application structure](../../main.go) before starting.

## Deterministic Requirements
- Follow Effective Go conventions and idioms
- Use interfaces to define behavior contracts
- Implement resource cleanup with `defer`
- Use `context.Context` for request scoping and cancellation
- Implement structured logging with appropriate levels
- Use proper HTTP status codes and error responses
- RESTful API design
- gofmt

## Structured Output
Generate code with:
- [ ] Wrapped errors with context
- [ ] Table-driven unit tests
- [ ] Package/module documentation
- [ ] Integration tests for API endpoints
- [ ] Graceful shutdown handling
