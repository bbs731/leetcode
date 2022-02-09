package leetcode

import "sort"

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		pos := sort.SearchInts(dp, nums[i])
		if pos == len(dp) {
			dp = append(dp, nums[i])

		} else {
			// pos < len(dp)
			dp[pos] = nums[i]
		}
	}
	return len(dp)
}
