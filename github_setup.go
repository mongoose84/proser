package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func setupGithubFolder(config ProjectConfig) error {
	githubDir := ".github"
	if err := os.MkdirAll(githubDir, 0755); err != nil {
		return fmt.Errorf("failed to create .github directory: %w", err)
	}

	// Create copilot-instructions.md
	if err := createCopilotInstructions(githubDir, config); err != nil {
		return err
	}

	// Create agent instructions
	agentsDir := filepath.Join(githubDir, "agents")
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		return fmt.Errorf("failed to create agents directory: %w", err)
	}

	if err := createAgentInstructions(agentsDir, config); err != nil {
		return err
	}

	return nil
}

func createCopilotInstructions(githubDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# GitHub Copilot Instructions for %s\n\n", config.ProjectName))

	if config.Description != "" {
		sb.WriteString("## Project Overview\n")
		sb.WriteString(config.Description + "\n\n")
	}

	if config.Language != "" {
		sb.WriteString("## Programming Language\n")
		sb.WriteString(config.Language + "\n\n")
	}

	if config.CodeStyle != "" {
		sb.WriteString("## Code Style Guidelines\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	if config.APIRules != "" {
		sb.WriteString("## API Design Rules\n")
		sb.WriteString(config.APIRules + "\n\n")
	}

	if config.Security != "" {
		sb.WriteString("## Security Requirements\n")
		sb.WriteString(config.Security + "\n\n")
	}

	sb.WriteString("\n## PROSE Principles\n")
	sb.WriteString("When operating within this project, adhere to these core PROSE principles:\n")
	sb.WriteString("- **Progressive Disclosure**: Structure information to reveal complexity progressively for efficient context utilization.\n")
	sb.WriteString("- **Reduced Scope**: Match task size to context capacity to manage complexity.\n")
	sb.WriteString("- **Orchestrated Composition**: Favor simple, composable components over complex, collapsed structures.\n")
	sb.WriteString("- **Safety Boundaries**: Maintain autonomy within established guardrails for reliability and security.\n")
	sb.WriteString("- **Explicit Hierarchy**: Increase specificity as scope narrows to support modularity and inheritance.\n\n")

	sb.WriteString("## General Instructions for AI Assistants\n\n")
	sb.WriteString("When working on this project, please:\n\n")

	i := 1
	if config.Language != "" {
		sb.WriteString(fmt.Sprintf("%d. **Follow Language Best Practices**: Write idiomatic %s code following community standards\n", i, config.Language))
		i++
	}
	if config.CodeStyle != "" {
		sb.WriteString(fmt.Sprintf("%d. **Maintain Code Style**: Adhere to the specified code style guidelines: %s\n", i, config.CodeStyle))
		i++
	}
	if config.APIRules != "" {
		sb.WriteString(fmt.Sprintf("%d. **API Consistency**: Follow these API rules: %s\n", i, config.APIRules))
		i++
	}
	if config.Security != "" {
		sb.WriteString(fmt.Sprintf("%d. **Security First**: Always consider security implications: %s\n", i, config.Security))
		i++
	}
	sb.WriteString(fmt.Sprintf("%d. **Documentation**: Keep documentation up-to-date with code changes\n", i))
	i++
	sb.WriteString(fmt.Sprintf("%d. **Testing**: Write tests for new functionality and bug fixes\n", i))
	i++
	sb.WriteString(fmt.Sprintf("%d. **Error Handling**: Implement proper error handling and logging\n", i))
	i++
	sb.WriteString(fmt.Sprintf("%d. **Performance**: Consider performance implications of changes\n", i))

	sb.WriteString("\n## Code Review Checklist\n\n")
	sb.WriteString("Before submitting changes, ensure:\n")
	if config.CodeStyle != "" {
		sb.WriteString("- [ ] Code follows style guidelines\n")
	}
	if config.Security != "" {
		sb.WriteString("- [ ] Security requirements are met\n")
	}
	if config.APIRules != "" {
		sb.WriteString("- [ ] API changes are documented\n")
	}
	sb.WriteString("- [ ] Tests are included and passing\n")
	sb.WriteString("- [ ] Error handling is appropriate\n")
	sb.WriteString("- [ ] Documentation is updated\n")

	filePath := filepath.Join(githubDir, "copilot-instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func createAgentInstructions(agentsDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Agent Instructions for %s\n\n", config.ProjectName))
	sb.WriteString("## Context\n")
	if config.Language != "" {
		sb.WriteString(fmt.Sprintf("This is a %s project with specific guidelines that all agents should follow.\n\n", config.Language))
	} else {
		sb.WriteString("This project has specific guidelines that all agents should follow.\n\n")
	}

	sb.WriteString("## Project Configuration\n\n")

	if config.Language != "" {
		sb.WriteString(fmt.Sprintf("### Language: %s\n", config.Language))
		sb.WriteString(fmt.Sprintf("All code should be written in %s following best practices for this language.\n\n", config.Language))
	}

	if config.CodeStyle != "" {
		sb.WriteString("### Code Style\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	if config.APIRules != "" {
		sb.WriteString("### API Design\n")
		sb.WriteString(config.APIRules + "\n\n")
	}

	if config.Security != "" {
		sb.WriteString("### Security\n")
		sb.WriteString(config.Security + "\n\n")
	}

	if config.CustomRules != "" {
		sb.WriteString("### Custom Rules\n")
		sb.WriteString(config.CustomRules + "\n\n")
	}

	sb.WriteString("## PROSE Operational Principles\n")
	sb.WriteString("Agents must operate according to the following principles:\n")
	sb.WriteString("1. **Progressive Disclosure**: Information should be structured to reveal complexity only as needed.\n")
	sb.WriteString("2. **Reduced Scope**: Align task breadth with the current context capacity.\n")
	sb.WriteString("3. **Orchestrated Composition**: Build complex systems by composing simple, well-defined components.\n")
	sb.WriteString("4. **Safety Boundaries**: Exercise autonomy strictly within the project's security and logic guardrails.\n")
	sb.WriteString("5. **Explicit Hierarchy**: Honor the specific rules of the current directory over general project rules.\n\n")

	sb.WriteString(`## Agent Responsibilities

### Code Generation
- Generate code that follows the language and style guidelines
- Include appropriate error handling
- Add inline comments for complex logic
- Consider edge cases

### Code Review
- Check for adherence to style guidelines
- Verify security requirements are met
- Ensure API consistency
- Validate test coverage

### Documentation
- Keep README and documentation in sync with code
- Document API endpoints and parameters
- Include usage examples

### Testing
- Write unit tests for new functions
- Include integration tests where appropriate
- Ensure tests are maintainable

## Communication Style
- Be clear and concise
- Explain technical decisions
- Highlight security considerations
- Suggest improvements when appropriate
`)

	filePath := filepath.Join(agentsDir, "general-instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}
