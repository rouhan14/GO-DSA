package main

func bottomInsert(stk *Stack, element int) {
	//Implement your solution here
	var temp int
	if stk.Length() == 0 {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		bottomInsert(stk, element)
		stk.Push(temp)
	}

}
