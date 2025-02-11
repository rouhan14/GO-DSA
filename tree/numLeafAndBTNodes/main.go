package main

func (t *Tree) NumLeafNodes() int {
	return numLeafNodes(t.root)
}

func numLeafNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left == nil && curr.right == nil {
		return 1
	}
	return numLeafNodes(curr.left) + numLeafNodes(curr.right)
}

func (t *Tree) NumFullNodesBT() int {
	return numFullNodesBT(t.root)
}

func numFullNodesBT(curr *Node) int {
	if curr == nil {
		return 0
	}
	if curr.left != nil && curr.right != nil {
		return 1 + numFullNodesBT(curr.left) + numFullNodesBT(curr.right)
	}
	return numFullNodesBT(curr.left) + numFullNodesBT(curr.right)
}
