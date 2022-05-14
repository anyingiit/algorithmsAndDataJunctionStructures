package BST

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Queue"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack"
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

func (n *node) Minimum() (result *node) {
	if n == nil {
		return nil
	}
	if n.Left == nil {
		return n
	}
	return n.Left.Minimum()
}

func (n *node) Maximum() (result *node) {
	if n == nil {
		return nil
	}
	if n.Right == nil {
		return n
	}
	return n.Right.Maximum()
}

func (n *node) RemoveMin() (successor *node, deletedNode *node) {
	if n == nil {
		return nil, nil
	}
	if n.Left == nil {
		return n.Right, n
	}
	n.Left, deletedNode = n.Left.RemoveMin()
	return n, deletedNode
}

func (n *node) RemoveMax() (successor *node, deletedNode *node) {
	if n == nil {
		return nil, nil
	}
	if n.Right == nil {
		return n.Left, n
	}
	n.Right, deletedNode = n.Right.RemoveMax()
	return n, deletedNode
}

func (n *node) Remove(e int) (successor *node, deletedNode *node) {
	if n == nil {
		return nil, nil
	}

	if n.e > e {
		n.Left, deletedNode = n.Left.Remove(e)
	} else if n.e < e {
		n.Right, deletedNode = n.Right.Remove(e)
	} else { // targetNode.e == e
		if n.Left == nil {
			return n.Right, n
		} else if n.Right == nil {
			return n.Left, n
		} else { // n.Left != nil && n.Right != nil
			n.Right, successor = n.Right.RemoveMin()
			successor.Left = n.Left
			successor.Right = n.Right
			return successor, n
		}
	}
	return n, deletedNode
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

// Push element to BST
//
// warning: this BST cannot support repeat element, so this method well return do you really success push to this BST
// if this method return is false, then because you try push to a repeat element to this BST
func (b *BST) Push(e int) (isSuccess bool) {
	isSuccess = false

	var push func(node *node) *node
	push = func(node *node) *node {
		if node == nil {
			b.size++
			isSuccess = true
			return newNode(e)
		}
		if node.e > e {
			node.Left = push(node.Left)
		} else if node.e < e {
			node.Right = push(node.Right)
		}
		return node
	}

	b.root = push(b.root)
	return isSuccess
}

func (b *BST) Has(e int) bool {
	var has func(node *node) bool
	has = func(node *node) bool {
		if node == nil {
			return false
		}
		if node.e > e {
			return has(node.Left)
		} else if node.e < e {
			return has(node.Right)
		}
		return true
	}

	return has(b.root)
}

func (b *BST) PrintPreOrder() {
	var preOrder func(node *node)
	preOrder = func(node *node) {
		if node == nil {
			return
		}
		fmt.Println(node.e)
		preOrder(node.Left)
		preOrder(node.Right)
	}

	preOrder(b.root)
}

func (b *BST) PrintInOrder() {
	var inOrder func(node *node)
	inOrder = func(node *node) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		fmt.Println(node.e)
		inOrder(node.Right)
	}
	inOrder(b.root)
}

func (b *BST) PrintPostOrder() {
	var postOrder func(node *node)
	postOrder = func(node *node) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		fmt.Println(node.e)
	}

	postOrder(b.root)
}

func (b *BST) GetPreOrderNR() (result []int) {
	stack := Stack.NewStack()
	stack.Push(b.root)
	for !stack.IsEmpty() {
		curNode := stack.Pop().(*node)
		if curNode == nil {
			continue
		}
		result = append(result, curNode.e)
		stack.Push(curNode.Right)
		stack.Push(curNode.Left)
	}
	return result
}

func (b *BST) GetLevelOrderNR() (result []int) {
	queue := Queue.NewQueue()
	queue.Push(b.root)
	for !queue.IsEmpty() {
		curNode := queue.Pop().(*node)
		if curNode == nil {
			continue
		}
		result = append(result, curNode.e)
		queue.Push(curNode.Left)
		queue.Push(curNode.Right)
	}
	return result
}

func (b *BST) GetPreOrderFormatString() (result string) {
	generateDepthString := func(depth int) (depthString string) {
		for i := 0; i < depth; i++ {
			depthString += "--"
		}
		return depthString
	}
	var getPreOrderFormatString func(node *node, depth int) (result string)
	getPreOrderFormatString = func(node *node, depth int) (result string) {
		if node == nil {
			return generateDepthString(depth) + "null" + "\n"
		}
		return generateDepthString(depth) + strconv.Itoa(node.e) + "\n" + getPreOrderFormatString(node.Left, depth+1) + getPreOrderFormatString(node.Right, depth+1)
	}

	return getPreOrderFormatString(b.root, 0)
}

func (b *BST) String() string {
	return b.GetPreOrderFormatString()
}

func (b *BST) Minimum() (minimumElement int, err error) { //TODO: 到底是返回节点还是值?
	if b.IsEmpty() {
		return 0, fmt.Errorf("get Minimum failed: the BST is empty")
	}
	return b.root.Minimum().e, nil
}

func (b *BST) Maximum() (maximumElement int, err error) { //TODO: 到底是返回节点还是值?
	if b.IsEmpty() {
		return 0, fmt.Errorf("get Maximum failed: the BST is empty")
	}
	return b.root.Maximum().e, nil
}

func (b *BST) RemoveMin() (deletedNode *node, err error) {
	if b.IsEmpty() {
		return nil, fmt.Errorf("cant remove min: BST is empty")
	}

	b.root, deletedNode = b.root.RemoveMin()

	if deletedNode == nil {
		return nil, fmt.Errorf("cannot remove min: other error")
	}
	b.size--
	return deletedNode, nil
}

func (b *BST) RemoveMax() (deletedNode *node, err error) {
	if b.IsEmpty() {
		return nil, fmt.Errorf("cant remove min: BST is empty")
	}

	b.root, deletedNode = b.root.RemoveMax()

	if deletedNode == nil {
		return nil, fmt.Errorf("cannot remove max: other error")
	}
	b.size--
	return deletedNode, nil
}

func (b *BST) Remove(e int) error { //TODO: logic error
	var deletedNode *node
	b.root, deletedNode = b.root.Remove(e)

	if deletedNode == nil {
		return fmt.Errorf("cannot remove element %d: cant find this element", e)
	}

	b.size--

	return nil
}
