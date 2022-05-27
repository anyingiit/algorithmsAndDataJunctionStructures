package QuickUnionV2

import "fmt"

// UnionFind for quick union 基于size的优化
type UnionFind struct {
	parent []int
	sz     []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	sz := make([]int, size)
	for i := range sz {
		sz[i] = 1
	}
	return &UnionFind{
		parent: parent,
		sz:     sz,
	}
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

	if u.sz[pRoot] < u.sz[qRoot] {
		u.parent[pRoot] = qRoot
		u.sz[qRoot] += u.sz[pRoot]
	} else {
		u.parent[qRoot] = pRoot
		u.sz[pRoot] += u.sz[qRoot]
	}
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
