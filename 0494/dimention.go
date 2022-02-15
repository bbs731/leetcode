package leetcode


// dp[i][target] = dp[i-1][target+nums[i]] + dp[i-1][target-nums[i]]
// 这个降维的 solution 是错了， 我们降不了维度
func findTargetSumWays(nums []int, target int) int {
	// we find the target + 2000

	n := len(nums)
	dp := make([]int, 4001)

	dp[2000] = 1

	for i := 1; i <= n; i++ {
		for j := 1000; j <= 3000; j++ {
			dp[j] = dp[j-nums[i-1]] + dp[j+nums[i-1]]
		}
	}
	return dp[target+2000]
}

