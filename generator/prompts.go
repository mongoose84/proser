package generator

import (
	"fmt"
	"strings"
)

// PromptsGenerator generates .github/prompts/*.prompt.md files
type PromptsGenerator struct{}

// Name returns the generator name
func (g *PromptsGenerator) Name() string {
	return "prompts"
}

// Generate creates prompt template files
func (g *PromptsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	if !ctx.Config.HasPrompts() {
		return map[string]string{}, nil
	}

	files := make(map[string]string)
	cfg := ctx.Config.Prompts

	if cfg.EnableCodeReview {
		files[".github/prompts/code-review.prompt.md"] = generateCodeReviewPrompt(ctx)
	}

	if cfg.EnableFeatureSpec {
		files[".github/prompts/feature-spec.prompt.md"] = generateFeatureSpecPrompt(ctx)
	}

	if cfg.EnableRefactor {
		files[".github/prompts/refactor.prompt.md"] = generateRefactorPrompt(ctx)
	}

	if cfg.EnableBugFix {
		files[".github/prompts/bug-fix.prompt.md"] = generateBugFixPrompt(ctx)
	}

	if cfg.EnablePRDescription {
		files[".github/prompts/pr-description.prompt.md"] = generatePRDescriptionPrompt(ctx)
	}

	return files, nil
}

func generateCodeReviewPrompt(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("mode: agent\n")
	sb.WriteString("model: gpt-4\n")
	sb.WriteString("tools: ['file-search', 'semantic-search', 'changes', 'problems']\n")
	sb.WriteString("description: 'Structured code review workflow with validation gates'\n")
	sb.WriteString("---\n")
	sb.WriteString("# Code Review Workflow\n\n")

	sb.WriteString("## Context Loading Phase\n")
	sb.WriteString("1. Review [project guidelines](../../docs/standards/)\n")
	sb.WriteString("2. Check [changed files](changes) in current PR\n")
	sb.WriteString("3. Analyze [existing issues](problems) and warnings\n")
	if ctx.Config.General.CodeStyle != "" {
		sb.WriteString(fmt.Sprintf("4. Verify adherence to: %s\n", ctx.Config.General.CodeStyle))
	}
	if ctx.Config.General.Security != "" {
		sb.WriteString(fmt.Sprintf("5. Check security requirements: %s\n", ctx.Config.General.Security))
	}
	sb.WriteString("\n")

	sb.WriteString("## Review Checklist\n")
	sb.WriteString("### Code Quality\n")
	sb.WriteString("- [ ] Code follows project style guidelines\n")
	sb.WriteString("- [ ] Functions/methods have clear, single responsibilities\n")
	sb.WriteString("- [ ] Variable and function names are descriptive\n")
	sb.WriteString("- [ ] No unnecessary complexity or over-engineering\n")
	sb.WriteString("- [ ] Code is DRY (Don't Repeat Yourself)\n\n")

	sb.WriteString("### Security\n")
	sb.WriteString("- [ ] No hard-coded credentials or secrets\n")
	sb.WriteString("- [ ] Input validation is present\n")
	sb.WriteString("- [ ] No SQL injection vulnerabilities\n")
	sb.WriteString("- [ ] Authentication/authorization checks in place\n")
	sb.WriteString("- [ ] Sensitive data is properly handled\n\n")

	sb.WriteString("### Testing\n")
	sb.WriteString("- [ ] Unit tests cover new/modified code\n")
	sb.WriteString("- [ ] Edge cases are tested\n")
	sb.WriteString("- [ ] Tests are meaningful and not just for coverage\n")
	sb.WriteString("- [ ] Integration tests updated if needed\n\n")

	sb.WriteString("### Documentation\n")
	sb.WriteString("- [ ] Public APIs are documented\n")
	sb.WriteString("- [ ] Complex logic has explanatory comments\n")
	sb.WriteString("- [ ] README updated if needed\n")
	sb.WriteString("- [ ] CHANGELOG updated for user-facing changes\n\n")

	sb.WriteString("### Performance\n")
	sb.WriteString("- [ ] No obvious performance bottlenecks\n")
	sb.WriteString("- [ ] Database queries are optimized\n")
	sb.WriteString("- [ ] No N+1 query problems\n")
	sb.WriteString("- [ ] Resource cleanup (connections, files) is handled\n\n")

	sb.WriteString("## Deterministic Execution\n")
	sb.WriteString("Use semantic search to find similar patterns: `semantic-search \"<pattern>\"`\n")
	sb.WriteString("Use file search to locate related files: `file-search \"**/*.test.*\"`\n\n")

	sb.WriteString("## Structured Output Requirements\n")
	sb.WriteString("Provide review feedback in the following format:\n\n")
	sb.WriteString("### Summary\n")
	sb.WriteString("[High-level assessment of the changes]\n\n")
	sb.WriteString("### Critical Issues\n")
	sb.WriteString("[Issues that must be fixed before merging]\n\n")
	sb.WriteString("### Suggestions\n")
	sb.WriteString("[Recommended improvements]\n\n")
	sb.WriteString("### Positive Observations\n")
	sb.WriteString("[Good patterns or improvements worth noting]\n\n")

	sb.WriteString("## Human Validation Gate\n")
	sb.WriteString("🚨 **STOP**: Review feedback before posting.\n")
	sb.WriteString("Confirm: Feedback is constructive, specific, and actionable.\n")

	return sb.String()
}

