package generator

// DefaultSkipDirs contains directory names that should never receive AGENT.md files.
// These are dependency, build output, and cache directories.
var DefaultSkipDirs = map[string]bool{
	// JavaScript / Node.js
	"node_modules":     true,
	"bower_components": true,
	".next":            true,
	".nuxt":            true,

	// .NET / C#
	"obj": true,
	"bin": true,

	// Build outputs (general)
	"dist":   true,
	"build":  true,
	"out":    true,
	"target": true,
	"output": true,

	// Go
	"vendor": true,

	// Python
	"__pycache__":   true,
	".venv":         true,
	"venv":          true,
	"env":           true,
	".eggs":         true,
	"*.egg-info":    true,
	".pytest_cache": true,

	// Java / JVM
	".gradle": true,
	".mvn":    true,

	// IDE / editor
	".idea":   true,
	".vscode": true,

	// Version control
	".git": true,

	// Coverage / test output
	"coverage":    true,
	".nyc_output": true,

	// Misc
	"tmp":  true,
	"temp": true,
	"logs": true,
}

// ShouldSkipDir returns true if the directory name should be skipped
func ShouldSkipDir(name string, skipList map[string]bool) bool {
	if skipList == nil {
		skipList = DefaultSkipDirs
	}
	return skipList[name]
}
