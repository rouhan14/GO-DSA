package main

func IsBalancedParenthesis(expn string) bool {
    stk := new(Stack)
    for _, char := range expn {
        // Push opening parentheses to the stack
        if char == '{' || char == '[' || char == '(' {
            stk.Push(char)
        } else if char == '}' || char == ']' || char == ')' {
            // If stack is empty or doesn't match the top element, it's unbalanced
            if stk.IsEmpty() {
                return false
            }

            top := stk.Top()
            // Check if the popped element matches the current closing parenthesis
            if (char == '}' && top != '{') || (char == ']' && top != '[') || (char == ')' && top != '(') {
                return false
            }

            stk.Pop() // Pop the matching opening parenthesis
        }
    }

    // If stack is empty, all parentheses were balanced
    return stk.IsEmpty()
}
