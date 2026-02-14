package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
	fmt.Println("📄 AGENT.md files created in subdirectories (4 levels deep)")
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

func setupGithubFolder(config ProjectConfig) error {
	githubDir := ".github"
	if err := os.MkdirAll(githubDir, 0755); err != nil {
		return fmt.Errorf("failed to create .github directory: %w", err)
	}

	// Create copilot-instructions.md
	if err := createCopilotInstructions(githubDir, config); err != nil {
		return err
	}

	// Create agent instructions
	agentsDir := filepath.Join(githubDir, "agents")
	if err := os.MkdirAll(agentsDir, 0755); err != nil {
		return fmt.Errorf("failed to create agents directory: %w", err)
	}

	if err := createAgentInstructions(agentsDir, config); err != nil {
		return err
	}

	// Create workflows directory with example
	workflowsDir := filepath.Join(githubDir, "workflows")
	if err := os.MkdirAll(workflowsDir, 0755); err != nil {
		return fmt.Errorf("failed to create workflows directory: %w", err)
	}

	if err := createExampleWorkflow(workflowsDir, config); err != nil {
		return err
	}

	return nil
}

func createCopilotInstructions(githubDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# GitHub Copilot Instructions for %s\n\n", config.ProjectName))

	if config.Description != "" {
		sb.WriteString("## Project Overview\n")
		sb.WriteString(config.Description + "\n\n")
	}

	if config.Language != "" {
		sb.WriteString("## Programming Language\n")
		sb.WriteString(config.Language + "\n\n")
	}

	if config.CodeStyle != "" {
		sb.WriteString("## Code Style Guidelines\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	if config.APIRules != "" {
		sb.WriteString("## API Design Rules\n")
		sb.WriteString(config.APIRules + "\n\n")
	}

	if config.Security != "" {
		sb.WriteString("## Security Requirements\n")
		sb.WriteString(config.Security + "\n\n")
	}

	if config.CustomRules != "" {
		sb.WriteString("## Additional Guidelines\n")
		sb.WriteString(config.CustomRules + "\n\n")
	}

	sb.WriteString("## General Instructions for AI Assistants\n\n")
	sb.WriteString("When working on this project, please:\n\n")

	i := 1
	if config.Language != "" {
		sb.WriteString(fmt.Sprintf("%d. **Follow Language Best Practices**: Write idiomatic %s code following community standards\n", i, config.Language))
		i++
	}
	if config.CodeStyle != "" {
		sb.WriteString(fmt.Sprintf("%d. **Maintain Code Style**: Adhere to the specified code style guidelines: %s\n", i, config.CodeStyle))
		i++
	}
	if config.APIRules != "" {
		sb.WriteString(fmt.Sprintf("%d. **API Consistency**: Follow these API rules: %s\n", i, config.APIRules))
		i++
	}
	if config.Security != "" {
		sb.WriteString(fmt.Sprintf("%d. **Security First**: Always consider security implications: %s\n", i, config.Security))
		i++
	}
	sb.WriteString(fmt.Sprintf("%d. **Documentation**: Keep documentation up-to-date with code changes\n", i))
	i++
	sb.WriteString(fmt.Sprintf("%d. **Testing**: Write tests for new functionality and bug fixes\n", i))
	i++
	sb.WriteString(fmt.Sprintf("%d. **Error Handling**: Implement proper error handling and logging\n", i))
	i++
	sb.WriteString(fmt.Sprintf("%d. **Performance**: Consider performance implications of changes\n", i))

	sb.WriteString("\n## Code Review Checklist\n\n")
	sb.WriteString("Before submitting changes, ensure:\n")
	if config.CodeStyle != "" {
		sb.WriteString("- [ ] Code follows style guidelines\n")
	}
	if config.Security != "" {
		sb.WriteString("- [ ] Security requirements are met\n")
	}
	if config.APIRules != "" {
		sb.WriteString("- [ ] API changes are documented\n")
	}
	sb.WriteString("- [ ] Tests are included and passing\n")
	sb.WriteString("- [ ] Error handling is appropriate\n")
	sb.WriteString("- [ ] Documentation is updated\n")

	filePath := filepath.Join(githubDir, "copilot-instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func createAgentInstructions(agentsDir string, config ProjectConfig) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Agent Instructions for %s\n\n", config.ProjectName))
	sb.WriteString("## Context\n")
	if config.Language != "" {
		sb.WriteString(fmt.Sprintf("This is a %s project with specific guidelines that all agents should follow.\n\n", config.Language))
	} else {
		sb.WriteString("This project has specific guidelines that all agents should follow.\n\n")
	}

	sb.WriteString("## Project Configuration\n\n")

	if config.Language != "" {
		sb.WriteString(fmt.Sprintf("### Language: %s\n", config.Language))
		sb.WriteString(fmt.Sprintf("All code should be written in %s following best practices for this language.\n\n", config.Language))
	}

	if config.CodeStyle != "" {
		sb.WriteString("### Code Style\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	if config.APIRules != "" {
		sb.WriteString("### API Design\n")
		sb.WriteString(config.APIRules + "\n\n")
	}

	if config.Security != "" {
		sb.WriteString("### Security\n")
		sb.WriteString(config.Security + "\n\n")
	}

	if config.CustomRules != "" {
		sb.WriteString("### Custom Rules\n")
		sb.WriteString(config.CustomRules + "\n\n")
	}

	sb.WriteString(`## Agent Responsibilities

### Code Generation
- Generate code that follows the language and style guidelines
- Include appropriate error handling
- Add inline comments for complex logic
- Consider edge cases

### Code Review
- Check for adherence to style guidelines
- Verify security requirements are met
- Ensure API consistency
- Validate test coverage

### Documentation
- Keep README and documentation in sync with code
- Document API endpoints and parameters
- Include usage examples

### Testing
- Write unit tests for new functions
- Include integration tests where appropriate
- Ensure tests are maintainable

## Communication Style
- Be clear and concise
- Explain technical decisions
- Highlight security considerations
- Suggest improvements when appropriate
`)

	filePath := filepath.Join(agentsDir, "general-instructions.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}

func createExampleWorkflow(workflowsDir string, config ProjectConfig) error {
	var workflowContent string

	// Create a simple CI workflow based on the language
	switch strings.ToLower(config.Language) {
	case "go", "golang":
		workflowContent = `name: Go CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.21'
    
    - name: Build
      run: go build -v ./...
    
    - name: Test
      run: go test -v ./...
    
    - name: Lint
      run: |
        go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
        golangci-lint run ./...
`
	case "python":
		workflowContent = `name: Python CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Python
      uses: actions/setup-python@v4
      with:
        python-version: '3.11'
    
    - name: Install dependencies
      run: |
        python -m pip install --upgrade pip
        pip install pytest pylint
        if [ -f requirements.txt ]; then pip install -r requirements.txt; fi
    
    - name: Lint
      run: pylint **/*.py
    
    - name: Test
      run: pytest
`
	case "javascript", "typescript", "node", "nodejs":
		workflowContent = `name: Node.js CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18'
    
    - name: Install dependencies
      run: npm ci
    
    - name: Lint
      run: npm run lint
    
    - name: Test
      run: npm test
    
    - name: Build
      run: npm run build
`
	default:
		workflowContent = fmt.Sprintf(`name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    # Add your build, test, and lint steps here for %s
`, config.Language)
	}

	filePath := filepath.Join(workflowsDir, "ci.yml")
	return os.WriteFile(filePath, []byte(workflowContent), 0644)
}

func createAgentMdFiles(config ProjectConfig) error {
	fmt.Println("\n📂 Creating AGENT.md files in directory structure...")

	// Get all directories in the current path, up to 4 levels deep
	dirs, err := findDirectories(".", 4)
	if err != nil {
		return fmt.Errorf("failed to find directories: %w", err)
	}

	if len(dirs) == 0 {
		fmt.Println("ℹ️  No subdirectories found. Skipping AGENT.md creation.")
		return nil
	}

	for _, dir := range dirs {
		if err := createAgentMdInDirectory(dir, config); err != nil {
			fmt.Printf("⚠️  Warning: failed to create AGENT.md in %s: %v\n", dir, err)
		} else {
			fmt.Printf("  ✓ Created AGENT.md in %s\n", dir)
		}
	}

	return nil
}

func findDirectories(root string, maxDepth int) ([]string, error) {
	var dirs []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden directories (but not the root directory itself)
		if info.IsDir() && path != root && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}

		if info.IsDir() && path != root {
			// Calculate depth
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				// If we can't determine relative path, skip this directory
				return filepath.SkipDir
			}
			depth := strings.Count(relPath, string(os.PathSeparator)) + 1

			if depth <= maxDepth {
				dirs = append(dirs, path)
			} else {
				return filepath.SkipDir
			}
		}

		return nil
	})

	return dirs, err
}

