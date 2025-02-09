package main

func reverseQueue(que *Queue) {
	//Implement your solution here
	stk := new(Stack)
	if !que.IsEmpty() {
		var size_of_que = que.size
		for i := 0; i < size_of_que; i++ {
			stk.Push(que.Dequeue().(int))
		}

		for i := 0; i < size_of_que; i++ {
			que.Enqueue(stk.Pop())
		}
	}
}
