package OrderSlice

import (
	"fmt"
)

// GetNoRepeat 获取有序数组的不重复元素, 如果输入参数不具有顺序则返回错误
func GetNoRepeat(orderedSlice []int) ([]int, error) {
	if !IsOrder(orderedSlice) {
		return nil, fmt.Errorf("can not get no repeat slice: input silice not is a ordered slice")
	}
	return GetNoRepeatNoCheck(orderedSlice), nil
}

// GetNoRepeatNoCheck 获取有序数组的不重复元素, 但是不检查参数是否具有顺序
//
// 请务必确保输入参数具有顺序, 否则将获得错误的返回
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
