package leetcode

/*
还有 O(n) 的 solution
 */
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func jump(nums []int) int {
	steps := make([]int, len(nums))

	for i := 0; i < len(steps); i++ {
		steps[i] = len(steps)
	}
	steps[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+nums[i] && j < len(nums); j++ {
			steps[j] = min(steps[j], steps[i]+1)
		}
	}

	return steps[len(steps)-1]
}
