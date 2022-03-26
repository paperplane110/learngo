package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func createNode(value int) *Node {
	// 使用工厂函数来实现自定义的构造函数
	// 返回了局部变量的地址
	return &Node{Value: value}
}
