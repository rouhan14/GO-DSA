package main

import "fmt"

func (t *Tree) PrintBreadthFirst() {
	que := new(Queue)
	var temp *Node
	if t.root != nil {
		que.Enqueue(t.root)
	}

	for que.Length() != 0 {
		temp2 := que.Dequeue()
		temp = temp2.(*Node)
		fmt.Print(temp.value, " ")
		if temp.left != nil {
			que.Enqueue(temp.left)
		}
		if temp.right != nil {
			que.Enqueue(temp.right)
		}
	}
}
