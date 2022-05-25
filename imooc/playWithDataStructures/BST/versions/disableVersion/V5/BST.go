package V5

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/versions/disableVersion/V5/E"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Queue"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack"
	"fmt"
)

type node struct {
	e     E.E
	Left  *node
	Right *node
}

func newNode(e E.E) *node {
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

func (n *node) Remove(e E.E) (successor *node, deletedNode *node) {
	if n == nil {
		return nil, nil
	}

	if n.e.Id() > e.Id() {
		n.Left, deletedNode = n.Left.Remove(e)
	} else if n.e.Id() < e.Id() {
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
func (b *BST) Push(e E.E) (isSuccess bool) {
	isSuccess = false

	var push func(node *node) *node
	push = func(node *node) *node {
		if node == nil {
			b.size++
			isSuccess = true
			return newNode(e)
		}
		if node.e.Id() > e.Id() {
			node.Left = push(node.Left)
		} else if node.e.Id() < e.Id() {
			node.Right = push(node.Right)
		}
		return node
	}

	b.root = push(b.root)
	return isSuccess
}

func (b *BST) Has(e E.E) bool {
	var has func(node *node) bool
	has = func(node *node) bool {
		if node == nil {
			return false
		}
		if node.e.Id() > e.Id() {
			return has(node.Left)
		} else if node.e.Id() < e.Id() {
			return has(node.Right)
		}
		return true
	}

	return has(b.root)
}

func (b *BST) GetPreOrder() (result []E.E) {
	var preOrder func(node *node)
	preOrder = func(node *node) {
		if node == nil {
			return
		}
		result = append(result, node.e)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(b.root)

	return result
}
func (b *BST) GetInOrder() (result []E.E) {
	var inOrder func(node *node)
	inOrder = func(node *node) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		result = append(result, node.e)
		inOrder(node.Right)
	}
	inOrder(b.root)

	return result
}
func (b *BST) GetPostOrder() (result []E.E) {
	var postOrder func(node *node)
	postOrder = func(node *node) {
		if node == nil {
			return
		}
		postOrder(node.Left)
		postOrder(node.Right)
		result = append(result, node.e)
	}
	postOrder(b.root)

	return result
}

func (b *BST) PrintPreOrder() {
	bstPreOrder := b.GetPreOrder()
	for _, e := range bstPreOrder {
		fmt.Println(e)
	}
}
func (b *BST) PrintInOrder() {
	bstInOrder := b.GetInOrder()
	for _, e := range bstInOrder {
		fmt.Println(e)
	}
}
func (b *BST) PrintPostOrder() {
	bstPostOrder := b.GetPostOrder()
	for _, e := range bstPostOrder {
		fmt.Println(e)
	}
}

func (b *BST) GetPreOrderNR() (result []E.E) {
	stack := Stack.NewStack()
	stack.Enqueue(b.root)
	for !stack.IsEmpty() {
		curNode := stack.Dequeue().(*node)
		if curNode == nil {
			continue
		}
		result = append(result, curNode.e)
		stack.Enqueue(curNode.Right)
		stack.Enqueue(curNode.Left)
	}
	return result
}
func (b *BST) GetLevelOrderNR() (result []E.E) {
	queue := Queue.NewQueue()
	queue.Enqueue(b.root)
	for !queue.IsEmpty() {
		curNode := queue.Dequeue().(*node)
		if curNode == nil {
			continue
		}
		result = append(result, curNode.e)
		queue.Enqueue(curNode.Left)
		queue.Enqueue(curNode.Right)
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
		return generateDepthString(depth) + node.e.String() + "\n" + getPreOrderFormatString(node.Left, depth+1) + getPreOrderFormatString(node.Right, depth+1)
	}

	return getPreOrderFormatString(b.root, 0)
}

func (b *BST) String() string {
	return b.GetPreOrderFormatString()
}

func (b *BST) Minimum() (minimumElement E.E, err error) { //TODO: 到底是返回节点还是值?
	if b.IsEmpty() {
		return nil, fmt.Errorf("get Minimum failed: the BST_V5 is empty")
	}
	return b.root.Minimum().e, nil
}
func (b *BST) Maximum() (maximumElement E.E, err error) { //TODO: 到底是返回节点还是值?
	if b.IsEmpty() {
		return nil, fmt.Errorf("get Maximum failed: the BST_V5 is empty")
	}
	return b.root.Maximum().e, nil
}

func (b *BST) RemoveMin() (deletedElement E.E, err error) {
	if b.IsEmpty() {
		return nil, fmt.Errorf("cant remove min: BST_V5 is empty")
	}

	var deletedNode *node
	b.root, deletedNode = b.root.RemoveMin()

	if deletedNode == nil {
		return nil, fmt.Errorf("cannot remove min: other error")
	}
	b.size--
	return deletedNode.e, nil
}
func (b *BST) RemoveMax() (deletedElement E.E, err error) {
	if b.IsEmpty() {
		return nil, fmt.Errorf("cant remove min: BST_V5 is empty")
	}

	var deletedNode *node
	b.root, deletedNode = b.root.RemoveMax()

	if deletedNode == nil {
		return nil, fmt.Errorf("cannot remove max: other error")
	}
	b.size--
	return deletedNode.e, nil
}

func (b *BST) Remove(e E.E) error { //TODO: logic error
	var deletedNode *node
	b.root, deletedNode = b.root.Remove(e)

	if deletedNode == nil {
		return fmt.Errorf("cannot remove element %d: cant find this element", e)
	}

	b.size--

	return nil
}
