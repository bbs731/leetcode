package leetcode

import (
	"fmt"
	"math"
)

// dp[i][k][0] // 0 stands for ith day finished does not hold stocks, 1 is when ith day is finished hold stocks
//dp[i][j][1] = max( dp[i-1][j][1], dp[i-1][j-1]][0] - prices[i])
//dp[i][j][0] = max( dp[i-1][j][0], dp[i-1][j][1] +price[i])

func maxProfit(k int, prices []int) int {
	n := len(prices)

	if n <= 1 {
		return 0
	}
	dp := make([][][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	// 初始化
	dp[0][0][0] = 0
	dp[0][0][1] = -prices[0] // buy

	for i := 1; i <= k; i++ {
		dp[0][i][0] = math.MinInt64 / 2
		dp[0][i][1] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		// j = 0 // special treatment
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]-prices[i]) // for buy
		for j := 1; j <= k; j++ {
			//buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			//sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]-prices[i])
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i]) // sell // 真他妈的神奇啊， sell 的 j 的 index 要大于 buy 的 index， 这样设计，答案才是真确的。 // 其实想想，也是
			// 挺自然的， 因为 dp[0][0][0] 已经占了 0th 交易的 sell index, 所以 sell index 需要增加

		}
	}
	fmt.Println(dp)
	ans := 0
	for j := 0; j <= k; j++ {
		ans = max(ans, dp[n-1][j][0])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
