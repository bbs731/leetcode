package leetcode

import "math"

/*
x1, 和 x2 总是相等的， 所以可以降到 3 维
和 741 采樱桃，是一个题型， 但是，741 更难一些！
 */
func cherryPickup(grid [][]int) int {
	N := len(grid)
	M := len(grid[0])
	minu := math.MinInt32

	dp := make([][][][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([][][]int, M)
		for j := 0; j < M; j++ {
			dp[i][j] = make([][]int, N)
			for k := 0; k < N; k++ {
				dp[i][j][k] = make([]int, M)
				for p := 0; p < M; p++ {
					dp[i][j][k][p] = minu
				}
			}
		}
	}

	var dfs func(int, int, int, int) int
	dfs = func(x1, y1, x2, y2 int) int {
		res := 0
		if x1 == N-1 {
			res += grid[x1][y1]
			if y1 != y2 {
				res += grid[x2][y2]
			}
			return res
		}

		// check y1 and y2 invalid
		if y1 < 0 || y1 >= M || y2 < 0 || y2 >= M {
			return 0
		}

		if dp[x1][y1][x2][y2] != minu {
			return dp[x1][y1][x2][y2]
		}

		ans := dfs(x1+1, y1, x2+1, y2)
		if y1-1 >= 0 {
			if y2-1 >= 0 {
				ans = max(ans, dfs(x1+1, y1-1, x2+1, y2-1))
			}
			ans = max(ans, dfs(x1+1, y1-1, x2+1, y2))
			if y2+1 < M {
				ans = max(ans, dfs(x1+1, y1-1, x2+1, y2+1))
			}
		}
		if y1+1 < M {
			if y2-1 >= 0 {
				ans = max(ans, dfs(x1+1, y1+1, x2+1, y2-1))
			}
			ans = max(ans, dfs(x1+1, y1+1, x2+1, y2))
			if y2+1 < M {
				ans = max(ans, dfs(x1+1, y1+1, x2+1, y2+1))
			}
		}

		if y2-1 >= 0 {
			ans = max(ans, dfs(x1+1, y1, x2+1, y2-1))
		}
		if y2+1 < M {
			ans = max(ans, dfs(x1+1, y1, x2+1, y2+1))
		}

		ans += grid[x1][y1]
		if y1 != y2 {
			ans += grid[x2][y2]
		}
		dp[x1][y1][x2][y2] = ans
		return ans
	}

	return dfs(0, 0, 0, M-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
