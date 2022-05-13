package OrderSlice

import (
	"fmt"
)

func GetNoRepeat(orderedSlice []int) ([]int, error) {
	if !IsOrder(orderedSlice) {
		return nil, fmt.Errorf("can not get no repeat slice: input silice not is a ordered slice")
	}
	return GetNoRepeatNoCheck(orderedSlice), nil
}

func GetNoRepeatNoCheck(orderedSlice []int) []int {
	if len(orderedSlice) == 0 || len(orderedSlice) == 1 {
		return orderedSlice
	}

	cp := make([]int, len(orderedSlice))
	copy(cp, orderedSlice)

	i, j := 0, 0

	for j < len(cp) {
		if cp[i] != cp[j] {
			i++
			cp[i] = cp[j]
		}
		j++
	}
	return cp[:i+1]
}
