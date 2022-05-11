package Queue

import (
	"fmt"
)

// Queue 后进先出
type Queue []interface{}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(e interface{}) {
	*q = append(*q, e)
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

// Pop 如果队列为空还尝试出队列的话, 程序就挂了
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) Test() {
	fmt.Printf("q is %T\n", q)   //q is *Queue.Queue
	fmt.Printf("*q is %T\n", *q) //*q is Queue.Queue

	// 由此可以得出结论, 指针接收者的方法时, 取q就是指针, 取*q就是值
}
