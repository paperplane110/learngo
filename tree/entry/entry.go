package main

import (
	"fmt"
	"github.com/paperplane110/learngo/tree"
)

func main() {
	var root tree.Node

	root = tree.Node{Value: 1}
	root.Right = &tree.Node{}
	root.Left = &tree.Node{2, nil, nil}
	root.Right.Right = new(tree.Node)
	fmt.Println(root)

	fmt.Println()
	root.Right.SetValue(3)
	root.Right.Print()
	root.Traverse()
}
