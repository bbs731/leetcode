package leetcode
/* wrong answer
*/
func removeBoxes(boxes []int) int {
	n := len(boxes)
	dp := make([][]int, n)

	// 初始化
	for i := 0; i < n; i++ {
		//counts[boxes[i]][i][i] = 1
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	getter := func(i, j int) int {
		if i >= n || j >= n {
			return 0
		}
		return dp[i][j]
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			// 只有两种情况, 合并 boxes[i] 和不合并 boxes[i]
			// boxes[i] 和   boxes[i+1 : j] 之间相同的数合并 (合并的话，有多中情况，是和前 1，2，3...个合并） 或者不合并
			numi := boxes[i]
			dp[i][j] = 1 + dp[i+1][j] // 不和别人合并
			//if numi == boxes[j] {
			//	dp[i][j]= max(dp[i][j], 2*2 + dp[i+1][j-1])
			//}
			accumulates := 0
			cnts := 1
			last := i
			for k := i + 1; k <= j; k++ {
				if boxes[k] == numi {
					cnts++
					accumulates += dp[last+1][k-1]
					last = k
					dp[i][j] = max(dp[i][j], cnts*cnts+accumulates+getter(k+1, j))
					dp[i][j] = max(dp[i][j], 2*2+dp[i+1][k-1]+getter(k+1, j))

				}
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
