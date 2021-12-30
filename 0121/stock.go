package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	ans := 0
	N := len(prices)
	min := prices[0]

	for i := 1; i < N; i++ {
		if prices[i] < min {
			min = prices[i]
		} else {
			ans = max(ans, prices[i]-min)
		}
	}
	return ans
}
