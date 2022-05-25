package HeapSortByDESC

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/OrderSlice"
	"math/rand"
	"testing"
)

func TestGetSortByDESC(t *testing.T) {
	tests := func() (result []int) {
		for i := 0; i < 1000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()

	maxHeapArr := GetSortByDESC(tests)

	if !OrderSlice.IsDESC(maxHeapArr) {
		t.Errorf("check GetSortByDESC failed: check return slice not is a order by ASC slice")
	}
}

func TestSetSortByDESC(t *testing.T) {
	tests := func() (result []int) {
		for i := 0; i < 1000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()

	SetSortByDESC(tests)

	if !OrderSlice.IsDESC(tests) {
		t.Errorf("check SetSortByDESC failed: check return slice not is a order by ASC slice")
	}
}

func BenchmarkGetSortByDESC(b *testing.B) {
	tests := func() (result []int) {
		for i := 0; i < 100000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sortedArr := GetSortByDESC(tests)

		if !OrderSlice.IsDESC(sortedArr) {
			b.Errorf("check GetSortByDESC failed: check return slice not is a order by ASC slice")
		}
	}
}
