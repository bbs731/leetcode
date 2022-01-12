package leetcode

import "math"

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	dp[0] = 0

	getter := func(i int) int {
		if i < 0 {
			return math.MaxInt64
		}
		return dp[i]
	}
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt64
		for j := 0; j < len(coins); j++ {
			if getter(i-coins[j]) != math.MaxInt64 {
				dp[i] = min(dp[i], getter(i-coins[j])+1)
			}
		}
	}

	if dp[amount] == math.MaxInt64 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
