package leetcode

func canPartition(nums []int) bool {
	n := len(nums)

	dp := make([][]bool, len(nums)+1)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	// 初始化
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			//dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][target]
}
