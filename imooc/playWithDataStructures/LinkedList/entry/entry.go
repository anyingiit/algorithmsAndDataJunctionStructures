package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/LinkedList"
	"fmt"
)

func main() {
	linkedList := LinkedList.NewLinkedList()
	linkedList.Add(10)
	linkedList.Add(5)
	linkedList.Add(3)
	linkedList.Add(1)
	linkedList.Add(9)
	fmt.Println("get all:", linkedList.GetAll())
	fmt.Println("linkedList size:", linkedList.Size())
	fmt.Println("has 10?", linkedList.Has(10))
	fmt.Println("has 90?", linkedList.Has(90))
	e, err := linkedList.Del(10)
	fmt.Print("del 10:")
	if err != nil {
		fmt.Println("failed:", err)
	} else {
		fmt.Println("success: return is", e)
	}
	fmt.Println("linkedList size:", linkedList.Size())

	fmt.Println("has 5?", linkedList.Has(5))
	fmt.Print("set element 5 to 233: ")
	err = linkedList.Set(5, 233)
	if err != nil {
		fmt.Println("failed:", err)
	} else {
		fmt.Printf("success: now has5? %v has233? %v\n", linkedList.Has(5), linkedList.Has(233))
	}
	fmt.Println("get all:", linkedList.GetAll())
}
