package filesystem

import (
	"io/fs"
	"os"
)

// FileSystem abstracts filesystem operations for testability
type FileSystem interface {
	// WriteFile writes data to a file at the specified path
	WriteFile(path string, data []byte, perm os.FileMode) error

	// MkdirAll creates a directory hierarchy
	MkdirAll(path string, perm os.FileMode) error

	// Walk traverses a directory tree
	Walk(root string, fn func(path string, info fs.FileInfo, err error) error) error

	// Stat returns file information
	Stat(path string) (fs.FileInfo, error)
}