func generateFeatureSpecPrompt(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("mode: agent\n")
	sb.WriteString("model: gpt-4\n")
	sb.WriteString("tools: ['file-search', 'semantic-search', 'codebase']\n")
	sb.WriteString("description: 'Feature implementation workflow with specification-first approach'\n")
	sb.WriteString("---\n")
	sb.WriteString("# Feature Implementation from Specification\n\n")

	sb.WriteString("## Context Loading Phase\n")
	sb.WriteString("1. Review [project specification](${specFile})\n")
	sb.WriteString("2. Analyze [existing codebase patterns](../../src/)\n")
	sb.WriteString("3. Check [API documentation](../../docs/api/)\n")
	sb.WriteString("4. Review [architecture guidelines](../../docs/architecture/)\n\n")

	sb.WriteString("## Planning Phase\n")
	sb.WriteString("### Requirements Analysis\n")
	sb.WriteString("- [ ] Understand problem statement\n")
	sb.WriteString("- [ ] Identify affected components\n")
	sb.WriteString("- [ ] List dependencies and integrations\n")
	sb.WriteString("- [ ] Identify breaking changes\n\n")

	sb.WriteString("### Technical Design\n")
	sb.WriteString("- [ ] Define data models and types\n")
	sb.WriteString("- [ ] Design API contracts\n")
	sb.WriteString("- [ ] Plan database changes\n")
	sb.WriteString("- [ ] Consider error handling\n")
	sb.WriteString("- [ ] Plan for testing\n\n")

	sb.WriteString("## Deterministic Execution\n")
	sb.WriteString("Use semantic search to find similar implementations:\n")
	sb.WriteString("`semantic-search \"similar feature implementation\"`\n\n")
	sb.WriteString("Use file search to locate test patterns:\n")
	sb.WriteString("`file-search \"**/*.test.{js,ts,go,py}\"`\n\n")

	sb.WriteString("## Implementation Checklist\n")
	if ctx.Config.HasBackend() {
		sb.WriteString("### Backend Implementation\n")
		sb.WriteString("- [ ] Create/update data models\n")
		sb.WriteString("- [ ] Implement business logic\n")
		sb.WriteString("- [ ] Add API endpoints\n")
		sb.WriteString("- [ ] Handle errors properly\n")
		sb.WriteString("- [ ] Add validation\n\n")
	}

	if ctx.Config.HasFrontend() {
		sb.WriteString("### Frontend Implementation\n")
		sb.WriteString("- [ ] Create/update components\n")
		sb.WriteString("- [ ] Implement state management\n")
		sb.WriteString("- [ ] Add API integration\n")
		sb.WriteString("- [ ] Handle loading and error states\n")
		sb.WriteString("- [ ] Ensure accessibility\n\n")
	}

	sb.WriteString("### Testing\n")
	sb.WriteString("- [ ] Write unit tests (>90% coverage target)\n")
	sb.WriteString("- [ ] Add integration tests\n")
	sb.WriteString("- [ ] Test error scenarios\n")
	sb.WriteString("- [ ] Verify edge cases\n\n")

	sb.WriteString("### Documentation\n")
	sb.WriteString("- [ ] Update API documentation\n")
	sb.WriteString("- [ ] Add inline code comments\n")
	sb.WriteString("- [ ] Update README if needed\n")
	sb.WriteString("- [ ] Add usage examples\n\n")

	sb.WriteString("## Structured Output Requirements\n")
	sb.WriteString("Generate implementation with:\n")
	sb.WriteString("1. Feature code in appropriate module\n")
	sb.WriteString("2. Comprehensive unit tests\n")
	sb.WriteString("3. Integration tests for API endpoints\n")
	sb.WriteString("4. Documentation updates\n\n")

	sb.WriteString("## Human Validation Gate\n")
	sb.WriteString("🚨 **STOP**: Review implementation plan before proceeding to code generation.\n")
	sb.WriteString("Confirm: Architecture alignment, test strategy, and breaking change impact.\n")

	return sb.String()
}

