package leetcode

/*
dfs 总是让你怀疑人生！ 
 */
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	m := len(grid1)
	n := len(grid1[0])

	visited := make([][]bool, m)
	for k := 0; k < m; k++ {
		visited[k] = make([]bool, n)
	}

	var getter = func(a, b int) bool {
		if a < 0 || a >= m || b < 0 || b >= n {
			return true
		}
		return visited[a][b]
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		visited[i][j] = true
		check := true
		if grid1[i][j] == 0 {
			check = false
		}

		direction := [][2]int{[2]int{0, -1}, [2]int{0, 1}, [2]int{1, 0}, [2]int{-1, 0}}
		for k := 0; k < len(direction); k++ {
			a, b := direction[k][0]+i, direction[k][1]+j
			if getter(a, b) == false && grid2[a][b] == 1 {
				if dfs(a, b) == false {
					check = false
				}
			}
		}
		return check
	}

	components := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 && visited[i][j] == false {
				if dfs(i, j) {
					components++
				}
			}
		}
	}

	return components
}
