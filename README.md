# PROSER

**PRO**SE framework **SE**tup and **R**esource generator

Automate how to add all the PROSE framework files to your projects with interactive configuration.

## What is PROSER?

PROSER is a command-line tool that helps you set up GitHub Copilot Workspace (PROSE framework) configuration files for your projects. It interactively collects information about your project and generates customized configuration files tailored to your needs.

## Features

- 🎯 **Interactive Setup**: Asks relevant questions about your project
- 📝 **Customized Templates**: Generates project-specific configuration files
- 🤖 **Agent Instructions**: Creates AGENT.md files in subdirectories (up to 4 levels deep)
- 🔧 **GitHub Integration**: Sets up .github folder with Copilot instructions
- 🚀 **CI/CD Ready**: Includes example workflow files based on your language

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

## Usage

1. Navigate to your project directory:
```bash
cd /path/to/your/project
```

2. Run PROSER:
```bash
proser
```

3. Answer the interactive prompts about your project:
   - Project name
   - Project description
   - Primary programming language
   - Code style guidelines
   - API design rules
   - Security requirements
   - Additional custom rules

4. PROSER will generate:
   - `.github/copilot-instructions.md` - Main Copilot instructions
   - `.github/agents/general-instructions.md` - Agent-specific instructions
   - `.github/workflows/ci.yml` - Example CI workflow for your language
   - `AGENT.md` files in subdirectories (up to 4 levels deep)

## Example

```bash
$ proser
===========================================
PROSER - PROSE Framework Setup Tool
===========================================

Please answer the following questions about your project:

Project name [my-project]: awesome-api
Project description [A software project]: REST API for awesome service
Primary programming language (e.g., Go, Python, JavaScript, Java) [Go]: Go
Code style guidelines (e.g., follow PEP8, use gofmt, ESLint rules) [Follow standard formatting]: Use gofmt and golangci-lint
API design rules (e.g., RESTful, GraphQL standards, versioning strategy) [RESTful API design]: RESTful with versioned endpoints
Security requirements (e.g., authentication methods, data encryption, OWASP compliance) [Follow OWASP top 10]: JWT auth, TLS 1.3, OWASP compliance
Additional custom rules or guidelines [None]: Use structured logging

📋 Configuration Summary:
  Project: awesome-api
  Language: Go
  Code Style: Use gofmt and golangci-lint
  API Rules: RESTful with versioned endpoints
  Security: JWT auth, TLS 1.3, OWASP compliance

📝 Generating files based on your configuration...

📂 Creating AGENT.md files in directory structure...
  ✓ Created AGENT.md in src
  ✓ Created AGENT.md in src/core
  ✓ Created AGENT.md in src/core/api
  ✓ Created AGENT.md in src/core/api/handlers

✅ Setup complete!
📁 Files created in .github/
📄 AGENT.md files created in subdirectories (4 levels deep)

🎉 Your project is now configured for PROSE framework!
```

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
