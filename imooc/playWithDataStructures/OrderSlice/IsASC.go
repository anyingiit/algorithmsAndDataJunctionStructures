package OrderSlice

func IsASC(x []int) bool {
	if len(x) == 0 || len(x) == 1 {
		return true
	}
	for i := 0; i < len(x)-1; i++ {
		if x[i] < x[i+1] {
			return false
		}
	}
	return true
}
