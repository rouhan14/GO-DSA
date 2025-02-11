package main

func (t *Tree) NumNodes() int {
	return numNodes(t.root)
}

func numNodes(curr *Node) int {
	if curr == nil {
		return 0
	}
	return (1 + numNodes(curr.right) + numNodes(curr.left))
}

func (t *Tree) SumAllBT() int {
	return sumAllBT(t.root)
}

func sumAllBT(curr *Node) int {
	if curr == nil {
		return 0
	}
	return (curr.value + sumAllBT(curr.left) + sumAllBT(curr.right))
}
