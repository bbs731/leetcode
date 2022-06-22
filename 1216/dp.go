package leetcode

/*
可以啊兄弟，开挂了吗？
 */
func isValidPalindrome(s string, k int) bool {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i >= j {
				dp[i][j] = 0
			} else {
				dp[i][j] = 2 * n // infinity
			}
		}
	}

	for l := 2; l <= n; l++ { // 这里，注意从2 起。
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = min(dp[i][j], dp[i+1][j-1])
			} else {
				dp[i][j] = min(dp[i][j], 2+dp[i+1][j-1])
			}
			dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
			dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
		}
	}

	return dp[0][n-1] <= k
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
