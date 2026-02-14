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
	// Skip if no frontend configuration
	if !ctx.Config.HasFrontend() {
		return map[string]string{}, nil
	}

	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString("applyTo: \"**/*.{jsx,tsx,css,js,ts,html,vue,scss,sass,less}\"\n")
	sb.WriteString("description: \"Frontend development guidelines for UI components\"\n")
	sb.WriteString("---\n")
	sb.WriteString("# Frontend Development Guidelines\n\n")

	sb.WriteString("## Context Loading\n")
	sb.WriteString("Review [project structure](../../README.md) and \n")
	if ctx.Config.Frontend.Framework != "" && ctx.Config.Frontend.Framework != "Vanilla" {
		sb.WriteString(fmt.Sprintf("[%s component patterns](../src/) before starting.\n\n", ctx.Config.Frontend.Framework))
	} else {
		sb.WriteString("[component patterns](../src/) before starting.\n\n")
	}

	sb.WriteString("## Technology Stack\n")
	sb.WriteString(fmt.Sprintf("- **Language**: %s\n", ctx.Config.Frontend.Language))
	if ctx.Config.Frontend.Framework != "" && ctx.Config.Frontend.Framework != "Vanilla" {
		sb.WriteString(fmt.Sprintf("- **Framework**: %s\n", ctx.Config.Frontend.Framework))
	}
	if ctx.Config.Frontend.BuildTool != "" {
		sb.WriteString(fmt.Sprintf("- **Build Tool**: %s\n", ctx.Config.Frontend.BuildTool))
	}
	sb.WriteString("\n")

	sb.WriteString("## Deterministic Requirements\n")
	sb.WriteString("- Use consistent component structure and naming conventions\n")
	sb.WriteString("- Implement proper state management patterns\n")
	sb.WriteString("- Apply responsive design principles\n")
	sb.WriteString("- Ensure accessibility (WCAG guidelines)\n")
	sb.WriteString("- Use semantic HTML elements\n")

	// Language-specific best practices
	lang := strings.ToLower(ctx.Config.Frontend.Language)
	switch lang {
	case "javascript", "js":
		sb.WriteString("- Follow modern JavaScript best practices (ES6+)\n")
		sb.WriteString("- Use proper module imports/exports\n")
	case "typescript", "ts":
		sb.WriteString("- Follow TypeScript best practices with strict mode\n")
		sb.WriteString("- Use proper type definitions and interfaces\n")
		sb.WriteString("- Leverage TypeScript's type system for runtime safety\n")
	default:
		sb.WriteString(fmt.Sprintf("- Follow %s best practices and conventions\n", ctx.Config.Frontend.Language))
	}
	sb.WriteString("\n")

	if ctx.Config.General.CodeStyle != "" {
		sb.WriteString("## Project Code Style\n")
		sb.WriteString(ctx.Config.General.CodeStyle + "\n\n")
	}

	// Framework-specific guidelines
	if ctx.Config.Frontend.Framework != "" && ctx.Config.Frontend.Framework != "Vanilla" {
		framework := strings.ToLower(ctx.Config.Frontend.Framework)
		switch framework {
		case "react":
			sb.WriteString("## React Guidelines\n")
			sb.WriteString("- Prefer functional components with hooks over class components\n")
			sb.WriteString("- Use proper prop validation (PropTypes or TypeScript)\n")
			sb.WriteString("- Follow React hook rules and best practices\n")
			sb.WriteString("- Use React.memo() for performance optimization when needed\n")
		case "vue":
			sb.WriteString("## Vue.js Guidelines\n")
			sb.WriteString("- Use Vue 3 Composition API when possible\n")
			sb.WriteString("- Follow Vue single-file component structure\n")
			sb.WriteString("- Use proper reactive data patterns\n")
			sb.WriteString("- Implement proper component lifecycle management\n")
		case "angular":
			sb.WriteString("## Angular Guidelines\n")
			sb.WriteString("- Follow Angular style guide and conventions\n")
			sb.WriteString("- Use dependency injection properly\n")
			sb.WriteString("- Implement reactive forms and proper validation\n")
			sb.WriteString("- Use RxJS observables for async operations\n")
		default:
			sb.WriteString(fmt.Sprintf("## %s Guidelines\n", ctx.Config.Frontend.Framework))
			sb.WriteString(fmt.Sprintf("- Follow %s best practices and patterns\n", ctx.Config.Frontend.Framework))
			sb.WriteString("- Maintain consistent code organization\n")
		}
	} else {
		sb.WriteString("## Component Guidelines\n")
		sb.WriteString("- Keep components small and focused on single responsibility\n")
		sb.WriteString("- Use vanilla JavaScript/DOM manipulation best practices\n")
	}
	sb.WriteString("- Use proper event handling and cleanup\n")
	sb.WriteString("- Implement error boundaries where appropriate\n\n")

	sb.WriteString("## Structured Output\n")
	sb.WriteString("Generate code with:\n")
	sb.WriteString("- [ ] Component documentation with usage examples\n")
	if ctx.Config.Frontend.Language != "" {
		lang := strings.ToLower(ctx.Config.Frontend.Language)
		if lang == "javascript" || lang == "typescript" || lang == "react" || lang == "js" || lang == "ts" {
			sb.WriteString("- [ ] Proper TypeScript/JSDoc annotations\n")
			sb.WriteString("- [ ] Unit tests with Jest/React Testing Library\n")
		} else if lang == "python" {
			sb.WriteString("- [ ] Python docstrings and type hints\n")
			sb.WriteString("- [ ] Unit tests with pytest or unittest\n")
		} else {
			sb.WriteString("- [ ] Proper documentation and type annotations\n")
			sb.WriteString("- [ ] Unit tests with appropriate testing framework\n")
		}
	} else {
		sb.WriteString("- [ ] Proper documentation and annotations\n")
		sb.WriteString("- [ ] Unit tests for components\n")
	}
	sb.WriteString("- [ ] Accessibility attributes (aria-labels, roles, etc.)\n")
	sb.WriteString("- [ ] Loading and error states for async operations\n")

	return map[string]string{
		".github/instructions/frontend.instructions.md": sb.String(),
	}, nil
}
