package leetcode

func flip(row []int) {
	for i := 0; i < len(row); i++ {
		if row[i] == 0 {
			row[i] = 1
		} else {
			row[i] = 0
		}
	}

}

func matrixScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			flip(grid[i])
		}
	}
	// then check column
	for i := 1; i < n; i++ {
		zeros := 0
		ones := 0
		for j := 0; j < m; j++ {
			if grid[j][i] == 0 {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			// flip column i
			for j := 0; j < m; j++ {
				if grid[j][i] == 0 {
					grid[j][i] = 1
				} else {
					grid[j][i] = 0
				}
			}
		}
	}

	// count rows and added all rows

	ans := 0
	for i := 0; i < m; i++ {
		row := 0
		for j := 0; j < n; j++ {
			row = row * 2
			row += grid[i][j]
		}
		ans += row
	}
	return ans
}
