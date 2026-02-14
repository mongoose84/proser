package filesystem

import (
	"io/fs"
	"os"
	"path/filepath"
)

// OsFileSystem implements FileSystem using real OS operations
type OsFileSystem struct{}

// NewOsFileSystem creates a new OsFileSystem
func NewOsFileSystem() *OsFileSystem {
	return &OsFileSystem{}
}

// WriteFile writes data to a file
func (fs *OsFileSystem) WriteFile(path string, data []byte, perm os.FileMode) error {
	return os.WriteFile(path, data, perm)
}

// MkdirAll creates a directory hierarchy
func (fs *OsFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// Walk traverses a directory tree
func (fs *OsFileSystem) Walk(root string, fn func(path string, info fs.FileInfo, err error) error) error {
	return filepath.Walk(root, fn)
}

// Stat returns file information
func (fs *OsFileSystem) Stat(path string) (fs.FileInfo, error) {
	return os.Stat(path)
}
