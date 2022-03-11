package leetcode

import "fmt"

/*
一样的算法， 优雅觉得解决
 */

func isInterleave(s1 string, s2 string, s3 string) bool {
	n1 := len(s1)
	n2 := len(s2)

	if len(s3) != n1+n2 {
		return false
	}
	dp := make([][]bool, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]bool, n2+1)
	}
	dp[0][0] = true

	for i := 0; i <= n1; i++ {
		for j := 0; j <= n2; j++ {
			k := i + j - 1 // s3 current index // 由 s1[i] 和 s2[j] 组成
			if k < 0 {
				continue
			}
			if i >= 1 && s3[k] == s1[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if j >= 1 && s3[k] == s2[j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[n1][n2]

}
