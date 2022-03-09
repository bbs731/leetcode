package leetcode

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	// dp[i][0] 代表第 i 天结束之后，没有持有股票时的利润。  dp[i][1] 则表示 ith天结束之后，持有股票是的利润
	// 我们知道 dp[i] 只和 dp[i-1] 的状态有关系
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	// 初始化
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// 这个状态转移函数， 简直是太牛B了
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}

	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
