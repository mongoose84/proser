package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func createFrontendInstructions(instructionsDir string, config ProjectConfig) error {
	// Skip if no frontend
	if config.FrontendLanguage == "" || strings.ToLower(config.FrontendLanguage) == "skip" {
		return nil
	}

	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("applyTo: \"**/*.{jsx,tsx,css,js,ts,html,vue,scss,sass,less}\"\n")
	sb.WriteString("description: \"Frontend development guidelines for UI components\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Frontend Development Guidelines\n\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review [project structure](../../README.md) and \n")
	if config.FrontendFramework != "" && config.FrontendFramework != "Vanilla" {
		sb.WriteString(fmt.Sprintf("[%s component patterns](../src/) before starting.\n\n", config.FrontendFramework))
	} else {
		sb.WriteString("[component patterns](../src/) before starting.\n\n")
	}

	sb.WriteString("## Technology Stack\n")
	sb.WriteString(fmt.Sprintf("- **Language**: %s\n", config.FrontendLanguage))
	if config.FrontendFramework != "" && config.FrontendFramework != "Vanilla" {
		sb.WriteString(fmt.Sprintf("- **Framework**: %s\n", config.FrontendFramework))
	}
	if config.FrontendBuildTool != "" {
		sb.WriteString(fmt.Sprintf("- **Build Tool**: %s\n", config.FrontendBuildTool))
	}
	sb.WriteString("\n")

	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Use consistent component structure and naming conventions\n")
	sb.WriteString("- Implement proper state management patterns\n")
	sb.WriteString("- Apply responsive design principles\n")
	sb.WriteString("- Ensure accessibility (WCAG guidelines)\n")
	sb.WriteString("- Use semantic HTML elements\n")

	// Language-specific best practices
	lang := strings.ToLower(config.FrontendLanguage)
	switch lang {
	case "javascript", "js":
		sb.WriteString("- Follow modern JavaScript best practices (ES6+)\n")
		sb.WriteString("- Use proper module imports/exports\n")
	case "typescript", "ts":
		sb.WriteString("- Follow TypeScript best practices with strict mode\n")
		sb.WriteString("- Use proper type definitions and interfaces\n")
		sb.WriteString("- Leverage TypeScript's type system for runtime safety\n")
	default:
		sb.WriteString(fmt.Sprintf("- Follow %s best practices and conventions\n", config.FrontendLanguage))
	}
	sb.WriteString("\n")

	if config.CodeStyle != "" {
		sb.WriteString("## Project Code Style\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	// Framework-specific guidelines
	if config.FrontendFramework != "" && config.FrontendFramework != "Vanilla" {
		framework := strings.ToLower(config.FrontendFramework)
		switch framework {
		case "react":
			sb.WriteString("## React Guidelines\n")
			sb.WriteString("- Prefer functional components with hooks over class components\n")
			sb.WriteString("- Use proper prop validation (PropTypes or TypeScript)\n")
			sb.WriteString("- Follow React hook rules and best practices\n")
			sb.WriteString("- Use React.memo() for performance optimization when needed\n")
		case "vue":
			sb.WriteString("## Vue.js Guidelines\n")
			sb.WriteString("- Use Vue 3 Composition API when possible\n")
			sb.WriteString("- Follow Vue single-file component structure\n")
			sb.WriteString("- Use proper reactive data patterns\n")
			sb.WriteString("- Implement proper component lifecycle management\n")
		case "angular":
			sb.WriteString("## Angular Guidelines\n")
			sb.WriteString("- Follow Angular style guide and conventions\n")
			sb.WriteString("- Use dependency injection properly\n")
			sb.WriteString("- Implement reactive forms and proper validation\n")
			sb.WriteString("- Use RxJS observables for async operations\n")
		default:
			sb.WriteString(fmt.Sprintf("## %s Guidelines\n", config.FrontendFramework))
			sb.WriteString(fmt.Sprintf("- Follow %s best practices and patterns\n", config.FrontendFramework))
			sb.WriteString("- Maintain consistent code organization\n")
		}
	} else {
		sb.WriteString("## Component Guidelines\n")
		sb.WriteString("- Keep components small and focused on single responsibility\n")
		sb.WriteString("- Use vanilla JavaScript/DOM manipulation best practices\n")
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
	// Skip if no backend
	if config.BackendLanguage == "" || strings.ToLower(config.BackendLanguage) == "skip" {
		return nil
	}

	var sb strings.Builder

	sb.WriteString("---\n")
	// Define file extensions based on backend language
	lang := strings.ToLower(config.BackendLanguage)
	switch lang {
	case "go":
		sb.WriteString("applyTo: \"**/*.go\"\n")
	case "python":
		sb.WriteString("applyTo: \"**/*.py\"\n")
	case "java":
		sb.WriteString("applyTo: \"**/*.java\"\n")
	case "javascript", "node.js", "node", "js":
		sb.WriteString("applyTo: \"**/*.js\"\n")
	case "typescript", "ts":
		sb.WriteString("applyTo: \"**/*.ts\"\n")
	case "rust", "rs":
		sb.WriteString("applyTo: \"**/*.rs\"\n")
	case "c#", "csharp":
		sb.WriteString("applyTo: \"**/*.cs\"\n")
	default:
		sb.WriteString("applyTo: \"**/*.{go,py,java,rs,js,ts}\"\n")
	}
	sb.WriteString("description: \"Backend development guidelines for server-side languages\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Backend Development Guidelines\n\n")

	sb.WriteString("## Technology Stack\n")
	sb.WriteString(fmt.Sprintf("- **Language**: %s\n", config.BackendLanguage))
	if config.BackendFramework != "" && config.BackendFramework != "None" {
		sb.WriteString(fmt.Sprintf("- **Framework**: %s\n", config.BackendFramework))
	}
	if config.BackendDatabase != "" {
		sb.WriteString(fmt.Sprintf("- **Database**: %s\n", config.BackendDatabase))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
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

	// Language-specific guidelines
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
		sb.WriteString(fmt.Sprintf("## %s Guidelines\n", config.BackendLanguage))
		sb.WriteString(fmt.Sprintf("- Follow %s best practices and idioms\n", config.BackendLanguage))
		sb.WriteString("- Use appropriate design patterns\n")
		sb.WriteString("- Implement proper resource management\n\n")
	}

	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Implement proper error handling and logging\n")
	sb.WriteString("- Use structured logging with appropriate log levels\n")
	sb.WriteString("- Apply dependency injection patterns\n")
	sb.WriteString("- Implement proper HTTP status codes and responses\n")
	switch lang {
	case "go":
		sb.WriteString("- Use context.Context for request scoping and cancellation\n")
	case "python":
		sb.WriteString("- Use asyncio for asynchronous operations when appropriate\n")
	case "java":
		sb.WriteString("- Use CompletableFuture for asynchronous operations\n")
	case "javascript", "typescript", "node", "js", "ts":
		sb.WriteString("- Use Promise-based patterns for asynchronous operations\n")
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

	sb.WriteString("## Technology Stack\n")
	if config.TestingFramework != "" {
		sb.WriteString(fmt.Sprintf("- **Testing Framework**: %s\n", config.TestingFramework))
	}
	if config.TestingStrategy != "" {
		sb.WriteString(fmt.Sprintf("- **Testing Strategy**: %s\n", config.TestingStrategy))
	}
	// Include relevant backend/frontend languages for test context
	if config.BackendLanguage != "" && strings.ToLower(config.BackendLanguage) != "skip" {
		sb.WriteString(fmt.Sprintf("- **Backend Language**: %s\n", config.BackendLanguage))
	}
	if config.FrontendLanguage != "" && strings.ToLower(config.FrontendLanguage) != "skip" {
		sb.WriteString(fmt.Sprintf("- **Frontend Language**: %s\n", config.FrontendLanguage))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review [project structure](../../README.md) and \n")
	sb.WriteString("[testing patterns](../../) before writing tests.\n\n")

	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Follow the AAA pattern: Arrange, Act, Assert\n")
	sb.WriteString("- Write descriptive test names that explain what is being tested\n")
	sb.WriteString("- Mock external dependencies appropriately\n")
	sb.WriteString("- Ensure tests are deterministic and repeatable\n")
	sb.WriteString("- Test both happy paths and error conditions\n")

	// Framework-specific testing patterns
	framework := strings.ToLower(config.TestingFramework)
	switch framework {
	case "jest":
		sb.WriteString("- Use describe/it blocks for test organization\n\n")
		sb.WriteString("## Jest Testing Patterns\n")
		sb.WriteString("- Use Jest's built-in assertion methods\n")
		sb.WriteString("- Use `beforeEach` and `afterEach` hooks for setup\n")
		sb.WriteString("- Use `test.each` or similar for data-driven tests\n")
		sb.WriteString("- Use Jest mocks and spies for dependencies\n\n")
	case "pytest":
		sb.WriteString("- Use parameterized tests for multiple scenarios\n\n")
		sb.WriteString("## Pytest Testing Patterns\n")
		sb.WriteString("- Use pytest fixtures for test setup and teardown\n")
		sb.WriteString("- Use `pytest.mark.parametrize` for data-driven tests\n")
		sb.WriteString("- Use `mock` library for mocking dependencies\n")
		sb.WriteString("- Use pytest plugins for additional functionality\n\n")
	case "junit":
		sb.WriteString("- Use parameterized tests for multiple scenarios\n\n")
		sb.WriteString("## JUnit Testing Patterns\n")
		sb.WriteString("- Use JUnit 5 annotations and assertions\n")
		sb.WriteString("- Use `@BeforeEach` and `@AfterEach` for setup/teardown\n")
		sb.WriteString("- Use `@ParameterizedTest` for data-driven tests\n")
		sb.WriteString("- Use Mockito for mocking dependencies\n\n")
	case "go testing":
		sb.WriteString("- Use table-driven tests for multiple scenarios\n\n")
		sb.WriteString("## Go Testing Patterns\n")
		sb.WriteString("- Use `testing.T` for unit tests and `testing.B` for benchmarks\n")
		sb.WriteString("- Follow Go naming conventions for test functions\n")
		sb.WriteString("- Use `testify` or similar assertion libraries when helpful\n")
		sb.WriteString("- Implement setup and teardown in test functions\n\n")
	default:
		if config.TestingFramework != "" {
			sb.WriteString(fmt.Sprintf("- Use appropriate testing patterns for %s\n\n", config.TestingFramework))
			sb.WriteString(fmt.Sprintf("## %s Testing Patterns\n", config.TestingFramework))
			sb.WriteString(fmt.Sprintf("- Follow %s best practices and conventions\n", config.TestingFramework))
			sb.WriteString("- Use appropriate assertion and mocking libraries\n\n")
		} else {
			sb.WriteString("- Use appropriate testing patterns for your framework\n\n")
		}
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
