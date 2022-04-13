package leetcode

func findDerangement(n int) int {

	mod := 1000000007
	// dp[i] = (n-1)*dp[i-1] + dp[i-2]
	dp := make([]int, max(4, n+1))
	dp[0] = 1
	dp[1] = 0
	dp[2] = 1
	dp[3] = 2
	for i := 4; i <= n; i++ {
		dp[i] = (i - 1) * (dp[i-1] + dp[i-2])
		dp[i] %= mod
	}
	return dp[n]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
