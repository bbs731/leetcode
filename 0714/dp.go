package leetcode

/*
f[i][0] 第i天 有股票  =  max{ f[i-1][0], f[i-1][1] - prices[i] - fee}
f[i][1] 第i天 没有股票 = max { f[i-1][1], f[i-1][0] + prices[i] }
*/

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0] - fee
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return max(dp[n-1][0], dp[n-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
