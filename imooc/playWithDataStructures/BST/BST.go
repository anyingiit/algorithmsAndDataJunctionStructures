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

func (b *BST) Push(e int) {
	var push func(node *node) *node
	push = func(node *node) *node {
		if node == nil {
			b.size++
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

//// RemoveTargetNodeMin 如果随便传入一个node, 而且对这个node执行删除min成功了, 会对本二分搜索树size--, 这不合理
//func (b *BST) RemoveTargetNodeMin(targetNode *node) (deletedNode *node, err error) {
//	if targetNode == nil {
//		return nil, fmt.Errorf("cant remove target node min: target node is nil")
//	}
//
//	targetNode, deletedNode = targetNode.RemoveMin()
//
//	if deletedNode == nil {
//		return nil, fmt.Errorf("cant remove target node min: other err")
//	}
//	b.size--
//	return targetNode, nil
//}
//
//// RemoveTargetNodeMax 如果随便传入一个node, 而且对这个node执行删除min成功了, 会对本二分搜索树size--, 这不合理
//func (b *BST) RemoveTargetNodeMax(targetNode *node) (deletedNode *node, err error) {
//	if targetNode == nil {
//		return nil, fmt.Errorf("cant remove target node max: target node is nil")
//	}
//
//	targetNode, deletedNode = targetNode.RemoveMax()
//
//	if deletedNode == nil {
//		return nil, fmt.Errorf("cant remove target node max: other err")
//	}
//	b.size--
//	return targetNode, nil
//}

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

//func (b *BST) remove(targetNode *node, e int) *node { //TODO: logic error
//	if targetNode == nil {
//		return nil
//	}
//
//	if targetNode.e > e {
//		targetNode.Left = b.remove(targetNode.Left, e)
//	} else if targetNode.e < e {
//		targetNode.Right = b.remove(targetNode.Right, e)
//	} else { // targetNode.e == e
//		if targetNode.Left == nil {
//			b.size--
//			return targetNode.Right
//		} else if targetNode.Right == nil {
//			b.size--
//			return targetNode.Left
//		} else {
//			var successor *node
//			targetNode.Right, successor = targetNode.Right.RemoveMin()
//			b.size--
//			// 这里踩了个坑, 注意注意!! !!!!必须要先对右节点进行赋值!!!!, 因为当node.Right没有孩子的情况下, successor和node.Right指向的节点是相同的, 在这种情况下, 如果先对successor设置左节点, 就相当于`targetNode.Right.Left=targetNode.Left`, 那么当进行node.Right为根的删除操作时, 函数会错误的视为node.Left(或子树中)为node.Right子树的最小的节点作为返回, 那么返回值就是node.Left, 返回值被赋予给了successor.Right. 假设一种情况: 当前节点的左孩子是叶子节点时, 先进行`successor.Left = targetNode.Left`, 在进行`successor.Right = b.RemoveMin(targetNode.Right)`时, successor.Left的结果是正确的的, 但是successor.Right的结果一定是successor.Left, 这是错误的, 期望successor.Right为空
//			//		if targetNode.Right.Left == nil && targetNode.Right.Left == nil{
//			//			successor.Left = targetNode.Left
//			//			successor.Right = b.RemoveMin(targetNode.Right)
//			//			// [actually]
//			//			//		successor.Left == targetNode.Left
//			//			//		successor.Right == targetNode.Left
//			//			// [expect]
//			//			//		success.Left == targetNode.Left
//			//			//		success.Right == nil
//			//		}
//			successor.Right = targetNode.Right
//			successor.Left = targetNode.Left
//			return successor
//		}
//	}
//	return targetNode
//}

func (b *BST) Remove(e int) error { //TODO: logic error
	var deletedNode *node
	b.root, deletedNode = b.root.Remove(e)

	if deletedNode == nil {
		return fmt.Errorf("cannot remove element %d: cant find this element", e)
	}

	b.size--

	return nil
}
