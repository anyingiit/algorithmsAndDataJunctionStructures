package Stack

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack/LinkStack"
)

type Stack interface {
	Enqueue(e interface{})
	Dequeue() interface{}
	Size() int
	IsEmpty() bool
}

func NewStack() Stack {
	return LinkStack.NewLinkStack()
}
