package MaxHeap

import "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/MaxHeap/V1"

type MaxHeap interface {
	// Size return the MaxHeap size
	Size() int
	// IsEmpty if the MaxHeap is empty then return true else return false
	IsEmpty() bool
	// Add for MaxHeap add any element
	Add(e int)
	// ExtractMax pop max element
	// if the MaxHeap is empty then return err
	ExtractMax() (int, error)
	// FindMax return max element
	// if the MaxHeap is empty then return err
	FindMax() (int, error)
	// Replace pop max element and add an element
	// if the MaxHeap is empty then return err
	Replace(e int) (int, error)
}

func NewMaxHeap() MaxHeap {
	return V1.NewMaxHeap()
}

func NewMaxHeapForArr(x []int) MaxHeap {
	return V1.NewMaxHeapForArr(x)
}
