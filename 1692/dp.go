package leetcode

func waysToDistribute(n int, k int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	mod := 1000000000 + 7

	var dfs func(int, int) int

	dfs = func(n int, k int) int {
		if n < k {
			return 0
		}
		if k == 1 {
			return 1
		}
		if n == k {
			return 1
		}
		if dp[n][k] != 0 {
			return dp[n][k]
		}
		dp[n][k] = (k*dfs(n-1, k) + dfs(n-1, k-1)) % mod
		return dp[n][k]
	}
	return dfs(n, k) % mod
}
