package tree

// Display the tree much like the output of "tree" command
//
// TODO: Print rendered variables alongside the paths in the tree once rendering is implemented

func (t *Tree) Display() (display string) {
	display = "\n" + t.Root.Name + "\n" + displayNode(t.Root, "")
	return
}

func displayNode(node *Node, prefix string) (display string) {
	count := len(node.children)

	for idx, child := range node.children {
		if idx == count-1 {
			display += prefix + "└── " + child.Name + "\n"
			display += displayNode(child, prefix+"    ")
		} else {
			display += prefix + "├── " + child.Name + "\n"
			display += displayNode(child, prefix+"│   ")
		}
	}

	return
}
