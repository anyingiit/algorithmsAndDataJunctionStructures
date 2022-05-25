package Queue

import "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Queue/LinkQueue"

type Queue interface {
	// Enqueue push a element to Queue
	Enqueue(e interface{})
	// Dequeue pop head element
	Dequeue() interface{}
	// GetFront return head element
	GetFront() interface{}
	Size() int
	IsEmpty() bool
}

func NewQueue() Queue {
	return LinkQueue.NewLinkQueue()
}
