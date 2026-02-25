package generator

import (
	"fmt"
	"strings"
)

// SpecsGenerator generates .github/specs/*.spec.md files
type SpecsGenerator struct{}

// Name returns the generator name
func (g *SpecsGenerator) Name() string {
	return "specs"
}

// Generate creates specification template files
func (g *SpecsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	if !ctx.Config.HasSpecs() {
		return map[string]string{}, nil
	}

	files := make(map[string]string)
	cfg := ctx.Config.Specs

	if cfg.EnableFeatureTemplate {
		files[".github/specs/feature-template.spec.md"] = generateFeatureTemplateSpec(ctx)
	}

	if cfg.EnableAPIEndpoint && ctx.Config.HasBackend() {
		files[".github/specs/api-endpoint.spec.md"] = generateAPIEndpointSpec(ctx)
	}

	if cfg.EnableComponent && ctx.Config.HasFrontend() {
		files[".github/specs/component.spec.md"] = generateComponentSpec(ctx)
	}

	return files, nil
}

func generateFeatureTemplateSpec(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("# Feature: [Feature Name]\n\n")

	sb.WriteString("## Problem\n")
	sb.WriteString("[What problem does this feature solve? What user need does it address?]\n\n")

	sb.WriteString("## Solution\n")
	sb.WriteString("[How will this feature work? What is the high-level approach?]\n\n")

	sb.WriteString("## User Stories\n")
	sb.WriteString("- As a [user type], I want to [action] so that [benefit]\n")
	sb.WriteString("- As a [user type], I want to [action] so that [benefit]\n\n")

	sb.WriteString("## Technical Changes\n\n")

	if ctx.Config.HasBackend() {
		sb.WriteString("### Backend\n")
		sb.WriteString(fmt.Sprintf("**Stack**: %s", ctx.Config.Backend.Language))
		if ctx.Config.Backend.Framework != "" && ctx.Config.Backend.Framework != "None" {
			sb.WriteString(fmt.Sprintf(" with %s", ctx.Config.Backend.Framework))
		}
		sb.WriteString("\n\n")
		sb.WriteString("**Components**:\n")
		sb.WriteString("- [ ] Models/Data: `[file paths]`\n")
		sb.WriteString("- [ ] Business Logic: `[file paths]`\n")
		sb.WriteString("- [ ] API Endpoints: `[file paths]`\n")
		if ctx.Config.Backend.Database != "" {
			sb.WriteString("- [ ] Database Changes: `[migrations/schema]`\n")
		}
		sb.WriteString("\n")
	}

	if ctx.Config.HasFrontend() {
		sb.WriteString("### Frontend\n")
		sb.WriteString(fmt.Sprintf("**Stack**: %s", ctx.Config.Frontend.Framework))
		if ctx.Config.Frontend.Language != "" {
			sb.WriteString(fmt.Sprintf(" (%s)", ctx.Config.Frontend.Language))
		}
		sb.WriteString("\n\n")
		sb.WriteString("**Components**:\n")
		sb.WriteString("- [ ] UI Components: `[file paths]`\n")
		sb.WriteString("- [ ] State Management: `[file paths]`\n")
		sb.WriteString("- [ ] API Integration: `[file paths]`\n\n")
	}

	sb.WriteString("## Testing\n")
	if ctx.Config.Testing.Framework != "" {
		sb.WriteString(fmt.Sprintf("**Framework**: %s\n\n", ctx.Config.Testing.Framework))
	}
	sb.WriteString("- [ ] Unit tests for core logic\n")
	sb.WriteString("- [ ] Integration tests for APIs/components\n")
	sb.WriteString("- [ ] Edge cases and error scenarios\n\n")

	sb.WriteString("## Acceptance Criteria\n")
	sb.WriteString("- [ ] [Specific, measurable criterion]\n")
	sb.WriteString("- [ ] [Specific, measurable criterion]\n")
	sb.WriteString("- [ ] All tests pass\n")
	sb.WriteString("- [ ] Documentation updated\n\n")

	sb.WriteString("## Dependencies\n")
	sb.WriteString("- [ ] [Internal dependency or external library]\n\n")

	sb.WriteString("## Notes\n")
	sb.WriteString("[Any additional context, edge cases, or considerations]\n")

	return sb.String()
}

