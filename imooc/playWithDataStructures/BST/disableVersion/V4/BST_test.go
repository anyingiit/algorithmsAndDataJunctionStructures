package V4

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BubbleSort"
	"fmt"
	"math/rand"
	"testing"
)

func TestBST_Size(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(500))
	}

	noRepeat := make(map[int]bool)

	for _, e := range tests {
		noRepeat[e] = true
	}

	noRepeatCount := len(noRepeat)
	//t.Logf("tests len %d, no repeat num count %d\n", len(tests), noRepeatCount)

	bst := NewBST()
	for _, e := range tests {
		bst.Push(e)
	}

	if bst.Size() != noRepeatCount {
		t.Errorf("check size failed, actually %d expect %d", bst.Size(), noRepeatCount)
	}
}

func TestBST_Has(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(500))
	}

	noRepeat := make(map[int]bool)

	for _, e := range tests {
		noRepeat[e] = true
	}
	//t.Logf("tests len %d, no repeat num count %d\n", len(tests), len(noRepeat))

	bst := NewBST()

	for key := range noRepeat {
		bst.Push(key)
	}

	for key := range noRepeat {
		if !bst.Has(key) {
			t.Errorf("BST Has test failed, expect has %d but not have", key)
		}
	}

}

func TestBST_String(t *testing.T) {
	bst := NewBST()

	bst.Push(41)
	bst.Push(22)
	bst.Push(58)
	bst.Push(15)
	bst.Push(33)
	bst.Push(13)
	bst.Push(50)
	bst.Push(42)
	bst.Push(53)

	expect := "41\n--22\n----15\n------13\n--------null\n--------null\n------null\n----33\n------null\n------null\n--58\n----50\n------42\n--------null\n--------null\n------53\n--------null\n--------null\n----null\n"
	if fmt.Sprint(bst) != expect {
		t.Errorf("check BST String failed:\nexpect\n\n\"%s\"\n\nactually\n\"%s\"\n\n", expect, fmt.Sprint(bst))
	}

}

func TestBST_RemoveMin(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		//{
		//	[]int{5},
		//	[]int{5},
		//},
		{
			[]int{5, 8, 6, 1, 4},
			[]int{1, 4, 5, 6, 8},
		},
		{
			[]int{5, 8, 6, 1, 4, 6},
			[]int{1, 4, 5, 6, 8},
		},
	}

	tests = append(tests, func() test {
		//t.Log("now created lang random data")

		var input []int
		for i := 0; i < 1000; i++ {
			input = append(input, rand.Intn(10000))
		}

		var noRepeat [10000]int
		for _, e := range input {
			noRepeat[e] = 1
		}

		var expect []int
		for i, e := range noRepeat {
			if e == 1 {
				expect = append(expect, i)
			}
		}

		BubbleSort.SortByDESC(expect)

		//t.Logf("lang random data:\ninput: %v\nexpect: %v", input, expect)
		return test{
			input:  input,
			expect: expect,
		}
	}())

	//t.Log("now created lang random data")
	//for _, tt := range tests {
	//	fmt.Println(tt.input)
	//	fmt.Println(tt.expect)
	//	fmt.Println()
	//}

	for _, tt := range tests {
		bst := NewBST()
		for _, e := range tt.input {
			bst.Push(e)
		}
		var actually []int
		for !bst.IsEmpty() {
			min, err := bst.RemoveMin()
			if err != nil {
				t.Errorf("test BST remove min failed: excute RemoveMin has error %s", err.Error())
			}
			actually = append(actually, min)
		}
		if len(actually) != len(tt.expect) {
			t.Errorf("test BST remove min failed: actually delete element count check failed, actually %d expect %d", len(actually), len(tt.expect))
		}
		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST remove min failed: check remove min return element failed, actully %d expect %d", actually[i], e)
			}
		}
	}
}

