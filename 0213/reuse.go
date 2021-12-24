package leetcode

/*
reuse 198 的结果
 */
func _rob(nums []int) int {

	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}

func rob2(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	/*
		1. 不盗取 n-1 房间
		2. 不盗取  0 房间
	 */
	return max(_rob(nums[:len(nums)-1]), _rob(nums[1:]))
}
