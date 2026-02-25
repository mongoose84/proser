package generator

import (
	"fmt"
	"strings"
)

// AgentsGenerator generates .github/agents/*.agent.md files
type AgentsGenerator struct{}

// Name returns the generator name
func (g *AgentsGenerator) Name() string {
	return "agents"
}

// Generate creates agent files
func (g *AgentsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	if !ctx.Config.HasAgents() {
		return map[string]string{}, nil
	}

	files := make(map[string]string)
	cfg := ctx.Config.Agents

	if cfg.EnableArchitect {
		files[".github/agents/architect.agent.md"] = generateArchitectAgent(ctx)
	}

	if cfg.EnableFrontend && ctx.Config.HasFrontend() {
		files[".github/agents/frontend-engineer.agent.md"] = generateFrontendEngineerAgent(ctx)
	}

	if cfg.EnableBackend && ctx.Config.HasBackend() {
		files[".github/agents/backend-engineer.agent.md"] = generateBackendEngineerAgent(ctx)
	}

	if cfg.EnableCodeReviewer {
		files[".github/agents/code-reviewer.agent.md"] = generateCodeReviewerAgent(ctx)
	}

	if cfg.EnableTechnicalWriter {
		files[".github/agents/technical-writer.agent.md"] = generateTechnicalWriterAgent(ctx)
	}

	if cfg.EnableDevOps {
		files[".github/agents/devops-engineer.agent.md"] = generateDevOpsEngineerAgent(ctx)
	}

	if cfg.EnableTester {
		files[".github/agents/tester.agent.md"] = generateTesterAgent(ctx)
	}

	return files, nil
}

