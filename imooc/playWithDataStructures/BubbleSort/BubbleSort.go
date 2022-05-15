package BubbleSort

// SetSortByDESC input int slice x, set slice x ordered by DESC.
//
// warning: this func well change origin slice x.
//
// DESC: 降序排序
func SetSortByDESC(x []int) {
	// 每次一次循环都会把除去已排序元素以外的其他元素中的最大元素, "顶"到已排序元素以外的其他元素中的最后一个
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
}

// SetSortByASC input int slice x, set slice x ordered by ASC.
//
// warning: this func well change origin slice x.
//
// ASC: 升序排序
func SetSortByASC(x []int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] < x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
}

// GetSortByDESC input int slice x, return ordered by DESC slice result.
//
// tips: this func not well change origin slice x.
//
// DESC: 降序排序
func GetSortByDESC(x []int) (result []int) {
	result = make([]int, len(x))
	copy(result, x)

	SetSortByDESC(result)
	return result
}

// GetSortByASC input int slice x, return ordered by ASC slice result.
//
// tips: this func not well change origin slice x.
//
// ASC: 升序排序
func GetSortByASC(x []int) (result []int) {
	result = make([]int, len(x))
	copy(result, x)

	SetSortByASC(result)
	return result
}
