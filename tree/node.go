package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createNode(value int) *treeNode {
	// 使用工厂函数来实现自定义的构造函数
	// 返回了局部变量的地址
	return &treeNode{value: value}
}

func main() {
	var root treeNode

	root = treeNode{value: 1}
	root.right = &treeNode{}
	root.left = &treeNode{2, nil, nil}
	root.right.right = new(treeNode)
	fmt.Println(root)

	fmt.Println()
	root.right.setValue(3)
	root.right.print()
	root.traverse()
}
