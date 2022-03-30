package leetcode

func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)

	dp := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= l; i++ {
		for k := m; k >= 0; k-- {
			for p := n; p >= 0; p-- {
				// not take item i
				ones, zeros := counts(strs[i-1])
				if k >= zeros && p >= ones {
					// we can take item i
					dp[k][p] = max(dp[k][p], dp[k-zeros][p-ones]+1)

				}
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func counts(s string) (int, int) {
	ones, zeros := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zeros++
		} else if s[i] == '1' {
			ones++
		}
	}
	return ones, zeros
}
