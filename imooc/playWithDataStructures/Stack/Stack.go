package Stack

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack/LinkStack"
)

type Stack interface {
	Push(e interface{})
	Pop() interface{}
	Size() int
	IsEmpty() bool
}

func NewStack() Stack {
	return LinkStack.NewLinkStack()
}
