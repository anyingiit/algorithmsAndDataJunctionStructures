package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack/SliceStack"
	"fmt"
)

func main() {
	stack := SliceStack.NewSliceStack()
	stack.Push(5)
	stack.Push(5)
	stack.Push(5)
	stack.Push(5)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())

}
