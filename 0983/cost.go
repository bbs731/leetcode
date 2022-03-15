package leetcode

import "sort"

func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n)
	dp[0] = min(min(costs[0], costs[1]), costs[2]) // 这个 test case 太坏了！

	for i := 1; i < n; i++ {
		d := days[i]
		dp[i] = dp[i-1] + costs[0]
		l7 := d - 7 + 1
		l7i := sort.SearchInts(days, l7)
		if l7i == 0 {
			dp[i] = min(dp[i], costs[1])
		} else {
			dp[i] = min(dp[l7i-1]+costs[1], dp[i])
		}
		l30 := d - 30 + 1
		l30i := sort.SearchInts(days, l30)
		if l30i == 0 {
			dp[i] = min(dp[i], costs[2])
		} else {
			dp[i] = min(dp[l30i-1]+costs[2], dp[i])
		}
	}
	return dp[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
