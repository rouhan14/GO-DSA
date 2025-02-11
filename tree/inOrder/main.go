package main

import "fmt"

func (t *Tree) PrintInOrder() {
	printInOrder(t.root)
}

func printInOrder(n *Node) {

	//Implement your solution here
	if n == nil {
		return
	}
	printInOrder(n.left)
	fmt.Print(n.value, " ")
	printInOrder(n.right)
}
