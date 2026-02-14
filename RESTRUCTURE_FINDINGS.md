# PROSER Restructure Findings & Implementation Plan

## 1. Current State Summary

PROSER is a Go CLI tool that generates GitHub Copilot PROSE framework files (`.github/` instructions, `AGENT.md` files) based on interactive user input. The codebase consists of 4 files in a flat structure:

| File | Responsibility |
|---|---|
| `main.go` | `ProjectConfig` struct, user input collection, orchestration |
| `github_setup.go` | `.github/` folder creation, `copilot-instructions.md` generation |
| `instructions.go` | Frontend, backend, and testing `.instructions.md` generation |
| `agent_md.go` | `AGENT.md` file generation in subdirectories |

---

## 2. Critical Bug

**The code does not compile.** `config.Language` is referenced in `instructions.go` (lines 107, 108, 267, 268, 397, 398) and `agent_md.go` (lines 89, 91), but the `ProjectConfig` struct has no `Language` field — only `FrontendLanguage` and `BackendLanguage`. The existing `proser` binary is stale. This must be fixed before any restructure.

---

## 3. Analysis

### 3.1 Maintainability — Poor

- **Monolithic `ProjectConfig`**: A single flat struct with 13 fields mixes unrelated concerns (general, frontend, backend, testing). Adding a new project type (e.g., client+server) means adding more fields to an already bloated struct.
- **Giant switch statements**: Language- and framework-specific logic is embedded as enormous switch/case blocks inside content-generation functions (`createBackendInstructions` is ~170 lines). Modifying one language's output risks breaking another.
- **String building mixed with logic**: Business logic (which sections to include) and template content (actual markdown text) are interleaved in the same functions using `strings.Builder`. There is no separation of data from rendering.
- **Flat file structure**: All Go code lives in one package at the root. There is no logical grouping.

### 3.2 Testability — Very Poor

- **No tests exist** anywhere in the project.
- **Direct filesystem I/O**: Functions call `os.WriteFile` and `os.MkdirAll` directly, making unit testing impossible without real filesystem operations.
- **Direct stdin I/O**: `collectUserInput()` reads from `os.Stdin` via `bufio.NewReader(os.Stdin)` — completely untestable without stdin mocking.
- **Functions produce side effects only**: Generator functions accept a config and write files. They return nothing testable. There is no way to inspect the generated content without reading back from disk.
- **No interfaces**: No abstraction boundary exists that could be mocked or stubbed.

### 3.3 Reuse of Code — Poor

- **Duplicated switch statements**: The same language-matching switch (go/python/java/js/ts) is repeated in `createFrontendInstructions`, `createBackendInstructions`, `createTestingInstructions`, and `createAgentMdInDirectory` — at least 4 places.
- **Duplicated section patterns**: "Code Style", "Security", "Structured Output" sections are generated with nearly identical code in multiple functions.
- **No shared abstractions**: No template engine, no builder abstraction, no shared section renderers.

### 3.4 Engineering Best Practices — Needs Improvement

- **No CI/CD pipeline** for the tool itself.
- **No input validation**: User input is accepted as-is with no validation.
- **Inconsistent error handling**: Some errors are returned, some are printed with `fmt.Printf`. No error wrapping with context in several places.
- **No structured logging**: Only `fmt.Println`.
- **README is out of sync with code**: README mentions `Language` field, workflows, and agents — features that don't match the current code.

### 3.5 SOLID Principles — Multiple Violations

| Principle | Status | Issue |
|---|---|---|
| **Single Responsibility** | Violated | `main.go` owns config definition, user input, and orchestration. `instructions.go` generates 3 separate file types. |
| **Open/Closed** | Severely Violated | Adding a new project type, a new language, or a new PROSE file type (prompts, chatmodes, specs, skills) requires modifying existing functions and the config struct. Nothing is extensible without modification. |
| **Liskov Substitution** | N/A | No interfaces or inheritance exist. |
| **Interface Segregation** | Violated | The monolithic `ProjectConfig` is passed to every function, even though each only uses a subset of fields. |
| **Dependency Inversion** | Violated | High-level orchestration depends directly on low-level details (`os.WriteFile`, `os.Stdin`, `filepath.Walk`). |

---

## 4. Restructure Proposal

### 4.1 New Directory Structure

