package main

func sortStack(stk *Stack) {
	//Implement your solution here
	var temp int
	if stk.Length() > 0 {
		temp = stk.Pop()
		sortStack(stk)
		sortedInsert(stk, temp)
	}

}