func generateAPIEndpointSpec(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("# API Endpoint: [Endpoint Name]\n\n")

	sb.WriteString("## Overview\n")
	sb.WriteString("**Purpose**: [What this endpoint does]\n\n")
	sb.WriteString("**Endpoint**: `[METHOD] /api/v1/[resource]`\n\n")
	sb.WriteString(fmt.Sprintf("**Stack**: %s", ctx.Config.Backend.Language))
	if ctx.Config.Backend.Framework != "" && ctx.Config.Backend.Framework != "None" {
		sb.WriteString(fmt.Sprintf(" with %s", ctx.Config.Backend.Framework))
	}
	sb.WriteString("\n\n")

	sb.WriteString("## Authentication\n")
	sb.WriteString("- **Required**: [Yes/No]\n")
	sb.WriteString("- **Method**: [JWT/API Key/OAuth2/None]\n")
	sb.WriteString("- **Permissions**: [Required roles]\n\n")

	sb.WriteString("## Request\n\n")

	sb.WriteString("### URL Parameters\n")
	sb.WriteString("| Parameter | Type | Required | Description |\n")
	sb.WriteString("|-----------|------|----------|-------------|\n")
	sb.WriteString("| `id` | integer | Yes | [Description] |\n\n")

	sb.WriteString("### Request Body\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"field1\": \"string\",\n")
	sb.WriteString("  \"field2\": 123\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("### Validation\n")
	sb.WriteString("| Field | Type | Required | Rules | Description |\n")
	sb.WriteString("|-------|------|----------|-------|-------------|\n")
	sb.WriteString("| `field1` | string | Yes | Max 255 | [Description] |\n\n")

	sb.WriteString("## Response\n\n")

	sb.WriteString("### Success (200 OK)\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"success\": true,\n")
	sb.WriteString("  \"data\": {\n")
	sb.WriteString("    \"id\": 123,\n")
	sb.WriteString("    \"field1\": \"value\"\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("### Error Responses\n")
	sb.WriteString("- **400 Bad Request**: Invalid input\n")
	sb.WriteString("- **401 Unauthorized**: Missing/invalid auth\n")
	sb.WriteString("- **404 Not Found**: Resource not found\n")
	sb.WriteString("- **500 Internal Error**: Server error\n\n")

	sb.WriteString("## Implementation\n\n")

	sb.WriteString("**Files**:\n")
	sb.WriteString("- Controller: `[file path]`\n")
	sb.WriteString("- Service: `[file path]`\n")
	sb.WriteString("- Model: `[file path]`\n")
	sb.WriteString("- Tests: `[file path]`\n\n")

	sb.WriteString("**Logic Flow**:\n")
	sb.WriteString("1. Validate input\n")
	sb.WriteString("2. Check authentication/authorization\n")
	sb.WriteString("3. [Business logic steps]\n")
	sb.WriteString("4. Return response\n\n")

	sb.WriteString("## Testing\n")
	sb.WriteString("- [ ] Valid request returns 200\n")
	sb.WriteString("- [ ] Invalid input returns 400\n")
	sb.WriteString("- [ ] Unauthorized returns 401\n")
	sb.WriteString("- [ ] Edge cases handled\n\n")

	sb.WriteString("## Security\n")
	if ctx.Config.General.Security != "" {
		sb.WriteString(fmt.Sprintf("**Project Requirements**: %s\n\n", ctx.Config.General.Security))
	}
	sb.WriteString("- [ ] Input validation\n")
	sb.WriteString("- [ ] SQL injection prevention\n")
	sb.WriteString("- [ ] Authentication checks\n")
	sb.WriteString("- [ ] Rate limiting\n")

	return sb.String()
}

func generateComponentSpec(ctx GenerateContext) string {
	var sb strings.Builder
	frontend := ctx.Config.Frontend

	sb.WriteString("# Component: [ComponentName]\n\n")

	sb.WriteString("## Overview\n")
	sb.WriteString("**Purpose**: [What this component does]\n\n")
	sb.WriteString(fmt.Sprintf("**Framework**: %s\n", frontend.Framework))
	sb.WriteString(fmt.Sprintf("**Language**: %s\n\n", frontend.Language))

	sb.WriteString("**Type**: [ ] Presentational | [ ] Container | [ ] Layout | [ ] Page\n\n")

	sb.WriteString("**Location**: `src/components/[ComponentName]/[ComponentName].tsx`\n\n")

	sb.WriteString("## Props\n")

	if strings.Contains(strings.ToLower(frontend.Language), "typescript") ||
		strings.Contains(strings.ToLower(frontend.Language), "ts") {
		sb.WriteString("```typescript\n")
		sb.WriteString("interface ComponentNameProps {\n")
		sb.WriteString("  prop1: string;\n")
		sb.WriteString("  prop2?: number;\n")
		sb.WriteString("  onClick?: (event: React.MouseEvent) => void;\n")
		sb.WriteString("  children?: React.ReactNode;\n")
		sb.WriteString("}\n")
		sb.WriteString("```\n\n")
	} else {
		sb.WriteString("```javascript\n")
		sb.WriteString("PropTypes = {\n")
		sb.WriteString("  prop1: PropTypes.string.isRequired,\n")
		sb.WriteString("  prop2: PropTypes.number,\n")
		sb.WriteString("  onClick: PropTypes.func,\n")
		sb.WriteString("  children: PropTypes.node\n")
		sb.WriteString("}\n")
		sb.WriteString("```\n\n")
	}

	sb.WriteString("## State\n")
	sb.WriteString("**Local State**:\n")
	sb.WriteString("- `[stateName]`: [description]\n\n")

	sb.WriteString("**Global State** (if needed):\n")
	sb.WriteString("- Store: `[store name]`\n")
	sb.WriteString("- Actions: `[list actions]`\n\n")

	sb.WriteString("## Visual Design\n")
	sb.WriteString("```\n")
	sb.WriteString("[Layout sketch]\n")
	sb.WriteString("┌─────────────────────────┐\n")
	sb.WriteString("│  Header                 │\n")
	sb.WriteString("├─────────────────────────┤\n")
	sb.WriteString("│  Content                │\n")
	sb.WriteString("└─────────────────────────┘\n")
	sb.WriteString("```\n\n")

	sb.WriteString("## Behavior\n")
	sb.WriteString("**User Interactions**:\n")
	sb.WriteString("- [Action]: [Result]\n\n")

	sb.WriteString("**Event Handlers**:\n")
	sb.WriteString("- `handle[Action]`: [description]\n\n")

	sb.WriteString("## Accessibility\n")
	sb.WriteString("- [ ] ARIA labels present\n")
	sb.WriteString("- [ ] Keyboard navigation works\n")
	sb.WriteString("- [ ] Screen reader compatible\n")
	sb.WriteString("- [ ] Focus management handled\n\n")

	sb.WriteString("## Testing\n")
	if ctx.Config.Testing.Framework != "" {
		sb.WriteString(fmt.Sprintf("**Framework**: %s\n\n", ctx.Config.Testing.Framework))
	}
	sb.WriteString("- [ ] Renders without errors\n")
	sb.WriteString("- [ ] Handles props correctly\n")
	sb.WriteString("- [ ] Event handlers fire\n")
	sb.WriteString("- [ ] Edge cases handled\n\n")

	sb.WriteString("## Implementation Checklist\n")
	sb.WriteString("- [ ] Component file created\n")
	sb.WriteString("- [ ] Types/PropTypes defined\n")
	sb.WriteString("- [ ] Styles implemented\n")
	sb.WriteString("- [ ] Tests written\n")
	sb.WriteString("- [ ] Accessibility verified\n")
	sb.WriteString("- [ ] Documentation updated\n\n")

	sb.WriteString("## Usage Example\n")
	if strings.Contains(strings.ToLower(frontend.Language), "typescript") ||
		strings.Contains(strings.ToLower(frontend.Language), "ts") {
		sb.WriteString("```tsx\n")
	} else {
		sb.WriteString("```jsx\n")
	}
	sb.WriteString("<ComponentName\n")
	sb.WriteString("  prop1=\"value\"\n")
	sb.WriteString("  prop2={123}\n")
	sb.WriteString("  onClick={handleClick}\n")
	sb.WriteString(">\n")
	sb.WriteString("  Content\n")
	sb.WriteString("</ComponentName>\n")
	sb.WriteString("```\n")

	return sb.String()
}
