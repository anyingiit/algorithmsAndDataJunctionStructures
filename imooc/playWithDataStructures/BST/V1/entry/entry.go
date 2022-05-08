package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/V1"
	"fmt"
)

func main() {
	bst := V1.NewBST()
	bst.Push(41)
	bst.Push(22)
	bst.Push(58)
	bst.Push(15)
	bst.Push(33)
	bst.Push(13)
	bst.Push(37)
	bst.Push(50)
	bst.Push(42)
	bst.Push(53)
	bst.Print()
	fmt.Println("bst size:", bst.GetSize())
}
