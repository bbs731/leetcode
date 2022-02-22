package leetcode

/*
这是一个 DP 问题。 是变种的 前缀和问题！
 */
func maxSumAfterOperation(nums []int) int {

	// 前缀和
	dp := make([]int, len(nums))
	changed := make([]int, len(nums))
	// 初始化
	if nums[0] > 0 {
		dp[0] = nums[0]
	}
	changed[0] = nums[0] * nums[0]

	ans := changed[0]

	for i := 1; i < len(nums); i++ {
		// update dp
		if dp[i-1]+nums[i] > 0 {
			dp[i] = dp[i-1] + nums[i]
		}
		// update changed
		changed[i] = max(dp[i-1]+nums[i]*nums[i], changed[i-1]+nums[i])
		ans = max(dp[i], ans)
		ans = max(changed[i], ans)

	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
