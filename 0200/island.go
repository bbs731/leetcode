package leetcode

/*
	写错了好几遍。
 */
var visited [][]bool

func DFS(grid [][]byte, i, j int) {
	visited[i][j] = true
	// for neighbours
	if i+1 < len(grid) && visited[i+1][j] == false && grid[i+1][j] == '1' {
		DFS(grid, i+1, j)
	}
	if j+1 < len(grid[0]) && visited[i][j+1] == false && grid[i][j+1] == '1' {
		DFS(grid, i, j+1)
	}
	if i-1 >= 0 && visited[i-1][j] == false && grid[i-1][j] == '1' {
		DFS(grid, i-1, j)
	}
	if j-1 >= 0 && visited[i][j-1] == false && grid[i][j-1] == '1' {
		DFS(grid, i, j-1)
	}
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	num := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == false && grid[i][j] == '1' {
				DFS(grid, i, j)
				num++
			}
		}
	}

	return num
}
