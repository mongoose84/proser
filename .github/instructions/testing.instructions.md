---
applyTo: "**/test/**"
description: "Testing guidelines for all test files and directories"
---
# Testing Guidelines

## Technology Stack
- **Testing Framework**: Go testing
- **Testing Strategy**: unit tests
- **Backend Language**: Go

## Context Loading
Review [project structure](../../README.md) and 
[testing patterns](../../) before writing tests.

## Deterministic Requirements
- Follow the AAA pattern: Arrange, Act, Assert
- Write descriptive test names that explain what is being tested
- Mock external dependencies appropriately
- Ensure tests are deterministic and repeatable
- Test both happy paths and error conditions
- Use table-driven tests for multiple scenarios

## Go Testing Patterns
- Use `testing.T` for unit tests and `testing.B` for benchmarks
- Follow Go naming conventions for test functions
- Use `testify` or similar assertion libraries when helpful
- Implement setup and teardown in test functions

## Code Style in Tests
Standard formatting for Go
Apply the same style guidelines to test code.

## Test Organization
- Group related tests in the same file
- Use clear, descriptive test names
- Keep tests focused on single functionality
- Use test helpers for common setup

## Structured Output
Generate tests with:
- [ ] Table-driven test patterns where appropriate
- [ ] Benchmark tests for performance-critical code
- [ ] Clear test documentation and comments
- [ ] Proper setup and cleanup
- [ ] Edge case and error condition coverage
- [ ] Integration tests for complex workflows
- [ ] Mock implementations for external dependencies
