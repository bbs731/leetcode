package leetcode

func maxProfit(prices []int) int {
	N := len(prices)
	profit := 0

	if len(prices) == 0 {
		return 0
	}

	// prepare dp
	less := prices[0]
	dp := make([]int, N)

	for i := 1; i < len(prices); i++ {
		if prices[i] > less {
			dp[i] = max(dp[i-1], prices[i]-less)
		} else {
			less = prices[i]
			dp[i] = dp[i-1]
		}
	}

	// now prepare dp2

	dp2 := make([]int, N)
	big := prices[N-1]
	for i := N - 2; i >= 0; i-- {
		if prices[i] >= big {
			big = prices[i]
			dp2[i] = dp2[i+1]
		} else {
			dp2[i] = max(dp2[i+1], big-prices[i])
		}

	}

	// now we have N-1 choice to break the problem into 2 subproblem

	for i := 0; i < N-1; i++ {
		prices1 := dp[i]
		prices2 := dp2[i+1]
		profit = max(profit, prices1+prices2)
	}
	return max(profit, dp[N-1])
}

func min(a, b int) int {
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
