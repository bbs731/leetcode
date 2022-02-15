package leetcode

// index 从 0 开始
func findTargetSumWays(nums []int, target int) int {
	// we find the target + 2000

	n := len(nums)
	dp := make([][]int, n+1)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, 4001)
	}
	// 初始化：
	// index 从0 开始， 这里初始化，有一个陷阱。

	if nums[0] == 0 {
		dp[0][2000] = 2
	} else {
		dp[0][nums[0]+2000 ] = 1
		dp[0][2000-nums[0]] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1000; j <= 3000; j++ {
			dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j+nums[i]]
		}
	}
	return dp[n-1][target+2000]
}
