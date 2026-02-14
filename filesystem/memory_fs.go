package filesystem

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

// MemoryFileSystem implements FileSystem in-memory for testing
type MemoryFileSystem struct {
	files map[string][]byte
	dirs  map[string]bool
}

// NewMemoryFileSystem creates a new in-memory filesystem
func NewMemoryFileSystem() *MemoryFileSystem {
	return &MemoryFileSystem{
		files: make(map[string][]byte),
		dirs:  make(map[string]bool),
	}
}

// WriteFile writes data to an in-memory file
func (mfs *MemoryFileSystem) WriteFile(path string, data []byte, perm os.FileMode) error {
	// Normalize path
	path = filepath.Clean(path)

	// Create parent directories automatically
	dir := filepath.Dir(path)
	if err := mfs.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Store file content
	mfs.files[path] = data
	return nil
}

// MkdirAll creates a directory hierarchy in memory
func (mfs *MemoryFileSystem) MkdirAll(path string, perm os.FileMode) error {
	// Normalize path
	path = filepath.Clean(path)

	// Create all parent directories
	parts := strings.Split(path, string(filepath.Separator))
	currentPath := ""
	for _, part := range parts {
		if part == "" {
			continue
		}
		if currentPath == "" {
			currentPath = part
		} else {
			currentPath = filepath.Join(currentPath, part)
		}
		mfs.dirs[currentPath] = true
	}
	mfs.dirs[path] = true
	return nil
}

// Walk traverses the directory tree in memory
func (mfs *MemoryFileSystem) Walk(root string, fn func(path string, info fs.FileInfo, err error) error) error {
	root = filepath.Clean(root)

	// Collect all paths that start with root
	var paths []string

	// Add the root directory itself
	if mfs.dirs[root] || root == "." {
		paths = append(paths, root)
	}

	// Add all subdirectories
	for dir := range mfs.dirs {
		if strings.HasPrefix(dir, root+string(filepath.Separator)) || dir == root {
			paths = append(paths, dir)
		}
	}

	// Add all files
	for file := range mfs.files {
		if strings.HasPrefix(file, root+string(filepath.Separator)) || filepath.Dir(file) == root {
			paths = append(paths, file)
		}
	}

	// Sort paths to ensure consistent order
	sort.Strings(paths)

	// Call the walk function for each path
	for _, path := range paths {
		var info fs.FileInfo
		if _, isFile := mfs.files[path]; isFile {
			info = &memoryFileInfo{
				name:  filepath.Base(path),
				size:  int64(len(mfs.files[path])),
				isDir: false,
			}
		} else {
			info = &memoryFileInfo{
				name:  filepath.Base(path),
				isDir: true,
			}
		}

		if err := fn(path, info, nil); err != nil {
			if err == filepath.SkipDir {
				// Skip this directory and its contents
				continue
			}
			return err
		}
	}

	return nil
}

// Stat returns file information for a path
func (mfs *MemoryFileSystem) Stat(path string) (fs.FileInfo, error) {
	path = filepath.Clean(path)

	// Check if it's a file
	if data, exists := mfs.files[path]; exists {
		return &memoryFileInfo{
			name:  filepath.Base(path),
			size:  int64(len(data)),
			isDir: false,
		}, nil
	}

	// Check if it's a directory
	if mfs.dirs[path] {
		return &memoryFileInfo{
			name:  filepath.Base(path),
			isDir: true,
		}, nil
	}

	// Special case: root directory always exists
	if path == "." {
		return &memoryFileInfo{
			name:  ".",
			isDir: true,
		}, nil
	}

	return nil, fmt.Errorf("path does not exist: %s", path)
}

// ReadFile reads a file from memory (helper for testing)
func (mfs *MemoryFileSystem) ReadFile(path string) ([]byte, error) {
	path = filepath.Clean(path)
	data, exists := mfs.files[path]
	if !exists {
		return nil, fmt.Errorf("file not found: %s", path)
	}
	return data, nil
}

// memoryFileInfo implements fs.FileInfo for in-memory files
type memoryFileInfo struct {
	name  string
	size  int64
	isDir bool
}

func (mfi *memoryFileInfo) Name() string       { return mfi.name }
func (mfi *memoryFileInfo) Size() int64        { return mfi.size }
func (mfi *memoryFileInfo) Mode() fs.FileMode  { return 0644 }
func (mfi *memoryFileInfo) ModTime() time.Time { return time.Now() }
func (mfi *memoryFileInfo) IsDir() bool        { return mfi.isDir }
func (mfi *memoryFileInfo) Sys() interface{}   { return nil }
