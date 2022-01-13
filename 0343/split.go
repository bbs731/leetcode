package leetcode

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		if i != n {
			dp[i] = i
		}
		for j := 1; j < i/2+1; j++ {
			dp[i] = max(dp[i], dp[j]*dp[i-j])
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
