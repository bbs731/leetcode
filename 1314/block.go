package leetcode

/*

这是一个， 前缀数组，扩展到了二维的情况的时候，的一个应用的例子。

用 dp[i][j] 来表示，矩阵中， element (i,j) 和 （0， 0）作为两个顶点围成的长方形 的所有 elements 的 sum 和。

dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i][j]

我们安装这个递推的式子， 把所有的 dp[i][j] 都计算出来，时间复杂度 O（mn).

然后再利用 dp 的值，来做本道题的查询

wow！  太牛了， 一次提交就过了！

 */
func matrixBlockSum(mat [][]int, k int) [][]int {

	m := len(mat)
	n := len(mat[0])

	dp := make([][]int, m)
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		ans[i] = make([]int, n)
	}

	getter := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return dp[i][j]
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = getter(i-1, j) + getter(i, j-1) - getter(i-1, j-1) + mat[i][j]
		}
	}

	// 现在我们有了二维数组的前缀和, 在 DP 里面

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = query(dp, i, j, k, m, n)
		}
	}

	return ans
}

func query(dp [][]int, i, j, k, m, n int) int {

	// left up corner
	lux := max(i-k, 0)
	luy := max(j-k, 0)

	// right down corner
	rdx := min(i+k, m-1)
	rdy := min(j+k, n-1)

	getter := func(i, j int) int {
		if i < 0 || j < 0 {
			return 0
		}
		return dp[i][j]
	}

	return dp[rdx][rdy] - getter(lux-1, rdy) - getter(rdx, luy-1) + getter(lux-1, luy-1)
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
