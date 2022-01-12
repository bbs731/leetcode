package leetcode

import "math"

/*
我操，你额太NB了， 不要给自己设置不必要的条件。

譬如 i > j 时尝试 delete 不尝试 replace 等等
 */

func minDistance(word1 string, word2 string) int {
	N1 := len(word1)
	N2 := len(word2)

	if N1 == 0 {
		return N2
	}
	if N2 == 0 {
		return N1
	}

	dp := make([][]int, N1)
	for i := 0; i < N1; i++ {
		dp[i] = make([]int, N2)
		for j := 0; j < N2; j++ {
			dp[i][j] = math.MaxInt64
		}
	}

	getter := func(i, j int) int {
		if i < 0 {
			if j < 0 {
				return 0
			}
			return j + 1
		} else {
			if j < 0 {
				return i + 1
			}
		}
		return dp[i][j]
	}
	for i := 0; i < N1; i++ {
		for j := 0; j < N2; j++ {
			if word1[i] == word2[j] {
				dp[i][j] = min(dp[i][j], getter(i-1, j-1))
			} else {
				//try replace
				dp[i][j] = min(dp[i][j], getter(i-1, j-1)+1)
				// try delete ?
				dp[i][j] = min(dp[i][j], getter(i-1, j)+1)
				// insert ?
				dp[i][j] = min(dp[i][j], getter(i, j-1)+1)
			}
		}
	}

	return dp[N1-1][N2-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
