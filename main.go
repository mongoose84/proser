package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mongoose84/proser/config"
)

func main() {
	fmt.Println("===========================================")
	fmt.Println("PROSER - PROSE Framework Setup Tool")
	fmt.Println("===========================================")
	fmt.Println()

	// Parse target path argument
	targetPath := "."
	if len(os.Args) > 1 {
		targetPath = os.Args[1]
	}

	// Resolve to absolute path
	absTarget, err := filepath.Abs(targetPath)
	if err != nil {
		fmt.Printf("❌ Error resolving target path: %v\n", err)
		os.Exit(1)
	}

	// Verify target path exists and is a directory
	fileInfo, err := os.Stat(absTarget)
	if err != nil {
		fmt.Printf("❌ Error accessing target path: %v\n", err)
		os.Exit(1)
	}
	if !fileInfo.IsDir() {
		fmt.Printf("❌ Target path is not a directory: %s\n", absTarget)
		os.Exit(1)
	}

	fmt.Printf("📁 Target directory: %s\n\n", absTarget)

	config := collectUserInput()

	fmt.Println("\n📝 Generating files based on your configuration...")

	if err := setupGithubFolder(absTarget, config); err != nil {
		fmt.Printf("❌ Error setting up .github folder: %v\n", err)
		os.Exit(1)
	}

	if err := createAgentMdFiles(absTarget, config); err != nil {
		fmt.Printf("❌ Error creating AGENT.md files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n✅ Setup complete!")
	fmt.Println("📁 Files created in .github/")
	fmt.Println("📄 AGENT.md files created in subdirectories (3 levels deep)")
	fmt.Println("\n🎉 Your project is now configured for PROSE framework!")
}

func collectUserInput() config.ProjectConfig {
	reader := bufio.NewReader(os.Stdin)
	cfg := config.ProjectConfig{
		General: config.GeneralConfig{},
		Testing: config.TestingConfig{},
	}

	fmt.Println("Please answer the following questions about your project:")
	fmt.Println()

	// General project information
	fmt.Println("=== General Project Information ===")
	cfg.General.ProjectName = promptUser(reader, "Project name", "my-project")
	cfg.General.Description = promptUser(reader, "Project description", "A software project")
	cfg.General.CodeStyle = promptUser(reader, "General code style guidelines (e.g., follow PEP8, use gofmt, ESLint rules)", "Follow standard formatting")
	cfg.General.Security = promptUser(reader, "Security requirements (e.g., authentication methods, data encryption, OWASP compliance)", "Follow OWASP top 10")
	cfg.General.CustomRules = promptUser(reader, "Additional custom rules or guidelines", "None")

	// Frontend-specific questions
	fmt.Println()
	fmt.Println("=== Frontend Configuration ===")
	frontendLanguage := promptUser(reader, "Frontend language (e.g., JavaScript, TypeScript, or 'skip' if no frontend)", "JavaScript")
	if strings.ToLower(frontendLanguage) != "skip" && frontendLanguage != "" {
		cfg.Frontend = &config.FrontendConfig{
			Language:  frontendLanguage,
			Framework: promptUser(reader, "Frontend framework (e.g., React, Vue, Angular, Vanilla)", "React"),
			BuildTool: promptUser(reader, "Frontend build tool (e.g., Webpack, Vite, Parcel, Create React App)", "Vite"),
		}
	}

	// Backend-specific questions
	fmt.Println()
	fmt.Println("=== Backend Configuration ===")
	backendLanguage := promptUser(reader, "Backend language (e.g., Go, Python, Java, Node.js, or 'skip' if no backend)", "Go")
	if strings.ToLower(backendLanguage) != "skip" && backendLanguage != "" {
		cfg.Backend = &config.BackendConfig{
			Language:  backendLanguage,
			Framework: promptUser(reader, "Backend framework (e.g., Express, Flask, Spring, Gin, FastAPI)", "None"),
			Database:  promptUser(reader, "Primary database (e.g., PostgreSQL, MongoDB, MySQL, SQLite)", "PostgreSQL"),
			APIRules:  promptUser(reader, "API design rules (e.g., RESTful, GraphQL standards, versioning strategy)", "RESTful API design"),
		}
	}

	// Testing-specific questions
	fmt.Println()
	fmt.Println("=== Testing Configuration ===")
	cfg.Testing.Framework = promptUser(reader, "Primary testing framework (e.g., Jest, pytest, JUnit, Go testing)", "Jest")
	cfg.Testing.Strategy = promptUser(reader, "Testing strategy focus (e.g., Unit tests, Integration tests, E2E, TDD)", "Unit and Integration tests")

	fmt.Println("\n📋 Configuration Summary:")
	fmt.Printf("  Project: %s\n", cfg.General.ProjectName)
	if cfg.General.Description != "" {
		fmt.Printf("  Description: %s\n", cfg.General.Description)
	}
	if cfg.HasFrontend() {
		fmt.Printf("  Frontend: %s", cfg.Frontend.Language)
		if cfg.Frontend.Framework != "" {
			fmt.Printf(" with %s", cfg.Frontend.Framework)
		}
		if cfg.Frontend.BuildTool != "" {
			fmt.Printf(" (%s)", cfg.Frontend.BuildTool)
		}
		fmt.Println()
	}
	if cfg.HasBackend() {
		fmt.Printf("  Backend: %s", cfg.Backend.Language)
		if cfg.Backend.Framework != "" && cfg.Backend.Framework != "None" {
			fmt.Printf(" with %s", cfg.Backend.Framework)
		}
		if cfg.Backend.Database != "" {
			fmt.Printf(" + %s", cfg.Backend.Database)
		}
		fmt.Println()
	}
	if cfg.Testing.Framework != "" {
		fmt.Printf("  Testing: %s (%s)\n", cfg.Testing.Framework, cfg.Testing.Strategy)
	}
	if cfg.General.CodeStyle != "" {
		fmt.Printf("  Code Style: %s\n", cfg.General.CodeStyle)
	}
	if cfg.HasBackend() && cfg.Backend.APIRules != "" {
		fmt.Printf("  API Rules: %s\n", cfg.Backend.APIRules)
	}
	if cfg.General.Security != "" {
		fmt.Printf("  Security: %s\n", cfg.General.Security)
	}
	if cfg.General.CustomRules != "" && cfg.General.CustomRules != "None" {
		fmt.Printf("  Custom Rules: %s\n", cfg.General.CustomRules)
	}

	return cfg
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
