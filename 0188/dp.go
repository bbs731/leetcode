package leetcode

import "math"

//buy[i][j]=max{buy[i−1][j],sell[i−1][j]−price[i]}
//sell[i][j]=max{sell[i−1][j],buy[i−1][j−1]+price[i]}
func maxProfit(k int, prices []int) int {
	n := len(prices)

	buy := make([][]int, n)
	sell := make([][]int, n)

	for i := 0; i < n; i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}

	// 初始化
	buy[0][0] = -prices[0]
	sell[0][0] = 0
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}

	ans := 0
	for j := 0; j <= k; j++ {
		ans = max(ans, sell[n-1][j])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