func generateRefactorPrompt(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("mode: agent\n")
	sb.WriteString("model: gpt-4\n")
	sb.WriteString("tools: ['file-search', 'semantic-search', 'codebase', 'editFiles', 'runTests']\n")
	sb.WriteString("description: 'Code refactoring workflow with safety checks'\n")
	sb.WriteString("---\n")
	sb.WriteString("# Code Refactoring Workflow\n\n")

	sb.WriteString("## Context Loading Phase\n")
	sb.WriteString("1. Review [target code](${targetFile})\n")
	sb.WriteString("2. Identify all [usages](semantic-search \"function_name\")\n")
	sb.WriteString("3. Check [test coverage](${testFiles})\n")
	sb.WriteString("4. Review [related documentation](../../docs/)\n\n")

	sb.WriteString("## Refactoring Planning\n")
	sb.WriteString("### Analysis\n")
	sb.WriteString("- [ ] Identify code smells and issues\n")
	sb.WriteString("- [ ] Map dependencies and impacts\n")
	sb.WriteString("- [ ] Verify test coverage exists\n")
	sb.WriteString("- [ ] List breaking changes\n\n")

	sb.WriteString("### Refactoring Strategy\n")
	sb.WriteString("- [ ] Define refactoring objectives\n")
	sb.WriteString("- [ ] Plan incremental steps\n")
	sb.WriteString("- [ ] Identify safe transformation patterns\n")
	sb.WriteString("- [ ] Plan for backward compatibility if needed\n\n")

	sb.WriteString("## Safe Refactoring Principles\n")
	sb.WriteString("1. **Small Steps**: Make small, incremental changes\n")
	sb.WriteString("2. **Test First**: Ensure tests pass before and after\n")
	sb.WriteString("3. **One Change**: One refactoring technique at a time\n")
	sb.WriteString("4. **Verify Often**: Run tests after each change\n\n")

	sb.WriteString("## Common Refactoring Patterns\n")
	sb.WriteString("### Simplification\n")
	sb.WriteString("- Extract method/function\n")
	sb.WriteString("- Inline method/function\n")
	sb.WriteString("- Consolidate duplicate code\n")
	sb.WriteString("- Simplify conditional expressions\n\n")

	sb.WriteString("### Organization\n")
	sb.WriteString("- Move method/function\n")
	sb.WriteString("- Rename for clarity\n")
	sb.WriteString("- Extract class/module\n")
	sb.WriteString("- Organize imports\n\n")

	sb.WriteString("## Execution Steps\n")
	sb.WriteString("1. Run existing tests to establish baseline\n")
	sb.WriteString("2. Apply refactoring incrementally\n")
	sb.WriteString("3. Run tests after each change\n")
	sb.WriteString("4. Update documentation\n")
	sb.WriteString("5. Final test run\n\n")

	sb.WriteString("## Validation Checklist\n")
	sb.WriteString("- [ ] All tests pass\n")
	sb.WriteString("- [ ] No functionality has changed\n")
	sb.WriteString("- [ ] Code is more readable\n")
	sb.WriteString("- [ ] Complexity has reduced\n")
	sb.WriteString("- [ ] Performance is maintained or improved\n")
	sb.WriteString("- [ ] Documentation is updated\n\n")

	sb.WriteString("## Human Validation Gate\n")
	sb.WriteString("🚨 **STOP**: Review refactoring plan before execution.\n")
	sb.WriteString("Confirm: Test coverage is adequate, changes are incremental, rollback plan exists.\n")

	return sb.String()
}

