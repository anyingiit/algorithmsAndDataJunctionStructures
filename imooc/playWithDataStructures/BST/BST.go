package BST

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/versions/V5"
	V5E "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/versions/V5/E"
	V5EInt "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/versions/V5/E/EInt"
	"fmt"
)

type E = V5E.E

type EInt = V5EInt.EInt

// BST no repeat BST
type BST interface {
	fmt.Stringer
	Size() int
	IsEmpty() bool
	// Push element to BST
	//
	// warning: this BST cannot support repeat element, so this method well return do you really success push to this BST
	// if this method return is false, then because you try push to a repeat element to this BST
	Push(e E) (isSuccess bool)
	Has(e E) bool
	GetPreOrder() (result []E)
	GetInOrder() (result []E)
	GetPostOrder() (result []E)
	PrintPreOrder()
	PrintInOrder()
	PrintPostOrder()
	GetPreOrderNR() (result []E)
	GetLevelOrderNR() (result []E)
	GetPreOrderFormatString() string
	Minimum() (minimumElement E, err error)
	Maximum() (maximumElement E, err error)
	RemoveMin() (deletedElement E, err error)
	RemoveMax() (deletedElement E, err error)
	Remove(e E) error
}

// NewBST return new no repeat BST
func NewBST() BST {
	return V5.NewBST()
}
