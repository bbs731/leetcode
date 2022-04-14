package leetcode


/* wrong answer
休息一下吧!
 */
func removeBoxes(boxes []int) int {
	n := len(boxes)

	// counts[num][i][j],  range [i, j] 之间 , num 出现的次数
	counts := make([][][]int, 101)
	// dp[num][i][j] ,  range [i, j] 之间， 最后以 num merge 结尾获得的最大分数
	dp := make([][][]int, 101)

	for i := 0; i <= 100; i++ {
		counts[i] = make([][]int, n)
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			counts[i][j] = make([]int, n)
			dp[i][j] = make([]int, n)
		}
	}

	// 初始化
	for i := 0; i < n; i++ {
		counts[boxes[i]][i][i] = 1
		dp[boxes[i]][i][i] = 1
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			for k := 1; k <= 100; k++ {
				if boxes[i] == k {
					counts[k][i][j] = 1 + counts[k][i+1][j]
				} else {
					counts[k][i][j] = counts[k][i+1][j]
				}
			}
		}
	}

	for l := 2; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			numi, numj := boxes[i], boxes[j]
			for k := 1; k <= 100; k++ {
				//dp[k][i][j] == ?
				if k == numi && k == numj {
					row := counts[k][i][j]
					dp[k][i][j] = dp[k][i+1][j-1] + row*row
				} else if k == numi {
					row := counts[k][i][j]
					dp[k][i][j] = dp[k][i+1][j] + row*row
				} else if k == numj {
					row := counts[k][i][j]
					dp[k][i][j] = dp[k][i][j-1] + row*row
				} else {
					if counts[k][i][j] == 0 {
						dp[k][i][j] = 0
					} else {
						dp[k][i][j] = dp[k][i+1][j-1] + 1 + 1
					}
				}
			}
		}
	}

	ans := 0

	for k := 1; k <= 100; k++ {
		ans = max(ans, dp[k][0][n-1])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
