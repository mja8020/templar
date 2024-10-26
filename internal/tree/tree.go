package tree

import (
	"os"
	"path/filepath"
)

type Tree struct {
	Paths []string
}

func NewTree(path string) (t *Tree, err error) {
	var paths []string

	err = recursePath(path, &paths)
	if err != nil {
		return &Tree{}, err
	}

	t = &Tree{paths}
	return
}

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
