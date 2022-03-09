package leetcode

func maxProfit(prices []int) int {
	// 初始化是难点啊！
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0 // buy2 的初始化，太难了， 怎么想到的？

	for i := 1; i < len(prices); i++ {
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, prices[i]+buy1)
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, prices[i]+buy2)
	}
	return max(sell1, sell2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
