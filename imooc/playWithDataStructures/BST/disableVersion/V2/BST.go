package V2

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Queue/SliceQueue"
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

// push 以当前节点为根, 向左子树或者又子树添加节点, 最后返回当前节点, 如果当前节点不存在则返回一个新的节点
func (n *node) push(e int) *node {
	if n == nil {
		return newNode().setE(e)
	}
	if e < n.E {
		// 对于当前时刻来说是否存在n.Left是未知的, 我们假设执行一下
		//		如果节点是不存在的:
		//			那么n.Left.push()就会返回一个新节点指针, 并且这个节点的值是我们传入的值
		//		如果节点是存在的:
		//			那么就又会进行判断, 如果判断需要将元素添加至左孩子那么其又会再次调用n.Left = n.Left.push(e), 直至调用到某一个Push函数中时刻, 发现当前节点为空, 那么返回一个新的节点指针, 并且这个节点的值是我们传入的值
		//			执行完上述步骤后(即返回新节点那个时刻的函数之后), 再次回到当前位置, 继续执行直至最后一行`return n`返回当前节点的指针, 直到某个个节点为顶点, 顶点也会返回自己的指针, 此时在BST中我们将接收了这个指针, 并且将其赋值给root节点
		//			* 通过观察可以发现
		//				* 如果当前节点不为空, 那么返回的指针总是不变的
		//				* 如果当前节点为空, 那么就会返回全新指针
		n.Left = n.Left.push(e)
	} else if e > n.E {
		n.Right = n.Right.push(e)
	}
	return n
}

// has 以当前节点为根, 查找是否存在e
func (n *node) has(e int) bool {
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
		return n.Left.has(e)
	} else {
		return n.Right.has(e)
	}
}

// setE 设定当前节点的值, 并且返回该节点
func (n *node) setE(e int) *node {
	n.E = e
	return n
}

// preOrder 以当前节点为根, 进行前序遍历
func (n *node) preOrder() {
	if n == nil {
		return
	}
	fmt.Println(n.E)
	n.Left.preOrder()
	n.Right.preOrder()
}

// inOrder 以当前节点为根, 进行中序遍历
func (n *node) inOrder() {
	if n == nil {
		return
	}
	n.Left.inOrder()
	fmt.Println(n.E)
	n.Right.inOrder()
}

// postOrder 以当前节点为根, 进行后续遍历
func (n *node) postOrder() {
	if n == nil {
		return
	}
	n.Left.postOrder()
	n.Right.postOrder()
	fmt.Println(n.E)
}

//func (n *node) preOrderNR() {
//	if n == nil {
//		return
//	}
//	stack := Stack.Stack{}
//	stack.push(n)
//	for stack.Size() > 0 {
//		curNode := stack.Pop().(*node)
//		fmt.Println(curNode.E)
//		if curNode.Right != nil {
//			stack.push(curNode.Right)
//		}
//		if curNode.Left != nil {
//			stack.push(curNode.Left)
//		}
//	}
//}

// preOrderNR 以当前节点为根, 进行非递归前序遍历
func (n *node) preOrderNR() {
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

// levelOrder 以当前节点为根, 进行层级遍历
func (n *node) levelOrder() {
	queue := SliceQueue.NewSliceQueue()
	queue.Push(n)
	for !queue.IsEmpty() {
		head := queue.Pop().(*node)
		if head == nil {
			continue
		}
		fmt.Println(head.E)
		queue.Push(head.Left)
		queue.Push(head.Right)
	}
}

// generateBSTString 以当前节点为跟, 通过前序遍历的方式生成字符串
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

// minimum 以当前节点作为根节点, 寻找最小的值并返回. 调用时务必确保被调用节点不为空
func (n *node) minimum() *node {
	if n.Left == nil {
		return n
	}
	return n.Left.minimum()
}

// minimum 以当前节点作为根节点, 寻找最小的值并返回. 调用时务必确保被调用节点不为空
func (n *node) maximum() *node {
	if n.Right == nil {
		return n
	}
	return n.Right.maximum()
}

func (n *node) removeMin() *node {
	if n.Left == nil {
		return n.Right
	}
	n.Left = n.Left.removeMin()
	return n
}

func (n *node) removeMax() *node {
	// 如果右侧节点为空, 那么自己就是最大的节点
	// 		所以返回当前节点的左侧作为上一次调用者的右侧节点(即对于上一个调用者来说将当前节点的右孩子替换成了原本右孩子的左孩子. 这也是没问题的, 因为对于调用的右孩子来说, 当前节点是一定小于又子树中的任意孩子的, 其中就当前包括原右侧节点的左孩子)
	if n.Right == nil {
		return n.Left
	}
	n.Right = n.Right.removeMax()
	return n
}

// 以当前节点为根, 返回删除目标后的当前节点
func (n *node) remove(e int) *node {
	if n.E == e { //TODO: logic err
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		} else {
			successor := n.Right.minimum()
			successor.Left = n.Left
			successor.Right = n.Right.removeMin()
			return successor
		}
	}
	if n.E > e {
		n.Left = n.Left.remove(e)
	} else {
		n.Right = n.Right.remove(e)
	}
	return n
}

// BST 二分搜索树
type BST struct {
	root *node
	size int
}

// NewBST 生成新的二叉搜索树
func NewBST() *BST {
	return &BST{
		root: nil,
		size: 0,
	}
}

// Push 添加元素
func (b *BST) Push(e int) {
	b.root = b.root.push(e)
	b.size++
}

// PreOrder 对整个二叉搜索树进行前序遍历
func (b *BST) PreOrder() {
	b.root.preOrder()
}

// InOrder 对整个二叉搜索树进行中序遍历
func (b *BST) InOrder() {
	b.root.inOrder()
}

// PostOrder 对整个二叉搜索树进行后续遍历
func (b *BST) PostOrder() {
	b.root.postOrder()
}

// PreOrderNR 对整个二叉搜索树进行非递归的前序遍历
func (b *BST) PreOrderNR() {
	b.root.preOrderNR()
}

// LevelOrder 对整个BST进行遍历
func (b *BST) LevelOrder() {
	b.root.levelOrder()
}

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.Size() == 0
}

func (b *BST) Has(e int) bool {
	return b.root.has(e)
}

// String 按照前序遍历的方式生成字符串, 并实现Stringer接口
func (b *BST) String() string {
	return b.root.generateBSTString(0)
}

func (b *BST) Minimum() (e int, err error) {
	if b.IsEmpty() {
		return 0, fmt.Errorf("cat search minimum, because BST is empty")
	}
	return b.root.minimum().E, nil
}

func (b *BST) Maximum() (e int, err error) {
	if b.IsEmpty() {
		return 0, fmt.Errorf("cat search maximum, because BST is empty")
	}
	return b.root.maximum().E, nil
}

func (b *BST) RemoveMin() (e int, err error) {
	minimum, err := b.Minimum()
	if err != nil {
		return 0, err
	}
	b.root = b.root.removeMin()
	b.size--
	return minimum, nil
}

func (b *BST) RemoveMax() (e int, err error) {
	maximum, err := b.Maximum()
	if err != nil {
		return 0, err
	}
	b.root = b.root.removeMax()
	b.size--
	return maximum, nil
}

func (b *BST) Remove(e int) error {
	has := b.root.has(e)
	if !has {
		return fmt.Errorf("remote element failed: cant has element")
	}
	b.root = b.root.remove(e)
	b.size--
	return nil
}
