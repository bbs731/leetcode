package leetcode

import (
	"math"
)

func minCostII(costs [][]int) int {
	n := len(costs)
	k := len(costs[0])
	f := make([]int, k)
	g := make([]int, k)

	for i := 0; i < k; i++ {
		f[i] = costs[0][i]
	}

	for i := 1; i < n; i++ {
		copy(g, f)

		for j := 0; j < k; j++ {
			f[j] = update(j, k, g) + costs[i][j]
		}
	}

	return min(f...)
}

func update(color int, k int, g []int) int {

	ans := math.MaxInt64
	for i := 0; i < k; i++ {
		if i != color {
			ans = min(ans, g[i])
		}
	}

	return ans

}

func min(s ...int) int {
	ans := s[0]
	for i := 1; i < len(s); i++ {
		if ans > s[i] {
			ans = s[i]
		}
	}
	return ans
}
