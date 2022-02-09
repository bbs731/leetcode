package leetcode

import "sort"

/*
太难了， 想不到啊！
 */
func maxEnvelopes(envelopes [][]int) int {

	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return false
	})

	dp := make([]int, 0)
	for i := 0; i < len(envelopes); i++ {
		height := envelopes[i][1]
		pos := sort.SearchInts(dp, height)
		if pos < len(dp) {
			dp[pos] = height
		} else {
			dp = append(dp, height)
		}
	}
	return len(dp)
}
