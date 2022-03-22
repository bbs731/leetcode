package leetcode

/*
用 dp[i][j] 表示长度， 有点画蛇添足的意思。 改进 r[i][j] bool 那个版本吧，有时间再试一次。

本做法是用的， range 区间 dp 的模板！
 */
func longestPalindrome(s string) string {
	n := len(s)
	ans := s[:1]
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j < n {
				if s[i] == s[j] {
					if l == 2 {
						dp[i][j] = 2
						if 2 > len(ans) {
							ans = s[i : j+1]
						}
					} else {
						if dp[i+1][j-1] == l-2 {
							dp[i][j] = dp[i+1][j-1] + 2
							if dp[i][j] > len(ans) {
								ans = s[i : j+1]
							}
						}
					}
				} else {
					dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				}
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
