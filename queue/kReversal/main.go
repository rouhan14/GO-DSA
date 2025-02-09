package main

func reverseKElementInQueue(que *Queue, k int) {
	if k <= 0 {
		return
	}

	stk := &Stack{}
	que2 := &Queue{}

	size_of_que := que.size

	// Push first k elements into stack
	for i := 0; i < k; i++ {
		stk.Push(que.Dequeue().(int))
	}

	// Enqueue back the reversed elements
	for i := 0; i < k; i++ {
		que.Enqueue(stk.Pop())
	}
	// Move the remaining elements to another queue
	for i := 0; i < size_of_que-k; i++ {
		que2.Enqueue(que.Dequeue())
	}

	// Enqueue back the remaining elements to maintain order
	for i := 0; i < que2.size+1; i++ {
		que.Enqueue(que2.Dequeue())
	}
}
