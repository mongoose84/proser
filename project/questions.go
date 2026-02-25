package project

import "github.com/mongoose84/proser/input"

// generalQuestions returns questions for general project configuration
func generalQuestions() []input.Question {
	return []input.Question{
		{Key: "project_name", Prompt: "Project name", DefaultValue: "my-project"},
		{Key: "description", Prompt: "Project description", DefaultValue: "A software project"},
		{Key: "code_style", Prompt: "General code style guidelines (e.g., follow PEP8, use gofmt, ESLint rules)", DefaultValue: "Follow standard formatting"},
		{Key: "security", Prompt: "Security requirements (e.g., authentication methods, data encryption, OWASP compliance)", DefaultValue: "Follow OWASP top 10"},
		{Key: "custom_rules", Prompt: "Additional custom rules or guidelines", DefaultValue: "None"},
	}
}

// frontendQuestions returns questions for frontend configuration
func frontendQuestions() []input.Question {
	return []input.Question{
		{Key: "frontend_language", Prompt: "Frontend language (e.g., JavaScript, TypeScript, or 'skip' if no frontend)", DefaultValue: "JavaScript"},
		{Key: "frontend_framework", Prompt: "Frontend framework (e.g., React, Vue, Angular, Vanilla)", DefaultValue: "React"},
		{Key: "frontend_build_tool", Prompt: "Frontend build tool (e.g., Webpack, Vite, Parcel, Create React App)", DefaultValue: "Vite"},
	}
}

// backendQuestions returns questions for backend configuration
func backendQuestions() []input.Question {
	return []input.Question{
		{Key: "backend_language", Prompt: "Backend language (e.g., Go, Python, Java, Node.js, or 'skip' if no backend)", DefaultValue: "Go"},
		{Key: "backend_framework", Prompt: "Backend framework (e.g., Express, Flask, Spring, Gin, FastAPI)", DefaultValue: "None"},
		{Key: "backend_database", Prompt: "Primary database (e.g., PostgreSQL, MongoDB, MySQL, SQLite)", DefaultValue: "PostgreSQL"},
		{Key: "api_rules", Prompt: "API design rules (e.g., RESTful, GraphQL standards, versioning strategy)", DefaultValue: "RESTful API design"},
	}
}

// testingQuestions returns questions for testing configuration
func testingQuestions() []input.Question {
	return []input.Question{
		{Key: "testing_framework", Prompt: "Primary testing framework (e.g., Jest, pytest, JUnit, Go testing)", DefaultValue: "Jest"},
		{Key: "testing_strategy", Prompt: "Testing strategy focus (e.g., Unit tests, Integration tests, E2E, TDD)", DefaultValue: "Unit and Integration tests"},
	}
}

// agentsQuestions returns questions for agent configuration
func agentsQuestions() []input.Question {
	return []input.Question{
		{Key: "enable_agents", Prompt: "Enable agents (role-based AI assistants)? (yes/no/skip)", DefaultValue: "yes"},
		{Key: "agent_architect", Prompt: "Enable architect agent (planning specialist)?", DefaultValue: "yes"},
		{Key: "agent_frontend", Prompt: "Enable frontend engineer agent?", DefaultValue: "yes"},
		{Key: "agent_backend", Prompt: "Enable backend engineer agent?", DefaultValue: "yes"},
		{Key: "agent_code_reviewer", Prompt: "Enable code reviewer agent?", DefaultValue: "yes"},
		{Key: "agent_technical_writer", Prompt: "Enable technical writer agent (documentation)?", DefaultValue: "yes"},
		{Key: "agent_devops", Prompt: "Enable DevOps engineer agent (infrastructure)?", DefaultValue: "no"},
	}
}

// promptsQuestions returns questions for prompt template configuration
func promptsQuestions() []input.Question {
	return []input.Question{
		{Key: "enable_prompts", Prompt: "Enable prompt templates (reusable workflows)? (yes/no/skip)", DefaultValue: "yes"},
		{Key: "prompt_code_review", Prompt: "Enable code review prompt template?", DefaultValue: "yes"},
		{Key: "prompt_feature_spec", Prompt: "Enable feature specification prompt template?", DefaultValue: "yes"},
		{Key: "prompt_refactor", Prompt: "Enable refactor prompt template?", DefaultValue: "yes"},
		{Key: "prompt_bug_fix", Prompt: "Enable bug fix prompt template?", DefaultValue: "yes"},
		{Key: "prompt_pr_description", Prompt: "Enable PR description prompt template?", DefaultValue: "yes"},
	}
}

// specsQuestions returns questions for spec template configuration
func specsQuestions() []input.Question {
	return []input.Question{
		{Key: "enable_specs", Prompt: "Enable specification templates (project planning)? (yes/no/skip)", DefaultValue: "yes"},
		{Key: "spec_feature_template", Prompt: "Enable feature template spec?", DefaultValue: "yes"},
		{Key: "spec_api_endpoint", Prompt: "Enable API endpoint spec?", DefaultValue: "yes"},
		{Key: "spec_component", Prompt: "Enable component spec (frontend)?", DefaultValue: "yes"},
	}
}
