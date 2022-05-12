package BubbleSort

import (
	"fmt"
)

func ExampleSortByASC() {
	s := []int{5, 8, 1, 4, 2}
	SortByASC(s)
	fmt.Println(s)

	//Output:
	//[8 5 4 2 1]
}

func ExampleSortByDESC() {
	s := []int{5, 8, 1, 4, 2}
	SortByDESC(s)
	fmt.Println(s)

	//Output:
	//[1 2 4 5 8]
}
