package leetcode

/*
wow! my god!  one submit pass!
 */
func numDistinct(s string, t string) int {
	n := len(s)
	m := len(t)

	// dp[i][j] 表示 s[:i] 中子串  t[:j] 出现的次数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	//dp[0][0] = 1 //
	//初始化
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][m]
}
