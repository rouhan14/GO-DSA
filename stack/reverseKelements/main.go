package main

func reverseKElementInStack(stk *Stack, k int) {
    //Implement your solution here
    var temp int
    que := &Queue{}
    if (stk.Length() != 0) {
      for i:=0; i < k; i++ {
        temp = stk.Pop()
        que.Enqueue(temp)
      }
      
      for !que.IsEmpty() {
        stk.Push(que.Dequeue().(int))
      }
    }
}
