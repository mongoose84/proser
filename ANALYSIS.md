# PROSER Capability Analysis

## Overview

This document analyzes what PROSER currently generates versus the full set of GitHub Copilot
customization files a software project can leverage. The goal is to identify gaps and prioritize
what to build next.

---

## 1. Current State — What PROSER Generates Today

### 1.1 Implemented Generators (actively producing output)

| Generator | Output File(s) | Status |
|---|---|---|
| `CopilotInstructionsGenerator` | `.github/copilot-instructions.md` | **Complete** |
| `FrontendInstructionsGenerator` | `.github/instructions/frontend.instructions.md` | **Complete** |
| `BackendInstructionsGenerator` | `.github/instructions/backend.instructions.md` | **Complete** |
| `TestingInstructionsGenerator` | `.github/instructions/testing.instructions.md` | **Complete** |
| `AgentMdGenerator` | `<subdir>/AGENT.md` (up to 3 levels deep) | **Complete** |

### 1.2 Stub Generators (defined in `generator/future.go`, return empty output)

| Generator | Intended Output | Status |
|---|---|---|
| `PromptsGenerator` | `.github/prompts/*.prompt.md` | **Stub — not implemented** |
| `ChatModesGenerator` | `.github/chatmodes/*.chatmode.md` | **Stub — not implemented** |
| `SpecsGenerator` | `.github/specs/*.spec.md` | **Stub — not implemented** |
| `SkillsGenerator` | `.github/skills/*.skill.md` | **Stub — not implemented** |

> None of the stub generators are wired into any `ProjectType.Generators()` list, so they
> never execute even if they returned content.

### 1.3 Project Types

| Type | Generators Used |
|---|---|
| `fullstack` | copilot-instructions, frontend-instructions, backend-instructions, testing-instructions, agent-md |
| `frontend` | copilot-instructions, frontend-instructions, testing-instructions, agent-md |
| `backend` | copilot-instructions, backend-instructions, testing-instructions, agent-md |

### 1.4 Language / Framework Coverage

**Languages**: Go, Python, Java, JavaScript, TypeScript, Rust, C#

**Frontend Frameworks**: React, Vue, Angular, Vanilla

**Testing Frameworks**: Jest, pytest, JUnit, Go testing

**Backend Frameworks**: Referenced in prompts (Gin, Flask, Spring, Express, FastAPI) but the
language registry only stores framework `Guidelines` — no deep framework-specific template
content is generated beyond a single bullet point.

### 1.5 Input Collection

All input is gathered through sequential text prompts via `InteractiveCollector`. There is no
YAML/JSON config file import, no auto-detection from existing project files (e.g., reading
`package.json`, `go.mod`), and no `--non-interactive` flag.

### 1.6 Architecture Strengths

- Clean `Generator` interface: `Name() string` + `Generate(ctx) (map[string]string, error)`
- `FileSystem` abstraction with `MemoryFileSystem` for testing
- Data-driven `language.Registry` with `LanguageInfo` and `FrameworkInfo`
- `ProjectType` registry pattern — easy to add new types
- `Writer` orchestrates generation and file writing uniformly

---

## 2. Desired State — Full GitHub Copilot Customization Surface

The GitHub Copilot PROSE framework supports the following file types and locations:

### 2.1 Instructions (scoped context)

```
.github/
├── copilot-instructions.md                   # Global repo-wide rules
└── instructions/
    ├── frontend.instructions.md              # applyTo: "**/*.{jsx,tsx,css}"
    ├── backend.instructions.md               # applyTo: "**/*.{py,go,java}"
    ├── testing.instructions.md               # applyTo: "**/test/**"
    ├── docs.instructions.md                  # applyTo: "**/*.md"
    ├── devops.instructions.md                # applyTo: "**/*.{yml,yaml,Dockerfile}"
    └── database.instructions.md              # applyTo: "**/*.sql"
```

**Current coverage**: copilot-instructions ✅, frontend ✅, backend ✅, testing ✅  
**Missing**: docs, devops/infra, database, and any custom domain-specific instruction files.

### 2.2 Chat Modes (persona-based agents)

```
.github/chatmodes/
├── architect.chatmode.md              # Planning specialist
├── frontend-engineer.chatmode.md      # UI specialist
├── backend-engineer.chatmode.md       # API specialist
├── technical-writer.chatmode.md       # Documentation specialist
├── code-reviewer.chatmode.md          # Review specialist
└── devops-engineer.chatmode.md        # Infrastructure specialist
```

**Current coverage**: Stub generator exists but produces nothing. Not wired into any project type.

### 2.3 Prompts (reusable prompt templates)

