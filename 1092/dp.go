package leetcode

import (
	"bytes"
)

/*
完全参考的， 算法导论，里面的，  LCS  实现！
看来，你的基础还是不过关啊！
 */

func shortestCommonSupersequence(str1 string, str2 string) string {

	m := len(str1)
	n := len(str2)

	dp := make([][]int, m+1)
	b := make([][]int, m+1)

	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		b[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				b[i][j] = 3 // take both str1 and str2
			} else if dp[i-1][j] >= dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
				b[i][j] = 1 // take str1
			} else {
				dp[i][j] = dp[i][j-1]
				b[i][j] = 2 // stake str2
			}
		}
	}

	var buff bytes.Buffer
	var left string

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i == 0 {
			left = str2[:j]
			return
		}
		if j == 0 {
			left = str1[:i]
			return
		}
		if b[i][j] == 3 {
			dfs(i-1, j-1)
			buff.WriteByte(str1[i-1])
		} else if b[i][j] == 2 {
			dfs(i, j-1)
			// take str2
			buff.WriteByte(str2[j-1])

		} else {
			dfs(i-1, j)
			// take str 1
			buff.WriteByte(str1[i-1])
		}
	}

	dfs(m, n)
	return left + buff.String()

}
