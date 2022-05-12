package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/V3"
	"fmt"
)

func main() {
	bst := V3.NewBST()

	bst.Push(41)
	bst.Push(22)
	bst.Push(58)
	bst.Push(15)
	bst.Push(33)
	bst.Push(13)
	bst.Push(50)
	bst.Push(42)
	bst.Push(53)

	bst.PreOrder()
	fmt.Println()

	bst.InOrder()
	fmt.Println()

	bst.PostOrder()
	fmt.Println()

	fmt.Println(bst)

	minimum, err := bst.Minimum()
	if err != nil {
		panic(err)
	}
	fmt.Println("minimum:", minimum)

	maximum, err := bst.Maximum()
	if err != nil {
		panic(err)
	}
	fmt.Println("maximum:", maximum)

	err = bst.Remove(50)
	if err != nil {
		fmt.Println(err)
	}

	err = bst.Remove(233)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bst)
}
