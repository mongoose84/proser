package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ProjectConfig struct {
	Language    string
	CodeStyle   string
	APIRules    string
	Security    string
	ProjectName string
	Description string
	CustomRules string
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

	config.ProjectName = promptUser(reader, "Project name", "my-project")
	config.Description = promptUser(reader, "Project description", "A software project")
	config.Language = promptUser(reader, "Primary programming language (e.g., Go, Python, JavaScript, Java)", "Go")
	config.CodeStyle = promptUser(reader, "Code style guidelines (e.g., follow PEP8, use gofmt, ESLint rules)", "Follow standard formatting")
	config.APIRules = promptUser(reader, "API design rules (e.g., RESTful, GraphQL standards, versioning strategy)", "RESTful API design")
	config.Security = promptUser(reader, "Security requirements (e.g., authentication methods, data encryption, OWASP compliance)", "Follow OWASP top 10")
	config.CustomRules = promptUser(reader, "Additional custom rules or guidelines", "None")

	fmt.Println("\n📋 Configuration Summary:")
	fmt.Printf("  Project: %s\n", config.ProjectName)
	if config.Language != "" {
		fmt.Printf("  Language: %s\n", config.Language)
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
	if config.CustomRules != "" {
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
