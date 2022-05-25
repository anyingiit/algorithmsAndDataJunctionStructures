package PriorityQueue

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/PriorityQueue/MaxHeapPriorityQueue"
)

type PriorityQueue interface {
	// Enqueue push a element to Queue
	Enqueue(e int)
	// Dequeue pop head element
	// if PriorityQueue is empty then return error
	Dequeue() (int, error)
	// GetFront return head element
	// if PriorityQueue is empty then return error
	GetFront() (int, error)
	Size() int
	IsEmpty() bool
}

func NewPriorityQueue() PriorityQueue {
	return MaxHeapPriorityQueue.NewPriorityQueue()
}
