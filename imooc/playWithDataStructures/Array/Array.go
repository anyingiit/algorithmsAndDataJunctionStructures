package Array

import "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Array/Versions/V1"

type Array interface {
	Size() int
	IsEmpty() bool
	Append(e int)
	Has(e int) bool
	RemoveFirst(e int) (index int, err error)
	Get(index int) (e int, err error)
	IndexFirst(e int) (index int, err error)
	Set(index int, newE int) (err error)
}

func NewArray() Array {
	return V1.NewArray()
}

func NewArrayCustomCap(cap int) Array {
	return V1.NewArrayCustomCap(cap)
}
