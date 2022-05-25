package HeapSortByDESC

func SetSortByDESC(x []int) {
	toMaxHeap(x)
	maxHeapSize := len(x)
	for maxHeapSize > 0 {
		x[0], x[maxHeapSize-1] = x[maxHeapSize-1], x[0]
		maxHeapSize--
		maxHeapSiftDown(x, 0, maxHeapSize)
	}
}
func GetSortByDESC(src []int) []int {
	srcCopy := make([]int, len(src))
	copy(srcCopy, src)

	toMaxHeap(srcCopy)
	maxHeapSize := len(srcCopy)
	for maxHeapSize > 0 {
		srcCopy[0], srcCopy[maxHeapSize-1] = srcCopy[maxHeapSize-1], srcCopy[0]
		maxHeapSize--
		maxHeapSiftDown(srcCopy, 0, maxHeapSize)
	}
	return srcCopy
}

func maxHeapSiftDown(x []int, index int, xLen int) {
	inspect := index
	for (inspect*2 + 1) < xLen {
		maxIndex := inspect*2 + 1
		if (inspect*2+2) < xLen && x[inspect*2+2] > x[maxIndex] {
			maxIndex = inspect*2 + 2
		}

		if x[maxIndex] > x[inspect] {
			x[maxIndex], x[inspect] = x[inspect], x[maxIndex]
			inspect = maxIndex
		} else {
			break
		}
	}
}

func toMaxHeap(x []int) {
	for i := len(x) - 1; i >= 0; i-- {
		maxHeapSiftDown(x, i, len(x))
	}
}
