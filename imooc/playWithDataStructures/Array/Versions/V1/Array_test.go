package V1

import (
	"testing"
)

func TestArray_Append(t *testing.T) {
	type test struct {
		input []int
	}
	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		for _, e := range tt.input {
			if !array.Has(e) {
				t.Errorf("test failed, expect has %d but not has", e)
			}
		}
	}
}

func TestArray_RemoveFirst(t *testing.T) {
	type test struct {
		input             []int
		removeTarget      []int
		expectRemoveIndex []int
		expect            []int
	}
	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 1, 2, 5},
			[]int{4, 1, 5, 23},
			[]int{0, 0, 5, 1},
			[]int{6, 1, 1, 2},
		},
	}

	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		for i, e := range tt.removeTarget {
			first, err := array.RemoveFirst(e)
			if err != nil {
				t.Errorf("test Set failed: excute RemoveFirst has err %s. remove rarget %d", err.Error(), e)
			}
			if first != tt.expectRemoveIndex[i] {
				t.Errorf("test Set failed: check return first remove value index failed expect %d actually %d", tt.expectRemoveIndex[i], first)
			}
		}

		for i, e := range tt.expect {
			get, err := array.Get(i)
			if err != nil {
				t.Errorf("test Set failed: excute Get has err %s. remove index %d", err.Error(), i)
			}

			if get != e {
				t.Errorf("test Set failed: expect %d actually %d", e, get)
			}
		}
	}
}

func TestArray_Set(t *testing.T) {
	type test struct {
		input   []int
		indexTo map[int]int
		expect  []int
	}
	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 2, 5},
			map[int]int{
				0: 9,
				6: 2,
				3: 8,
			},
			[]int{9, 1, 6, 8, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		for index, newValue := range tt.indexTo {
			err := array.Set(index, newValue)
			if err != nil {
				t.Errorf("test Set failed: excute Set has err %s. index %d", err.Error(), index)
			}
		}

		for i, e := range tt.expect {
			get, err := array.Get(i)
			if err != nil {
				t.Errorf("test Set failed: excute Get has err %s. index %d", err.Error(), i)
			}
			if get != e {
				t.Errorf("test Set failed: expect %d actually %d", e, get)
			}
		}

	}
}

func TestArray_Has(t *testing.T) {
	type test struct {
		input []int
	}

	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		for _, e := range tt.input {
			if !array.Has(e) {
				t.Errorf("test Has failed: expect has element %d but not have", e)
			}
		}

	}
}

func TestArray_Get(t *testing.T) {
	type test struct {
		input []int
	}

	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		for i, e := range tt.input {
			get, err := array.Get(i)
			if err != nil {
				t.Errorf("test Get failed: excute Get has err %s", err.Error())
			}
			if get != e {
				t.Errorf("test Get failed: expect %d actully %d", e, get)
			}
		}

	}

}

func TestArray_Index(t *testing.T) {
	type test struct {
		input       []int
		expectIndex []int
	}

	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 2, 5},
			[]int{0, 1, 2, 3, 1, 5, 6},
		},
	}

	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		for i, e := range tt.expectIndex {
			indexFirst, err := array.IndexFirst(tt.input[i])
			if err != nil {
				t.Errorf("test IndexFirst failed: excute IndexFirst has err: %s", err.Error())
			}
			if indexFirst != e {
				t.Errorf("test IndexFirst failed: actually %d expectIndex %d", indexFirst, i)
			}
		}
	}
}

func TestArray_Size(t *testing.T) {
	type test struct {
		input []int
	}
	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		if len(tt.input) != array.Size() {
			t.Errorf("test size failed: expect size %d actually %d", len(tt.input), array.Size())
		}
	}

}

func TestArray_IsEmpty(t *testing.T) {
	type test struct {
		input []int
	}
	tests := []test{
		{
			[]int{4, 1, 6, 23, 1, 2, 5},
		},
	}

	for _, tt := range tests {
		array := NewArray()
		for _, e := range tt.input {
			array.Append(e)
		}

		for _, e := range tt.input {
			_, err := array.RemoveFirst(e)
			if err != nil {
				t.Errorf("test IsEmpty failed: excute RemoveFist has err %s", err.Error())
			}
		}

		if !array.IsEmpty() {
			t.Errorf("test IsEmpty failed: expect array is empty actually size is %d", array.Size())
		}
	}
}

//TODO: 测试resize是否能正确的工作
