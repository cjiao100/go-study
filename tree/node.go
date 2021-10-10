package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// Print    (node Node) 是一个接收者，类似java的类方法
func (node Node) Print() {
	fmt.Println(node.Value)
}

// SetValue 因为go都是值传递，如果需要改传入的对象里的值，所以需要接收指针
func (node *Node) SetValue(value int) {
	node.Value = value
}

// CreateTreeNode 工厂函数，可以替代构造函数的作用，其实本质就是创建了一个新的对象，然后将对象的地址返回
func CreateTreeNode(value int) *Node {
	return &Node{Value: value}
}
