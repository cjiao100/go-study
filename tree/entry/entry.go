package main

import (
	"fmt"

	"github.com/cjiao100/learngo/tree"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

type myTreeNode2 struct {
	*tree.Node
}

func (myNode *myTreeNode2) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := myTreeNode{myNode.Left}
	left.postOrder()
	right := myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

// Traverse 相当于函数重载
func (myNode *myTreeNode2) Traverse() {
	fmt.Println("this methods is shadowed")
}

func main() {
	var root tree.Node

	fmt.Println(root)

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}

	// go语言的指针可以直接.引用
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateTreeNode(2)

	fmt.Println(root)

	root.Print()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	// 在go中通过地址或者直接通过值调用，都是一样的
	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(200)
	pRoot.Print()

	root.Traverse()

	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()

	// 内嵌写法，类似继承
	myRoot2 := myTreeNode2{&root}
	myRoot2.Traverse()
	myRoot2.Node.Traverse()

	//var baseRoot *tree.Node
	//baseRoot = &myRoot2

	nodes := []tree.Node{
		{Value: 3},
		{},
		// 这种时候可以忽略属性名，只写value
		{5, nil, &root},
	}

	fmt.Println(nodes)
}
