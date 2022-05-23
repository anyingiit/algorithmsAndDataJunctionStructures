package V1

import (
	"fmt"
)

// MaxHeap 数组实现的最大堆
type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{nil}
}

// NewMaxHeapForArr Heapify
func NewMaxHeapForArr(e []int) *MaxHeap {
	maxHeap := NewMaxHeap()
	maxHeap.data = make([]int, len(e))
	copy(maxHeap.data, e)
	// Heapify 将乱序的数组转换为一个二叉堆
	// * 也可以创建一个新的maxHeap然后使用Add向其添加元素, 但是那样没有这样快
	// 为什么不从最后一个元素使用siftUp直接遍历完成呢?
	//		因为siftUp会假设当前元素上边的元素顺序是正确的
	//		而现在的情况是 顺序是错误的
	// 为什么不从第一个元素使用siftDown直到遍历完成?
	//		因为siftDown会假设当前元素下面的元素是正确的的
	//		而现在的实际情况是 顺序是错误的
	// 根据以上两点, 正序siftDown不行, 倒序siftUp不行, 倒序使用siftDown行, 那么正序使用siftUp行不行呢?
	//		经过实际验证, 是可行的, 但是应该效率没倒序siftDown高
	//cpu: Intel(R) Core(TM) i7-4790K CPU @ 4.00GHz
	//BenchmarkNewMaxHeapForArr
	//BenchmarkNewMaxHeapForArr-8        98644             13001 ns/op
	//PASS
	for i := maxHeap.Size() - 1; i >= 0; i-- {
		maxHeap.siftDown(i)
	}
	//cpu: Intel(R) Core(TM) i7-4790K CPU @ 4.00GHz
	//BenchmarkNewMaxHeapForArr
	//BenchmarkNewMaxHeapForArr-8        80864             14386 ns/op
	//PASS
	//for i := 0; i < maxHeap.Size()-1; i++ {
	//	maxHeap.siftUp(i)
	//}
	return maxHeap
}

func (m *MaxHeap) Size() int {
	return len(m.data)
}

func (m *MaxHeap) IsEmpty() bool {
	return m.Size() == 0
}

func (m *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

func (m *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

// Add 向最大堆中添加元素
func (m *MaxHeap) Add(e int) {
	m.data = append(m.data, e)
	m.siftUp(m.Size() - 1)
}

// siftUp 使用非递归实现上浮(siftUp)
func (m *MaxHeap) siftUp(index int) {
	inspect := index
	for {
		if inspect <= 0 || m.data[inspect] <= m.data[m.parent(inspect)] {
			break
		}
		m.data[inspect], m.data[m.parent(inspect)] = m.data[m.parent(inspect)], m.data[inspect]
		inspect = m.parent(inspect)
	}
}

// ExtractMax 提取最大元素
func (m *MaxHeap) ExtractMax() (int, error) {
	maxTemp, err := m.FindMax()
	if err != nil {
		return 0, fmt.Errorf("cannot extract max: %s", err.Error())
	}

	m.data[0] = m.data[m.Size()-1]
	m.data = m.data[:m.Size()-1]
	m.siftDown(0)
	return maxTemp, nil
}

// siftDown 使用非递归下浮(siftDown)
func (m *MaxHeap) siftDown(index int) {
	inspect := index
	for m.hasIndex(m.leftChild(inspect)) { // 如果存在左孩子就说明两件事: 1. 左孩子存在 2. 还没到底, inspect还不是叶子结点
		maxIndex := m.leftChild(inspect)
		if m.hasIndex(m.rightChild(inspect)) && m.data[m.rightChild(inspect)] > m.data[maxIndex] {
			maxIndex = m.rightChild(inspect)
		}

		if m.data[maxIndex] > m.data[inspect] {
			m.data[inspect], m.data[maxIndex] = m.data[maxIndex], m.data[inspect]
			inspect = maxIndex
		} else { // 当左右孩子中最大的元素都比inspect的值要小或等于时则遍历结束
			break
		}
	}
}

func (m *MaxHeap) hasIndex(index int) bool {
	if index < 0 || index > m.Size()-1 {
		return false
	}
	return true
}

func (m *MaxHeap) FindMax() (int, error) {
	if m.IsEmpty() {
		return 0, fmt.Errorf("cannot find max element: the MaxHeap is empty")
	}
	return m.data[0], nil
}

// Replace 提取最大元素并添加一个新元素
func (m *MaxHeap) Replace(e int) (int, error) {
	maxTemp, err := m.FindMax()
	if err != nil {
		return 0, fmt.Errorf("cannot Replace: %s", err.Error())
	}
	m.data[0] = e
	m.siftDown(0)
	return maxTemp, nil
}
