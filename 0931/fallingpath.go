package leetcode

import "math"


/*
典型的 DP 问题， 关键看计算顺序
 */
 
func minFallingPathSum(matrix [][]int) int {
	N := len(matrix)

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	getter := func(i, j int) int {
		if i >= N {
			return 0
		}
		if j < 0 || j >= N {
			return math.MaxInt64
		}
		return dp[i][j]
	}

	// DP 的关键，就是计算顺序
	for i := N - 1; i >= 0; i-- {
		for j := 0; j < N; j++ {
			dp[i][j] = min(getter(i+1, j), min(getter(i+1, j-1), getter(i+1, j+1))) + matrix[i][j]
		}
	}

	ans := math.MaxInt64
	for i := 0; i < N; i++ {
		ans = min(ans, dp[0][i])
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
