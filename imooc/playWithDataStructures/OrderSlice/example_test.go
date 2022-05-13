package OrderSlice

import "fmt"

func ExampleGetNoRepeat() {
	s1 := []int{1, 1, 2, 2, 3, 3, 4, 5}
	s2 := []int{5, 5, 4, 4, 3, 3, 2, 1}

	s1NoRepeat, err := GetNoRepeat(s1)
	if err != nil {
		panic(err)
	}

	fmt.Println(s1NoRepeat)

	s2NoRepeat, err := GetNoRepeat(s2)
	if err != nil {
		panic(err)
	}
	fmt.Println(s2NoRepeat)

	//Output:
	//[1 2 3 4 5]
	//[5 4 3 2 1]
}
