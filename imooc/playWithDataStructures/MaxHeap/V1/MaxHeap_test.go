package V1

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/OrderSlice"
	"math/rand"
	"testing"
)

func TestMaxHeap_parent(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]int{0, 0, 0, 1, 1, 2, 2},
		},
	}
	maxHeap := NewMaxHeap()
	for _, tt := range tests {
		for i, e := range tt.input {
			if maxHeap.parent(e) != tt.expect[i] {
				t.Errorf("test parent failed: expect %d actually %d", tt.expect[i], maxHeap.parent(e))
			}
		}
	}
}
func TestMaxHeap_leftChild(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]int{1, 3, 5, 7, 9, 11, 13},
		},
	}
	maxHeap := NewMaxHeap()
	for _, tt := range tests {
		for i, e := range tt.input {
			if maxHeap.leftChild(e) != tt.expect[i] {
				t.Errorf("test leftChild failed: expect %d actually %d", tt.expect[i], maxHeap.parent(e))
			}
		}
	}
}
func TestMaxHeap_rightChild(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]int{2, 4, 6, 8, 10, 12, 14},
		},
	}
	maxHeap := NewMaxHeap()
	for _, tt := range tests {
		for i, e := range tt.input {
			if maxHeap.rightChild(e) != tt.expect[i] {
				t.Errorf("test leftChild failed: expect %d actually %d", tt.expect[i], maxHeap.parent(e))
			}
		}
	}
}

func TestMaxHeap_Add(t *testing.T) {
	type test struct {
		input []int
	}
	var tests []test

	tests = append(tests, func() (tt test) {
		for i := 0; i < 1000; i++ {
			tt.input = append(tt.input, rand.Intn(700))
		}
		return tt
	}())

	for _, tt := range tests {
		maxHeap := NewMaxHeap()

		for _, e := range tt.input {
			maxHeap.Add(e)
		}

		for i := maxHeap.Size() - 1; i > 0; i-- {
			if maxHeap.data[i] > maxHeap.data[maxHeap.parent(i)] {
				t.Errorf("test add failed: child bigger to parent")
			}
		}
	}
}

func TestMaxHeap_ExtractMax(t *testing.T) {
	type test struct {
		input []int
	}
	var tests []test

	tests = append(tests, func() (tt test) {
		for i := 0; i < 1000; i++ {
			tt.input = append(tt.input, rand.Intn(700))
		}
		return tt
	}())

	for _, tt := range tests {
		maxHeap := NewMaxHeap()

		for _, e := range tt.input {
			maxHeap.Add(e)
		}

		var actually []int
		for !maxHeap.IsEmpty() {
			extractMax, err := maxHeap.ExtractMax()
			if err != nil {
				t.Errorf("test extract max failed: excute ExtractMax has err %s", err.Error())
			}
			actually = append(actually, extractMax)
		}

		if !OrderSlice.IsASC(actually) {
			t.Errorf("test extract max failed: actually not is order by ASC")
		}
	}
}
