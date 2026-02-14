package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mongoose84/proser/config"
	"github.com/mongoose84/proser/filesystem"
)

func createAgentMdFiles(targetPath string, cfg config.ProjectConfig, fs filesystem.FileSystem) error {
	fmt.Println("\n📂 Creating AGENT.md files in directory structure...")

	// Get all directories in the target path, up to 3 levels deep
	dirs, err := findDirectories(targetPath, 3, fs)
	if err != nil {
		return fmt.Errorf("failed to find directories: %w", err)
	}

	if len(dirs) == 0 {
		fmt.Println("ℹ️  No subdirectories found. Skipping AGENT.md creation.")
		return nil
	}

	for _, dir := range dirs {
		if err := createAgentMdInDirectory(dir, cfg, fs); err != nil {
			fmt.Printf("⚠️  Warning: failed to create AGENT.md in %s: %v\n", dir, err)
		} else {
			fmt.Printf("  ✓ Created AGENT.md in %s\n", dir)
		}
	}

	return nil
}

func findDirectories(root string, maxDepth int, fs filesystem.FileSystem) ([]string, error) {
	var dirs []string

	err := fs.Walk(root, func(path string, info os.FileInfo, err error) error {
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

func createAgentMdInDirectory(dir string, cfg config.ProjectConfig, fs filesystem.FileSystem) error {
	dirName := filepath.Base(dir)
	relPath := dir

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

	filePath := filepath.Join(dir, "AGENT.md")
	return fs.WriteFile(filePath, []byte(sb.String()), 0644)
}
