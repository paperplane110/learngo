package main

import (
	"fmt"
	"github.com/paperplane110/learngo/tree"
)

//type MyNode struct {
//	node *tree.Node
//}
//
//func (node *MyNode) postOrder() {
//	if node.node == nil {
//		return
//	}
//	left := MyNode{node.node.Left}
//	right := MyNode{node.node.Right}
//	left.postOrder()
//	right.postOrder()
//	node.node.Print()
//}

type MyNode struct {
	*tree.Node // this is embedded
}

func (node *MyNode) postOrder() {
	if node.Node == nil {
		return
	}
	left := MyNode{node.Left}
	right := MyNode{node.Right}
	left.postOrder()
	right.postOrder()
	node.Print()
}

func main() {
	root := MyNode{&tree.Node{Value: 3}}
	root.Right = &tree.Node{}
	root.Left = &tree.Node{2, nil, nil}
	root.Right.Right = new(tree.Node)
	fmt.Println(root)

	fmt.Println()
	root.Right.SetValue(3)
	root.Right.Print()
	root.Traverse()

	fmt.Println()
	myRoot := root
	myRoot.postOrder()
}
