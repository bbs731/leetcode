package leetcode

/*
	f[i][0]  第i 天结束之后， 有股票：  f[i][0] = max{ f[i-1][0] , f[i-1][1] - price[i]}
	f[i][1]  第i 天结束之后， 没有股票，但是可以买：  f[i][1] = max { f[i-1][1], f[i-1][2]}
	f[i][2]  第i 天结束之后， 没有股票，不可以买：  f[i][2] = f[i-1][0] + price[i]
 */
func maxProfit(prices []int) int {

	// 初始化 f[0][0], f[0][1], f[0][2], 为 a, b ,c
	a, b, c := -prices[0], 0, 0

	for i := 1; i < len(prices); i++ {
		preva := a
		a = max(a, b-prices[i])
		b = max(b, c)
		c = preva + prices[i]
	}

	return max(a, max(b, c))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
