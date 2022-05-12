package V2

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BubbleSort"
	"fmt"
	"math/rand"
	"testing"
)

func TestBST_RemoveMin(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}
	bst := NewBST()
	for _, e := range tests {
		bst.Push(e)
	}

	//fmt.Println(bst)

	var bstOut []int
	for !bst.IsEmpty() {
		min, err := bst.RemoveMin()
		if err != nil {
			t.Errorf("delete bst min element failed: %s", err.Error())
		}
		bstOut = append(bstOut, min)
	}

	if len(bstOut) != len(tests) {
		t.Errorf("bst out len check failed, expect %d actually %d", len(tests), len(bstOut))
	}

	BubbleSort.SortByDESC(tests)

	for i, expect := range tests {
		if bstOut[i] != expect {
			t.Errorf("check bst remove min failed, expect %d, actually %d", expect, bstOut[i])
		}
	}

}

func TestBST_RemoveMax(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}
	bst := NewBST()
	for _, e := range tests {
		bst.Push(e)
	}

	var bstOut []int
	for !bst.IsEmpty() {
		max, err := bst.RemoveMax()
		if err != nil {
			t.Errorf("delete bst max element failed: %s", err.Error())
		}
		bstOut = append(bstOut, max)
	}

	if len(bstOut) != len(tests) {
		t.Errorf("bst out len check failed, expect %d actually %d", len(tests), len(bstOut))
	}

	BubbleSort.SortByASC(tests)

	for i, expect := range tests {
		if bstOut[i] != expect {
			t.Errorf("check bst remove max failed, expect %d, actually %d", expect, bstOut[i])
		}
	}

}

func TestBST_Push(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(100))
	}

	bst := NewBST()
	for _, e := range tests {
		bst.Push(e)
	}

	fmt.Println(bst)
	expectSize := len(tests)
	actuallySize := 0
	for i := 0; i < expectSize; i++ {
		if _, err := bst.RemoveMin(); err == nil {
			actuallySize++
		}
	}

	fmt.Println(bst)

	if expectSize != actuallySize {
		t.Errorf("test BST Push failed: check BST size failed, expectSize %d actuallySize %d", expectSize, actuallySize)
	}
}

func TestBST_Remove(t *testing.T) {
	bst := NewBST()
	err := bst.Remove(233)
	if err == nil {
		t.Errorf("test space BST failed, expect err not is nil, actually err is nil")
	}
	bst.Push(10)
	err = bst.Remove(10)
	if err != nil {
		t.Errorf("test BST remove element failed: %s", err.Error())
	}

	if bst.Size() != 0 {
		t.Errorf("test BST remove element failed: BST size check failed, actually size is %d, expect bst size is 0", bst.Size())
	}

	if bst.Has(10) {
		t.Errorf("test BST remove element failed: actually bst has element 10, expect not have element 10")
	}

	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}
	for _, e := range tests {
		bst.Push(e)
	}
	fmt.Printf("tests size %d, bst size %d", len(tests), bst.Size())
	//for _, e := range tests {
	//	err := bst.Remove(e)
	//	if err != nil {
	//		t.Errorf("test BST remove element failed: expect remove element not have err, actuall has err %s", err.Error())
	//	}
	//}
}
