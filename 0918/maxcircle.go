package leetcode
/*
好题， 打击信心。
反复做
 */
func maxSubarraySumCircular(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if sum[i-1] > 0 {
			sum[i] = nums[i] + sum[i-1]
		} else {
			sum[i] = nums[i]
		}
		ans = max(ans, sum[i])
	}

	rightsum := make([]int, len(nums))
	rightsum[len(nums)-1 ] = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		rightsum[i] = rightsum[i+1] + nums[i]
	}

	maxrightsum := make([]int, len(nums))
	maxrightsum[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i > 0; i-- {
		maxrightsum[i] = max(maxrightsum[i+1], rightsum[i])
	}

	leftsum := 0
	for i := 0; i < len(sum)-1; i++ {
		leftsum += nums[i]
		ans = max(ans, leftsum+maxrightsum[i+1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
