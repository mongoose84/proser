package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mongoose84/proser/config"
)

func setupGithubFolder(cfg config.ProjectConfig) error {
	githubDir := ".github"
	if err := os.MkdirAll(githubDir, 0755); err != nil {
		return fmt.Errorf("failed to create .github directory: %w", err)
	}

	// Create copilot-instructions.md (global rules)
	if err := createCopilotInstructions(githubDir, cfg); err != nil {
		return err
	}

	// Create instructions directory
	instructionsDir := filepath.Join(githubDir, "instructions")
	if err := os.MkdirAll(instructionsDir, 0755); err != nil {
		return fmt.Errorf("failed to create instructions directory: %w", err)
	}

	// Create specific instruction files based on configuration
	if cfg.HasFrontend() {
		if err := createFrontendInstructions(instructionsDir, cfg); err != nil {
			return err
		}
	}

	if cfg.HasBackend() {
		if err := createBackendInstructions(instructionsDir, cfg); err != nil {
			return err
		}
	}

	// Always create testing instructions if we have any technology
	if cfg.HasFrontend() || cfg.HasBackend() {
		if err := createTestingInstructions(instructionsDir, cfg); err != nil {
			return err
		}
	}

	return nil
}

func createCopilotInstructions(githubDir string, cfg config.ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Global Repository Instructions\n\n"))

	sb.WriteString("## Project Overview\n")
	if cfg.General.Description != "" {
		sb.WriteString(cfg.General.Description + "\n\n")
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
	if cfg.HasFrontend() || cfg.HasBackend() {
		sb.WriteString("## Technology Stack\n")
		if cfg.HasBackend() {
			sb.WriteString(fmt.Sprintf("- **Backend**: %s", cfg.Backend.Language))
			if cfg.Backend.Framework != "" && cfg.Backend.Framework != "None" {
				sb.WriteString(fmt.Sprintf(" with %s", cfg.Backend.Framework))
			}
			sb.WriteString("\n")
		}
		if cfg.HasFrontend() {
			sb.WriteString(fmt.Sprintf("- **Frontend**: %s", cfg.Frontend.Language))
			if cfg.Frontend.Framework != "" && cfg.Frontend.Framework != "Vanilla" {
				sb.WriteString(fmt.Sprintf(" with %s", cfg.Frontend.Framework))
			}
			sb.WriteString("\n")
		}
		if cfg.HasBackend() && cfg.Backend.Database != "" {
			sb.WriteString(fmt.Sprintf("- **Database**: %s\n", cfg.Backend.Database))
		}
		if cfg.Testing.Framework != "" {
			sb.WriteString(fmt.Sprintf("- **Testing**: %s\n", cfg.Testing.Framework))
		}
		sb.WriteString("\n")
	}

	if cfg.General.CodeStyle != "" {
		sb.WriteString("## Code Style\n")
		sb.WriteString(cfg.General.CodeStyle + "\n\n")
	}

	if cfg.HasBackend() && cfg.Backend.APIRules != "" {
		sb.WriteString("## API Guidelines\n")
		sb.WriteString(cfg.Backend.APIRules + "\n\n")
	}

	if cfg.General.Security != "" {
		sb.WriteString("## Security Requirements\n")
		sb.WriteString(cfg.General.Security + "\n\n")
	}

	if cfg.General.CustomRules != "" && cfg.General.CustomRules != "None" {
		sb.WriteString("## Custom Project Rules\n")
		sb.WriteString(cfg.General.CustomRules + "\n\n")
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
