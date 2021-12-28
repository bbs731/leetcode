package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func jump2(nums []int) int {

	far := 0
	step := 0
	position := 0
	for i := 0; i < len(nums)-1; i++ {
		position = max(position, nums[i]+i)
		if i == far {
			far = position
			step++
		}
	}
	return step
}
