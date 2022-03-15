package leetcode

import (
	"sort"
)

/*
只能说 backtrace 简直太恶心了！ 怎么办呢？

需要深刻的理解， dfs 和 DP 的区别。 什么区别呢？

 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	n := len(candidates)
	ans := make([][]int, 0)
	seen := make(map[int]bool)
	visited := make(map[int]bool)

	var dfs func(sum int, start int, prefix []int)

	dfs = func(sum int, start int, prefix []int) {
		if start == n {
			return
		}
		h := calculatehash(prefix)
		if _, ok := visited[sum*1000*1000+start*1000+h]; ok {
			// visited before
			return
		}
		if sum+candidates[start] > target {
			return
		}
		if sum+candidates[start] == target {
			p := make([]int, 0, len(prefix)+1)
			p = append(p, prefix...)
			p = append(p, candidates[start])
			h := calculatehash(p)
			if _, ok := seen[h]; !ok {
				ans = append(ans, p)
			}
			seen[h] = true

			return
		}

		cprefix := make([]int, 0, len(prefix)+1)
		cprefix = append(cprefix, prefix...)
		cprefix = append(cprefix, candidates[start])
		dfs(sum+candidates[start], start+1, cprefix)
		// restore
		dfs(sum, start+1, prefix)

		visited[sum*1000*1000+start*1000+h] = true
	}

	dfs(0, 0, []int{})
	return ans
}

func calculatehash(p []int) int {
	n := len(p)
	base := 53
	h := 0
	for i := 0; i < n; i++ {
		h = h*base + p[i]
	}
	return h
}
