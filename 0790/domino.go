package leetcode

// dp[n] = dp[n-1] + (dp[n-2]*2) + (dp[n-3]*2)
func numTilings(n int) int {
	mod := 1000000007
	dp := [4]int{1, 0, 0, 0}
	for i := 0; i < n; i++ {
		dp[0], dp[1], dp[2], dp[3] =
			dp[3]+dp[0], dp[0]+dp[2], dp[0]+dp[1], dp[0]+dp[1]+dp[2]
		dp[0] %= mod
		dp[1] %= mod
		dp[2] %= mod
		dp[3] %= mod
	}
	return dp[0]
}
