package main

func sortedInsert(stk *Stack, element int) {
	//Implement your solution here
	var temp int
	if stk.Length() == 0 || element > stk.Top() {
		stk.Push(element)
	} else {
		temp = stk.Pop()
		sortedInsert(stk, element)
		stk.Push(temp)
	}
}
