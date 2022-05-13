package OrderSlice

import "fmt"

func ExampleGetNoRepeat() {
	var s1 []int
	s2 := []int{1}
	s3 := []int{1, 1, 2, 2, 3, 3, 4, 5}
	s4 := []int{5, 5, 4, 4, 3, 3, 2, 1}

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

	s3NoRepeat, err := GetNoRepeat(s3)
	if err != nil {
		panic(err)
	}

	fmt.Println(s3NoRepeat)

	s4NoRepeat, err := GetNoRepeat(s4)
	if err != nil {
		panic(err)
	}
	fmt.Println(s4NoRepeat)

	//Output:
	//[]
	//[1]
	//[1 2 3 4 5]
	//[5 4 3 2 1]
}

func ExampleIsDESC() {
	var s1 []int
	s2 := []int{1}
	s3 := []int{1, 1}
	s4 := []int{1, 1, 2, 2, 3, 3, 4, 5}
	s5 := []int{5, 2, 1, 3}

	fmt.Println(IsDESC(s1))
	fmt.Println(IsDESC(s2))
	fmt.Println(IsDESC(s3))
	fmt.Println(IsDESC(s4))
	fmt.Println(IsDESC(s5))

	//Output:
	//true
	//true
	//true
	//true
	//false
}

func ExampleIsASC() {
	var s1 []int
	s2 := []int{1}
	s3 := []int{1, 1}
	s4 := []int{5, 5, 4, 4, 3, 3, 2, 1}
	s5 := []int{5, 2, 1, 3}

	fmt.Println(IsASC(s1))
	fmt.Println(IsASC(s2))
	fmt.Println(IsASC(s3))
	fmt.Println(IsASC(s4))
	fmt.Println(IsASC(s5))

	//Output:
	//true
	//true
	//true
	//true
	//false
}

func ExampleIsOrder() {
	var s1 []int
	s2 := []int{1}
	s3 := []int{1, 1}
	s4 := []int{1, 1, 2, 2, 3, 3, 4, 5}
	s5 := []int{5, 5, 4, 4, 3, 3, 2, 1}
	s6 := []int{5, 2, 1, 3}

	fmt.Println(IsOrder(s1))
	fmt.Println(IsOrder(s2))
	fmt.Println(IsOrder(s3))
	fmt.Println(IsOrder(s4))
	fmt.Println(IsOrder(s5))
	fmt.Println(IsOrder(s6))

	//Output:
	//true
	//true
	//true
	//true
	//true
	//false
}
