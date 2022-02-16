package leetcode

func longestLine(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])

	var getter = func(i, j int) int {
		if i < 0 || i >= m {
			return 0
		}
		if j < 0 || j >= n {
			return 0
		}
		return mat[i][j]
	}

	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 4) // store 3 directions
		}
	}

	dirs := [][]int{[]int{0, -1}, []int{-1, 0}, []int{-1, -1}, []int{-1, 1}}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				// check 4 directions
				for k := 0; k < len(dirs); k++ {
					if getter(i+dirs[k][0], j+dirs[k][1]) == 1 {
						dp[i][j][k] = dp[i+dirs[k][0]][j+dirs[k][1]][k] + 1
					} else {
						dp[i][j][k] = 1
					}
					ans = max(ans, dp[i][j][k])
				}
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