func TestBST_RemoveMax(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{5},
			[]int{5},
		},
		{
			[]int{5, 8, 6, 1, 4},
			[]int{8, 6, 5, 4, 1},
		},
		{
			[]int{5, 8, 6, 1, 4, 6},
			[]int{8, 6, 5, 4, 1},
		},
	}

	tests = append(tests, func() test {
		//t.Log("now created lang random data")

		var input []int
		for i := 0; i < 1000; i++ {
			input = append(input, rand.Intn(10000))
		}

		var noRepeat [10000]int
		for _, e := range input {
			noRepeat[e] = 1
		}

		var expect []int
		for i, e := range noRepeat {
			if e == 1 {
				expect = append(expect, i)
			}
		}

		BubbleSort.SortByASC(expect)

		//t.Logf("lang random data:\ninput: %v\nexpect: %v", input, expect)
		return test{
			input:  input,
			expect: expect,
		}
	}())

	//t.Log("now created lang random data")
	//for _, tt := range tests {
	//	fmt.Println(tt.input)
	//	fmt.Println(tt.expect)
	//	fmt.Println()
	//}

	for _, tt := range tests {
		bst := NewBST()
		for _, e := range tt.input {
			bst.Push(e)
		}
		var actually []int
		for !bst.IsEmpty() {
			max, err := bst.RemoveMax()
			if err != nil {
				t.Errorf("test BST remove max failed: excute RemoveMax has error %s", err.Error())
			}
			actually = append(actually, max)
		}
		if len(actually) != len(tt.expect) {
			t.Errorf("test BST remove max failed: actually delete element count check failed, actually %d expect %d", len(actually), len(tt.expect))
		}
		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST remove max failed: check remove max return element failed, actully %d expect %d", actually[i], e)
			}
		}
	}
}

func TestBST_Remove(t *testing.T) {
	type test struct {
		input        []int
		deleteTarget []int
		expect       []int
	}

	tests := []test{
		{
			[]int{5},
			[]int{5},
			[]int{},
		},
		{
			[]int{5, 8, 6, 1, 4},
			[]int{5},
			[]int{1, 4, 6, 8},
		},
		{
			[]int{5, 8, 6, 1, 4, 6},
			[]int{5},
			[]int{1, 4, 6, 8},
		},
	}

	tests = append(tests, func() test {
		//t.Log("now created lang random data")

		var input []int
		for i := 0; i < 1000; i++ {
			input = append(input, rand.Intn(10000))
		}

		var noRepeat [10000]int
		for _, e := range input {
			noRepeat[e] = 1
		}

		var deleteTarget []int

		var expect []int
		for i, e := range noRepeat {
			if e == 1 {
				if rand.Intn(100) > 50 {
					deleteTarget = append(deleteTarget, i)
					continue
				}
				expect = append(expect, i)
			}
		}

		BubbleSort.SortByDESC(expect)

		//t.Logf("lang random data:\ninput: %v\nexpect: %v", input, expect)
		return test{
			input:        input,
			deleteTarget: deleteTarget,
			expect:       expect,
		}
	}())

	//t.Log("now created lang random data")
	//for _, tt := range tests {
	//	fmt.Println(tt.input)
	//	fmt.Println(tt.deleteTarget)
	//	fmt.Println(tt.expect)
	//	fmt.Println()
	//}

	for _, tt := range tests {
		bst := NewBST()
		for _, e := range tt.input {
			bst.Push(e)
		}
		for _, e := range tt.deleteTarget {
			err := bst.Remove(e)
			if err != nil {
				t.Errorf("test BST remove element failed: %s, input %v deleteTarget %v expect %v", err.Error(), tt.input, tt.deleteTarget, tt.expect)
			}
		}

		var actually []int

		for !bst.IsEmpty() {
			min, err := bst.RemoveMin()
			if err != nil {
				t.Errorf("test BST remove failed: BST not is empty, but excute RemoveMax has error %s", err.Error())
			}
			actually = append(actually, min)
		}
		if len(actually) != len(tt.expect) {
			t.Errorf("test BST remove failed: actually delete element count check failed, actually %d expect %d", len(actually), len(tt.expect))
		}
		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST remove failed: check remove return element failed, actully %d expect %d", actually[i], e)
			}
		}
	}
}

