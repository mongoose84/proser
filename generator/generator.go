package generator

import (
	"github.com/mongoose84/proser/config"
	"github.com/mongoose84/proser/filesystem"
)

// GenerateContext provides generators with everything they need
type GenerateContext struct {
	Config     config.ProjectConfig
	TargetPath string // absolute path to the target project root
	FS         filesystem.FileSystem
}

// Generator interface for all PROSE file type generators
type Generator interface {
	// Name returns a human-readable name for this generator
	Name() string

	// Generate produces all files for this generator type.
	// Returns a map of relative file paths to their content.
	Generate(ctx GenerateContext) (map[string]string, error)
}
