package HeapSortByDESC

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/OrderSlice"
	"math/rand"
	"testing"
)

func TestGetSort(t *testing.T) {
	tests := func() (result []int) {
		for i := 0; i < 1000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()

	maxHeapArr := GetSort(tests)

	if !OrderSlice.IsDESC(maxHeapArr) {
		t.Errorf("check GetHeapSort failed: check return slice not is a order by ASC slice")
	}
}

func TestSetSort(t *testing.T) {
	tests := func() (result []int) {
		for i := 0; i < 1000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()

	SetSort(tests)

	if !OrderSlice.IsDESC(tests) {
		t.Errorf("check GetHeapSort failed: check return slice not is a order by ASC slice")
	}
}

func BenchmarkGetSort(b *testing.B) {
	tests := func() (result []int) {
		for i := 0; i < 100000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sortedArr := GetSort(tests)

		if !OrderSlice.IsDESC(sortedArr) {
			b.Errorf("check GetHeapSort failed: check return slice not is a order by ASC slice")
		}
	}
}
