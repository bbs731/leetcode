package leetcode

func maximalSquare(matrix [][]byte) int {
	M := len(matrix)
	N := len(matrix[0])

	ans := 0
	// dp[i][j] 代表， 坐标 i, j 的方格作为正方形的，右下角的正方形的长度.
	dp := make([][]int, M)
	for i := 0; i < M; i++ {
		dp[i] = make([]int, N)
	}

	//  初始化， first row and first column
	for i := 0; i < M; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			ans = 1
		}
	}
	for j := 0; j < N; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			ans = 1
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = 1
			ans = max(ans, 1)
			if i-1 >= 0 && j-1 >= 0 {
				// dp[i][j]  和  dp[i-1][j], dp[i][j-1] 还有 dp[i-1][j-1] 有关系， 但是，找到这个关系，太难了！ 哎！
				// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1  是这样吗？ 这个竟然是对的！ shit a!
				// 被，这个转态方程， 折磨的要死啊！
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				ans = max(ans, dp[i][j])
			}
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
