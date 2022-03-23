package leetcode

/*
wrong answer

看官网吧，关键那个 Dp 的公式，不可能想出来啊！
 */
func countPalindromicSubsequences(s string) int {
	n := len(s)
	mod := 1000000007

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			if s[i] == s[j] {
				dp[i][j] = (dp[i+1][j-1]*2 + 2) % mod
				// minus overlap
				// find the overlap
				m, n := i+1, j-1
				for ; m < j; m++ {
					if s[m] == s[i] {
						break
					}
				}
				for ; n > i; n-- {
					if s[n] == s[i] {
						break
					}
				}
				if m == j || n == i || m > n {
					continue
				}
				dp[i][j] = (dp[i][j] - dp[m][n]) % mod

			} else {
				dp[i][j] = (dp[i+1][j] + dp[i][j-1] - dp[i+1][j-1]) % mod
			}
		}
	}
	return dp[0][n-1]
}
