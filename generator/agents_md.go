package generator

import (
	"fmt"
	"strings"
)

// AgentsMdGenerator generates an AGENTS.md file at the project root.
// Per the PROSE Explicit Hierarchy constraint, AGENTS.md lives at the project
// directory level (not in every subdirectory). The generated file is intentionally
// minimal — it scaffolds the standard structure and asks the user to fill it out
// with their AI agent.
type AgentsMdGenerator struct{}

// Name returns the generator name.
func (g *AgentsMdGenerator) Name() string {
	return "agents-md"
}

// Generate creates a single AGENTS.md at the project root.
func (g *AgentsMdGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	cfg := ctx.Config

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# %s\n\n", cfg.General.ProjectName))

	if cfg.General.Description != "" {
		sb.WriteString(cfg.General.Description + "\n\n")
	}

	// Technology overview
	if cfg.HasBackend() || cfg.HasFrontend() {
		sb.WriteString("## Tech Stack\n")
		if cfg.HasBackend() {
			sb.WriteString(fmt.Sprintf("- **Backend**: %s", cfg.Backend.Language))
			if cfg.Backend.Framework != "" && cfg.Backend.Framework != "None" {
				sb.WriteString(fmt.Sprintf(" (%s)", cfg.Backend.Framework))
			}
			sb.WriteString("\n")
		}
		if cfg.HasFrontend() {
			sb.WriteString(fmt.Sprintf("- **Frontend**: %s", cfg.Frontend.Language))
			if cfg.Frontend.Framework != "" && cfg.Frontend.Framework != "Vanilla" {
				sb.WriteString(fmt.Sprintf(" (%s)", cfg.Frontend.Framework))
			}
			sb.WriteString("\n")
		}
		if cfg.HasBackend() && cfg.Backend.Database != "" {
			sb.WriteString(fmt.Sprintf("- **Database**: %s\n", cfg.Backend.Database))
		}
		sb.WriteString("\n")
	}

	// Conventions
	sb.WriteString("## Conventions\n")
	if cfg.General.CodeStyle != "" {
		sb.WriteString(fmt.Sprintf("- Code style: %s\n", cfg.General.CodeStyle))
	}
	if cfg.HasBackend() && cfg.Backend.APIRules != "" {
		sb.WriteString(fmt.Sprintf("- API design: %s\n", cfg.Backend.APIRules))
	}
	if cfg.General.Security != "" {
		sb.WriteString(fmt.Sprintf("- Security: %s\n", cfg.General.Security))
	}
	if cfg.Testing.Framework != "" {
		sb.WriteString(fmt.Sprintf("- Testing: %s (%s)\n", cfg.Testing.Framework, cfg.Testing.Strategy))
	}
	sb.WriteString("\n")

	// Prompt for AI-driven completion
	sb.WriteString("## Next Steps\n\n")
	sb.WriteString("This file was scaffolded by [proser](https://github.com/mongoose84/proser).\n")
	sb.WriteString("Use the prompt below with your AI agent to fill it out:\n\n")
	sb.WriteString("---\n\n")
	sb.WriteString("```text\n")
	sb.WriteString("Read the AGENTS.md file in this project and the PROSE specification at\n")
	sb.WriteString("https://danielmeppiel.github.io/awesome-ai-native/docs/prose/\n\n")
	sb.WriteString("Then analyze this repository and rewrite AGENTS.md to follow the PROSE\n")
	sb.WriteString("Explicit Hierarchy convention. Include:\n\n")
	sb.WriteString("- A brief project overview and architecture summary\n")
	sb.WriteString("- Key directories and what they own\n")
	sb.WriteString("- Domain-specific coding conventions and patterns\n")
	sb.WriteString("- How to build, test, and run the project\n")
	sb.WriteString("- Any guardrails or things agents should avoid\n\n")
	sb.WriteString("Keep it concise — context is a scarce resource. Remove this prompt\n")
	sb.WriteString("section when you are done.\n\n")
	sb.WriteString("If the project has distinct top-level directories (e.g. client/ and\n")
	sb.WriteString("server/), create a separate AGENTS.md in each with domain-specific\n")
	sb.WriteString("guidance that inherits from this root file.\n")
	sb.WriteString("```\n\n")
	sb.WriteString("---\n\n")
	sb.WriteString("Learn more: https://danielmeppiel.github.io/awesome-ai-native/docs/prose/\n")

	return map[string]string{
		"AGENTS.md": sb.String(),
	}, nil
}
