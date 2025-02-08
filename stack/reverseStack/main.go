package main

func reverseStack(stk *Stack) {
	//Implement your solution here
	var temp int
	if stk.Length() != 0 {
	  temp = stk.Pop()
	  reverseStack(stk)
	  bottomInsert(stk, temp)
	}
	
}