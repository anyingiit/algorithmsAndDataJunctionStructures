package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Queue/SliceQueue"
	"fmt"
)

func main() {
	queue := SliceQueue.NewSliceQueue()
	queue.Test()

	queue.Enqueue(5)
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Enqueue(0)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println("queue size:", queue.Size())
}
