package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack/SliceStack"
	"fmt"
)

func main() {
	stack := SliceStack.NewSliceStack()
	stack.Enqueue(5)
	stack.Enqueue(5)
	stack.Enqueue(5)
	stack.Enqueue(5)
	fmt.Println(stack.Dequeue())
	fmt.Println(stack.Dequeue())
	fmt.Println(stack.Dequeue())
	fmt.Println(stack.Dequeue())
	fmt.Println(stack.Size())

}
