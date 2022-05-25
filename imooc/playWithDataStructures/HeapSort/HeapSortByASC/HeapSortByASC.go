package HeapSortByASC

func SetSortByASC(x []int) {
	toMinHeap(x)
	maxHeapSize := len(x)
	for maxHeapSize > 0 {
		x[0], x[maxHeapSize-1] = x[maxHeapSize-1], x[0]
		maxHeapSize--
		minHeapSiftDown(x, 0, maxHeapSize)
	}
}
func GetSortByASC(src []int) []int {
	srcCopy := make([]int, len(src))
	copy(srcCopy, src)

	toMinHeap(srcCopy)
	maxHeapSize := len(srcCopy)
	for maxHeapSize > 0 {
		srcCopy[0], srcCopy[maxHeapSize-1] = srcCopy[maxHeapSize-1], srcCopy[0]
		maxHeapSize--
		minHeapSiftDown(srcCopy, 0, maxHeapSize)
	}
	return srcCopy
}

func minHeapSiftDown(x []int, index int, xLen int) {
	inspect := index
	for (inspect*2 + 1) < xLen {
		minIndex := inspect*2 + 1
		if (inspect*2+2) < xLen && x[inspect*2+2] < x[minIndex] {
			minIndex = inspect*2 + 2
		}

		if x[minIndex] < x[inspect] {
			x[minIndex], x[inspect] = x[inspect], x[minIndex]
			inspect = minIndex
		} else {
			break
		}
	}
}

func toMinHeap(x []int) {
	for i := len(x) - 1; i >= 0; i-- {
		minHeapSiftDown(x, i, len(x))
	}
}
