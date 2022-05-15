package V5

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/versions/V5/E/Eint"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BubbleSort"
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/OrderSlice"
	"fmt"
	"math/rand"
	"testing"
)

func TestBST_Size(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(500))
	}
	noRepeat, err := OrderSlice.GetNoRepeat(BubbleSort.GetSortByDESC(tests))
	if err != nil {
		t.Errorf("check size failed: make no repeat elements failed %s", err.Error())
	}

	noRepeatCount := len(noRepeat)
	//t.Logf("tests len %d, no repeat num count %d\n", len(tests), noRepeatCount)

	bst := NewBST()
	for _, e := range tests {
		bst.Push(EInt.NewE(e))
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
		bst.Push(EInt.NewE(key))
	}

	for key := range noRepeat {
		if !bst.Has(EInt.NewE(key)) {
			t.Errorf("BST_V5 Has test failed, expect has %d but not have", key)
		}
	}

}

func TestBST_Push(t *testing.T) {
	type test struct {
		input              []int
		expectPushResponse []bool
		expectOutput       []int
	}

	tests := []test{
		{
			[]int{5},
			[]bool{true},
			[]int{5},
		},
		{
			[]int{1, 2},
			[]bool{true, true},
			[]int{1, 2},
		},
		{
			[]int{5, 8, 6, 1, 4},
			[]bool{true, true, true, true, true},
			[]int{1, 4, 5, 6, 8},
		},
		{
			[]int{5, 8, 6, 1, 4, 6},
			[]bool{true, true, true, true, true, false},
			[]int{1, 4, 5, 6, 8},
		},
		{
			[]int{5, 6, 4, 6, 2, 1, 4, 6},
			[]bool{true, true, true, false, true, true, false, false},
			[]int{1, 2, 4, 5, 6},
		},
		{
			[]int{9, 1, 3, 6, 5, 4, 1},
			[]bool{true, true, true, true, true, true, false},
			[]int{1, 3, 4, 5, 6, 9},
		},
		{
			[]int{4, 5, 2, 9, 4, 7, 2, 6, 4, 1, 2, 3, 6, 5, 4, 2, 2, 2, 2},
			[]bool{true, true, true, true, false, true, false, true, false, true, false, true, false, false, false, false, false, false, false},
			[]int{1, 2, 3, 4, 5, 6, 7, 9},
		},
	}

	for _, tt := range tests {
		bst := NewBST()

		var actuallyResponse []bool
		for _, e := range tt.input {
			actuallyResponse = append(actuallyResponse, bst.Push(EInt.NewE(e)))
		}

		if len(actuallyResponse) != len(tt.expectPushResponse) {
			t.Errorf("test BST_V5 Push failed: push response len check failed, actually %v expect %v. input %v expectPushResponse %v actuallyResponse %v", len(actuallyResponse), len(tt.expectPushResponse), tt.input, tt.expectPushResponse, actuallyResponse)
		}

		for i, e := range tt.expectPushResponse {
			if actuallyResponse[i] != e {
				t.Errorf("test BST_V5 Push failed: push response check failed, actually %v expect %v. input %v expectPushResponse %v actuallyResponse %v", actuallyResponse[i], e, tt.input, tt.expectPushResponse, actuallyResponse)
			}
		}

		var actuallyOutput []int

		for !bst.IsEmpty() {
			min, err := bst.RemoveMin()
			if err != nil {
				t.Errorf("test BST_V5 Push failed: RemoveMin failed %s, input %v", err.Error(), tt.input)
			}
			actuallyOutput = append(actuallyOutput, min.Data().(int))
		}

		if len(actuallyOutput) != len(tt.expectOutput) {
			t.Errorf("test BST_V5 Push failed: check expect output len failed, actually %d expect %d. input %v expectOutput %v actuallyOutput %v", len(actuallyOutput), len(tt.expectOutput), tt.input, tt.expectOutput, actuallyOutput)
		}

		for i, e := range tt.expectOutput {
			if actuallyOutput[i] != e {
				t.Errorf("test BST_V5 Push failed: check expect output element failed, actually %d expect %d. input %v expectOutput %v actuallyOutput %v", actuallyOutput[i], e, tt.input, tt.expectOutput, actuallyOutput)
			}
		}
	}
}

func TestBST_String(t *testing.T) {
	bst := NewBST()

	bst.Push(EInt.NewE(41))
	bst.Push(EInt.NewE(22))
	bst.Push(EInt.NewE(58))
	bst.Push(EInt.NewE(15))
	bst.Push(EInt.NewE(33))
	bst.Push(EInt.NewE(13))
	bst.Push(EInt.NewE(50))
	bst.Push(EInt.NewE(42))
	bst.Push(EInt.NewE(53))

	expect := "41\n--22\n----15\n------13\n--------null\n--------null\n------null\n----33\n------null\n------null\n--58\n----50\n------42\n--------null\n--------null\n------53\n--------null\n--------null\n----null\n"
	if fmt.Sprint(bst) != expect {
		t.Errorf("check BST_V5 String failed:\nexpect\n\n\"%s\"\n\nactually\n\"%s\"\n\n", expect, fmt.Sprint(bst))
	}

}

