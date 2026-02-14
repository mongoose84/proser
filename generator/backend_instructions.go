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
	// Skip if no backend configuration
	if !ctx.Config.HasBackend() {
		return map[string]string{}, nil
	}

	var sb strings.Builder

	sb.WriteString("---\n")
	// Define file extensions based on backend language
	lang := strings.ToLower(ctx.Config.Backend.Language)
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
	sb.WriteString(fmt.Sprintf("- **Language**: %s\n", ctx.Config.Backend.Language))
	if ctx.Config.Backend.Framework != "" && ctx.Config.Backend.Framework != "None" {
		sb.WriteString(fmt.Sprintf("- **Framework**: %s\n", ctx.Config.Backend.Framework))
	}
	if ctx.Config.Backend.Database != "" {
		sb.WriteString(fmt.Sprintf("- **Database**: %s\n", ctx.Config.Backend.Database))
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
		sb.WriteString(fmt.Sprintf("## %s Guidelines\n", ctx.Config.Backend.Language))
		sb.WriteString(fmt.Sprintf("- Follow %s best practices and idioms\n", ctx.Config.Backend.Language))
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

	if ctx.Config.General.CodeStyle != "" {
		sb.WriteString("## Project Code Style\n")
		sb.WriteString(ctx.Config.General.CodeStyle + "\n\n")
	}

	if ctx.Config.Backend.APIRules != "" {
		sb.WriteString("## API Development\n")
		sb.WriteString(ctx.Config.Backend.APIRules + "\n")
		sb.WriteString("- Use proper HTTP methods (GET, POST, PUT, DELETE, PATCH)\n")
		sb.WriteString("- Validate all input data with appropriate error messages\n")
		sb.WriteString("- Use structured JSON responses\n\n")
	}

	if ctx.Config.General.Security != "" {
		sb.WriteString("## Security Requirements\n")
		sb.WriteString(ctx.Config.General.Security + "\n\n")
	}

	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate code with:\n")
	if ctx.Config.Backend.Language != "" {
		lang := strings.ToLower(ctx.Config.Backend.Language)
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

	return map[string]string{
		".github/instructions/backend.instructions.md": sb.String(),
	}, nil
}
