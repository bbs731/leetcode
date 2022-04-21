package leetcode

/*
 O(m*n*n)
 */
func countCornerRectangles(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	counts := make(map[[2]int]int)
	for i := 0; i < m; i++ {
		for c := 0; c < n; c++ {
			if grid[i][c] == 1 {
				for p := c + 1; p < n; p++ {
					if grid[i][p] == 1 {
						counts[[2]int{c, p}] += 1
					}
				}
			}
		}
	}

	sum := 0
	for _, v := range counts {
		sum += v * (v - 1) / 2
	}
	return sum
}
