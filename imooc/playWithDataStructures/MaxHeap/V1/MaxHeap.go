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
	if m.IsEmpty() {
		return 0, fmt.Errorf("cannot extract max, beacuse the Max Heap is empty")
	}
	max := m.data[0]
	m.data[0] = m.data[m.Size()-1]
	m.data = m.data[:m.Size()-1]
	m.siftDown(0)
	return max, nil
}

// siftDown 使用非递归下浮(siftDown)
func (m *MaxHeap) siftDown(index int) {
	inspect := index
	for {
		// 坐标为叶子节点时遍历时结束
		if !m.hasIndex(m.leftChild(inspect)) && !m.hasIndex(m.rightChild(inspect)) {
			break
		}

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
