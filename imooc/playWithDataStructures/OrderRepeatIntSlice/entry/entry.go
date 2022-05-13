package main

import (
	"algorithmsAndDataJunctionStructures/imooc/playWithDataStructures/OrderRepeatIntSlice"
	"fmt"
)

func main() {
	s := []int{1, 1, 2, 2, 3, 3, 4, 5}
	noRepeat := OrderRepeatIntSlice.GetNoRepeat(s)
	fmt.Println(noRepeat)

	s1 := []int{1, 1}
	noRepeat = OrderRepeatIntSlice.GetNoRepeat(s1)
	fmt.Println(noRepeat)
}
