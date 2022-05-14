package BST

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/disableVersion/V4"
	"fmt"
)

// BST no repeat BST
type BST interface {
	fmt.Stringer
	Size() int
	IsEmpty() bool
	// Push element to BST
	//
	// warning: this BST cannot support repeat element, so this method well return do you really success push to this BST
	// if this method return is false, then because you try push to a repeat element to this BST
	Push(e int) (isSuccess bool)
	Has(e int) bool
	GetPreOrder() (result []int)
	GetInOrder() (result []int)
	GetPostOrder() (result []int)
	PrintPreOrder()
	PrintInOrder()
	PrintPostOrder()
	GetPreOrderNR() (result []int)
	GetLevelOrderNR() (result []int)
	GetPreOrderFormatString() string
	Minimum() (minimumElement int, err error)
	Maximum() (maximumElement int, err error)
	RemoveMin() (deletedElement int, err error)
	RemoveMax() (deletedElement int, err error)
	Remove(e int) error
}

// NewBST return new no repeat BST
func NewBST() BST {
	return V4.NewBST()
}
