package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func createFrontendInstructions(instructionsDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("applyTo: \"**/*.{jsx,tsx,css,js,ts,html}\"\n")
	sb.WriteString("description: \"Frontend development guidelines for UI components\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Frontend Development Guidelines\n\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review [project structure](../../README.md) and \n")
	sb.WriteString("[component patterns](../src/) before starting.\n\n")

	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Use consistent component structure and naming conventions\n")
	sb.WriteString("- Implement proper state management patterns\n")
	sb.WriteString("- Apply responsive design principles\n")
	sb.WriteString("- Ensure accessibility (WCAG guidelines)\n")
	sb.WriteString("- Use semantic HTML elements\n")

	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		if lang == "javascript" || lang == "typescript" || lang == "react" || lang == "js" || lang == "ts" {
			sb.WriteString("- Follow modern JavaScript/TypeScript best practices\n")
			sb.WriteString("- Use ES6+ features appropriately\n")
		} else if lang == "python" {
			sb.WriteString("- Follow Python web framework conventions (Flask/Django)\n")
			sb.WriteString("- Use Jinja2 templating best practices\n")
		} else if lang == "java" {
			sb.WriteString("- Follow Java web framework conventions (Spring/JSF)\n")
			sb.WriteString("- Use proper MVC patterns\n")
		} else {
			sb.WriteString(fmt.Sprintf("- Follow %s frontend framework best practices\n", config.Language))
		}
	}
	sb.WriteString("\n")

	if config.CodeStyle != "" {
		sb.WriteString("## Project Code Style\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	// Dynamic component guidelines based on language
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		if lang == "javascript" || lang == "typescript" || lang == "react" || lang == "js" || lang == "ts" {
			sb.WriteString("## React/JavaScript Guidelines\n")
			sb.WriteString("- Prefer functional components over class components\n")
			sb.WriteString("- Use hooks appropriately (useState, useEffect, custom hooks)\n")
			sb.WriteString("- Implement proper prop validation\n")
		} else {
			sb.WriteString("## Component Guidelines\n")
			sb.WriteString(fmt.Sprintf("- Follow %s component patterns and conventions\n", config.Language))
			sb.WriteString("- Maintain clear separation of concerns\n")
		}
	} else {
		sb.WriteString("## Component Guidelines\n")
		sb.WriteString("- Keep components small and focused\n")
		sb.WriteString("- Maintain clear separation of concerns\n")
	}
	sb.WriteString("- Use proper event handling and cleanup\n")
	sb.WriteString("- Implement error boundaries where appropriate\n\n")

	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate code with:\n")
	sb.WriteString("- [ ] Component documentation with usage examples\n")
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		if lang == "javascript" || lang == "typescript" || lang == "react" || lang == "js" || lang == "ts" {
			sb.WriteString("- [ ] Proper TypeScript/JSDoc annotations\n")
			sb.WriteString("- [ ] Unit tests with Jest/React Testing Library\n")
		} else if lang == "python" {
			sb.WriteString("- [ ] Python docstrings and type hints\n")
			sb.WriteString("- [ ] Unit tests with pytest or unittest\n")
		} else {
			sb.WriteString("- [ ] Proper documentation and type annotations\n")
			sb.WriteString("- [ ] Unit tests with appropriate testing framework\n")
		}
	} else {
		sb.WriteString("- [ ] Proper documentation and annotations\n")
		sb.WriteString("- [ ] Unit tests for components\n")
	}
	sb.WriteString("- [ ] Accessibility attributes (aria-labels, roles, etc.)\n")
	sb.WriteString("- [ ] Loading and error states for async operations\n")

	filePath := filepath.Join(instructionsDir, "frontend.instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func createBackendInstructions(instructionsDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("applyTo: \"**/*.{go,py,java,rs,js,ts}\"\n")
	sb.WriteString("description: \"Backend development guidelines for server-side languages\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Backend Development Guidelines\n\n")

	sb.WriteString("## Context Loading\n")
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		switch lang {
		case "go":
			sb.WriteString("Review [Go module dependencies](../../go.mod) and \n")
			sb.WriteString("[main application structure](../../main.go) before starting.\n\n")
		case "python":
			sb.WriteString("Review [Python dependencies](../../requirements.txt) and \n")
			sb.WriteString("[main application structure](../../) before starting.\n\n")
		case "java":
			sb.WriteString("Review [Maven/Gradle dependencies](../../pom.xml) and \n")
			sb.WriteString("[main application structure](../../src/main/java/) before starting.\n\n")
		case "javascript", "typescript", "node", "js", "ts":
			sb.WriteString("Review [Node.js dependencies](../../package.json) and \n")
			sb.WriteString("[main application structure](../../) before starting.\n\n")
		default:
			sb.WriteString("Review [project dependencies](../../) and \n")
			sb.WriteString("[main application structure](../../) before starting.\n\n")
		}
	} else {
		sb.WriteString("Review [project dependencies](../../) and \n")
		sb.WriteString("[main application structure](../../) before starting.\n\n")
	}

	// Language-specific guidelines
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		switch lang {
		case "go":
			sb.WriteString("## Go-Specific Guidelines\n")
			sb.WriteString("- Follow Go conventions and idioms (effective Go)\n")
			sb.WriteString("- Package names should be lowercase, single words\n")
			sb.WriteString("- Use receiver names consistently\n")
			sb.WriteString("- Prefer composition over inheritance\n")
			sb.WriteString("- Use interfaces to define behavior contracts\n")
			sb.WriteString("- Implement proper resource cleanup with defer\n\n")
		case "python":
			sb.WriteString("## Python-Specific Guidelines\n")
			sb.WriteString("- Follow PEP 8 style guidelines\n")
			sb.WriteString("- Use type hints for function signatures\n")
			sb.WriteString("- Follow naming conventions (snake_case)\n")
			sb.WriteString("- Use context managers for resource management\n")
			sb.WriteString("- Prefer list comprehensions and generator expressions\n\n")
		case "java":
			sb.WriteString("## Java-Specific Guidelines\n")
			sb.WriteString("- Follow Java naming conventions (camelCase, PascalCase)\n")
			sb.WriteString("- Use proper exception handling with try-with-resources\n")
			sb.WriteString("- Implement equals() and hashCode() consistently\n")
			sb.WriteString("- Use dependency injection frameworks appropriately\n")
			sb.WriteString("- Follow SOLID principles\n\n")
		case "javascript", "typescript", "node", "js", "ts":
			sb.WriteString("## Node.js/JavaScript Guidelines\n")
			sb.WriteString("- Use async/await for asynchronous operations\n")
			sb.WriteString("- Implement proper error handling with try-catch\n")
			sb.WriteString("- Use ES6+ features appropriately\n")
			sb.WriteString("- Follow Node.js best practices for modules\n")
			sb.WriteString("- Use middleware patterns for request processing\n\n")
		default:
			sb.WriteString(fmt.Sprintf("## %s Guidelines\n", config.Language))
			sb.WriteString(fmt.Sprintf("- Follow %s best practices and idioms\n", config.Language))
			sb.WriteString("- Use appropriate design patterns\n")
			sb.WriteString("- Implement proper resource management\n\n")
		}
	}

	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Implement proper error handling and logging\n")
	sb.WriteString("- Use structured logging with appropriate log levels\n")
	sb.WriteString("- Apply dependency injection patterns\n")
	sb.WriteString("- Implement proper HTTP status codes and responses\n")
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		if lang == "go" {
			sb.WriteString("- Use context.Context for request scoping and cancellation\n")
		} else if lang == "python" {
			sb.WriteString("- Use asyncio for asynchronous operations when appropriate\n")
		} else if lang == "java" {
			sb.WriteString("- Use CompletableFuture for asynchronous operations\n")
		} else if lang == "javascript" || lang == "typescript" || lang == "node" || lang == "js" || lang == "ts" {
			sb.WriteString("- Use Promise-based patterns for asynchronous operations\n")
		}
	}
	sb.WriteString("\n")

	if config.CodeStyle != "" {
		sb.WriteString("## Project Code Style\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	if config.APIRules != "" {
		sb.WriteString("## API Development\n")
		sb.WriteString(config.APIRules + "\n")
		sb.WriteString("- Use proper HTTP methods (GET, POST, PUT, DELETE, PATCH)\n")
		sb.WriteString("- Validate all input data with appropriate error messages\n")
		sb.WriteString("- Use structured JSON responses\n\n")
	}

	if config.Security != "" {
		sb.WriteString("## Security Requirements\n")
		sb.WriteString(config.Security + "\n\n")
	}

	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate code with:\n")
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		switch lang {
		case "go":
			sb.WriteString("- [ ] Comprehensive error handling with wrapped errors\n")
			sb.WriteString("- [ ] Unit tests with table-driven test patterns\n")
			sb.WriteString("- [ ] Benchmark tests for performance-critical code\n")
		case "python":
			sb.WriteString("- [ ] Proper exception handling with specific exception types\n")
			sb.WriteString("- [ ] Unit tests with pytest or unittest\n")
			sb.WriteString("- [ ] Performance profiling for critical code\n")
		case "java":
			sb.WriteString("- [ ] Comprehensive exception handling with custom exceptions\n")
			sb.WriteString("- [ ] Unit tests with JUnit and appropriate mocking\n")
			sb.WriteString("- [ ] Performance tests with JMH for critical code\n")
		case "javascript", "typescript", "node", "js", "ts":
			sb.WriteString("- [ ] Proper error handling with try-catch and error objects\n")
			sb.WriteString("- [ ] Unit tests with Jest, Mocha, or similar framework\n")
			sb.WriteString("- [ ] Performance monitoring for critical operations\n")
		default:
			sb.WriteString("- [ ] Comprehensive error handling\n")
			sb.WriteString("- [ ] Unit tests with appropriate testing framework\n")
			sb.WriteString("- [ ] Performance tests where applicable\n")
		}
	} else {
		sb.WriteString("- [ ] Comprehensive error handling\n")
		sb.WriteString("- [ ] Unit tests with appropriate patterns\n")
		sb.WriteString("- [ ] Performance tests for critical code\n")
	}
	sb.WriteString("- [ ] Proper package/module documentation and examples\n")
	sb.WriteString("- [ ] Integration tests for API endpoints\n")
	sb.WriteString("- [ ] Logging with structured fields\n")
	sb.WriteString("- [ ] Graceful shutdown handling for services\n")

	filePath := filepath.Join(instructionsDir, "backend.instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func createTestingInstructions(instructionsDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("applyTo: \"**/test/**\"\n")
	sb.WriteString("description: \"Testing guidelines for all test files and directories\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Testing Guidelines\n\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review [project structure](../../README.md) and \n")
	sb.WriteString("[testing patterns](../../) before writing tests.\n\n")

	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Follow the AAA pattern: Arrange, Act, Assert\n")
	sb.WriteString("- Write descriptive test names that explain what is being tested\n")
	sb.WriteString("- Mock external dependencies appropriately\n")
	sb.WriteString("- Ensure tests are deterministic and repeatable\n")
	sb.WriteString("- Test both happy paths and error conditions\n")

	// Language-specific testing patterns
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		switch lang {
		case "go":
			sb.WriteString("- Use table-driven tests for multiple scenarios\n\n")
			sb.WriteString("## Go Testing Patterns\n")
			sb.WriteString("- Use `testing.T` for unit tests and `testing.B` for benchmarks\n")
			sb.WriteString("- Follow Go naming conventions for test functions\n")
			sb.WriteString("- Use `testify` or similar assertion libraries when helpful\n")
			sb.WriteString("- Implement setup and teardown in test functions\n\n")
		case "python":
			sb.WriteString("- Use parameterized tests for multiple scenarios\n\n")
			sb.WriteString("## Python Testing Patterns\n")
			sb.WriteString("- Use `pytest` or `unittest` for test framework\n")
			sb.WriteString("- Use fixtures for test setup and teardown\n")
			sb.WriteString("- Use `pytest.mark.parametrize` for data-driven tests\n")
			sb.WriteString("- Use `mock` library for mocking dependencies\n\n")
		case "java":
			sb.WriteString("- Use parameterized tests for multiple scenarios\n\n")
			sb.WriteString("## Java Testing Patterns\n")
			sb.WriteString("- Use JUnit 5 for test framework\n")
			sb.WriteString("- Use `@BeforeEach` and `@AfterEach` for setup/teardown\n")
			sb.WriteString("- Use `@ParameterizedTest` for data-driven tests\n")
			sb.WriteString("- Use Mockito for mocking dependencies\n\n")
		case "javascript", "typescript", "node", "js", "ts":
			sb.WriteString("- Use describe/it blocks for test organization\n\n")
			sb.WriteString("## JavaScript/Node.js Testing Patterns\n")
			sb.WriteString("- Use Jest, Mocha, or similar testing framework\n")
			sb.WriteString("- Use `beforeEach` and `afterEach` hooks for setup\n")
			sb.WriteString("- Use test.each or similar for data-driven tests\n")
			sb.WriteString("- Use appropriate mocking libraries (Jest mocks, Sinon, etc.)\n\n")
		default:
			sb.WriteString(fmt.Sprintf("- Use appropriate testing patterns for %s\n\n", config.Language))
			sb.WriteString(fmt.Sprintf("## %s Testing Patterns\n", config.Language))
			sb.WriteString(fmt.Sprintf("- Use appropriate testing framework for %s\n", config.Language))
			sb.WriteString("- Follow language-specific testing conventions\n")
			sb.WriteString("- Use appropriate mocking/stubbing libraries\n\n")
		}
	} else {
		sb.WriteString("- Use appropriate testing patterns for your language\n\n")
	}

	if config.CodeStyle != "" {
		sb.WriteString("## Code Style in Tests\n")
		sb.WriteString(config.CodeStyle + "\n")
		sb.WriteString("Apply the same style guidelines to test code.\n\n")
	}

	sb.WriteString("## Test Organization\n")
	sb.WriteString("- Group related tests in the same file\n")
	sb.WriteString("- Use clear, descriptive test names\n")
	sb.WriteString("- Keep tests focused on single functionality\n")
	sb.WriteString("- Use test helpers for common setup\n\n")

	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate tests with:\n")
	if config.Language != "" {
		lang := strings.ToLower(config.Language)
		switch lang {
		case "go":
			sb.WriteString("- [ ] Table-driven test patterns where appropriate\n")
			sb.WriteString("- [ ] Benchmark tests for performance-critical code\n")
		case "python":
			sb.WriteString("- [ ] Pytest fixtures for test setup\n")
			sb.WriteString("- [ ] Parameterized tests for multiple scenarios\n")
		case "java":
			sb.WriteString("- [ ] JUnit test classes with proper annotations\n")
			sb.WriteString("- [ ] Parameterized tests for data-driven scenarios\n")
		case "javascript", "typescript", "node", "js", "ts":
			sb.WriteString("- [ ] Describe/it block structure for organization\n")
			sb.WriteString("- [ ] Async/await patterns for testing async code\n")
		default:
			sb.WriteString("- [ ] Appropriate test structure for your language\n")
			sb.WriteString("- [ ] Language-specific testing patterns\n")
		}
	} else {
		sb.WriteString("- [ ] Appropriate test patterns for your language\n")
		sb.WriteString("- [ ] Framework-specific testing structure\n")
	}
	sb.WriteString("- [ ] Clear test documentation and comments\n")
	sb.WriteString("- [ ] Proper setup and cleanup\n")
	sb.WriteString("- [ ] Edge case and error condition coverage\n")
	sb.WriteString("- [ ] Integration tests for complex workflows\n")
	sb.WriteString("- [ ] Mock implementations for external dependencies\n")

	filePath := filepath.Join(instructionsDir, "testing.instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}
