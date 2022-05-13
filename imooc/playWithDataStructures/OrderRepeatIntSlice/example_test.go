package OrderRepeatIntSlice

import "fmt"

func ExampleGetNoRepeat() {
	s1 := []int{1, 1, 2, 2, 3, 3, 4, 5}
	s2 := []int{5, 5, 4, 4, 3, 3, 2, 1}

	s1NoRepeat := GetNoRepeat(s1)
	s2NoRepeat := GetNoRepeat(s2)
	fmt.Println(s1NoRepeat)
	fmt.Println(s2NoRepeat)

	//Output:
	//[1 2 3 4 5]
	//[5 4 3 2 1]
}
