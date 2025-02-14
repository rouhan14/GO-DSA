package main

import "math"

func (t *Tree) CeilBST(val int) int {

	ceil := math.MinInt32

	curr := t.root

	for curr != nil {
		if curr.value <= val {
			curr = curr.right
		} else {
			ceil = curr.value
			curr = curr.left
		}
	}
	if ceil < 0 {
		return -1
	}

	return ceil
}

func (t *Tree) FloorBST(val int) int {
	curr := t.root
	floor := math.MaxInt32
	for curr != nil {
		if curr.value >= val {
			curr = curr.left
		} else {
			floor = curr.value
			curr = curr.right
		}
	}
	if floor == math.MaxInt32 {
		return -1
	}
	return floor
}
