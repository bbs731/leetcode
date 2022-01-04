package leetcode

/*

f[i][0] 第i天 有股票  =  max{ f[i-1][0], f[i-1][1] - prices[i] - fee}
f[i][1] 第i天 没有股票 = max { f[i-1][1], f[i-1][0] + prices[i] }

 */
func maxProfit(prices []int, fee int) int {
	a, b := -prices[0]-fee, 0

	for i := 1; i < len(prices); i++ {
		preva := a
		a = max(a, b-prices[i]-fee)
		b = max(b, preva+prices[i])
	}
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
