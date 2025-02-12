package main

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
}

func CreateBinarySearchTree(arr []int) *Tree {
    t := new(Tree)
    size := len(arr)
    t.root = createBinarySearchTreeUtil(arr, 0, size-1)
    return t
}
func createBinarySearchTreeUtil(arr []int, start int, end int) *Node {
    if start > end {
        return nil
    }
    mid := (start + end) / 2
    curr := new(Node)
    curr.value = arr[mid]
    curr.left = createBinarySearchTreeUtil(arr, start, mid-1)
    curr.right = createBinarySearchTreeUtil(arr, mid+1, end)
    return curr
}

func FindMax(curr *Node) *Node {
    var node *Node = curr
    if node == nil {
        return nil
    }
    for node.right != nil {
        node = node.right
    }
    return node
}

func FindMin(curr *Node) *Node {
    var node *Node = curr
    if node == nil {
        return nil
    }
    for node.left != nil {
        node = node.left
    }
    return node
}