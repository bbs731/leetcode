package leetcode

func minInsertions(s string) int {
	n := len(s)
	//ans := math.MaxInt64 >> 1
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ { // 把 j 的 check 加在这里， 看起来更简洁吗？ 好像是的呢
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
