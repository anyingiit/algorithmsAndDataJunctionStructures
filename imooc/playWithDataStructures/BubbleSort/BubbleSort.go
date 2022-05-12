package BubbleSort

// SortByDESC 降序排序
func SortByDESC(x []int) {
	// 每次一次循环都会把除去已排序元素以外的其他元素中的最大元素, "顶"到已排序元素以外的其他元素中的最后一个
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
}

// SortByASC 増序排序
func SortByASC(x []int) {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] < x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
}