```
.github/prompts/
├── code-review.prompt.md              # Structured review workflow
├── feature-spec.prompt.md             # Spec-first methodology
├── async-implementation.prompt.md     # GitHub Coding Agent delegation
├── refactor.prompt.md                 # Refactoring workflow
├── bug-fix.prompt.md                  # Bug investigation template
└── pr-description.prompt.md           # PR body generation
```

**Current coverage**: Stub generator exists but produces nothing. Not wired into any project type.

### 2.4 Specifications (project planning templates)

```
.github/specs/
├── feature-template.spec.md           # Standard feature specification
├── api-endpoint.spec.md               # API-specific specification
├── component.spec.md                  # UI component specification
└── migration.spec.md                  # Database/system migration spec
```

**Current coverage**: Stub generator exists but produces nothing. Not wired into any project type.

### 2.5 Nested Agent Files (directory-scoped context)

```
src/
├── AGENT.md
├── components/
│   └── AGENT.md
├── services/
│   └── AGENT.md
└── utils/
    └── AGENT.md
```

**Current coverage**: ✅ Implemented — traverses up to 3 levels, skips build/vendor dirs.  
**Gaps**: Content is generic. Could be smarter by inferring directory purpose from file contents
(e.g., a dir full of `*_test.go` gets testing-focused AGENT.md).

### 2.6 Skills (tool definitions)

```
.github/skills/
├── database-query.skill.md            # SQL generation skill
├── api-client.skill.md                # HTTP client generation
└── test-data.skill.md                 # Test fixture generation
```

**Current coverage**: Stub generator exists but produces nothing. Not wired. Skills are a newer
Copilot concept and may not be widely adopted yet — lower priority.

---

## 3. Gap Analysis

### 3.1 Feature Gap Matrix

| Feature | Current | Desired | Priority | Effort |
|---|---|---|---|---|
| **Copilot Instructions** | ✅ Full | ✅ | — | — |
| **Frontend Instructions** | ✅ Full | ✅ | — | — |
| **Backend Instructions** | ✅ Full | ✅ | — | — |
| **Testing Instructions** | ✅ Full | ✅ | — | — |
| **Agent MD files** | ✅ Generic | Context-aware | Medium | Medium |
| **Chat Modes** | ❌ Stub | 4-6 role-based modes | **High** | Medium |
| **Prompt Templates** | ❌ Stub | 4-6 reusable prompts | **High** | Medium |
| **Spec Templates** | ❌ Stub | 3-4 spec templates | Medium | Low |
| **Skills** | ❌ Stub | 2-3 skill definitions | Low | Low |
| **Docs Instructions** | ❌ Missing | `.instructions.md` for docs | Medium | Low |
| **DevOps Instructions** | ❌ Missing | `.instructions.md` for CI/CD | Medium | Low |
| **Database Instructions** | ❌ Missing | `.instructions.md` for SQL | Low | Low |

### 3.2 Infrastructure / UX Gaps

| Gap | Description | Priority | Effort |
|---|---|---|---|
| **Auto-detection** | Scan project files (`package.json`, `go.mod`, `Cargo.toml`) to pre-fill config instead of asking everything interactively | **High** | Medium |
| **Config file import** | Support a `.proser.yml` or `.proser.json` so teams can commit and share their config, enabling `proser --config .proser.yml` | **High** | Medium |
| **Non-interactive mode** | `--non-interactive` flag using defaults or config file for CI pipelines | Medium | Low |
| **Incremental updates** | Detect existing `.github/` files and merge/update instead of overwrite | Medium | High |
| **Dry-run mode** | `--dry-run` to preview generated files without writing | Medium | Low |
| **Language registry usage** | The `language.Registry` is defined but never used by generators — generators hardcode language-specific content via switch statements instead of looking up `LanguageInfo` | **High** | Medium |
| **Framework registry usage** | `FrameworkInfo` is registered but never consumed — framework-specific instructions are hardcoded in generators | **High** | Medium |
| **Template engine** | No Go template (`text/template`) usage — all content is built with `strings.Builder`. Templates would make customization and contribution easier | Medium | High |
| **Tests** | No test files exist in the repository | **High** | High |
| **Custom project types** | Users cannot define their own project types beyond fullstack/frontend/backend (e.g., mobile, CLI, library, monorepo) | Medium | Medium |

### 3.3 Content Quality Gaps

| Area | Current State | Improvement |
|---|---|---|
| **Instructions frontmatter** | ✅ Has `applyTo` and `description` | Could add `priority`, `version` |
| **AGENT.md content** | Generic boilerplate regardless of directory | Should infer purpose from filenames/contents |
| **Chat mode structure** | N/A | Needs `role`, `constraints`, `tools` sections per Copilot spec |
| **Prompt structure** | N/A | Needs `mode`, `tools`, `description` frontmatter per Copilot spec |
| **applyTo patterns** | Basic language globs | Could be more precise (e.g., exclude test files from backend instructions) |

---

