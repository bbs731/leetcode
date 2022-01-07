package leetcode

/*
动态规则的递推公式我们已经有了。 然后我们需要考虑的，就是计算顺序的问题了:  index 按照什么样的顺序计算， 能让子问题的解先计算出来
 */
// dp[m][n] = dp[m-1][n] + dp[m][n-1]
func uniquePaths(m int, n int) int {


	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}
	//for j := 1; j <= n; j++ {
	//	dp[1][j] = 1
	//}

	for i := 1; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]
}
