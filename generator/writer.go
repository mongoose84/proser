package generator

import (
	"fmt"
	"path/filepath"

	"github.com/mongoose84/proser/filesystem"
)

// Writer writes generated files to the filesystem
type Writer struct {
	FS filesystem.FileSystem
}

// NewWriter creates a new Writer
func NewWriter(fs filesystem.FileSystem) *Writer {
	return &Writer{FS: fs}
}

// WriteFiles writes the generated files from a generator to the filesystem
func (w *Writer) WriteFiles(targetPath string, files map[string]string) error {
	for relPath, content := range files {
		fullPath := filepath.Join(targetPath, relPath)

		// Create parent directory
		dir := filepath.Dir(fullPath)
		if err := w.FS.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		// Write file
		if err := w.FS.WriteFile(fullPath, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", fullPath, err)
		}
	}
	return nil
}

// RunGenerator executes a generator and writes its output
func (w *Writer) RunGenerator(gen Generator, ctx GenerateContext) error {
	files, err := gen.Generate(ctx)
	if err != nil {
		return fmt.Errorf("generator %s failed: %w", gen.Name(), err)
	}

	if err := w.WriteFiles(ctx.TargetPath, files); err != nil {
		return fmt.Errorf("failed to write files for generator %s: %w", gen.Name(), err)
	}

	return nil
}
