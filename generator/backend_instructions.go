package generator

import (
	"fmt"
	"strings"
)

// BackendInstructionsGenerator generates backend-specific instructions
type BackendInstructionsGenerator struct{}

// Name returns the generator name
func (g *BackendInstructionsGenerator) Name() string {
	return "backend-instructions"
}

// Generate creates backend instructions content
func (g *BackendInstructionsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	if !ctx.Config.HasBackend() {
		return map[string]string{}, nil
	}

	cfg := ctx.Config
	lang := strings.ToLower(cfg.Backend.Language)

	var sb strings.Builder

	// --- Frontmatter ---
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("applyTo: \"%s\"\n", backendApplyTo(lang)))
	sb.WriteString(fmt.Sprintf("description: \"%s backend development guidelines with context engineering\"\n",
		cfg.Backend.Language))
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("# %s Backend Development Guidelines\n\n", cfg.Backend.Language))

	sb.WriteString("Inherits from [global instructions](../copilot-instructions.md).\n\n")

	// --- Context Loading ---
	sb.WriteString("## Context Loading\n")
	switch lang {
	case "go":
		sb.WriteString("Review [module dependencies](../../go.mod) and\n")
		sb.WriteString("[application structure](../../main.go) before starting.\n\n")
	case "python":
		sb.WriteString("Review [dependencies](../../requirements.txt) and\n")
		sb.WriteString("[application structure](../../) before starting.\n\n")
	case "java":
		sb.WriteString("Review [build dependencies](../../pom.xml) and\n")
		sb.WriteString("[application structure](../../src/main/java/) before starting.\n\n")
	case "javascript", "typescript", "node", "node.js", "js", "ts":
		sb.WriteString("Review [dependencies](../../package.json) and\n")
		sb.WriteString("[application structure](../../src/) before starting.\n\n")
	default:
		sb.WriteString("Review [project dependencies](../../) and\n")
		sb.WriteString("[application structure](../../) before starting.\n\n")
	}

	// --- Deterministic Requirements ---
	sb.WriteString("## Deterministic Requirements\n")

	// Language-specific requirements
	switch lang {
	case "go":
		sb.WriteString("- Follow Effective Go conventions and idioms\n")
		sb.WriteString("- Use interfaces to define behavior contracts\n")
		sb.WriteString("- Implement resource cleanup with `defer`\n")
		sb.WriteString("- Use `context.Context` for request scoping and cancellation\n")
	case "python":
		sb.WriteString("- Follow PEP 8 style guidelines\n")
		sb.WriteString("- Use type hints for function signatures\n")
		sb.WriteString("- Use context managers for resource management\n")
		sb.WriteString("- Use asyncio for async operations when appropriate\n")
	case "java":
		sb.WriteString("- Follow Java naming conventions (camelCase, PascalCase)\n")
		sb.WriteString("- Use try-with-resources for resource management\n")
		sb.WriteString("- Apply dependency injection and SOLID principles\n")
	case "javascript", "typescript", "node", "node.js", "js", "ts":
		sb.WriteString("- Use async/await for asynchronous operations\n")
		sb.WriteString("- Follow ES6+ module patterns\n")
		sb.WriteString("- Implement proper error handling with try-catch\n")
	case "rust", "rs":
		sb.WriteString("- Use Result and Option types for error handling\n")
		sb.WriteString("- Follow Rust ownership and borrowing patterns\n")
	default:
		sb.WriteString(fmt.Sprintf("- Follow %s best practices and idioms\n", cfg.Backend.Language))
	}

	// Framework-specific requirements
	if cfg.Backend.Framework != "" && cfg.Backend.Framework != "None" {
		sb.WriteString(fmt.Sprintf("- Follow %s patterns and conventions\n", cfg.Backend.Framework))
	}

	// Universal backend requirements
	sb.WriteString("- Implement structured logging with appropriate levels\n")
	sb.WriteString("- Use proper HTTP status codes and error responses\n")

	if cfg.Backend.APIRules != "" {
		sb.WriteString(fmt.Sprintf("- %s\n", cfg.Backend.APIRules))
	}
	if cfg.General.CodeStyle != "" {
		sb.WriteString(fmt.Sprintf("- %s\n", cfg.General.CodeStyle))
	}
	if cfg.General.Security != "" {
		sb.WriteString(fmt.Sprintf("- %s\n", cfg.General.Security))
	}
	sb.WriteString("\n")

	// --- Structured Output ---
	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate code with:\n")

	switch lang {
	case "go":
		sb.WriteString("- [ ] Wrapped errors with context\n")
		sb.WriteString("- [ ] Table-driven unit tests\n")
	case "python":
		sb.WriteString("- [ ] Specific exception types\n")
		sb.WriteString("- [ ] Pytest-style unit tests\n")
	case "java":
		sb.WriteString("- [ ] Custom exception hierarchy\n")
		sb.WriteString("- [ ] JUnit tests with appropriate mocking\n")
	case "javascript", "typescript", "node", "node.js", "js", "ts":
		sb.WriteString("- [ ] Typed error objects\n")
		sb.WriteString("- [ ] Unit tests with Jest or Mocha\n")
	default:
		sb.WriteString("- [ ] Comprehensive error handling\n")
		sb.WriteString("- [ ] Unit tests with appropriate framework\n")
	}

	sb.WriteString("- [ ] Package/module documentation\n")
	sb.WriteString("- [ ] Integration tests for API endpoints\n")
	sb.WriteString("- [ ] Graceful shutdown handling\n")

	return map[string]string{
		".github/instructions/backend.instructions.md": sb.String(),
	}, nil
}

// backendApplyTo returns the applyTo glob for a given backend language
func backendApplyTo(lang string) string {
	switch lang {
	case "go":
		return "**/*.go"
	case "python":
		return "**/*.py"
	case "java":
		return "**/*.java"
	case "javascript", "node.js", "node", "js":
		return "**/*.js"
	case "typescript", "ts":
		return "**/*.ts"
	case "rust", "rs":
		return "**/*.rs"
	case "c#", "csharp":
		return "**/*.cs"
	default:
		return "**/*.{go,py,java,rs,js,ts}"
	}
}
