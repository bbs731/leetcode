package leetcode

/*
A DP problem
f[i][j][k] = sum (f[i'][j'][k-1])
(i', j') 是上下左右的邻居。  [k] 的维度，用不着，所以 f 可以是一个二维的数组 [][]int
 */
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {

	if maxMove == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化 border
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] += 1
			}
			if i == m-1 {
				dp[i][j] += 1
			}
			if j == 0 {
				dp[i][j] += 1
			}
			if j == n-1 {
				dp[i][j] += 1
			}
		}
	}
	var cdp [][]int
	var getter = func(i, j int) int {
		if i < 0 || i >= m {
			return 1
		}
		if j < 0 || j >= n {
			return 1
		}
		return cdp[i][j]
	}

	dirs := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	for k := 2; k <= maxMove; k++ {
		cdp = makeAcopy(dp, m, n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				dp[i][j] = 0
				for d := 0; d < len(dirs); d++ {
					dp[i][j] += getter(dirs[d][0]+i, dirs[d][1]+j)
				}
			}
		}
	}

	return dp[startRow][startColumn] % (1000000000 + 7)
}

func makeAcopy(dp [][]int, m, n int) [][]int {
	cdp := make([][]int, m)
	for i := 0; i < m; i++ {
		cdp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cdp[i][j] = dp[i][j] % (1000000000 + 7)
		}
	}
	return cdp
}
