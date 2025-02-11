package main

import "fmt"

func (t *Tree) PrintDepthFirst() {

	//Implement your solution here
	stk := new(Stack)
	if t.root != nil {
		stk.Push(t.root)
	}

	for stk.Length() != 0 {
		temp := stk.Pop()
		fmt.Print(temp.value, " ")
		if temp.right != nil {
			stk.Push(temp.right)
		}
		if temp.left != nil {
			stk.Push(temp.left)
		}
	}
}
