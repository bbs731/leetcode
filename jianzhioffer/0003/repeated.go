package offer

func findRepeatNumber(nums []int) int {

	table := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if _, ok := table[nums[i]]; ok {
			return nums[i]
		}
		table[nums[i]] = true
	}
	return nums[0]
}
