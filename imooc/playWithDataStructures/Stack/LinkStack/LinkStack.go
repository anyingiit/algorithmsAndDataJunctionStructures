package LinkStack

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

func (n *node) setE(e interface{}) *node {
	n.e = e
	return n
}

func (n *node) setNext(next *node) *node {
	n.next = next
	return n
}

func (n *node) pushHead(e interface{}) *node {
	return newNode(e).setNext(n)
}

type LinkStack struct {
	head *node
	size int
}

func NewLinkStack() *LinkStack {
	return &LinkStack{
		head: nil,
		size: 0,
	}
}

func (l *LinkStack) Size() int {
	return l.size
}

func (l *LinkStack) IsEmpty() bool {
	return l.Size() == 0
}

// Push element to head and LinkStack size++
func (l *LinkStack) Push(e interface{}) {
	l.head = l.head.pushHead(e)
	l.size++
}

// Pop
//
// Get head element and LinkStack size--
//
// Warning: You need make sure the LinkStack not is empty
func (l *LinkStack) Pop() interface{} {
	e := l.head.e
	l.head = l.head.next
	l.size--
	return e
}