func generateBugFixPrompt(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("mode: agent\n")
	sb.WriteString("model: gpt-4\n")
	sb.WriteString("tools: ['file-search', 'semantic-search', 'codebase', 'problems', 'testFailure', 'editFiles', 'runTests']\n")
	sb.WriteString("description: 'Systematic bug investigation and fix workflow'\n")
	sb.WriteString("---\n")
	sb.WriteString("# Bug Fix Workflow\n\n")

	sb.WriteString("## Context Loading Phase\n")
	sb.WriteString("1. Review [bug report](${issueLink})\n")
	sb.WriteString("2. Check [related code](semantic-search \"${component}\")\n")
	sb.WriteString("3. Review [test failures](testFailure)\n")
	sb.WriteString("4. Check [recent changes](changes) to affected code\n\n")

	sb.WriteString("## Investigation Phase\n")
	sb.WriteString("### Reproduce the Bug\n")
	sb.WriteString("- [ ] Understand expected behavior\n")
	sb.WriteString("- [ ] Identify actual behavior\n")
	sb.WriteString("- [ ] Create minimal reproduction case\n")
	sb.WriteString("- [ ] Document reproduction steps\n\n")

	sb.WriteString("### Root Cause Analysis\n")
	sb.WriteString("- [ ] Trace code execution path\n")
	sb.WriteString("- [ ] Identify point of failure\n")
	sb.WriteString("- [ ] Understand why the bug occurs\n")
	sb.WriteString("- [ ] Check for similar bugs elsewhere\n\n")

	sb.WriteString("## Fix Strategy\n")
	sb.WriteString("### Planning\n")
	sb.WriteString("- [ ] Determine fix approach\n")
	sb.WriteString("- [ ] Identify affected components\n")
	sb.WriteString("- [ ] Consider edge cases\n")
	sb.WriteString("- [ ] Plan for regression prevention\n\n")

	sb.WriteString("### Implementation\n")
	sb.WriteString("1. Write a failing test that reproduces the bug\n")
	sb.WriteString("2. Implement the fix\n")
	sb.WriteString("3. Verify the test now passes\n")
	sb.WriteString("4. Run all tests to check for regressions\n")
	sb.WriteString("5. Add additional tests for edge cases\n\n")

	sb.WriteString("## Fix Validation Checklist\n")
	sb.WriteString("- [ ] Bug is reproducible before fix\n")
	sb.WriteString("- [ ] Bug is fixed after changes\n")
	sb.WriteString("- [ ] New tests prevent regression\n")
	sb.WriteString("- [ ] All existing tests pass\n")
	sb.WriteString("- [ ] No new issues introduced\n")
	sb.WriteString("- [ ] Edge cases are handled\n")
	sb.WriteString("- [ ] Documentation updated if needed\n\n")

	sb.WriteString("## Deterministic Execution\n")
	sb.WriteString("Use semantic search to find related code:\n")
	sb.WriteString("`semantic-search \"similar functionality\"`\n\n")
	sb.WriteString("Use file search to find test files:\n")
	sb.WriteString("`file-search \"**/*test*\"`\n\n")

	sb.WriteString("## Structured Output Requirements\n")
	sb.WriteString("Provide fix with:\n")
	sb.WriteString("1. Clear description of root cause\n")
	sb.WriteString("2. Explanation of fix approach\n")
	sb.WriteString("3. Code changes\n")
	sb.WriteString("4. Test that reproduces and validates fix\n")
	sb.WriteString("5. Any documentation updates\n\n")

	sb.WriteString("## Human Validation Gate\n")
	sb.WriteString("🚨 **STOP**: Review fix strategy before implementation.\n")
	sb.WriteString("Confirm: Root cause is understood, fix is minimal, tests are comprehensive.\n")

	return sb.String()
}

