package BubbleSort

import (
	"math/rand"
	"testing"
)

// TestSortByDESC
//
// 1. 排序后长度测试
// 2. 排序后顺序测试
func TestSortByDESC(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}

	originLen := len(tests)

	SortByDESC(tests)

	if len(tests) != originLen {
		t.Errorf("len test failed expect %d, actually %d", originLen, len(tests))
	}

	for i := 0; i < len(tests)-1; i++ {
		if tests[i] > tests[i+1] {
			t.Errorf("test failed")
		}
	}
}

// TestSortByASC
//
// 1. 排序后长度测试
// 2. 排序后顺序测试
func TestSortByASC(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}

	originLen := len(tests)

	SortByASC(tests)

	if len(tests) != originLen {
		t.Errorf("len test failed expect %d, actually %d", originLen, len(tests))
	}

	for i := 0; i < len(tests)-1; i++ {
		if tests[i] < tests[i+1] {
			t.Errorf("test failed")
		}
	}

}
