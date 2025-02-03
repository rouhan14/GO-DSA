package main

type Node struct {
 value int
 next *Node
}

type StackLinkedList struct {
 head *Node
 size int
}
