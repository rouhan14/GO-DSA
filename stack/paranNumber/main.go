package main

import "fmt"

func printParenthesisNumber(expn string, size int) {
	//Uncomment the ch variable and use it iterate through the string
	var ch byte
	stk := new(Stack)
	output := "" //use output variable to save and print the output string
	var count int = 1

	//Implement your solution here
	for i := 0; i < size; i++ {
		ch = expn[i]

		if ch == '(' {
			output += fmt.Sprintf("%v", count)
			stk.Push(count)
			count += 1
		} else if ch == ')' {
			output += fmt.Sprintf("%v", stk.Top())
			stk.Pop()
		}
	}

	fmt.Println(output)
}
