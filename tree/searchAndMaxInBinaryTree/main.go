package main

import "math"

func (t *Tree) SearchBT(value int) bool {
	return searchBT(t.root, value)
}

func searchBT(root *Node, value int) bool {
	var left, right bool
	if root == nil {
		return false
	}
	if root.value == value {
		return true
	}
	left = searchBT(root.left, value)
	if left {
		return true
	}
	right = searchBT(root.right, value)
	if right {
		return true
	}
	return false
}

func (t *Tree) FindMaxBT() int {
	return findMaxBT(t.root)
}

func findMaxBT(curr *Node) int {
	if curr == nil {
		return math.MinInt32
	}
	max := curr.value
	left := findMaxBT(curr.left)
	right := findMaxBT(curr.right)
	if left > max {
		max = left
	}
	if right > max {
		max = right
	}
	return max
}
