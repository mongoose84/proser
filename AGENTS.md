# PROSER - PROSE Framework Setup and Resource Generator

**Purpose**: Interactive CLI tool that automates GitHub Copilot PROSE file generation for projects

## Project Overview

PROSER collects project information via interactive prompts and generates customized PROSE configuration files (copilot-instructions.md, .instructions.md files, agent definitions) tailored to your technology stack and requirements.

### Architecture Philosophy
- **Generator-based**: Modular generators implement a common interface
- **Filesystem abstraction**: Enables testing without disk I/O
- **Data-driven**: Language/framework metadata in registries, not hardcoded
- **SOLID principles**: Clean separation of concerns, testability-first

## Repository Structure

```
proser/
├── main.go              # CLI entry point and orchestration
├── config/              # Project configuration structs
├── input/               # User input collection interfaces
├── project/             # Project type definitions (fullstack, backend-only, etc.)
├── generator/           # PROSE file generators (instructions, agents, specs)
├── language/            # Language and framework metadata registry
├── filesystem/          # FS abstraction (real + in-memory for testing)
└── .github/
    ├── agents/          # Custom agent definitions (.agent.md files)
    └── instructions/    # Domain-specific instruction files
```

### Key Directories

| Directory | Responsibility |
|-----------|---------------|
| `generator/` | Implements file generators for each PROSE primitive type |
| `config/` | Project configuration models and answer collection |
| `language/` | Metadata about languages, frameworks, testing tools |
| `project/` | Defines which generators run for project types |
| `filesystem/` | OS and memory filesystem implementations |

## Tech Stack

- **Language**: Go (idiomatic Go patterns, Effective Go conventions)
- **Testing**: Go standard testing package with table-driven tests
- **Build**: Native Go toolchain (go build, go test)
- **Formatting**: gofmt

## Development Guidelines

### Code Conventions
- Follow [backend instructions](.github/instructions/backend.instructions.md) for Go-specific patterns
- Follow [testing instructions](.github/instructions/testing.instructions.md) for test patterns
- Use interfaces to define behavior contracts
- Implement proper error wrapping with context
- Maintain clean separation: input collection, config management, generation logic

### Testing Strategy
- Unit tests for all generators
- In-memory filesystem for deterministic tests
- Table-driven test patterns
- Coverage for error paths and edge cases

### Building and Running
```bash
# Build
go build -o proser

# Run (current directory)
./proser

# Run (target directory)
./proser /path/to/project

# Run tests
go test ./...

# Format code
gofmt -w .
```

## Custom Agents

This project uses specialized agents for different development workflows:

| Agent | Purpose | Tools |
|-------|---------|-------|
| [architect.agent.md](.github/agents/architect.agent.md) | System architecture and planning | codebase, search, problems |
| [backend-engineer.agent.md](.github/agents/backend-engineer.agent.md) | Go backend implementation | Full development toolset |
| [code-reviewer.agent.md](.github/agents/code-reviewer.agent.md) | Code quality and review | codebase, search, problems |

### When to Use Which Agent
- **Architect**: Design decisions, planning, architecture evaluation
- **Backend Engineer**: Implementing features, fixing bugs, refactoring
- **Code Reviewer**: Pre-merge review, testing coverage, quality gates

## Instructions Hierarchy

```
.github/
└── instructions/
    ├── backend.instructions.md     # Applies to **/*.go
    └── testing.instructions.md     # Applies to **/*_test.go
```

Domain-specific instructions inherit from [copilot-instructions.md](.github/copilot-instructions.md) at the root.

## Agent Boundaries

### ✅ Agents CAN
- Read any file in the repository
- Search codebase semantically or with grep
- Analyze architecture and patterns
- Generate new code following conventions
- Run tests and builds
- Create new files in appropriate locations

### ❌ Agents SHOULD NOT
- Modify go.mod dependencies without explicit request
- Change core interfaces without architectural review
- Skip test coverage for new generators
- Bypass filesystem abstraction in generator code
- Create files outside project directory

## Security and Safety

- **Input validation**: All user inputs validated before use
- **Path safety**: Resolves to absolute paths, validates directories exist
- **No remote execution**: Purely local file generation
- **Read-only analysis**: Agents can analyze but require confirmation for destructive changes

## Progressive Disclosure

When exploring this codebase:

1. **Start with**: [README.md](README.md) for high-level features and architecture
2. **Understand the data model**: [config/config.go](config/config.go) defines project configuration structure
3. **Explore generation**: [generator/generator.go](generator/generator.go) shows the Generator interface
4. **See implementations**: Individual generator files show how each PROSE file type is created
5. **Review testing**: `*_test.go` files demonstrate testing patterns

## Common Tasks

### Adding a New Generator
1. Create new file in `generator/` implementing `Generator` interface
2. Add to appropriate project type in `project/types.go`
3. Create table-driven tests using memory filesystem
4. Update this AGENTS.md with new generator purpose

### Adding a New Language/Framework
1. Update `language/languages.go` with metadata
2. Ensure generator templates handle new options
3. Add test cases for new language paths

### Adding a New Project Type
1. Define in `project/types.go`
2. Specify which generators apply
3. Add questions in `project/questions.go` if needed

## References

- [PROSE Specification](https://danielmeppiel.github.io/awesome-ai-native/docs/prose/)
- [GitHub Copilot Documentation](https://docs.github.com/en/copilot)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Testing](https://go.dev/doc/tutorial/add-a-test)

## Context Engineering Notes

This file follows PROSE principles:
- **Progressive Disclosure**: Links to deeper detail rather than overwhelming upfront
- **Explicit Hierarchy**: Root-level guidance, domain-specific in subdirectories
- **Safety Boundaries**: Clear CAN/CANNOT agent boundaries
- **Orchestrated Composition**: Modular structure enables independent agent work
- **Reduced Scope**: Focused sections, right-sized for context windows

---

*This AGENTS.md file provides universal discovery and project-wide context for AI agents working with this codebase.*
