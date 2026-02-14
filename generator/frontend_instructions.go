package generator

import (
	"fmt"
	"strings"
)

// FrontendInstructionsGenerator generates frontend-specific instructions
type FrontendInstructionsGenerator struct{}

// Name returns the generator name
func (g *FrontendInstructionsGenerator) Name() string {
	return "frontend-instructions"
}

// Generate creates frontend instructions content
func (g *FrontendInstructionsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	if !ctx.Config.HasFrontend() {
		return map[string]string{}, nil
	}

	cfg := ctx.Config
	lang := strings.ToLower(cfg.Frontend.Language)

	var sb strings.Builder

	// --- Frontmatter ---
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("applyTo: \"%s\"\n", frontendApplyTo(lang, cfg.Frontend.Framework)))
	sb.WriteString(fmt.Sprintf("description: \"%s development guidelines with context engineering\"\n",
		cfg.Frontend.Language))
	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("# %s Development Guidelines\n\n", cfg.Frontend.Language))

	// --- Context Loading ---
	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review [project conventions](../../README.md)")
	if cfg.Frontend.Framework != "" && cfg.Frontend.Framework != "Vanilla" {
		sb.WriteString(fmt.Sprintf(" and\n[%s component patterns](../../src/) before starting.\n\n",
			cfg.Frontend.Framework))
	} else {
		sb.WriteString(" and\n[component patterns](../../src/) before starting.\n\n")
	}

	// --- Deterministic Requirements ---
	sb.WriteString("## Deterministic Requirements\n")

	// Language-specific requirements
	switch lang {
	case "typescript", "ts":
		sb.WriteString("- Use strict TypeScript configuration\n")
		sb.WriteString("- Define proper type definitions and interfaces\n")
		sb.WriteString("- Avoid `any` — leverage the type system for safety\n")
	case "javascript", "js":
		sb.WriteString("- Follow modern JavaScript best practices (ES6+)\n")
		sb.WriteString("- Use proper module imports/exports\n")
	default:
		sb.WriteString(fmt.Sprintf("- Follow %s best practices and conventions\n", cfg.Frontend.Language))
	}

	// Framework-specific requirements
	if cfg.Frontend.Framework != "" && cfg.Frontend.Framework != "Vanilla" {
		switch strings.ToLower(cfg.Frontend.Framework) {
		case "react":
			sb.WriteString("- Prefer functional components with hooks\n")
			sb.WriteString("- Implement error boundaries for React components\n")
		case "vue":
			sb.WriteString("- Use Vue 3 Composition API\n")
			sb.WriteString("- Follow single-file component structure\n")
		case "angular":
			sb.WriteString("- Follow the Angular style guide\n")
			sb.WriteString("- Use dependency injection and RxJS observables\n")
		default:
			sb.WriteString(fmt.Sprintf("- Follow %s patterns and conventions\n", cfg.Frontend.Framework))
		}
	}

	// Universal frontend requirements
	sb.WriteString("- Ensure accessibility (WCAG guidelines, semantic HTML)\n")
	sb.WriteString("- Apply responsive design principles\n")

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
	case "typescript", "ts":
		sb.WriteString("- [ ] JSDoc comments for all public APIs\n")
		sb.WriteString("- [ ] Type exports in appropriate index files\n")
	case "javascript", "js":
		sb.WriteString("- [ ] JSDoc comments for all public APIs\n")
	default:
		sb.WriteString("- [ ] Documentation and type annotations for public APIs\n")
	}

	sb.WriteString("- [ ] Unit tests in appropriate test directory\n")
	sb.WriteString("- [ ] Accessibility attributes (aria-labels, roles)\n")
	sb.WriteString("- [ ] Loading and error states for async operations\n")

	return map[string]string{
		".github/instructions/frontend.instructions.md": sb.String(),
	}, nil
}

// frontendApplyTo returns the applyTo glob for a given frontend language/framework
func frontendApplyTo(lang, framework string) string {
	fw := strings.ToLower(framework)
	switch {
	case fw == "vue":
		return "**/*.{vue,js,ts,css,scss,sass,less}"
	case fw == "angular":
		return "**/*.{ts,html,css,scss,sass,less}"
	}
	switch lang {
	case "typescript", "ts":
		return "**/*.{ts,tsx,css,scss,sass,less}"
	case "javascript", "js":
		return "**/*.{js,jsx,css,scss,sass,less}"
	default:
		return "**/*.{js,jsx,ts,tsx,css,html,vue,scss,sass,less}"
	}
}
