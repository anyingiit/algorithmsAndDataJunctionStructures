package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/disableVersion/V2"
	"fmt"
)

func main() {
	bst := V2.NewBST()
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
	fmt.Println("bst size:", bst.Size())
	fmt.Println("has 50?", bst.Has(50))
	fmt.Println("has 0?", bst.Has(0))

	// 老师留的作业, 在完成二分搜索树学习后可以尝试完成
	// 如何通过字符的方式打印易读的二叉树?

	fmt.Println(bst)

	bst.InOrder()
	fmt.Println()

	bst.PostOrder()
	fmt.Println()

	bst.PreOrderNR()
	fmt.Println()

	bst.LevelOrder()
	fmt.Println()

	minimum, err := bst.Minimum()
	fmt.Print("bst minimum: ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(minimum)
	}

	maximum, err := bst.Maximum()
	fmt.Print("bst maximum: ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maximum)
	}

	e, err := bst.RemoveMin()

	fmt.Print("bst delete minimum: ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(e)
	}

	minimum, err = bst.Minimum()
	fmt.Print("bst minimum: ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(minimum)
	}

	maximum, err = bst.Maximum()
	fmt.Print("bst maximum: ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maximum)
	}

	e, err = bst.RemoveMax()

	fmt.Print("bst delete maximum: ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(e)
	}

	maximum, err = bst.Maximum()
	fmt.Print("bst maximum: ")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(maximum)
	}
	fmt.Println()

	fmt.Println(bst)
	fmt.Println()

	err = bst.Remove(50)
	if err != nil {
		fmt.Println("delete element 50 failed:", err)
	}
	fmt.Println()

	fmt.Println(bst)

}
