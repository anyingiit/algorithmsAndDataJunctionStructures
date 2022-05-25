package SliceQueue

import (
	"fmt"
)

// SliceQueue 后进先出
type SliceQueue []interface{}

func (q *SliceQueue) GetFront() interface{} {
	return (*q)[0]
}

func NewSliceQueue() *SliceQueue {
	return &SliceQueue{}
}

// Enqueue push element to back and SliceQueue size++
func (q *SliceQueue) Enqueue(e interface{}) {
	*q = append(*q, e)
}

func (q *SliceQueue) Size() int {
	return len(*q)
}

func (q *SliceQueue) IsEmpty() bool {
	return q.Size() == 0
}

// Dequeue
//
// Get head element and SliceQueue size--
//
// Warning: You need make sure the LinkQueue not is empty
func (q *SliceQueue) Dequeue() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *SliceQueue) Test() {
	fmt.Printf("q is %T\n", q)   //q is *SliceQueue.SliceQueue
	fmt.Printf("*q is %T\n", *q) //*q is SliceQueue.SliceQueue

	// 由此可以得出结论, 指针接收者的方法时, 取q就是指针, 取*q就是值
}
