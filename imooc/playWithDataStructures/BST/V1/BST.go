package V1

import "fmt"

// node 二分搜索树节点
type node struct {
	E     int
	Left  *node
	Right *node
}

func newNode() *node {
	return &node{
		E:     0,
		Left:  nil,
		Right: nil,
	}
}

func (n *node) Push(e int) {
	if e < n.E {
		if n.Left == nil {
			n.Left = newNode()
			n.Left.E = e
			return
		} else {
			n.Left.Push(e)
		}
	} else {
		if n.Right == nil {
			n.Right = newNode()
			n.Right.E = e
			return
		} else {
			n.Right.Push(e)
		}
	}
}

func (n *node) Print() {
	if n == nil {
		return
	}
	fmt.Println(n.E)
	n.Left.Print()
	n.Right.Print()
}

// BST 二分搜索树
type BST struct {
	root *node
	size int
}

func NewBST() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}

func (b *BST) Push(e int) {
	if b.root == nil {
		b.root = newNode()
		b.root.E = e
		b.size++
		return
	}
	b.root.Push(e)
	b.size++
}

func (b *BST) Print() {
	b.root.Print()
}

func (b *BST) GetSize() int {
	return b.size
}
