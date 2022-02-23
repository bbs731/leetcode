package leetcode

func mctFromLeafValues(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	largest := make([][]int, n)
	sum := 0
	// 初始化
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		largest[i] = make([]int, n)
		dp[i][i] = arr[i]
		largest[i][i] = arr[i]
		sum += arr[i]
	}
	//for i := 0; i < n; i++ {
	//	for j := i + 1; j < n; j++ {
	//		largest[i][j] = max(largest[i][j-1], arr[j])
	//	}
	//}
	for l := 1; l < n; l++ {
		for i := 0; i < n; i++ {
			// j -i = l
			j := l + i
			for k := i; k < j && j < n; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+largest[i][k]*largest[k+1][j])
				largest[i][j] = max(largest[i][k], largest[k+1][j]) // 其实 largest[i][j] 计算一遍就可以了，不用计算k 遍
			}
		}
	}

	return dp[0][n-1] - sum
}

func min(a, b int) int {
	if a == 0 {
		return b
	}
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