func TestBST_GetPreOrder(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{5},
			[]int{5},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{5, 6, 4, 6, 2, 1, 4, 6},
			[]int{5, 4, 2, 1, 6},
		},
		{
			[]int{9, 1, 3, 6, 5, 4, 1},
			[]int{9, 1, 3, 6, 5, 4},
		},
		{
			[]int{4, 5, 2, 9, 4, 7, 2, 6, 4, 1, 2, 3, 6, 5, 4, 2, 2, 2, 2},
			[]int{4, 2, 1, 3, 5, 9, 7, 6},
		},
	}

	for _, tt := range tests {
		bst := NewBST()
		for _, e := range tt.input {
			bst.Push(e)
		}

		actually := bst.GetPreOrder()

		if len(actually) != len(tt.expect) {
			t.Errorf("test BST GetPreOrder failed: response len check failed, actually %d expect %d. input %v expect %v actually %v", len(actually), len(tt.expect), tt.input, tt.expect, actually)
		}

		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST GetPreOrder failed: response element check failed, actually %d expect %d. input %v expect %v actually %v", actually[i], e, tt.input, tt.expect, actually)
			}
		}
	}
}

func TestBST_GetInOrder(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{5},
			[]int{5},
		},
		{
			[]int{1, 2},
			[]int{1, 2},
		},
		{
			[]int{5, 6, 4, 6, 2, 1, 4, 6},
			[]int{1, 2, 4, 5, 6},
		},
		{
			[]int{9, 1, 3, 6, 5, 4, 1},
			[]int{1, 3, 4, 5, 6, 9},
		},
		{
			[]int{4, 5, 2, 9, 4, 7, 2, 6, 4, 1, 2, 3, 6, 5, 4, 2, 2, 2, 2},
			[]int{1, 2, 3, 4, 5, 6, 7, 9},
		},
	}

	for _, tt := range tests {
		bst := NewBST()
		for _, e := range tt.input {
			bst.Push(e)
		}

		actually := bst.GetInOrder()

		if len(actually) != len(tt.expect) {
			t.Errorf("test BST GetInOrder failed: response len check failed, actually %d expect %d. input %v expect %v actually %v", len(actually), len(tt.expect), tt.input, tt.expect, actually)
		}

		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST GetInOrder failed: response element check failed, actually %d expect %d. input %v expect %v actually %v", actually[i], e, tt.input, tt.expect, actually)
			}
		}
	}
}

func TestBST_GetPostOrder(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{5},
			[]int{5},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{5, 6, 4, 6, 2, 1, 4, 6},
			[]int{1, 2, 4, 6, 5},
		},
		{
			[]int{9, 1, 3, 6, 5, 4, 1},
			[]int{4, 5, 6, 3, 1, 9},
		},
		{
			[]int{4, 5, 2, 9, 4, 7, 2, 6, 4, 1, 2, 3, 6, 5, 4, 2, 2, 2, 2},
			[]int{1, 3, 2, 6, 7, 9, 5, 4},
		},
	}

	for _, tt := range tests {
		bst := NewBST()
		for _, e := range tt.input {
			bst.Push(e)
		}

		actually := bst.GetPostOrder()

		if len(actually) != len(tt.expect) {
			t.Errorf("test BST GetPostOrder failed: response len check failed, actually %d expect %d. input %v expect %v actually %v", len(actually), len(tt.expect), tt.input, tt.expect, actually)
		}

		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST GetPostOrder failed: response element check failed, actually %d expect %d. input %v expect %v actually %v", actually[i], e, tt.input, tt.expect, actually)
			}
		}
	}
}