## 4. Recommended Implementation Roadmap

### Phase 1 — High-Value, Medium-Effort (immediate)

1. **Implement `ChatModesGenerator`** — Generate 4-6 project-aware chat modes:
   - `architect.chatmode.md` — planning, cannot execute code
   - `frontend-engineer.chatmode.md` — UI work scoped to frontend files
   - `backend-engineer.chatmode.md` — API work scoped to backend files
   - `technical-writer.chatmode.md` — documentation only
   - Wire into all three project types with appropriate subset

2. **Implement `PromptsGenerator`** — Generate 3-4 reusable prompt templates:
   - `code-review.prompt.md` — structured review with validation checkpoints
   - `feature-spec.prompt.md` — spec-first methodology
   - `async-implementation.prompt.md` — GitHub Coding Agent delegation
   - Wire into all project types

3. **Wire language registry into generators** — Replace hardcoded switch statements with
   `Registry.LookupLanguage()` / `Registry.LookupFramework()` lookups. This makes adding new
   languages a data-only change rather than touching every generator.

4. **Add unit tests** — At minimum: generator output tests, config parsing tests, registry
   lookup tests. The `MemoryFileSystem` already exists for this purpose.

### Phase 2 — Usability Improvements (next)

5. **Auto-detection** — Scan the target directory for `package.json`, `go.mod`, `Cargo.toml`,
   `pom.xml`, etc. and pre-populate `ProjectConfig` fields. Fall back to interactive prompts
   for anything not detected.

6. **Implement `SpecsGenerator`** — Generate spec templates:
   - `feature-template.spec.md`
   - `api-endpoint.spec.md` (backend projects)
   - `component.spec.md` (frontend projects)

7. **Config file support** — Read/write `.proser.yml` so teams can version-control their setup.

8. **Dry-run mode** — `--dry-run` flag that prints file paths and content without writing.

### Phase 3 — Polish and Extend (later)

9. **Additional instruction files** — `docs.instructions.md`, `devops.instructions.md`,
   `database.instructions.md`

10. **Smart AGENT.md** — Analyze directory contents to generate context-aware agent instructions
    rather than generic boilerplate.

11. **Template engine migration** — Move content generation from `strings.Builder` to
    `text/template` files in the `template/` directory for easier customization.

12. **Custom project types** — Allow users to define project types via config file.

13. **Incremental updates** — Detect existing files and offer merge strategies.

14. **Skills generator** — Lower priority; implement when the skills spec stabilizes.

---

## 5. File Tree — Current vs Target

```
Current Output                          Target Output
─────────────────────────────           ─────────────────────────────────────────
.github/                                .github/
├── copilot-instructions.md      ✅     ├── copilot-instructions.md
└── instructions/                       ├── instructions/
    ├── frontend.instructions.md ✅     │   ├── frontend.instructions.md
    ├── backend.instructions.md  ✅     │   ├── backend.instructions.md
    └── testing.instructions.md  ✅     │   ├── testing.instructions.md
                                        │   ├── docs.instructions.md          ❌
                                        │   ├── devops.instructions.md        ❌
                                        │   └── database.instructions.md      ❌
                                        ├── chatmodes/
                                        │   ├── architect.chatmode.md         ❌
                                        │   ├── frontend-engineer.chatmode.md ❌
                                        │   ├── backend-engineer.chatmode.md  ❌
                                        │   └── technical-writer.chatmode.md  ❌
                                        ├── prompts/
                                        │   ├── code-review.prompt.md         ❌
                                        │   ├── feature-spec.prompt.md        ❌
                                        │   └── async-implementation.prompt.md❌
                                        ├── specs/
                                        │   ├── feature-template.spec.md      ❌
                                        │   ├── api-endpoint.spec.md          ❌
                                        │   └── component.spec.md             ❌
                                        └── skills/                           ❌
                                            └── ...
<subdir>/AGENT.md                ✅     <subdir>/AGENT.md (context-aware)
```

---

## 6. Summary

PROSER has a solid architectural foundation — the `Generator` interface, filesystem abstraction,
language registry, and project-type registry are all well-designed extension points. The core
instruction generators (copilot, frontend, backend, testing) and the AGENT.md walker are
complete and functional.

The biggest gaps are:

1. **Four stub generators** that exist structurally but produce no output (chatmodes, prompts,
   specs, skills)
2. **Language/framework registry is unused** — generators hardcode language-specific content
   instead of consulting the registry
3. **No project auto-detection** — everything is manual interactive input
4. **No tests** — despite having a `MemoryFileSystem` designed for testability
5. **No template engine** — the `template/` directory is empty; all content is `strings.Builder`

The path to a complete Copilot customization generator requires implementing the four stub
generators (prioritizing chatmodes and prompts), wiring the registry into generators, and
adding auto-detection + config file support for better UX.
