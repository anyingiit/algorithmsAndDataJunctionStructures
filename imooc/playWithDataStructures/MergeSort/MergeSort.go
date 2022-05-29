package MergeSort

type mergeFunc func(x []int, l, mid, r int)

func sort(x []int, mergeFunc mergeFunc) {
	mergeSort(x, 0, len(x)-1, mergeFunc)
}

func mergeSort(x []int, l, r int, mergeFunc mergeFunc) {
	if l >= r {
		return
	}
	mid := (r-l)/2 + l // TIPS: 当处理上亿级别规模的数据的时候如果使用 (l + r) / 2 获取mid值, l + r的过程可能会产生溢出, 如果想规避这个问题可以将刚刚的算式优化为 l + (r - l) / 2, 这样可以避免溢出
	mergeSort(x, l, mid, mergeFunc)
	mergeSort(x, mid+1, r, mergeFunc)
	mergeFunc(x, l, mid, r)
}

func SortByDESC(x []int) {
	mergeFunc := mergeFunc(func(x []int, l, mid, r int) {
		xCopy := make([]int, r-l+1)
		copy(xCopy, x[l:r+1])

		xCopyMid := mid - l
		xCopyL := 0
		xCopyR := xCopyMid + 1
		k := l
		for {
			if xCopyL > xCopyMid && xCopyR > len(xCopy)-1 {
				break
			}

			if xCopyR > len(xCopy)-1 || (xCopyL <= xCopyMid && xCopy[xCopyL] < xCopy[xCopyR]) { // 如果R越界了或者我自身存在并且比R对应的元素小
				x[k] = xCopy[xCopyL]
				xCopyL++
			} else { // R一定没有越界, 并且L越界了或者R比L对应的元素小
				x[k] = xCopy[xCopyR]
				xCopyR++
			}
			k++
		}
	})

	sort(x, mergeFunc)
}

func SortByASC(x []int) {
	mergeFunc := mergeFunc(func(x []int, l, mid, r int) {
		xCopy := make([]int, r-l+1)
		copy(xCopy, x[l:r+1])

		xCopyMid := mid - l
		xCopyL := 0
		xCopyR := xCopyMid + 1
		k := l
		for {
			if xCopyL > xCopyMid && xCopyR > len(xCopy)-1 {
				break
			}

			// 真值推算使用到了德摩根律
			if xCopyR > len(xCopy)-1 || (xCopyL <= xCopyMid && xCopy[xCopyL] > xCopy[xCopyR]) { // 如果R越界了或者我自身存在并且比R对应的元素大
				x[k] = xCopy[xCopyL]
				xCopyL++
			} else { // R一定没有越界, 并且L越界了或者R比L对应的元素大
				x[k] = xCopy[xCopyR]
				xCopyR++
			}
			k++
		}
	})

	sort(x, mergeFunc)
}
