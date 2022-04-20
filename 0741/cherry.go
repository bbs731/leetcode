package leetcode

import (
	"fmt"
	"math"
)

/*
这道题，太难了， 做不出来！
 */
func cherryPickup(grid [][]int) int {
	N := len(grid)
	dp := make([][][]int, N)
	min := math.MinInt32 >> 1
	for i := 0; i < N; i++ {
		dp[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = make([]int, N)
			for k := 0; k < N; k++ {
				dp[i][j][k] = min
			}
		}
	}

	var dfs func(int, int, int) int
	dfs = func(x1 int, y1 int, x2 int) int {
		y2 := x1 + y1 - x2

		if x1 >= N || x2 >= N || y1 >= N || y2 >= N || grid[x1][y1] == -1 || grid[x2][y2] == -1 {
			return -1
		}
		if x1 == N-1 && y1 == N-1 {
			return grid[x1][y1]
		}
		if dp[x1][y1][x2] != min {
			return dp[x1][y1][x2]
		}
		ans := max(dfs(x1+1, y1, x2), dfs(x1+1, y1, x2+1))
		ans = max(ans, dfs(x1, y1+1, x2+1))
		ans = max(ans, dfs(x1, y1+1, x2))

		// 这个在做什么？
		if ans < 0 {
			dp[x1][y1][x2] = -9
			return -9
		}

		ans += grid[x1][y1]
		if x1 != x2 {
			ans += grid[x2][y2]
		}

		dp[x1][y1][x2] = ans
		return ans
	}

	res := max(0, dfs(0, 0, 0))
	fmt.Println(dp)

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
