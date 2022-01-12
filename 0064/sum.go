package leetcode

import "math"

/*
典型的DP 问题，
抓住 DP的要点， 1. 从子问题，求得当前问题的解， （转态方程） 2. 求解的顺序
 */

func minPathSum(grid [][]int) int {
	N := len(grid)
	M := len(grid[0])

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}

	getter := func(i, j int) int {
		if i < 0 || i >= N {
			return math.MaxInt64
		}

		if j < 0 || j >= M {
			return math.MaxInt64
		}

		return dp[i][j]
	}

	//dp[N-1][M-1] = grid[N-1][M-1]
	for i := N - 1; i >= 0; i-- {
		for j := M - 1; j >= 0; j-- {
			if i == N-1 && j == M-1 {
				dp[i][j] = grid[i][j]
				continue
			}
			dp[i][j] = min(getter(i+1, j), getter(i, j+1)) + grid[i][j]
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
