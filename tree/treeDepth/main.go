package main

func (t *Tree) TreeDepth() int {
    return treeDepth(t.root)
}

func treeDepth(root *Node) int {
  if root == nil {
    return -1
  }
  leftDepth:=treeDepth(root.left)
  rightDepth:=treeDepth(root.right)
  if leftDepth > rightDepth {
    return leftDepth+1
  }
  return rightDepth+1
}