# PROSER

**PRO**SE framework **SE**tup and **R**esource generator

Automate setting up GitHub Copilot PROSE files for your projects with interactive configuration.

## What is PROSER?

PROSER is a command-line tool that helps you set up GitHub Copilot Workspace (Based on this brilliant documentation: https://danielmeppiel.github.io/awesome-ai-native/docs/prose/) It interactively collects information about your project and generates customized configuration files tailored to your technology stack and requirements.

## Features

- 🎯 **Interactive Setup**: Asks relevant questions based on your project type
- 📝 **Customized Templates**: Generates project-specific configuration files
- 🏗️ **Multi-Project Types**: Supports fullstack, frontend-only, and backend-only projects
- 🤖 **Agent Instructions**: Creates AGENT.md files in subdirectories (up to 3 levels deep)
- 🔧 **GitHub Integration**: Sets up .github folder with Copilot instructions
- 🧩 **Extensible Architecture**: Easy to add new generators and project types
- ✅ **SOLID Principles**: Clean, maintainable, and testable codebase

## Architecture

PROSER follows a modular, generator-based architecture:

- **Generator Interface**: All PROSE file types implement a common interface
- **FileSystem Abstraction**: Enables testing without touching disk
- **Language Registry**: Data-driven language and framework metadata
- **Project Types**: Define which generators run for different project types
- **Input Collection**: Abstract user input for testability

### Project Structure

```
proser/
├── main.go                   # Entry point and orchestrator
├── config/                   # Project configuration
├── input/                    # Input collection interfaces
├── project/                  # Project type definitions
├── generator/                # Generator implementations
│   ├── copilot_instructions.go
│   ├── frontend_instructions.go
│   ├── backend_instructions.go
│   ├── testing_instructions.go
│   ├── agent_md.go
│   └── future.go            # Stub generators for future file types
├── language/                 # Language & framework registry
├── filesystem/               # Filesystem abstraction
└── template/                 # Shared template helpers
```

## Installation

### From Source

```bash
git clone https://github.com/mongoose84/proser.git
cd proser
go build -o proser
```

### Using Go Install

```bash
go install github.com/mongoose84/proser@latest
```

> **Note**: After installing, if you get `command not found`, you need to add Go's bin directory to your PATH.

#### Troubleshooting: Command Not Found

If `proser` command is not found after installation, your Go bin directory is not in your PATH. Here are the solutions:

**Option 1: Add Go bin to PATH (Recommended)**

Find your Go bin directory:
```bash
go env GOPATH
```

Then add the bin directory to your PATH. Add this line to your shell config file (`~/.bashrc`, `~/.zshrc`, or `~/.profile`):

```bash
# For bash/zsh
export PATH="$PATH:$(go env GOPATH)/bin"
```

Then reload your shell:
```bash
source ~/.bashrc  # or ~/.zshrc or ~/.profile
```

**Option 2: Run directly from GOPATH**

```bash
$(go env GOPATH)/bin/proser
```

**Option 3: Use `go run` (Development)**

```bash
# From a local clone
cd /path/to/proser
go run main.go
```

## Usage

Run PROSER in your project directory (or specify a target path):

```bash
# Run in current directory
proser

# Run in a specific directory
proser /path/to/your/project

# Show help
proser --help
```

### Interactive Prompts

PROSER will ask you to select a project type and then collect information specific to that type:

**Project Types:**
- **Fullstack**: Frontend + Backend application
- **Frontend**: Frontend-only application
- **Backend**: Backend/API service only

**Questions include:**
- Project name and description
- Code style guidelines
- Security requirements
- Frontend language, framework, and build tool (if applicable)
- Backend language, framework, and database (if applicable)
- Testing framework and strategy

### Generated Files

PROSER creates the following files:

- `.github/copilot-instructions.md` - Global repository instructions
- `.github/instructions/frontend.instructions.md` - Frontend-specific guidelines (if applicable)
- `.github/instructions/backend.instructions.md` - Backend-specific guidelines (if applicable)
- `.github/instructions/testing.instructions.md` - Testing guidelines
- `AGENT.md` files in subdirectories (automatically skips node_modules, vendor, bin, etc.)

## Example

```bash
$ proser
===========================================
PROSER - PROSE File Setup Tool
===========================================

📁 Target directory: /home/user/awesome-api

Select project type:
  1. fullstack - Full-stack application with frontend and backend
  2. frontend - Frontend application only
  3. backend - Backend/API service only

Enter project type (fullstack/frontend/backend) [fullstack]: backend

Please answer the following questions about your project:

Project name [my-project]: awesome-api
Project description [A software project]: REST API for awesome service
General code style guidelines [Follow standard formatting]: Use gofmt and golangci-lint
Security requirements [Follow OWASP top 10]: JWT auth, TLS 1.3, OWASP compliance
Additional custom rules or guidelines [None]: Use structured logging
Backend language [Go]: Go
Backend framework [None]: Gin
Primary database [PostgreSQL]: PostgreSQL
API design rules [RESTful API design]: RESTful with versioned endpoints
Primary testing framework [Jest]: Go testing
Testing strategy focus [Unit and Integration tests]: Table-driven unit tests and integration tests

📋 Configuration Summary:
  Project: awesome-api
  Description: REST API for awesome service
  Backend: Go with Gin + PostgreSQL
  Testing: Go testing (Table-driven unit tests and integration tests)
  Code Style: Use gofmt and golangci-lint
  API Rules: RESTful with versioned endpoints
  Security: JWT auth, TLS 1.3, OWASP compliance
  Custom Rules: Use structured logging

📝 Generating files based on your configuration...
  ✓ Generated copilot-instructions files
  ✓ Generated backend-instructions files
  ✓ Generated testing-instructions files
  ✓ Generated agent-md files

✅ Setup complete!
📁 Files created in .github/
📄 AGENT.md files created in subdirectories

🎉 Your project is now configured for PROSE Architectural Style for AI-Native Development!
```

## Supported Languages and Frameworks

### Languages
- Go
- Python
- Java
- JavaScript/Node.js
- TypeScript
- Rust
- C#

### Frontend Frameworks
- React
- Vue
- Angular
- Vanilla JS/TS

### Testing Frameworks
- Jest (JavaScript/TypeScript)
- pytest (Python)
- JUnit (Java)
- Go testing (Go)

## Extending PROSER

### Adding a New Language

Add a new language definition in `language/languages.go`:

```go
r.RegisterLanguage(&LanguageInfo{
    Name:           "ruby",
    Aliases:        []string{"rb"},
    FileExtensions: []string{".rb"},
    Guidelines: []string{
        "- Follow Ruby best practices and idioms",
        "- Use proper exception handling",
    },
    BestPractices: []string{
        "- Follow Ruby style guide",
        "- Use RuboCop for linting",
    },
    TestingPatterns: []string{
        "- [ ] RSpec test suites",
    },
})
```

### Adding a New Project Type

Create a new type in `project/types.go`:

```go
type MobileProject struct{}

func (p *MobileProject) Name() string {
    return "mobile"
}

func (p *MobileProject) Description() string {
    return "Mobile application (iOS/Android)"
}

func (p *MobileProject) Generators() []generator.Generator {
    return []generator.Generator{
        &generator.CopilotInstructionsGenerator{},
        &generator.MobileInstructionsGenerator{}, // Create this
        &generator.AgentMdGenerator{},
    }
}

func (p *MobileProject) Questions() []input.Question {
    return append(generalQuestions(), mobileQuestions()...)
}
```

Register it in `project/project_type.go`:

```go
func init() {
    Register(&MobileProject{})
}
```

### Adding a New Generator

Create a new generator in `generator/`:

```go
type PromptsGenerator struct{}

func (g *PromptsGenerator) Name() string {
    return "prompts"
}

func (g *PromptsGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
    files := make(map[string]string)
    // Generate .github/prompts/*.prompt.md files
    files[".github/prompts/refactor.prompt.md"] = "# Refactoring Prompt\n..."
    return files, nil
}
```

## Contributing

Contributions are welcome! Please ensure:

1. Code follows Go best practices
2. New features include tests
3. Architecture principles are maintained (SOLID, DRY, etc.)
4. Documentation is updated

## License

MIT License - See LICENSE file for details

## Credits

Inspired by the GitHub Copilot Workspace PROSE Architectural Style for AI-Native Development described here: https://danielmeppiel.github.io/awesome-ai-native/docs/prose/

## Generated Files

### `.github/copilot-instructions.md`
Main instructions for GitHub Copilot containing:
- Project overview
- Language and code style guidelines
- API design rules
- Security requirements
- Code review checklist

### `.github/agents/general-instructions.md`
Detailed agent instructions including:
- Project context
- Agent responsibilities
- Communication style
- Specific guidelines for code generation, review, and documentation

### `.github/workflows/ci.yml`
Language-specific CI workflow template supporting:
- Go
- Python
- JavaScript/TypeScript/Node.js
- Generic template for other languages

### `AGENT.md` files
Created in each subdirectory (up to 4 levels deep) with:
- Directory-specific context
- Guidelines relevant to that location
- Testing and documentation requirements

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - feel free to use this tool in your projects!

## Support

If you encounter any issues or have questions, please open an issue on GitHub.
