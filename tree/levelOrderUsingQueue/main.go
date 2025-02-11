package main

import "fmt"

func (t *Tree) PrintLevelOrderLineByLine() {
	//Implement your solution here
	que1 := new(Queue)
	que2 := new(Queue)
	var temp *Node

	if t.root != nil {
		que1.Enqueue(t.root)
	}

	for que2.Length() != 0 || que1.Length() != 0 {
		for que1.Length() != 0 {
			temp2 := que1.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que2.Enqueue(temp.left)
			}
			if temp.right != nil {
				que2.Enqueue(temp.right)
			}
		}
		fmt.Print("; ")
		for que2.Length() != 0 {

			temp2 := que2.Dequeue()
			temp = temp2.(*Node)
			fmt.Print(temp.value, " ")
			if temp.left != nil {
				que1.Enqueue(temp.left)
			}
			if temp.right != nil {
				que1.Enqueue(temp.right)
			}
		}
		fmt.Print("; ")
	}
}
