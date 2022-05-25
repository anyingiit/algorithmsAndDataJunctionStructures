package MaxHeapPriorityQueue

import "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/MaxHeap"

type MaxHeapPriorityQueue struct {
	maxHeap MaxHeap.MaxHeap
}

func NewPriorityQueue() *MaxHeapPriorityQueue {
	return &MaxHeapPriorityQueue{MaxHeap.NewMaxHeap()}
}

func (p *MaxHeapPriorityQueue) Enqueue(e int) {
	p.maxHeap.Add(e)
}

func (p *MaxHeapPriorityQueue) Dequeue() (int, error) {
	return p.maxHeap.ExtractMax()
}

func (p *MaxHeapPriorityQueue) GetFront() (int, error) {
	return p.maxHeap.FindMax()
}

func (p *MaxHeapPriorityQueue) Size() int {
	return p.maxHeap.Size()
}

func (p *MaxHeapPriorityQueue) IsEmpty() bool {
	return p.maxHeap.IsEmpty()
}
