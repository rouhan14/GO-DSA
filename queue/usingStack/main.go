package main

type QueueUsingStack struct {
	stk1 Stack
	// add more values here if need be
	stk2 Stack
	
   }
   
   func (que *QueueUsingStack) Add(value int) {
	 //Implement your solution here
	 que.stk1.Push(value)
   
   }
   
   func (que *QueueUsingStack) Remove() int {
	 //Implement your solution here
	 for i:=0; i < que.stk1.Length() + 1; i++ {
	   que.stk2.Push(que.stk1.Pop())
	 }
	 
	 retVal := que.stk2.Top()
	 
	 for i:=0; i < que.stk2.Length() + 1; i++ {
	   que.stk1.Push(que.stk2.Pop())
	 }
   
	 return retVal
   }
   
   func (que *QueueUsingStack) Length() int {
	 //Implement your solution here
   
	 return que.stk1.Length()
   }
   
   func (que *QueueUsingStack) IsEmpty() bool {
	 //Implement your solution here
	 que.stk1.IsEmpty()
   
	 return que.stk1.IsEmpty()
   } 
   