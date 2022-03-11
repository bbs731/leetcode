package leetcode

/*
你TMD， 真是个天才!!!!
 */
func isScramble(s1 string, s2 string) bool {
	n := len(s1)
	if n != len(s2) {
		return false
	}
	// dp[i][j][len] // s1 start with index 1 , s2 start with index j, and string length is len is sramble ?

	dp := make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
		}
	}

	for l := 1; l <= n; l++ {
		for i := 0; i < n && i+l-1 < n; i++ {
			for j := 0; j < n && j+l-1 < n; j++ {
				if l == 1 { // 这个当初始化了吗？
					if s1[i] == s2[j] {
						dp[i][j][1] = true
					}
				} else {
					// otherwise l >=2
					for cut := 1; cut <= l-1; cut++ { // possible k cut
						// no scramble case
						dp[i][j][l] = dp[i][j][l] || (dp[i][j][cut] && dp[i+cut][j+cut][l-cut])
						// scramble case
						dp[i][j][l] = dp[i][j][l] || (dp[i][j+l-cut][cut] && dp[i+cut][j][l-cut])
					}
				}
			}
		}
	}

	return dp[0][0][n]
}
