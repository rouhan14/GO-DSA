package main

type MinStack struct {
	maxSize int
	// Initialize your data structures here
	mainStack    Stack
	minimumStack Stack
}

// removes and returns value from stack
func (m *MinStack) Pop() int {
	// write your code here
	_ = m.minimumStack.Pop()
	top := m.mainStack.Top()
	_ = m.mainStack.Pop()
	return top
}

// Pushes value into the stack
func (m *MinStack) Push(value int) {
	// write your code here
	m.mainStack.Push(value)
	if !m.minimumStack.IsEmpty() && m.minimumStack.Top() < value {
		m.minimumStack.Push(m.minimumStack.Top())
	} else {
		m.minimumStack.Push(value)
	}
}

// returns minimum value in O(1)
func (m *MinStack) Min() int {
	// write your code here
	return m.minimumStack.Top()
}
