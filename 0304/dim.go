package leetcode

type NumMatrix struct {
	dp [][]int
}

/*
调整了一下  index 和以为的保持一致， 这样就不需要 getter function 来检查 index < 0 的情况了。
 */
func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])

	numMatrix := NumMatrix{
		dp: make([][]int, m+1),
	}

	for i := 0; i <= m; i++ {
		numMatrix.dp[i] = make([]int, n+1)
	}

	// 让我们来，创建二维的前缀数组和。 （可以参考 1314 那道题)
	//for i := 1; i <= m; i++ {
	//	for j := 1; j <= n; j++ {
	//		numMatrix.dp[i][j] = numMatrix.dp[i-1][j] + numMatrix.dp[i][j-1] - numMatrix.dp[i-1][j-1] + matrix[i-1][j-1]
	//	}
	//}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			numMatrix.dp[i+1][j+1] = numMatrix.dp[i][j+1] + numMatrix.dp[i+1][j] - numMatrix.dp[i][j] + matrix[i][j]
		}
	}
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return this.dp[row2+1][col2+1] - this.dp[row1][col2+1] - this.dp[row2+1][col1] + this.dp[row1][col1]
}