func TestBST_RemoveMin(t *testing.T) {
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

		BubbleSort.SetSortByDESC(expect)

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
			bst.Push(EInt.NewE(e))
		}
		var actually []int
		for !bst.IsEmpty() {
			min, err := bst.RemoveMin()
			if err != nil {
				t.Errorf("test BST_V5 remove min failed: excute RemoveMin has error %s", err.Error())
			}
			actually = append(actually, min.Data().(int))
		}
		if len(actually) != len(tt.expect) {
			t.Errorf("test BST_V5 remove min failed: actually delete element count check failed, actually %d expect %d", len(actually), len(tt.expect))
		}
		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST_V5 remove min failed: check remove min return element failed, actully %d expect %d", actually[i], e)
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

		BubbleSort.SetSortByASC(expect)

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
			bst.Push(EInt.NewE(e))
		}
		var actually []int
		for !bst.IsEmpty() {
			max, err := bst.RemoveMax()
			if err != nil {
				t.Errorf("test BST_V5 remove max failed: excute RemoveMax has error %s", err.Error())
			}
			actually = append(actually, max.Data().(int))
		}
		if len(actually) != len(tt.expect) {
			t.Errorf("test BST_V5 remove max failed: actually delete element count check failed, actually %d expect %d", len(actually), len(tt.expect))
		}
		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST_V5 remove max failed: check remove max return element failed, actully %d expect %d", actually[i], e)
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

		BubbleSort.SetSortByDESC(expect)

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
			bst.Push(EInt.NewE(e))
		}
		for _, e := range tt.deleteTarget {
			err := bst.Remove(EInt.NewE(e))
			if err != nil {
				t.Errorf("test BST_V5 remove element failed: %s, input %v deleteTarget %v expect %v", err.Error(), tt.input, tt.deleteTarget, tt.expect)
			}
		}

		var actually []int

		for !bst.IsEmpty() {
			min, err := bst.RemoveMin()
			if err != nil {
				t.Errorf("test BST_V5 remove failed: BST_V5 not is empty, but excute RemoveMax has error %s", err.Error())
			}
			actually = append(actually, min.Data().(int))
		}
		if len(actually) != len(tt.expect) {
			t.Errorf("test BST_V5 remove failed: actually delete element count check failed, actually %d expect %d", len(actually), len(tt.expect))
		}
		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test BST_V5 remove failed: check remove return element failed, actully %d expect %d", actually[i], e)
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
			bst.Push(EInt.NewE(e))
		}

		actually := bst.GetPreOrder()

		if len(actually) != len(tt.expect) {
			t.Errorf("test BST_V5 GetPreOrder failed: response len check failed, actually %d expect %d. input %v expect %v actually %v", len(actually), len(tt.expect), tt.input, tt.expect, actually)
		}

		for i, e := range tt.expect {
			if actually[i].Data().(int) != e {
				t.Errorf("test BST_V5 GetPreOrder failed: response element check failed, actually %d expect %d. input %v expect %v actually %v", actually[i], e, tt.input, tt.expect, actually)
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
			bst.Push(EInt.NewE(e))
		}

		actually := bst.GetInOrder()

		if len(actually) != len(tt.expect) {
			t.Errorf("test BST_V5 GetInOrder failed: response len check failed, actually %d expect %d. input %v expect %v actually %v", len(actually), len(tt.expect), tt.input, tt.expect, actually)
		}

		for i, e := range tt.expect {
			if actually[i].Data().(int) != e {
				t.Errorf("test BST_V5 GetInOrder failed: response element check failed, actually %d expect %d. input %v expect %v actually %v", actually[i], e, tt.input, tt.expect, actually)
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
			bst.Push(EInt.NewE(e))
		}

		actually := bst.GetPostOrder()

		if len(actually) != len(tt.expect) {
			t.Errorf("test BST_V5 GetPostOrder failed: response len check failed, actually %d expect %d. input %v expect %v actually %v", len(actually), len(tt.expect), tt.input, tt.expect, actually)
		}

		for i, e := range tt.expect {
			if actually[i].Data().(int) != e {
				t.Errorf("test BST_V5 GetPostOrder failed: response element check failed, actually %d expect %d. input %v expect %v actually %v", actually[i], e, tt.input, tt.expect, actually)
			}
		}
	}
}
