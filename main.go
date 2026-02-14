package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ProjectConfig struct {
	Language      string
	CodeStyle     string
	APIRules      string
	Security      string
	ProjectName   string
	Description   string
	CustomRules   string
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

	fmt.Println("Please answer the following questions about your project:\n")

	config.ProjectName = promptUser(reader, "Project name", "my-project")
	config.Description = promptUser(reader, "Project description", "A software project")
	config.Language = promptUser(reader, "Primary programming language (e.g., Go, Python, JavaScript, Java)", "Go")
	config.CodeStyle = promptUser(reader, "Code style guidelines (e.g., follow PEP8, use gofmt, ESLint rules)", "Follow standard formatting")
	config.APIRules = promptUser(reader, "API design rules (e.g., RESTful, GraphQL standards, versioning strategy)", "RESTful API design")
	config.Security = promptUser(reader, "Security requirements (e.g., authentication methods, data encryption, OWASP compliance)", "Follow OWASP top 10")
	config.CustomRules = promptUser(reader, "Additional custom rules or guidelines", "None")

	fmt.Println("\n📋 Configuration Summary:")
	fmt.Printf("  Project: %s\n", config.ProjectName)
	fmt.Printf("  Language: %s\n", config.Language)
	fmt.Printf("  Code Style: %s\n", config.CodeStyle)
	fmt.Printf("  API Rules: %s\n", config.APIRules)
	fmt.Printf("  Security: %s\n", config.Security)

	return config
}

func promptUser(reader *bufio.Reader, prompt, defaultValue string) string {
	fmt.Printf("%s [%s]: ", prompt, defaultValue)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
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
	content := fmt.Sprintf(`# GitHub Copilot Instructions for %s

## Project Overview
%s

## Programming Language
%s

## Code Style Guidelines
%s

## API Design Rules
%s

## Security Requirements
%s

## Additional Guidelines
%s

## General Instructions for AI Assistants

When working on this project, please:

1. **Follow Language Best Practices**: Write idiomatic %s code following community standards
2. **Maintain Code Style**: Adhere to the specified code style guidelines: %s
3. **API Consistency**: Follow these API rules: %s
4. **Security First**: Always consider security implications: %s
5. **Documentation**: Keep documentation up-to-date with code changes
6. **Testing**: Write tests for new functionality and bug fixes
7. **Error Handling**: Implement proper error handling and logging
8. **Performance**: Consider performance implications of changes

## Code Review Checklist

Before submitting changes, ensure:
- [ ] Code follows style guidelines
- [ ] Security requirements are met
- [ ] API changes are documented
- [ ] Tests are included and passing
- [ ] Error handling is appropriate
- [ ] Documentation is updated
`, 
		config.ProjectName,
		config.Description,
		config.Language,
		config.CodeStyle,
		config.APIRules,
		config.Security,
		config.CustomRules,
		config.Language,
		config.CodeStyle,
		config.APIRules,
		config.Security,
	)

	filePath := filepath.Join(githubDir, "copilot-instructions.md")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func createAgentInstructions(agentsDir string, config ProjectConfig) error {
	// Create a general agent instruction file
	content := fmt.Sprintf(`# Agent Instructions for %s

## Context
This is a %s project with specific guidelines that all agents should follow.

## Project Configuration

### Language: %s
All code should be written in %s following best practices for this language.

### Code Style
%s

### API Design
%s

### Security
%s

### Custom Rules
%s

## Agent Responsibilities

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
`,
		config.ProjectName,
		config.Language,
		config.Language,
		config.Language,
		config.CodeStyle,
		config.APIRules,
		config.Security,
		config.CustomRules,
	)

	filePath := filepath.Join(agentsDir, "general-instructions.md")
	return os.WriteFile(filePath, []byte(content), 0644)
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
        go install golang.org/x/lint/golint@latest
        golint ./...
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
		fmt.Println("ℹ️  No subdirectories found. Creating sample directory structure...")
		sampleDirs := []string{"src", "src/core", "src/core/api", "src/core/api/handlers"}
		for _, dir := range sampleDirs {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return fmt.Errorf("failed to create sample directory %s: %w", dir, err)
			}
		}
		dirs = sampleDirs
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
			relPath, _ := filepath.Rel(root, path)
			depth := strings.Count(relPath, string(os.PathSeparator)) + 1

			if depth <= maxDepth {
				dirs = append(dirs, path)
			} else if depth > maxDepth {
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

	content := fmt.Sprintf(`# Agent Instructions for %s

## Directory Context
This directory is part of the %s project located at: %s

## Purpose
This directory contains code and resources related to: %s

## Guidelines

### Code Style
%s

### Language
All code in this directory should be written in %s.

### API Rules
%s

### Security
%s

## Working in This Directory

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
`,
		dirName,
		config.ProjectName,
		relPath,
		dirName,
		config.CodeStyle,
		config.Language,
		config.APIRules,
		config.Security,
	)

	filePath := filepath.Join(dir, "AGENT.md")
	return os.WriteFile(filePath, []byte(content), 0644)
}
