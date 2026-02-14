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

	sb.WriteString("## Problem Statement\n")
	sb.WriteString("[Describe the problem this feature solves. What user need or business requirement ")
	sb.WriteString("does this address? Be specific about the current pain points.]\n\n")

	sb.WriteString("## Proposed Solution\n")
	sb.WriteString("[Describe the high-level solution approach. What is being built and how does it ")
	sb.WriteString("solve the problem?]\n\n")

	sb.WriteString("## User Stories\n")
	sb.WriteString("### Primary User Story\n")
	sb.WriteString("As a [user type], I want to [action] so that [benefit].\n\n")

	sb.WriteString("### Additional User Stories\n")
	sb.WriteString("- As a [user type], I want to [action] so that [benefit]\n")
	sb.WriteString("- As a [user type], I want to [action] so that [benefit]\n\n")

	sb.WriteString("## Requirements\n\n")

	sb.WriteString("### Functional Requirements\n")
	sb.WriteString("1. [Specific functionality requirement]\n")
	sb.WriteString("2. [Specific functionality requirement]\n")
	sb.WriteString("3. [Specific functionality requirement]\n\n")

	sb.WriteString("### Non-Functional Requirements\n")
	sb.WriteString("- **Performance**: [Performance requirements, e.g., response time]\n")
	sb.WriteString("- **Security**: [Security requirements, e.g., authentication, authorization]\n")
	sb.WriteString("- **Scalability**: [Scalability requirements]\n")
	sb.WriteString("- **Accessibility**: [Accessibility requirements, e.g., WCAG compliance]\n")
	sb.WriteString("- **Compatibility**: [Browser/platform compatibility requirements]\n\n")

	sb.WriteString("## Technical Approach\n\n")

	if ctx.Config.HasBackend() {
		sb.WriteString("### Backend Changes\n")
		sb.WriteString(fmt.Sprintf("**Language**: %s\n", ctx.Config.Backend.Language))
		if ctx.Config.Backend.Framework != "" && ctx.Config.Backend.Framework != "None" {
			sb.WriteString(fmt.Sprintf("**Framework**: %s\n\n", ctx.Config.Backend.Framework))
		}
		sb.WriteString("**Components**:\n")
		sb.WriteString("- [ ] Data models: `[file path]`\n")
		sb.WriteString("- [ ] Business logic: `[file path]`\n")
		sb.WriteString("- [ ] API endpoints: `[file path]`\n")
		sb.WriteString("- [ ] Database migrations: `[file path]`\n\n")
	}

	if ctx.Config.HasFrontend() {
		sb.WriteString("### Frontend Changes\n")
		sb.WriteString(fmt.Sprintf("**Framework**: %s\n\n", ctx.Config.Frontend.Framework))
		sb.WriteString("**Components**:\n")
		sb.WriteString("- [ ] UI components: `[file path]`\n")
		sb.WriteString("- [ ] State management: `[file path]`\n")
		sb.WriteString("- [ ] API integration: `[file path]`\n")
		sb.WriteString("- [ ] Styles: `[file path]`\n\n")
	}

	if ctx.Config.HasBackend() {
		sb.WriteString("### Database Changes\n")
		sb.WriteString(fmt.Sprintf("**Database**: %s\n\n", ctx.Config.Backend.Database))
		sb.WriteString("**Schema Changes**:\n")
		sb.WriteString("- [ ] New tables: [list tables]\n")
		sb.WriteString("- [ ] Modified tables: [list tables]\n")
		sb.WriteString("- [ ] New indexes: [list indexes]\n")
		sb.WriteString("- [ ] Data migrations needed: [describe]\n\n")

		sb.WriteString("### API Contracts\n")
		sb.WriteString("#### Endpoint 1\n")
		sb.WriteString("```\n")
		sb.WriteString("POST /api/[endpoint]\n")
		sb.WriteString("```\n")
		sb.WriteString("**Request**:\n")
		sb.WriteString("```json\n")
		sb.WriteString("{\n")
		sb.WriteString("  \"field\": \"value\"\n")
		sb.WriteString("}\n")
		sb.WriteString("```\n")
		sb.WriteString("**Response**:\n")
		sb.WriteString("```json\n")
		sb.WriteString("{\n")
		sb.WriteString("  \"success\": true,\n")
		sb.WriteString("  \"data\": {}\n")
		sb.WriteString("}\n")
		sb.WriteString("```\n\n")
	}

	sb.WriteString("## Testing Strategy\n\n")

	sb.WriteString("### Unit Tests\n")
	sb.WriteString(fmt.Sprintf("**Framework**: %s\n\n", ctx.Config.Testing.Framework))
	sb.WriteString("- [ ] Test data models and validation\n")
	sb.WriteString("- [ ] Test business logic functions\n")
	sb.WriteString("- [ ] Test utility functions\n")
	sb.WriteString("- [ ] Test component rendering (frontend)\n\n")

	sb.WriteString("### Integration Tests\n")
	sb.WriteString("- [ ] Test API endpoints end-to-end\n")
	sb.WriteString("- [ ] Test database operations\n")
	sb.WriteString("- [ ] Test authentication/authorization\n")
	sb.WriteString("- [ ] Test error scenarios\n\n")

	sb.WriteString("### Manual Testing Scenarios\n")
	sb.WriteString("1. [Test scenario 1]\n")
	sb.WriteString("2. [Test scenario 2]\n")
	sb.WriteString("3. [Edge case testing]\n\n")

	sb.WriteString("## Validation Criteria\n")
	sb.WriteString("Feature is considered complete when:\n")
	sb.WriteString("- [ ] All functional requirements are implemented\n")
	sb.WriteString("- [ ] All tests pass (>90% code coverage)\n")
	sb.WriteString("- [ ] API documentation is updated\n")
	sb.WriteString("- [ ] User-facing documentation is updated\n")
	sb.WriteString("- [ ] Code review is completed\n")
	sb.WriteString("- [ ] Security review is completed\n")
	sb.WriteString("- [ ] Performance requirements are met\n")
	sb.WriteString("- [ ] Accessibility requirements are met\n\n")

	sb.WriteString("## Dependencies\n")
	sb.WriteString("### Internal Dependencies\n")
	sb.WriteString("- [ ] [Feature/component name]\n")
	sb.WriteString("- [ ] [Feature/component name]\n\n")

	sb.WriteString("### External Dependencies\n")
	sb.WriteString("- [ ] [Third-party library/service]\n")
	sb.WriteString("- [ ] [Third-party library/service]\n\n")

	sb.WriteString("## Risks and Mitigations\n")
	sb.WriteString("| Risk | Impact | Probability | Mitigation |\n")
	sb.WriteString("|------|--------|-------------|------------|\n")
	sb.WriteString("| [Risk description] | High/Medium/Low | High/Medium/Low | [Mitigation strategy] |\n\n")

	sb.WriteString("## Timeline and Milestones\n")
	sb.WriteString("- [ ] **Phase 1**: Design and specification (X days)\n")
	sb.WriteString("- [ ] **Phase 2**: Backend implementation (X days)\n")
	sb.WriteString("- [ ] **Phase 3**: Frontend implementation (X days)\n")
	sb.WriteString("- [ ] **Phase 4**: Testing and refinement (X days)\n")
	sb.WriteString("- [ ] **Phase 5**: Documentation and deployment (X days)\n\n")

	sb.WriteString("## Open Questions\n")
	sb.WriteString("- [ ] [Question that needs to be resolved]\n")
	sb.WriteString("- [ ] [Question that needs to be resolved]\n\n")

	sb.WriteString("## Handoff Checklist\n")
	sb.WriteString("Before implementation begins:\n")
	sb.WriteString("- [ ] Architecture approved by team lead\n")
	sb.WriteString("- [ ] Database schema finalized\n")
	sb.WriteString("- [ ] API contracts reviewed\n")
	sb.WriteString("- [ ] Security considerations addressed\n")
	sb.WriteString("- [ ] UI/UX design approved (if applicable)\n")
	sb.WriteString("- [ ] All dependencies identified\n")
	sb.WriteString("- [ ] Implementation ready for assignment\n\n")

	sb.WriteString("## References\n")
	sb.WriteString("- [Related documentation]\n")
	sb.WriteString("- [Design mockups]\n")
	sb.WriteString("- [Technical research]\n")

	return sb.String()
}

