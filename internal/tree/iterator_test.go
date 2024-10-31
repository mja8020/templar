package tree

import (
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

	tree, err := NewTree("../../fixtures/1_basic")
	require.NoError(t, err)

	expecting := [][]string{
		{"", "/", "/workspaces/Templar/fixtures/1_basic", ""},
		{"a", "/a", "/workspaces/Templar/fixtures/1_basic/a", "/"},
		{"b", "/a/b", "/workspaces/Templar/fixtures/1_basic/a/b", "/a"},
		{"c", "/a/c", "/workspaces/Templar/fixtures/1_basic/a/c", "/a"},
		{"d", "/a/c/d", "/workspaces/Templar/fixtures/1_basic/a/c/d", "/a/c"},
		{"e", "/e", "/workspaces/Templar/fixtures/1_basic/e", "/"},
		{"f", "/e/f", "/workspaces/Templar/fixtures/1_basic/e/f", "/e"},
	}

	for i, n := range tree.Iterator() {
		assert.Empty(t, nil, "")

		assert.Equal(t, expecting[i][0], n.Name)
		assert.Equal(t, expecting[i][1], n.Label)
		assert.Equal(t, expecting[i][2], n.Path)
		assert.Equal(t, expecting[i][3], n.Parent)
	}
}
