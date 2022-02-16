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

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				// check 4 directions
				// check row
				if getter(i, j-1) == 1 {
					dp[i][j][0] = dp[i][j-1][0] + 1
				} else {
					dp[i][j][0] = 1
				}
				// column
				if getter(i-1, j) == 1 {
					dp[i][j][1] = dp[i-1][j][1] + 1
				} else {
					dp[i][j][1] = 1
				}
				// up left
				if getter(i-1, j-1) == 1 {
					dp[i][j][2] = dp[i-1][j-1][2] + 1
				} else {
					dp[i][j][2] = 1
				}
				// up right
				if getter(i-1, j+1) == 1 {
					dp[i][j][3] = dp[i-1][j+1][3] + 1
				} else {
					dp[i][j][3] = 1
				}

				ans = max(ans, dp[i][j][0], dp[i][j][1], dp[i][j][2], dp[i][j][3])
			}
		}
	}
	return ans
}

func max(nums ...int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
