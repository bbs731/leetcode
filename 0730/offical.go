package leetcode

/*
oh , my god ! 能他妈的看懂答案吗？
 */
func countPalindromicSubsequences(s string) int {
	n := len(s)
	mod := 1000000007

	//
	dp := make([][][]int, 4)
	for k := 0; k < 4; k++ {
		dp[k] = make([][]int, n)
		for i := 0; i < n; i++ {
			dp[k][i] = make([]int, n)
		}

	}
	// 初始化
	for i := 0; i < n; i++ {
		k := s[i] - 'a'
		dp[k][i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			for k := 0; k < 4; k++ {
				c := 'a' + k
				if int(s[i]) != c {
					dp[k][i][j] = dp[k][i+1][j]
				} else if int(s[j]) != c {
					dp[k][i][j] = dp[k][i][j-1]
				} else {
					for m := 0; m < 4; m++ {
						dp[k][i][j] = (dp[k][i][j] + dp[m][i+1][j-1]) % mod
					}
					dp[k][i][j] += 2

				}
			}
		}
	}

	total := 0
	for k := 0; k < 4; k++ {
		total = (total + dp[k][0][n-1]) % mod
	}
	return total
}
