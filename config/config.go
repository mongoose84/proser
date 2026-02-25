package config

// GeneralConfig holds general project configuration
type GeneralConfig struct {
	ProjectName string
	Description string
	CodeStyle   string
	Security    string
	CustomRules string
}

// FrontendConfig holds frontend-specific configuration
type FrontendConfig struct {
	Language  string // js, ts, etc.
	Framework string // React, Vue, Angular, etc.
	BuildTool string // Webpack, Vite, etc.
}

// BackendConfig holds backend-specific configuration
type BackendConfig struct {
	Language  string // Go, Python, Java, Node.js, etc.
	Framework string // Express, Flask, Spring, etc.
	Database  string // PostgreSQL, MongoDB, etc.
	APIRules  string
}

// TestingConfig holds testing-specific configuration
type TestingConfig struct {
	Framework string // Jest, pytest, JUnit, etc.
	Strategy  string // Unit, Integration, E2E focus
}

// AgentsConfig holds agent configuration
type AgentsConfig struct {
	EnableArchitect       bool
	EnableFrontend        bool
	EnableBackend         bool
	EnableCodeReviewer    bool
	EnableTechnicalWriter bool
	EnableDevOps          bool
	EnableTester          bool
}

// PromptsConfig holds prompt templates configuration
type PromptsConfig struct {
	EnableCodeReview    bool
	EnableFeatureSpec   bool
	EnableRefactor      bool
	EnableBugFix        bool
	EnablePRDescription bool
}

// SpecsConfig holds specification templates configuration
type SpecsConfig struct {
	EnableFeatureTemplate bool
	EnableAPIEndpoint     bool
	EnableComponent       bool
}

// ProjectConfig is the main configuration structure
type ProjectConfig struct {
	General  GeneralConfig
	Frontend *FrontendConfig // nil if no frontend
	Backend  *BackendConfig  // nil if no backend
	Testing  TestingConfig
	Agents   *AgentsConfig  // nil if no agents
	Prompts  *PromptsConfig // nil if no prompts
	Specs    *SpecsConfig   // nil if no specs
}

// HasFrontend returns true if the project has frontend configuration
func (c *ProjectConfig) HasFrontend() bool {
	return c.Frontend != nil
}

// HasBackend returns true if the project has backend configuration
func (c *ProjectConfig) HasBackend() bool {
	return c.Backend != nil
}

// HasAgents returns true if the project has agents configuration
func (c *ProjectConfig) HasAgents() bool {
	return c.Agents != nil
}

// HasPrompts returns true if the project has prompts configuration
func (c *ProjectConfig) HasPrompts() bool {
	return c.Prompts != nil
}

// HasSpecs returns true if the project has specs configuration
func (c *ProjectConfig) HasSpecs() bool {
	return c.Specs != nil
}
