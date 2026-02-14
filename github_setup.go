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

	// Create copilot-instructions.md (global rules)
	if err := createCopilotInstructions(githubDir, config); err != nil {
		return err
	}

	// Create instructions directory
	instructionsDir := filepath.Join(githubDir, "instructions")
	if err := os.MkdirAll(instructionsDir, 0755); err != nil {
		return fmt.Errorf("failed to create instructions directory: %w", err)
	}

	// Create specific instruction files
	if err := createFrontendInstructions(instructionsDir, config); err != nil {
		return err
	}

	if err := createBackendInstructions(instructionsDir, config); err != nil {
		return err
	}

	if err := createTestingInstructions(instructionsDir, config); err != nil {
		return err
	}

	return nil
}

func createCopilotInstructions(githubDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Global Repository Instructions\n\n"))

	sb.WriteString("## Project Overview\n")
	if config.Description != "" {
		sb.WriteString(config.Description + "\n\n")
	} else {
		sb.WriteString(fmt.Sprintf("This is a %s project with specific development guidelines.\n\n", config.Language))
	}

	sb.WriteString("## Universal Principles\n")
	sb.WriteString("- Write clean, maintainable, and well-documented code\n")
	sb.WriteString("- Follow established project conventions and patterns\n")
	sb.WriteString("- Ensure proper error handling and logging\n")
	sb.WriteString("- Implement comprehensive testing for all new features\n")
	sb.WriteString("- Maintain consistent code formatting and style\n\n")

	if config.Language != "" {
		sb.WriteString(fmt.Sprintf("## Primary Language: %s\n", config.Language))
		sb.WriteString(fmt.Sprintf("Follow %s best practices and idiomatic patterns.\n\n", config.Language))
	}

	if config.CodeStyle != "" {
		sb.WriteString("## Code Style\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	if config.APIRules != "" {
		sb.WriteString("## API Guidelines\n")
		sb.WriteString(config.APIRules + "\n\n")
	}

	if config.Security != "" {
		sb.WriteString("## Security Requirements\n")
		sb.WriteString(config.Security + "\n\n")
	}

	if config.CustomRules != "" {
		sb.WriteString("## Custom Project Rules\n")
		sb.WriteString(config.CustomRules + "\n\n")
	}

	sb.WriteString("## Documentation Standards\n")
	sb.WriteString("- Include clear README files for major components\n")
	sb.WriteString("- Document all public APIs and interfaces\n")
	sb.WriteString("- Provide usage examples where appropriate\n")
	sb.WriteString("- Keep documentation up-to-date with code changes\n\n")

	sb.WriteString("## Performance Considerations\n")
	sb.WriteString("- Write efficient algorithms and data structures\n")
	sb.WriteString("- Consider scalability implications\n")
	sb.WriteString("- Optimize for both development and runtime performance\n")
	sb.WriteString("- Profile and benchmark critical code paths\n")

	filePath := filepath.Join(githubDir, "copilot-instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}
