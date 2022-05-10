package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/Stack"
	"fmt"
)

func main() {
	stack := Stack.NewStack()
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
