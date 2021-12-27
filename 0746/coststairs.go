package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	a, b, c := 0, 0, 0
	for i := 2; i < len(cost); i++ {
		c = min(a+cost[i-2], b+cost[i-1])
		a, b = b, c
	}

	return min(a+cost[n-2], b+cost[n-1])
}
