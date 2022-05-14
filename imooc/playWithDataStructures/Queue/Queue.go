package Queue

import "algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Queue/LinkQueue"

type Queue interface {
	Push(e interface{})
	Pop() interface{}
	Size() int
	IsEmpty() bool
}

func NewQueue() Queue {
	return LinkQueue.NewLinkQueue()
}
