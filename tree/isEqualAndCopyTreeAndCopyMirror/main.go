package main

func (t *Tree) IsEqual(t2 *Tree) bool {
	return isEqual(t.root, t2.root)
}

func isEqual(node1 *Node, node2 *Node) bool {
	//Implement your solution here
	if node1 == nil && node2 == nil {
		return true
	} else if node1 == nil || node2 == nil {
		return false
	} else {
		return ((node1.value == node2.value) &&
			isEqual(node1.left, node2.left) &&
			isEqual(node1.right, node2.right))
	}
}

func (t *Tree) CopyTree() *Tree {
	Tree2 := new(Tree)
	Tree2.root = copyTree(t.root)
	return Tree2
}

func copyTree(curr *Node) *Node {
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.left = copyTree(curr.left)
		temp.right = copyTree(curr.right)
		return temp
	}
	return nil
}

func (t *Tree) CopyMirrorTree() *Tree {
	tree := new(Tree)
	tree.root = copyMirrorTree(t.root)
	return tree
}

func copyMirrorTree(curr *Node) *Node {
	//Implement your solution here
	var temp *Node
	if curr != nil {
		temp = new(Node)
		temp.value = curr.value
		temp.left = copyMirrorTree(curr.right)
		temp.right = copyMirrorTree(curr.left)
		return temp
	}

	return nil
}
