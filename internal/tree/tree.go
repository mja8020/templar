package tree

import (
	"os"
	"path/filepath"
)

type NodeSet []*Node

/*
Traverse: "../../fixtures/1_basic" with a full path of "/workspaces/Templar/fixtures/1_basic"

Root
{
	Name: ""
	Label: "/"
	Path: "/workspaces/Templar/fixtures/1_basic"
	Parent: ""
}

Folder
{
	Name: "a"
	Label: "/a"
	Path: "/workspaces/Templar/fixtures/1_basic/a"
	Parent: "/"
}
*/

type Node struct {
	Name   string // Basename, empty when root
	Path   string // Full filesystemn path
	Label  string // Path relative to root minus leading . (on windows use / for separator here)
	Parent string // Label of the parent (string not reference to avoid circular dependencies)

	children NodeSet
}

type Tree struct {
	Size int

	Root *Node
}

func NewTree(path string) (t *Tree, err error) {
	t = &Tree{}

	path, err = getRootDirectory(path)
	if err != nil {
		return
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return
	}

	root := &Node{
		Name:   "",
		Label:  "/",
		Path:   path,
		Parent: "",
	}

	count, err := recursePath(root)
	if err != nil {
		return
	}

	t.Root = root
	t.Size = count + 1 // include root in size

	return
}

func recursePath(parent *Node) (count int, err error) {
	entries, err := os.ReadDir(parent.Path)
	if err != nil {
		return
	}

	for _, e := range entries {
		if e.IsDir() {
			// Otherwise to have root as / every node would start with //
			label := parent.Label + "/" + e.Name()
			if parent.Label == "/" {
				label = parent.Label + e.Name()
			}

			node := &Node{
				Name:   e.Name(),
				Path:   filepath.Join(parent.Path, e.Name()),
				Label:  label,
				Parent: parent.Label,
			}

			parent.children = append(parent.children, node)

			var childCount int

			childCount, err = recursePath(node)
			if err != nil {
				return
			}

			count += childCount + 1
		}
	}

	return
}
