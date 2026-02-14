package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ProjectConfig struct {
	// General project info
	ProjectName string
	Description string
	CodeStyle   string
	Security    string
	CustomRules string

	// Frontend-specific
	FrontendLanguage  string // js, ts, etc.
	FrontendFramework string // React, Vue, Angular, etc.
	FrontendBuildTool string // Webpack, Vite, etc.

	// Backend-specific
	BackendLanguage  string // Go, Python, Java, Node.js, etc.
	BackendFramework string // Express, Flask, Spring, etc.
	BackendDatabase  string // PostgreSQL, MongoDB, etc.
	APIRules         string

	// Testing-specific
	TestingFramework string // Jest, pytest, JUnit, etc.
	TestingStrategy  string // Unit, Integration, E2E focus
}

func main() {
	fmt.Println("===========================================")
	fmt.Println("PROSER - PROSE Framework Setup Tool")
	fmt.Println("===========================================")
	fmt.Println()

	config := collectUserInput()

	fmt.Println("\n📝 Generating files based on your configuration...")

	if err := setupGithubFolder(config); err != nil {
		fmt.Printf("❌ Error setting up .github folder: %v\n", err)
		os.Exit(1)
	}

	if err := createAgentMdFiles(config); err != nil {
		fmt.Printf("❌ Error creating AGENT.md files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Setup complete!")
	fmt.Println("📁 Files created in .github/")
	fmt.Println("📄 AGENT.md files created in subdirectories (3 levels deep)")
	fmt.Println("\n🎉 Your project is now configured for PROSE framework!")
}

func collectUserInput() ProjectConfig {
	reader := bufio.NewReader(os.Stdin)
	config := ProjectConfig{}

	fmt.Println("Please answer the following questions about your project:")
	fmt.Println()

	// General project information
	fmt.Println("=== General Project Information ===")
	config.ProjectName = promptUser(reader, "Project name", "my-project")
	config.Description = promptUser(reader, "Project description", "A software project")
	config.CodeStyle = promptUser(reader, "General code style guidelines (e.g., follow PEP8, use gofmt, ESLint rules)", "Follow standard formatting")
	config.Security = promptUser(reader, "Security requirements (e.g., authentication methods, data encryption, OWASP compliance)", "Follow OWASP top 10")
	config.CustomRules = promptUser(reader, "Additional custom rules or guidelines", "None")

	// Frontend-specific questions
	fmt.Println()
	fmt.Println("=== Frontend Configuration ===")
	config.FrontendLanguage = promptUser(reader, "Frontend language (e.g., JavaScript, TypeScript, or 'skip' if no frontend)", "JavaScript")
	if strings.ToLower(config.FrontendLanguage) != "skip" && config.FrontendLanguage != "" {
		config.FrontendFramework = promptUser(reader, "Frontend framework (e.g., React, Vue, Angular, Vanilla)", "React")
		config.FrontendBuildTool = promptUser(reader, "Frontend build tool (e.g., Webpack, Vite, Parcel, Create React App)", "Vite")
	}

	// Backend-specific questions
	fmt.Println()
	fmt.Println("=== Backend Configuration ===")
	config.BackendLanguage = promptUser(reader, "Backend language (e.g., Go, Python, Java, Node.js, or 'skip' if no backend)", "Go")
	if strings.ToLower(config.BackendLanguage) != "skip" && config.BackendLanguage != "" {
		config.BackendFramework = promptUser(reader, "Backend framework (e.g., Express, Flask, Spring, Gin, FastAPI)", "None")
		config.BackendDatabase = promptUser(reader, "Primary database (e.g., PostgreSQL, MongoDB, MySQL, SQLite)", "PostgreSQL")
		config.APIRules = promptUser(reader, "API design rules (e.g., RESTful, GraphQL standards, versioning strategy)", "RESTful API design")
	}

	// Testing-specific questions
	fmt.Println()
	fmt.Println("=== Testing Configuration ===")
	config.TestingFramework = promptUser(reader, "Primary testing framework (e.g., Jest, pytest, JUnit, Go testing)", "Jest")
	config.TestingStrategy = promptUser(reader, "Testing strategy focus (e.g., Unit tests, Integration tests, E2E, TDD)", "Unit and Integration tests")

	fmt.Println("\n📋 Configuration Summary:")
	fmt.Printf("  Project: %s\n", config.ProjectName)
	if config.Description != "" {
		fmt.Printf("  Description: %s\n", config.Description)
	}
	if config.FrontendLanguage != "" && strings.ToLower(config.FrontendLanguage) != "skip" {
		fmt.Printf("  Frontend: %s", config.FrontendLanguage)
		if config.FrontendFramework != "" {
			fmt.Printf(" with %s", config.FrontendFramework)
		}
		if config.FrontendBuildTool != "" {
			fmt.Printf(" (%s)", config.FrontendBuildTool)
		}
		fmt.Println()
	}
	if config.BackendLanguage != "" && strings.ToLower(config.BackendLanguage) != "skip" {
		fmt.Printf("  Backend: %s", config.BackendLanguage)
		if config.BackendFramework != "" && config.BackendFramework != "None" {
			fmt.Printf(" with %s", config.BackendFramework)
		}
		if config.BackendDatabase != "" {
			fmt.Printf(" + %s", config.BackendDatabase)
		}
		fmt.Println()
	}
	if config.TestingFramework != "" {
		fmt.Printf("  Testing: %s (%s)\n", config.TestingFramework, config.TestingStrategy)
	}
	if config.CodeStyle != "" {
		fmt.Printf("  Code Style: %s\n", config.CodeStyle)
	}
	if config.APIRules != "" {
		fmt.Printf("  API Rules: %s\n", config.APIRules)
	}
	if config.Security != "" {
		fmt.Printf("  Security: %s\n", config.Security)
	}
	if config.CustomRules != "" && config.CustomRules != "None" {
		fmt.Printf("  Custom Rules: %s\n", config.CustomRules)
	}

	return config
}

func promptUser(reader *bufio.Reader, prompt, defaultValue string) string {
	fmt.Printf("%s [%s] (type 'skip' to omit): ", prompt, defaultValue)
	input, err := reader.ReadString('\n')
	if err != nil {
		// In case of read error, return default value
		return defaultValue
	}
	input = strings.TrimSpace(input)
	if strings.ToLower(input) == "skip" {
		return ""
	}
	if input == "" {
		return defaultValue
	}
	return input
}
