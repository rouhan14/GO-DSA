func (t *Tree) PrintPostOrder() {
    printPostOrder(t.root)
}

func printPostOrder(n *Node) {

    //Implement your solution here
    if n==nil {
      return
    }


    //Implement your solution here
    printPostOrder(n.left)
    printPostOrder(n.right)
    fmt.Print(n.value, " ")
}