```
proser/
├── main.go                          # Entry point only — parse flags, wire dependencies, run
├── go.mod
├── README.md
│
├── config/
│   ├── config.go                    # ProjectConfig struct (composed of sub-configs)
│   └── config_test.go
│
├── input/
│   ├── collector.go                 # InputCollector interface + interactive implementation
│   └── collector_test.go
│
├── project/
│   ├── project_type.go              # ProjectType interface + registry
│   ├── frontend.go                  # Frontend project type
│   ├── backend.go                   # Backend project type
│   ├── fullstack.go                 # Fullstack (frontend + backend) project type
│   ├── client_server.go             # Client + Server project type (future)
│   └── project_type_test.go
│
├── generator/
│   ├── generator.go                 # Generator interface + base helpers
│   ├── instructions.go              # InstructionsGenerator (frontend, backend, testing .instructions.md)
│   ├── copilot_instructions.go      # CopilotInstructionsGenerator (copilot-instructions.md)
│   ├── agent_md.go                  # AgentMdGenerator (AGENT.md files)
│   ├── prompts.go                   # PromptsGenerator (future: .prompt.md files)
│   ├── chatmodes.go                 # ChatModesGenerator (future: .chatmode.md files)
│   ├── specs.go                     # SpecsGenerator (future: .spec.md files)
│   ├── skills.go                    # SkillsGenerator (future: .skill.md files)
│   └── generator_test.go
│
├── language/
│   ├── registry.go                  # Language/framework metadata registry (data-driven)
│   ├── languages.go                 # Built-in language definitions (Go, Python, JS, etc.)
│   └── registry_test.go
│
├── filesystem/
│   ├── fs.go                        # FileSystem interface (Write, MkdirAll, Walk)
│   ├── os_fs.go                     # Real OS implementation
│   ├── memory_fs.go                 # In-memory implementation for testing
│   └── fs_test.go
│
└── template/
    ├── renderer.go                  # Template rendering helpers
    ├── sections.go                  # Shared section builders (code style, security, etc.)
    └── renderer_test.go
```

### 4.2 Key Interfaces

#### 4.2.1 FileSystem Interface

Abstracts all filesystem operations so generators can be tested without touching disk.

```go
// filesystem/fs.go
package filesystem

type FileSystem interface {
    WriteFile(path string, data []byte, perm os.FileMode) error
    MkdirAll(path string, perm os.FileMode) error
    Walk(root string, fn filepath.WalkFunc) error
}
```

Provide two implementations:
- `OsFileSystem` — wraps real `os` calls (production use)
- `MemoryFileSystem` — stores files in a `map[string][]byte` (testing use)

#### 4.2.2 Generator Interface

Each PROSE file type (instructions, prompts, chatmodes, specs, skills) implements this interface.

```go
// generator/generator.go
package generator

// GenerateContext provides generators with everything they need:
// the project config, the target path, and filesystem access for
// generators that need to discover the target project's structure.
type GenerateContext struct {
    Config     config.ProjectConfig
    TargetPath string              // absolute path to the target project root
    FS         filesystem.FileSystem
}

type Generator interface {
    // Name returns a human-readable name for this generator (e.g., "instructions", "prompts")
    Name() string

    // Generate produces all files for this generator type.
    // Returns a map of relative file paths to their content.
    Generate(ctx GenerateContext) (map[string]string, error)
}
```

**Why `GenerateContext` includes the filesystem**: Most generators (instructions, prompts, etc.) produce a static set of files and only use `ctx.Config`. But `AgentMdGenerator` needs to **walk the target project's directory tree at runtime** to discover which directories exist — the output files depend on the target's structure, not just the config. Passing `ctx.FS` lets it call `Walk()` to enumerate directories dynamically. For testing, inject a `MemoryFileSystem` pre-populated with a fake directory tree.

**Critical design choice**: `Generate` returns a `map[string]string` (path → content) instead of writing directly to the filesystem. This separates content generation from I/O, making generators trivially testable by inspecting the returned map. A separate `Writer` component handles the actual file writing.

#### 4.2.3 ProjectType Interface

Defines what generators apply to a given project type and how input is collected.

```go
// project/project_type.go
package project

type ProjectType interface {
    // Name returns the project type identifier (e.g., "frontend", "backend", "fullstack", "client-server")
    Name() string

    // Description returns a user-facing description
    Description() string

    // Generators returns the set of generators that should run for this project type
    Generators() []generator.Generator

    // Questions returns the input questions specific to this project type
    Questions() []input.Question
}
```

New project types (e.g., `client-server`) are added by implementing this interface and registering them — no modification of existing code required.

#### 4.2.4 InputCollector Interface

Abstracts user input for testability.

```go
// input/collector.go
package input

type Question struct {
    Key          string
    Prompt       string
    DefaultValue string
}

type InputCollector interface {
    // Collect asks the user all questions and returns answers keyed by Question.Key
    Collect(questions []Question) (map[string]string, error)
}
```

