package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mongoose84/proser/config"
)

// AgentMdGenerator generates AGENT.md files in subdirectories
type AgentMdGenerator struct {
	MaxDepth int             // Maximum directory depth to traverse (default: 3)
	SkipDirs map[string]bool // Directories to skip (default: DefaultSkipDirs)
}

// Name returns the generator name
func (g *AgentMdGenerator) Name() string {
	return "agent-md"
}

// Generate creates AGENT.md files in subdirectories
func (g *AgentMdGenerator) Generate(ctx GenerateContext) (map[string]string, error) {
	files := make(map[string]string)
	cfg := ctx.Config

	// Use defaults if not set
	maxDepth := g.MaxDepth
	if maxDepth == 0 {
		maxDepth = 3
	}
	skipDirs := g.SkipDirs
	if skipDirs == nil {
		skipDirs = DefaultSkipDirs
	}

	// Find all directories up to maxDepth
	dirs, err := g.findDirectories(ctx.TargetPath, maxDepth, skipDirs, ctx.FS)
	if err != nil {
		return nil, fmt.Errorf("failed to find directories: %w", err)
	}

	// Generate AGENT.md for each directory
	for _, dir := range dirs {
		content := g.generateAgentMdContent(dir, ctx.TargetPath, cfg)

		// Create relative path from target
		relPath, err := filepath.Rel(ctx.TargetPath, dir)
		if err != nil {
			continue // Skip if we can't create relative path
		}

		filePath := filepath.Join(relPath, "AGENT.md")
		files[filePath] = content
	}

	return files, nil
}

// findDirectories walks the directory tree and returns directories up to maxDepth
func (g *AgentMdGenerator) findDirectories(root string, maxDepth int, skipDirs map[string]bool, fs interface{}) ([]string, error) {
	var dirs []string

	// Type assert to filesystem walker
	type Walker interface {
		Walk(root string, fn func(path string, info os.FileInfo, err error) error) error
	}

	walker, ok := fs.(Walker)
	if !ok {
		return nil, fmt.Errorf("filesystem does not support Walk")
	}

	err := walker.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == root {
			return nil
		}

		if info.IsDir() {
			// Skip hidden directories
			if strings.HasPrefix(info.Name(), ".") {
				return filepath.SkipDir
			}

			// Skip directories in skip list
			if ShouldSkipDir(info.Name(), skipDirs) {
				return filepath.SkipDir
			}

			// Calculate depth
			relPath, err := filepath.Rel(root, path)
			if err != nil {
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

// generateAgentMdContent creates the AGENT.md content for a directory
func (g *AgentMdGenerator) generateAgentMdContent(dir string, targetPath string, cfg config.ProjectConfig) string {
	dirName := filepath.Base(dir)

	// Create relative path from target
	relPath, err := filepath.Rel(targetPath, dir)
	if err != nil {
		relPath = dir
	}

	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# Agent Instructions for %s\n\n", dirName))
	sb.WriteString("## Directory Context\n")
	sb.WriteString(fmt.Sprintf("This directory is part of the %s project located at: %s\n\n", cfg.General.ProjectName, relPath))
	sb.WriteString("## Purpose\n")
	sb.WriteString(fmt.Sprintf("This directory contains code and resources related to: %s\n\n", dirName))

	sb.WriteString("## Guidelines\n\n")

	if cfg.General.CodeStyle != "" {
		sb.WriteString("### Code Style\n")
		sb.WriteString(cfg.General.CodeStyle + "\n\n")
	}

	if cfg.HasBackend() && cfg.Backend.APIRules != "" {
		sb.WriteString("### API Rules\n")
		sb.WriteString(cfg.Backend.APIRules + "\n\n")
	}

	if cfg.General.Security != "" {
		sb.WriteString("### Security\n")
		sb.WriteString(cfg.General.Security + "\n\n")
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

	return sb.String()
}
