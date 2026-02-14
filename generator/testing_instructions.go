package generator

import (
	"fmt"
	"strings"
)

// TestingInstructionsGenerator generates testing-specific instructions
type TestingInstructionsGenerator struct{}

// Name returns the generator name
func (g *TestingInstructionsGenerator) Name() string {
	return "testing-instructions"
}

// Generate creates testing instructions content
func (g *TestingInstructionsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("applyTo: \"**/test/**\"\n")
	sb.WriteString("description: \"Testing guidelines for all test files and directories\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Testing Guidelines\n\n")

	sb.WriteString("## Technology Stack\n")
	if ctx.Config.Testing.Framework != "" {
		sb.WriteString(fmt.Sprintf("- **Testing Framework**: %s\n", ctx.Config.Testing.Framework))
	}
	if ctx.Config.Testing.Strategy != "" {
		sb.WriteString(fmt.Sprintf("- **Testing Strategy**: %s\n", ctx.Config.Testing.Strategy))
	}
	// Include relevant backend/frontend languages for test context
	if ctx.Config.HasBackend() {
		sb.WriteString(fmt.Sprintf("- **Backend Language**: %s\n", ctx.Config.Backend.Language))
	}
	if ctx.Config.HasFrontend() {
		sb.WriteString(fmt.Sprintf("- **Frontend Language**: %s\n", ctx.Config.Frontend.Language))
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
	framework := strings.ToLower(ctx.Config.Testing.Framework)
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
		if ctx.Config.Testing.Framework != "" {
			sb.WriteString(fmt.Sprintf("- Use appropriate testing patterns for %s\n\n", ctx.Config.Testing.Framework))
			sb.WriteString(fmt.Sprintf("## %s Testing Patterns\n", ctx.Config.Testing.Framework))
			sb.WriteString(fmt.Sprintf("- Follow %s best practices and conventions\n", ctx.Config.Testing.Framework))
			sb.WriteString("- Use appropriate assertion and mocking libraries\n\n")
		} else {
			sb.WriteString("- Use appropriate testing patterns for your framework\n\n")
		}
	}

	if ctx.Config.General.CodeStyle != "" {
		sb.WriteString("## Code Style in Tests\n")
		sb.WriteString(ctx.Config.General.CodeStyle + "\n")
		sb.WriteString("Apply the same style guidelines to test code.\n\n")
	}

	sb.WriteString("## Test Organization\n")
	sb.WriteString("- Group related tests in the same file\n")
	sb.WriteString("- Use clear, descriptive test names\n")
	sb.WriteString("- Keep tests focused on single functionality\n")
	sb.WriteString("- Use test helpers for common setup\n\n")

	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate tests with:\n")
	if ctx.Config.Backend != nil {
		lang := strings.ToLower(ctx.Config.Backend.Language)
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

	return map[string]string{
		".github/instructions/testing.instructions.md": sb.String(),
	}, nil
}
