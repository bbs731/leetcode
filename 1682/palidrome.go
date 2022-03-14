package leetcode

/*
Typical DP problem 单串，但是多个维度的 DP 问题
 */
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	last := make([][]map[byte]struct{}, n) // save the longest palindrome, the outmost chars

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		last[i] = make([]map[byte]struct{}, n)
		for j := 0; j < n; j++ {
			last[i][j] = make(map[byte]struct{})
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := l - 1 + i
			if j < n {
				// dp[i][j] = max{ dp[i+1][j], dp[i][j-1], dp[i+1][j-1] }
				dp[i][j] = dp[i+1][j]
				last[i][j] = copyamap(last[i+1][j])

				if dp[i][j-1] > dp[i][j] {
					dp[i][j] = dp[i][j-1]
					last[i][j] = copyamap(last[i][j-1])
				} else if dp[i][j-1] == dp[i][j] {
					last[i][j] = mergemap(last[i][j], last[i][j-1])
				}

				if s[i] == s[j] {
					// check 没有连续的相同的字符
					_, ok := last[i+1][j-1][s[i]]
					if len(last[i+1][j-1]) > 1 || !ok {
						if dp[i+1][j-1]+2 > dp[i][j] {
							dp[i][j] = dp[i+1][j-1] + 2
							last[i][j] = map[byte]struct{}{s[i]: struct{}{}}
						} else if dp[i+1][j-1]+2 == dp[i][j] {
							last[i][j] = mergemap(last[i][j], map[byte]struct{}{s[i]: struct{}{}})
						}
					}
				}
			}
		}
	}
	return dp[0][n-1]
}

func copyamap(candidates map[byte]struct{}) map[byte]struct{} {
	acopy := make(map[byte]struct{})
	for k, v := range candidates {
		acopy[k] = v
	}
	return acopy
}
func mergemap(target, candidates map[byte]struct{}) map[byte]struct{} {
	for k, v := range candidates {
		target[k] = v
	}
	return target
}
