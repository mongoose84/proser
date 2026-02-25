package generator

import (
	"fmt"
	"strings"

	"github.com/mongoose84/proser/config"
)

// TestingInstructionsGenerator generates testing-specific instructions
type TestingInstructionsGenerator struct{}

// Name returns the generator name
func (g *TestingInstructionsGenerator) Name() string {
	return "testing-instructions"
}

// Generate creates testing instructions content
func (g *TestingInstructionsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	cfg := ctx.Config

	var sb strings.Builder

	// --- Frontmatter ---
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("applyTo: \"%s\"\n", testingApplyTo(cfg)))
	sb.WriteString("description: \"Testing guidelines with context engineering\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Testing Guidelines\n\n")

	sb.WriteString("Inherits from [global instructions](../copilot-instructions.md).\n\n")

	// --- Context Loading ---
	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review [project conventions](../../README.md) and\n")
	if cfg.Testing.Framework != "" {
		sb.WriteString(fmt.Sprintf("[existing %s tests](../../) before writing tests.\n\n",
			cfg.Testing.Framework))
	} else {
		sb.WriteString("[existing tests](../../) before writing tests.\n\n")
	}

	// --- Deterministic Requirements ---
	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Follow the AAA pattern: Arrange, Act, Assert\n")
	sb.WriteString("- Write descriptive test names that explain the scenario\n")
	sb.WriteString("- Mock external dependencies — keep tests isolated\n")
	sb.WriteString("- Ensure tests are deterministic and repeatable\n")
	sb.WriteString("- Cover both happy paths and error conditions\n")

	// Framework-specific requirements (folded in as bullets)
	framework := strings.ToLower(cfg.Testing.Framework)
	switch framework {
	case "jest":
		sb.WriteString("- Use `describe`/`it` blocks for organization\n")
		sb.WriteString("- Use `beforeEach`/`afterEach` for setup and teardown\n")
	case "pytest":
		sb.WriteString("- Use pytest fixtures for setup and teardown\n")
		sb.WriteString("- Use `@pytest.mark.parametrize` for data-driven tests\n")
	case "junit":
		sb.WriteString("- Use JUnit 5 annotations (`@BeforeEach`, `@ParameterizedTest`)\n")
		sb.WriteString("- Use Mockito for mocking dependencies\n")
	case "go testing":
		sb.WriteString("- Use table-driven tests for multiple scenarios\n")
		sb.WriteString("- Use `testing.T` for unit tests, `testing.B` for benchmarks\n")
	default:
		if cfg.Testing.Framework != "" {
			sb.WriteString(fmt.Sprintf("- Follow %s conventions and patterns\n", cfg.Testing.Framework))
		}
	}

	if cfg.Testing.Strategy != "" {
		sb.WriteString(fmt.Sprintf("- %s\n", cfg.Testing.Strategy))
	}
	if cfg.General.CodeStyle != "" {
		sb.WriteString(fmt.Sprintf("- %s\n", cfg.General.CodeStyle))
	}
	sb.WriteString("\n")

	// --- Structured Output ---
	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate tests with:\n")

	// Language-aware checklist
	if cfg.HasBackend() {
		lang := strings.ToLower(cfg.Backend.Language)
		switch lang {
		case "go":
			sb.WriteString("- [ ] Table-driven test patterns\n")
			sb.WriteString("- [ ] Benchmark tests for performance-critical code\n")
		case "python":
			sb.WriteString("- [ ] Pytest fixtures and parametrized cases\n")
		case "java":
			sb.WriteString("- [ ] JUnit test classes with proper annotations\n")
		case "javascript", "typescript", "node", "js", "ts":
			sb.WriteString("- [ ] Async/await patterns for async code\n")
		}
	}

	sb.WriteString("- [ ] Setup and teardown for shared state\n")
	sb.WriteString("- [ ] Edge case and error condition coverage\n")
	sb.WriteString("- [ ] Mock implementations for external dependencies\n")
	sb.WriteString("- [ ] Clear test documentation\n")

	return map[string]string{
		".github/instructions/testing.instructions.md": sb.String(),
	}, nil
}

// testingApplyTo returns a sensible glob for test files
func testingApplyTo(cfg config.ProjectConfig) string {
	if cfg.HasBackend() {
		lang := strings.ToLower(cfg.Backend.Language)
		switch lang {
		case "go":
			return "**/*_test.go"
		case "python":
			return "**/test_*.py"
		case "java":
			return "**/test/**/*.java"
		}
	}
	return "**/test/**"
}
