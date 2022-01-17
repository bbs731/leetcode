package offer

func search(nums []int, target int) int {

	times := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			times++
		}
	}
	return times
}
