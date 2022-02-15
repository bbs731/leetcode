package leetcode

// index 从 1 开始
func findTargetSumWays(nums []int, target int) int {
	// we find the target + 2000

	n := len(nums)
	dp := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 4001)
	}
	// 初始化：
	dp[0][2000] = 1
	for i := 1; i <= n; i++ {
		for j := 1000; j <= 3000; j++ {
			dp[i][j] = dp[i-1][j-nums[i-1]] + dp[i-1][j+nums[i-1]]
		}
	}
	return dp[n][target+2000]
}
