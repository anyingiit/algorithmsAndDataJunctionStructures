package BST

import (
	"fmt"
	"strconv"
)

type node struct {
	e     int
	Left  *node
	Right *node
}

func newNode(e int) *node {
	return &node{
		e:     e,
		Left:  nil,
		Right: nil,
	}
}

// BST 不重复的二分搜索树
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

func (b *BST) Size() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.Size() == 0
}

func (b *BST) push(node *node, e int) *node {
	if node == nil {
		b.size++
		return newNode(e)
	}
	if node.e > e {
		node.Left = b.push(node.Left, e)
	} else if node.e < e {
		node.Right = b.push(node.Right, e)
	}
	return node
}

func (b *BST) Push(e int) {
	b.root = b.push(b.root, e)
}

func (b *BST) has(node *node, e int) bool {
	if node == nil {
		return false
	}
	if node.e > e {
		return b.has(node.Left, e)
	} else if node.e < e {
		return b.has(node.Right, e)
	}
	return true
}

func (b *BST) Has(e int) bool {
	return b.has(b.root, e)
}

func (b *BST) preOrder(node *node) {
	if node == nil {
		return
	}
	fmt.Println(node.e)
	b.preOrder(node.Left)
	b.preOrder(node.Right)
}

func (b *BST) PreOrder() {
	b.preOrder(b.root)
}

func (b *BST) inOrder(node *node) {
	if node == nil {
		return
	}
	b.inOrder(node.Left)
	fmt.Println(node.e)
	b.inOrder(node.Right)
}

func (b *BST) InOrder() {
	b.inOrder(b.root)
}

func (b *BST) postOrder(node *node) {
	if node == nil {
		return
	}
	b.postOrder(node.Left)
	b.postOrder(node.Right)
	fmt.Println(node.e)
}

func (b *BST) PostOrder() {
	b.postOrder(b.root)
}

// TODO: func (b *BST) InOrderNR(){}

func (b *BST) generateBSTString(node *node, depth int) (result string) {
	if node == nil {
		return b.generateDepthString(depth) + "null" + "\n"
	}
	return b.generateDepthString(depth) + strconv.Itoa(node.e) + "\n" + b.generateBSTString(node.Left, depth+1) + b.generateBSTString(node.Right, depth+1)
}

func (b *BST) generateDepthString(depth int) (depthString string) {
	for i := 0; i < depth; i++ {
		depthString += "--"
	}
	return depthString
}

func (b *BST) GenerateBSTString() (result string) {
	return b.generateBSTString(b.root, 0)
}

func (b *BST) String() string {
	return b.GenerateBSTString()
}

func (b *BST) minimum(node *node) *node {
	if node.Left == nil {
		return node
	}
	return b.minimum(node.Left)
}

func (b *BST) Minimum() (minimum int, err error) { //TODO: 到底是返回节点还是值?
	if b.IsEmpty() {
		return 0, fmt.Errorf("get minimum failed: the BST is empty")
	}
	return b.minimum(b.root).e, nil
}

func (b *BST) maximum(node *node) *node {
	if node.Right == nil {
		return node
	}
	return b.maximum(node.Right)
}

func (b *BST) Maximum() (maximum int, err error) { //TODO: 到底是返回节点还是值?
	if b.IsEmpty() {
		return 0, fmt.Errorf("get maximum failed: the BST is empty")
	}
	return b.maximum(b.root).e, nil
}

func (b *BST) removeMin(node *node) *node { //TODO: logic err
	if node.Left == nil {
		b.size--
		return node.Right
	}
	node.Left = b.removeMin(node.Left)
	return node
}

func (b *BST) RemoveMin() (e int, err error) {
	minimum, err := b.Minimum()
	if err != nil {
		return 0, err
	}
	b.root = b.removeMin(b.root)
	return minimum, nil
}

func (b *BST) removeMax(node *node) *node {
	if node.Right == nil {
		b.size--
		return node.Left
	}
	node.Right = b.removeMax(node.Right)
	return node
}

func (b *BST) RemoveMax() (e int, err error) {
	maximum, err := b.Maximum()
	if err != nil {
		return 0, err
	}
	b.root = b.removeMax(b.root)
	return maximum, nil
}

func (b *BST) remove(node *node, e int) *node { //TODO: logic error
	if node == nil {
		return nil
	}

	if node.e > e {
		node.Left = b.remove(node.Left, e)
	} else if node.e < e {
		node.Right = b.remove(node.Right, e)
	} else { // node.e == e
		if node.Left == nil {
			b.size--
			return node.Right
		} else if node.Right == nil {
			b.size--
			return node.Left
		} else {
			successor := b.minimum(node.Right)
			// 这里踩了个坑, 注意注意!! !!!!必须要先对右节点进行赋值!!!!, 因为当node.Right没有孩子的情况下, successor和node.Right指向的节点是相同的, 在这种情况下, 如果先对successor设置左节点, 就相当于`node.Right.Left=node.Left`, 那么当进行node.Right为根的删除操作时, 函数会错误的视为node.Left(或子树中)为node.Right子树的最小的节点作为返回, 那么返回值就是node.Left, 返回值被赋予给了successor.Right. 假设一种情况: 当前节点的左孩子是叶子节点时, 先进行`successor.Left = node.Left`, 在进行`successor.Right = b.removeMin(node.Right)`时, successor.Left的结果是正确的的, 但是successor.Right的结果一定是successor.Left, 这是错误的, 期望successor.Right为空
			//		if node.Right.Left == nil && node.Right.Left == nil{
			//			successor.Left = node.Left
			//			successor.Right = b.removeMin(node.Right)
			//			// [actually]
			//			//		successor.Left == node.Left
			//			//		successor.Right == node.Left
			//			// [expect]
			//			//		success.Left == node.Left
			//			//		success.Right == nil
			//		}
			successor.Right = b.removeMin(node.Right) // removeMin 本身会进行size--, 所以不需要进行size--了
			successor.Left = node.Left
			return successor
		}
	}
	return node
}

func (b *BST) Remove(e int) error { //TODO: logic error
	originSize := b.Size()
	b.root = b.remove(b.root, e)
	if b.Size() != originSize-1 {
		return fmt.Errorf("failed remove %d", e)
	}
	return nil
}
