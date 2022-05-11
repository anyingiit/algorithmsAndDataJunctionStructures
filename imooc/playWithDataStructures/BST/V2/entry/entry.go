package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/BST/V2"
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
	fmt.Println("bst size:", bst.GetSize())
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
}
