package MergeSort

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/OrderSlice"
	"math/rand"
	"testing"
)

func TestSortByDESC(t *testing.T) {
	type test struct {
		input []int
	}

	tests := []test{
		{
			[]int{5, 2, 5, 6, 8, 1, 2, 0},
		},
	}

	tests = append(tests, func() test {
		var result []int
		for i := 0; i < 1000000; i++ {
			result = append(result, rand.Intn(1000000))
		}
		return test{result}
	}())

	for _, tt := range tests {
		SortByDESC(tt.input)

		if !OrderSlice.IsDESC(tt.input) {
			t.Errorf("test SortByDESC failed: check is order by DESC failed")
		}
	}
}

func TestSortByASC(t *testing.T) {
	type test struct {
		input []int
	}

	tests := []test{
		{
			[]int{5, 2, 5, 6, 8, 1, 2, 0},
		},
	}

	tests = append(tests, func() test {
		var result []int
		for i := 0; i < 1000000; i++ {
			result = append(result, rand.Intn(1000000))
		}
		return test{result}
	}())

	for _, tt := range tests {
		SortByASC(tt.input)

		if !OrderSlice.IsASC(tt.input) {
			t.Errorf("test SortByDESC failed: check is order by DESC failed")
		}
	}
}
