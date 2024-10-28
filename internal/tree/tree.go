package tree

import (
	"os"
	"path/filepath"
)

type NodeSet []*Node

type Node struct {
	Name string

	children NodeSet
}

type Tree struct {
	Size int

	root *Node
}

func NewTree(path string) (t *Tree, err error) {
	t = &Tree{}

	path, err = getRootDirectory(path)
	if err != nil {
		return
	}

	root := &Node{
		Name: path,
	}

	count, err := recursePath(path, root)
	if err != nil {
		return
	}

	t.root = root
	t.Size = count + 1 // include root in size

	return
}

func recursePath(path string, parent *Node) (count int, err error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return
	}

	for _, e := range entries {
		if e.IsDir() {

			node := &Node{
				Name: e.Name(),
			}

			parent.children = append(parent.children, node)

			var childCount int

			childCount, err = recursePath(filepath.Join(path, e.Name()), node)
			if err != nil {
				return
			}

			count += childCount + 1
		}
	}

	return
}
