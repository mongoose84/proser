---
applyTo: "**/*_test.go"
description: "Testing guidelines with context engineering"
---
# Testing Guidelines

## Context Loading
Review [project conventions](../../README.md) and
[existing Go testing tests](../../) before writing tests.

## Deterministic Requirements
- Follow the AAA pattern: Arrange, Act, Assert
- Write descriptive test names that explain the scenario
- Mock external dependencies — keep tests isolated
- Ensure tests are deterministic and repeatable
- Cover both happy paths and error conditions
- Use table-driven tests for multiple scenarios
- Use `testing.T` for unit tests, `testing.B` for benchmarks
- unit tests
- gofmt

## Structured Output
Generate tests with:
- [ ] Table-driven test patterns
- [ ] Benchmark tests for performance-critical code
- [ ] Setup and teardown for shared state
- [ ] Edge case and error condition coverage
- [ ] Mock implementations for external dependencies
- [ ] Clear test documentation
