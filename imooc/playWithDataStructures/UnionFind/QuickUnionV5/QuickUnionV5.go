package QuickUnionV5

import "fmt"

// UnionFind for quick union 基于rank and Path Compression(相较于V4版, 该版本会将树保持在两层的高度)的优化, 虽然树的高度更低了, 但是压缩的过程也有了更多的开销, 所以实际表现上略微差V4一点点
type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	for i := range parent {
		parent[i] = i
	}
	rank := make([]int, size)
	for i := range rank {
		rank[i] = 1
	}
	return &UnionFind{
		parent: parent,
		rank:   rank,
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

	if u.rank[pRoot] < u.rank[qRoot] {
		// 0	0
		// 0	0
		// 0
		u.parent[pRoot] = qRoot
	} else if u.rank[qRoot] < u.rank[pRoot] {
		// 0	0
		// 0	0
		// 		0
		u.parent[qRoot] = pRoot
	} else { //u.rank[qRoot] == u.rank[pRoot]
		// 0	0
		// 0	0
		// 0	0
		u.parent[pRoot] = qRoot
		u.rank[qRoot] += 1
	}
	return nil
}

func (u *UnionFind) find(p int) (int, error) {
	if p < 0 || p > u.GetSize()-1 {
		return 0, fmt.Errorf("p out of bound")
	}

	var f func(p int) int
	f = func(p int) int {
		if u.parent[p] == p {
			return p
		}
		u.parent[p] = f(u.parent[p])
		return u.parent[p]
	}

	return f(p), nil
}
