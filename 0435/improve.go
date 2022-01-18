package leetcode

import "sort"

/*
O(nlgn) solution
 */
func eraseOverlapIntervals(intervals [][]int) int {
	// 按照 end 来排序 intervals
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	// 这里面最关键的性质是， end 最小的 Interval 是最左边的 interval (这是一个 贪心的选择）

	ans := 0
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < current[1] {
			// i intersect with 最左边的 Interval
			ans++
		} else {
			current = intervals[i]
		}
	}
	return ans
}
