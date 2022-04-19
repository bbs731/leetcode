package leetcode

import "sort"

/*
参考 1235  similar DP problem
 */

func maxValue(events [][]int, k int) int {

	n := len(events)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}

	sort.Slice(events, func(i int, j int) bool {
		return events[i][1] < events[j][1]
	})

	for i := 1; i <= n; i++ {
		pos := sort.Search(n, func(mid int) bool {
			return events[mid][1] > events[i-1][0]-1 // 看题， [1, 1] start time = end time = 1 interval
		})
		for q := 1; q <= k && q <= i; q++ {
			dp[i][q] = max(dp[i-1][q], dp[pos][q-1]+events[i-1][2])
		}
	}
	return dp[n][k]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
