package tree

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIterator(t *testing.T) {
	/*
	   name -> label -> path
	     -> / -> /workspaces/Templar/fixtures/1_basic
	   a -> /a -> /workspaces/Templar/fixtures/1_basic/a
	   b -> /a/b -> /workspaces/Templar/fixtures/1_basic/a/b
	   c -> /a/c -> /workspaces/Templar/fixtures/1_basic/a/c
	   d -> /a/c/d -> /workspaces/Templar/fixtures/1_basic/a/c/d
	   e -> /e -> /workspaces/Templar/fixtures/1_basic/e
	   f -> /e/f -> /workspaces/Templar/fixtures/1_basic/e/f
	*/

	fixture := "../../fixtures/1_basic"
	path, err := filepath.Abs(fixture)
	require.NoError(t, err)

	tree, err := NewTree(fixture)
	require.NoError(t, err)

	// Full path depends where test is running from
	expecting := [][]string{
		{"", "/", path, ""},
		{"a", "/a", filepath.Join(path, "a"), "/"},
		{"b", "/a/b", filepath.Join(path, "a", "b"), "/a"},
		{"c", "/a/c", filepath.Join(path, "a", "c"), "/a"},
		{"d", "/a/c/d", filepath.Join(path, "a", "c", "d"), "/a/c"},
		{"e", "/e", filepath.Join(path, "e"), "/"},
		{"f", "/e/f", filepath.Join(path, "e", "f"), "/e"},
	}

	for i, n := range tree.Iterator() {
		assert.Empty(t, nil, "")

		assert.Equal(t, expecting[i][0], n.Name)
		assert.Equal(t, expecting[i][1], n.Label)
		assert.Equal(t, expecting[i][2], n.Path)
		assert.Equal(t, expecting[i][3], n.Parent)
	}
}
