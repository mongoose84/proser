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

	// Create specific instruction files based on configuration
	if config.FrontendLanguage != "" && strings.ToLower(config.FrontendLanguage) != "skip" {
		if err := createFrontendInstructions(instructionsDir, config); err != nil {
			return err
		}
	}

	if config.BackendLanguage != "" && strings.ToLower(config.BackendLanguage) != "skip" {
		if err := createBackendInstructions(instructionsDir, config); err != nil {
			return err
		}
	}

	// Always create testing instructions if we have any technology
	if (config.FrontendLanguage != "" && strings.ToLower(config.FrontendLanguage) != "skip") ||
		(config.BackendLanguage != "" && strings.ToLower(config.BackendLanguage) != "skip") {
		if err := createTestingInstructions(instructionsDir, config); err != nil {
			return err
		}
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
		sb.WriteString("This project follows specific development guidelines for different technology stacks.\n\n")
	}

	sb.WriteString("## Universal Principles\n")
	sb.WriteString("- Write clean, maintainable, and well-documented code\n")
	sb.WriteString("- Follow established project conventions and patterns\n")
	sb.WriteString("- Ensure proper error handling and logging\n")
	sb.WriteString("- Implement comprehensive testing for all new features\n")
	sb.WriteString("- Maintain consistent code formatting and style\n\n")

	// Technology stack overview
	hasFrontend := config.FrontendLanguage != "" && strings.ToLower(config.FrontendLanguage) != "skip"
	hasBackend := config.BackendLanguage != "" && strings.ToLower(config.BackendLanguage) != "skip"

	if hasFrontend || hasBackend {
		sb.WriteString("## Technology Stack\n")
		if hasBackend {
			sb.WriteString(fmt.Sprintf("- **Backend**: %s", config.BackendLanguage))
			if config.BackendFramework != "" && config.BackendFramework != "None" {
				sb.WriteString(fmt.Sprintf(" with %s", config.BackendFramework))
			}
			sb.WriteString("\n")
		}
		if hasFrontend {
			sb.WriteString(fmt.Sprintf("- **Frontend**: %s", config.FrontendLanguage))
			if config.FrontendFramework != "" && config.FrontendFramework != "Vanilla" {
				sb.WriteString(fmt.Sprintf(" with %s", config.FrontendFramework))
			}
			sb.WriteString("\n")
		}
		if config.BackendDatabase != "" {
			sb.WriteString(fmt.Sprintf("- **Database**: %s\n", config.BackendDatabase))
		}
		if config.TestingFramework != "" {
			sb.WriteString(fmt.Sprintf("- **Testing**: %s\n", config.TestingFramework))
		}
		sb.WriteString("\n")
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

	if config.CustomRules != "" && config.CustomRules != "None" {
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
