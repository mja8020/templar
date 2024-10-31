package tree

import "iter"

func traverse(node *Node) []*Node {
	results := NodeSet{node}

	for _, c := range node.children {
		results = append(results, traverse(c)...)
	}

	return results
}

// Iterator - Iterates over the tree returning the Parent and the current Node
func (t *Tree) Iterator() iter.Seq2[int, *Node] {
	// FIXME: This isn't an actual yield/iterator since we enumerate all nodes first
	nodeList := traverse(t.Root)

	return func(yield func(int, *Node) bool) {
		for i, node := range nodeList {
			if !yield(i, node) {
				return
			}
		}
	}
}
