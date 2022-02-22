package leetcode

import "sort"

/*
greedy algorithms
 */
func findLongestChain(pairs [][]int) int {

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	current := 0
	cnt := 1
	for i := 1; i < len(pairs); i++ {
		if pairs[i][0] > pairs[current][1] {
			current = i
			cnt++
		}
	}
	return cnt
}
