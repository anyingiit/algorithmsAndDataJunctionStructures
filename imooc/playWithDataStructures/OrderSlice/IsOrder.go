package OrderSlice

func IsOrder(x []int) bool {
	return IsDESC(x) || IsASC(x)
}