Provide:
- `InteractiveCollector` — reads from a provided `io.Reader` (default `os.Stdin`)
- For tests, pass a `strings.Reader` with predefined answers

#### 4.2.5 Language Registry (Data-Driven)

Replaces all duplicated switch statements with a lookup-based system.

```go
// language/registry.go
package language

type LanguageInfo struct {
    Name            string
    FileExtensions  []string            // e.g., []string{"go"}
    Guidelines      []string            // language-specific guideline lines
    TestingPatterns []string            // language-specific testing pattern lines
    ContextFiles    []string            // e.g., []string{"go.mod", "main.go"}
    OutputChecklist []string            // structured output checklist items
}

type FrameworkInfo struct {
    Name       string
    Language   string
    Guidelines []string
}

type Registry struct {
    languages  map[string]*LanguageInfo
    frameworks map[string]*FrameworkInfo
}

func (r *Registry) Register(lang *LanguageInfo)
func (r *Registry) RegisterFramework(fw *FrameworkInfo)
func (r *Registry) Lookup(name string) (*LanguageInfo, bool)
func (r *Registry) LookupFramework(name string) (*FrameworkInfo, bool)
```

Each language's data is defined declaratively in `language/languages.go`. Adding a new language means adding a new struct literal — no switch statements to modify.

#### 4.2.6 Directory Skip List

The `AgentMdGenerator` must skip directories that are build artifacts, dependency caches, or otherwise not authored source code. Define a default skip list and allow user override.

```go
// generator/skip.go
package generator

// DefaultSkipDirs contains directory names that should never receive AGENT.md files.
// These are dependency, build output, and cache directories.
var DefaultSkipDirs = map[string]bool{
    // JavaScript / Node.js
    "node_modules": true,
    "bower_components": true,
    ".next": true,
    ".nuxt": true,

    // .NET / C#
    "obj": true,
    "bin": true,

    // Build outputs (general)
    "dist": true,
    "build": true,
    "out": true,
    "target": true,
    "output": true,

    // Go
    "vendor": true,

    // Python
    "__pycache__": true,
    ".venv": true,
    "venv": true,
    "env": true,
    ".eggs": true,
    "*.egg-info": true,

    // Java / JVM
    ".gradle": true,
    ".mvn": true,

    // IDE / editor
    ".idea": true,
    ".vscode": true,

    // Version control
    ".git": true,

    // Coverage / test output
    "coverage": true,
    ".nyc_output": true,

    // Misc
    "tmp": true,
    "temp": true,
    "logs": true,
}

// ShouldSkipDir returns true if the directory name is in the skip list.
func ShouldSkipDir(name string, skipList map[string]bool) bool {
    return skipList[name]
}
```

The `AgentMdGenerator` checks each directory name against this skip list during its `Walk()`. The skip list is a `map[string]bool` so it can be extended by the user (e.g., via a `.proserignore` file or CLI flag in the future).

### 4.3 Revised ProjectConfig (Composed)

```go
// config/config.go
package config

type GeneralConfig struct {
    ProjectName string
    Description string
    CodeStyle   string
    Security    string
    CustomRules string
}

type FrontendConfig struct {
    Language  string
    Framework string
    BuildTool string
}

type BackendConfig struct {
    Language  string
    Framework string
    Database  string
    APIRules  string
}

type TestingConfig struct {
    Framework string
    Strategy  string
}

type ProjectConfig struct {
    General  GeneralConfig
    Frontend *FrontendConfig  // nil if no frontend
    Backend  *BackendConfig   // nil if no backend
    Testing  TestingConfig
}

func (c *ProjectConfig) HasFrontend() bool { return c.Frontend != nil }
func (c *ProjectConfig) HasBackend() bool  { return c.Backend != nil }
```

Using pointers for optional sections (nil = not applicable) is clearer than checking for "skip" strings scattered throughout the code.

### 4.4 Target Path as CLI Argument

The tool must accept an optional target path argument instead of always operating on the current working directory. This is standard CLI practice and enables running proser from anywhere.

```
Usage: proser [options] [target-path]

Arguments:
  target-path    Path to the project to set up (default: current directory)

Options:
  --skip-dirs    Comma-separated list of additional directory names to skip
                 when generating AGENT.md files (added to built-in defaults)
```

If no target path is given, default to `"."` (CWD). Resolve the target path to an absolute path immediately and pass it through the `GenerateContext` to all generators. All generated file paths are relative to this target.

### 4.5 Orchestrator (main.go)

The `main` function becomes a thin wiring layer:

```go
func main() {
    // 0. Parse CLI arguments
    targetPath := "."
    if len(os.Args) > 1 {
        targetPath = os.Args[1]
    }
    absTarget, err := filepath.Abs(targetPath)

    // 1. Create dependencies
    fs := filesystem.NewOsFileSystem()
    langRegistry := language.NewDefaultRegistry()
    collector := input.NewInteractiveCollector(os.Stdin)

    // 2. Let user pick project type
    projectType := project.SelectProjectType(collector)

    // 3. Collect input for that project type
    answers, err := collector.Collect(projectType.Questions())
    cfg := config.FromAnswers(answers)

    // 4. Build the generation context
    ctx := generator.GenerateContext{
        Config:     cfg,
        TargetPath: absTarget,
        FS:         fs,
    }

    // 5. Run all generators for this project type
    for _, gen := range projectType.Generators() {
        files, err := gen.Generate(ctx)
        for relPath, content := range files {
            fullPath := filepath.Join(absTarget, relPath)
            fs.MkdirAll(filepath.Dir(fullPath), 0755)
            fs.WriteFile(fullPath, []byte(content), 0644)
        }
    }
}
```

### 4.6 Adding New PROSE File Types (Prompts, ChatModes, Specs, Skills)

Each new type is a new struct implementing `Generator`:

```go
// generator/prompts.go
type PromptsGenerator struct {
    registry *language.Registry
}

func (g *PromptsGenerator) Name() string { return "prompts" }

func (g *PromptsGenerator) Generate(cfg config.ProjectConfig) (map[string]string, error) {
    files := make(map[string]string)
    // Generate .github/prompts/*.prompt.md files
    // ...
    return files, nil
}
```

Then register it in the relevant `ProjectType.Generators()` list. **No existing code is modified.**

### 4.7 Adding a New Project Type (e.g., Client+Server)

```go
// project/client_server.go
type ClientServerProject struct{}

func (p *ClientServerProject) Name() string        { return "client-server" }
func (p *ClientServerProject) Description() string { return "Client application with a backend server" }

func (p *ClientServerProject) Generators() []generator.Generator {
    return []generator.Generator{
        &generator.CopilotInstructionsGenerator{},
        &generator.FrontendInstructionsGenerator{},
        &generator.BackendInstructionsGenerator{},
        &generator.TestingInstructionsGenerator{},
        &generator.AgentMdGenerator{},
        // future:
        // &generator.PromptsGenerator{},
        // &generator.SkillsGenerator{},
    }
}

func (p *ClientServerProject) Questions() []input.Question {
    // Combines frontend + backend + general questions
    return append(append(
        generalQuestions(),
        frontendQuestions()...),
        backendQuestions()...,
    )
}
```

Register it: `project.Register(&ClientServerProject{})`. **No existing code is modified.**

---

## 5. Implementation Steps (Ordered)

Implement these steps in order. Each step should result in a compilable, working state.

### Step 1: Fix the compilation bug
- In `instructions.go` and `agent_md.go`, replace all references to `config.Language` with the correct field. In `createFrontendInstructions`, use `config.FrontendLanguage`. In `createBackendInstructions`, use `config.BackendLanguage`. In `createTestingInstructions`, use `config.BackendLanguage` (or whichever is set). In `createAgentMdInDirectory`, remove the `Language` block or derive it from `FrontendLanguage`/`BackendLanguage`.
- Verify the project compiles with `go build`.

### Step 2: Create the package structure
- Create directories: `config/`, `input/`, `project/`, `generator/`, `language/`, `filesystem/`, `template/`.
- Move `ProjectConfig` to `config/config.go` and split into sub-structs (`GeneralConfig`, `FrontendConfig`, `BackendConfig`, `TestingConfig`).
- Add `HasFrontend()` and `HasBackend()` helper methods.
- Update all files to use the new config package import.

### Step 3: Add target path CLI argument
- Parse an optional positional argument for the target project path (default: `"."`)
- Resolve it to an absolute path using `filepath.Abs()`.
- Pass the target path through to all functions that currently hardcode `"."` or rely on CWD.
- Verify it works: `go build && ./proser /tmp/test-project`.

### Step 4: Introduce the FileSystem interface
- Create `filesystem/fs.go` with the `FileSystem` interface.
- Create `filesystem/os_fs.go` with the real OS implementation.
- Create `filesystem/memory_fs.go` with the in-memory implementation. The `MemoryFileSystem` must support `Walk()` over its virtual directory tree so `AgentMdGenerator` can be tested against fake project structures.
- Refactor all `os.WriteFile` and `os.MkdirAll` calls to use the interface.

