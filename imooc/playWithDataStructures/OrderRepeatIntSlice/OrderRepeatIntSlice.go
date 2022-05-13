package OrderRepeatIntSlice

// GetNoRepeat 输入有序的数组(正序和道具都可), 返回无重复的有序数组(不改变输入顺序)
func GetNoRepeat(origin []int) (result []int) {
	if len(origin) == 0 {
		return nil
	}
	if len(origin) == 1 {
		return origin
	}

	cp := make([]int, len(origin))
	copy(cp, origin)

	i, j := 0, 0

	for j < len(cp) {
		if cp[i] != cp[j] {
			i++
			cp[i] = cp[j]
		}
		j++
	}
	return cp[:i+1]
}
