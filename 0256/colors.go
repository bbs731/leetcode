package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(costs [][]int) int {
	n := len(costs)
	f := [3]int{costs[0][0], costs[0][1], costs[0][2]}

	for i := 1; i < n; i++ {
		f0, f1, f2 := f[0], f[1], f[2]
		f[0] = min(f1, f2) + costs[i][0]
		f[1] = min(f0, f2) + costs[i][1]
		f[2] = min(f0, f1) + costs[i][2]
	}

	return min(min(f[0], f[1]), f[2])
}