func generatePRDescriptionPrompt(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("mode: agent\n")
	sb.WriteString("model: gpt-4\n")
	sb.WriteString("tools: ['changes', 'codebase', 'semantic-search']\n")
	sb.WriteString("description: 'Generate comprehensive pull request descriptions'\n")
	sb.WriteString("---\n")
	sb.WriteString("# Pull Request Description Generator\n\n")

	sb.WriteString("## Context Loading Phase\n")
	sb.WriteString("1. Review [changed files](changes)\n")
	sb.WriteString("2. Analyze [commit messages](git log)\n")
	sb.WriteString("3. Check [related issues](${issueLinks})\n")
	sb.WriteString("4. Understand [project context](../../README.md)\n\n")

	sb.WriteString("## PR Description Structure\n\n")

	sb.WriteString("### Title\n")
	sb.WriteString("Create a clear, concise title following the format:\n")
	sb.WriteString("`[Type] Brief description of changes`\n\n")
	sb.WriteString("Types: `feat`, `fix`, `refactor`, `docs`, `test`, `chore`\n\n")

	sb.WriteString("### Description Template\n")
	sb.WriteString("```markdown\n")
	sb.WriteString("## Overview\n")
	sb.WriteString("[Brief summary of what this PR does]\n\n")

	sb.WriteString("## Problem Statement\n")
	sb.WriteString("[What problem does this solve? Link to issue if applicable]\n\n")

	sb.WriteString("## Solution\n")
	sb.WriteString("[How does this PR solve the problem?]\n\n")

	sb.WriteString("## Changes\n")
	sb.WriteString("### Added\n")
	sb.WriteString("- [New features or functionality]\n\n")

	sb.WriteString("### Modified\n")
	sb.WriteString("- [Changed functionality]\n\n")

	sb.WriteString("### Removed\n")
	sb.WriteString("- [Deleted functionality]\n\n")

	sb.WriteString("## Testing\n")
	sb.WriteString("- [ ] Unit tests added/updated\n")
	sb.WriteString("- [ ] Integration tests added/updated\n")
	sb.WriteString("- [ ] Manual testing performed\n\n")

	sb.WriteString("### Test Plan\n")
	sb.WriteString("[How to test these changes]\n\n")

	sb.WriteString("## Breaking Changes\n")
	sb.WriteString("[List any breaking changes, or \"None\"]\n\n")

	sb.WriteString("## Documentation\n")
	sb.WriteString("- [ ] Documentation updated\n")
	sb.WriteString("- [ ] API docs updated (if applicable)\n")
	sb.WriteString("- [ ] README updated (if applicable)\n\n")

	sb.WriteString("## Checklist\n")
	sb.WriteString("- [ ] Code follows project style guidelines\n")
	sb.WriteString("- [ ] Self-review completed\n")
	sb.WriteString("- [ ] Tests pass locally\n")
	sb.WriteString("- [ ] No new warnings\n")
	sb.WriteString("- [ ] Documentation is clear\n\n")

	sb.WriteString("## Related Issues\n")
	sb.WriteString("Closes #[issue_number]\n")
	sb.WriteString("Related to #[issue_number]\n")
	sb.WriteString("```\n\n")

	sb.WriteString("## Content Guidelines\n")
	sb.WriteString("- Be specific and factual\n")
	sb.WriteString("- Explain the \"why\" not just the \"what\"\n")
	sb.WriteString("- Include screenshots for UI changes\n")
	sb.WriteString("- Link to related documentation\n")
	sb.WriteString("- Mention performance implications\n")
	sb.WriteString("- Note any deployment considerations\n\n")

	sb.WriteString("## Human Validation Gate\n")
	sb.WriteString("🚨 **STOP**: Review generated description.\n")
	sb.WriteString("Confirm: Description is accurate, complete, and helpful for reviewers.\n")

	return sb.String()
}
