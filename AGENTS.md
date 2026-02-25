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

Domain-specific instructions in [.github/instructions/](.github/instructions/):
- [Backend Guidelines](.github/instructions/backend.instructions.md)
- [Testing Guidelines](.github/instructions/testing.instructions.md)

All inherit from [copilot-instructions.md](.github/copilot-instructions.md).

## Agent Boundaries

### ✅ Agents CAN
- Read any file in the repository
- Search codebase and analyze patterns
- Generate code following conventions
- Run tests and builds

### ❌ Agents SHOULD NOT
- Modify dependencies without explicit request
- Change core architecture without discussion
- Skip test coverage for new features

## Progressive Disclosure

1. Start: [README.md](README.md)
2. Global: [copilot-instructions.md](.github/copilot-instructions.md)
3. Domain-specific: [.github/instructions/](.github/instructions/)

## References

- [PROSE Specification](https://danielmeppiel.github.io/awesome-ai-native/docs/prose/)
- [GitHub Copilot Documentation](https://docs.github.com/en/copilot)

## Context Engineering

This file follows PROSE principles:
- **Progressive Disclosure**: Links to detail rather than overwhelming upfront
- **Reduced Scope**: Focused, essential content only
- **Explicit Hierarchy**: Root guidance, domain-specific in subdirectories

---

*This AGENTS.md file provides universal discovery and project-wide context for AI agents working with this codebase.*
