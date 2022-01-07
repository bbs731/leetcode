package leetcode
/*
wow!  一次就过， 真是走了狗屎运!
 */
// dp[m][n] = dp[m-1][n] + dp[m][n-1]
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	found := false
	for i := 1; i <= m; i++ {
		if found {
			//dp[i][1] = 0
			continue
		}
		dp[i][1] = 1
		if isObstacle(obstacleGrid, i, 1) {
			dp[i][1] = 0
			found = true
		}
	}

	found = false
	for j := 1; j <= n; j++ {
		if found {
			//dp[1][j] = 0
			continue
		}
		dp[1][j] = 1
		if isObstacle(obstacleGrid, 1, j) {
			dp[1][j] = 0
			found = true
		}
	}

	for i := 1; i <= m; i++ {
		for j := 2; j <= n; j++ {
			if isObstacle(obstacleGrid, i, j) {
				//dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m][n]

}

func isObstacle(grid [][]int, m, n int) bool {
	if grid[m-1][n-1] == 1 {
		return true
	}
	return false
}
