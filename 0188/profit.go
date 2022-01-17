package leetcode

/*
	因为需要 loop prev 数组， 所以，时间复杂度， 比官网的略高
	O (nk) 的解， 看官网吧!
 */
func maxProfit(k int, prices []int) int {
	N := len(prices)

	if N == 0 {
		return 0
	}

	// dp[i][k] 代表， k 此交易， 的最大收益值
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, k+1)
	}

	getter := func(i, k int) int {
		if i <= 0 || k <= 0 {
			return 0
		}
		return dp[i][k]
	}
	// dp[i][k] = max { dp[i][k-1] (这个保证了 dp[i][k] 是递增函数, dp[i-1][k] }
	// 另外一种情况就是 prices[i] 参与了一次股票交易， 我们用 prev[i] 数组，表示， 比 prices[i] 小的买入点的index.

	prev := make([]int, N)
	prev[0] = -1

	// 初始化 prev
	for i := 1; i < N; i++ {
		previous := i - 1
		for previous != -1 {
			if prices[i] > prices[previous] {
				break
			} else {
				previous = prev[previous]
			}
		}
		prev[i] = previous
	}

	for j := 1; j <= k; j++ {
		for i := 1; i < N; i++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j])

			// prices[i] 参与交易
			previous := prev[i]
			profit := 0
			for previous != -1 {
				profit = prices[i] - prices[previous]
				//dp[i][j] = max(dp[i][j], profit+dp[previous-1][j-1])
				dp[i][j] = max(dp[i][j], profit+getter(previous-1, j-1))
				previous = prev[previous]
			}
		}
	}

	return dp[N-1][k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
