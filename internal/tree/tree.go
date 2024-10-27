package tree

import (
	"os"
	"path/filepath"
)

type Tree struct {
	Paths []string
}

// Generate a tree from the root of the project
// Return pointer in case we need future mutablility
func NewTree() (t *Tree, err error) {
	var paths []string

	t = &Tree{}

	path, err := getRootDirectory()
	if err != nil {
		return
	}

	err = recursePath(path, &paths)
	if err != nil {
		return
	}

	t.Paths = paths
	return
}

// Recursively walk the directory tree, appending paths to the slice
// Using a slice pointer should help with very large trees
func recursePath(path string, paths *[]string) (err error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	*paths = append(*paths, path)

	for _, e := range entries {
		if e.IsDir() {
			recursePath(filepath.Join(path, e.Name()), paths)
		}
	}

	return
}
