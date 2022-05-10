package V2

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack"
	"fmt"
	"strconv"
)

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

func (n *node) Push(e int) *node {
	if n == nil {
		return newNode().SetE(e)
	}
	if e < n.E {
		// 对于当前时刻来说是否存在n.Left是未知的, 我们假设执行一下
		//		如果节点是不存在的:
		//			那么n.Left.Push()就会返回一个新节点指针, 并且这个节点的值是我们传入的值
		//		如果节点是存在的:
		//			那么就又会进行判断, 如果判断需要将元素添加至左孩子那么其又会再次调用n.Left = n.Left.Push(e), 直至调用到某一个Push函数中时刻, 发现当前节点为空, 那么返回一个新的节点指针, 并且这个节点的值是我们传入的值
		//			执行完上述步骤后(即返回新节点那个时刻的函数之后), 再次回到当前位置, 继续执行直至最后一行`return n`返回当前节点的指针, 直到某个个节点为顶点, 顶点也会返回自己的指针, 此时在BST中我们将接收了这个指针, 并且将其赋值给root节点
		//			* 通过观察可以发现
		//				* 如果当前节点不为空, 那么返回的指针总是不变的
		//				* 如果当前节点为空, 那么就会返回全新指针
		n.Left = n.Left.Push(e)
	} else {
		n.Right = n.Right.Push(e)
	}
	return n
}

func (n *node) Has(e int) bool {
	if n == nil {
		return false
	}
	if n.E == e {
		return true
	}
	// 当节点的值大于欲查询的值时, 从左侧寻找
	// 因为节点一定会大于左孩子的值
	//
	// 比节点大的值一定在节点的右子树中, 因为节点得值一定小于右子树
	// 比节点小的值一定在节点的左子树中, 因为节点的值一定大于左子树
	if n.E > e {
		return n.Left.Has(e)
	} else {
		return n.Right.Has(e)
	}
}

func (n *node) SetE(e int) *node {
	n.E = e
	return n
}

// PreOrder 前序遍历
func (n *node) PreOrder() {
	if n == nil {
		return
	}
	fmt.Println(n.E)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *node) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Println(n.E)
	n.Right.InOrder()
}

func (n *node) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Println(n.E)
}

//func (n *node) PreOrderNR() {
//	if n == nil {
//		return
//	}
//	stack := Stack.Stack{}
//	stack.Push(n)
//	for stack.Size() > 0 {
//		curNode := stack.Pop().(*node)
//		fmt.Println(curNode.E)
//		if curNode.Right != nil {
//			stack.Push(curNode.Right)
//		}
//		if curNode.Left != nil {
//			stack.Push(curNode.Left)
//		}
//	}
//}

// PreOrderNR 非递归前序遍历
func (n *node) PreOrderNR() {
	// 创建一个栈, 使用这个栈模拟递归用到的系统栈
	stack := Stack.NewStack()
	stack.Push(n) // 甭管当前节点是否为空都存进去
	for !stack.IsEmpty() {
		curNode := stack.Pop().(*node)
		if curNode == nil { // 不管是不是空都存了进来, 所以判断下, 如果是空的就跳过
			continue
		}
		fmt.Println(curNode.E)
		stack.Push(curNode.Right) // 甭管左孩子是否为空都存进去
		stack.Push(curNode.Left)  // 甭管右孩子是否为空都存进去
	}
}

// generateBSTString 通过前序遍历的方式生成字符串
func (n *node) generateBSTString(depth int) (result string) {
	if n == nil {
		return result + n.generateDepthString(depth) + "null\n"
	}

	//result = result + n.generateDepthString(depth) + strconv.Itoa(n.E) + "\n"
	//result = result + n.Left.generateBSTString(depth+1)
	//result = result + n.Right.generateBSTString(depth+1)
	return (n.generateDepthString(depth) + strconv.Itoa(n.E) + "\n") + n.Left.generateBSTString(depth+1) + n.Right.generateBSTString(depth+1)
}

func (n *node) generateDepthString(depth int) (result string) {
	for i := 0; i < depth; i++ {
		result += "--"
	}
	return result
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
	b.root = b.root.Push(e)
	b.size++
}

// PreOrder 前序遍历
func (b *BST) PreOrder() {
	b.root.PreOrder()
}

func (b *BST) InOrder() {
	b.root.InOrder()
}

func (b *BST) PostOrder() {
	b.root.PostOrder()
}

func (b *BST) PreOrderNR() {
	b.root.PreOrderNR()
}

func (b *BST) GetSize() int {
	return b.size
}

func (b *BST) Has(e int) bool {
	return b.root.Has(e)
}

// String 按照前序遍历的方式生成字符串, 并实现Stringer接口
func (b *BST) String() string {
	return b.root.generateBSTString(0)
}
