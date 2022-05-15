package BubbleSort

import (
	"fmt"
)

func ExampleSetSortByASC() {
	s := []int{5, 8, 1, 4, 2}
	SetSortByASC(s)
	fmt.Println(s)

	//Output:
	//[8 5 4 2 1]
}

func ExampleSetSortByDESC() {
	s := []int{5, 8, 1, 4, 2}
	SetSortByDESC(s)
	fmt.Println(s)

	//Output:
	//[1 2 4 5 8]
}

func ExampleGetSortByASC() {
	s := []int{5, 8, 1, 4, 2}
	result := GetSortByASC(s)
	fmt.Println("not well changed origin slice s:", s)
	fmt.Println("result:", result)

	//Output:
	//not well changed origin slice s: [5 8 1 4 2]
	//result: [8 5 4 2 1]
}

func ExampleGetSortByDESC() {
	s := []int{5, 8, 1, 4, 2}
	result := GetSortByDESC(s)
	fmt.Println("not well changed origin slice s:", s)
	fmt.Println("result:", result)

	//Output:
	//not well changed origin slice s: [5 8 1 4 2]
	//result: [1 2 4 5 8]
}
