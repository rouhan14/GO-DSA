package main

import "fmt"

func (t *Tree) FindMinNode() *Node {

	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}
	for node.left != nil {
		node = node.left
	}

	return node
}

func (t *Tree) FindMaxNode() *Node {
	var node *Node = t.root
	if node == nil {
		fmt.Println("EmptyTreeException")
		return nil
	}
	for node.right != nil {
		node = node.right
	}

	return node
}
