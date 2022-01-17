package leetcode

import "math"

/* 这个版本，是基础版， 时间复杂度是 O(n^2) 在 n = 10^5 的情况下， 会超时，
需要做优化
1. 计算 prices1 的时候， 可以保存个数组，dp[i] 表示 i 之前的一次交易的最大值，这样的话，计算prices1 就是一个简单的查询
2. 同样的道理， 保存 dp2[i] 的数组， 表示  i 到 N-1 之间交易的最大值， 这样计算 prices2 就变成一个简单的查询
*/

// 这个函数，只考虑一次交易
func unit(prices []int) int {
	less := math.MaxInt64
	profit := 0

	if len(prices) == 0 {
		return profit
	}

	for i := 0; i < len(prices); i++ {
		if prices[i] > less {
			profit = max(profit, prices[i]-less)
		} else {
			less = prices[i]
		}
	}
	return profit
}

func maxProfit(prices []int) int {
	N := len(prices)
	// now we have N-1 choice to break the problem into 2 subproblem

	profit := 0

	for i := 1; i <= N-1; i++ {
		prices1 := unit(prices[:i+1])
		prices2 := unit(prices[i:])

		profit = max(profit, prices1+prices2)

	}
	return profit
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
