package leetcode

func closedIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(int, int) bool
	dfs = func(row, col int) bool {
		visited[row][col] = true
		// reach border
		border := false
		if row == 0 || row == m-1 || col == 0 || col == n-1 {
			border = true
		}

		var getter = func(r, c int) bool {
			if r < 0 || r >= m || c < 0 || c >= n {
				return true
			}
			return visited[r][c]
		}
		dir := [][]int{[]int{0, 1}, []int{0, -1}, []int{1, 0}, []int{-1, 0}}
		for i := 0; i < len(dir); i++ {
			a := dir[i][0] + row
			b := dir[i][1] + col
			if getter(a, b) == false && grid[a][b] == 0 {
				if dfs(a, b) == true {
					border = true
				}
			}
		}
		return border
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == false && grid[i][j] == 0 {
				if dfs(i, j) == false {
					ans++
				}
			}
		}
	}

	return ans
}
