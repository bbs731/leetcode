package leetcode

import "sort"

/*
O(n^2) 的解法， 会超出时间限制
 */
func maxEnvelopes(envelopes [][]int) int {

	ans := 0
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		} else if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return false
	})

	check := func(m, n int) bool {
		if envelopes[m][0] == envelopes[n][0] {
			return false
		}
		if envelopes[m][1] >= envelopes[n][1] {
			return false
		}
		return true
	}

	cnt := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		cnt[i] = 1

		for j := i - 1; j >= 0; j-- {
			if check(j, i) {
				cnt[i] = max(cnt[i], cnt[j]+1)
			}
		}
		ans = max(ans, cnt[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
