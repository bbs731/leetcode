package leetcode

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])

	numMatrix := NumMatrix{
		dp: make([][]int, m),
	}

	for i := 0; i < m; i++ {
		numMatrix.dp[i] = make([]int, n)
	}

	// 让我们来，创建二维的前缀数组和。 （可以参考 1314 那道题)

	getter := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return numMatrix.dp[i][j]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			numMatrix.dp[i][j] = getter(i-1, j) + getter(i, j-1) - getter(i-1, j-1) + matrix[i][j]
		}
	}

	return numMatrix

}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	getter := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return this.dp[i][j]
	}

	return this.dp[row2][col2] - getter(row1-1, col2) - getter(row2, col1-1) + getter(row1-1, col1-1)
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
