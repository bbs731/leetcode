package leetcode

func maxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var getter = func(r, c int) bool {
		if r < 0 || r >= m || c < 0 || c >= n {
			return true
		}
		return visited[r][c]
	}
	var dfs func(i, j int) int
	dfs = func(r, c int) int {
		visited[r][c] = true
		total := 1

		dir := [][2]int{[2]int{0, -1}, [2]int{0, 1}, [2]int{1, 0}, [2]int{-1, 0}}
		for i := 0; i < len(dir); i++ {
			a := r + dir[i][0]
			b := c + dir[i][1]
			if getter(a, b) == false && grid[a][b] == 1 {
				total += dfs(a, b)
			}
		}
		return total
	}

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == false && grid[i][j] == 1 {
				total = max(total, dfs(i, j))
			}
		}
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
