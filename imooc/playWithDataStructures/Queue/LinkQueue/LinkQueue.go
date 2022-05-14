package LinkQueue

type node struct {
	e    interface{}
	next *node
}

func newNode(e interface{}) *node {
	return &node{
		e:    e,
		next: nil,
	}
}

func (n *node) pushBack(e interface{}) *node {
	if n == nil {
		return newNode(e)
	}
	n.next = n.next.pushBack(e)
	return n
}

type LinkQueue struct {
	head *node
	size int
}

func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		head: nil,
		size: 0,
	}
}

func (l *LinkQueue) Size() int {
	return l.size
}

func (l *LinkQueue) IsEmpty() bool {
	return l.Size() == 0
}

// Push element to back and LinkQueue size++
func (l *LinkQueue) Push(e interface{}) {
	l.head = l.head.pushBack(e)
	l.size++
}

// Pop
//
// Get head element and LinkQueue size--
//
// Warning: You need make sure the LinkQueue not is empty
func (l *LinkQueue) Pop() interface{} {
	e := l.head.e
	l.head = l.head.next
	l.size--
	return e
}
