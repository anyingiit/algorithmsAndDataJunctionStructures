package V1

import "fmt"

type Array struct {
	arr  []int
	size int
	cap  int
}

func NewArrayCustomCap(cap int) *Array {
	return &Array{
		arr:  make([]int, cap),
		size: 0,
		cap:  cap,
	}
}

func NewArray() *Array {
	return &Array{
		arr:  make([]int, 10),
		size: 0,
		cap:  10,
	}
}

// resize newSize >= size
func (a *Array) resize(newSize int) error {
	if newSize < a.Size() {
		return fmt.Errorf("cannot resize array, because must newSize >= size")
	}
	newSlice := make([]int, newSize)
	copy(newSlice, a.arr)
	a.arr = newSlice
	a.cap = newSize
	return nil
}

func (a *Array) Size() int {
	return a.size
}

func (a *Array) IsEmpty() bool {
	return a.Size() == 0
}

func (a *Array) Append(e int) {
	if a.Size() == a.cap {
		err := a.resize(a.cap * 2)
		if err != nil {
			panic(fmt.Errorf("internal error: %d", err))
		}
	}
	a.arr[a.Size()] = e
	a.size++
}

func (a *Array) Has(e int) bool {
	for i := 0; i < a.Size(); i++ {
		if a.arr[i] == e {
			return true
		}
	}
	return false
}

func (a *Array) RemoveFirst(e int) (index int, err error) {
	i, err := a.IndexFirst(e)
	if err != nil {
		return 0, err
	}

	// 如果实际存储元素的个数减去一(因为即将要删除一个元素), 小于目前容量的四分之一,则认为需要缩小空间了,
	// 执行resize, 重置容量到当前容量的二分之一
	if a.Size()-1 < a.cap/4 {
		err := a.resize(a.cap / 2)
		if err != nil {
			panic(fmt.Errorf("internal error: %d", err))
		}
	}

	if i == a.Size()-1 {
		a.size--
		return i, nil
	}

	for j := i; j < a.Size()-1; j++ {
		a.arr[j] = a.arr[j+1]
	}
	a.size--
	return i, nil
}

func (a *Array) Get(index int) (e int, err error) {
	if index < 0 || index > a.Size()-1 {
		return 0, fmt.Errorf("get index %d failed: plaese make sure index >= 0 and index < size", index)
	}
	return a.arr[index], nil
}

func (a *Array) IndexFirst(e int) (index int, err error) {
	for i := 0; i < a.Size(); i++ {
		if a.arr[i] == e {
			return i, nil
		}
	}
	return 0, fmt.Errorf("remove element %d failed: cannot find this element", e)
}

func (a *Array) Set(index int, newE int) (err error) {
	if index < 0 || index > a.Size()-1 {
		return fmt.Errorf("get index %d failed: plaese make sure index >= 0 and index < size", index)
	}
	a.arr[index] = newE
	return nil
}

func (a *Array) Swap(i, j int) error {
	if i < 0 || i > a.Size()-1 || j < 0 || j > a.Size()-1 {
		return fmt.Errorf("swap index i=%d, j=%d failed: plaese make sure i >= 0 and i < size and j >= 0 and j < size", i, j)
	}
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
	return nil
}
