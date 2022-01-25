package leetcode

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(nums); i += 2 {
		ans += min(nums[i], nums[i+1])
	}

	return ans
}
