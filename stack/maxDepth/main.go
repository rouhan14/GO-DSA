package main

func maxDepthParenthesis(expn string, size int) int {
	//Implement your solution here
	stk := new(Stack)
	maxDepth := 0
	depth := 0
	var ch byte
	
	for i := 0; i<size;i++ {
	  ch = expn[i]
	  if(ch == '(') {
		stk.Push(ch)
		depth += 1
	  } else if ch == ')' {
		stk.Pop()
		depth -= 1
	  }
	  if depth > maxDepth {
		maxDepth = depth
	  }
	}
	return maxDepth //Return 0 if depth of parenthesis is zero
  }