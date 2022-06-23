package leetcode

/*
我觉得这个题是难题！
其中 dp[i][m] m 这个index 用到了 max(j, m) 还是挺特别的，在别的例子中没见到过！
 */
func stoneGameII(piles []int) int {
	n := len(piles)
	prefixsum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixsum[i] = prefixsum[i-1] + piles[i-1]
	}

	dp := make([][]int, n) // state dp[i][m]
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for k := 1; k <= n; k++ {
			if i+2*k >= n {
				dp[i][k] = prefixsum[n] - prefixsum[i]
			} else {
				value := 10000000
				for j := 1; j <= 2*k; j++ {
					// we have these candidates, and we need to find the mininum of them
					if dp[i+j][max(j, k)] < value {
						value = dp[i+j][max(j, k)]
					}
				}
				dp[i][k] = prefixsum[n] - prefixsum[i] - value
			}
		}
	}
	return dp[0][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
