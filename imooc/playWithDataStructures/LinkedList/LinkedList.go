package LinkedList

import "fmt"

type node struct {
	e    int
	next *node
}

func newNode(e int) *node {
	return &node{
		e:    e,
		next: nil,
	}
}

func (n *node) add(e int) *node {
	if n == nil {
		return newNode(e)
	}
	n.next = n.next.add(e)
	return n
}

func (n *node) has(e int) bool {
	if n == nil {
		return false
	}
	if n.e == e {
		return true
	}
	return n.next.has(e)
}

func (n *node) del(e int) (*node, int, error) {
	if n == nil {
		return nil, 0, fmt.Errorf("not has target element %d", e)
	}
	if n.e == e {
		return n.next, n.e, nil
	}
	return n.next.del(e)
}

func (n *node) set(e int, newValue int) error {
	if n == nil {
		return fmt.Errorf("can not set element %d: not has target element %d", e, e)
	}
	if n.e == e {
		n.e = newValue
		return nil
	}
	return n.next.set(e, newValue)
}

// getAll 由于链表本身是不保证顺序的, 所以返回的结果是倒着的也没有问题
func (n *node) getAll() (result []int) {
	if n == nil {
		return result
	}
	return append(n.next.getAll(), n.e)
}

type LinkedList struct {
	root *node
	size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(e int) {
	l.root = l.root.add(e)
	l.size++
}

func (l *LinkedList) Has(e int) bool {
	return l.root.has(e)
}

func (l *LinkedList) Size() int {
	return l.size
}

func (l *LinkedList) Del(e int) (int, error) {
	root, e, err := l.root.del(e)
	l.root = root
	if err == nil {
		l.size--
	}
	return e, err
}

func (l *LinkedList) Set(e, newValue int) error {
	return l.root.set(e, newValue)
}

// GetAll 获得所有元素. 但是顺序从存储结构上来说, 是倒序的.
//
// 存储顺序: 1 -> 5 -> 2 -> 10 -> 3
//
// 返回顺序: 3, 10, 2, 5, 1
func (l *LinkedList) GetAll() (result []int) {
	return l.root.getAll()
}
