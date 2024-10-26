package tree

import "testing"

func TestBasicTree(t *testing.T) {
	tree, err := NewTree("../../fixtures/basic")
	if err != nil {
		t.Errorf("Error creating tree: %s", err)
	}

	if len(tree.Paths) != 4 {
		t.Errorf("Expected 4 paths, got %d", len(tree.Paths))
	}
}
