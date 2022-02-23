package leetcode

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])

	visited := make([][]bool, m)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		dp[i] = make([]int, n)
	}

	var getAround func(int, int) [][]int
	getAround = func(i, j int) [][]int {
		dirs := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
		candidates := make([][]int, 0)
		for k := 0; k < len(dirs); k++ {
			if dirs[k][0]+i >= 0 && dirs[k][0]+i < m && dirs[k][1]+j >= 0 && dirs[k][1]+j < n && matrix[dirs[k][0]+i][dirs[k][1]+j] > matrix[i][j] {
				candidates = append(candidates, []int{dirs[k][0] + i, dirs[k][1] + j})
			}
		}
		return candidates
	}

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n {
			return 0
		}
		candidates := getAround(i, j)
		for _, candidate := range candidates {
			if visited[candidate[0]][candidate[1]] {
				dp[i][j] = max(dp[i][j], dp[candidate[0]][candidate[1]]+1)
			} else {
				dp[i][j] = max(dp[i][j], dfs(candidate[0], candidate[1])+1)
			}
		}
		visited[i][j] = true
		return dp[i][j]
	}

	ans := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == false {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
