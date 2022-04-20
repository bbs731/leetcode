package leetcode

func countCornerRectangles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	counts := 0
	for i := 1; i < m; i++ {
		for c := 0; c < n; c++ {
			if grid[i][c] == 1 {
				for p := c + 1; p < n; p++ {
					if grid[i][p] == 1 {
						for r := 0; r < i; r++ {
							if grid[r][c] == 1 && grid[r][p] == 1 {
								counts++
							}
						}
					}

				}
			}
		}
	}
	return counts
}
