package leetcode

// dp[i][j] =
/*
case:	p[j] is not . or *
		dp[i-1][j-1] if s[i] == p[j]

case:  p[j] is .
		dp[i-1][j-1]

case: p[j] is *, then we have two subcase use
	 1. dp[i][j-2]  不用
     2. dp[i-1][j]  if  s[i] == p[j-1]

 */
func isMatch(s string, p string) bool {

	n := len(s)
	m := len(p)

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}

	// 初始化
	dp[0][0] = true // no pattern match for empty string

	match := func(a, b int) bool {
		if a == 0 {
			return false
		}
		if p[b-1] == '.' {
			return true
		}
		return s[a-1] == p[b-1]
	}

	for i := 0; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '*' {
				//dp[i][j] = dp[i-1][j] || dp[i][j-1]
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if match(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}

			} else if match(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}

	return dp[n][m]
}
