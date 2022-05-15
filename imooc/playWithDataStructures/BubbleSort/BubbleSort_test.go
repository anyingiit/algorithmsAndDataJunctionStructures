package BubbleSort

import (
	"math/rand"
	"testing"
)

// TestSortByDESC
//
// 1. 排序后长度测试
// 2. 排序后顺序测试
func TestSetSortByDESC(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}

	originLen := len(tests)

	SetSortByDESC(tests)

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
func TestSetSortByASC(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}

	originLen := len(tests)

	SetSortByASC(tests)

	if len(tests) != originLen {
		t.Errorf("len test failed expect %d, actually %d", originLen, len(tests))
	}

	for i := 0; i < len(tests)-1; i++ {
		if tests[i] < tests[i+1] {
			t.Errorf("test failed")
		}
	}

}

func TestGetSortByDESC(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}

	testsCopy := make([]int, len(tests))
	copy(testsCopy, tests)

	descResult := GetSortByDESC(tests)

	if len(tests) != len(testsCopy) {
		t.Errorf("check GetSortByDESC is changed origin slice failed, originLen %d actuallyLen %d. origin %v actully %v", len(testsCopy), len(tests), testsCopy, tests)
	}
	for i, e := range testsCopy {
		if tests[i] != e {
			t.Errorf("check GetSortByDESC is changed origin slice failed, actually %v expect %v. tests %v testsCopy %v", tests[i], e, tests, testsCopy)
		}
	}

	if len(descResult) != len(tests) {
		t.Errorf("len test failed expect %d, actually %d", len(tests), len(descResult))
	}
	for i := 0; i < len(descResult)-1; i++ {
		if descResult[i] > descResult[i+1] {
			t.Errorf("test failed")
		}
	}
}

func TestGetSortByASC(t *testing.T) {
	var tests []int
	for i := 0; i < 1000; i++ {
		tests = append(tests, rand.Intn(10000))
	}

	testsCopy := make([]int, len(tests))
	copy(testsCopy, tests)

	ascResult := GetSortByASC(tests)

	if len(tests) != len(testsCopy) {
		t.Errorf("check GetSortByDESC is changed origin slice failed, originLen %d actuallyLen %d. origin %v actully %v", len(testsCopy), len(tests), testsCopy, tests)
	}
	for i, e := range testsCopy {
		if tests[i] != e {
			t.Errorf("check GetSortByDESC is changed origin slice failed, actually %v expect %v. tests %v testsCopy %v", tests[i], e, tests, testsCopy)
		}
	}

	if len(ascResult) != len(tests) {
		t.Errorf("len test failed expect %d, actually %d", len(tests), len(ascResult))
	}
	for i := 0; i < len(ascResult)-1; i++ {
		if ascResult[i] < ascResult[i+1] {
			t.Errorf("test failed")
		}
	}
}
