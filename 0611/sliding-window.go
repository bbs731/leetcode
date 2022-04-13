package leetcode

import "sort"

func triangleNumber(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	counts := 0

	for a := 0; a < n-2; a++ {
		for b := a + 1; b < n-1; b++ {
			var c int
			for c = b + 1; c < n && nums[a]+nums[b] > nums[c]; c++ {
			}

			counts += max(c-1-b-1+1, 0)
		}
	}
	return counts
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
