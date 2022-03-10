package leetcode

/*
	f[i][0]  第i 天结束之后， 有股票：  f[i][0] = max{ f[i-1][0] , f[i-1][1] - price[i]}
	f[i][1]  第i 天结束之后， 没有股票，但是可以买：  f[i][1] = max { f[i-1][1], f[i-1][2]}
	f[i][2]  第i 天结束之后， 没有股票，不可以买：  f[i][2] = f[i-1][0] + price[i]
 */
func maxProfit(prices []int) int {
	n := len(prices)

	f := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		f[i] = make([]int, 3)
	}
	f[0][0], f[0][1], f[0][2] = -prices[0], 0, 0

	for i := 1; i < len(prices); i++ {
		f[i][0] = max(f[i-1][0], f[i-1][1]-prices[i])
		f[i][1] = max(f[i-1][1], f[i-1][2])
		f[i][2] = f[i-1][0] + prices[i]
	}
	// 初始化 f[0][0], f[0][1], f[0][2], 为 a, b ,c

	return max(f[n-1][2], max(f[n-1][1], f[n-1][0]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