### Step 5: Create the Language Registry
- Create `language/registry.go` with the `Registry`, `LanguageInfo`, and `FrameworkInfo` types.
- Create `language/languages.go` with all current language/framework data (Go, Python, Java, JavaScript/TypeScript, Rust, C#, React, Vue, Angular).
- This data should be extracted from the current switch statements.

### Step 6: Create the Generator interface and migrate existing generators
- Define the `Generator` interface and `GenerateContext` in `generator/generator.go`.
- Create `generator/skip.go` with the `DefaultSkipDirs` map and `ShouldSkipDir()` function.
- Migrate `createFrontendInstructions` → `generator.FrontendInstructionsGenerator`.
- Migrate `createBackendInstructions` → `generator.BackendInstructionsGenerator`.
- Migrate `createTestingInstructions` → `generator.TestingInstructionsGenerator`.
- Migrate `createCopilotInstructions` → `generator.CopilotInstructionsGenerator`.
- Migrate `createAgentMdFiles` → `generator.AgentMdGenerator`.
  - `AgentMdGenerator` must use `ctx.FS.Walk(ctx.TargetPath, ...)` to discover directories dynamically.
  - During the walk, skip any directory whose name is in `DefaultSkipDirs` (in addition to hidden dirs starting with `.`).
  - This ensures it works correctly on any arbitrary project structure without creating files in `node_modules/`, `bin/`, `obj/`, `dist/`, `vendor/`, `__pycache__/`, etc.
- Each generator should use the language registry instead of switch statements.
- Each generator's `Generate()` accepts `GenerateContext` and returns `map[string]string` (path → content).

### Step 7: Create shared template/section helpers
- Create `template/sections.go` with reusable section builders: `CodeStyleSection()`, `SecuritySection()`, `TechStackSection()`, `StructuredOutputSection()`.
- Refactor generators to use these shared builders instead of duplicating code.

### Step 8: Create the InputCollector interface
- Define `InputCollector` and `Question` in `input/collector.go`.
- Implement `InteractiveCollector` that accepts an `io.Reader`.
- Migrate `collectUserInput()` and `promptUser()` into this package.

### Step 9: Create the ProjectType interface and registry
- Define `ProjectType` interface in `project/project_type.go`.
- Implement `FrontendProject`, `BackendProject`, `FullstackProject` types.
- Add a `Register()` function and a `SelectProjectType()` function that lets users pick.
- Each project type returns its applicable generators and questions.

### Step 10: Create stub generators for future PROSE file types
- Create empty generator implementations for: `PromptsGenerator`, `ChatModesGenerator`, `SpecsGenerator`, `SkillsGenerator`.
- Each should implement `Generator` and return an empty map with a TODO comment.
- This proves the architecture supports expansion without modifying existing code.

### Step 11: Rewrite main.go as thin orchestrator
- `main.go` should only: parse CLI args (target path, skip dirs), create dependencies, select project type, collect input, build `GenerateContext`, run generators, write output.
- The target path is resolved to absolute and passed into `GenerateContext`.
- All logic should live in the packages.

### Step 12: Add tests
- `filesystem/fs_test.go` — Test `MemoryFileSystem` operations.
- `language/registry_test.go` — Test registry lookup, missing language fallback.
- `generator/generator_test.go` — Test each generator produces expected file paths and content against `MemoryFileSystem`. Use table-driven tests.
- `input/collector_test.go` — Test `InteractiveCollector` with a `strings.Reader` providing canned answers.
- `config/config_test.go` — Test `FromAnswers()`, `HasFrontend()`, `HasBackend()`.
- `project/project_type_test.go` — Test that each project type returns the correct generator set.

### Step 13: Update README.md
- Reflect the new architecture and supported features.
- Document how to add new project types and generators.

---

## 6. Summary of Principles Applied

| Principle | How It's Achieved |
|---|---|
| **Single Responsibility** | Each package owns one concern. Each generator handles one file type. |
| **Open/Closed** | New project types, generators, and languages are added by creating new files and registering them — no modification of existing code. |
| **Liskov Substitution** | All `Generator` implementations are interchangeable. All `ProjectType` implementations are interchangeable. |
| **Interface Segregation** | Generators receive only the config sub-structs they need via the composed `ProjectConfig`. `FileSystem` interface is minimal. |
| **Dependency Inversion** | Generators depend on `FileSystem` interface, not `os`. Input depends on `io.Reader`, not `os.Stdin`. Orchestration depends on interfaces, not concrete types. |
| **Testability** | Every component can be tested in isolation via interfaces and dependency injection. `MemoryFileSystem` and `strings.Reader` enable pure unit tests. |
| **DRY** | Language registry eliminates duplicated switch statements. Shared template sections eliminate duplicated markdown builders. |