func createAgentMdInDirectory(dir string, config ProjectConfig) error {
	dirName := filepath.Base(dir)
	relPath := dir

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Agent Instructions for %s\n\n", dirName))
	sb.WriteString("## Directory Context\n")
	sb.WriteString(fmt.Sprintf("This directory is part of the %s project located at: %s\n\n", config.ProjectName, relPath))
	sb.WriteString("## Purpose\n")
	sb.WriteString(fmt.Sprintf("This directory contains code and resources related to: %s\n\n", dirName))

	sb.WriteString("## Guidelines\n\n")

	if config.CodeStyle != "" {
		sb.WriteString("### Code Style\n")
		sb.WriteString(config.CodeStyle + "\n\n")
	}

	if config.Language != "" {
		sb.WriteString("### Language\n")
		sb.WriteString(fmt.Sprintf("All code in this directory should be written in %s.\n\n", config.Language))
	}

	if config.APIRules != "" {
		sb.WriteString("### API Rules\n")
		sb.WriteString(config.APIRules + "\n\n")
	}

	if config.Security != "" {
		sb.WriteString("### Security\n")
		sb.WriteString(config.Security + "\n\n")
	}

	sb.WriteString(`## Working in This Directory

When making changes in this directory:

1. Understand the purpose and scope of files here
2. Follow the project-wide guidelines specified in /.github/copilot-instructions.md
3. Ensure changes are consistent with the parent directory structure
4. Update documentation if you modify public interfaces
5. Add or update tests as needed
6. Consider the impact on dependent modules

## Testing
- Tests for this directory should be located nearby
- Run tests before committing changes
- Ensure new functionality is covered by tests

## Documentation
Keep inline documentation up-to-date and clear.
`)

	filePath := filepath.Join(dir, "AGENT.md")
	return os.WriteFile(filePath, []byte(sb.String()), 0644)
}
