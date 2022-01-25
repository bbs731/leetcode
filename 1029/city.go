package leetcode

import "sort"

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
func twoCitySchedCost(costs [][]int) int {
	ans := 0
	sort.Slice(costs, func(i, j int) bool {
		return (max(costs[i][0], costs[i][1]) - min(costs[i][0], costs[i][1])) >  // 这里要倒序啊！
			(max(costs[j][0], costs[j][1]) - min(costs[j][0], costs[j][1]))
	})

	cntA := 0
	cntB := 0
	n := len(costs)
	for i := 0; i < len(costs); i++ {
		if cntA == n/2 {
			ans += costs[i][1]
			cntB++
			continue
		}
		if cntB == n/2 {
			ans += costs[i][0]
			cntA++
			continue
		}
		if costs[i][0] > costs[i][1] {
			ans += costs[i][1]
			cntB++

		} else {
			ans += costs[i][0]
			cntA++
		}
	}
	return ans
}