func generateAPIEndpointSpec(ctx GenerateContext) string {
	var sb strings.Builder

	sb.WriteString("# API Endpoint Specification: [Endpoint Name]\n\n")

	sb.WriteString("## Overview\n")
	sb.WriteString("**Purpose**: [Brief description of what this endpoint does]\n\n")
	sb.WriteString("**Endpoint**: `[METHOD] /api/v1/[resource]`\n\n")
	sb.WriteString(fmt.Sprintf("**Framework**: %s\n", ctx.Config.Backend.Language))
	if ctx.Config.Backend.Framework != "" && ctx.Config.Backend.Framework != "None" {
		sb.WriteString(fmt.Sprintf("**Backend Framework**: %s\n", ctx.Config.Backend.Framework))
	}
	sb.WriteString(fmt.Sprintf("**Database**: %s\n\n", ctx.Config.Backend.Database))

	sb.WriteString("## Authentication\n")
	sb.WriteString("- **Required**: Yes/No\n")
	sb.WriteString("- **Method**: [JWT Bearer Token / API Key / OAuth2 / None]\n")
	sb.WriteString("- **Permissions**: [Required roles or permissions]\n\n")

	sb.WriteString("## Request\n\n")

	sb.WriteString("### HTTP Method\n")
	sb.WriteString("`[GET | POST | PUT | PATCH | DELETE]`\n\n")

	sb.WriteString("### URL Parameters\n")
	sb.WriteString("| Parameter | Type | Required | Description |\n")
	sb.WriteString("|-----------|------|----------|-------------|\n")
	sb.WriteString("| `id` | integer | Yes | [Description] |\n")
	sb.WriteString("| `filter` | string | No | [Description] |\n\n")

	sb.WriteString("### Query Parameters\n")
	sb.WriteString("| Parameter | Type | Required | Default | Description |\n")
	sb.WriteString("|-----------|------|----------|---------|-------------|\n")
	sb.WriteString("| `page` | integer | No | 1 | Page number for pagination |\n")
	sb.WriteString("| `limit` | integer | No | 20 | Items per page |\n")
	sb.WriteString("| `sort` | string | No | `created_at` | Sort field |\n\n")

	sb.WriteString("### Request Headers\n")
	sb.WriteString("```\n")
	sb.WriteString("Content-Type: application/json\n")
	sb.WriteString("Authorization: Bearer <token>\n")
	sb.WriteString("```\n\n")

	sb.WriteString("### Request Body\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"field1\": \"string\",\n")
	sb.WriteString("  \"field2\": 123,\n")
	sb.WriteString("  \"field3\": {\n")
	sb.WriteString("    \"nested\": \"value\"\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("### Request Body Schema\n")
	sb.WriteString("| Field | Type | Required | Validation | Description |\n")
	sb.WriteString("|-------|------|----------|------------|-------------|\n")
	sb.WriteString("| `field1` | string | Yes | Max 255 chars | [Description] |\n")
	sb.WriteString("| `field2` | integer | Yes | Min 1, Max 1000 | [Description] |\n\n")

	sb.WriteString("## Response\n\n")

	sb.WriteString("### Success Response (200 OK)\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"success\": true,\n")
	sb.WriteString("  \"data\": {\n")
	sb.WriteString("    \"id\": 123,\n")
	sb.WriteString("    \"field1\": \"value\",\n")
	sb.WriteString("    \"created_at\": \"2026-02-14T10:00:00Z\"\n")
	sb.WriteString("  },\n")
	sb.WriteString("  \"meta\": {\n")
	sb.WriteString("    \"timestamp\": \"2026-02-14T10:00:00Z\"\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("### Error Responses\n\n")

	sb.WriteString("#### 400 Bad Request\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"success\": false,\n")
	sb.WriteString("  \"error\": {\n")
	sb.WriteString("    \"code\": \"INVALID_INPUT\",\n")
	sb.WriteString("    \"message\": \"Validation failed\",\n")
	sb.WriteString("    \"details\": [\n")
	sb.WriteString("      {\n")
	sb.WriteString("        \"field\": \"field1\",\n")
	sb.WriteString("        \"message\": \"Field is required\"\n")
	sb.WriteString("      }\n")
	sb.WriteString("    ]\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("#### 401 Unauthorized\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"success\": false,\n")
	sb.WriteString("  \"error\": {\n")
	sb.WriteString("    \"code\": \"UNAUTHORIZED\",\n")
	sb.WriteString("    \"message\": \"Authentication required\"\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("#### 403 Forbidden\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"success\": false,\n")
	sb.WriteString("  \"error\": {\n")
	sb.WriteString("    \"code\": \"FORBIDDEN\",\n")
	sb.WriteString("    \"message\": \"Insufficient permissions\"\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("#### 404 Not Found\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"success\": false,\n")
	sb.WriteString("  \"error\": {\n")
	sb.WriteString("    \"code\": \"NOT_FOUND\",\n")
	sb.WriteString("    \"message\": \"Resource not found\"\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("#### 500 Internal Server Error\n")
	sb.WriteString("```json\n")
	sb.WriteString("{\n")
	sb.WriteString("  \"success\": false,\n")
	sb.WriteString("  \"error\": {\n")
	sb.WriteString("    \"code\": \"INTERNAL_ERROR\",\n")
	sb.WriteString("    \"message\": \"An unexpected error occurred\"\n")
	sb.WriteString("  }\n")
	sb.WriteString("}\n")
	sb.WriteString("```\n\n")

	sb.WriteString("## Implementation Details\n\n")

	sb.WriteString("### File Structure\n")
	sb.WriteString("```\n")
	sb.WriteString("- Controller/Handler: [file path]\n")
	sb.WriteString("- Service/Business Logic: [file path]\n")
	sb.WriteString("- Data Model: [file path]\n")
	sb.WriteString("- Validation: [file path]\n")
	sb.WriteString("- Tests: [file path]\n")
	sb.WriteString("```\n\n")

	sb.WriteString("### Database Operations\n")
	sb.WriteString("1. [Operation description]\n")
	sb.WriteString("2. [Operation description]\n\n")

	sb.WriteString("### Business Logic\n")
	sb.WriteString("1. Validate input parameters\n")
	sb.WriteString("2. Check authentication and authorization\n")
	sb.WriteString("3. [Business logic step]\n")
	sb.WriteString("4. [Business logic step]\n")
	sb.WriteString("5. Return response\n\n")

	sb.WriteString("## Testing\n\n")

	sb.WriteString("### Unit Tests\n")
	sb.WriteString("- [ ] Test input validation\n")
	sb.WriteString("- [ ] Test business logic\n")
	sb.WriteString("- [ ] Test error handling\n")
	sb.WriteString("- [ ] Test edge cases\n\n")

	sb.WriteString("### Integration Tests\n")
	sb.WriteString("- [ ] Test successful request/response\n")
	sb.WriteString("- [ ] Test authentication failure\n")
	sb.WriteString("- [ ] Test authorization failure\n")
	sb.WriteString("- [ ] Test validation errors\n")
	sb.WriteString("- [ ] Test database failure scenarios\n\n")

	sb.WriteString("### Example Test Cases\n")
	sb.WriteString("```\n")
	sb.WriteString("Test: Valid request returns 200\n")
	sb.WriteString("Given: Valid authentication token and request body\n")
	sb.WriteString("When: POST request is made\n")
	sb.WriteString("Then: Response status is 200 and data is returned\n\n")

	sb.WriteString("Test: Missing required field returns 400\n")
	sb.WriteString("Given: Valid authentication token but missing required field\n")
	sb.WriteString("When: POST request is made\n")
	sb.WriteString("Then: Response status is 400 with validation error\n")
	sb.WriteString("```\n\n")

	sb.WriteString("## Security Considerations\n")
	if ctx.Config.General.Security != "" {
		sb.WriteString(fmt.Sprintf("**Project Security Requirements**: %s\n\n", ctx.Config.General.Security))
	}
	sb.WriteString("- [ ] Input validation and sanitization\n")
	sb.WriteString("- [ ] SQL injection prevention\n")
	sb.WriteString("- [ ] XSS prevention\n")
	sb.WriteString("- [ ] CSRF protection (if applicable)\n")
	sb.WriteString("- [ ] Rate limiting\n")
	sb.WriteString("- [ ] Authentication enforcement\n")
	sb.WriteString("- [ ] Authorization checks\n")
	sb.WriteString("- [ ] Sensitive data handling\n\n")

	sb.WriteString("## Performance Considerations\n")
	sb.WriteString("- **Expected Load**: [requests per second]\n")
	sb.WriteString("- **Response Time Target**: [milliseconds]\n")
	sb.WriteString("- **Database Query Optimization**: [describe]\n")
	sb.WriteString("- **Caching Strategy**: [if applicable]\n\n")

	sb.WriteString("## Dependencies\n")
	sb.WriteString("- [ ] [External service or API]\n")
	sb.WriteString("- [ ] [Database table]\n")
	sb.WriteString("- [ ] [Other endpoint]\n\n")

	sb.WriteString("## Documentation\n")
	sb.WriteString("- [ ] API documentation updated\n")
	sb.WriteString("- [ ] OpenAPI/Swagger spec updated\n")
	sb.WriteString("- [ ] Postman collection updated\n")
	sb.WriteString("- [ ] README updated if needed\n\n")

	sb.WriteString("## Rollout Plan\n")
	sb.WriteString("- [ ] Deploy to development environment\n")
	sb.WriteString("- [ ] Run integration tests\n")
	sb.WriteString("- [ ] Deploy to staging environment\n")
	sb.WriteString("- [ ] Manual testing in staging\n")
	sb.WriteString("- [ ] Deploy to production\n")
	sb.WriteString("- [ ] Monitor for errors\n")

	return sb.String()
}

func generateComponentSpec(ctx GenerateContext) string {
	var sb strings.Builder
	frontend := ctx.Config.Frontend

	sb.WriteString("# Component Specification: [ComponentName]\n\n")

	sb.WriteString("## Overview\n")
	sb.WriteString("**Purpose**: [Brief description of what this component does]\n\n")
	sb.WriteString(fmt.Sprintf("**Framework**: %s\n", frontend.Framework))
	sb.WriteString(fmt.Sprintf("**Language**: %s\n\n", frontend.Language))

	sb.WriteString("## Component Details\n\n")

	sb.WriteString("### Type\n")
	sb.WriteString("- [ ] Presentational (UI only, no business logic)\n")
	sb.WriteString("- [ ] Container (manages state and business logic)\n")
	sb.WriteString("- [ ] Layout (structural component)\n")
	sb.WriteString("- [ ] Page (route component)\n\n")

	sb.WriteString("### Location\n")
	sb.WriteString("**File Path**: `src/components/[ComponentName]/[ComponentName].tsx`\n\n")

	sb.WriteString("### Props Interface\n")

	if strings.Contains(strings.ToLower(frontend.Language), "typescript") ||
		strings.Contains(strings.ToLower(frontend.Language), "ts") {
		sb.WriteString("```typescript\n")
		sb.WriteString("interface ComponentNameProps {\n")
		sb.WriteString("  /** [Description of prop] */\n")
		sb.WriteString("  prop1: string;\n")
		sb.WriteString("  \n")
		sb.WriteString("  /** [Description of prop] */\n")
		sb.WriteString("  prop2?: number;\n")
		sb.WriteString("  \n")
		sb.WriteString("  /** Click handler */\n")
		sb.WriteString("  onClick?: (event: React.MouseEvent) => void;\n")
		sb.WriteString("  \n")
		sb.WriteString("  /** Child elements */\n")
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

	sb.WriteString("### State Management\n")
	sb.WriteString("**Local State**:\n")
	sb.WriteString("- `[stateName]`: [description]\n")
	sb.WriteString("- `[stateName]`: [description]\n\n")

	sb.WriteString("**Global State** (if applicable):\n")
	sb.WriteString("- Store: `[store name]`\n")
	sb.WriteString("- Selectors: `[list selectors]`\n")
	sb.WriteString("- Actions: `[list actions]`\n\n")

	sb.WriteString("## Visual Design\n\n")

	sb.WriteString("### Layout\n")
	sb.WriteString("```\n")
	sb.WriteString("[ASCII art or description of component layout]\n")
	sb.WriteString("┌─────────────────────────────┐\n")
	sb.WriteString("│  Header                     │\n")
	sb.WriteString("├─────────────────────────────┤\n")
	sb.WriteString("│  Content Area               │\n")
	sb.WriteString("│                             │\n")
	sb.WriteString("├─────────────────────────────┤\n")
	sb.WriteString("│  Actions                    │\n")
	sb.WriteString("└─────────────────────────────┘\n")
	sb.WriteString("```\n\n")

	sb.WriteString("### Styling\n")
	sb.WriteString("**Approach**: [CSS Modules / Styled Components / Tailwind / etc.]\n\n")
	sb.WriteString("**Style File**: `[ComponentName].module.css` or styled component definition\n\n")
	sb.WriteString("**Design Tokens**:\n")
	sb.WriteString("- Colors: [list relevant colors]\n")
	sb.WriteString("- Spacing: [spacing values]\n")
	sb.WriteString("- Typography: [font styles]\n\n")

	sb.WriteString("### Responsive Behavior\n")
	sb.WriteString("- **Mobile** (< 768px): [description]\n")
	sb.WriteString("- **Tablet** (768px - 1024px): [description]\n")
	sb.WriteString("- **Desktop** (> 1024px): [description]\n\n")

	sb.WriteString("## Behavior\n\n")

	sb.WriteString("### User Interactions\n")
	sb.WriteString("1. **[Action]**: [Description of what happens]\n")
	sb.WriteString("2. **[Action]**: [Description of what happens]\n\n")

	sb.WriteString("### Event Handlers\n")
	sb.WriteString("- `handle[Action]`: [description]\n")
	sb.WriteString("- `handle[Action]`: [description]\n\n")

	sb.WriteString("### Side Effects\n")
	sb.WriteString("- [ ] API calls on mount\n")
	sb.WriteString("- [ ] Subscriptions/intervals\n")
	sb.WriteString("- [ ] Event listeners\n")
	sb.WriteString("- [ ] Cleanup requirements\n\n")

	sb.WriteString("## Data Flow\n\n")

	sb.WriteString("### Input\n")
	sb.WriteString("- Props from parent component\n")
	sb.WriteString("- Data from API: `[endpoint]`\n")
	sb.WriteString("- Data from store: `[store slice]`\n\n")

	sb.WriteString("### Output\n")
	sb.WriteString("- Events emitted to parent: `[list events]`\n")
	sb.WriteString("- State updates triggered: `[list updates]`\n")
	sb.WriteString("- API calls made: `[list endpoints]`\n\n")

	sb.WriteString("## Accessibility\n\n")

	sb.WriteString("### ARIA Attributes\n")
	sb.WriteString("- `aria-label`: [description]\n")
	sb.WriteString("- `aria-describedby`: [description]\n")
	sb.WriteString("- `role`: [appropriate role]\n\n")

	sb.WriteString("### Keyboard Navigation\n")
	sb.WriteString("- `Tab`: [behavior]\n")
	sb.WriteString("- `Enter/Space`: [behavior]\n")
	sb.WriteString("- `Escape`: [behavior]\n\n")

	sb.WriteString("### Screen Reader Support\n")
	sb.WriteString("- [ ] All interactive elements have labels\n")
	sb.WriteString("- [ ] Dynamic content changes are announced\n")
	sb.WriteString("- [ ] Focus management is handled properly\n\n")

	sb.WriteString("## Testing\n\n")

	sb.WriteString("### Unit Tests\n")
	sb.WriteString(fmt.Sprintf("**Framework**: %s\n\n", ctx.Config.Testing.Framework))
	sb.WriteString("- [ ] Renders without errors\n")
	sb.WriteString("- [ ] Handles all props correctly\n")
	sb.WriteString("- [ ] Calls event handlers when expected\n")
	sb.WriteString("- [ ] Handles edge cases (null, undefined, etc.)\n\n")

	sb.WriteString("### Integration Tests\n")
	sb.WriteString("- [ ] Integrates with parent components\n")
	sb.WriteString("- [ ] API calls work correctly\n")
	sb.WriteString("- [ ] State management works correctly\n\n")

	sb.WriteString("### Visual Regression Tests\n")
	sb.WriteString("- [ ] Default state\n")
	sb.WriteString("- [ ] Interactive states (hover, focus, active)\n")
	sb.WriteString("- [ ] Error state\n")
	sb.WriteString("- [ ] Loading state\n")
	sb.WriteString("- [ ] Responsive breakpoints\n\n")

	sb.WriteString("## Performance\n\n")

	sb.WriteString("### Optimization Strategies\n")
	sb.WriteString("- [ ] Memoization (`useMemo`, `useCallback`, `React.memo`)\n")
	sb.WriteString("- [ ] Lazy loading for heavy components\n")
	sb.WriteString("- [ ] Code splitting if needed\n")
	sb.WriteString("- [ ] Virtualization for long lists\n\n")

	sb.WriteString("### Performance Targets\n")
	sb.WriteString("- Initial render: [target ms]\n")
	sb.WriteString("- Re-render time: [target ms]\n")
	sb.WriteString("- Bundle size impact: [target KB]\n\n")

	sb.WriteString("## Dependencies\n\n")

	sb.WriteString("### External Libraries\n")
	sb.WriteString("- [ ] [Library name] - [purpose]\n\n")

	sb.WriteString("### Internal Dependencies\n")
	sb.WriteString("- [ ] [Component/hook/utility] - [purpose]\n\n")

	sb.WriteString("## Implementation Checklist\n")
	sb.WriteString("- [ ] Component file created\n")
	sb.WriteString("- [ ] Types/PropTypes defined\n")
	sb.WriteString("- [ ] Styles implemented\n")
	sb.WriteString("- [ ] Logic implemented\n")
	sb.WriteString("- [ ] Unit tests written\n")
	sb.WriteString("- [ ] Integration tests written\n")
	sb.WriteString("- [ ] Accessibility verified\n")
	sb.WriteString("- [ ] Documentation updated\n")
	sb.WriteString("- [ ] Storybook story created (if applicable)\n")
	sb.WriteString("- [ ] Code review completed\n\n")

	sb.WriteString("## Examples\n\n")

	sb.WriteString("### Basic Usage\n")
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
	sb.WriteString("```\n\n")

	sb.WriteString("### Advanced Usage\n")
	if strings.Contains(strings.ToLower(frontend.Language), "typescript") ||
		strings.Contains(strings.ToLower(frontend.Language), "ts") {
		sb.WriteString("```tsx\n")
	} else {
		sb.WriteString("```jsx\n")
	}
	sb.WriteString("// Example with complex props\n")
	sb.WriteString("```\n\n")

	sb.WriteString("## Future Enhancements\n")
	sb.WriteString("- [ ] [Potential improvement]\n")
	sb.WriteString("- [ ] [Potential improvement]\n")

	return sb.String()
}
