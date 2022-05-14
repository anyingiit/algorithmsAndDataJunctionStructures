package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Queue/SliceQueue"
	"fmt"
)

func main() {
	queue := SliceQueue.NewSliceQueue()
	queue.Test()

	queue.Push(5)
	queue.Push(4)
	queue.Push(3)
	queue.Push(2)
	queue.Push(1)
	queue.Push(0)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	fmt.Println("queue size:", queue.Size())
}
