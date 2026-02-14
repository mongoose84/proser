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
