package leetcode

import "math"

func minimumTotal(triangle [][]int) int {
	N := len(triangle)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, i+1)
	}
	// 初始化
	for i := 0; i < N; i++ {
		for j := 0; j < i+1; j++ {
			if i == N-1 {
				dp[N-1][j] = triangle[N-1][j]
			} else {
				dp[i][j] = math.MaxInt64
			}
		}
	}

	for i := N - 2; i >= 0; i-- {
		m := len(triangle[i])
		for j := 0; j < m; j++ {
			dp[i][j] = min(dp[i][j], triangle[i][j]+min(dp[i+1][j], dp[i+1][j+1]))
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
