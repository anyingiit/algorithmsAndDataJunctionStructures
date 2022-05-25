package SliceStack

import "testing"

func TestSliceStack_Push(t *testing.T) {
	type test struct {
		input  []int
		expect []int
	}

	tests := []test{
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{7, 6, 5, 4, 3, 2, 1},
		},
		{
			[]int{8, 6, 4, 8, 9, 7},
			[]int{7, 9, 8, 4, 6, 8},
		},
	}

	for _, tt := range tests {
		stack := NewSliceStack()
		for _, e := range tt.input {
			stack.Enqueue(e)
		}

		var actually []int
		for !stack.IsEmpty() {
			actually = append(actually, stack.Dequeue().(int))
		}

		if len(actually) != len(tt.expect) {
			t.Errorf("test SliceStack failed: check actually len failed actually len %d expect len %d, input %v", len(actually), len(tt.expect), tt.input)
		}

		for i, e := range tt.expect {
			if actually[i] != e {
				t.Errorf("test SliceStack failed: pop got element check failed actually %d expect %d, input %v expect %v", actually[i], e, tt.input, tt.expect)
			}
		}
	}
}
