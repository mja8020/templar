package tree

func (t *Tree) Display() (display string) {
	display = "\n" + t.root.Name + "\n" + displayNode(t.root, "")
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
