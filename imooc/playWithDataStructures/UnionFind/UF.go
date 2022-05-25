package UnionFind

import "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/UnionFind/V1"

// UF 对于当下固定的元素进行并或者查的操作
type UF interface {
	GetSize() int
	// IsConnected 两个元素是否所述一个集合(是否可以连接的)
	IsConnected(p, q int) (bool, error)
	// UnionElements 将两个元素并在一起
	UnionElements(p, q int) error
}

// NewUnionFind 传入固定数据(ID)数量size
func NewUnionFind(size int) UF {
	return V1.NewUnionFind(size)
}
