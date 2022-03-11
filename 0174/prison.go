package leetcode

import "math"

/*
 You are the knight, submission once got pass!
 */
func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	m := len(dungeon[0])

	dp := make([][]int, n)
	// dp[i][j] 代表， 进入房间 (i,j) 的时候，至少需要带的血量. dp[][] 有个性质， 每个 elments >=1
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	if dungeon[n-1][m-1] >= 0 {
		dp[n-1][m-1] = 1
	} else {
		dp[n-1][m-1] = -dungeon[n-1][m-1] + 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if i == n-1 && j == m-1 {
				continue
			}
			// 初始化
			dp[i][j] = math.MaxInt64
			// 考虑，下一步向右边走。
			if j+1 < m {
				c := dp[i][j+1] - dungeon[i][j]
				if c <= 0 {
					c = 1
				}
				dp[i][j] = min(dp[i][j], c)
			}
			if i+1 < n {
				c := dp[i+1][j] - dungeon[i][j]
				if c <= 0 {
					c = 1
				}
				dp[i][j] = min(dp[i][j], c)
			}
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
