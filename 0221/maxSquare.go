package leetcode
/*

我的答案，要比别人的慢：

原因在于：
				for ; x >= i-l; x-- {
					if matrix[x][j] == '0' {
						break
					}
				}
				for ; y >= j-l; y-- {
					if matrix[i][y] == '0' {
						break
					}
				}
				dp[i][j] = min(i-x, j-y)


漏算的这些处理， 还是个循环， 可以用 dp[i-1][j] 和 dp[i][j-1] 等价的代替掉

 */
func maximalSquare(matrix [][]byte) int {
	M := len(matrix)
	N := len(matrix[0])
	ans := 0

	dp := make([][]int, M)
	for i := 0; i < M; i++ {
		dp[i] = make([]int, N)
	}

	getter := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return dp[i][j]
	}

	checker := func(i, j, l int) bool {
		// check column
		for x := i - l; x <= i; x++ {
			if matrix[x][j] != '1' {
				return false
			}
		}

		for y := j - l; y <= j; y++ {
			if matrix[i][y] != '1' {
				return false
			}
		}
		return true
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			// matrix[i][j] == '1'
			if getter(i-1, j-1) != 0 {
				l := getter(i-1, j-1)
				if checker(i, j, l) {
					dp[i][j] = l + 1
					ans = max(ans, dp[i][j])
					continue
				}
				//checker is false
				// 这里漏算了
				// check colmn and rows which one is shorter
				x, y := i, j
				for ; x >= i-l; x-- {
					if matrix[x][j] == '0' {
						break
					}
				}
				for ; y >= j-l; y-- {
					if matrix[i][y] == '0' {
						break
					}
				}
				dp[i][j] = min(i-x, j-y)

				//if i-1 >= 0 && j-1 >= 0 && matrix[i-1][j] == '1' && matrix[i][j-1] == '1' {
				//	dp[i][j] = 2
				//} else {
				//	dp[i][j] = 1
				//}
			} else {
				dp[i][j] = 1
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans * ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
