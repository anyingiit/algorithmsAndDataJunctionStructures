package V1

import "fmt"

// UnionFind for quick find
type UnionFind struct {
	id []int
}

func NewUnionFind(size int) *UnionFind {
	list := make([]int, size)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
	return &UnionFind{list}
}

func (u *UnionFind) GetSize() int {
	return len(u.id)
}

func (u *UnionFind) IsConnected(p, q int) (bool, error) {
	pId, err := u.find(p)
	if err != nil {
		return false, err
	}
	qId, err := u.find(q)
	if err != nil {
		return false, err
	}
	return pId == qId, nil
}

func (u *UnionFind) UnionElements(p, q int) error {
	pId, err := u.find(p)
	if err != nil {
		return err
	}
	qId, err := u.find(q)
	if err != nil {
		return err
	}
	if pId == qId {
		return nil
	}
	for i := 0; i < u.GetSize(); i++ {
		if u.id[i] == qId {
			u.id[i] = pId
		}
	}
	return nil
}

// find 查找元素p对应的集合编号
func (u *UnionFind) find(p int) (int, error) {
	if p < 0 || p > u.GetSize()-1 {
		return 0, fmt.Errorf("p out of bound")
	}
	return u.id[p], nil
}
