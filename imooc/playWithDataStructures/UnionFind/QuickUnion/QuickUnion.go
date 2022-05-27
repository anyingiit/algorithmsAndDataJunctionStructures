package QuickUnion

import "fmt"

// UnionFind for quick union
type UnionFind struct {
	parent []int
}

func NewUnionFind(size int) *UnionFind {
	list := make([]int, size)
	for i := range list {
		list[i] = i
	}
	return &UnionFind{list}
}

func (u *UnionFind) GetSize() int {
	return len(u.parent)
}

func (u *UnionFind) IsConnected(p, q int) (bool, error) {
	pRoot, err := u.find(p)
	if err != nil {
		return false, fmt.Errorf("cannot check p, q is connected, because find p failed: %s", err.Error())
	}
	qRoot, err := u.find(q)
	if err != nil {
		return false, fmt.Errorf("cannot check p, q is connected, because find q failed: %s", err.Error())
	}
	return pRoot == qRoot, nil
}

func (u *UnionFind) UnionElements(p, q int) error {
	pRoot, err := u.find(p)
	if err != nil {
		return fmt.Errorf("cannot union element p, q is connected, because find p failed: %s", err.Error())
	}
	qRoot, err := u.find(q)
	if err != nil {
		return fmt.Errorf("cannot union element p, q is connected, because find q failed: %s", err.Error())
	}
	u.parent[pRoot] = qRoot
	return nil
}

func (u *UnionFind) find(p int) (int, error) {
	if p < 0 || p > u.GetSize()-1 {
		return 0, fmt.Errorf("p out of bound")
	}

	var f func(q int) int
	f = func(q int) int {
		if u.parent[q] == q {
			return q
		}
		return f(u.parent[q])
	}

	return f(p), nil
}
