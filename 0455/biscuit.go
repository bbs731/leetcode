package leetcode

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	start := 0

	for i := 0; i < len(g); i++ {
		if start >= len(s) {
			break
		}
		for j := start; j < len(s); j++ {
			if s[j] >= g[i] {
				start = j + 1
				ans++
				break
			}
		}
	}
	return ans
}
