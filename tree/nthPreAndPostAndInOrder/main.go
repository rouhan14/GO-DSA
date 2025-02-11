package main

import "fmt"

func (t *Tree) NthPreOrder(index int) {
	var counter int
	nthPreOrder(t.root, index, &counter)
}

func nthPreOrder(node *Node, index int, counter *int) {
	//Implement your solution here

	if node != nil {
		(*counter)++
		if *counter == index {
			fmt.Print(node.value)
		}
		nthPreOrder(node.left, index, counter)
		nthPreOrder(node.right, index, counter)
	}
}

func (t *Tree) NthPostOrder(index int) {
	var counter int
	nthPostOrder(t.root, index, &counter)
}

func nthPostOrder(node *Node, index int, counter *int) {
	//Implement your solution here

	if node != nil {
		nthPostOrder(node.left, index, counter)
		nthPostOrder(node.right, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Print(node.value)
		}
	}
}

func (t *Tree) NthInOrder(index int) {
	var counter int
	nthInOrder(t.root, index, &counter)
}

func nthInOrder(node *Node, index int, counter *int) {
	//Implement your solution here
	if node != nil {
		nthInOrder(node.left, index, counter)
		(*counter)++
		if *counter == index {
			fmt.Print(node.value)
		}
		nthInOrder(node.right, index, counter)
	}
}
