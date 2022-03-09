package leetcode

const INFINITY = 10000000

func minCost(houses []int, cost [][]int, m int, n int, target int) int {

	// dp[i][b][c] // b stands for block, c stands for color, i stands for house index
	dp := make([][][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([][]int, m+1)
		for b := 0; b <= m; b++ {
			dp[i][b] = make([]int, n+1)
			for c := 0; c <= n; c++ {
				dp[i][b][c] = INFINITY
			}
		}
	}
	// 初始化 house 0
	for c := 1; c <= n; c++ {
		if houses[0] == 0 {
			dp[0][1][c] = cost[0][c-1]
		} else {
			dp[0][1][houses[0]] = 0 // 去年夏天已经涂过了，不用钱
		}
	}

	for i := 1; i < m; i++ {
		// house i select the color
		var colors []int
		// 第 ith 的 house cost 需要些处理， 如果去年夏天刷过，cost 为0， 如果没刷过 house[i]==0，才计算 cost
		costi := make([]int, n+1)
		if houses[i] == 0 {
			//colors = make([]int, n+1)
			for k := 1; k <= n; k++ {
				//colors[k] = k
				colors = append(colors, k)
				costi[k-1] = cost[i][k-1]
			}
		} else {
			colors = append(colors, houses[i])
		}
		// loop the color
		for _, c := range colors {
			// loop the block number
			for b := i + 1; b >= 1; b-- {
				// we can say dp[i] only rely on dp[i-1]
				j := i - 1
				if houses[j] == 0 {
					for jc := 1; jc <= n; jc++ {
						if jc == c {
							//dp[i][b][c] = min(dp[i][b][c], dp[i-1][b][c]+cost[i][c-1])
							dp[i][b][c] = min(dp[i][b][c], dp[i-1][b][c]+costi[c-1])
						} else {
							//dp[i][b][c] = min(dp[i][b][c], dp[i-1][b-1][jc]+cost[i][c-1])
							dp[i][b][c] = min(dp[i][b][c], dp[i-1][b-1][jc]+costi[c-1])
						}
					}

				} else {
					if houses[j] == c {
						//dp[i][b][c] = min(dp[i][b][c], dp[i-1][b][c]+cost[i][c-1])
						dp[i][b][c] = min(dp[i][b][c], dp[i-1][b][c]+costi[c-1])
					} else {
						//dp[i][b][c] = min(dp[i][b][c], dp[i-1][b-1][houses[j]]+cost[i][c-1])
						dp[i][b][c] = min(dp[i][b][c], dp[i-1][b-1][houses[j]]+costi[c-1])
					}
				}
			}
		}
	}

	ans := INFINITY
	for c := 1; c <= n; c++ {
		ans = min(ans, dp[m-1][target][c])
	}
	if ans == INFINITY {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
