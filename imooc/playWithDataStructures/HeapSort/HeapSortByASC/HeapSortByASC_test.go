package HeapSortByASC

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/OrderSlice"
	"math/rand"
	"testing"
)

func TestGetSortByASC(t *testing.T) {
	tests := func() (result []int) {
		for i := 0; i < 1000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()

	sortedArr := GetSortByASC(tests)

	if !OrderSlice.IsASC(sortedArr) {
		t.Errorf("check GetSortByASC failed: check return slice not is a order by ASC slice")
	}
}

func TestSetSortByASC(t *testing.T) {
	tests := func() (result []int) {
		for i := 0; i < 1000; i++ {
			result = append(result, rand.Intn(700))
		}
		return result
	}()

	SetSortByASC(tests)

	if !OrderSlice.IsASC(tests) {
		t.Errorf("check SetSortByASC failed: check return slice not is a order by ASC slice")
	}
}
