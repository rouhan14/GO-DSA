package main

func (t *Tree) Find(value int) bool {
	var temp *Node = t.root
	for temp != nil {
		if temp.value == value {
			return true
		} else if temp.value > value {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}
	return false
}
