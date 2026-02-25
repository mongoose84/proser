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
$ proser /path/to/my-project

Select project type:
  1. fullstack - Full-stack application with frontend and backend
  2. frontend - Frontend application only  
  3. backend - Backend/API service only

Enter project type: backend

Project name: awesome-api
Backend language: Go
Backend framework: Gin
Primary database: PostgreSQL
Testing framework: Go testing

✓ Generated .github/copilot-instructions.md
✓ Generated .github/instructions/backend.instructions.md
✓ Generated .github/instructions/testing.instructions.md
✓ Generated AGENTS.md

✅ Setup complete!
```

## Supported Technologies

**Languages**: Go, Python, Java, JavaScript/TypeScript, Rust, C#

**Frontend**: React, Vue, Angular, Svelte

**Testing**: Go testing, Jest, pytest, JUnit

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

## Generated Files

PROSER creates a complete PROSE framework structure:

- `.github/copilot-instructions.md` - Root instructions with project overview, tech stack, code style
- `.github/instructions/*.instructions.md` - Domain-specific guidelines (backend, frontend, testing)
- `.github/agents/*.agent.md` - Specialized agent definitions (optional)
- `.github/prompts/*.prompt.md` - Workflow templates (optional)
- `.github/specs/*.spec.md` - Specification templates (optional)
- `AGENTS.md` - Project discovery file at root

## Contributing

Contributions welcome! Ensure:
- Code follows Go best practices
- New features include tests
- Documentation is updated

## License

MIT License

## Credits

Inspired by [PROSE Architectural Style](https://danielmeppiel.github.io/awesome-ai-native/docs/prose/) for AI-Native Development.
