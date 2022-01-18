package leetcode

import (
	"sort"
)

/*
时间复杂度是 O(n^2) 会超时
 */
func eraseOverlapIntervals(intervals [][]int) int {

	// less function 比较， start 还是 end 是没有区别的， 这里sort 只为了，能固定 intervals 之间的顺序
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })

	//DP 记录， 以 i interval 结尾的， 能组成最多不想交 intervals 的个数
	dp := make([]int, len(intervals))

	for i := 1; i < len(intervals); i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				// i and j not intersect
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	largest := dp[0]
	for i := 1; i < len(intervals); i++ {
		if largest > dp[i] {
			largest = dp[i]
		}
	}

	return len(intervals) - largest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
