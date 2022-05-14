package DeleteOrderArrayRepeats

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