func generateArchitectAgent(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("description: 'System architect and planning specialist'\n")
	sb.WriteString("tools: ['changes', 'codebase', 'search', 'problems']\n")
	sb.WriteString("model: Claude Sonnet 4\n")
	sb.WriteString("---\n\n")

	sb.WriteString("System architect specializing in high-level design, architecture decisions, and technical strategy.\n\n")

	sb.WriteString("## Project Context\n")
	sb.WriteString(fmt.Sprintf("- **Project**: %s\n", ctx.Config.General.ProjectName))
	if ctx.Config.General.Description != "" {
		sb.WriteString(fmt.Sprintf("- **Description**: %s\n", ctx.Config.General.Description))
	}
	if ctx.Config.HasBackend() {
		sb.WriteString(fmt.Sprintf("- **Backend**: %s", ctx.Config.Backend.Language))
		if ctx.Config.Backend.Framework != "" && ctx.Config.Backend.Framework != "None" {
			sb.WriteString(fmt.Sprintf(" (%s)", ctx.Config.Backend.Framework))
		}
		sb.WriteString("\n")
	}
	if ctx.Config.HasFrontend() {
		sb.WriteString(fmt.Sprintf("- **Frontend**: %s", ctx.Config.Frontend.Language))
		if ctx.Config.Frontend.Framework != "" && ctx.Config.Frontend.Framework != "Vanilla" {
			sb.WriteString(fmt.Sprintf(" (%s)", ctx.Config.Frontend.Framework))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review:\n")
	sb.WriteString("1. [README](../../README.md) and [copilot-instructions](../../.github/copilot-instructions.md)\n")
	sb.WriteString("2. Existing code patterns and architecture\n")
	sb.WriteString("3. Current technical constraints and requirements\n\n")

	sb.WriteString("## Approach\n")
	sb.WriteString("- Focus on planning and design before implementation\n")
	sb.WriteString("- Consider trade-offs between approaches\n")
	sb.WriteString("- Document architectural decisions with rationale\n")

	return sb.String()
}

func generateFrontendEngineerAgent(ctx GenerateContext) string {
	var sb strings.Builder
	frontend := ctx.Config.Frontend

	sb.WriteString("---\n")
	sb.WriteString("description: 'Frontend development specialist with UI/UX focus'\n")
	sb.WriteString("tools: ['changes', 'codebase', 'editFiles', 'runCommands', 'runTasks',\n")
	sb.WriteString("        'search', 'problems', 'testFailure', 'terminalLastCommand']\n")
	sb.WriteString("model: Claude Sonnet 4\n")
	sb.WriteString("---\n\n")

	sb.WriteString(fmt.Sprintf("Frontend specialist: %s", frontend.Framework))
	if frontend.Language != "" {
		sb.WriteString(fmt.Sprintf(" with %s", frontend.Language))
	}
	sb.WriteString(". Focuses on UI/UX, component architecture, accessibility, and performance.\n\n")

	sb.WriteString("## Project Stack\n")
	sb.WriteString(fmt.Sprintf("- **Framework**: %s\n", frontend.Framework))
	sb.WriteString(fmt.Sprintf("- **Language**: %s\n", frontend.Language))
	if frontend.BuildTool != "" {
		sb.WriteString(fmt.Sprintf("- **Build Tool**: %s\n", frontend.BuildTool))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review:\n")
	sb.WriteString("1. [Frontend instructions](../../.github/instructions/frontend.instructions.md)\n")
	sb.WriteString("2. [Project overview](../../README.md)\n")
	sb.WriteString("3. Existing component patterns\n\n")

	sb.WriteString("## Approach\n")
	sb.WriteString("- Follow component-based architecture\n")
	sb.WriteString("- Ensure accessibility (WCAG) standards\n")
	sb.WriteString("- Optimize for performance and UX\n")
	sb.WriteString("- Write comprehensive component tests\n")

	return sb.String()
}

func generateBackendEngineerAgent(ctx GenerateContext) string {
	var sb strings.Builder
	backend := ctx.Config.Backend

	sb.WriteString("---\n")
	sb.WriteString("description: 'Backend development specialist with security focus'\n")
	sb.WriteString("tools: ['changes', 'codebase', 'editFiles', 'runCommands', 'runTasks',\n")
	sb.WriteString("        'search', 'problems', 'testFailure', 'terminalLastCommand']\n")
	sb.WriteString("model: Claude Sonnet 4\n")
	sb.WriteString("---\n\n")

	sb.WriteString(fmt.Sprintf("Backend specialist: %s", backend.Language))
	if backend.Framework != "" && backend.Framework != "None" {
		sb.WriteString(fmt.Sprintf(" with %s", backend.Framework))
	}
	sb.WriteString(". Focuses on API design, database optimization, security, and testing.\n\n")

	sb.WriteString("## Project Stack\n")
	sb.WriteString(fmt.Sprintf("- **Language**: %s\n", backend.Language))
	if backend.Framework != "" && backend.Framework != "None" {
		sb.WriteString(fmt.Sprintf("- **Framework**: %s\n", backend.Framework))
	}
	if backend.Database != "" {
		sb.WriteString(fmt.Sprintf("- **Database**: %s\n", backend.Database))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review:\n")
	sb.WriteString("1. [Backend instructions](../../.github/instructions/backend.instructions.md)\n")
	sb.WriteString("2. [Testing instructions](../../.github/instructions/testing.instructions.md)\n")
	sb.WriteString("3. [Project overview](../../README.md)\n")
	sb.WriteString("4. Existing patterns and architecture\n\n")

	sb.WriteString("## Approach\n")
	sb.WriteString("- Follow security-first development\n")
	sb.WriteString("- Implement proper error handling and logging\n")
	sb.WriteString("- Write comprehensive tests (unit + integration)\n")
	sb.WriteString("- Optimize database queries and API performance\n")

	return sb.String()
}

func generateCodeReviewerAgent(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("description: 'Code review specialist focused on quality and best practices'\n")
	sb.WriteString("tools: ['changes', 'codebase', 'search', 'problems']\n")
	sb.WriteString("model: Claude Sonnet 4\n")
	sb.WriteString("---\n\n")

	sb.WriteString("Code review specialist focused on quality, security, maintainability, and best practices. ")
	sb.WriteString("Provides constructive, actionable feedback.\n\n")

	sb.WriteString("## Project Standards\n")
	if ctx.Config.General.CodeStyle != "" {
		sb.WriteString(fmt.Sprintf("- **Code Style**: %s\n", ctx.Config.General.CodeStyle))
	}
	if ctx.Config.General.Security != "" {
		sb.WriteString(fmt.Sprintf("- **Security**: %s\n", ctx.Config.General.Security))
	}
	if ctx.Config.Testing.Framework != "" {
		sb.WriteString(fmt.Sprintf("- **Testing**: %s\n", ctx.Config.Testing.Framework))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review:\n")
	sb.WriteString("1. [Global instructions](../../.github/copilot-instructions.md)\n")
	if ctx.Config.HasBackend() {
		sb.WriteString("2. [Backend instructions](../../.github/instructions/backend.instructions.md)\n")
	}
	if ctx.Config.HasFrontend() {
		sb.WriteString("2. [Frontend instructions](../../.github/instructions/frontend.instructions.md)\n")
	}
	sb.WriteString("3. Changed files and context\n\n")

	sb.WriteString("## Review Focus\n")
	sb.WriteString("- [ ] Project style guidelines followed\n")
	sb.WriteString("- [ ] Security best practices applied\n")
	sb.WriteString("- [ ] Proper error handling\n")
	sb.WriteString("- [ ] Tests present and meaningful\n")
	sb.WriteString("- [ ] Clear documentation\n")
	sb.WriteString("- [ ] Performance considerations\n\n")

	sb.WriteString("## Approach\n")
	sb.WriteString("- Be constructive and specific\n")
	sb.WriteString("- Explain reasoning behind suggestions\n")
	sb.WriteString("- Prioritize critical issues\n")
	sb.WriteString("- Acknowledge improvements\n")

	return sb.String()
}

func generateTechnicalWriterAgent(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("description: 'Documentation specialist focused on clear technical writing'\n")
	sb.WriteString("tools: ['changes', 'codebase', 'editFiles', 'search']\n")
	sb.WriteString("model: Claude Sonnet 4\n")
	sb.WriteString("---\n\n")

	sb.WriteString("Technical documentation specialist focused on clear, comprehensive, and maintainable documentation. ")
	sb.WriteString("Writes for diverse audiences (developers, users, contributors).\n\n")

	sb.WriteString("## Project Context\n")
	sb.WriteString(fmt.Sprintf("- **Project**: %s\n", ctx.Config.General.ProjectName))
	if ctx.Config.General.Description != "" {
		sb.WriteString(fmt.Sprintf("- **Description**: %s\n", ctx.Config.General.Description))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review:\n")
	sb.WriteString("1. [README](../../README.md) and existing docs\n")
	sb.WriteString("2. [Project instructions](../../.github/copilot-instructions.md)\n")
	sb.WriteString("3. Code structure and patterns\n\n")

	sb.WriteString("## Documentation Standards\n")
	sb.WriteString("- Use clear, concise language\n")
	sb.WriteString("- Include practical examples\n")
	sb.WriteString("- Keep docs in sync with code\n")
	sb.WriteString("- Use proper markdown formatting\n")
	sb.WriteString("- Add diagrams where helpful\n\n")

	sb.WriteString("## Approach\n")
	sb.WriteString("- Start with user perspective\n")
	sb.WriteString("- Organize information logically\n")
	sb.WriteString("- Provide context and rationale\n")
	sb.WriteString("- Include troubleshooting guidance\n")

	return sb.String()
}

func generateDevOpsEngineerAgent(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("description: 'DevOps and infrastructure specialist'\n")
	sb.WriteString("tools: ['changes', 'codebase', 'editFiles', 'runCommands', 'runTasks',\n")
	sb.WriteString("        'search', 'problems', 'terminalLastCommand']\n")
	sb.WriteString("model: Claude Sonnet 4\n")
	sb.WriteString("---\n\n")

	sb.WriteString("DevOps specialist focused on CI/CD, deployment automation, infrastructure as code, and reliability. ")
	sb.WriteString("Prioritizes automation, monitoring, and operational excellence.\n\n")

	sb.WriteString("## Project Context\n")
	sb.WriteString(fmt.Sprintf("- **Project**: %s\n", ctx.Config.General.ProjectName))
	if ctx.Config.HasBackend() {
		sb.WriteString(fmt.Sprintf("- **Backend**: %s\n", ctx.Config.Backend.Language))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review:\n")
	sb.WriteString("1. [README](../../README.md) for build/deployment info\n")
	sb.WriteString("2. CI/CD configurations (.github/workflows, .gitlab-ci.yml, etc.)\n")
	sb.WriteString("3. Existing infrastructure code\n\n")

	sb.WriteString("## Best Practices\n")
	sb.WriteString("- Automate repetitive tasks\n")
	sb.WriteString("- Version control infrastructure\n")
	sb.WriteString("- Implement monitoring and alerting\n")
	sb.WriteString("- Follow security best practices\n")
	sb.WriteString("- Test changes before production\n\n")

	sb.WriteString("## Approach\n")
	sb.WriteString("- Prioritize reliability and stability\n")
	sb.WriteString("- Implement gradual rollouts\n")
	sb.WriteString("- Maintain clear audit trails\n")
	sb.WriteString("- Optimize for cost and performance\n")

	return sb.String()
}

func generateTesterAgent(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("description: 'QA and testing specialist focused on quality assurance'\n")
	sb.WriteString("tools: ['changes', 'codebase', 'editFiles', 'runCommands', 'runTasks',\n")
	sb.WriteString("        'search', 'problems', 'testFailure', 'terminalLastCommand']\n")
	sb.WriteString("model: Claude Sonnet 4\n")
	sb.WriteString("---\n\n")

	sb.WriteString("QA specialist focused on test automation, quality assurance, and comprehensive testing strategies. ")
	sb.WriteString("Prioritizes coverage, reliability, and defect prevention.\n\n")

	sb.WriteString("## Project Testing\n")
	if ctx.Config.Testing.Framework != "" {
		sb.WriteString(fmt.Sprintf("- **Framework**: %s\n", ctx.Config.Testing.Framework))
		sb.WriteString(fmt.Sprintf("- **Strategy**: %s\n", ctx.Config.Testing.Strategy))
	}
	if ctx.Config.HasBackend() {
		sb.WriteString(fmt.Sprintf("- **Backend**: %s\n", ctx.Config.Backend.Language))
	}
	if ctx.Config.HasFrontend() {
		sb.WriteString(fmt.Sprintf("- **Frontend**: %s\n", ctx.Config.Frontend.Framework))
	}
	sb.WriteString("\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review:\n")
	sb.WriteString("1. [Testing instructions](../../.github/instructions/testing.instructions.md)\n")
	sb.WriteString("2. [Project overview](../../README.md)\n")
	sb.WriteString("3. Existing test patterns and coverage\n\n")

	sb.WriteString("## Testing Principles\n")
	sb.WriteString("- Test behavior, not implementation\n")
	sb.WriteString("- Follow testing pyramid (unit > integration > e2e)\n")
	sb.WriteString("- Test edge cases and error conditions\n")
	sb.WriteString("- Keep tests fast and independent\n")
	sb.WriteString("- Use descriptive test names\n\n")

	sb.WriteString("## Test Checklist\n")
	sb.WriteString("- [ ] Unit tests cover core logic\n")
	sb.WriteString("- [ ] Integration tests verify interactions\n")
	sb.WriteString("- [ ] Edge cases and errors tested\n")
	sb.WriteString("- [ ] Tests are deterministic\n")
	sb.WriteString("- [ ] Proper setup/teardown\n")
	sb.WriteString("- [ ] Dependencies mocked appropriately\n\n")

	sb.WriteString("## Approach\n")
	sb.WriteString("- Understand requirements first\n")
	sb.WriteString("- Write tests as documentation\n")
	sb.WriteString("- Focus on reliability and maintainability\n")
	sb.WriteString("- Prevent flaky tests\n")
	sb.WriteString("- Provide clear bug reports\n")

	return sb.String()
}
